/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("user not found")

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(
	ctx context.Context,
	u *User,
) error {
	return r.db.QueryRow(
		ctx,
		`
		INSERT INTO streamforge.users (
			id,
			kratos_identity_id,
			email,
			first_name,
			last_name,
			email_verified
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING created_at, updated_at
		`,
		u.ID,
		u.KratosIdentityID,
		u.Email,
		u.FirstName,
		u.LastName,
		u.EmailVerified,
	).Scan(
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id string,
) (*User, error) {
	var u User

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			kratos_identity_id,
			email,
			first_name,
			last_name,
			email_verified,
			created_at,
			updated_at
		FROM streamforge.users
		WHERE id = $1
		`,
		id,
	).Scan(
		&u.ID,
		&u.KratosIdentityID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.EmailVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *PostgresRepository) GetByKratosIdentityID(
	ctx context.Context,
	kratosIdentityID string,
) (*User, error) {
	var u User

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			kratos_identity_id,
			email,
			first_name,
			last_name,
			email_verified,
			created_at,
			updated_at
		FROM streamforge.users
		WHERE kratos_identity_id = $1
		`,
		kratosIdentityID,
	).Scan(
		&u.ID,
		&u.KratosIdentityID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.EmailVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *PostgresRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*User, error) {
	var u User

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			kratos_identity_id,
			email,
			first_name,
			last_name,
			email_verified,
			created_at,
			updated_at
		FROM streamforge.users
		WHERE email = $1
		`,
		email,
	).Scan(
		&u.ID,
		&u.KratosIdentityID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.EmailVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}
