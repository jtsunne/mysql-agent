package cmd

import (
	"fmt"
	"mysql-agent/config"
	"mysql-agent/database"
	"mysql-agent/logs"
	"net/http"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "startserver",
	Short: "Starts the HTTP server to show MySQL Slave Status",
	Run: func(cmd *cobra.Command, args []string) {
		startHTTPServer()
	},
}

func startHTTPServer() {
	config := config.LoadConfig()
	logger := logs.NewLogger(config)
	db := database.Connect(config, logger)

	http.HandleFunc("/mysql_status", func(w http.ResponseWriter, r *http.Request) {
		status, err := database.GetSlaveStatus(db)
		if err != nil {
			logger.Errorf("Error getting slave status: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(status.ToJSON()))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "/ Not Found. Use /mysql_status", http.StatusNotFound)
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
