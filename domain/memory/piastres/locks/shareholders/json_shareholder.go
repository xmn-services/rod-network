package shareholders

import (
	"time"
)

// JSONShareHolder represents a json representation of the ShareHolder instance
type JSONShareHolder struct {
	Key       string    `json:"key"`
	Power     uint      `json:"power"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONShareHolderFromShareHolder(ins ShareHolder) *JSONShareHolder {
	key := ins.Key().String()
	power := ins.Power()
	createdOn := ins.CreatedOn()
	return createJSONShareHolder(key, power, createdOn)
}

func createJSONShareHolder(
	key string,
	power uint,
	createdOn time.Time,
) *JSONShareHolder {
	out := JSONShareHolder{
		Key:       key,
		Power:     power,
		CreatedOn: createdOn,
	}

	return &out
}
