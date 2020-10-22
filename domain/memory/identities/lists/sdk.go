package lists

import lists "github.com/xmn-services/rod-network/domain/memory/identities/lists/list"

// Lists represents lists
type Lists interface {
	All() []lists.List
}
