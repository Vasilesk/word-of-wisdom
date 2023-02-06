package typeutils

type Byter interface {
	Bytes() []byte
}

func NewByter(bb []byte) Byter {
	return bytesHolder(bb)
}

func NewByterFromString(s string) Byter {
	return bytesHolder([]byte(s))
}

type bytesHolder []byte

func (b bytesHolder) Bytes() []byte {
	return b
}
