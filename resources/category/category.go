package category

// Category defines the structure of a categor record
type Category struct {
	name        string
	numListings int
}

// categories is a map of category name to the category record
var categories map[string]*Category

func init() {
	categories = make(map[string]*Category)
	categoryHeap = &categoryHeapType{}
}

// GetCategoryIndex returns the index of categories by name
func GetCategoryIndex() map[string]*Category {
	return categories
}
