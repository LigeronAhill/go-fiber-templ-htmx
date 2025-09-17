package vacancy

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

const limit = 10

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

func (r *Repository) GetAll(page int) ([]*Vacancy, error) {
	offset := (page - 1) * limit
	query := `
		SELECT * FROM vacancies ORDER BY created_at DESC LIMIT @limit OFFSET @offset;
	`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := r.pool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[Vacancy])
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (r *Repository) TotalPages() (int, error) {
	query := "SELECT COUNT(*) FROM vacancies;"
	row := r.pool.QueryRow(context.Background(), query)
	var totalCount int
	err := row.Scan(&totalCount)
	if err != nil {
		r.log.Error().Err(err).Msg("Failed to count vacancies")
		return 0, err
	}
	totalPages := (totalCount + limit - 1) / limit
	return totalPages, nil
}
