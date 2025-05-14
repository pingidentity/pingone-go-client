package acctest

import (
	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/pingone"
)

func TestClient(svcConfig *config.Configuration) (*pingone.APIClient, error) {
	configuration := pingone.NewConfiguration(svcConfig)
	configuration.AppendUserAgent("testing")
	return pingone.NewAPIClient(configuration)
}
