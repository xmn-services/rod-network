package states

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToState converts js data to a State instance
func (app *adapter) ToState(js []byte) (State, error) {
	ins := new(state)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a state instance to js data
func (app *adapter) ToJSON(state State) ([]byte, error) {
	return json.Marshal(state)
}
