package transactions

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestTransaction_withFees_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()

	// shareholder's PK:
	pk := signature.NewPrivateKeyFactory().Create()
	pubKey := pk.PublicKey()
	pubKeyHash, err := hashAdapter.FromBytes([]byte(pubKey.String()))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// shareholders:
	power := uint(1)
	holders := []shareholders.ShareHolder{
		shareholders.CreateShareHolderForTests(power, *pubKeyHash),
	}

	// transaction expense bill lock:
	treeshold := uint(1)
	lock := locks.CreateLockForTests(holders, treeshold)

	// transaction expense bill:
	trxExpenseBillAmount := uint64(11)
	trxExpenseBill := bills.CreateBillForTests(lock, trxExpenseBillAmount)

	// transaction expense cancel lock:
	cancelTreeshold := uint(1)
	trxExpenseCancelLock := locks.CreateLockForTests(holders, cancelTreeshold)

	// transaction expense:
	trxExpenseContent := expenses.CreateContentForTests(trxExpenseBillAmount, trxExpenseBill, trxExpenseCancelLock)

	trxExpenseSig, err := pk.RingSign(trxExpenseContent.From().Lock().Hash().String(), []signature.PublicKey{
		pubKey,
	})

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	trxFee := expenses.CreateExpenseForTests(trxExpenseContent, []signature.RingSignature{
		trxExpenseSig,
	})

	trxFees := []expenses.Expense{
		trxFee,
	}

	// transaction:
	executesOnTrigger := true
	amountPubKeyInRing := uint(20)
	trxIns, _ := CreateTransactionWithFeesForTests(amountPubKeyInRing, executesOnTrigger, trxFees)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// repository and service:
	repository, service := CreateRepositoryServiceForTests(20)
	err = service.Save(trxIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTrx, err := repository.Retrieve(trxIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !trxIns.Hash().Compare(retTrx.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(trxIns)
	if err != nil {
		t.Errorf("the Transaction instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(transaction)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Transaction instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, trxIns)
}
