package main

import "time"

func QuizRunner(questions, answers []string, answerCh chan string) int {
	var goodAnswers int
	for i := 0; i < len(questions); i++ {
		select {
		case answer := <-answerCh:
			if answer == answers[i] {
				goodAnswers++
			}
		case <-time.After(time.Second):
		}
	}

	return goodAnswers
}
