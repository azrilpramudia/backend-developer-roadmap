package postgres

import (
	"context"
	"database/sql"

	"github.com/azrilpramudia/go-restful-api/internal/product"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) product.Repository {
	return &productRepo{db: db}
}

func (r *productRepo) Create(ctx context.Context, p *product.Product) error {
	query := `INSERT INTO products (name, price, stock, created_at, updated_at)
						VALUES ($1, $2, $3, now(), now())
						RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query, p.Name, p.Price, p.Stock).
		Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func (r *productRepo) GetByID(ctx context.Context, id int64) (*product.Product, error) {
	query := `SELECT id, name, price, stock, created_at, updated_at
						FROM products WHERE id = $1`
	p := &product.Product{}
	err := r.db.QueryRowContext(ctx, query, id).
			Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return p, nil
} 

func (r *productRepo) GetAll(ctx context.Context) ([]*product.Product, error) {
	query := `SELECT id, name, price, stock, created_at, updated_at FROM products ORDER BY id`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*product.Product
	for rows.Next() {
		p := &product.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *productRepo) Update(ctx context.Context, p *product.Product) error {
	query := `UPDATE products SET name=$1, price=$2, stock=$3, updated_at=now()
						WhERE id=$4 RETURNING updated_at`
	return r.db.QueryRowContext(ctx, query, p.Name, p.Price, p.Stock, p.ID).Scan(&p.UpdatedAt)
}

func (r *productRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM products WHERE id=$1`, id)
	return err
}