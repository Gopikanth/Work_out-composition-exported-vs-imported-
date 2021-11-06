package main

import (
	"fmt"
	"myapp/pack"
)

var planets = []pack.Planet{
	{
		Name:                   "jupiter",
		Size:                   100000,
		Distance_between_earth: 5.6},
	{
		Name:                   "mars",
		Size:                   1000,
		Distance_between_earth: 7.4},
	{
		Name:                   "venus",
		Size:                   1000,
		Distance_between_earth: 10.6},
}

func main() {
	system := pack.Solar_system{
		All_planents: planets,
	}
	fmt.Println(system.All())
	fmt.Println(system.Farest())
}
