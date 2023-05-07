package main

import "math"

//Shape はインターフェイスです。
type Shape interface {
	Area() float64
}

// Rectangle は四角形を表す
type Rectangle struct {
	Width  float64
	Height float64
}

//Area は長方形の面積を計算
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter は長方形の周囲を計算
// 引数をRectangle構造体にすることで、誤って使われるリスクを無くす
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle は円形を表す
type Circle struct {
	Radius float64
}

//Area は円形の面積を計算
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle は三角形を表す
type Triangle struct {
	Base   float64
	Height float64
}

//Area は三角形の面積を計算
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
