package fetch

import (
	"net/http"
	"net/url"
	"path"

	"github.com/EntryDSM/Corgi/config"
)

var AuthenticatedReqeust *http.Request

func InitializeRequestWithAuth() {
	repositoryURL, err := url.Parse(config.CorgiConfig.Registry)
	if err != nil {
		panic("Malformed Repository URL.")
	}

	repositoryURL.Path = path.Join(repositoryURL.Path, resourceOfRepositoryList)
	AuthenticatedReqeust, err := http.NewRequest("", repositoryURL.Path, nil)
	if err != nil {
		panic(err.Error())
	}

	AuthenticatedReqeust.SetBasicAuth(config.CorgiConfig.User, config.CorgiConfig.Password)
}
