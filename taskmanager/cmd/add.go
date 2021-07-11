package cmd

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Put a new task to TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		db, err := bolt.Open("todo.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer func(db *bolt.DB) {
			_ = db.Close()
		}(db)

		_ = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TODO"))
			id, _ := b.NextSequence()
			err := b.Put(itob(int(id)), []byte(task))
			fmt.Printf("You have added %s\n", task)
			return err
		})
	},
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
