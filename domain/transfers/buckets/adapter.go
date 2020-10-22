package buckets

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToBucket converts js data to a Bucket instance
func (app *adapter) ToBucket(js []byte) (Bucket, error) {
	ins := new(bucket)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a bucket instance to js data
func (app *adapter) ToJSON(bucket Bucket) ([]byte, error) {
	return json.Marshal(bucket)
}
