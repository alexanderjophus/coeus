package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStars(t *testing.T) {
	td := []struct {
		name              string
		decisions         Decisions
		originalStarCount []StarCount
		expectedOutput    []StarCount
	}{
		{
			name:              "Add stars to empty struct",
			decisions:         goldenDecisions(),
			originalStarCount: []StarCount{},
			expectedOutput: func() []StarCount {
				s := goldenStarCount()
				s[0].FirstStar = 1
				s[1].SecondStar = 1
				s[2].ThirdStar = 1
				return s
			}(),
		},
		{
			name:              "Increases star count from 0s",
			decisions:         goldenDecisions(),
			originalStarCount: goldenStarCount(),
			expectedOutput: func() []StarCount {
				s := goldenStarCount()
				s[0].FirstStar = 1
				s[1].SecondStar = 1
				s[2].ThirdStar = 1
				return s
			}(),
		},
		{
			name:      "Increases star count from 1s",
			decisions: goldenDecisions(),
			originalStarCount: func() []StarCount {
				s := goldenStarCount()
				s[0].FirstStar = 1
				s[1].SecondStar = 1
				s[2].ThirdStar = 1
				return s
			}(),
			expectedOutput: func() []StarCount {
				s := goldenStarCount()
				s[0].FirstStar = 2
				s[1].SecondStar = 2
				s[2].ThirdStar = 2
				return s
			}(),
		},
		{
			name: "Increases star count for other stars",
			decisions: func() Decisions {
				d := goldenDecisions()
				d.FirstStar, d.SecondStar, d.ThirdStar = d.SecondStar, d.ThirdStar, d.FirstStar
				return d
			}(),
			originalStarCount: func() []StarCount {
				s := goldenStarCount()
				s[0].FirstStar = 1
				s[1].SecondStar = 1
				s[2].ThirdStar = 1
				return s
			}(),
			expectedOutput: func() []StarCount {
				s := goldenStarCount()
				s[0].FirstStar = 1
				s[0].ThirdStar = 1
				s[1].SecondStar = 1
				s[1].FirstStar = 1
				s[2].ThirdStar = 1
				s[2].SecondStar = 1
				return s
			}(),
		},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			output := addStars(tc.originalStarCount, &tc.decisions)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}

func goldenDecisions() Decisions {
	return Decisions{
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
	}
}

func goldenStarCount() []StarCount {
	return []StarCount{
		StarCount{
			ID:         8470860,
			FullName:   "Jaroslav Halak",
			Link:       "/api/v1/people/8470860",
			FirstStar:  0,
			SecondStar: 0,
			ThirdStar:  0,
		},
		StarCount{
			ID:         8473419,
			FullName:   "Brad Marchand",
			Link:       "/api/v1/people/8473419",
			FirstStar:  0,
			SecondStar: 0,
			ThirdStar:  0,
		},
		StarCount{
			ID:         8477956,
			FullName:   "David Pastrnak",
			Link:       "/api/v1/people/8477956",
			FirstStar:  0,
			SecondStar: 0,
			ThirdStar:  0,
		},
	}
}
