package main

import "fmt"

func vote(totSiswa, totA, totB int, hasil *string) {
    totPemilih := totA + totB

    if totPemilih >= int(0.6*float64(totSiswa)) { 
        if totA > totB {
            *hasil = "Kandidat A menang"
        } else if totB > totA {
            *hasil = "Kandidat B menang"
        } else {
            *hasil = "tidak ada pemenang"
        }
    } else {
        *hasil = "tidak ada pemenang"
    }
}

func main() {
    var totSiswa, totA, totB int
    var hasil string

    fmt.Scan(&totSiswa, &totA, &totB)
    vote(totSiswa, totA, totB, &hasil)
    fmt.Println(hasil)
}
