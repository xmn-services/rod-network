package servers

// Application represents a server application
type Application interface {
	Start() error
	Stop() error
}
