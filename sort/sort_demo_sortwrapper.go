/*
*Desc:
*CreateBy:Cooyw
*Time:2018/9/20
*/
package main

import (
	"fmt"
	"sort"
)

type Grams int

func (g Grams) String() string {
	return fmt.Sprintf("%dg", int(g))
}

type Organ struct {
	Name   string
	Weight Grams
}
type Organs []*Organ

func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//以name 方式排序
type ByOrganName struct{ Organs }

func (s ByOrganName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

//以Weight 方式排序
type ByOrganWeight struct{ Organs }

func (s ByOrganWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	s := []*Organ{
		{"brain", 1340},
		{"heart", 290},
		{"liver", 1494},
		{"pancreas", 131},
		{"prostate", 62},
		{"spleen", 162},
	}

	sort.Sort(ByOrganWeight{s})
	printOrgans(s)

	fmt.Println("-------------")
	sort.Sort(ByOrganName{s})
	printOrgans(s)

}

func printOrgans(s []*Organ) {
	for _, o := range s {
		fmt.Printf("%-10s(%v)\n", o.Name, o.Weight)
	}
}
