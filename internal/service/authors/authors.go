package authors

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"

	"eight/internal/models"
)

type HandlerAuthors struct {
	store store
}

func NewService(pqdriver *pgxpool.Pool, db *sql.DB) (*HandlerAuthors, error) {
	authorStore, err := newStore(pqdriver, db)
	if err != nil {
		return nil, err
	}

	return &HandlerAuthors{
		store: authorStore,
	}, nil
}

func (a *HandlerAuthors) AllAuthors(ctx context.Context) (models.AuthorSlice, error) {
	authors, err := a.store.All(ctx)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (a *HandlerAuthors) CreateAuthor(ctx context.Context, author *models.Author) (*models.Author,
	error) {
	return a.store.CreateAuthor(ctx, author)
}

func (a *HandlerAuthors) GetAuthor(ctx context.Context, authorID int64) (*models.Author, error) {
	return a.store.GetAuthor(ctx, authorID)
}