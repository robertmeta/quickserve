package main

import "fmt"

// Define a type named "servabledir" as a slice of strings
type servableDir []string

// Now, for our new type, implement the two methods of
// the flag.Value interface...
// The first method is String() string
func (s *servableDir) String() string {
	return fmt.Sprintf("%s", *s)
}

// The second method is Set(value string) error
func (s *servableDir) Set(value string) error {
	*s = append(*s, value)
	return nil
}
