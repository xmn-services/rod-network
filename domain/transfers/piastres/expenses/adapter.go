package expenses

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToExpense converts js data to a Expense instance
func (app *adapter) ToExpense(js []byte) (Expense, error) {
	ins := new(expense)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToJSON converts a expense instance to js data
func (app *adapter) ToJSON(expense Expense) ([]byte, error) {
	return json.Marshal(expense)
}
