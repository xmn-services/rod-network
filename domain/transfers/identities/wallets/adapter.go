package wallets

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToWallet converts js data to a Wallet instance
func (app *adapter) ToWallet(js []byte) (Wallet, error) {
	ins := new(wallet)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a wallet instance to js data
func (app *adapter) ToJSON(wallet Wallet) ([]byte, error) {
	return json.Marshal(wallet)
}
