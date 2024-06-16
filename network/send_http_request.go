package network

import (
	"io"
	"net/http"
)

func SendHttpRequest(
	endpoint string,
) ([]byte, error) {

	var httpResponse *http.Response
	var httpError error
	httpResponse, httpError = http.Get(endpoint)
	if httpError != nil {
		return nil, httpError
	}
	
	defer httpResponse.Body.Close()

	var ioBody []byte
	var ioError error
	ioBody, ioError = io.ReadAll(httpResponse.Body)
	if ioError != nil {
		return nil, ioError
	}

	return ioBody, nil
}