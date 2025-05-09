package acctest

import (
	"github.com/pingidentity/pingone-go-client/pingone"
)

func TestClient() (*pingone.APIClient, error) {
	configuration := pingone.NewConfiguration()
	configuration.AppendUserAgent("testing")
	configuration.Debug = true
	return pingone.NewAPIClient(configuration)
}
