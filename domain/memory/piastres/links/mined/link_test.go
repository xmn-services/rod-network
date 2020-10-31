package mined

import (
	"testing"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/links"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/hash"
)

func TestLink_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()

	// shareholder's PK:
	pk := signature.NewPrivateKeyFactory().Create()
	pubKey := pk.PublicKey()
	pubKeyHash, _ := hashAdapter.FromBytes([]byte(pubKey.String()))

	// shareholders:
	power := uint(1)
	holders := []shareholders.ShareHolder{
		shareholders.CreateShareHolderForTests(power, *pubKeyHash),
	}

	// genesis lock:
	treeshold := uint(1)
	lock := locks.CreateLockForTests(holders, treeshold)

	// genesis bill:
	billAmount := uint64(5000)
	genBill := bills.CreateBillForTests(lock, billAmount)

	// genesis:
	blockDiffBase := uint(1)
	blockDiffIncreasePerTrx := float64(0.00001)
	linkDiff := uint(1)
	gen := genesis.CreateGenesisForTests(blockDiffBase, blockDiffIncreasePerTrx, linkDiff, genBill)

	// transaction expense bill:
	trxExpenseBillAmount := uint64(11)
	trxExpenseBill := bills.CreateBillForTests(lock, trxExpenseBillAmount)

	// transaction expense:
	trxExpenseContent := expenses.CreateContentForTests(trxExpenseBillAmount, []bills.Bill{
		trxExpenseBill,
	})

	trxExpenseSig, _ := pk.RingSign(trxExpenseBill.Lock().Hash().String(), []signature.PublicKey{
		pubKey,
	})

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
	trx := transactions.CreateTransactionWithFeesForTests(amountPubKeyInRing, executesOnTrigger, trxFees)

	// block:
	additional := uint(0)
	trxList := []transactions.Transaction{
		trx,
	}

	nextBlock := blocks.CreateBlockForTests(gen, additional, trxList)

	// link:
	index := uint(2)
	prevLink, _ := hashAdapter.FromBytes([]byte("prev link hash"))
	link := links.CreateLinkForTests(*prevLink, nextBlock, index)

	// mined link:
	mining := "232"
	minedLink := CreateLinkForTests(link, mining)

	// encode then decode:
	adapter := NewAdapter()
	encoded, err := adapter.Encode(minedLink)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	decoded, err := adapter.Decode(encoded)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	reEncoded, err := adapter.Encode(decoded)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if encoded != reEncoded {
		t.Errorf("the encoding and decoding process failed to work")
		return
	}
}
