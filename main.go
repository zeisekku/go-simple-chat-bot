package main

import "github.com/zeisekku/go-simple-chat-bot/chatbot"

func main() {
	chatbot.Greet("Jacek", "2023") // change it as you need
	chatbot.ShowName()
	chatbot.GuessAge()
	chatbot.Count()
	chatbot.StartQuiz()
	chatbot.SayGoodbye()
}
