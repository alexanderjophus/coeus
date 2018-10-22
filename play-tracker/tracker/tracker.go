package tracker

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
)

type LiveFeedResponse struct {
	Copyright string   `json:"copyright"`
	GamePk    int      `json:"gamePk"`
	Link      string   `json:"link"`
	MetaData  MetaData `json:"metaData"`
	LiveData  LiveFeed `json:"liveData"`
}

type MetaData struct {
	Wait      int    `json:"wait"`
	TimeStamp string `json:"timeStamp"`
}

type LiveFeed struct {
	Plays Plays `json:"plays"`
}

type Plays struct {
	AllPlays      []Play          `json:"allPlays"`
	ScoringPlays  []int           `json:"scoringPlays"`
	PenaltyPlays  []int           `json:"penaltyPlays"`
	PlaysByPeriod []PlaysByPeriod `json:"playsByPeriod"`
}

type Play struct {
	Players []Player `json:"players"`
	Result  struct {
		Event       string `json:"event"`
		EventCode   string `json:"eventCode"`
		EventTypeID string `json:"eventTypeId"`
		Description string `json:"description"`
	} `json:"result"`
	About struct {
		EventIdx            int       `json:"eventIdx"`
		EventID             int       `json:"eventId"`
		Period              int       `json:"period"`
		PeriodType          string    `json:"periodType"`
		OrdinalNum          string    `json:"ordinalNum"`
		PeriodTime          string    `json:"periodTime"`
		PeriodTimeRemaining string    `json:"periodTimeRemaining"`
		DateTime            time.Time `json:"dateTime"`
		Goals               struct {
			Away int `json:"away"`
			Home int `json:"home"`
		} `json:"goals"`
	} `json:"about"`
	Coordinates struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"coordinates"`
	Team struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Link    string `json:"link"`
		TriCode string `json:"triCode"`
	} `json:"team"`
}

type Player struct {
	Player struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"player"`
	PlayerType string `json:"playerType"`
}

type PlaysByPeriod struct {
	StartIndex int   `json:"startIndex"`
	Plays      []int `json:"plays"`
	EndIndex   int   `json:"endIndex"`
}

// Exec does stuff
func Exec(l LiveFeedResponse, w io.Writer) error {
	c := csv.NewWriter(w)
	defer c.Flush()

	c.Write([]string{"X", "Y", "type", "period", "time_remaining"})

	for _, play := range l.LiveData.Plays.AllPlays {
		c.Write([]string{
			fmt.Sprintf("%.1f", play.Coordinates.X),
			fmt.Sprintf("%.1f", play.Coordinates.Y),
			play.Result.EventTypeID,
			strconv.Itoa(play.About.Period),
			play.About.PeriodTimeRemaining,
		})
	}
	return nil
}
