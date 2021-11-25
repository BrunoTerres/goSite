package models

import (
	"std/bruno/goSite/db"
)

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SelectTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}
