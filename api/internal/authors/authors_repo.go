package authors

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type AuthorsRepository struct {
	log  *zap.Logger
	conn *pgx.Conn
}

func NewAuthorsRepository(log *zap.Logger, conn *pgx.Conn) *AuthorsRepository {
	return &AuthorsRepository{log: log, conn: conn}
}

func (r *AuthorsRepository) FindAll() ([]Author, error) {
	sql := `SELECT id, name FROM authors`

	items := make([]Author, 0)
	err := pgxscan.Select(context.Background(), r.conn, &items, sql)

	return items, err
}

func (r *AuthorsRepository) FindById(id string) (Author, error) {
	sql := `SELECT id, name FROM authors WHERE id = $1`

	var item Author
	err := pgxscan.Get(context.Background(), r.conn, &item, sql, id)

	return item, err
}

func (r *AuthorsRepository) Create(dto CreateAuthorDto) (Author, error) {
	sql := `INSERT INTO authors (name, bio) VALUES ($1, $2) RETURNING id, name, bio`

	var item Author
	err := pgxscan.Get(context.Background(), r.conn, &item, sql, dto.Name, dto.Bio)

	return item, err
}

func (r *AuthorsRepository) UpdateById(id string, dto UpdateAuthorDto) (Author, error) {
	sql := `
		UPDATE authors
    SET
        name = CASE WHEN $2 THEN $3 ELSE name END,
				bio = CASE WHEN $4 THEN $5 ELSE bio END
    WHERE id = $1
    RETURNING id, name, bio;
	`

	var item Author
	err := pgxscan.Get(
		context.Background(),
		r.conn,
		&item,
		sql,
		id,
		dto.Name != nil,
		dto.Name,
		dto.Bio != nil,
		dto.Bio,
	)

	return item, err
}

func (r *AuthorsRepository) DeleteById(id string) error {
	sql := `DELETE FROM authors WHERE id = $1`

	_, err := r.conn.Exec(context.Background(), sql, id)

	return err
}
