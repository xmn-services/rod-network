package expenses

// JSONExpense represents a jsonExpense instance
type JSONExpense struct {
	Content    *JSONContent `json:"content"`
	Signatures []string     `json:"signatures"`
}

func createJSONExpenseFromExpense(expense Expense) *JSONExpense {
	content := createJSONContentFromContent(expense.Content())
	signatures := []string{}
	sigs := expense.Signatures()
	for _, oneSig := range sigs {
		signatures = append(signatures, oneSig.String())
	}

	return createJSONExpense(content, signatures)
}

func createJSONExpense(
	content *JSONContent,
	signatures []string,
) *JSONExpense {
	out := JSONExpense{
		Content:    content,
		Signatures: signatures,
	}

	return &out
}
