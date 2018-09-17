/**
 * Description: 
 * User: 1067
 * Date: 2018-09-17
 * Time: 10:19
 */

package main

import (
	"sort"
	"fmt"
)

// A couple of type definitions to make the units clear.#两个类型定义，使单元更加清晰。
type earthMass float64

type au float64

// A Planet defines the properties of a solar system object.#行星定义了太阳系天体的属性。
type Planet struct {
	name     string
	mass     earthMass
	distance au
}

//By是定义其Planet参数顺序的“less”函数的类型。
type By func(p1, p2 *Planet) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
// Sort是函数类型上的一个方法，By，根据函数对参数片进行排序。
func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by,
	}
	sort.Sort(ps)
}

type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool
}

func (s *planetSorter) Len() int {
	return len(s.planets)
}

// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

var planets = []Planet{
	{"Mercury", 0.055, 0.4},
	{"Venus", 0.815, 0.7},
	{"Earth", 1.0, 1.0},
	{"Mars", 0.107, 1.5},
}

func main() {
	sortByName := func(p1, p2 *Planet) bool { return p1.name < p2.name }
	sortByMass := func(p1, p2 *Planet) bool { return p1.mass < p2.mass }
	sortByDistance := func(p1, p2 *Planet) bool { return p1.distance < p2.distance }

	//ps := planetSorter{
	//	planets: planets,
	//	by:      sortByName,
	//}
	By(sortByName).Sort(planets)
	fmt.Println(planets)

	By(sortByMass).Sort(planets)
	fmt.Println(planets)


	By(sortByDistance).Sort(planets)
	fmt.Println(planets)
}
