package category

import "context"

type Repository interface {
	GetAllData(ctx context.Context, limit string) ([]CategoryDTO, error)
	AddData(ctx context.Context, dt []interface{}) error
}
