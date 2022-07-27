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
	"github.com/nintran52/one-talent-tutorial/internal/data"
	dbutil "github.com/nintran52/one-talent-tutorial/internal/util/db"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("seed called")
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func seedCmdFunc(cmd *cobra.Command, args []string) {
	if err := applyFixtures(); err != nil {
		fmt.Printf("Error while seeding fixtures: %v", err)
		os.Exit(1)
	}
	fmt.Println("Seeded all fixtures.")
}

func applyFixtures() error {
	ctx := context.Background()
	config := config.DefaultServiceConfigFromEnv()
	db, err := sql.Open("postgres", config.Database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	// insert fixtures in an auto-managed db transaction
	return dbutil.WithTransaction(ctx, db, func(tx boil.ContextExecutor) error {

		fixtures := data.Upserts()

		for _, fixture := range fixtures {
			if err := fixture.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
				fmt.Printf("Failed to upsert fixture: %v\n", err)
				return err
			}
		}

		fmt.Printf("Upserted %d fixtures.\n", len(fixtures))
		return nil

	})
}
