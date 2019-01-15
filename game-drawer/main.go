package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/trelore/coeus/game-drawer/drawer"
)

var (
	inDir  = "./play-tracker/"
	outDir = "./out/"
)

func main() {
	log.Printf("service starting; in %s\n", inDir)

	dir, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatalf("read dir: %s", err)
	}
	for _, d := range dir {
		fileName := inDir + d.Name()
		inFile, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()

		r := csv.NewReader(inFile)

		events, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(outDir + strings.TrimSuffix(d.Name(), filepath.Ext(d.Name())) + ".svg")
		if err != nil {
			log.Fatal(err)
		}

		err = drawer.Exec(out, events)
		if err != nil {
			log.Fatal(err)
		}
	}
}
