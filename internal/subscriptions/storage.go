package subscriptions

import "context"

type Repository interface {
	Create(ctx context.Context, sub *Subscription) (*Subscription, error)
	GetList(ctx context.Context, is_deleted bool) ([]Subscription, error)
	Get(ctx context.Context, id string, is_deleted bool) (*Subscription, error)
	Update(ctx context.Context, sub *Subscription) error
	Delete(ctx context.Context, sub *Subscription) error
}
