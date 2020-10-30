package genesis

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets"
	wallet_bills "github.com/xmn-services/rod-network/domain/memory/identities/wallets/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
)

type current struct {
	hashAdapter        hash.Adapter
	shareHolderBuilder shareholders.Builder
	lockBuilder        locks.Builder
	billBuilder        bills.Builder
	genesisBuilder     genesis.Builder
	genesisRepository  genesis.Repository
	genesisService     genesis.Service
	identityRepository identities.Repository
	walletBillBuilder  wallet_bills.Builder
	walletBuilder      wallets.Builder
	identityBuilder    identities.Builder
	identityService    identities.Service
}

func createCurrent() Current {
	out := current{}
	return &out
}

// Init initializes the genesis block
func (app *current) Init(
	name string,
	password string,
	seed string,
	power uint,
	walletName string,
	amountUnits uint64,
	blockDifficultyBase uint,
	blockDifficultyIncreasePerTrx float64,
	linkDifficulty uint,
) error {
	_, err := app.genesisRepository.Retrieve()
	if err == nil {
		return errors.New("the genesis block has already been created")
	}

	identity, err := app.identityRepository.Retrieve(name, seed, password)
	if err != nil {
		return err
	}

	createdOn := time.Now().UTC()
	pk := signature.NewPrivateKeyFactory().Create()
	pubKey := pk.PublicKey()
	pubKeyHash, err := app.hashAdapter.FromString(pubKey.String())
	if err != nil {
		return err
	}

	shareHolder, err := app.shareHolderBuilder.Create().WithKey(*pubKeyHash).WithPower(power).CreatedOn(createdOn).Now()
	if err != nil {
		return err
	}

	lock, err := app.lockBuilder.Create().WithShareHolders([]shareholders.ShareHolder{
		shareHolder,
	}).WithTreeshold(power).CreatedOn(createdOn).Now()

	if err != nil {
		return err
	}

	bill, err := app.billBuilder.Create().WithLock(lock).WithAmount(amountUnits).CreatedOn(createdOn).Now()
	if err != nil {
		return err
	}

	gen, err := app.genesisBuilder.Create().
		WithBlockDifficultyBase(blockDifficultyBase).
		WithBlockDifficultyIncreasePerTrx(blockDifficultyIncreasePerTrx).
		WithLinkDifficulty(linkDifficulty).
		WithBill(bill).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return err
	}

	walletBill, err := app.walletBillBuilder.Create().WithBill(bill).WithPrivateKeys([]signature.PrivateKey{
		pk,
	}).CreatedOn(createdOn).Now()

	if err != nil {
		return err
	}

	wallet, err := app.walletBuilder.Create().WithBills([]wallet_bills.Bill{
		walletBill,
	}).WithName(walletName).CreatedOn(createdOn).Now()

	if err != nil {
		return err
	}

	root := identity.Root()
	identityCreatedOn := identity.CreatedOn()
	lastUpdatedOn := time.Now().UTC()
	wallets := []wallets.Wallet{}
	if identity.HasWallets() {
		wallets = identity.Wallets()
	}

	wallets = append(wallets, wallet)
	identityBuilder := app.identityBuilder.Create().WithSeed(seed).WithName(name).WithRoot(root).WithWallets(wallets).CreatedOn(identityCreatedOn).LastUpdatedOn(lastUpdatedOn)
	if identity.HasBuckets() {
		buckets := identity.Buckets()
		identityBuilder.WithBuckets(buckets)
	}

	updatedIdentity, err := identityBuilder.Now()
	if err != nil {
		return err
	}

	err = app.identityService.Update(updatedIdentity, name, seed, password)
	if err != nil {
		return err
	}

	return app.genesisService.Save(gen)
}
