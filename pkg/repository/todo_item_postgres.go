package repository

import (
	"TODOapi"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) CreateItem(listId int, item TODOapi.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAllItems(userId, listId int) ([]TODOapi.TodoItem, error) {
	var items []TODOapi.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.lists_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	err := r.db.Select(&items, query, listId, userId)
	if err != nil {
		return nil, err
	}

	return items, nil
}
