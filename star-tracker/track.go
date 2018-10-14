package main

// LiveFeedData is the response from a /${GAME_ID}/feed/live endpoint
type LiveFeedData struct {
	LiveData LiveData `json:"liveData"`
}

// LiveData holds a bunch of structs like plays and decisions
type LiveData struct {
	Decisions Decisions `json:"decisions"`
}

// Decisions reflects the decisions in the game, stars/winning & losing goalie
type Decisions struct {
	FirstStar  Star `json:"firstStar"`
	SecondStar Star `json:"secondStar"`
	ThirdStar  Star `json:"thirdStar"`
}

// Star holds information about a player
type Star struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Link     string `json:"link"`
}

func (s *Stars) update(decision Decisions) {
	// inefficient just wanting to get this working to move onto more meaningful transformations
	// also a defect here, needs splitting out into a func and testing independently
	foundFirst, foundSecond, foundThird := false, false, false
	for k := range s.Stars {
		if s.Stars[k].ID == decision.FirstStar.ID {
			s.Stars[k].FirstStar++
			foundFirst = true
			continue
		}
		if s.Stars[k].ID == decision.SecondStar.ID {
			s.Stars[k].SecondStar++
			foundSecond = true
			continue
		}
		if s.Stars[k].ID == decision.ThirdStar.ID {
			s.Stars[k].ThirdStar++
			foundThird = true
			continue
		}
	}

	if !foundFirst {
		s.Stars = append(s.Stars, StarCount{
			ID:         decision.FirstStar.ID,
			Link:       decision.FirstStar.Link,
			FullName:   decision.FirstStar.FullName,
			FirstStar:  1,
			SecondStar: 0,
			ThirdStar:  0,
		})
	}
	if !foundSecond {
		s.Stars = append(s.Stars, StarCount{
			ID:         decision.SecondStar.ID,
			Link:       decision.SecondStar.Link,
			FullName:   decision.SecondStar.FullName,
			FirstStar:  0,
			SecondStar: 1,
			ThirdStar:  0,
		})
	}
	if !foundThird {
		s.Stars = append(s.Stars, StarCount{
			ID:         decision.ThirdStar.ID,
			Link:       decision.ThirdStar.Link,
			FullName:   decision.ThirdStar.FullName,
			FirstStar:  0,
			SecondStar: 0,
			ThirdStar:  1,
		})
	}
}
