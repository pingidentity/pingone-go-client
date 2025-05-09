package acctest

import (
	"context"

	"github.com/pingidentity/pingone-go-client/pingone"
)

func TestClient(ctx context.Context) *pingone.APIClient {
	configuration := pingone.NewConfiguration()
	configuration.AppendUserAgent("testing")
	configuration.Debug = true
	return pingone.NewAPIClient(configuration)
}
