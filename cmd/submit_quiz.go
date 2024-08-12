package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var userID string
var responses string

var submitAnswersCmd = &cobra.Command{
	Use:   "submitAnswers",
	Short: "Submit answers to the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		submitAnswers()
	},
}

func init() {
	rootCmd.AddCommand(submitAnswersCmd)

	submitAnswersCmd.Flags().StringVarP(&userID, "user", "u", "", "User ID")
	submitAnswersCmd.Flags().StringVarP(&responses, "responses", "r", "", "Comma-separated list of responses in format 'questionID:answerID,questionID:answerID,...'")
	submitAnswersCmd.MarkFlagRequired("user")
	submitAnswersCmd.MarkFlagRequired("responses")
}

func submitAnswers() {
	if userID == "" || responses == "" {
		log.Fatalf("User ID and responses are required")
	}

	responseList := strings.Split(responses, ",")
	parsedResponses := parseUserResponses(responseList)

	payload := map[string]interface{}{
		"user_id":   userID,
		"responses": parsedResponses,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/api/quiz-me/quiz/responses", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Failed to submit answers: %v", err)
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

	fmt.Println("Submission Result:")
	fmt.Println(prettyJSON.String())
}

func parseUserResponses(responseList []string) []map[string]int {
	var parsedResponses []map[string]int

	for _, response := range responseList {
		var qID, aID int
		fmt.Sscanf(response, "%d:%d", &qID, &aID)
		parsedResponses = append(parsedResponses, map[string]int{"question_id": qID, "answer_id": aID})
	}

	return parsedResponses
}
