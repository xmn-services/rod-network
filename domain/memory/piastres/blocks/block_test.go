package blocks

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestBlock_Success(t *testing.T) {
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

	// transaction expense:
	trxExpenseContent := expenses.CreateContentForTests(trxExpenseBillAmount, []bills.Bill{
		trxExpenseBill,
	})

	trxExpenseSig, err := pk.RingSign(trxExpenseBill.Lock().Hash().String(), []signature.PublicKey{
		pubKey,
	})

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	trxFee := expenses.CreateExpenseForTests(trxExpenseContent, [][]signature.RingSignature{
		[]signature.RingSignature{
			trxExpenseSig,
		},
	})

	trxFees := []expenses.Expense{
		trxFee,
	}

	// transaction:
	executesOnTrigger := true
	amountPubKeyInRing := uint(20)
	trxIns := transactions.CreateTransactionWithFeesForTests(amountPubKeyInRing, executesOnTrigger, trxFees)

	// transactions:
	trx := []transactions.Transaction{
		trxIns,
	}

	// genesis bill:
	amount := uint64(56)
	bill := bills.CreateBillForTests(lock, amount)

	// genesis:
	blockDiffBase := uint(1)
	blockDiffIncreasePerTrx := float64(1.0)
	linkDiff := uint(1)
	genesisIns := genesis.CreateGenesisForTests(blockDiffBase, blockDiffIncreasePerTrx, linkDiff, bill)

	// block:
	additional := uint(0)
	blockIns := CreateBlockForTests(genesisIns, additional, trx)

	// repository and service:
	genService, repository, service := CreateRepositoryServiceForTests()
	err = service.Save(blockIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// save the genesis;
	err = genService.Save(genesisIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTrx, err := repository.Retrieve(blockIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !blockIns.Hash().Compare(retTrx.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(blockIns)
	if err != nil {
		t.Errorf("the Block instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(block)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Block instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, blockIns)
}
