package main

import (
	"fmt"

	envfile "github.com/oksuriini/envarParser"
)

func main() {
	env, err := envfile.GetOneEnvar("EXAMPLE_KEY", "envarfile")
	if err != nil {
		fmt.Printf("error %v encountered", err)
		return
	}
	fmt.Println(env) // --> Prints "EXAMPLE_VALUE"
}
