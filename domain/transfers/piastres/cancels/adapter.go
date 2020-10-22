package cancels

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToCancel converts js data to a Cancel instance
func (app *adapter) ToCancel(js []byte) (Cancel, error) {
	ins := new(cancel)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a cancel instance to js data
func (app *adapter) ToJSON(cancel Cancel) ([]byte, error) {
	return json.Marshal(cancel)
}
