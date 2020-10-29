package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// needs to be var as we use ldflags to make this
// /pfs/out etc
var (
	inDir  = "./statsapi/api/v1/people/"
	outDir = "./out"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Printf("service starting; in %s\n", inDir)

	err := filepath.Walk(inDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}

		var s PlayerStatsGameLog
		err = extractDataFromJSONFile(&s, path)
		if err != nil {
			return err
		}

		err = os.MkdirAll(outDir, 0700)
		if err != nil {
			return err
		}
		fName := fmt.Sprintf("%s/%s.csv", outDir, parsePlayerIDFromPath(path))
		f, err := os.Create(fName)
		if err != nil {
			return err
		}
		defer f.Close()

		err = Exec(s, f)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func extractDataFromJSONFile(v interface{}, fileName string) error {
	log.Printf("Reading %s\n", fileName)
	source, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer source.Close()

	if err := json.NewDecoder(source).Decode(&v); err != nil {
		return err
	}
	return nil
}

func parsePlayerIDFromPath(path string) (id string) {
	path = strings.TrimPrefix(path, inDir)
	paths := strings.Split(path, "/")
	return paths[0]
}
