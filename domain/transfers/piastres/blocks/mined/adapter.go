package mined

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToBlock converts js data to a Block instance
func (app *adapter) ToBlock(js []byte) (Block, error) {
	ins := new(block)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a block instance to js data
func (app *adapter) ToJSON(block Block) ([]byte, error) {
	return json.Marshal(block)
}
