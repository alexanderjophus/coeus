package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/trelore/coeus/highlight-generator/thing"
)

var (
	inDir  = "./statsapi/api/v1/game/"
	outDir = "./out/"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Printf("service starting; in %s\n", inDir)

	dir, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatalf("read dir: %s", err)
	}

	os.Mkdir(outDir, 0700)
	outFile := outDir + "highlights.txt"
	out, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	log.Printf("writing to %s", outFile)

	for _, d := range dir {
		fileName := inDir + d.Name() + "/content.raw"
		var r thing.RawResponse
		err := extractDataFromJSONFile(&r, fileName)
		if err != nil {
			log.Fatal(err)
		}

		err = thing.Exec(out, r)
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
