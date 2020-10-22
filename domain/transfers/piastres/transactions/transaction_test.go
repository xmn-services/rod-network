package transactions

import (
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestTransaction_isExpense_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	fees, _ := hashAdapter.FromBytes([]byte("to build the fees hash..."))
	expense, _ := hashAdapter.FromBytes([]byte("to build the expense hash..."))

	triggersOn := time.Now().UTC().Add(time.Second * 24 * 60 * 60 * 5)

	pk := signature.NewPrivateKeyFactory().Create()
	second := signature.NewPrivateKeyFactory().Create()
	sig, _ := pk.RingSign(hsh.String(), []signature.PublicKey{
		pk.PublicKey(),
		second.PublicKey(),
	})

	createdOn := time.Now().UTC()
	transaction, err := NewBuilder().
		Create().
		WithHash(*hsh).
		TriggersOn(triggersOn).
		WithFees(*fees).
		WithExpense(*expense).
		WithSignature(sig).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !transaction.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if sig.String() != transaction.Signature().String() {
		t.Errorf("the signature is invalid")
		return
	}

	if !triggersOn.Equal(transaction.TriggersOn()) {
		t.Errorf("the triggersOn is invalid, expected: %s, returned: %s", triggersOn.String(), transaction.TriggersOn().String())
		return
	}

	if transaction.ExecutesOnTrigger() {
		t.Errorf("the executesOnTrigger was expected to be false")
		return
	}

	if !transaction.Fees().Compare(*fees) {
		t.Errorf("the fees is invalid")
		return
	}

	if !transaction.IsExpense() {
		t.Errorf("the expense was expected to be valid")
		return
	}

	if !transaction.Expense().Compare(*expense) {
		t.Errorf("the expense is invalid")
		return
	}

	if transaction.IsCancel() {
		t.Errorf("the cancel was NOT expected to be valid")
		return
	}

	if transaction.Cancel() != nil {
		t.Errorf("the cancel was expected to be nil")
		return
	}

	if !transaction.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), transaction.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(transaction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTransaction, err := repository.Retrieve(transaction.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, transaction, retTransaction)
}

func TestTransaction_isExpense_executesOnTrigger_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	fees, _ := hashAdapter.FromBytes([]byte("to build the fees hash..."))
	expense, _ := hashAdapter.FromBytes([]byte("to build the expense hash..."))

	triggersOn := time.Now().UTC().Add(time.Second * 24 * 60 * 60 * 5)

	pk := signature.NewPrivateKeyFactory().Create()
	second := signature.NewPrivateKeyFactory().Create()
	sig, _ := pk.RingSign(hsh.String(), []signature.PublicKey{
		pk.PublicKey(),
		second.PublicKey(),
	})

	createdOn := time.Now().UTC()
	transaction, err := NewBuilder().
		Create().
		WithHash(*hsh).
		TriggersOn(triggersOn).
		ExecutesOnTrigger().
		WithFees(*fees).
		WithExpense(*expense).
		WithSignature(sig).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !transaction.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if sig.String() != transaction.Signature().String() {
		t.Errorf("the signature is invalid")
		return
	}

	if !triggersOn.Equal(transaction.TriggersOn()) {
		t.Errorf("the triggersOn is invalid, expected: %s, returned: %s", triggersOn.String(), transaction.TriggersOn().String())
		return
	}

	if !transaction.ExecutesOnTrigger() {
		t.Errorf("the executesOnTrigger was expected to be true")
		return
	}

	if !transaction.Fees().Compare(*fees) {
		t.Errorf("the fees is invalid")
		return
	}

	if !transaction.IsExpense() {
		t.Errorf("the expense was expected to be valid")
		return
	}

	if !transaction.Expense().Compare(*expense) {
		t.Errorf("the expense is invalid")
		return
	}

	if transaction.IsCancel() {
		t.Errorf("the cancel was NOT expected to be valid")
		return
	}

	if transaction.Cancel() != nil {
		t.Errorf("the cancel was expected to be nil")
		return
	}

	if !transaction.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), transaction.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(transaction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTransaction, err := repository.Retrieve(transaction.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, transaction, retTransaction)
}

func TestTransaction_isCancel_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	fees, _ := hashAdapter.FromBytes([]byte("to build the fees hash..."))
	cancel, _ := hashAdapter.FromBytes([]byte("to build the cancel hash..."))

	triggersOn := time.Now().UTC().Add(time.Second * 24 * 60 * 60 * 5)

	pk := signature.NewPrivateKeyFactory().Create()
	second := signature.NewPrivateKeyFactory().Create()
	sig, _ := pk.RingSign(hsh.String(), []signature.PublicKey{
		pk.PublicKey(),
		second.PublicKey(),
	})

	createdOn := time.Now().UTC()
	transaction, err := NewBuilder().
		Create().
		WithHash(*hsh).
		TriggersOn(triggersOn).
		WithFees(*fees).
		WithCancel(*cancel).
		WithSignature(sig).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !transaction.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if sig.String() != transaction.Signature().String() {
		t.Errorf("the signature is invalid")
		return
	}

	if !triggersOn.Equal(transaction.TriggersOn()) {
		t.Errorf("the triggersOn is invalid, expected: %s, returned: %s", triggersOn.String(), transaction.TriggersOn().String())
		return
	}

	if transaction.ExecutesOnTrigger() {
		t.Errorf("the executesOnTrigger was expected to be false")
		return
	}

	if !transaction.Fees().Compare(*fees) {
		t.Errorf("the fees is invalid")
		return
	}

	if transaction.IsExpense() {
		t.Errorf("the expense was NOT expected to be valid")
		return
	}

	if transaction.Expense() != nil {
		t.Errorf("the expense was expected to be nil")
		return
	}

	if !transaction.IsCancel() {
		t.Errorf("the cancel was expected to be valid")
		return
	}

	if !transaction.Cancel().Compare(*cancel) {
		t.Errorf("the cancel is invalid")
		return
	}

	if !transaction.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), transaction.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(transaction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTransaction, err := repository.Retrieve(transaction.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, transaction, retTransaction)
}

func TestTransaction_isCancel_executesOnTrigger_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	fees, _ := hashAdapter.FromBytes([]byte("to build the fees hash..."))
	cancel, _ := hashAdapter.FromBytes([]byte("to build the cancel hash..."))

	triggersOn := time.Now().UTC().Add(time.Second * 24 * 60 * 60 * 5)

	pk := signature.NewPrivateKeyFactory().Create()
	second := signature.NewPrivateKeyFactory().Create()
	sig, _ := pk.RingSign(hsh.String(), []signature.PublicKey{
		pk.PublicKey(),
		second.PublicKey(),
	})

	createdOn := time.Now().UTC()
	transaction, err := NewBuilder().
		Create().
		WithHash(*hsh).
		TriggersOn(triggersOn).
		ExecutesOnTrigger().
		WithFees(*fees).
		WithCancel(*cancel).
		WithSignature(sig).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !transaction.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if sig.String() != transaction.Signature().String() {
		t.Errorf("the signature is invalid")
		return
	}

	if !triggersOn.Equal(transaction.TriggersOn()) {
		t.Errorf("the triggersOn is invalid, expected: %s, returned: %s", triggersOn.String(), transaction.TriggersOn().String())
		return
	}

	if !transaction.ExecutesOnTrigger() {
		t.Errorf("the executesOnTrigger was expected to be true")
		return
	}

	if !transaction.Fees().Compare(*fees) {
		t.Errorf("the fees is invalid")
		return
	}

	if transaction.IsExpense() {
		t.Errorf("the expense was NOT expected to be valid")
		return
	}

	if transaction.Expense() != nil {
		t.Errorf("the expense was expected to be nil")
		return
	}

	if !transaction.IsCancel() {
		t.Errorf("the cancel was expected to be valid")
		return
	}

	if !transaction.Cancel().Compare(*cancel) {
		t.Errorf("the cancel is invalid")
		return
	}

	if !transaction.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), transaction.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(transaction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTransaction, err := repository.Retrieve(transaction.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, transaction, retTransaction)
}
