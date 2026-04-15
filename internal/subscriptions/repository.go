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

func (repo *repository) Create(ctx context.Context, sub *Subscription) (*Subscription, error) {
	q := `
		CREATE subscriptions
			(TODO)
		VALUES
			()

	`
	repo.logger.DebugSQL(q)

	return nil, nil
}

func (repo *repository) GetList(ctx context.Context, is_deleted bool) ([]Subscription, error) {

	q := `
		SELECT
			TODO
		FROM subscriptions WHERE
			is_deleted=$1
	`
	repo.logger.DebugSQL(q)

	return []Subscription{}, nil
}

func (repo *repository) Get(ctx context.Context, id string, is_deleted bool) (*Subscription, error) {
	q := `
		SELECT
			TODO
		FROM subsctriptions WHERE
			id=$1 AND is_deleted=$2
	`
	repo.logger.DebugSQL(q)

	return nil, nil
}

func (repo *repository) Update(ctx context.Context, sub *Subscription) error {
	q := `
		UPDATE subsctriptions SET
			TODO
		WHERE
			id=$1
	`
	repo.logger.DebugSQL(q)

	return nil
}

func (repo *repository) Delete(ctx context.Context, sub *Subscription) error {
	// When deleting, change is_deleted to true
	q := `
		UPDATE subsctriptions SET
			is_deleted=true
		WHERE
			id=$1
	`
	repo.logger.DebugSQL(q)

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
