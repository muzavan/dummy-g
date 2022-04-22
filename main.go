package main

import "fmt"

func main() {
	fmt.Println(getTheBestMan("vim-go"))
}

func getTheBestMan(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
