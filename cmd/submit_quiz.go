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

var userID string
var responses []string // Expecting responses in format "questionID:answerID"

var submitAnswersCmd = &cobra.Command{
	Use:   "submitAnswers",
	Short: "Submit answers to the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		submitAnswers()
	},
}

func init() {
	rootCmd.AddCommand(submitAnswersCmd)

	// Add flags to accept user ID and responses
	submitAnswersCmd.Flags().StringVarP(&userID, "user", "u", "", "User ID")
	submitAnswersCmd.Flags().StringArrayVarP(&responses, "response", "r", []string{}, "Responses in format 'questionID:answerID'")

	submitAnswersCmd.MarkFlagRequired("user")
	submitAnswersCmd.MarkFlagRequired("response")
}

func submitAnswers() {
	if userID == "" || len(responses) == 0 {
		log.Fatalf("User ID and responses are required")
	}

	var parsedResponses []map[string]int
	userResponseMap := make(map[int]int)

	// Parse user responses and store in a map
	for _, response := range responses {
		var qID, aID int
		fmt.Sscanf(response, "%d:%d", &qID, &aID)
		userResponseMap[qID] = aID
		parsedResponses = append(parsedResponses, map[string]int{"question_id": qID, "answer_id": aID})
	}

	// Add missing answers with default value (-1)
	for _, question := range fetchedQuestions {
		if _, answered := userResponseMap[question.ID]; !answered {
			parsedResponses = append(parsedResponses, map[string]int{"question_id": question.ID, "answer_id": -1})
		}
	}

	payload := map[string]interface{}{
		"user_id":   userID,
		"responses": parsedResponses,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Make the POST request
	resp, err := http.Post("http://localhost:8080/api/quiz-me/responses", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Failed to submit answers: %v", err)
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

	fmt.Println("Submission Result:")
	fmt.Println(prettyJSON.String())
}
