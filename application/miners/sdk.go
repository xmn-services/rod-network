package miners

// maxMiningValue represents the max mining value before adding another miner number to the slice
const maxMiningValue = 2147483647

// maxMiningTries represents the max mining characters to try before abandonning
const maxMiningTries = 2147483647

// miningBeginValue represents the first value of the hash that is expected on mining
const miningBeginValue = 0

// Application represents the miner application
type Application interface {
	Current() Current
}

// Current represents the current application
type Current interface {
	Block(trx []string) error
	Link(prevMinedLink string, nextBlock string) error
}
