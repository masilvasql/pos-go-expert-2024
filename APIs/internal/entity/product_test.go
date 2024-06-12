package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Notebook", 2000)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotNil(t, product.CreatedAt)
	assert.Equal(t, "Notebook", product.Name)
	assert.Equal(t, 2000, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 2000)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "name is required")
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	_, err := NewProduct("Produto", 0)
	assert.EqualError(t, err, "price is required")
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	_, err := NewProduct("Produto", -10)
	assert.EqualError(t, err, "invalid Price")
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Produto", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
