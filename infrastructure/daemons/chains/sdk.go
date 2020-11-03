package chains

// Application represents a chain application
type Application interface {
	Start() error
	Stop() error
}
