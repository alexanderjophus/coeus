package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

var (
	inDir  = "/pfs/statsapi/api/v1/game/"
	outDir = "/pfs/out/"
)

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

func main() {
	Exec(outDir+"stars.csv", inDir)
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
	return &s, nil
}
