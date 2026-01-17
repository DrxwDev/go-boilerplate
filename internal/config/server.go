package config

import "time"

type ServerConfig struct {
	ReadTimeout       time.Duration
	WriteTimout       time.Duration
	IdleTimeot        time.Duration
	ReadHeaderTimeout time.Duration
}
