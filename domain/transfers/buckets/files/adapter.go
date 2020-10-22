package files

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToFile converts js data to a File instance
func (app *adapter) ToFile(js []byte) (File, error) {
	ins := new(file)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a file instance to js data
func (app *adapter) ToJSON(file File) ([]byte, error) {
	return json.Marshal(file)
}
