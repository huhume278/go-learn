package main

import "fmt"

type Config struct {
	path string
}

func (c *Config) Path() string {
	if c == nil {
		return "usr/home"
	}
	return c.path
}

func main() {
	var c1 *Config
	var c2 = &Config{
		path: "/export",
	}
	fmt.Printf(c1.Path(), c2.Path())
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
