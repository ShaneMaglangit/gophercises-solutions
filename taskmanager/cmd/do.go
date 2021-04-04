package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do [order]",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		rmId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Invalid parameter")
		}

		db, err := bolt.Open("todo.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TODO"))
			index := 1
			err = b.ForEach(func(k, v []byte) error {
				if index == rmId {
					err = b.Delete(k)
					fmt.Printf("You have completed %s\n", string(v))
				}
				index++
				return err
			})
			return err
		})
	},
}
