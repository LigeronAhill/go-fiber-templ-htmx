package vacancy

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Repository struct {
	pool *pgxpool.Pool
	log  *zerolog.Logger
}

func NewRepo(pool *pgxpool.Pool, log *zerolog.Logger) *Repository {
	return &Repository{
		pool,
		log,
	}
}

func (r *Repository) AddVacancy(form *VacancyFormCreate) error {
	query := `
	INSERT INTO vacancies (role, company, type, salary, location, email, created_at)
	VALUES (@role, @company, @type, @salary, @location, @email, @created_at)
	`
	args := pgx.NamedArgs{
		"role":       form.Role,
		"company":    form.Company,
		"type":       form.Type,
		"salary":     form.Salary,
		"location":   form.Location,
		"email":      form.Email,
		"created_at": time.Now(),
	}
	_, err := r.pool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}
	return nil
}
