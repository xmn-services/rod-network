package peers

import (
	"strings"

	"github.com/xmn-services/rod-network/libs/file"
)

type repository struct {
	adapter          Adapter
	fileRepository   file.Repository
	filePathWithName string
}

func createRepository(
	adapter Adapter,
	fileRepository file.Repository,
	filePathWithName string,
) Repository {
	out := repository{
		adapter:          adapter,
		fileRepository:   fileRepository,
		filePathWithName: filePathWithName,
	}

	return &out
}

// Retrieve retrieve peers
func (app *repository) Retrieve() (Peers, error) {
	data, err := app.fileRepository.Retrieve(app.filePathWithName)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	lines := strings.Split(string(data), "\n")
	for _, oneLine := range lines {
		line := strings.TrimSpace(oneLine)
		if line == "" {
			continue
		}

		urls = append(urls, line)
	}

	return app.adapter.ToPeers(urls)
}
