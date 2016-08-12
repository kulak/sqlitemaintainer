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
	var db_file string
	var sql_dir string
	var help bool
	flag.StringVar(&db_file, "db", "sqlite.db",
		"Name of SQLite database file to create or upgrade.")
	flag.StringVar(&sql_dir, "dir", "sql",
		"Name of the directory with SQL scripts.")
	flag.BoolVar(&help, "h", false, "Prints usage.")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	fmt.Printf("\tDatabase file: %s\n\tDirectory: %s\n is about to be upgraded.\n",
		db_file, sql_dir)
	fmt.Println("Press ENTER to proceed or terminate the application to quit.")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	version, err := sqlitemaint.UpgradeSQLite(db_file, sql_dir)
	if err != nil {
		log.Fatalln("Failed to upgrade DB.  Error: %v.", err)
	}
	fmt.Printf("Current version is %v.\n", version)
}
