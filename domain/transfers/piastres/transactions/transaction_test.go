package transactions

import (
	"testing"
	"time"

	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestTransaction_isBucket_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	fee, _ := hashAdapter.FromBytes([]byte("to build the fees hash..."))
	bucket, _ := hashAdapter.FromBytes([]byte("to build the bucket hash..."))

	fees := []hash.Hash{
		*fee,
	}

	createdOn := time.Now().UTC()
	transaction, err := NewBuilder().
		Create().
		WithHash(*hsh).
		WithFees(fees).
		WithBucket(*bucket).
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

	if !transaction.HasFees() {
		t.Errorf("the fees were expected to be valid")
		return
	}

	trFees := transaction.Fees()
	if len(trFees) != len(fees) {
		t.Errorf("%d fees were expected, %d returned", len(fees), len(trFees))
		return
	}

	for index, oneTrFee := range fees {
		if !oneTrFee.Compare(trFees[index]) {
			t.Errorf("the fee (index: %d) is invalid", index)
			return
		}
	}

	if !transaction.HasBucket() {
		t.Errorf("the bucket was expected to be valid")
		return
	}

	if !transaction.Bucket().Compare(*bucket) {
		t.Errorf("the bucket is invalid")
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
