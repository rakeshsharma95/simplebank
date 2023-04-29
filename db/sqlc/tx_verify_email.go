package db

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
)

type VerifyEmailTxParams struct {
	EmailId    int64
	SecretCode string
}

type VerifyEmailTxResult struct {
	User        User
	VerifyEmail VerifyEmail
}

func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		log.Info().Int64("id", arg.EmailId).Err(err).
			Str("secret_code", arg.SecretCode).Msg("verify data email  UpdateVerifyEmail")
		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		log.Err(err).Msg("verify data email  UpdateVerifyEmail")
		if err != nil {
			return err
		}

		// log.Info().Int64("id", result.VerifyEmail.ID).Err(err).
		// 	Str("email", result.VerifyEmail.Email).Msg("verify data email  UpdateVerifyEmail")

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Username: result.VerifyEmail.Username,
			IsEmailVerified: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		})

		return err
	})

	return result, err
}
