package expenses

// JSONExpense represents a jsonExpense instance
type JSONExpense struct {
	Content    *JSONContent `json:"content"`
	Signatures [][]string   `json:"signatures"`
}

func createJSONExpenseFromExpense(expense Expense) *JSONExpense {
	content := createJSONContentFromContent(expense.Content())
	signatures := [][]string{}
	sigs := expense.Signatures()
	for _, oneSigs := range sigs {
		signatureList := []string{}
		for _, oneSig := range oneSigs {
			signatureList = append(signatureList, oneSig.String())
		}

		signatures = append(signatures, signatureList)
	}

	return createJSONExpense(content, signatures)
}

func createJSONExpense(
	content *JSONContent,
	signatures [][]string,
) *JSONExpense {
	out := JSONExpense{
		Content:    content,
		Signatures: signatures,
	}

	return &out
}
