package db

import (
	"time"
)

type TransactionId = uint64

type Transaction struct {
	Id        TransactionId
	Title     string
	OwnerId   UserId
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type TransactionCreate struct {
	Title   string
	OwnerId uint
}

type TransactionUpdate struct {
	Id    TransactionId
	Title *string
}

func CreateTransaction(value TransactionCreate) (TransactionId, error) {

	result, err := DB.Exec(
		`
		INSERT INTO transactions
		(title, owner_id)
		VALUES (?, ?)
		`,
		value.Title, value.OwnerId,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return TransactionId(id), nil
}

func GetTransaction(id TransactionId) (*Transaction, error) {

	var row = DB.QueryRow(
		`
		SELECT
			id,
			owner_id,
			title,
			created_at,
			updated_at,
			deleted_at
		FROM transactions
		WHERE id = ?
			AND deleted_at IS NULL
		`,
		id,
	)

	var result Transaction
	var err = row.Scan(&result.Id, &result.OwnerId, &result.Title,
		&result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateTransaction(value TransactionUpdate) error {

	_, err := DB.Exec(
		`
		UPDATE transactions
		SET title = ?
		WHERE id = ?
			AND deleted_at IS NULL
		`,
		value.Title,
		value.Id,
	)

	return err
}

func DeleteTransaction(id TransactionId) error {

	_, err := DB.Exec(
		`
		UPDATE transactions
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE id = ?
			AND deleted_at IS NULL
		`,
		id,
	)

	return err
}
