package transactions

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToTransaction converts js data to a Transaction instance
func (app *adapter) ToTransaction(js []byte) (Transaction, error) {
	ins := new(transaction)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a transaction instance to js data
func (app *adapter) ToJSON(transaction Transaction) ([]byte, error) {
	return json.Marshal(transaction)
}
