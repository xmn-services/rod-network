package entries

import "time"

type jsonEntry struct {
	Hash         string    `json:"hash"`
	Name         string    `json:"string"`
	Transactions []string  `json:"transactions"`
	Description  string    `json:"description"`
	CreatedOn    time.Time `json:"created_on"`
}

func createJSONEntryFromEntry(ins Entry) *jsonEntry {
	hash := ins.Hash().String()
	name := ins.Name()
	description := ""
	if ins.HasDescription() {
		description = ins.Description()
	}

	lst := []string{}
	transactions := ins.Transactions()
	for _, oneTrx := range transactions {
		lst = append(lst, oneTrx.String())
	}

	createdOn := ins.CreatedOn()
	return createJSONEntry(hash, name, lst, description, createdOn)
}

func createJSONEntry(
	hash string,
	name string,
	transactions []string,
	description string,
	createdOn time.Time,
) *jsonEntry {
	out := jsonEntry{
		Hash:         hash,
		Name:         name,
		Transactions: transactions,
		Description:  description,
		CreatedOn:    createdOn,
	}

	return &out
}
