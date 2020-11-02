package piastres

import (
	"math/rand"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/domain/memory/identities"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
)

type current struct {
	pkFactory                                    signature.PrivateKeyFactory
	bucketRepository                             buckets.Repository
	lockBuilder                                  locks.Builder
	expenseBuilder                               expenses.Builder
	expenseContentBuilder                        expenses.ContentBuilder
	trxBuilder                                   transactions.Builder
	identity                                     identities.Identity
	amountAdditionalPubKeysPerShareHolderPerRing int
}

func createCurrent(
	identity identities.Identity,
) Current {
	out := current{
		identity: identity,
	}

	return &out
}

// Bucket executes a bucket transaction
func (app *current) Bucket(absolutePath string, fees []Fee) error {
	bucket, err := app.identity.Buckets().Fetch(absolutePath)
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
			walletBills, err := app.identity.Wallets().Fetch(amount)
			if err != nil {
				return err
			}

			pks := []signature.PrivateKey{}
			bills := []bills.Bill{}
			for _, oneWalletBill := range walletBills {
				bill := oneWalletBill.Bill()
				pk := oneWalletBill.PrivateKey()

				bills = append(bills, bill)
				pks = append(pks, pk)
			}

			remainingPK := app.pkFactory.Create()
			ringPublicKeys := []signature.PublicKey{}
			lockPublicKeys := []hash.Hash{}

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
			for _, onePK := range pks {
				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)
				insert := r1.Intn(app.amountAdditionalPubKeysPerShareHolderPerRing)

				pubKeys := []signature.PublicKey{}
				for i := 0; i < app.amountAdditionalPubKeysPerShareHolderPerRing; i++ {
					pubKey := app.pkFactory.Create().PublicKey()
					pubKeys = append(pubKeys, pubKey)

					// insert the current PK's pubkey:
					if i == insert {
						pubKeys = append(pubKeys, onePK.PublicKey())
					}
				}

				ringSig, err := onePK.RingSign(msg, pubKeys)
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

	return app.identity.Wallets().Transact(trx)
}
