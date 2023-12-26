package db

import (
	"bd-backend/internal/admin/category"
	"bd-backend/pkg/logging"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	client *pgxpool.Pool
	logger *logging.Logger
}

func NewRepository(client *pgxpool.Pool, logger *logging.Logger) category.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r repository) GetAllData(ctx context.Context, limit string) ([]category.CategoryDTO, error) {

	var (
		ns []category.CategoryDTO
	)

	q := `select
				id,value
			 from timtables
			 order by id desc
			 limit $1;`

	rows, err := r.client.Query(ctx, q, limit)
	defer rows.Close()

	if err != nil {
		return ns, nil
	}

	for rows.Next() {
		var n category.CategoryDTO
		errN := rows.Scan(
			&n.UUID, &n.Value,
		)
		if errN != nil {
			r.logger.Error("errN :::", errN)
			continue
		}

		ns = append(ns, n)
	}
	return ns, nil
}

func (r repository) AddData(ctx context.Context, dt []interface{}) error {

	q := `INSERT INTO timtables (value) VALUES ($1);`

	_, err := r.client.Exec(ctx, q, dt)
	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}
