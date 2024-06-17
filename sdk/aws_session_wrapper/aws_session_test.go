package awssessionwrapper

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
)

func TestAWSSessionProperties(t *testing.T) {
	var sess AWSSession

	// Check ConnectProxy Properties
	if _, ok := interface{}(sess).(interface {
		ConnectProxy(AWSCredentails) (*session.Session, error)
	}); !ok {
		t.Errorf("AWSSession does not have method ConnectProxy")
	}

	// Check ConnectProxy Properties
	if _, ok := interface{}(sess).(interface {
		ConnectSessionByProfile(AWSCredentails) (*session.Session)
	}); !ok {
		t.Errorf("AWSSession does not have method ConnectSessionByProfile")
	}

	// Check ConnectProxy Properties
	if _, ok := interface{}(sess).(interface {
		ConnectSessionByCredentials(AWSCredentails) (*session.Session)
	}); !ok {
		t.Errorf("AWSSession does not have method ConnectSessionByCredentials")
	}
}