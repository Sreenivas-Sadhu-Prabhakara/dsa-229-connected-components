package main

import "testing"

func TestCountComponents(t *testing.T) {
	cases := []struct {
		n     int
		edges [][2]int
		want  int
	}{
		{5, [][2]int{{0, 1}, {1, 2}, {3, 4}}, 2},
		{5, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 1},
		{4, nil, 4}, // edgeless: every node its own component
		{1, nil, 1},
		{0, nil, 0},
		{6, [][2]int{{0, 1}, {2, 3}, {4, 5}}, 3},
		{3, [][2]int{{0, 1}, {1, 0}}, 2}, // parallel edges don't merge further
	}
	for _, c := range cases {
		if got := CountComponents(c.n, c.edges); got != c.want {
			t.Errorf("CountComponents(%d,%v) = %d, want %d", c.n, c.edges, got, c.want)
		}
	}
}
