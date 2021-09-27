package checkout

import "viniti.us/hashout/models/checkout"

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return UseCase{repo: repo}
}

func (u UseCase) Checkout(c *checkout.Cart) error {
	products, err := u.repo.Find(c.Items)
	if err != nil {
		return err
	}

	// Fetch discount for each product

	// Check for black friday

	c.Items = products
	c.CalculateTotals()

	return nil
}
