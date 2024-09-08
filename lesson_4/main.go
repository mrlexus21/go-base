package main

import (
	"fmt"
)
import _ "github.com/mitchellh/mapstructure"

type Point struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func (p Point) method() {
	fmt.Println("call Point method")
}

func main() {
	/*arr := []string{"a", "b", "c"}
	for _, l := range arr {
		fmt.Println(l)
	}*/
	/*pointsMap := map[string]Point{
		"b": {
			X: 13,
			Y: 15,
		},
	}
	otherPointsMap := make(map[int]Point)
	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(pointsMap)
	fmt.Println(pointsMap["a"])
	otherPointsMap[1] = Point{1, 2}
	fmt.Println(otherPointsMap)
	fmt.Println(otherPointsMap[1])*/

	/*var oneMorePointsMap map[string]Point
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point{"a": {1, 2}}
	}
	fmt.Println(oneMorePointsMap)

	key := "a"
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("key=%s exist in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("key=%s not exist in map\n", key)
		fmt.Println(value)
	}*/

	pointsMap := map[string]int{
		"xx": 1,
		"yy": 2,
	}
	/*p1 := Point{}
	mapstructure.Decode(pointsMap, &p1)
	fmt.Println(p1)*/
	for k, v := range pointsMap {
		fmt.Println(k)
		fmt.Println(v)
	}
}
