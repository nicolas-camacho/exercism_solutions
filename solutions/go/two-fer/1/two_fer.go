package twofer

// ShareWith is a function to use the name param to share one with me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
