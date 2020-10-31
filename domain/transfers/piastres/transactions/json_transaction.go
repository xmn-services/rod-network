package transactions

import (
	"time"
)

type jsonTransaction struct {
	Hash      string    `json:"hash"`
	Fees      []string  `json:"fees"`
	Bucket    string    `json:"bucket"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONTransactionFromTransaction(ins Transaction) *jsonTransaction {
	hash := ins.Hash().String()

	strFees := []string{}
	if ins.HasFees() {
		fees := ins.Fees()
		for _, oneFee := range fees {
			strFees = append(strFees, oneFee.String())
		}
	}

	bucket := ""
	if ins.HasBucket() {
		bucket = ins.Bucket().String()
	}

	createdOn := ins.CreatedOn()
	return createJSONTransaction(hash, strFees, bucket, createdOn)
}

func createJSONTransaction(
	hash string,
	fees []string,
	bucket string,
	createdOn time.Time,
) *jsonTransaction {
	out := jsonTransaction{
		Hash:      hash,
		Fees:      fees,
		Bucket:    bucket,
		CreatedOn: createdOn,
	}

	return &out
}
