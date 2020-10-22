package chains

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToChain converts js data to a Chain instance
func (app *adapter) ToChain(js []byte) (Chain, error) {
	ins := new(chain)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a chain instance to js data
func (app *adapter) ToJSON(chain Chain) ([]byte, error) {
	return json.Marshal(chain)
}
