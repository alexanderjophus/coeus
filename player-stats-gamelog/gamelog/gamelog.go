package gamelog

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

// PlayerStatsGameLog json struct to represent /api/v1/people/XXXXXXXX/stats?stats=gameLog
type PlayerStatsGameLog struct {
	Copyright string `json:"copyright"`
	Stats     []Stat `json:"stats"`
}

type Stat struct {
	Type struct {
		DisplayName string `json:"displayName"`
	} `json:"type"`
	Splits []Split `json:"splits"`
}

type Split struct {
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
}

// Exec does stuff
func Exec(s PlayerStatsGameLog, w io.Writer) error {
	c := csv.NewWriter(w)
	defer c.Flush()

	for _, stat := range s.Stats {
		for _, split := range stat.Splits {
			err := c.Write([]string{
				split.Date,
				split.Opponent.Name,
				strconv.FormatBool(split.IsHome),
				split.Stat.TimeOnIce,
				strconv.Itoa(split.Stat.Goals),
				strconv.Itoa(split.Stat.Assists),
				strconv.Itoa(split.Stat.Shots),
				strconv.Itoa(split.Stat.Pim),
				strconv.Itoa(split.Stat.Hits),
				strconv.Itoa(split.Stat.PowerPlayGoals),
				strconv.Itoa(split.Stat.PowerPlayPoints),
				split.Stat.PowerPlayTimeOnIce,
				split.Stat.EvenTimeOnIce,
				split.Stat.PenaltyMinutes,
				strconv.Itoa(split.Stat.GameWinningGoals),
				strconv.Itoa(split.Stat.OverTimeGoals),
				strconv.Itoa(split.Stat.ShortHandedGoals),
				strconv.Itoa(split.Stat.ShortHandedPoints),
				split.Stat.ShortHandedTimeOnIce,
				strconv.Itoa(split.Stat.Blocked),
				strconv.Itoa(split.Stat.PlusMinus),
				strconv.Itoa(split.Stat.Shifts),
			})
			if err != nil {
				log.Printf("failed to write game %s to file", split.Date)
			}
		}
	}
	return nil
}
