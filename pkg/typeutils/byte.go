package typeutils

type Byter interface {
	Bytes() []byte
}

type bytesHolder []byte

func NewByter(bb []byte) Byter {
	return bytesHolder(bb)
}

func (b bytesHolder) Bytes() []byte {
	return b
}
