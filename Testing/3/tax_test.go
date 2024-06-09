package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax := CalculateTax(1000.00)
	assert.Equal(t, 10.0, tax)

	tax = CalculateTax(0)
	assert.Equal(t, 0.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	respoisoty := &TaxRepositoryMock{}
	respoisoty.On("SaveTax", 10.0).Return(nil)
	respoisoty.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, respoisoty)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, respoisoty)
	assert.Error(t, err, "error saving tax")

	respoisoty.AssertExpectations(t)
	respoisoty.AssertNumberOfCalls(t, "SaveTax", 2)

}
