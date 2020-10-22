package shareholders

import (
	"time"
)

type jsonShareHolder struct {
	Hash      string    `json:"hash"`
	Key       string    `json:"key"`
	Power     uint      `json:"power"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONShareHolderFromShareHolder(ins ShareHolder) *jsonShareHolder {
	hash := ins.Hash().String()
	key := ins.Key().String()
	power := ins.Power()
	createdOn := ins.CreatedOn()
	return createJSONShareHolder(hash, key, power, createdOn)
}

func createJSONShareHolder(
	hash string,
	key string,
	power uint,
	createdOn time.Time,
) *jsonShareHolder {
	out := jsonShareHolder{
		Hash:      hash,
		Key:       key,
		Power:     power,
		CreatedOn: createdOn,
	}

	return &out
}
