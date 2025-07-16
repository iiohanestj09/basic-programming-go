package main

import "fmt"

type Mahasiswa struct {
    Nama string
    Umur int
}

func (m *Mahasiswa) TambahUmur() {
    m.Umur += 1
}

func main() {
    mhs := Mahasiswa{Nama: "Safira", Umur: 20}
    fmt.Println("Sebelum:", mhs.Umur)
    mhs.TambahUmur()
    fmt.Println("Sesudah:", mhs.Umur)
}
