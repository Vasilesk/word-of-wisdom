package wisdom

import "github.com/vasilesk/word-of-wisdom/internal/repo/wisdomwords"

func mapResponseRandom(w wisdomwords.Wisdom) ResponseRandom {
	return ResponseRandom{Wisdom: Wisdom{Text: w.Text}}
}
