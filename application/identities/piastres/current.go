package piastres

import (
	"math/rand"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
)

type current struct {
	hashAdapter                                  hash.Adapter
	pkFactory                                    signature.PrivateKeyFactory
	lockBuilder                                  locks.Builder
	expenseBuilder                               expenses.Builder
	expenseContentBuilder                        expenses.ContentBuilder
	trxBuilder                                   transactions.Builder
	identityRepository                           identities.Repository
	identityService                              identities.Service
	name                                         string
	password                                     string
	seed                                         string
	amountAdditionalPubKeysPerShareHolderPerRing int
}

func createCurrent(
	hashAdapter hash.Adapter,
	pkFactory signature.PrivateKeyFactory,
	lockBuilder locks.Builder,
	expenseBuilder expenses.Builder,
	expenseContentBuilder expenses.ContentBuilder,
	trxBuilder transactions.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
	name string,
	password string,
	seed string,
	amountAdditionalPubKeysPerShareHolderPerRing int,
) Current {
	out := current{
		hashAdapter:           hashAdapter,
		pkFactory:             pkFactory,
		lockBuilder:           lockBuilder,
		expenseBuilder:        expenseBuilder,
		expenseContentBuilder: expenseContentBuilder,
		trxBuilder:            trxBuilder,
		identityRepository:    identityRepository,
		identityService:       identityService,
		name:                  name,
		password:              password,
		seed:                  seed,
		amountAdditionalPubKeysPerShareHolderPerRing: amountAdditionalPubKeysPerShareHolderPerRing,
	}

	return &out
}

// Bucket executes a bucket transaction
func (app *current) Bucket(absolutePath string, fees []Fee) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	bucket, err := identity.Buckets().Fetch(absolutePath)
	if err != nil {
		return err
	}

	bucketHash := bucket.Hash()
	createdOn := time.Now().UTC()
	builder := app.trxBuilder.Create().WithBucket(bucketHash).CreatedOn(createdOn)
	if len(fees) > 0 {
		expFees := []expenses.Expense{}
		for _, oneFee := range fees {
			amount := oneFee.Amount()
			feeLock := oneFee.Lock()
			walletBills, err := identity.Wallets().Fetch(amount)
			if err != nil {
				return err
			}

			bills := []bills.Bill{}
			ringPublicKeys := []signature.PublicKey{}

			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			maxPubKeysInRing := r1.Intn(app.amountAdditionalPubKeysPerShareHolderPerRing)
			maxIndex := int(maxPubKeysInRing / len(walletBills))
			for _, oneWalletBill := range walletBills {
				pubKeys := oneWalletBill.RingKeys()
				for index, onePubKey := range pubKeys {
					if maxIndex >= (index + 1) {
						break
					}

					ringPublicKeys = append(ringPublicKeys, onePubKey)
				}

				bill := oneWalletBill.Bill()
				bills = append(bills, bill)
			}

			remainingPK := app.pkFactory.Create()
			ringPublicKeys = append(ringPublicKeys, remainingPK.PublicKey())

			// shuffle the slice of ring public keys:
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(ringPublicKeys), func(i, j int) {
				ringPublicKeys[i], ringPublicKeys[j] = ringPublicKeys[j], ringPublicKeys[i]
			})

			lockPublicKeys := []hash.Hash{}
			for _, oneRingPublicKey := range ringPublicKeys {
				hsh, err := app.hashAdapter.FromString(oneRingPublicKey.String())
				if err != nil {
					return err
				}

				lockPublicKeys = append(lockPublicKeys, *hsh)
			}

			lockCreatedOn := time.Now().UTC()
			remaining, err := app.lockBuilder.Create().WithPublicKeys(lockPublicKeys).CreatedOn(lockCreatedOn).Now()
			if err != nil {
				return err
			}

			createdOn := time.Now().UTC()
			expenseContent, err := app.expenseContentBuilder.Create().WithAmount(amount).From(bills).WithLock(feeLock).WithRemaining(remaining).CreatedOn(createdOn).Now()
			if err != nil {
				return err
			}

			msg := expenseContent.Hash().String()
			signatures := []signature.RingSignature{}
			for _, oneWalletBill := range walletBills {
				pk := oneWalletBill.PrivateKey()
				pubKeys := oneWalletBill.RingKeys()
				ringSig, err := pk.RingSign(msg, pubKeys)
				if err != nil {
					return err
				}

				signatures = append(signatures, ringSig)
			}

			expense, err := app.expenseBuilder.Create().WithContent(expenseContent).WithSignatures(signatures).Now()
			if err != nil {
				return err
			}

			expFees = append(expFees, expense)

		}

		builder.WithFees(expFees)
	}

	trx, err := builder.Now()
	if err != nil {
		return err
	}

	err = identity.Wallets().Transact(trx)
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity.Hash(), identity, app.password, app.password)
}
