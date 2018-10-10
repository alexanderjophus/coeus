package main

import (
	"reflect"
	"testing"
)

func TestAddStars(t *testing.T) {
	td := []struct {
		name           string
		output         Output
		decisions      Decisions
		expectedOutput Output
	}{
		{
			name:   "Add stars to empty output",
			output: Output{},
			decisions: Decisions{
				FirstStar: Star{
					ID:       8470860,
					FullName: "Jaroslav Halak",
					Link:     "/api/v1/people/8470860",
				},
				SecondStar: Star{
					ID:       8473419,
					FullName: "Brad Marchand",
					Link:     "/api/v1/people/8473419",
				},
				ThirdStar: Star{
					ID:       8477956,
					FullName: "David Pastrnak",
					Link:     "/api/v1/people/8477956",
				},
			},
			expectedOutput: Output{
				FirstStars: []StarCount{
					StarCount{
						ID:       8470860,
						FullName: "Jaroslav Halak",
						Link:     "/api/v1/people/8470860",
						Count:    1,
					},
				},
				SecondStars: []StarCount{
					StarCount{
						ID:       8473419,
						FullName: "Brad Marchand",
						Link:     "/api/v1/people/8473419",
						Count:    1,
					},
				},
				ThirdStars: []StarCount{
					StarCount{
						ID:       8477956,
						FullName: "David Pastrnak",
						Link:     "/api/v1/people/8477956",
						Count:    1,
					},
				},
			},
		},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			tc.output.addStars(tc.decisions)
			if !reflect.DeepEqual(tc.output, tc.expectedOutput) {
				t.Errorf("expected %v\ngot %v", tc.expectedOutput, tc.output)
			}
		})
	}
}
