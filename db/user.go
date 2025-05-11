package db

type UserId = uint64

type User struct {
	Id    UserId
	Name  string
	Email string
	Phone *string
}

type UserCreate struct {
	Name     string
	Email    string
	Phone    *string
	Password string
}

type UserUpdate struct {
	Id       UserId
	Name     *string
	Email    *string
	Phone    *string
	Password *string
}

func CreateUser(value UserCreate) (UserId, error) {

	tx, err := DB.Begin()

	if err != nil {
		return 0, err
	}

	result, err := DB.Exec(
		`
		INSERT INTO users
		(name, email, phone)
		VALUES ($1, $2, $3)
		`,
		value.Name,
		value.Email,
		value.Phone,
	)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = DB.Exec(
		`
		INSERT INTO secrets
		(user_id, password)
		VALUES ($1, $2)
		`,
		id,
		value.Password,
	)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return UserId(id), nil
}

func GetUser(id UserId) (*User, error) {

	row := DB.QueryRow(
		`
		SELECT
			name,
			email,
			phone
		FROM users
		WHERE id = $1
			AND deleted_at IS NULL
		`,
		id,
	)

	var result User
	result.Id = id

	err := row.Scan(&result.Name, &result.Email, &result.Phone)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateUser(value UserUpdate) error {

	_, err := DB.Exec(
		`
		UPDATE users
		SET name = $1,
			email = $2,
			phone = $3
		WHERE id = $4,
			AND deleted_at IS NULL
		`,
		value.Name,
		value.Email,
		value.Phone,
		value.Id,
	)

	return err
}

func DeleteUser(id UserId) error {

	_, err := DB.Exec(
		`
		UPDATE users
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
			AND deleted_at IS NULL
		`,
		id,
	)

	return err
}
