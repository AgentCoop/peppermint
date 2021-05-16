package client

type ResponseHeader interface {

}

type ResponseData interface {
	Populate()
}

type Response interface {
	ResponseHeader
	ResponseData
}

