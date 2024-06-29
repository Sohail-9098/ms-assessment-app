package question

var QuestionSet = []MultipleChoiceQuestion{
	{
		QuestionText:  "What is the capital of France?",
		Options:       [4]string{"Berlin", "Madrid", "Paris", "Lisbon"},
		Mark:          5,
		correctAnswer: "Paris",
	},
	{
		QuestionText:  "Which planet is known as the Red Planet?",
		Options:       [4]string{"Earth", "Mars", "Jupiter", "Saturn"},
		Mark:          5,
		correctAnswer: "Mars",
	},
	{
		QuestionText:  "Who wrote 'To Kill a Mockingbird'?",
		Options:       [4]string{"Harper Lee", "Mark Twain", "Ernest Hemingway", "F. Scott Fitzgerald"},
		Mark:          5,
		correctAnswer: "Harper Lee",
	},
	{
		QuestionText:  "What is the largest ocean on Earth?",
		Options:       [4]string{"Atlantic Ocean", "Indian Ocean", "Arctic Ocean", "Pacific Ocean"},
		Mark:          5,
		correctAnswer: "Pacific Ocean",
	},
	{
		QuestionText:  "Who painted the Mona Lisa?",
		Options:       [4]string{"Leonardo da Vinci", "Vincent van Gogh", "Pablo Picasso", "Claude Monet"},
		Mark:          5,
		correctAnswer: "Leonardo da Vinci",
	},
}
