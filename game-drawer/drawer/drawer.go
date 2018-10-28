package drawer

import (
	"io"
	"log"
	"os"
	"strconv"

	svg "github.com/ajstarks/svgo"
)

// Exec does stuff
func Exec(w io.Writer, events [][]string) error {
	canvas := svg.New(w)
	canvas.Start(2000, 850)
	defer canvas.End()

	err := addBase(canvas.Writer)
	if err != nil {
		return err
	}
	for i := range events {
		if events[i][2] == "GOAL" {
			xi, yi, err := coordinates(events[i])
			if err != nil {
				log.Println(err)
				continue
			}
			addEvent(canvas, xi, yi, "G")
		}
	}
	return nil
}

func addBase(w io.Writer) error {
	f, err := os.Open("rink.svg")
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(w, f)
	return nil
}

func addEvent(canvas *svg.SVG, x, y int, symbol string) {
	xs := []int{20, 10, -10, -20, -10, 10}
	ys := []int{0, 18, 18, 0, -18, -18}
	for i := range xs {
		xs[i] += x
	}
	for i := range ys {
		ys[i] += y
	}
	canvas.Polygon(xs, ys, "stroke:black")
	canvas.Text(x-6, y+4, symbol, "stroke:white")
}

func coordinates(in []string) (int, int, error) {
	x, err := strconv.ParseFloat(in[0], 64)
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.ParseFloat(in[1], 64)
	if err != nil {
		return 0, 0, err
	}
	xi := (int(100 + x)) * 10
	yi := (int(42.5 + y)) * 10
	return xi, yi, nil
}
