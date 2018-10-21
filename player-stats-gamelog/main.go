package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexanderjosephtrelore/coeus/player-stats-gamelog/gamelog"
)

var (
	inDir = "./statsapi/api/v1/people/"
)

func main() {
	log.Printf("service starting; in %s\n", inDir)

	dir, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatalf("read dir: %s", err)
	}
	for _, d := range dir {
		fileName := inDir + d.Name() + "/stats/stats_gameLog.raw"
		var s gamelog.PlayerStatsGameLog
		err := extractDataFromJSONFile(&s, fileName)
		if err != nil {
			log.Fatal(err)
		}

		os.Mkdir("out", 0700)
		fName := fmt.Sprintf("out/%s.csv", d.Name())
		f, err := os.Create(fName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = gamelog.Exec(s, f)
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
