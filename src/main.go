package main

import (
	"DotEnvar/src/envfile"
	"fmt"
	"os"
)

func main() {
	envfile.ParseEnvFile(".dotenvar")

	fmt.Println(os.Getenv("API_KEY"))
}
