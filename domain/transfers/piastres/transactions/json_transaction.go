package transactions

import (
	"time"
)

type jsonTransaction struct {
	Hash      string    `json:"hash"`
	Address   string    `json:"address"`
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

	address := ""
	if ins.HasAddress() {
		address = ins.Address().String()
	}

	createdOn := ins.CreatedOn()
	return createJSONTransaction(hash, address, strFees, bucket, createdOn)
}

func createJSONTransaction(
	hash string,
	address string,
	fees []string,
	bucket string,
	createdOn time.Time,
) *jsonTransaction {
	out := jsonTransaction{
		Hash:      hash,
		Address:   address,
		Fees:      fees,
		Bucket:    bucket,
		CreatedOn: createdOn,
	}

	return &out
}
