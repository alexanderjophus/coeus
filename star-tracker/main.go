package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

func main() {
	os.Mkdir(outDir, 0700)

	log.Printf("Reading directory %s\n", inDir)

	dir, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatalf("read dir: %s", err)
	}

	for _, d := range dir {
		fileName := inDir + d.Name() + "/feed/live"
		err := readFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func readFile(src string) error {
	log.Printf("Reading %s\n", src)
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	var liveFeed LiveFeedData
	if err := json.NewDecoder(source).Decode(&liveFeed); err != nil {
		return err
	}

	log.Println(liveFeed)
	return nil
}
