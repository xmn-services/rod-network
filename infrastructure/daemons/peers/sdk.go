package peers

// Application represents a peer application
type Application interface {
	Start() error
	Stop() error
}
