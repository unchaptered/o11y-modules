package factory

import (
	awssessionwrapper "o11y-modules/sdk/aws_session_wrapper"
)

type SessionFactory struct{}

func (a *SessionFactory) GetAWSSession() awssessionwrapper.IAWSSession {
	return &awssessionwrapper.AWSSession{}
}