package genesis

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestGenesis_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pkFactory := signature.NewPrivateKeyFactory()
	firstPK := pkFactory.Create()
	secondPK := pkFactory.Create()
	thirdPK := pkFactory.Create()

	firstHash, _ := hashAdapter.FromBytes([]byte(firstPK.PublicKey().String()))
	secondHash, _ := hashAdapter.FromBytes([]byte(secondPK.PublicKey().String()))
	thirdHash, _ := hashAdapter.FromBytes([]byte(thirdPK.PublicKey().String()))

	shareholders := []shareholders.ShareHolder{
		shareholders.CreateShareHolderForTests(52, *firstHash),
		shareholders.CreateShareHolderForTests(49, *secondHash),
		shareholders.CreateShareHolderForTests(1, *thirdHash),
	}

	treeshold := uint(51)
	lock := locks.CreateLockForTests(shareholders, treeshold)

	amount := uint64(56)
	bill := bills.CreateBillForTests(lock, amount)

	// genesis:
	blockDiffBase := uint(1)
	blockDiffIncreasePerTrx := float64(1.0)
	linkDiff := uint(1)
	genesisIns := CreateGenesisForTests(blockDiffBase, blockDiffIncreasePerTrx, linkDiff, bill)

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(genesisIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBill, err := repository.Retrieve()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !genesisIns.Hash().Compare(retBill.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(genesisIns)
	if err != nil {
		t.Errorf("the Bill instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(genesis)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Bill instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, genesisIns)
}
