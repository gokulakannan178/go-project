// Golang Program to Add Two Matrix Using Multi-dimensional Arrays
package main

import (
	"fmt"
)

type Data struct {
	Vid  string
	Name string
}
type Inentory struct {
	ID   int
	Data []Data
}

shared.Shared
func main() {
	sh = make(shared.Shared)
	inventory := []Inentory{}
	var size []Inentory
	size = append(size, []Inentory{{1, []Data{{"1", "20"}}}, {2, []Data{{"1", "20"}}}}...)
	var Color []Inentory
	Color = append(Color, []Inentory{{3, []Data{{"1", "B"}}}, {4, []Data{{"1", "G"}}}}...)
	// var Mat []Inentory
	// Mat = append(Mat, []Inentory{{4, []Data{{"1", "#"}}}, {6, []Data{{"1", "*"}}}}...)
	Mul(size, Color)
	fmt.Println(Mul(size, Color))

}

func Mul(m1 []Inentory, m2 []Inentory) []Inentory {
	inventoryCount := 0
	retData := []Inentory{}
	for _, v := range m1 {
		var i Inentory
		inventoryCount++
		for _, v := range m2 {
			i.ID = inventoryCount
			i.Data = append(i.Data, v.Data)
		}
	}
	return retData
}
