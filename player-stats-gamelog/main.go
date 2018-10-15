package main

import (
	"encoding/json"
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
		fileName := inDir + d.Name() + "/stats?stats=gameLog"
		var s gamelog.PlayerStatsGameLog
		err := extractDataFromJSONFile(&s, fileName)
		if err != nil {
			log.Fatal(err)
		}

		gamelog.Exec(s)
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
