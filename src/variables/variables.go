package variables

import "path"

// Return path
var version string

var (
	//related
	video string

	//closely related
	duration int
	current  int
)

func GetFilePath(filepath string) string {
	var _, file string

	_, file = path.Split(filepath)

	return file
}

func CalculateArea() int {
	width, height := 50, 10

	return width * height
}

func TypeConversion() int {
	speed, force := 100, 2.5
	speed = int(float64(speed) * force)
	return speed
}

func RawString() string {
	a := ``
	return a
}

