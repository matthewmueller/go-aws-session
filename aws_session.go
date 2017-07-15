package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// New session
func New(id string, secret string, region string) (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(id, secret, ""),
		Region:      aws.String(region),
	})
}

// Must create a new session
func Must(id string, secret string, region string) *session.Session {
	return session.Must(New(id, secret, region))
}
