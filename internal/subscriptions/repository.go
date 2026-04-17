package subscriptions

import (
	"TestTask/pkg/logging"
	"TestTask/pkg/postgresql"
	"context"
	"errors"
	"fmt"
	"strings"
)

var (
	DatabaseNotContentErr error = errors.New("database invalid: Not found")
	DatabaseNotFoundErr   error = errors.New("database invalid: Not content")
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(
	client postgresql.Client,
	logger *logging.Logger,
) Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (repo *repository) Create(ctx context.Context, dto *SubscriptionCreateDTO) (*Subscription, error) {
	q := `
		INSERT INTO subscriptions
			(name, price, user_id, start_at, end_at)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id
	`
	repo.logger.DebugSQL(q)

	sub := Subscription{
		Name:    dto.Name,
		Price:   dto.Price,
		UserID:  dto.UserID,
		StartAt: dto.StartAt,
		EndAt:   dto.EndAt,
	}
	err := repo.client.QueryRow(ctx, q, dto.Name, dto.Price, dto.UserID, dto.StartAt, dto.EndAt).Scan(&sub.ID)
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (repo *repository) GetList(ctx context.Context, is_deleted bool) ([]Subscription, error) {
	q := `
		SELECT
			id, name, price, user_id, start_at, end_at
		FROM subscriptions WHERE
			is_deleted=$1
		ORDER BY id ASC
	`
	repo.logger.DebugSQL(q)

	rows, err := repo.client.Query(ctx, q, is_deleted)
	if err != nil {
		return nil, err
	}

	subs := make([]Subscription, 0)
	for rows.Next() {
		var sub Subscription

		err = rows.Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &sub.StartAt, &sub.EndAt)
		if err != nil {
			return nil, err
		}

		subs = append(subs, sub)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subs, nil
}

func (repo *repository) Get(ctx context.Context, id int, is_deleted bool) (*Subscription, error) {
	q := `
		SELECT
			id, name, price, user_id, start_at, end_at
		FROM subscriptions WHERE
			id=$1 AND is_deleted=$2
	`
	repo.logger.DebugSQL(q)

	sub := Subscription{}

	err := repo.client.QueryRow(ctx, q, id, is_deleted).Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &sub.StartAt, &sub.EndAt)
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (repo *repository) Update(ctx context.Context, dto *SubscriptionUpdateDTO) (*Subscription, error) {
	params := make([]string, 0)
	args := make([]any, 0)
	if dto.Name != "" {
		args = append(args, dto.Name)
		params = append(params, fmt.Sprintf("name=$%d", len(args)))
	}
	if dto.Price != 0 {
		args = append(args, dto.Price)
		params = append(params, fmt.Sprintf("price=$%d", len(args)))
	}
	if dto.UserID != "" {
		args = append(args, dto.UserID)
		params = append(params, fmt.Sprintf("user_id=$%d", len(args)))
	}
	if len(args) == 0 {
		return nil, DatabaseNotContentErr
	}
	params = append(params, "updated_at=CURRENT_TIMESTAMP")

	args = append(args, dto.ID)
	q := fmt.Sprintf(
		`UPDATE subscriptions SET %s WHERE id=$%d RETURNING %s`,
		strings.Join(params, ","),                    // Добавляем параметры после SET
		len(args),                                    // Добавляем id для WHERE
		"id, name, price, user_id, start_at, end_at", // Добавляем столбцы RETURNING
	)
	repo.logger.DebugSQL(q)

	sub := &Subscription{}
	err := repo.client.QueryRow(ctx, q, args...).Scan(
		&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &sub.StartAt, &sub.EndAt,
	)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (repo *repository) Delete(ctx context.Context, id int) error {
	// Для удаления изменяется параметр is_delete на True
	q := `
		UPDATE subscriptions SET
			is_deleted=true, updated_at=CURRENT_TIMESTAMP
		WHERE
			id=$1
	`
	repo.logger.DebugSQL(q)

	_, err := repo.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) Sum(ctx context.Context, dto *SubscriptionStatDTO, is_deleted bool) (int, error) {
	q := `
		SELECT
			SUM(price)
		FROM subscriptions WHERE
			user_id=$1 AND start_at >= $2 AND start_at < $3 AND is_deleted=$4
	`
	args := []any{dto.UserID, dto.StartDate, dto.StopDate, is_deleted}
	if dto.Name != "" {
		q += " AND name=$5"
		args = append(args, dto.Name)
	}
	repo.logger.DebugSQL(q)

	var sum int
	err := repo.client.QueryRow(ctx, q, args...).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}
