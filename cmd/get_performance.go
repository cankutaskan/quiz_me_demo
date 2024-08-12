package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var participantID string

var getPerformanceCmd = &cobra.Command{
	Use:   "getPerformance",
	Short: "Fetch the performance of a participant",
	Run: func(cmd *cobra.Command, args []string) {
		getPerformance()
	},
}

func init() {
	rootCmd.AddCommand(getPerformanceCmd)
	getPerformanceCmd.Flags().StringVarP(&participantID, "participant", "p", "", "Participant ID")
	getPerformanceCmd.MarkFlagRequired("participant")
}

func getPerformance() {
	if participantID == "" {
		log.Fatalf("Participant ID is required")
	}

	url := fmt.Sprintf("http://localhost:8080/api/quiz-me/quiz/performance/%s", participantID)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch performance: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, body, "", "  "); err != nil {
		log.Fatalf("Failed to format JSON response: %v", err)
	}

	fmt.Println("Performance Result:")
	fmt.Println(prettyJSON.String())
}
