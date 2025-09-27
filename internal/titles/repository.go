package titles

import (
	"context"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"ss-api/pkg/util"
)

type TitlesRepository struct {
	log  *zap.Logger
	conn *pgx.Conn
}

func NewTitlesRepository(log *zap.Logger, conn *pgx.Conn) *TitlesRepository {
	return &TitlesRepository{log: log, conn: conn}
}

func (r *TitlesRepository) FindAll() ([]Title, error) {
	sql := `SELECT id, name, id_author FROM titles`

	items := make([]Title, 0)
	err := pgxscan.Select(context.Background(), r.conn, &items, sql)

	return items, err
}

func (r *TitlesRepository) FindById(id string) (Title, error) {
	sql := `SELECT id, name, id_author FROM titles WHERE id = $1`

	var item Title
	err := pgxscan.Get(context.Background(), r.conn, &item, sql, id)

	if errors.Is(err, pgx.ErrNoRows) {
		return Title{}, fmt.Errorf("find title by id %s: %w", id, util.ErrNotFound)
	}

	return item, err
}

func (r *TitlesRepository) Create(dto CreateTitleDto) (Title, error) {
	sql := `INSERT INTO titles (name, id_author) VALUES ($1, $2) RETURNING id, name, id_author`

	var item Title
	err := pgxscan.Get(context.Background(), r.conn, &item, sql, dto.Name, dto.IdAuthor)

	return item, err
}

func (r *TitlesRepository) UpdateById(id string, dto UpdateTitleDto) (Title, error) {
	sql := `
		UPDATE titles
    SET
				name = coalesce($2, name),
				id_author = coalesce($2, id_author),
    WHERE id = $1
    RETURNING id, name, id_author;
	`

	var item Title
	err := pgxscan.Get(
		context.Background(),
		r.conn,
		&item,
		sql,
		id,
		dto.Name,
		dto.IdAuthor,
	)

	return item, err
}

func (r *TitlesRepository) DeleteById(id string) error {
	sql := `DELETE FROM titles WHERE id = $1`

	_, err := r.conn.Exec(context.Background(), sql, id)

	return err
}
