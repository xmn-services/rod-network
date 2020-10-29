package countries

import "github.com/xmn-services/rod-network/libs/entities"

// Country represents a contry
type Country interface {
	entities.Immutable
	Name() string
	Code() string
}
