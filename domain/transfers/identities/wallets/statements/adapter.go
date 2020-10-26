package statements

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToStatement converts js data to a Statement instance
func (app *adapter) ToStatement(js []byte) (Statement, error) {
	ins := new(statement)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a statement instance to js data
func (app *adapter) ToJSON(statement Statement) ([]byte, error) {
	return json.Marshal(statement)
}
