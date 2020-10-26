package bills

import "time"

type jsonBill struct {
	Hash        string    `json:"hash"`
	Bill        string    `json:"bill"`
	PrivateKeys []string  `json:"pks"`
	CreatedOn   time.Time `json:"created_on"`
}

func createJSONBillFromBill(ins Bill) *jsonBill {
	hash := ins.Hash().String()
	bill := ins.Bill().String()

	pksStr := []string{}
	pks := ins.PrivateKeys()
	for _, onePK := range pks {
		str := onePK.String()
		pksStr = append(pksStr, str)
	}

	createdOn := ins.CreatedOn()
	return createJSONBill(hash, bill, pksStr, createdOn)
}

func createJSONBill(
	hash string,
	bill string,
	pks []string,
	createdOn time.Time,
) *jsonBill {
	out := jsonBill{
		Hash:        hash,
		Bill:        bill,
		PrivateKeys: pks,
		CreatedOn:   createdOn,
	}

	return &out
}
