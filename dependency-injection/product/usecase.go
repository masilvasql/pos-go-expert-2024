package product

type ProductUseCase struct {
	repository ProductRepositoriyInterface
}

func NewProductUseCase(r ProductRepositoriyInterface) *ProductUseCase {
	return &ProductUseCase{r}
}

func (uc *ProductUseCase) GetProduct(id int) (Product, error) {

	return uc.repository.GetProduct(id)
}
