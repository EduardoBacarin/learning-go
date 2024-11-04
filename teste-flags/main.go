package main

import (
	"flag"
	"fmt"
)

var sobrenome string

func main() {
	var nomeFlag = flag.String("nome", "Nome", "Insira o nome")
	var idadeFlag = flag.Int("idade", 18, "Insira a idade")
	var valorFlag = flag.Float64("valor", 1, "Insira o valor")
	flag.StringVar(&sobrenome, "sobrenome", "Sobrenome", "Insira o sobrenome")

	flag.Parse()
	fmt.Println("O nome é ", *nomeFlag)
	fmt.Println("O sobrenome é ", *&sobrenome)
	fmt.Println("Idade inserida é ", *idadeFlag)
	fmt.Println("O valor é de R$", *valorFlag)
}
