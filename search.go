package goinsta

type Search struct {
	inst *Instagram
}

// NewSearch creates new Search structure
func NewSearch(inst *Instagram) *Search {
	search := &Search{
		inst: inst,
	}
	return search
}

// ByUser performs a search by username
func (search *Search) User(user string) error {
	// TODO
}
