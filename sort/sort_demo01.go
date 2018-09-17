/*
*Desc:
*CreateBy:Cooyw
*Time:2018/9/16
*/
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}


// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	people := []Person{
		{"Bob", 20},
		{"Eric", 19},
		{"Dale", 16},
		{"Curry", 18},
		{"Alice", 17},
	}
	fmt.Println(people)
	//var sortPeople ByAge
	//sortPeople = people
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)

}
