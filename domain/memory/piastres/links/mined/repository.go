package mined

import (
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/links"
	transfer_mined_link "github.com/xmn-services/rod-network/domain/transfers/piastres/links/mined"
)

type repository struct {
	builder        Builder
	linkRepository links.Repository
	trRepository   transfer_mined_link.Repository
}

func createRepository(
	builder Builder,
	linkRepository links.Repository,
	trRepository transfer_mined_link.Repository,
) Repository {
	out := repository{
		builder:        builder,
		linkRepository: linkRepository,
		trRepository:   trRepository,
	}

	return &out
}

// Retrieve retrieves a link by hash
func (app *repository) Retrieve(hash hash.Hash) (Link, error) {
	trLink, err := app.trRepository.Retrieve(hash)
	if err != nil {
		return nil, err
	}

	linkHash := trLink.Link()
	subLink, err := app.linkRepository.Retrieve(linkHash)
	if err != nil {
		return nil, err
	}

	mining := trLink.Mining()
	createdOn := trLink.CreatedOn()
	return app.builder.Create().WithLink(subLink).WithMining(mining).CreatedOn(createdOn).Now()
}
