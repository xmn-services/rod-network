package transactions

import (
	"time"
)

type jsonTransaction struct {
	Hash       string    `json:"hash"`
	Signature  string    `json:"signature"`
	TriggersOn time.Time `json:"triggers_on"`
	Fees       []string  `json:"fees"`
	Bucket     string    `json:"bucket"`
	Cancel     string    `json:"cancel"`
	CreatedOn  time.Time `json:"created_on"`
}

func createJSONTransactionFromTransaction(ins Transaction) *jsonTransaction {
	hash := ins.Hash().String()
	signature := ins.Signature().String()
	triggersOn := ins.TriggersOn()

	strFees := []string{}
	if ins.HasFees() {
		fees := ins.Fees()
		for _, oneFee := range fees {
			strFees = append(strFees, oneFee.String())
		}
	}

	bucket := ""
	if ins.IsBucket() {
		bucket = ins.Bucket().String()
	}

	cancel := ""
	if ins.IsCancel() {
		cancel = ins.Cancel().String()
	}

	createdOn := ins.CreatedOn()
	return createJSONTransaction(hash, signature, triggersOn, strFees, bucket, cancel, createdOn)
}

func createJSONTransaction(
	hash string,
	signature string,
	triggersOn time.Time,
	fees []string,
	bucket string,
	cancel string,
	createdOn time.Time,
) *jsonTransaction {
	out := jsonTransaction{
		Hash:       hash,
		Signature:  signature,
		TriggersOn: triggersOn,
		Fees:       fees,
		Bucket:     bucket,
		Cancel:     cancel,
		CreatedOn:  createdOn,
	}

	return &out
}
