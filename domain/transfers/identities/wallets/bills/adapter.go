package bills

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToBill converts js data to a Bill instance
func (app *adapter) ToBill(js []byte) (Bill, error) {
	ins := new(bill)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a bill instance to js data
func (app *adapter) ToJSON(bill Bill) ([]byte, error) {
	return json.Marshal(bill)
}
