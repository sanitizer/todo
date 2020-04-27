package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sanitizer/todo/model"
	"os"
	"time"
)

const (
	GET_ALL = `
		SELECT *
		FROM todo
		WHERE NOT is_deleted
		AND CASE WHEN $1
			THEN status = LOWER($2)
			ELSE TRUE
		END
		ORDER BY name
	`
	SOFT_DELETE = `
		UPDATE todo
		SET is_deleted = TRUE,
			update_dt = NOW()
		WHERE id = $1
		RETURNING id
	`
	CREATE = `
		INSERT INTO todo (
			name,
			description
		)
		VALUES (
			$1,
			$2
		)
		RETURNING *
	`
	UPDATE = `
		UPDATE todo
		SET name = $1,
			description = $2,
			update_dt = NOW()
		WHERE id = $3
		AND NOT is_deleted
		RETURNING *
	`
	COMPLETE = `
		UPDATE todo
		SET status = 'complete',
			update_dt = NOW()
		WHERE id = $1
		AND NOT is_deleted
		RETURNING *
	`
)

func Connection() *sql.DB {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"password=postgres host=%s dbname=postgres port=5432 sslmode=disable",
			os.Getenv("DB_HOST"),
		),
	)
	checkForErrorAndFail(err)

	return db
}

func GetAll(status string) []model.Todo {
	rows, err := query(
		GET_ALL,
		status != "",
		status,
	)
	defer rows.Close()

	checkForErrorAndFail(err)

	var todos []model.Todo
	for rows.Next() {
		todos = append(todos, *mapPojo(rows))
	}

	return todos
}

func Delete(todoId int64) int64 {
	rows, err := query(SOFT_DELETE, todoId)
	defer rows.Close()

	checkForErrorAndFail(err)

	for rows.Next() {
		return getId(rows)
	}

	return 0
}

func Create(todo model.Todo) *model.Todo {
	rows, err := query(
		CREATE,
		todo.Name,
		todo.Description,
	)
	defer rows.Close()

	checkForErrorAndFail(err)

	for rows.Next() {
		return mapPojo(rows)
	}

	return nil
}

func Update(todoId int64, todo model.Todo) *model.Todo {
	rows, err := query(
		UPDATE,
		todo.Name,
		todo.Description,
		todoId,
	)
	defer rows.Close()

	checkForErrorAndFail(err)

	for rows.Next() {
		return mapPojo(rows)
	}

	return nil
}

func Complete(todoId int64) *model.Todo {
	rows, err := query(
		COMPLETE,
		todoId,
	)
	defer rows.Close()

	checkForErrorAndFail(err)

	for rows.Next() {
		return mapPojo(rows)
	}

	return nil
}

func getId(rows *sql.Rows) int64 {
	var id int64
	err := rows.Scan(&id)

	checkForErrorAndFail(err)

	return id
}

func mapPojo(rows *sql.Rows) *model.Todo {
	var id int64
	var name string
	var description string
	var status string
	var isDeleted bool
	var createDt time.Time
	var updateDt time.Time

	err := rows.Scan(
		&id,
		&name,
		&description,
		&status,
		&isDeleted,
		&createDt,
		&updateDt,
	)

	checkForErrorAndFail(err)

	return &model.Todo {
		Id: id,
		Name: name,
		Description: description,
		Status: status,
		IsDeleted: isDeleted,
		CreateDt: createDt,
		UpdateDt: updateDt,
	}
}

func query(query string, args ...interface{}) (*sql.Rows, error) {
	db := Connection()
	defer db.Close()
	statement, err := db.Prepare(query)
	checkForErrorAndFail(err)

	return statement.Query(args...)
}
