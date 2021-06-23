package stream

func encLayer(msg interface{}, isSecure bool, key []byte) error {
	if !isSecure { return nil }

	//_, ok := msg.(codec.Packet)
	//if ok { return nil }
	//
	//if len(key) == 0 { return ErrEmptyEncryptionKey }
	//msg = codec.NewPacket(msg, key)
	return nil
}
