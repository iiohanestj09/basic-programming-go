package main

import "fmt"

func main() {
    var x int = 10
    var p *int

    p = &x

    fmt.Println("Alamat memori x:", p)
    fmt.Println("Nilai x melalui pointer:", *p)

    *p = 20
    fmt.Println("Nilai x setelah diubah melalui pointer:", x)
}
