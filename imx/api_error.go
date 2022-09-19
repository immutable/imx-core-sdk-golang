package imx

import (
	"fmt"
	"net/http"
)

type APIError struct {
	httpResponse *http.Response
	err          error
}

/*
NewAPIError returns an APIError when Api Request fails.
@param httpResponse httpResponse.
@param err Error message.
@return APIError
*/
func NewAPIError(httpResponse *http.Response, err error) error {
	return &APIError{
		httpResponse: httpResponse,
		err:          err,
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("error in api request: %v, HTTP response body: %v", e.err, e.httpResponse.Body)
}

func (e *APIError) Response() *http.Response {
	return e.httpResponse
}
