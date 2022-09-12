package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/topic-management-service/config"
	"github.com/hexcraft-biz/topic-management-service/features"
)

type TestStruct struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

func main() {
	cfg, _ := config.Load()
	cfg.DBOpen(false)
	cfg.InitRedis()

	engine := gin.Default()
	engine.SetTrustedProxies([]string{cfg.Env.TrustProxy})

	// Base features
	features.LoadCommon(engine, cfg)
	// Bridging features
	features.LoadBridging(engine, cfg)

	engine.Run(":" + cfg.Env.AppPort)
}
