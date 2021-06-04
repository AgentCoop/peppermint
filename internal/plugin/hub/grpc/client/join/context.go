package join

type joinContext struct {
	secret string
	nodeTags []string
	encKey []byte
}
