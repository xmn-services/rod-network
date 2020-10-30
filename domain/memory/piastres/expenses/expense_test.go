package expenses

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
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

	billAmount := uint64(56)
	bill := bills.CreateBillForTests(lock, billAmount)

	cancelTreeshold := uint(22)
	cancelLock := locks.CreateLockForTests(shareholders, cancelTreeshold)

	amount := billAmount - 1
	content := CreateContentForTests(amount, bill, cancelLock)

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

	expenseIns := CreateExpenseForTests(content, signatures)

	if !expenseIns.Content().Hash().Compare(content.Hash()) {
		t.Errorf("the content is invalid, expected: %s, returned: %s", content.Hash().String(), expenseIns.Content().Hash().String())
		return
	}

	if !reflect.DeepEqual(signatures, expenseIns.Signatures()) {
		t.Errorf("the ring signatures are invalid")
		return
	}

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(expenseIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retExpense, err := repository.Retrieve(expenseIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !expenseIns.Hash().Compare(retExpense.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(expenseIns)
	if err != nil {
		t.Errorf("the Expense instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(expense)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to an Expense instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, expenseIns)
}

/*
func TestContent_withZeroAmount_Success(t *testing.T) {
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

	amount := uint(0)
	expenseIns := CreateExpenseForTests(amount, bill)

	if !expenseIns.From().Hash().Compare(bill.Hash()) {
		t.Errorf("the from bill is invalid, expected: %s, returned: %s", bill.Hash().String(), expenseIns.From().Hash().String())
		return
	}

	if expenseIns.Amount() != amount {
		t.Errorf("the amount is invalid, expected: %d, returned: %d", amount, expenseIns.Amount())
		return
	}

	if expenseIns.HasRemaining() {
		t.Errorf("the expenseIns was not expected a remaining lock instance")
		return
	}

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(expenseIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retExpense, err := repository.Retrieve(expenseIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !expenseIns.Hash().Compare(retExpense.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}
}

func TestContent_withRemaining_Success(t *testing.T) {
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

	amount := billAmount - 1
	expenseIns := CreateExpenseWithRemainingForTests(amount, bill, lock)

	if !expenseIns.From().Hash().Compare(bill.Hash()) {
		t.Errorf("the from bill is invalid, expected: %s, returned: %s", bill.Hash().String(), expenseIns.From().Hash().String())
		return
	}

	if expenseIns.Amount() != amount {
		t.Errorf("the amount is invalid, expected: %d, returned: %d", amount, expenseIns.Amount())
		return
	}

	if !expenseIns.HasRemaining() {
		t.Errorf("the expenseIns was expecting a remaining lock instance")
		return
	}

	if !expenseIns.Remaining().Hash().Compare(lock.Hash()) {
		t.Errorf("the remaining lock is invalid, expected: %s, returned: %s", lock.Hash().String(), expenseIns.Remaining().Hash().String())
		return
	}

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(expenseIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retExpense, err := repository.Retrieve(expenseIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !expenseIns.Hash().Compare(retExpense.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}
}

func TestContent_amountIsTooBig_returnsError(t *testing.T) {
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

	amount := billAmount + 1
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().From(bill).WithAmount(amount).CreatedOn(createdOn).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestContent_withRemaining_nothingIsRemaining_returnsError(t *testing.T) {
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
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().From(bill).WithAmount(billAmount).CreatedOn(createdOn).WithRemaining(lock).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
*/
