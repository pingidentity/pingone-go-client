// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CheckCreated(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error) {
	checkAPISuccessWithResponseObject(t, responseDataObj, expectedResponseDataType, httpResponse, httpError, http.StatusCreated)
}

func CheckFound(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error) {
	checkAPISuccessWithResponseObject(t, responseDataObj, expectedResponseDataType, httpResponse, httpError, http.StatusOK)
}

func CheckReplaced(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error) {
	checkAPISuccessWithResponseObject(t, responseDataObj, expectedResponseDataType, httpResponse, httpError, http.StatusOK)
}

func CheckDeleted(t *testing.T, httpResponse *http.Response, httpError error) {
	checkAPISuccess(t, httpResponse, httpError, http.StatusNoContent)
}

func CheckNotFound(t *testing.T, responseDataObj any, httpResponse *http.Response, httpError error) {
	checkAPIFailure(t, responseDataObj, httpResponse, httpError, http.StatusNotFound)
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

func checkAPIFailure(t *testing.T, responseDataObj any, httpResponse *http.Response, httpError error, expectedStatusCode int) {
	require.NotNil(t, httpError)
	require.Nil(t, responseDataObj)
	require.NotNil(t, httpResponse)
	assert.Equal(t, expectedStatusCode, httpResponse.StatusCode)
}

func CheckPingOneAPIErrorResponse(t *testing.T, httpError error, expectedErrorType any, expectedErrorMessageRegex *regexp.Regexp) {
	require.IsType(t, expectedErrorType, httpError)

	// assert.NotEmpty(t, errorModel.Id)
	// assert.Equal(t, expectedErrorCode, *errorModel.Code)
	// assert.NotEmpty(t, errorModel.Message)
	// assert.Regexp(t, expectedErrorMessageRegex, *errorModel.Message)
}
