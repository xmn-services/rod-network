package daemons

// Builder represents a daemon application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// Application represents a daemon application
type Application interface {
	Start() error
	Stop() error
}
