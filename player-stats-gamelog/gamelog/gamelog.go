package gamelog

import (
	"encoding/json"
	"log"
)

// PlayerStatsGameLog json struct to represent /api/v1/people/XXXXXXXX/stats?stats=gameLog
type PlayerStatsGameLog struct {
	Copyright string `json:"copyright"`
	Stats     []struct {
		Type struct {
			DisplayName string `json:"displayName"`
		} `json:"type"`
		Splits []struct {
			Season string `json:"season"`
			Stat   struct {
				TimeOnIce            string  `json:"timeOnIce"`
				Assists              int     `json:"assists"`
				Goals                int     `json:"goals"`
				Pim                  int     `json:"pim"`
				Shots                int     `json:"shots"`
				Games                int     `json:"games"`
				Hits                 int     `json:"hits"`
				PowerPlayGoals       int     `json:"powerPlayGoals"`
				PowerPlayPoints      int     `json:"powerPlayPoints"`
				PowerPlayTimeOnIce   string  `json:"powerPlayTimeOnIce"`
				EvenTimeOnIce        string  `json:"evenTimeOnIce"`
				PenaltyMinutes       string  `json:"penaltyMinutes"`
				ShotPct              float64 `json:"shotPct"`
				GameWinningGoals     int     `json:"gameWinningGoals"`
				OverTimeGoals        int     `json:"overTimeGoals"`
				ShortHandedGoals     int     `json:"shortHandedGoals"`
				ShortHandedPoints    int     `json:"shortHandedPoints"`
				ShortHandedTimeOnIce string  `json:"shortHandedTimeOnIce"`
				Blocked              int     `json:"blocked"`
				PlusMinus            int     `json:"plusMinus"`
				Points               int     `json:"points"`
				Shifts               int     `json:"shifts"`
			} `json:"stat"`
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			Opponent struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"opponent"`
			Date   string `json:"date"`
			IsHome bool   `json:"isHome"`
			IsWin  bool   `json:"isWin"`
			IsOT   bool   `json:"isOT"`
			Game   struct {
				GamePk  int    `json:"gamePk"`
				Link    string `json:"link"`
				Content struct {
					Link string `json:"link"`
				} `json:"content"`
			} `json:"game"`
		} `json:"splits"`
	} `json:"stats"`
}

// Exec does stuff
func Exec(s PlayerStatsGameLog) {
	b, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
	}
	log.Println(b)
}
