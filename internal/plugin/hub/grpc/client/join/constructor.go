package join

func NewJoinContext(secret string, nodeTags []string) *joinContext {
	c := &joinContext{
		secret:   secret,
		nodeTags: nodeTags,
	}
	return c
}
