package wallets

import "time"

type jsonWallet struct {
	Hash        string    `json:"hash"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Statement   string    `json:"statement"`
	Bills       []string  `json:"bills"`
	CreatedOn   time.Time `json:"created_on"`
}

func createJSONWalletFromWallet(ins Wallet) *jsonWallet {
	hash := ins.Hash().String()
	name := ins.Name()
	description := ""
	if ins.HasDescription() {
		description = ins.Description()
	}

	statement := ins.Statement().String()

	lst := []string{}
	if ins.HasBills() {
		bills := ins.Bills()
		for _, oneBill := range bills {
			lst = append(lst, oneBill.String())
		}
	}

	createdOn := ins.CreatedOn()
	return createJSONWallet(hash, name, description, statement, lst, createdOn)
}

func createJSONWallet(
	hash string,
	name string,
	description string,
	statement string,
	bills []string,
	createdOn time.Time,
) *jsonWallet {
	out := jsonWallet{
		Hash:        hash,
		Name:        name,
		Description: description,
		Statement:   statement,
		Bills:       bills,
		CreatedOn:   createdOn,
	}

	return &out
}
