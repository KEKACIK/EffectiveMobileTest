package subscriptions

import "context"

type Repository interface {
	Create(ctx context.Context, dto *CreateSubscriptionDTO) (*Subscription, error)
	GetList(ctx context.Context, is_deleted bool) ([]Subscription, error)
	Get(ctx context.Context, id int, is_deleted bool) (*Subscription, error)
	Update(ctx context.Context, sub *Subscription) error
	Delete(ctx context.Context, id int) error
}
