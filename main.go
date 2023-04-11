package main

import (
	_ "Gym-backend/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"Gym-backend/internal/cmd"
	_ "Gym-backend/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
