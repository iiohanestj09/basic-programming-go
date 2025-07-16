package main

import "fmt"

func ubahNilai(x int) {
    x = x + 10
    fmt.Println("Nilai di dalam fungsi:", x)
}

func main() {
    nilai := 5
    ubahNilai(nilai)
    fmt.Println("Nilai setelah fungsi dipanggil:", nilai)
}
