package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create(username string, email string, password string) error
	GetByID() (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteByID(id int64) error
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, username, email, hashedPassword)
	if err != nil {
		fmt.Println("Error inserting user: ", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error getting rows affected: ", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil
	}

	fmt.Println("User created successfully, rows affected", rowsAffected)

	return nil
}

func (u *UserRepositoryImpl) GetByID() (*models.User, error) {
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"

	row := u.db.QueryRow(query, 1)

	fmt.Println("row.....", row)
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
	}

	fmt.Println("User fetched successfully:", user)

	return user, nil
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	query := "SELECT id, username, email, created_at, updated_at from users"

	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Failed to executing query:", err)
		return nil, err
	}

	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			fmt.Println("Error scanning row: ", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Row iteration error: ", err)
		return nil, err
	}

	fmt.Printf("✅ Successfully fetched %d users\n", len(users))
	return users, nil
}

func (u *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = ?"

	row := u.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password) //hashed password
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, err
		} else {
			fmt.Println("Error scanning user: ", err)
			return nil, err
		}
	}

	return user, nil
}

func (u *UserRepositoryImpl) DeleteByID(id int64) error {
	return nil
}
