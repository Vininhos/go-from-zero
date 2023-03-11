package main

import "fmt"

func main() {
    var numero int64 = -10000000
    fmt.Println(numero)

    var numero2 uint32 = 10000
    fmt.Println(numero2)

    var numero4 byte = 123
    fmt.Println(numero4)

    var numeroReal1 float32 = 123.45
    fmt.Println(numeroReal1)

    var numeroReal2 float64 = 12310923.45
    fmt.Println(numeroReal2)

    numeroReal3 := 12345.67
    fmt.Println(numeroReal3)

    var str string = "Texto"
    fmt.Println(str)

    str2 := "Texto2"
    fmt.Println(str2)

    char := 'B'
    fmt.Println(char)

    // FIM STRINGS
    var texto string
    fmt.Println(texto)

    var booleano1 bool = true
    fmt.Println(booleano1)

    var erro error = errors.New("Erro interno")
    fmt.Println(erro)
}