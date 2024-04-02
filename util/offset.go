package util

import (
	"math"
)

func CalculateOffset(points *[]Vec2) Vec2 {
	return Vec2{
		X: (*points)[1].X - (*points)[0].X,
		Y: (*points)[1].Y - (*points)[0].Y,
	}
}

func GetOffsets(points *[]Vec2) (map[int]Vec2, Vec2) {
	c1_c2_points := (*points)[0:2]
	c2_c3_points := (*points)[1:3]
	offsets := map[int]Vec2{}
	offsets[1] = Vec2{}
	offset_c1_c2 := CalculateOffset(&c1_c2_points)
	offsets[0] = offset_c1_c2
	for i, j := 0, len(c2_c3_points)-1; i < j; i, j = i+1, j-1 {
		c2_c3_points[i], c2_c3_points[j] = c2_c3_points[j], c2_c3_points[i]
	}
	offset_c2_c3 := CalculateOffset(&c2_c3_points)
	offsets[2] = offset_c2_c3
	maxOffset := Vec2{}
	for _, offset := range offsets {
		if math.Abs(float64(offset.X)) > math.Abs(float64(maxOffset.X)) {
			maxOffset.X = offset.X
		}
		if math.Abs(float64(offset.Y)) > math.Abs(float64(maxOffset.Y)) {
			maxOffset.Y = offset.Y
		}
	}
	return offsets, maxOffset
}
