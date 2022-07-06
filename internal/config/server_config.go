package config

import (
	"runtime"
	"time"

	"github.com/nintran52/one-talent-tutorial/internal/util"
)

type EchoServer struct {
	ListenAddress string
}

type Server struct {
	Echo     EchoServer
	Database Database
}

func DefaultServerConfigFromEnv() Server {
	return Server{
		Database: Database{
			Host:     util.GetEnv("PGHOST", "localhost"),
			Port:     util.GetEnvAsInt("PGPORT", 5432),
			Username: util.GetEnv("PGUSER", "dbuser"),
			Password: util.GetEnv("PGPASSWORD", " dbpass"),
			Database: util.GetEnv("PGDATABASE", "development"),
			AdditionalParams: map[string]string{
				"sslmode": util.GetEnv("PGSSLMODE", "disable"),
			},
			MaxOpenConns:    util.GetEnvAsInt("DB_MAX_OPEN_CONNS", runtime.NumCPU()*2),
			MaxIdleConns:    util.GetEnvAsInt("DB_MAX_IDLE_CONNS", 1),
			ConnMaxLifetime: time.Second * time.Duration(util.GetEnvAsInt("DB_CONN_MAX_LIFETIME_SEC", 60)),
		},
		Echo: EchoServer{
			ListenAddress: util.GetEnv("SERVER_ECHO_LISTEN_ADDRESS", ":8050"),
		},
	}
}
