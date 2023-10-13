package userrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
)

type UserRepoImpl struct {
}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

// Login implements UserRepo.
func (ur *UserRepoImpl) Login(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error) {
	SQL := "SELECT `id`, `nip`, `rank`, `email`, `password` FROM `user` WHERE email=? AND password=?"
	rows, err := tx.QueryContext(ctx, SQL, user.Email, user.Password)
	helper.PanicIfError(err)
	defer rows.Close()

	newUser := entity.User{}
	if rows.Next() {
		err := rows.Scan(&newUser.Id, &newUser.NIP, &newUser.Rank, &newUser.Email, &newUser.Password)
		helper.PanicIfError(err)
		return newUser, nil
	} else {
		return newUser, errors.New("user is not found")
	}
}

// Register implements UserRepo.
func (ur *UserRepoImpl) Register(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error) {
	user, err := ur.CheckEmailNIP(ctx, tx, user)
	helper.PanicIfError(err)
	SQL := "INSERT INTO `user`(`nip`, `rank`, `email`, `password`) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.NIP, user.Rank, user.Email, user.Password)
	if err != nil {
		return user, errors.New("can't insert new user")
	} else {
		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		user.Id = int(id)
		return user, nil
	}
}

// CheckEmailNIP implements UserRepo.
func (*UserRepoImpl) CheckEmailNIP(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error) {
	var exists bool
	SQL := "SELECT * FROM `user` WHERE email=?"
	query := tx.QueryRowContext(ctx, SQL, user.Email)

	err := query.Scan(&exists)
	helper.PanicIfError(err)

	if exists {
		return user, errors.New("email telah terdaftar")
	} else {
		SQL = "SELECT * FROM `user` WHERE nip=?"
		query = tx.QueryRowContext(ctx, SQL, user.NIP)

		err = query.Scan(&exists)
		helper.PanicIfError(err)

		if exists {
			return user, errors.New("nip telah terdaftar")
		} else {
			return user, nil
		}
	}
}

// Save implements UserRepo.
func (ur *UserRepoImpl) Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	SQL := "UPDATE user SET NIP=?, Rank=?, Email=?, Password=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, user.NIP, user.Rank, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

// Delete implements UserRepo.
func (ur *UserRepoImpl) Delete(ctx context.Context, tx *sql.Tx, user entity.User) {
	SQL := "delete from user where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

// FindById implements UserRepo.
func (*UserRepoImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error) {
	SQL := "select * from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.NIP, &user.Rank, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// FindAll implements UserRepo.
func (ur *UserRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.User {
	SQL := "select * from category where"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	users := []entity.User{}
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.NIP, &user.Rank, &user.Email, &user.Password)
		helper.PanicIfError(err)

		users = append(users, user)
	}
	return users
}

// FindByEmail implements UserRepo.
func (ur *UserRepoImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User, error) {
	SQL := "select * from category where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.NIP, &user.Rank, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// FindByNIP implements UserRepo.
func (ur *UserRepoImpl) FindByNIP(ctx context.Context, tx *sql.Tx, nip string) (entity.User, error) {
	SQL := "select * from category where nip = ?"
	rows, err := tx.QueryContext(ctx, SQL, nip)
	helper.PanicIfError(err)
	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.NIP, &user.Rank, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}
