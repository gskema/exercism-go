package queenattack

import (
	"errors"
)

const testVersion = 1

type Point struct {
	x, y int
}

var cols = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
}

var rows = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
}

func createPointChessPosition(pos string) (Point, bool) {
	if 2 != len(pos) {
		return Point{0, 0}, false
	}

	col, ok1 := cols[pos[:1]]
	row, ok2 := rows[pos[1:2]]

	if !ok1 || !ok2 {
		return Point{0, 0}, false
	}

	return Point{col, row}, true
}

func CanQueenAttack(pos1, pos2 string) (bool, error) {
	if pos1 == pos2 {
		return false, errors.New("Invalid positions")
	}

	q1, ok1 := createPointChessPosition(pos1)
	q2, ok2 := createPointChessPosition(pos2)

	if !ok1 || !ok2 {
		return false, errors.New("Invalid chess position")
	}

	return q1.x == q2.x || q1.y == q2.y || q1.x-q2.x == q1.y-q2.y || q1.x-q2.x == -(q1.y-q2.y), nil
}
