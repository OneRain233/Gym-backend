package cron

import (
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/os/gcron"
)

func OrderExpired() error {
	// check in **:00:00 an d **:30:00
	ctx := gctx.New()
	_, err := gcron.Add(ctx, "0 0,30 * * * *", func(ctx context.Context) {
		err := service.Order().CheckExpiredOrder(ctx)
		if err != nil {
			return
		}
	}, "order_expired")
	if err != nil {
		return err
	}
	return nil

}
