package links

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	transfer_link "github.com/xmn-services/rod-network/domain/transfers/piastres/links"
)

type repository struct {
	builder         Builder
	blockRepository blocks.Repository
	trRepository    transfer_link.Repository
}

func createRepository(
	builder Builder,
	blockRepository blocks.Repository,
	trRepository transfer_link.Repository,
) Repository {
	out := repository{
		builder:         builder,
		blockRepository: blockRepository,
		trRepository:    trRepository,
	}

	return &out
}

// Retrieve retrieves a link by hash
func (app *repository) Retrieve(hsh hash.Hash) (Link, error) {
	trLink, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	nextHash := trLink.Next()
	next, err := app.blockRepository.Retrieve(nextHash)
	if err != nil {
		return nil, err
	}

	prevLink := trLink.PreviousLink()
	index := trLink.Index()
	createdOn := trLink.CreatedOn()
	return app.builder.Create().
		WithNext(next).
		WithPreviousLink(prevLink).
		WithIndex(index).
		CreatedOn(createdOn).
		Now()
}
