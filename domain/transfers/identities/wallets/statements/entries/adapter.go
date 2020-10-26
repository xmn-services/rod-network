package entries

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToEntry converts js data to an Entry instance
func (app *adapter) ToEntry(js []byte) (Entry, error) {
	ins := new(entry)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a entry instance to js data
func (app *adapter) ToJSON(entry Entry) ([]byte, error) {
	return json.Marshal(entry)
}
