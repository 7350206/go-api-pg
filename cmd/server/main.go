// it'll be run first
// here will be instantiation for all components

package main

import "fmt"

// responsible for instantiation and app startup
func Run() error {
	fmt.Println("starting app")
	return nil
}

func main() {

	// handle errs while app starts, instead main() panic - catch err and [print]
	fmt.Println("main started")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
