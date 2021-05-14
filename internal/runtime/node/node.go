package node

type Node interface {
	GetEncKey()
}

type node struct {
	encKey []byte
}