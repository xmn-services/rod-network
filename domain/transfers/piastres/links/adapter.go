package links

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToLink converts js data to a Link instance
func (app *adapter) ToLink(js []byte) (Link, error) {
	ins := new(link)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a link instance to js data
func (app *adapter) ToJSON(link Link) ([]byte, error) {
	return json.Marshal(link)
}
