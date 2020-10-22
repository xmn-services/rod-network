package informations

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToInformation converts js data to a Information instance
func (app *adapter) ToInformation(js []byte) (Information, error) {
	ins := new(information)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a information instance to js data
func (app *adapter) ToJSON(information Information) ([]byte, error) {
	return json.Marshal(information)
}
