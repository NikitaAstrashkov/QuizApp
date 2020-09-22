package main

import (
	"fmt"
)

type QuizPoll struct {
	Question      string
	a             string
	b             string
	c             string
	CorrectAnswer string
}

var QuizData = []QuizPoll{
	{
		"Who wasn't the president of Ukraine?",
		"Kuchma",
		"Ahmetov",
		"Both were",
		"b",
	},
	{
		"Select one who were \"Hetman of the Troops of Zaporizhzhya\" of Ukrainian cossacs",
		"Samoylo Koshka",
		"Somko Yakim",
		"Mogyla Andrey",
		"a",
	},
	{
		"Which city was the capital of Ukraine from 1919 to 1934",
		"Kyiv",
		"Lviv",
		"Kharkiv",
		"c",
	},
}

func main() {
	var (
		CorrectAnswers = 0
		Questions      = 0
	)
	fmt.Println("Welcome to the quiz!")
	for _, Poll := range QuizData {
		Questions++
		var answer string
		fmt.Println(Poll.Question)
		fmt.Printf("a. %s\n", Poll.a)
		fmt.Printf("b. %s\n", Poll.b)
		fmt.Printf("c. %s\n", Poll.c)
		_, _ = fmt.Scan(&answer)
		if answer == Poll.CorrectAnswer {
			fmt.Printf("You're right! Answer %s is correct\n", answer)
			CorrectAnswers++
		} else {
			fmt.Printf("You're wrong. Correct answer is %s\n", Poll.CorrectAnswer)
		}
	}
	fmt.Printf("The quiz is over, your result: %d answers out of %d are correct.", CorrectAnswers, Questions)
}
