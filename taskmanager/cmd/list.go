package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all task in the list",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("todo.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TODO"))
			err := b.ForEach(func(k, v []byte) error {
				fmt.Println(string(v))
				return nil
			})
			return err
		})
	},
}
