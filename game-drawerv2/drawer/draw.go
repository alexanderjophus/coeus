package drawer

import (
	"image"
	"image/png"
	"io"
	"log"

	"github.com/dustin/go-heatmap"
	"github.com/dustin/go-heatmap/schemes"
)

func Exec(w io.Writer, r RawResponse, filter string) error {
	points := []heatmap.DataPoint{}
	for _, play := range r.LiveData.Plays.AllPlays {
		if filter == "" || filter == play.Result.EventTypeID {
			points = append(points, heatmap.P(play.Coordinates.X, play.Coordinates.Y))
		}
	}
	scheme := schemes.AlphaFire
	img := heatmap.Heatmap(image.Rect(0, 0, 850, 2000), points, 200, 128, scheme)

	log.Println("writing heatmap")
	png.Encode(w, img)

	return nil
}
