module some-project-protocol

go 1.14

require (
	github.com/cenkalti/backoff v2.1.1+incompatible // indirect
	github.com/go-acme/lego v2.5.0+incompatible // indirect
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.4 // indirect
	github.com/ijc/Gotty v0.0.0-20170406111628-a8b993ba6abd // indirect
	github.com/micro/micro/v2 v2.9.2-0.20200724154133-c8dfef1004ab // indirect
	github.com/pquerna/otp v1.2.0 // indirect
	google.golang.org/genproto v0.0.0-20200726014623-da3ae01ef02d // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
