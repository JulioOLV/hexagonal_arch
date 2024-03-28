package application_test

import (
	"testing"

	"github.com/JulioOLV/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than 0 to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Product Test"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than zero", err.Error())

	product.Id = "123"
	_, err = product.IsValid()
	require.Error(t, err)

	product.Id = uuid.NewV4().String()
	product.Name = ""
	require.Error(t, err)

	product.Name = "Product Test"
	product.Status = ""
	require.Error(t, err)
}