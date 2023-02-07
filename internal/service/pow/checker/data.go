package checker

import (
	"errors"
	"time"
)

type powData struct {
	Challenge  string
	ValidUntil time.Time
	IP         string
	URI        string
}

func (d powData) Map() map[string]interface{} {
	return map[string]interface{}{
		"challenge":  d.Challenge,
		"validUntil": d.ValidUntil.UnixNano(),
		"ip":         d.IP,
		"uri":        d.URI,
	}
}

func powDataFromMap(m map[string]interface{}) (*powData, error) {
	challengeI, ok := m["challenge"]
	if !ok {
		return nil, errors.New("challenge was not found")
	}

	challenge, ok := challengeI.(string)
	if !ok {
		return nil, errors.New("challenge casting error")
	}

	validUntilI, ok := m["validUntil"]
	if !ok {
		return nil, errors.New("validUntil was not found")
	}

	validUntilInt, ok := validUntilI.(int)
	if !ok {
		return nil, errors.New("validUntil casting error")
	}

	validUntil := time.Date(0, 0, 0, 0, 0, 0, validUntilInt, nil)

	ipI, ok := m["ip"]
	if !ok {
		return nil, errors.New("ip was not found")
	}

	ip, ok := ipI.(string)
	if !ok {
		return nil, errors.New("ip casting error")
	}

	uriI, ok := m["uri"]
	if !ok {
		return nil, errors.New("uri was not found")
	}

	uri, ok := uriI.(string)
	if !ok {
		return nil, errors.New("uri casting error")
	}

	return newPowData(challenge, validUntil, ip, uri), nil
}

func newPowData(challenge string, validUntil time.Time, ip string, uri string) *powData {
	return &powData{
		Challenge:  challenge,
		ValidUntil: validUntil,
		IP:         ip,
		URI:        uri,
	}
}
