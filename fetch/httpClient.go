package fetch

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"

	"github.com/EntryDSM/Corgi/config"
)

var CorgiHttpClient *http.Client

func InitializeHttpClient() {
	caCert, err := ioutil.ReadFile(config.CorgiConfig.ServerCertFilePath)
	if err != nil {
		panic(err.Error())
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	CorgiHttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}
}
