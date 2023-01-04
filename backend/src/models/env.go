package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	DB *pgxpool.Pool
}
