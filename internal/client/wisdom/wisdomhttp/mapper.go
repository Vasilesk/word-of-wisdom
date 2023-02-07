package wisdomhttp

import "github.com/vasilesk/word-of-wisdom/internal/client/wisdom"

func mapResponseRandom(r responseRandom) wisdom.ResponseRandom {
	return wisdom.ResponseRandom{Text: r.Text}
}
