package genesis

type block struct {
	base     uint
	increase float64
}

func createBlock(
	base uint,
	increase float64,
) Block {
	out := block{
		base:     base,
		increase: increase,
	}

	return &out
}

// Base returns the base
func (obj *block) Base() uint {
	return obj.base
}

// IncreasePerTrx returns the increase per trx
func (obj *block) IncreasePerTrx() float64 {
	return obj.increase
}
