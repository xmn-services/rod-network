package bills

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
)

func TestBill_Success(t *testing.T) {
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

	amount := uint(56)
	billIns := CreateBillForTests(lock, amount)

	if !billIns.Lock().Hash().Compare(lock.Hash()) {
		t.Errorf("the lock is invalid, expected: %s, returned: %s", lock.Hash().String(), billIns.Lock().Hash().String())
		return
	}

	if billIns.Amount() != amount {
		t.Errorf("the amount is invalid, expected: %d, returned: %d", amount, billIns.Amount())
		return
	}

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(billIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBill, err := repository.Retrieve(billIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !billIns.Hash().Compare(retBill.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(billIns)
	if err != nil {
		t.Errorf("the Bill instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(bill)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Bill instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, billIns)
}
