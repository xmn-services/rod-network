package expenses

import (
	"reflect"
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestExpense_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	lock, _ := hashAdapter.FromBytes([]byte("to build the lock hash..."))
	fromSingle, _ := hashAdapter.FromBytes([]byte("to build the from hash..."))
	amount := uint64(56)
	createdOn := time.Now().UTC()

	pk := signature.NewPrivateKeyFactory().Create()
	ring := []signature.PublicKey{
		pk.PublicKey(),
	}

	sig, _ := pk.RingSign(hsh.String(), ring)
	signatures := [][]signature.RingSignature{
		[]signature.RingSignature{
			sig,
		},
	}

	from := []hash.Hash{
		*fromSingle,
	}

	expense, err := NewBuilder().Create().WithHash(*hsh).From(from).WithLock(*lock).WithAmount(amount).WithSignatures(signatures).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !expense.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	retFrom := expense.From()
	if len(retFrom) != len(from) {
		t.Errorf("%d from hashes were expectyed, %d returned", len(from), len(retFrom))
		return
	}

	if !expense.Lock().Compare(*lock) {
		t.Errorf("the lock hash is invalid")
		return
	}

	if !reflect.DeepEqual(signatures, expense.Signatures()) {
		t.Errorf("the ring signatures are invalid")
		return
	}

	if expense.HasRemaining() {
		t.Errorf("the remaining hash was NOT expected to be valid")
		return
	}

	if expense.Remaining() != nil {
		t.Errorf("the remaining hash was expected to be nil")
		return
	}

	if expense.Amount() != amount {
		t.Errorf("the amount is invalid, expected: %d, returned: %d", amount, expense.Amount())
		return
	}

	if !expense.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), expense.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(expense)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retExpense, err := repository.Retrieve(expense.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, expense, retExpense)
}

func TestExpense_withRemaining_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	lock, _ := hashAdapter.FromBytes([]byte("to build the lock hash..."))
	fromSingle, _ := hashAdapter.FromBytes([]byte("to build the from hash..."))
	remaining, _ := hashAdapter.FromBytes([]byte("to build the remaining hash..."))
	amount := uint64(56)
	createdOn := time.Now().UTC()

	pk := signature.NewPrivateKeyFactory().Create()
	ring := []signature.PublicKey{
		pk.PublicKey(),
	}

	sig, _ := pk.RingSign(hsh.String(), ring)
	signatures := [][]signature.RingSignature{
		[]signature.RingSignature{
			sig,
		},
	}

	from := []hash.Hash{
		*fromSingle,
	}

	expense, err := NewBuilder().Create().WithHash(*hsh).From(from).WithLock(*lock).WithAmount(amount).WithRemaining(*remaining).WithSignatures(signatures).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !expense.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	retFrom := expense.From()
	if len(retFrom) != len(from) {
		t.Errorf("%d from hashes were expectyed, %d returned", len(from), len(retFrom))
		return
	}

	if !expense.Lock().Compare(*lock) {
		t.Errorf("the lock hash is invalid")
		return
	}

	if !reflect.DeepEqual(signatures, expense.Signatures()) {
		t.Errorf("the ring signatures are invalid")
		return
	}

	if !expense.HasRemaining() {
		t.Errorf("the remaining hash was expected to be valid")
		return
	}

	if !expense.Remaining().Compare(*remaining) {
		t.Errorf("the remaining hash is invalid")
		return
	}

	if expense.Amount() != amount {
		t.Errorf("the amount is invalid, expected: %d, returned: %d", amount, expense.Amount())
		return
	}

	if !expense.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), expense.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(expense)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retExpense, err := repository.Retrieve(expense.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, expense, retExpense)
}
