module github.com/AgentCoop/peppermint

go 1.15

replace github.com/AgentCoop/go-work => /home/pihpah/go/src/github.com/AgentCoop/go-work

require (
	github.com/AgentCoop/go-work v0.0.1
	github.com/golang/protobuf v1.5.2
	github.com/jessevdk/go-flags v1.4.0
	github.com/mattn/go-sqlite3 v1.14.6
	google.golang.org/grpc v1.35.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.26.0
)
