package category

// EnsureExists creates a category with the given name if it
// does not already exist
func EnsureExists(name string) error {
	if _, ok := categories[name]; ok {
		return nil
	}
	categories[name] = &Category{
		name: name,
	}
	return nil
}
