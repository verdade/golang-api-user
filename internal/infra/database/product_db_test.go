package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/verdade/golang-api-user/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Novo produto", 10)
	assert.NoError(t, err)
	productDb := NewProduct(db)
	err = productDb.Create(product)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
}

func TestProduct_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*10)
		assert.NoError(t, err)
		productDb := NewProduct(db)
		err = productDb.Create(product)
	}
	productDb := NewProduct(db)
	products, err := productDb.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products2, err := productDb.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products2, 10)
	assert.Equal(t, "Product 11", products2[0].Name)
	assert.Equal(t, "Product 20", products2[9].Name)

	products3, err := productDb.FindAll(2, 3, "asc")
	assert.NoError(t, err)
	assert.Len(t, products3, 3)
	assert.Equal(t, "Product 4", products3[0].Name)
	assert.Equal(t, "Product 6", products3[2].Name)
}

func TestProduct_FindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDb := NewProduct(db)
	product, err = productDb.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Product 1", product.Name)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.NoError(t, err)
	product, err = productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
}
