package product

import "context"

type Repository interface {
	Create(ctx context.Context, p *Product) error
	GetByID(ctx context.Context, id int64) (*Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id int64) error
}