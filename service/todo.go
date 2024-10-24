package service

import (
	"context"
	"database/sql"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
			insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
			confirm = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)


	// トランザクションの開始
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
			return nil, err
	}
	defer tx.Rollback()

	// TODO の挿入
	result, err := tx.ExecContext(ctx, insert, subject, description)
	if err != nil {
			return nil, err
	}

	// 挿入したレコードの ID を取得
	id, err := result.LastInsertId()
	if err != nil {
			return nil, err
	}

	// 挿入した TODO を取得
	row := tx.QueryRowContext(ctx, confirm, id)
	var todo model.TODO
	if err := row.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
	}

	// トランザクションのコミット
	if err := tx.Commit(); err != nil {
			return nil, err
	}

	// 作成した TODO を返す
	return &todo, nil
}


// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	return nil, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
