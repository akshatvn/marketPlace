package category

import (
	"container/heap"
	"errors"
	"fmt"
	"github.com/akshatvn/marketPlace/resources/user"
)

/**** START --- HEAP IMPLEMENTATION FOR KEEPING TRACK OF TOP CATEGORY ****/
type categoryHeapType []*Category

var categoryHeap *categoryHeapType

func (h categoryHeapType) Len() int {
	return len(h)
}

func (h categoryHeapType) Less(i, j int) bool {
	return h[i].numListings > h[j].numListings
}

func (h categoryHeapType) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *categoryHeapType) Push(x interface{}) {
	*h = append(*h, x.(*Category))
}

func (h *categoryHeapType) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

/**** END --- HEAP IMPLEMENTATION FOR KEEPING TRACK OF TOP CATEGORY ****/

// UpdateHeap alters number of listings of a given category,
// and then ensures heap invariant in O(log(n)) time
func UpdateHeap(name string, inc bool) {
	cat := categories[name]
	if inc {
		cat.numListings += 1
	} else {
		cat.numListings -= 1
	}
	heap.Fix(categoryHeap, 0)
}

// GetTopCategory returns the name of the category with the highest
// number of listings in O(1) time
func GetTopCategory(args ...string) error {
	if len(args) < 1 {
		return errors.New("Category Get needs at least 1 argument")
	}
	username := args[0]
	if !user.DoesExist(username) {
		return errors.New("Error - unknown user")
	}
	var maxCategory *Category
	for _, c := range categories {
		if maxCategory == nil  || c.numListings > maxCategory.numListings {
			maxCategory = c
		}
	}
	fmt.Println(maxCategory.name)
	return nil
}

