package cancels

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
)

func TestExpense_Success(t *testing.T) {
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

	billAmount := uint(56)
	bill := bills.CreateBillForTests(lock, billAmount)

	cancelInsTreeshold := uint(22)
	cancelInsLock := locks.CreateLockForTests(shareholders, cancelInsTreeshold)

	amount := billAmount - 1
	content := expenses.CreateContentForTests(amount, bill, cancelInsLock)

	ring := []signature.PublicKey{
		firstPK.PublicKey(),
		secondPK.PublicKey(),
		thirdPK.PublicKey(),
	}

	msg := content.From().Lock().Hash().String()
	firstSig, _ := firstPK.RingSign(msg, ring)
	secondSig, _ := secondPK.RingSign(msg, ring)
	signatures := []signature.RingSignature{
		firstSig,
		secondSig,
	}

	expense := expenses.CreateExpenseForTests(content, signatures)

	cancelInsMsg := content.Cancel().Hash().String()
	cancelInsFirstSig, _ := firstPK.RingSign(cancelInsMsg, ring)
	cancelInsSecondSig, _ := secondPK.RingSign(cancelInsMsg, ring)
	cancelInsSignatures := []signature.RingSignature{
		cancelInsFirstSig,
		cancelInsSecondSig,
	}

	cancelIns := CreateCancelForTests(expense, lock, cancelInsSignatures)

	if !cancelIns.Expense().Hash().Compare(expense.Hash()) {
		t.Errorf("the expense is invalid, expected: %s, returned: %s", expense.Hash().String(), cancelIns.Expense().Hash().String())
		return
	}

	if !cancelIns.Lock().Hash().Compare(lock.Hash()) {
		t.Errorf("the lock is invalid, expected: %s, returned: %s", lock.Hash().String(), cancelIns.Expense().Hash().String())
		return
	}

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(cancelIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retCancel, err := repository.Retrieve(cancelIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !cancelIns.Hash().Compare(retCancel.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(cancelIns)
	if err != nil {
		t.Errorf("the Cancel instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(cancel)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Cancel instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, cancelIns)
}
