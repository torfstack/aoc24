package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	testfile := "testinput.txt"
	part1(testfile)
	part2(testfile)

	inputfile := "input.txt"
	part1(inputfile)
	part2(inputfile)
}

func part2(file string) {
	m, p := parseInput(file)

	count := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m.HasObstacleAt(Position{X: i, Y: j}) {
				continue
			}

			m.AddObstacleAt(Position{X: i, Y: j})
			if m.PositionWillEndUpGoingInCircles(p) {
				count++
			}
			m.RemoveObstacleFrom(Position{X: i, Y: j})
			m.RefreshTiles()
		}
	}

	log.Printf("Part 2: %d", count)
}

func part1(file string) {
	m, p := parseInput(file)

	for m.Contains(p) {
		if m.GetTile(p).IsObstacle() {
			p.TakeStepBack()
			p.ChangeDirection()
		}
		m.GetTile(p).MarkForDirection(p.Dir)
		p.TakeStep()
	}

	log.Printf("Part 1: %d", m.CountSteppedOn())
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Position struct {
	X   int
	Y   int
	Dir Direction
}

func (p *Position) ChangeDirection() {
	p.Dir = (p.Dir + 1) % 4
}

func (p *Position) TakeStep() {
	switch p.Dir {
	case North:
		p.X--
	case East:
		p.Y++
	case South:
		p.X++
	case West:
		p.Y--
	}
}

func (p *Position) TakeStepBack() {
	switch p.Dir {
	case North:
		p.X++
	case East:
		p.Y--
	case South:
		p.X--
	case West:
		p.Y++
	}
}

type Tile struct {
	Content   rune
	SteppedOn map[Direction]bool
}

func NewTile(c rune) Tile {
	return Tile{Content: c, SteppedOn: make(map[Direction]bool)}
}

func (t *Tile) IsObstacle() bool {
	return t.Content == '#'
}

func (t *Tile) MarkForDirection(dir Direction) {
	t.SteppedOn[dir] = true
}

func (t *Tile) HasBeenSteppedOn(dir Direction) bool {
	return t.SteppedOn[dir]
}

type Map [][]Tile

func (m *Map) GetTile(p Position) *Tile {
	return &(*m)[p.X][p.Y]
}

func (m *Map) Contains(p Position) bool {
	return p.X >= 0 && p.X < len(*m) && p.Y >= 0 && p.Y < len((*m)[0])
}

func (m *Map) HasObstacleAt(p Position) bool {
	return m.GetTile(p).IsObstacle()
}

func (m *Map) CountSteppedOn() int {
	count := 0
	for _, row := range *m {
		for _, t := range row {
			if len(t.SteppedOn) > 0 {
				count++
			}
		}
	}
	return count
}

func (m *Map) PositionWillEndUpGoingInCircles(p Position) bool {
	for m.Contains(p) {
		if m.GetTile(p).IsObstacle() {
			p.TakeStepBack()
			p.ChangeDirection()
		}

		if m.GetTile(p).HasBeenSteppedOn(p.Dir) {
			return true
		}

		m.GetTile(p).MarkForDirection(p.Dir)
		p.TakeStep()
	}

	return false
}

func (m *Map) AddObstacleAt(pos Position) {
	(*m)[pos.X][pos.Y].Content = '#'
}

func (m *Map) RemoveObstacleFrom(pos Position) {
	(*m)[pos.X][pos.Y].Content = '.'
}

func (m *Map) RefreshTiles() {
	for i := 0; i < len(*m); i++ {
		for j := 0; j < len((*m)[i]); j++ {
			(*m)[i][j].SteppedOn = make(map[Direction]bool)
		}
	}
}

func parseInput(file string) (Map, Position) {
	var m Map
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	r := 0
	var pos Position
	for scanner.Scan() {
		line := scanner.Text()
		var row []Tile
		for i, c := range line {
			if c == '^' {
				pos = Position{X: r, Y: i, Dir: North}
				row = append(row, NewTile('.'))
				continue
			}
			row = append(row, NewTile(c))
		}
		m = append(m, row)
		r++
	}

	return m, pos
}
