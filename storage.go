package auth

import (
	"github.com/joaosoft/dbr"
)

type StoragePostgres struct {
	config *AuthConfig
	db     *dbr.Dbr
}

func NewStoragePostgres(config *AuthConfig) (*StoragePostgres, error) {
	dbr, err := dbr.New(dbr.WithConfiguration(config.Dbr))
	if err != nil {
		return nil, err
	}

	return &StoragePostgres{
		config: config,
		db:     dbr,
	}, nil
}

func (storage *StoragePostgres) GetUserByIdUserAndRefreshToken(idUser, refreshToken string) (*User, error) {
	user := &User{}
	count, err := storage.db.
		Select("*").
		From(authTableUser).
		Where("id_user = ?", idUser).
		Where("refresh_token = ?", refreshToken).
		Where("active").
		Load(user)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	return user, nil
}

func (storage *StoragePostgres) GetUserByEmailAndPassword(email, password string) (*User, error) {
	user := &User{}
	count, err := storage.db.
		Select("*").
		From(authTableUser).
		Where("email = ?", email).
		Where("password_hash = ?", password).
		Where("active").
		Load(user)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	return user, nil
}

func (storage *StoragePostgres) UpdateUserRefreshToken(idUser, refreshToken string) error {
	result, err := storage.db.
		Update(authTableUser).
		Set("refresh_token", refreshToken).
		Where("id_user = ?", idUser).
		Where("active").
		Exec()

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrorNotFound
	}

	return nil
}
