package client

type RequestHeader interface {

}

type RequestData interface {

}

type Request interface {
	RequestHeader
	RequestData
	ToGrpcRequest() interface{}
}
