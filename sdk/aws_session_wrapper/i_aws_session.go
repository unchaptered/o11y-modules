package awssessionwrapper

import (
	"github.com/aws/aws-sdk-go/aws/session"
)


type AWSCredentialType int
const (
	AWSCredentialTypeProfile AWSCredentialType = iota
	AWSCredentialTypeAccessKey
)

type AWSCredentails struct {
	Type               AWSCredentialType
	ProfileName        string
	AccessKeyID        string
	SecretAccessKey    string
	Region             string
}

type IAWSSession interface {
	ConnectProxy(credentialInfo AWSCredentails) (*session.Session, error)
	ConnectSessionByProfile(credentialInfo AWSCredentails) (*session.Session)
	ConnectSessionByCredentials(credentialInfo AWSCredentails) (*session.Session)
}