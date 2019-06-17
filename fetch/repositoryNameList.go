package fetch

import (
	"encoding/json"
	"net/http"
)

type RepositoryList struct {
	Repositories []string `json:"repositories"`
}

const resourceOfRepositoryList = "v2/_catalog"

func GetRepositoryNameList() (*RepositoryList, error) {
	req := http.Request(*AuthenticatedReqeust)
	req.Method = "GET"

	res, err := CorgiHttpClient.Do(&req)
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
