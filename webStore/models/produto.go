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

	selectAllProducts, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")

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

		p.Id = id
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

func DeleteProduct(id string) {
	db := db.ConnectToDatabase()
	deleteProduct, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func GetProduct(id string) Produto {
	db := db.ConnectToDatabase()

	produtoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	newProduct := Produto{}
	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		newProduct.Id = id
		newProduct.Nome = nome
		newProduct.Descricao = descricao
		newProduct.Preco = preco
		newProduct.Quantidade = quantidade
	}

	defer db.Close()
	return newProduct
}

func UpdateProduct(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectToDatabase()

	atualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
