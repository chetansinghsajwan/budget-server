package db

import (
	"time"

	"github.com/lib/pq"
)

type TransactionId = uint64

type Transaction struct {
	Id        TransactionId
	Title     string
	OwnerId   UserId
	Amount    uint64
	IsCredit  bool
	Time      *time.Time
	Tags      pq.StringArray
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type TransactionCreate struct {
	Title    string
	OwnerId  uint
	Amount   uint64
	IsCredit bool
	Time     *time.Time
	Tags     pq.StringArray
}

type TransactionUpdate struct {
	Id       TransactionId
	Title    *string
	Amount   *uint64
	IsCredit *bool
	Time     **time.Time
	Tags     *pq.StringArray
}

func CreateTransaction(value TransactionCreate) (TransactionId, error) {

	result, err := DB.Exec(
		`
		INSERT INTO transactions
		(title, owner_id, amount, is_credit, time, tags)
		VALUES ($1, $2, $3, $4, $5, $6)
		`,
		value.Title,
		value.OwnerId,
		value.Amount,
		value.IsCredit,
		value.Time,
		value.Tags,
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
			owner_id,
			title,
			(amount::numeric)::bigint as amount,
			is_credit,
			time,
			tags,
			created_at,
			updated_at,
			deleted_at
		FROM transactions
		WHERE id = $1
			AND deleted_at IS NULL
		`,
		id,
	)

	var result Transaction
	result.Id = id

	var err = row.Scan(&result.OwnerId, &result.Title,
		&result.Amount, &result.IsCredit, &result.Time, &result.Tags,
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
		SET title = $1
		WHERE id = $2
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
		WHERE id = $1
			AND deleted_at IS NULL
		`,
		id,
	)

	return err
}
