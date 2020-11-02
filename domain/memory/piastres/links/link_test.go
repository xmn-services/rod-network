package links

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestLink_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()

	// shareholder's PK:
	pk := signature.NewPrivateKeyFactory().Create()
	pubKey := pk.PublicKey()
	pubKeyHash, err := hashAdapter.FromBytes([]byte(pubKey.String()))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	pubKeys := []hash.Hash{
		*pubKeyHash,
	}

	// transaction expense bill lock:
	lock := locks.CreateLockForTests(pubKeys)

	// transaction expense bill:
	trxExpenseBillAmount := uint64(11)
	trxExpenseBill := bills.CreateBillForTests(lock, trxExpenseBillAmount)

	// transaction expense:
	trxExpenseContent := expenses.CreateContentForTests(trxExpenseBillAmount, []bills.Bill{
		trxExpenseBill,
	}, lock)

	trxExpenseSig, err := pk.RingSign(trxExpenseBill.Lock().Hash().String(), []signature.PublicKey{
		pubKey,
	})

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	trxFee := expenses.CreateExpenseForTests(trxExpenseContent, []signature.RingSignature{
		trxExpenseSig,
	})

	fees := []expenses.Expense{
		trxFee,
	}

	// transaction:
	executesOnTrigger := true
	amountPubKeyInRing := uint(20)
	trxIns := transactions.CreateTransactionWithFeesForTests(amountPubKeyInRing, executesOnTrigger, fees)

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
	blockIns := blocks.CreateBlockForTests(genesisIns, additional, trx)

	// link:
	index := uint(5)
	prevLink, _ := hashAdapter.FromBytes([]byte("this is the prev link"))
	linkIns := CreateLinkForTests(*prevLink, blockIns, index)

	// repository and service:
	genService, repository, service := CreateRepositoryServiceForTests()
	err = service.Save(linkIns)
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

	retTrx, err := repository.Retrieve(linkIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !linkIns.Hash().Compare(retTrx.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(linkIns)
	if err != nil {
		t.Errorf("the Link instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(link)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Link instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, linkIns)
}
