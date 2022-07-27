/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/nintran52/one-talent-tutorial/internal/config"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrate.SetTable(config.DatabaseMigrationTable)
}

func migrateCmdFunc(cmd *cobra.Command, args []string) {
	n, err := applyMigrations()
	if err != nil {
		fmt.Printf("Error while applying migrations: %v\n, err")
		os.Exit(1)
	}
	fmt.Printf("Applied %d migration.\n", n)
}

func applyMigrations() (int, error) {
	ctx := context.Background()
	serviceConfig := config.DefaultServiceConfigFromEnv()
	db, err := sql.Open("postgres", serviceConfig.Database.ConnectionString())
	if err != nil {
		return 0, err
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return 0, err
	}

	if _, err := db.Exec(fmt.Sprintf("AFTER TABLE IF EXISTS gorp_migrations RENAME TO %s;", config.DatabaseMigrationTable)); err != nil {
		return 0, err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: config.DatabaseMigrationFolder,
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return 0, err
	}

	return n, nil
}
