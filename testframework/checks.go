// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// CheckCreated verifies that an API operation successfully created a resource.
// The responseDataObj parameter contains the actual response data, expectedResponseDataType
// specifies the expected type of the response, httpResponse is the HTTP response object,
// and httpError is any error from the API call. This function validates that the operation
// returned HTTP 201 Created status and the response data matches the expected type.
func CheckCreated(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error) {
	checkAPISuccessWithResponseObject(t, responseDataObj, expectedResponseDataType, httpResponse, httpError, http.StatusCreated)
}

// CheckFound verifies that an API operation successfully retrieved a resource.
// The responseDataObj parameter contains the actual response data, expectedResponseDataType
// specifies the expected type of the response, httpResponse is the HTTP response object,
// and httpError is any error from the API call. This function validates that the operation
// returned HTTP 200 OK status and the response data matches the expected type.
func CheckFound(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error) {
	checkAPISuccessWithResponseObject(t, responseDataObj, expectedResponseDataType, httpResponse, httpError, http.StatusOK)
}

// CheckReplaced verifies that an API operation successfully replaced/updated a resource.
// The responseDataObj parameter contains the actual response data, expectedResponseDataType
// specifies the expected type of the response, httpResponse is the HTTP response object,
// and httpError is any error from the API call. This function validates that the operation
// returned HTTP 200 OK status and the response data matches the expected type.
func CheckReplaced(t *testing.T, responseDataObj any, expectedResponseDataType any, httpResponse *http.Response, httpError error) {
	checkAPISuccessWithResponseObject(t, responseDataObj, expectedResponseDataType, httpResponse, httpError, http.StatusOK)
}

// CheckDeleted verifies that an API operation successfully deleted a resource.
// The httpResponse parameter is the HTTP response object and httpError is any error
// from the API call. This function validates that the operation returned HTTP 204 No Content
// status, which is the standard response for successful deletion operations.
func CheckDeleted(t *testing.T, httpResponse *http.Response, httpError error) {
	checkAPISuccess(t, httpResponse, httpError, http.StatusNoContent)
}

// CheckNotFound verifies that an API operation correctly returned a not found error.
// The responseDataObj parameter contains the actual response data, httpResponse is the
// HTTP response object, and httpError is any error from the API call. This function
// validates that the operation returned HTTP 404 Not Found status.
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

// CheckPingOneAPIErrorResponse verifies that the error is of the expected type.
// The expectedErrorMessageRegex parameter is currently unused but reserved for future validation.
func CheckPingOneAPIErrorResponse(t *testing.T, httpError error, expectedErrorType any, _ *regexp.Regexp) {
	require.IsType(t, expectedErrorType, httpError)

	// assert.NotEmpty(t, errorModel.Id)
	// assert.Equal(t, expectedErrorCode, *errorModel.Code)
	// assert.NotEmpty(t, errorModel.Message)
	// assert.Regexp(t, expectedErrorMessageRegex, *errorModel.Message)
}
