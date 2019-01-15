package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/trelore/coeus/game-drawerv2/drawer"
)

var (
	inDir  = "./statsapi/api/v1/game/"
	outDir = "./out/"
)

func main() {
	log.Printf("service starting; in %s\n", inDir)

	var filter string
	flag.StringVar(&filter, "filter", "", "type of play to make heatmap for")
	flag.Parse()
	log.Printf("filter is %s", filter)

	dir, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatalf("read dir: %s", err)
	}
	for _, d := range dir {
		fileName := inDir + d.Name() + "/feed/live.raw" // needs changing (can glob be tighter?)
		var r drawer.RawResponse
		err := extractDataFromJSONFile(&r, fileName)
		if err != nil {
			log.Fatal(err)
		}

		os.Mkdir(outDir, 0700)
		gameID := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		outFile := outDir + gameID + ".png"
		out, err := os.Create(outFile)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		log.Printf("writing to %s", outFile)
		err = drawer.Exec(out, r, filter)
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
