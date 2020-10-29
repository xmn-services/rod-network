package persons

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/exchanges/methods/persons/areas"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Person represents an in-person cash transfer
type Person interface {
	entities.Immutable
	Area() areas.Area
	Time() time.Time
	MaxDelay() time.Duration
}
