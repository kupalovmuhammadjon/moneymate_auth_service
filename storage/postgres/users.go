package postgres

import (
	"auth_service/pkg/logger"
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"

	pb "auth_service/genproto/users"
)

type usersRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewUsersRepo(db *pgxpool.Pool, log logger.ILogger) *usersRepo {
	return &usersRepo{
		db:  db,
		log: log,
	}
}

func (u *usersRepo) GetUserProfile(ctx context.Context, request *pb.PrimaryKey) (*pb.User, error) {

	var (
		user      = pb.User{}
		query     string
		err       error
		firstName sql.NullString
		lastName  sql.NullString
		updatedAt sql.NullString
	)

	query = `
	select
		id,
		username,
		email,
		first_name,
		last_name,
		role,
		created_at::text,
		updated_at::text
	from 
		users 
	where
		id = $1
	`

	if err = u.db.QueryRow(ctx, query, request.Id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&firstName,
		&lastName,
		&user.Role,
		&user.CreatedAt,
		&updatedAt,
	); err != nil {
		u.log.Error("error while getting user id by username", logger.Error(err))
		return nil, err
	}
	user.FirstName = firstName.String
	user.LastName = lastName.String
	user.UpdatedAt = updatedAt.String

	return &user, nil
}

func (u *usersRepo) UpdateUserProfile(ctx context.Context, request *pb.UpdateUser) (*pb.UpdateProfileResponce, error) {

	var (
		user      = pb.UpdateProfileResponce{}
		query     string
		err       error
		updatedAt = sql.NullString{}
	)

	query = `
	update users set
		username   = $1,
		email      = $2,
		first_name = $3,
		last_name  = $4
	where id = $5 
	returning
		id,
		username,
		email,
		first_name,
		last_name,
		role,
		created_at::text,
		updated_at::text
	`

	if err = u.db.QueryRow(ctx, query,
		request.GetUsername(),
		request.GetEmail(),
		request.GetFirstName(),
		request.GetLastName(),
		request.GetId(),
	).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.CreatedAt,
		&updatedAt,
	); err != nil {
		u.log.Error("error while updating user info in storage layer", logger.Error(err))
		return nil, err
	}

	user.UpdatedAt = updatedAt.String

	return &user, nil
}

func (u *usersRepo) CheckPasswordExisis(ctx context.Context, request *pb.ChangePassword) (bool, error) {

	var (
		exists = sql.NullInt64{}
		query  string
		err    error
	)

	query = `
		select
			1
		from
			users
		where
			id = $1 and password_hash = $2
	`

	if err = u.db.QueryRow(ctx, query,
		request.GetUserId(),
		request.GetCurrentPassword()).Scan(
		&exists,
	); err != nil {
		u.log.Error("error while checking current password is currect in storage layer", logger.Error(err))
		return false, err
	}

	if exists.Int64 == 0 {
		return false, nil
	}
	return true, nil
}

func (u *usersRepo) ChangePassword(ctx context.Context, request *pb.ChangePassword) (*pb.Message, error) {

	var (
		query string
		err   error
	)

	query = `
		update 
			users 
		set 
			password_hash = $1
		where
			id = $2
	`

	if _, err = u.db.Exec(ctx, query,
		request.GetNewPassword(),
		request.GetUserId(),
	); err != nil {
		u.log.Error("error while changing password  in storage layer", logger.Error(err))
		return nil, err
	}

	return &pb.Message{
		Message: "assword successfully changed",
	}, nil
}
