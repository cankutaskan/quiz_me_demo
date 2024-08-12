# Quiz_me

**Quiz_me** is a simple quiz project that allows users to participate in quizzes, submit their answers, and see how their performance compares to others.

## Project Structure

- **cmd/**: Contains the command-line interface (CLI) commands for interacting with the Quiz_me API.
- **api/**: Contains the API logic and routes.
- **db/**: Contains the in-memory database logic and data models.
- **utils/**: Contains utility functions used across the project.

## Getting Started

### Prerequisites

- Go (Golang) installed on your machine.

### Running the Project

To run the Quiz_me API server:

```bash
go build
go run quiz_me
```

### Example commands
Response structure: {questionID}:{responseID} 
Both questionID and responseID provided in the fetched quiz. 

```bash
go run quiz_me getQuiz --count=3
go run quiz_me submitAnswers --user=user1 --responses=10:2,2:2,8:1
go run quiz_me getPerformance --participant=user1
```

### Notes
Questions are fetched randomly, and any missing responses are automatically populated with -1 as the answer for simplicity. Alternative to each question accomplished by using a field Category. Each category contains alternative questions and response always contain questions with unique IDs.

