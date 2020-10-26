package peers

import (
	"strings"

	"github.com/xmn-services/rod-network/libs/file"
)

type service struct {
	adapter          Adapter
	fileService      file.Service
	filePathWithName string
}

func createService(
	adapter Adapter,
	fileService file.Service,
	filePathWithName string,
) Service {
	out := service{
		adapter:          adapter,
		fileService:      fileService,
		filePathWithName: filePathWithName,
	}

	return &out
}

// Save save peers
func (app *service) Save(peers Peers) error {
	urls := app.adapter.ToURLs(peers)
	str := strings.Join(urls, "\n")
	return app.fileService.Save(app.filePathWithName, []byte(str))
}
