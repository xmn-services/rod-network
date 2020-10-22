package chunks

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToChunk converts js data to a Chunk instance
func (app *adapter) ToChunk(js []byte) (Chunk, error) {
	ins := new(chunk)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a chunk instance to js data
func (app *adapter) ToJSON(chunk Chunk) ([]byte, error) {
	return json.Marshal(chunk)
}
