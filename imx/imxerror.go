package imx

import (
	"fmt"
	"net/http"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
IMXError struct contains the details of the error for the API request.
For more information about the API error codes, see https://docs.x.immutable.com/docs/error-codes/
*/
type IMXError struct {
	HTTPStatusCode int // e.g. 400
	RequestURL     string
	api.APIError
}

/*
NewIMXError returns an IMXError when Api Request fails.
@param httpResponse httpResponse.
@param err Error message.
@return IMXError
*/
func NewIMXError(httpResponse *http.Response, err error) error {
	var apiError api.APIError
	if genericAPIError, ok := err.(*api.GenericOpenAPIError); ok {
		apiError, ok = genericAPIError.Model().(api.APIError)
		if !ok {
			apiError = *api.NewAPIError("", string(genericAPIError.Body()))
		}
	} else {
		apiError = *api.NewAPIError("", err.Error())
	}

	return &IMXError{
		HTTPStatusCode: httpResponse.StatusCode,
		RequestURL:     httpResponse.Request.URL.String(),
		APIError:       apiError,
	}
}

func (e *IMXError) Error() string {
	return fmt.Sprintf("Request URL: %v Status: %v, Message: %v", e.RequestURL, e.HTTPStatusCode, e.Message)
}
