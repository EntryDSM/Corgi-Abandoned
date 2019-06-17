package fetch

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/EntryDSM/Corgi/config"
)

type RepositoryList struct {
	Repositories []string `json:"repositories"`
}

const resourceOfRepositoryList = "v2/_catalog"

func GetRepositoryNameList(repositoryURLString string) (*RepositoryList, error) {
	repositoryURL, err := url.Parse(repositoryURLString)
	if err != nil {
		panic("Malformed Repository URL.")
	}

	repositoryURL.Path = path.Join(repositoryURL.Path, resourceOfRepositoryList)
	req, err := http.NewRequest("GET", repositoryURL.Path, nil)
	if err != nil {
		panic(err.Error())
	}

	req.SetBasicAuth(config.CorgiConfig.User, config.CorgiConfig.Password)
	res, err := CorgiHttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var repositoryList *RepositoryList = &RepositoryList{}
	err = json.NewDecoder(res.Body).Decode(repositoryList)
	if err != nil {
		return nil, err
	}

	return repositoryList, nil
}
