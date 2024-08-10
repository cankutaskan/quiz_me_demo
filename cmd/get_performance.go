package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var participantID string

// getPerformanceCmd represents the getPerformance command
var getPerformanceCmd = &cobra.Command{
	Use:   "getPerformance",
	Short: "Fetch the performance of a participant",
	Run: func(cmd *cobra.Command, args []string) {
		getPerformance()
	},
}

func init() {
	rootCmd.AddCommand(getPerformanceCmd)

	// Add a flag to specify the participant ID
	getPerformanceCmd.Flags().StringVarP(&participantID, "participant", "p", "", "Participant ID")
	getPerformanceCmd.MarkFlagRequired("participant")
}

func getPerformance() {
	if participantID == "" {
		log.Fatalf("Participant ID is required")
	}

	// Build the request URL
	url := fmt.Sprintf("http://localhost:8080/api/quiz-me/performance/%s", participantID)

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch performance: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Pretty print the JSON response
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, body, "", "  "); err != nil {
		log.Fatalf("Failed to format JSON response: %v", err)
	}

	fmt.Println("Performance Result:")
	fmt.Println(prettyJSON.String())
}
