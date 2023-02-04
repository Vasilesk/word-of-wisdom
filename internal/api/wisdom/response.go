package wisdom

type ResponseRandom struct {
	Wisdom
}

type Wisdom struct {
	Text string `json:"text"`
}
