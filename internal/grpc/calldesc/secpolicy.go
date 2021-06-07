package calldesc

func (s secPolicy) IsSecure() bool {
	return s.e2e_Enc
}

func (s secPolicy) EncKey() []byte {
	return s.encKey
}
