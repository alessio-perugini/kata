package main

import "testing"

func TestUniverse_LivingNeighbors(t *testing.T) {
	tests := []struct {
		name     string
		universe []uint8
		x        uint8
		y        uint8
		expected uint8
	}{
		{
			name:     "test top bound",
			universe: []uint8{0, setBit(0, 7), 0, 0, 0, 0, 0, 0},
			x:        0,
			y:        7,
			expected: 0,
		},
		{
			name:     "test bottom bound",
			universe: []uint8{0, 0, 0, 0, 0, 0, setBit(0, 7), 0},
			x:        7,
			y:        7,
			expected: 0,
		},
		{
			name:     "test left bound",
			universe: []uint8{0, setBit(0, 7), 0, 0, 0, 0, 0, 0},
			x:        1,
			y:        8,
			expected: 0,
		},
		{
			name:     "test left bound",
			universe: []uint8{0, setBit(0, 7), 0, 0, 0, 0, 0, 0},
			x:        1,
			y:        1,
			expected: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			u := &Universe{
				universe: tt.universe,
			}

			if got := u.LivingNeighbors(tt.x, tt.y); got != tt.expected {
				t.Errorf("Errorre uaglio!, got=%v expected=%v", got, tt.expected)
			}
		})
	}
}
