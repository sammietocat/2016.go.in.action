package search

// defaultMatcher implements the default matcher.
type defaultMatcher struct {
}

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// init registers the default matcher with the program.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
