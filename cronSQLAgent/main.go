package cronsqlagent

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// Create a cron job scheduler
	c := cron.New()

	// Add a cron job that runs the updateQuery function every minute
	c.AddFunc("@every 1m", updateQuery)

	// Start the cron job scheduler
	c.Start()

	// Run indefinitely, or until interrupted
	select {}
}

func updateQuery() {
	// Establish a connection to the PostgreSQL database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Prepare the update query
	query := "UPDATE messages SET is_published = true WHERE scheduled_at >= NOW() AND scheduled_at <= NOW() + INTERVAL '1 minute';"

	// Execute the update query
	_, err = db.Exec(query)
	if err != nil {
		log.Println("Error executing update query:", err)
		return
	}

	log.Println("Update query executed successfully at", time.Now())
}
