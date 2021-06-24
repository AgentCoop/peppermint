package utils

import (
	"fmt"
	"strings"
)

func Grpc_ExtractServerShortName(fullName string) string {
	parts := strings.Split(fullName, ".")
	if len(parts) == 0 {
		panic(fmt.Errorf("invalid gRPC server name %s", fullName))
	}
	return parts[len(parts) - 1]
}

func Grpc_MethodToServiceName(method string) string {
	parts := strings.Split(method, "/")
	return parts[1]
}
