/**
 * Description: 
 * User: 1067
 * Date: 2018-09-19
 * Time: 10:24
 */

package main

import (
	"sort"
	"fmt"
)

type Change struct {
	user     string
	language string
	lines    int
}

type lessFunc func(p, p1 *Change) bool

type multiSort struct {
	changes []Change
	less    []lessFunc
}

func (ms *multiSort) Sort(changes []Change) {
	ms.changes = changes
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSort {
	return &multiSort{
		less: less,
	}
}

func (ms *multiSort) Len() int {
	return len(ms.changes)
}

func (ms *multiSort) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

func (ms *multiSort) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

func main() {
	var changes = []Change{
		{"gri", "Go", 100},
		{"ken", "C", 150},
		{"glenda", "Go", 200},
		{"rsc", "Go", 200},
		{"r", "Go", 100},
		{"ken", "Go", 200},
		{"dmr", "C", 100},
		{"r", "C", 150},
		{"gri", "Smalltalk", 80},
	}

	user := func(c1, c2 *Change) bool {
		return c1.user < c2.user
	}

	language := func(c1, c2 *Change) bool {
		return c1.language < c2.language
	}

	decreasingLines := func(c1, c2 *Change) bool {
		return c1.lines < c2.lines
	}

	OrderedBy(user).Sort(changes)
	fmt.Println(changes)

	OrderedBy(language).Sort(changes)
	fmt.Println(changes)

	OrderedBy(decreasingLines).Sort(changes)
	fmt.Println(changes)

	//现根据user排序  在根据language排序
	OrderedBy(user, language).Sort(changes)
	fmt.Println(changes)

	//现根据user排序  在根据language排序 在根据lines排序
	OrderedBy(user, language, decreasingLines).Sort(changes)
	fmt.Println(changes)
}
