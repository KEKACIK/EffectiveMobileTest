package subscriptions

import "context"

type Repository interface {
	Create(ctx context.Context, dto *SubscriptionCreateDTO) (*Subscription, error)
	GetList(ctx context.Context, is_deleted bool) ([]Subscription, error)
	Get(ctx context.Context, id int, is_deleted bool) (*Subscription, error)
	Update(ctx context.Context, dto *SubscriptionUpdateDTO) (*Subscription, error)
	Delete(ctx context.Context, id int) error
}
