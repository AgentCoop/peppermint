package chello

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInvalidCreds  = status.Error(codes.PermissionDenied, "invalid join credentials provided")
	errAlreadyJoined = status.Error(codes.PermissionDenied, "already joined, disjoin first")
)
