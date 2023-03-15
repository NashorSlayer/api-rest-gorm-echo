package database

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // Standard library bindings for pgx
	"github.com/jmoiron/sqlx"
	"github.com/rarifbkn/api-rest-gorm-echo/settings"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host = %s user= %s password=%s dbname =%s port=%d ",
		s.DB.Host,
		s.DB.User,
		s.DB.Password,
		s.DB.Name,
		s.DB.Port,
	)
	return sqlx.ConnectContext(ctx, "pgx", connectionString)
}
