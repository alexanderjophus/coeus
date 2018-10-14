package main

import (
	"log"
)

var (
	inDir  = "./statsapi/api/v1/game/"
	outDir = "./out/"
)

func main() {
	log.Printf("service starting. in %s. out %s\n", inDir, outDir)
	Exec(outDir+"stars.csv", inDir)
}
