/**
 * User: RockyF
 * Date: 13-12-5
 * Time: 上午11:49
 */
package tbs

import "math"

type Point struct {
	X int32
	Y int32
}

type Size struct {
	Width		int32
	Height	int32
}

func CreatePoint(x int32, y int32) *Point{
	return &Point{X:x, Y:y};
}

func (this *Point)Add(point *Point){
	this.X = this.X + point.X
	this.Y = this.Y + point.Y
}

func (this *Point)Minus(point *Point){
	this.X = this.X - point.X
	this.Y = this.Y - point.Y
}

func (this * Point)Distance(point *Point) float64{
	return Distance(this, point)
}

func Distance(p1 * Point, p2 *Point) float64{
	return math.Sqrt(math.Pow((float64)(p2.Y-p1.Y), 2) + math.Pow((float64)(p2.X-p1.X), 2))
}

type Line struct {
	Pos1 Point
	Pos2 Point
}

func (this *Line)InLine(point *Point, width *int)bool{
	return false
}

type Rect struct {
	Pos				Point
	Size			Size
	Rotation	int32
}

func (this *Rect)GetLeft()int32{
	return this.Pos.X;
}

func (this *Rect)GetTop()int32{
	return this.Pos.Y;
}

func (this *Rect)GetWidth()int32{
	return this.Size.Width;
}

func (this *Rect)GetHeight()int32{
	return this.Size.Height;
}

func (this *Rect)Contains(point *Point)bool{
	if(this.Rotation == 0){
		return point.X > this.Pos.X && point.Y > this.Pos.Y && point.X < this.Pos.X + this.Size.Width && point.Y < this.Pos.Y + this.Size.Height
	}else{

	}

	return false;
}

type Polygon struct{
	PosArr  []*Point
}

func CreatePolygon(source []*Point) *Polygon{
	var polygon *Polygon
	polygon.PosArr = source

	return polygon
}

func (this *Polygon)Contains(point *Point) bool{
	count1 := 0
	count2 := 0
	len := len(this.PosArr)

	j := 0
	for i := 0;i < len;i = i + i {
		j = i - 1
		z := (point.Y - this.PosArr[j].X) * (this.PosArr[i].Y - this.PosArr[j].Y) - (point.Y - this.PosArr[j].Y) * (this.PosArr[i].X - this.PosArr[j].X);
		if z > 0{
			count1 = count1 + 1
		}else if z < 0{
			count2 = count2 + 1
		}
	}

	return (count1 == 0 || count2 == 0)
}

type Circle struct {
	Pos			Point
	Radius	int32
}

type Vector struct {
	Pos	Point
	Vx	float32
	Vy	float32
}
