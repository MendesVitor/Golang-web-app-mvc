package models

import (
	"store/config/db"
)

type Product struct {
	Name, Description string
	Price             float64
	Id, Quantity      int
}

func GetAllProducts() []Product {
	db := db.ConnDB()

	getAllProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for getAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getAllProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func NewProduct(name, description string, price float64, quantity int) {
	db := db.ConnDB()

	insertData, err := db.Prepare("insert into products (name,description,price,quantity) values($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnDB()

	delete, err := db.Prepare("delete from products where id =$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnDB()

	product, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for product.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = product.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}
	defer db.Close()
	return productToUpdate
}

func UpdatedProduct(name, description string, price float64, id, quantity int) {
	db := db.ConnDB()

	updateData, err := db.Prepare("update products set name=$1,description=$2,price=$3,quantity=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateData.Exec(name, description, price, quantity, id)
	defer db.Close()
}
