package main

import "github.com/lggomezml/dummyrepo-go/demo"
import "github.com/lggomezml/dummyrepo-go-dep/demodep"
import "fmt"

func main() {
	fmt.Println("local dep (dummyrepo-go): " + demo.GetGreeting())
	fmt.Println("inner dep (dummyrepo-go via dummyrepo-go-dep): " + demodep.GetDependencyGreeting())
}
