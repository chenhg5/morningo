package config

import "time"

type JwtConfig struct {
	SECRET string
	EXP    time.Duration // 过期时间
	ALG    string        // 算法
}

func GetJwtConfig() *JwtConfig {
	return &JwtConfig{
		SECRET: "morningo_session",
		EXP:    time.Hour,
		ALG:    "HS256",
	}
}
