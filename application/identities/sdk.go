package identities

import "github.com/xmn-services/rod-network/domain/memory/identities"

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithSeed(seed string) UpdateBuilder
	WithName(name string) UpdateBuilder
	WithPassword(password string) UpdateBuilder
	WithRoot(root string) UpdateBuilder
	WithIdentity(identity identities.Identity) UpdateBuilder
	Now() (Update, error)
}

// Update represents an identity update
type Update interface {
	Seed() string
	Name() string
	Password() string
	Root() string
	Identity() identities.Identity
}

// Application represents an identity application
type Application interface {
	Create() error
	Update(update Update) error
	Retrieve() (identities.Identity, error)
	Delete() error
}
