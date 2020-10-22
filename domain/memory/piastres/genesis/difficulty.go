package genesis

type difficulty struct {
	block Block
	link  uint
}

func createDifficulty(
	block Block,
	link uint,
) Difficulty {
	out := difficulty{
		block: block,
		link:  link,
	}

	return &out
}

// Block returns the block's difficulty
func (obj *difficulty) Block() Block {
	return obj.block
}

// Link returns the link's difficulty
func (obj *difficulty) Link() uint {
	return obj.link
}
