package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

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

type Stars struct {
	Stars []StarCount
}

// StarCount has the info of how many times that player got that # star of the game
type StarCount struct {
	ID         int
	Link       string
	FullName   string
	FirstStar  int
	SecondStar int
	ThirdStar  int
}

func Exec(totalLoc string, id string) {
	os.Mkdir(outDir, 0700)

	s, err := extractDataFromCSVFile(totalLoc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reading dir  %s\n", id)

	dir, err := ioutil.ReadDir(id)
	if err != nil {
		log.Fatalf("read dir: %s", err)
	}

	for _, d := range dir {
		fileName := id + d.Name() + "/feed/live"
		var liveFeed LiveFeedData
		err := extractDataFromJSONFile(&liveFeed, fileName)
		if err != nil {
			log.Fatal(err)
		}

		s.update(liveFeed.LiveData.Decisions)

	}
	//write stars back out to file
	f, err := os.OpenFile(outDir+"stars.csv", os.O_RDWR, os.ModeExclusive)
	if err != nil {
		log.Fatal(err)
	}
	err = gocsv.MarshalCSVWithoutHeaders(s.Stars, gocsv.NewSafeCSVWriter(csv.NewWriter(f)))
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
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

func extractDataFromJSONFile(v interface{}, src string) error {
	log.Printf("Reading %s\n", src)
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	if err := json.NewDecoder(source).Decode(&v); err != nil {
		return err
	}
	return nil
}

func extractDataFromCSVFile(src string) (*Stars, error) {
	log.Printf("Reading %s\n", src)
	source, err := os.Open(src)
	if err != nil {
		source, err = os.Create(src)
		if err != nil {
			return nil, err
		}
	}
	defer source.Close()

	var stars []StarCount
	s := Stars{
		Stars: stars,
	}

	f, err := source.Stat()
	if err != nil {
		return nil, err
	}
	if f.Size() == 0 {
		return &s, nil
	}

	err = gocsv.UnmarshalWithoutHeaders(source, &stars)
	if err != nil {
		return nil, err
	}

	s.Stars = stars
	log.Printf("Found %d stars", len(s.Stars))
	return &s, nil
}
