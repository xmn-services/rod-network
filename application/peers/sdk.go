package peers

// Application represents the peer application
type Application interface {
	Save(host string, port uint) error
}
