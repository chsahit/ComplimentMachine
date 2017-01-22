package controllers

type Sentiment struct {
	Trait       string
	Probability Prob   `json:"probability"`
	Label       string `json:"label"`
}

type Prob struct {
	Neg     float64 `json:"neg"`
	Neutral float64 `json:"neutral"`
	Pos     float64 `json:"pos"`
}

type Sentiments []Sentiment

func (slice Sentiments) Len() int {

	return len(slice)
}

func (slice Sentiments) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice Sentiments) Less(i, j int) bool {
	var score1 float64
	var score2 float64
	if slice[i].Probability.Pos > slice[i].Probability.Neutral &&
		slice[i].Probability.Pos > slice[i].Probability.Neg {
		score1 = 100 + slice[i].Probability.Pos
	} else if slice[i].Probability.Neutral > slice[i].Probability.Neg {
		score1 = slice[i].Probability.Neutral
	} else {
		score1 = slice[i].Probability.Neg - 100
	}
        if slice[i].Trait == "adult" || slice[i].Trait == "portrait" {
            score1 = -100;
        }
	if slice[j].Probability.Pos > slice[j].Probability.Neutral &&
		slice[j].Probability.Pos > slice[j].Probability.Neg {
		score2 = 100 + slice[j].Probability.Pos
	} else if slice[j].Probability.Neutral > slice[j].Probability.Neg {
		score2 = slice[j].Probability.Neutral
	} else {
		score2 = slice[j].Probability.Neg - 100
	}
        if slice[j].Trait == "adult" || slice[j].Trait == "portrait" {
            score2 = -100;
        }
	return score1 > score2

}
