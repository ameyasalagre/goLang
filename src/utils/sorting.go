package main
import (
	"fmt";
	"sort"
)
func main() {
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}

//run 
// set GOPATH=E:\Coding\Golang\HelloWorld\HelloWorld
// set GOBIN=E:\Coding\Golang\HelloWorld\HelloWorld\bin
// go run src\main\hello.go  =-----> Check dir path and set accordingly