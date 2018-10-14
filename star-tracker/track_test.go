package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStars(t *testing.T) {
	td := []struct {
		name           string
		decisions      Decisions
		Stars          Stars
		expectedOutput Stars
	}{
		{
			name:      "Add stars to empty struct",
			decisions: goldenDecisions(),
			Stars:     Stars{},
			expectedOutput: func() Stars {
				s := goldenStars()
				s.Stars[0].FirstStar = 1
				s.Stars[1].SecondStar = 1
				s.Stars[2].ThirdStar = 1
				return s
			}(),
		},
		{
			name:      "Increases star count from 0s",
			decisions: goldenDecisions(),
			Stars:     goldenStars(),
			expectedOutput: func() Stars {
				s := goldenStars()
				s.Stars[0].FirstStar = 1
				s.Stars[1].SecondStar = 1
				s.Stars[2].ThirdStar = 1
				return s
			}(),
		},
		{
			name:      "Increases star count from 1s",
			decisions: goldenDecisions(),
			Stars: func() Stars {
				s := goldenStars()
				s.Stars[0].FirstStar = 1
				s.Stars[1].SecondStar = 1
				s.Stars[2].ThirdStar = 1
				return s
			}(),
			expectedOutput: func() Stars {
				s := goldenStars()
				s.Stars[0].FirstStar = 2
				s.Stars[1].SecondStar = 2
				s.Stars[2].ThirdStar = 2
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
			Stars: func() Stars {
				s := goldenStars()
				s.Stars[0].FirstStar = 1
				s.Stars[1].SecondStar = 1
				s.Stars[2].ThirdStar = 1
				return s
			}(),
			expectedOutput: func() Stars {
				s := goldenStars()
				s.Stars[0].FirstStar = 1
				s.Stars[0].ThirdStar = 1
				s.Stars[1].SecondStar = 1
				s.Stars[1].FirstStar = 1
				s.Stars[2].ThirdStar = 1
				s.Stars[2].SecondStar = 1
				return s
			}(),
		},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			tc.Stars.update(tc.decisions)
			assert.Equal(t, tc.expectedOutput, tc.Stars)
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

func goldenStars() Stars {
	return Stars{
		Stars: []StarCount{
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
		},
	}
}
