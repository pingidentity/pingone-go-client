package acctest

import (
	"net/http"
	"testing"

	"github.com/pingidentity/pingone-go-client/pingone"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CheckCreated(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error) {
	checkAPISuccessWithResponseObject(t, responseDataObj, expectedResponseDataType, httpResponse, httpError, 201)
}

func CheckFound(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error) {
	checkAPISuccessWithResponseObject(t, responseDataObj, expectedResponseDataType, httpResponse, httpError, 200)
}

func CheckReplaced(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error) {
	checkAPISuccessWithResponseObject(t, responseDataObj, expectedResponseDataType, httpResponse, httpError, 200)
}

func CheckDeleted(t *testing.T, httpResponse *http.Response, httpError error) {
	checkAPISuccess(t, httpResponse, httpError, 204)
}

func CheckNotFound(t *testing.T, responseDataObj any, httpResponse *http.Response, httpError error) {
	checkAPIFailure(t, responseDataObj, httpResponse, httpError, 404, "404 Not Found")
}

func checkAPISuccess(t *testing.T, httpResponse *http.Response, httpError error, expectedStatusCode int) {
	require.Nil(t, httpError)
	require.NotNil(t, httpResponse)
	assert.Equal(t, expectedStatusCode, httpResponse.StatusCode)
}

func checkAPISuccessWithResponseObject(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error, expectedStatusCode int) {
	checkAPISuccess(t, httpResponse, httpError, expectedStatusCode)
	require.NotNil(t, responseDataObj)

	require.IsType(t, expectedResponseDataType, responseDataObj)
}

func checkAPIFailure(t *testing.T, responseDataObj any, httpResponse *http.Response, httpError error, expectedStatusCode int, expectedErrorString string) {
	require.NotNil(t, httpError)
	require.Nil(t, responseDataObj)
	require.NotNil(t, httpResponse)
	assert.Equal(t, expectedStatusCode, httpResponse.StatusCode)
	assert.EqualError(t, httpError, expectedErrorString)

	require.IsType(t, &pingone.APIError{}, httpError)
	assert.IsType(t, pingone.ErrorResponse{}, httpError.(*pingone.APIError).Model())
}
