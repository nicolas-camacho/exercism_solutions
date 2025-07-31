//Package accumulate is an implementation for exercism's accumulate exercise.
package accumulate

type function func(string) string

//Accumulate iterates through a slice of string and applies a function recived to each item.
func Accumulate(list []string, modifier function) (modified []string) {
	for _, item := range list {
		modified = append(modified, modifier(item))
	}
	return
}
