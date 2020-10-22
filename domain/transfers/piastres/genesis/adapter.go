package genesis

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToGenesis converts js data to a Genesis instance
func (app *adapter) ToGenesis(js []byte) (Genesis, error) {
	ins := new(genesis)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a genesis instance to js data
func (app *adapter) ToJSON(genesis Genesis) ([]byte, error) {
	return json.Marshal(genesis)
}
