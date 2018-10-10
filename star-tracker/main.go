package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

const inDir = "/pfs/statsapi/api/v1/game/"
const outDir = "/pfs/out/"

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

// this really wants to be a csv like
// ID, Link, FullName, FirstStars, SecondStars, ThirdStars
// 8477956, /api/v1/people/8477956, David Pastrnak, 0, 0, 1

// Output is the csv file we ammend to track star of the game players
type Output struct {
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

func main() {
	os.Mkdir(outDir, 0700)

	stars, err := extractDataFromCSVFile(outDir + "stars.csv")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reading dir  %s\n", inDir)

	dir, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatalf("read dir: %s", err)
	}

	for _, d := range dir {
		fileName := inDir + d.Name() + "/feed/live"
		var liveFeed LiveFeedData
		err := extractDataFromJSONFile(&liveFeed, fileName)
		if err != nil {
			log.Fatal(err)
		}

		// inefficient just wanting to get this working to move onto more meaningful transformations
		// also a defect here, needs splitting out into a func and testing independently
		foundFirst, foundSecond, foundThird := false, false, false
		for _, star := range stars {
			if star.ID == liveFeed.LiveData.Decisions.FirstStar.ID {
				star.FirstStar++
				foundFirst = true
			}
			if star.ID == liveFeed.LiveData.Decisions.SecondStar.ID {
				star.SecondStar++
				foundSecond = true
			}
			if star.ID == liveFeed.LiveData.Decisions.ThirdStar.ID {
				star.ThirdStar++
				foundThird = true
			}
		}
		if !foundFirst {
			stars = append(stars, StarCount{
				ID:         liveFeed.LiveData.Decisions.FirstStar.ID,
				Link:       liveFeed.LiveData.Decisions.FirstStar.Link,
				FullName:   liveFeed.LiveData.Decisions.FirstStar.FullName,
				FirstStar:  1,
				SecondStar: 0,
				ThirdStar:  0,
			})
		}
		if !foundSecond {
			stars = append(stars, StarCount{
				ID:         liveFeed.LiveData.Decisions.SecondStar.ID,
				Link:       liveFeed.LiveData.Decisions.SecondStar.Link,
				FullName:   liveFeed.LiveData.Decisions.SecondStar.FullName,
				FirstStar:  0,
				SecondStar: 1,
				ThirdStar:  0,
			})
		}
		if !foundThird {
			stars = append(stars, StarCount{
				ID:         liveFeed.LiveData.Decisions.ThirdStar.ID,
				Link:       liveFeed.LiveData.Decisions.ThirdStar.Link,
				FullName:   liveFeed.LiveData.Decisions.ThirdStar.FullName,
				FirstStar:  0,
				SecondStar: 0,
				ThirdStar:  1,
			})
		}
	}
	//write stars back out to file
	f, err := os.OpenFile(outDir+"stars.csv", os.O_RDWR, os.ModeExclusive)
	if err != nil {
		log.Fatal(err)
	}
	err = gocsv.MarshalCSVWithoutHeaders(stars, gocsv.NewSafeCSVWriter(csv.NewWriter(f)))
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
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

func extractDataFromCSVFile(src string) ([]StarCount, error) {
	log.Printf("Reading %s\n", src)
	source, err := os.Create(src)
	if err != nil {
		return nil, err
	}
	defer source.Close()

	var stars []StarCount

	f, err := source.Stat()
	if err != nil {
		return nil, err
	}
	if f.Size() == 0 {
		return stars, nil
	}

	if err := gocsv.UnmarshalWithoutHeaders(source, &stars); err != nil {
		return nil, err
	}
	return stars, nil
}
