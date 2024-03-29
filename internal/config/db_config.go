package config

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/nintran52/one-talent-tutorial/internal/util"
)

const DatabaseMigrationTable = "migrations"

var DatabaseMigrationFolder = filepath.Join(util.GetProjectRootDir(), "/migrations")

type Database struct {
	Host             string
	Port             int
	Username         string
	Password         string `json:"-"`
	Database         string
	AdditionalParams map[string]string `json:",omitempty"`
	MaxOpenConns     int
	MaxIdleConns     int
	ConnMaxLifetime  time.Duration
}

func (c Database) ConnectionString() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.Host, c.Port, c.Username, c.Password, c.Database))

	if _, ok := c.AdditionalParams["sslmode"]; !ok {
		b.WriteString(" sslmode=disable")
	}

	if len(c.AdditionalParams) > 0 {
		params := make([]string, 0, len(c.AdditionalParams))
		for param := range c.AdditionalParams {
			params = append(params, param)
		}

		sort.Strings(params)

		for _, param := range params {
			fmt.Fprintf(&b, " %s=%s", param, c.AdditionalParams[param])
		}
	}

	fmt.Printf("DB uri: %v", b.String())

	return b.String()
}
