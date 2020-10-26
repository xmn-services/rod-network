package identities

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToIdentity converts js data to a Identity instance
func (app *adapter) ToIdentity(js []byte) (Identity, error) {
	ins := new(identity)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a identity instance to js data
func (app *adapter) ToJSON(identity Identity) ([]byte, error) {
	return json.Marshal(identity)
}
