package repository

import (
	"context"
	"fmt"

	"service/internal/gallery/entity"
)

func (r *Repo) GetPhotos(ctx context.Context) ([]*entity.Photo, error) {
	q := "SELECT * FROM photo"

	var row []*entity.Photo

	if err := r.replica.SelectContext(ctx, &row, q); err != nil {
		return nil, fmt.Errorf("db query failed: %w", err)
	}

	return row, nil
}
