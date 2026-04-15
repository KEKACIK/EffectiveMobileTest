package subscriptions

import (
	"TestTask/pkg/logging"
	"TestTask/pkg/postgresql"
	"context"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (repo *repository) Create(ctx context.Context, dto *CreateSubscriptionDTO) (*Subscription, error) {
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

func (repo *repository) Get(ctx context.Context, id string, is_deleted bool) (*Subscription, error) {
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

func (repo *repository) Update(ctx context.Context, sub *Subscription) error {
	q := `
		UPDATE subscriptions SET
			name=$1, price=$2, user_id=$3, start_at=$4, end_at=$5
		WHERE
			id=$6
	`
	repo.logger.DebugSQL(q)

	_, err := repo.client.Exec(ctx, q, sub.Name, sub.Price, sub.UserID, sub.StartAt, sub.EndAt, sub.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) Delete(ctx context.Context, id int) error {
	// When deleting, change is_deleted to true
	q := `
		UPDATE subscriptions SET
			is_deleted=true
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

func NewRepository(
	client postgresql.Client,
	logger *logging.Logger,
) Repository {

	return &repository{
		client: client,
		logger: logger,
	}
}
