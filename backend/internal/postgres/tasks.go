package postgres

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"github.com/jackc/pgx/v5"
)

func (p *postgres) CreateTask(ctx context.Context, tx pgx.Tx, task *domain.Task) (int, error) {
	q1 := `insert into tasks (title, description, body, revenue, type, priority, status, thumbnail) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`
	row := tx.QueryRow(ctx, q1, task.Title, task.Description, task.Body, task.Revenue, task.Type, task.Priority, task.Status, task.Thumbnail)

	var taskId int
	if err := row.Scan(&taskId); err != nil {
		return 0, err
	}

	return taskId, nil
}

func (p *postgres) GetAllTasks(ctx context.Context, tx pgx.Tx) ([]*domain.Task, error) {
	q1 := `select * from tasks`
	rows, err := tx.Query(ctx, q1)
	if err != nil {
		return nil, err
	}

	var tasks []*domain.Task

	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		if err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.Body, &task.Revenue, &task.Type, &task.Priority, &task.Status, &task.Thumbnail); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (p *postgres) DeleteTask(ctx context.Context, tx pgx.Tx, taskId int) error {
	q1 := `delete from tasks where id=$1`
	_, err := tx.Exec(ctx, q1, taskId)
	return err
}
