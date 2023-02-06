package signer

import "fmt"

type Signer interface {
	Sign(data Data) (Signed, error)
	Restore(signed Signed) (Data, error)
}

type Data interface {
	Map() map[string]interface{}
}

type Signed interface {
	fmt.Stringer
}
