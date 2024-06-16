package network

import "testing"

func Test_SendHttpRequest_InvalidEndpoint(t *testing.T) {
	endpoint := "wrong_endpoint"

	_, httpError := SendHttpRequest(endpoint)
	if httpError == nil {
		t.Error(
			"SendHttpRequest는 유효하지 않은 Endpoint에서 에러를 발생시킵니다.",
			httpError.Error(),
		)
	}
}

func Test_SendHttpRequest_ValidEndpoint(t *testing.T) {
	endpoint := "https://naver.com"

	httpResponse, _ := SendHttpRequest(endpoint)
	if httpResponse == nil {
		t.Error(
			"SnedHttpRequest는 유효한 httpResponseBody를 응답해야 합니다.",
			httpResponse,
		)
	}
}