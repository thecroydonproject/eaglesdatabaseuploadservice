//package data contains data definitions
package data


//Game describes results of football game
type Game struct {

	Gamedate string `json:"gamedate"`
	Team string `json:"team"`
	Awayorhome string `json:"awayorhome"`
	Competition string `json:"competition"`
	Result string `json:"result"`
	Score1 string `json:"score1"`
	Score2 string `json:"score2"`
}
