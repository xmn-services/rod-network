package genesis

// Application represents the genesis application
type Application interface {
	Current() Current
}

// Current represents the current application
type Current interface {
	Init(
		name string,
		password string,
		seed string,
		walletName string,
		amountUnits uint64,
		blockDifficultyBase uint,
		blockDifficultyIncreasePerTrx float64,
		linkDifficulty uint,
	) error
}
