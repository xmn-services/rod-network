package locks

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToLock converts js data to a Lock instance
func (app *adapter) ToLock(js []byte) (Lock, error) {
	ins := new(lock)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a lock instance to js data
func (app *adapter) ToJSON(lock Lock) ([]byte, error) {
	return json.Marshal(lock)
}
