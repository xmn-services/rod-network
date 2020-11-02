package addresses

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToAddress converts a JSON to an address
func (app *adapter) ToAddress(js *JSONAddress) (Address, error) {
	return createAddressFromJSON(js)
}

// ToJSON converts an address to JSON
func (app *adapter) ToJSON(address Address) *JSONAddress {
	return createJSONAddressFromAddress(address)
}
