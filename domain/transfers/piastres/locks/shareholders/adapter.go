package shareholders

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToShareHolder converts js data to a ShareHolder instance
func (app *adapter) ToShareHolder(js []byte) (ShareHolder, error) {
	ins := new(shareHolder)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a shareHolder instance to js data
func (app *adapter) ToJSON(shareHolder ShareHolder) ([]byte, error) {
	return json.Marshal(shareHolder)
}
