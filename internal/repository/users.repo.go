package repo

import (
	"database/sql"

	"restaurant-management/internal/models"
)

type UserRepo interface {
	EmailExists(email string) (bool, error)
	PhoneExists(phone_number string) (bool, error)

	Add(user *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetById(userId string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Edit(userId string, user *models.User) (*models.User, error)
	Delete(userId string) (err error)
}

type userSql struct {
	conn *sql.DB
}

func (u *userSql) EmailExists(email string) (bool, error) {
	var userId string

	query := `SELECT id FROM users WHERE email = $1;`

	err := u.conn.QueryRow(query, email).Scan(&userId)

	if err != nil {
		if err == sql.ErrNoRows {
			// Email does not exist
			return false, nil
		}
		// An error occurred while executing the query
		return true, err
	}

	// Email already exists
	return true, nil
}

func (u *userSql) PhoneExists(phone_number string) (bool, error) {
	var userId string

	query := `SELECT id FROM users WHERE phone_number = $1;`

	err := u.conn.QueryRow(query, phone_number).Scan(&userId)

	if err != nil {
		if err == sql.ErrNoRows {
			// Email does not exist
			return false, nil
		}
		// An error occurred while executing the query
		return true, err
	}

	// Email already exists
	return true, nil
}

func (u *userSql) Add(user *models.User) (*models.User, error) {
	return u.addUser(user)
}

func (u *userSql) GetByEmail(email string) (usr *models.User, err error) {
	usr = new(models.User)

	query := `SELECT id, email, password, first_name, last_name, phone_number, role, date_created, date_updated FROM users WHERE email = $1`

	err = u.conn.QueryRow(query, email).Scan(&usr.Id, &usr.Email, &usr.Password, &usr.FirstName, &usr.LastName, &usr.PhoneNumber, &usr.Role, &usr.DateCreated, &usr.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (u *userSql) GetById(userId string) (usr *models.User, err error) {
	usr = new(models.User)

	query := `SELECT id, email, password, first_name, last_name, phone_number, role, date_created, date_updated FROM users WHERE id = $1`

	err = u.conn.QueryRow(query, userId).Scan(&usr.Id, &usr.Email, &usr.Password, &usr.FirstName, &usr.LastName, &usr.PhoneNumber, &usr.Role, &usr.DateCreated, &usr.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (u *userSql) GetAll() ([]*models.User, error) {
	var users []*models.User

	query := `SELECT id, email, password, first_name, last_name, phone_number, role, date_created, date_updated FROM users`

	rows, err := u.conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Role, &user.DateCreated, &user.DateUpdated)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, err
}

func (u *userSql) Edit(userId string, user *models.User) (usr *models.User, err error) {
	usr = new(models.User)

	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, phone_number = $4, date_updated = CURRENT_TIMESTAMP WHERE id = $5 RETURNING id, first_name, last_name, email, phone_number, password, role, date_created, date_updated`

	err = u.conn.QueryRow(query, user.FirstName, user.LastName, user.Email, user.PhoneNumber, userId).Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.PhoneNumber, &usr.Password, &usr.Role, &usr.DateCreated, &usr.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (u *userSql) Delete(userId string) (err error) {
	query := `DELETE FROM users WHERE id = $1`

	_, err = u.conn.Exec(query, userId)
	if err != nil {
		return
	}

	return
}

func NewUserRepo(conn *sql.DB) UserRepo {
	return &userSql{conn: conn}
}

func (u *userSql) addUser(user *models.User) (usr *models.User, err error) {
	usr = new(models.User)

	if user.Role != "ADMIN" {
		query := `INSERT INTO users(first_name, last_name, email, phone_number, password) VALUES ($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, email, phone_number, password, role, date_created, date_updated`

		err = u.conn.QueryRow(query, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Password).Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.PhoneNumber, &usr.Password, &usr.Role, &usr.DateCreated, &usr.DateUpdated)
		if err != nil {
			return
		}

		return
	}

	query := `INSERT INTO users(first_name, last_name, email, phone_number, password, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, first_name, last_name, email, phone_number, password, role, date_created, date_updated`

	err = u.conn.QueryRow(query, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Password, user.Role).Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.PhoneNumber, &usr.Password, &usr.Role, &usr.DateCreated, &usr.DateUpdated)
	if err != nil {
		return
	}

	return
}
