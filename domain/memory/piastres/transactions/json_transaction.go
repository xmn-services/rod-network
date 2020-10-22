package transactions

import "time"

// JSONTransaction represents a json transaction instance
type JSONTransaction struct {
	Content   *JSONContent `json:"content"`
	Signature string       `json:"signature"`
	CreatedOn time.Time    `json:"created_on"`
}

func createJSONTransactionFromTransaction(trx Transaction) *JSONTransaction {
	content := createJSONContentFromContent(trx.Content())
	sig := trx.Signature().String()
	createdOn := trx.CreatedOn()
	return createJSONTransaction(content, sig, createdOn)
}

func createJSONTransaction(
	content *JSONContent,
	signature string,
	createdOn time.Time,
) *JSONTransaction {
	out := JSONTransaction{
		Content:   content,
		Signature: signature,
		CreatedOn: createdOn,
	}

	return &out
}
