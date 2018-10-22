package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexanderjosephtrelore/coeus/play-tracker/tracker"
)

var (
	inDir  = "./statsapi/api/v1/game/"
	outDir = "./out"
)

func main() {
	log.Printf("service starting; in %s\n", inDir)

	dir, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatalf("read dir: %s", err)
	}
	for _, d := range dir {
		fileName := inDir + d.Name() + "/feed/live.raw" // needs changing (can glob be tighter?)
		var r tracker.LiveFeedResponse
		err := extractDataFromJSONFile(&r, fileName)
		if err != nil {
			log.Fatal(err)
		}

		os.Mkdir(outDir, 0700)
		fName := fmt.Sprintf("%s/%s.csv", outDir, d.Name())
		f, err := os.Create(fName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = tracker.Exec(r, f)
		if err != nil {
			log.Fatal(err)
		}
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
