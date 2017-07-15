# AWS Session

Convenience wrapper around the AWS Session SDK. 

## Install

```sh
go get github.com/matthewmueller/go-aws-session
```

## Create a new session

```go
// create it
session := aws_session.New("ACCESS_ID", "SECRET_KEY", "us-east-1")

// use it
s3Svc := s3.New(sess)
manager := s3manager.NewUploaderWithClient(s3Svc)
```

## License

MIT