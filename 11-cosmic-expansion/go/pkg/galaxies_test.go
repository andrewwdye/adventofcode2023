package pkg

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSolve1(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	assert.Equal(t, 374, lo.Must(Solve1(strings.NewReader(input))))
}

func TestParseUniverse(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	univ := ParseUniverse(strings.NewReader(input))
	assert.Len(t, univ.Galaxies, 9, univ)
	assert.Equal(t, 13, univ.Width)
	assert.Equal(t, 12, univ.Height)
}

func TestExpand(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	expected := `....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`
	assert.Equal(t, expected, strings.Join(expand(strings.Split(input, "\n")), "\n"))
}

func TestSolve2(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	assert.Equal(t, 1030, lo.Must(Solve2(strings.NewReader(input), 10)))
	assert.Equal(t, 8410, lo.Must(Solve2(strings.NewReader(input), 100)))
}

func TestParseUniverseBig(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	univ := ParseUniverseBig(strings.NewReader(input), 10)
	assert.Len(t, univ.Galaxies, 9, univ)
	assert.Equal(t, 10-3+3*10, univ.Width)
	assert.Equal(t, 10-2+2*10, univ.Height)
	assert.Equal(t, Location{12, 0}, univ.Galaxies[0])
}
