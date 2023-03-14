package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Vinícius"
	u.idade = 23
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua X", 10}

	usuario2 := usuario{"Vinícius", 23, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Vinícius"}
	fmt.Println(usuario3)
}
