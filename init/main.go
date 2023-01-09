package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	operations "ql/init/math"
)

type Carro struct {
	Nome string
}

type Gato struct {
	Nome string
}

type Cliente struct {
	Nome string
	Email string
	CPF int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo metodo: ", c.Nome)
}

type ClienteInternacional struct {
	Cliente
	País string `json:"país"`
}

func (c Carro) andar() {
	c.Nome = "Ka"
	fmt.Println(c.Nome, "andou")
}

func (g *Gato) miaa() {
	g.Nome = "Jujiu"
	fmt.Println(g.Nome, "Miaaauuu")
}

func main() {
	carro := Carro {
		Nome: "BMW",
	}
	carro.andar()
	fmt.Println(carro.Nome)


	gato := Gato {
		Nome: "Bilu",
	}
	gato.miaa()
	fmt.Println(gato.Nome)


	cliente := Cliente {
		Nome: "Níkollas",
		Email: "nikollas@gmail.com",
		CPF: 12345678900,
	}
	fmt.Println(cliente)
	cliente.Exibe()


	clienteInt := ClienteInternacional {
		Cliente: Cliente {
			Nome: "Garra",
			Email: "garra@gmail.com",
			CPF: 98765432100,
		},
		País: "EUA",
	}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d. País: %s\n", clienteInt.Nome, clienteInt.Email, clienteInt.CPF, clienteInt.País)
	clienteInt.Exibe()


	clienteJson, err := json.Marshal(clienteInt)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(clienteJson))


	fmt.Println("Hello World")
	a := "Nikollas"
	c := 3.234
	b := 5
	b += 2
	d := false
	e := `urgrgrgrdg`
	fmt.Println(a, b, c, d, c, e)
	fmt.Printf("%T, %T, %T, %T, %T, %T \n", a, b, c, d, c, e)


	resultSome := operations.Soma(1, 4)
	fmt.Println(resultSome)


	resultDiv, errDiv := operations.Division(1, 4)
	if errDiv != nil {
		log.Fatal(errDiv.Error())
	}
	fmt.Println(resultDiv)


	res, errHttp := http.Get("http://google.com.br")
	if errHttp != nil {
		log.Fatal("Erro ao consultar ", errHttp.Error())
	}
	fmt.Println(res.Header)


	somaTudo := operations.SomaTudo(1, 6, 7, 8, 4, 2)
	fmt.Println(somaTudo)


	i := 10
	fmt.Println(&i)
	var ponteiro *int = &i
	fmt.Println(*ponteiro)
	*ponteiro = 50
	fmt.Println(i)
	fmt.Println(*ponteiro)


	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)

	jsonCliente4 := `{"Nome":"Juji","Email":"juji@gmail.com","CPF":55555500,"país":"HOL"}`
	cliente4 := ClienteInternacional{}
	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println(cliente4)
}

func abc(a *int) {
	*a = 200
}