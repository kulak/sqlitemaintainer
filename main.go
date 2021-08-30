package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kulak/sqlitemaint"
)

func main() {
	// copy all args except for program name
	var dbFile string
	var sqlDir string
	var backup bool
	var help bool
	flag.StringVar(&dbFile, "db", "sqlite.db", "Name of SQLite database file to create or upgrade.")
	flag.StringVar(&sqlDir, "dir", "sql", "Name of the directory with SQL scripts.")
	flag.BoolVar(&backup, "backup", true, "Write a backup file with prefix 'Copy-of-'")
	flag.BoolVar(&help, "h", false, "Prints usage.")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	fmt.Printf("\tDatabase file: %s\n\tDirectory: %s\n is about to be upgraded.\n",
		dbFile, sqlDir)
	fmt.Println("Press ENTER to proceed or terminate the application to quit.")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	version, err := sqlitemaint.UpgradeSQLite(dbFile, sqlDir, backup)
	if err != nil {
		log.Fatalln("Failed to upgrade DB. Error:", err)
	}
	fmt.Printf("Current version is %v.\n", version)
}
