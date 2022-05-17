package models

import "GoLangAlura/webStore/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProducts() []Produto {
	db := db.ConnectToDatabase()

	selectAllProducts, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
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

func CreateNewProduct(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectToDatabase()

	insertOnDatabase, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertOnDatabase.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
