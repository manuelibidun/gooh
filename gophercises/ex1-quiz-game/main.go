package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

type QuizSet []Quiz

var filename string
var timeLimit int
var shuffleState bool

func init() {
	const (
		defaultInputFile    = "problems.csv"
		defaultTimeLimit    = 30
		defaultShuffleState = true
	)

	flag.StringVar(&filename,
		"filename",
		defaultInputFile,
		"Input file",
	)
	flag.StringVar(&filename,
		"f",
		defaultInputFile,
		"Input file",
	)
	flag.IntVar(&timeLimit,
		"timeLimit",
		defaultTimeLimit,
		"Time limit",
	)
	flag.IntVar(&timeLimit,
		"t",
		defaultTimeLimit,
		"Time limit",
	)
	flag.BoolVar(&shuffleState,
		"shuffle",
		defaultShuffleState,
		"Shuffle question (T/F)?",
	)
	flag.BoolVar(&shuffleState,
		"s",
		defaultShuffleState,
		"Shuffle question (T/F)?",
	)
}
func main() {
	flag.Parse()
	quizset := loadQuizFromFile(filename, shuffleState)
	var response string
	var correct int
	for _, q := range quizset {
		fmt.Printf("Solve: `%s ", q.Question+"` (or `q` to quit) ?")
		fmt.Scanf("%s", &response)
		if response == "q" {
			break
		}
		if q.Answer == response {
			correct += 1
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(quizset))
}

func loadQuizFromFile(f string, s bool) QuizSet {
	handle, err := os.Open(f)
	if err != nil {
		log.Fatal("Error opening file. Error code:", err)
		os.Exit(1)
	}

	defer handle.Close()

	cr := csv.NewReader(handle)
	data, err := cr.ReadAll()
	if err != nil {
		log.Fatal("Error reading file. Error code:", err)
		os.Exit(1)
	}
	qzset := getQuizSet(data)
	if shuffleState {
		qzset = shuffleQuizSet(qzset)
	}
	return qzset
}

func getQuizSet(d [][]string) QuizSet {
	var qzset QuizSet
	for _, row := range d {
		quiz := Quiz{
			Question: row[0],
			Answer:   row[1],
		}
		qzset = append(qzset, quiz)
	}
	return qzset
}

func shuffleQuizSet(quizset QuizSet) QuizSet {
	rand.Seed(time.Now().UnixNano())
	for i := range quizset {
		k := rand.Intn(len(quizset))
		quizset[i], quizset[k] = quizset[k], quizset[i]
	}
	return quizset
}
