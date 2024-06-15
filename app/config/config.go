package config

import (
	"context"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/supersupersimple/comment/ent"
)

const (
	Setup        = "setup"
	Session      = "session"
	AllowOrigins = "allow_origins"

	LimitPerBatch = "limit_per_batch"
	MaxLoopDepth  = "max_loop_depth"
)

func setDefaultConfig() {
	viper.SetDefault(Setup, true)

	viper.SetDefault(Session, uuid.New().String())
	viper.SetDefault(AllowOrigins, []string{"*"})

	viper.SetDefault(LimitPerBatch, 5)
	viper.SetDefault(MaxLoopDepth, 3)
}

func LoadConfig(client *ent.Client) error {
	setDefaultConfig()
	conf, err := client.Conf.Query().Only(context.Background())
	if ent.IsNotFound(err) {
		viper.Set(Setup, false)
		return nil
	} else if err != nil {
		return err
	}
	updateConfig(conf)
	return nil
}

func updateConfig(conf *ent.Conf) {
	viper.Set(Setup, true)
	viper.Set(Session, conf.CookieSecret)
	viper.Set(LimitPerBatch, conf.LimitPerBatch)
	viper.Set(MaxLoopDepth, conf.MaxLoopDepth)
	viper.Set(AllowOrigins, conf.AllowOrigins)
}
