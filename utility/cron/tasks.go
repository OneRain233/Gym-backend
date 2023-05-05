package cron

import (
	"Gym-backend/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/os/gcron"
)

func OrderExpired() error {
	// check in **:00:00 an d **:30:00
	ctx := gctx.New()
	_, err := gcron.Add(ctx, "*/60 * * * * *", func(ctx context.Context) {
		err := service.Order().CheckExpiredOrder(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
	}, "order_expired")
	if err != nil {
		return err
	}
	return nil
}

func SubscriptionExpired() error {
	ctx := gctx.New()
	// 00:00 every day
	_, err := gcron.Add(ctx, "0 0 0 * * *", func(ctx context.Context) {
		err := service.Subscription().CheckExpiredSubscription(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	if err != nil {
		return err
	}
	return nil
}
