package identities

import (
	"github.com/xmn-services/rod-network/application/identities/buckets"
	"github.com/xmn-services/rod-network/application/identities/piastres"
	"github.com/xmn-services/rod-network/application/identities/servers"
	"github.com/xmn-services/rod-network/domain/memory/identities"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// Application represents an identity application
type Application interface {
	Update(update Update) error
	Retrieve() (identities.Identity, error)
	Delete() error
	Bucket() buckets.Application
	Piastre() piastres.Application
	Server() servers.Application
}

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithSeed(seed string) UpdateBuilder
	WithName(name string) UpdateBuilder
	WithPassword(password string) UpdateBuilder
	WithRoot(root string) UpdateBuilder
	Now() (Update, error)
}

// Update represents an identity update
type Update interface {
	HasSeed() bool
	Seed() string
	HasName() bool
	Name() string
	HasPassword() bool
	Password() string
	HasRoot() bool
	Root() string
}
