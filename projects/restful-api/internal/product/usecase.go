package product

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("product not found")
var ErrInvalidInput = errors.New("invalid input")

type Usecase interface {
	Create(ctx context.Context, p *Product) error
	GetByID(ctx context.Context, id int64) (*Product, error)
	GetAll(ctx context.Context) ([]*Product, error) 
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id int64) error
}

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{repo : repo}
}

func (u *usecase) Create(ctx context.Context, p *Product) error {
	if p.Name == "" || p.Price <= 0 {
		return ErrInvalidInput
	}
	return u.repo.Create(ctx, p)
}

func (u *usecase) GetByID(ctx context.Context, id int64) (*Product, error) {
	p, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, ErrNotFound
	}
	return p, nil
}

func (u *usecase) GetAll(ctx context.Context) ([]*Product, error) {
	return u.repo.GetAll(ctx)
}

func (u *usecase) Update(ctx context.Context, p *Product) error {
	if p.Name == "" || p.Price <= 0 {
		return ErrInvalidInput
	}
	existing, err := u.repo.GetByID(ctx, p.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrNotFound
	}
	return u.repo.Update(ctx, p)
}

func (u *usecase) Delete(ctx context.Context, id int64) error {
	existing, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrNotFound
	}
	return u.repo.Delete(ctx, id)
}