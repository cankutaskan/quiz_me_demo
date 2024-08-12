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

var questionCount int
var fetchedQuiz Quiz

type Quiz struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var getQuizCmd = &cobra.Command{
	Use:   "getQuiz",
	Short: "Fetch a quiz from the API",
	Run: func(cmd *cobra.Command, args []string) {
		fetchQuiz()
	},
}

func init() {
	rootCmd.AddCommand(getQuizCmd)

	// Add a flag to specify the number of questions
	getQuizCmd.Flags().IntVarP(&questionCount, "count", "c", 10, "Number of questions to retrieve")
}

func fetchQuiz() {
	url := fmt.Sprintf("http://localhost:8080/api/quiz-me/quiz/%d", questionCount)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch quiz: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the JSON response into the fetchedQuiz object
	if err := json.Unmarshal(body, &fetchedQuiz); err != nil {
		log.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	// Pretty print the JSON response
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, body, "", "  "); err != nil {
		log.Fatalf("Failed to format JSON response: %v", err)
	}

	fmt.Println("Quiz Questions:")
	fmt.Println(prettyJSON.String())
}
