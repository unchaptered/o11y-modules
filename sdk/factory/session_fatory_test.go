package factory

import (
	awssessionwrapper "o11y-modules/sdk/aws_session_wrapper"
	"testing"
)

func TestSessionFactory(t *testing.T) {
	sessionFactory := SessionFactory{}
	session := sessionFactory.GetAWSSession()

	_, isOk := session.(*awssessionwrapper.AWSSession)
	if !isOk {
		t.Error(
			"SessionFactory{}.GetAWSSession()가 awssessionwrapper.AWSSession 리턴하지 않습니다.",
		)
	}
}