package awssessionwrapper

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)


type AWSSession struct {}

func (s AWSSession) ConnectProxy(credentialInfo AWSCredentails) (*session.Session, error) {

	var sessionData *session.Session

	switch credentialInfo.Type {
	case AWSCredentialTypeProfile:
		sessionData = s.ConnectSessionByProfile(credentialInfo)
	case AWSCredentialTypeAccessKey:
		sessionData = s.ConnectSessionByCredentials(credentialInfo)
	default:
		return nil, fmt.Errorf("Unsupported AWSCredentials.Type")
	}

	return sessionData, nil
}

func (s AWSSession) ConnectSessionByProfile(credentialInfo AWSCredentails) (*session.Session) {
	sessionData := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
				Profile: credentialInfo.ProfileName,
				Config: aws.Config{
					Region: aws.String(credentialInfo.Region),
				},
			},
		),
	)

	return sessionData
}

func (s AWSSession) ConnectSessionByCredentials(credentialInfo AWSCredentails) (*session.Session) {
	sessionData := session.Must(
		session.NewSession(
			&aws.Config{
				Credentials: credentials.NewStaticCredentials(
					credentialInfo.AccessKeyID,
					credentialInfo.SecretAccessKey,
					"",
				),
				Region: aws.String(credentialInfo.Region),
			},
		),
	)
	
	return sessionData
}