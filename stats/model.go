package stats

type stats struct {
	//UserProfile string
	Name  string `json:"name"`
	Score string `json:"score"`
}

type place struct {
	Stats    stats `json:"stats"`
	Position int   `json:"position"`
}

type top10 struct {
	Top10 []place
}
