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
