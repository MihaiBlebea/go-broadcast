---
Title: Build a Go Quiz Game with concurrent jobs from scratch in 6 easy steps
Summary: Coding a game that involves concurrency must be sooooo complicated... right? Not really! It's quite easy by using Go routines system. Let me show you how to do it.
Image: /static/img/quiz_game.jpg
Tags:
    - game
    - go
    - goroutine
Slug: go-concurrent-quiz-game
Published: "2020-09-26 14:47:00"
---
Coding a game is realy complicated...

This is what I used to tell myself when I just started to code. 

As time went by, I developed the habbit of always trying to code the same project when starting to learn a new language.

You might have already notice that every coding course starts with a "Hello world" or "To do list" demo app.

I have chosen the Quiz game as it gives me the chance to compare the same features across different languages.

For example, implementing a timeout timer in PHP might be more difficult then doing it in Go. But how much more difficult?

Let's give it a go... 😃

We will be starting this from an empty folder and we'll use a couple of waypoints to guilde us through the process:

1. Read questions and answers from a yaml file

2. Parse the yaml file to Go data structs

3. Build the game loop with inputs from the user and outputs to the console

4. Add a timer that will tell the user when the game time has expired

5. Move everything from the main file to separate packages and have as little as possible logic inside the entry file

6. Add unit tests and prepare for production

As a bonus, we will try to code everything in under 100 lines. This is just because I saw a lot of comments on different forums saying that Go is too verbouse.

Let's put that to the challenge.

I will build this in incremental steps, giving you the ability to check how it works on every step of the way.

We will start by creating a folder `go-quiz` with a `main.go` file inside it:

```go
package main

func main() {
    // Main logic goes here ...
}
```
Next, add this yaml file in the same folder:
```yaml
problems:
    - question: "What is the capital of UK"
      answer: "London"
    - question: "10 + 5"
      answer: "15"
    - question: "How many wheels does a motorcycle have"
      answer: "2"
```
You can add as many "problems" as you like. I started with 3 just because it's easier to test.

To finish the setup, just initialize the go module and install one single dependency package, by runing this commands:

- `go mod init github.com/<your-github-username>/go-quiz`

- `go get gopkg.in/yaml.v2`

Now setup is finished, let's start writing our first code.

## Read questions and answers from a yaml file

The first iteration of the main function should look like this:

**main.go:**

```go
package main

import (
	"flag"
	"fmt"
    "log"
    "ioutil"
)

func main() {
	// Get the name of the questions file and the timeout limit in sec
	file := flag.String("file", "questions.yaml", "File with questions")
	limit := flag.Int("limit", 10, "Timeout limit in sec")
	flag.Parse()

    // Read the external file.
    // Noticed that we de-referenced the file variable before using it
    // file is just the pointer to the string, will returns a memory address
    // Returns b which is a slice of bytes
	b, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
    }
}
```
It's time to have a look over what we just created.

On lines 12 and 13 we are defining two flags that will capture arguments passed when we run this application. The two resulting variable will be pointer to string and pointer to integer, so if you print them out, you will get the addresses in memory of the variables that those two point to.

Not very usefull. 

To get the actual values, you will have to dereference those pointers first.

You can use get the variable's value with `*file` and `*limit`.

Next, we call `flag.Parse()` to parse those two flags. Don't foget to call that method after defining the flags.

On line 20, we will use the value of file (path to file) to get the content of our *questions.yaml* file.

We check for errors and handle them if any, and that's it for now.

## Parse the yaml file to Go data structs

**main.go:**

```go
package main

import (
	"flag"
	"fmt"
    "log"
    "yaml"
)

// Quiz type to hold the list of problems
type Quiz struct {
	Problems []Problem `yaml:"problems"`
}

// Problem type with question and answer
type Problem struct {
	Question string `yaml:"question"`
	Answer   string `yaml:"answer"`
}

func main() {
    // Parse the flags here 
    // ...

	b, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
    }

    // Decode the yaml file using the yaml package that we required as dependency
    var quiz Quiz
	err = yaml.Unmarshal(b, &quiz)
	if err != nil {
		log.Fatal(err)
    }
    
    fmt.Println(quiz)

    // Game loop should go here
    // ...
}
```
Line 11 to line 19, we define two new data structs.

In Go, data is separated from the behaviour. If you look at languages like PHP, Node and Python, you notice that classes tend to contain data (attributes) and logic (methods), but not so much in Golang.

I rejected this system when I first looked at this language, but after a while I started to see the benefits of this approach.

It's much cleaner and from a coding perspective, it kind of makes sense.

Of course, you can still attach methods to this structs, the struct becoming a "receiver" of the method, but we will talk about this a bit later.

Another section that we added to the main.go file starts at line 31.

We created an empty Quiz struct and passed it by reference into the `yaml.Unmarshal` method. This takes care of parsing the raw data that comes out as a slice of bytes from `ioutil.ReadFile()` and "populates" our struct with it.

You will encounter this pattern a lot in Golang.

If you run this, as it stands, you will get a log of the Quiz struct with the data from the yaml file.

In the next iteration, we will start building the game loop witch will contain all our logic.

To make this more readable, I will replace the sections that we already built with placeholder comments to give you a better understanding of the big picture, but in the same time, keep this file short (and sweet 😃).

## Build the game loop and get user inputs

**main.go:**

```go
package main

import (
    // ...
)

func main() {

    // ...
    // Rest of the code reading and parsing the yaml file

    // Create a variable to hold the score
    var score int
	for i, problem := range quiz.Problems {
	    fmt.Printf("Question %d: %s?\n", i+1, problem.Question)

        // Reader reads the user input until he presses Enter (new line)
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        // Remove any whitespace or new lines from the input string
        input = strings.TrimSpace(input)

        // Checks the user input against the correct answer
        if answer.Input == problem.Answer {
            fmt.Println("Correct")
            score++ // Increments the score if the answer is correct
        } else {
            fmt.Println("Wrong")
        }
    }
    
    // Prints a message with the user score from the total possible
    fmt.Printf("Game over! Score is %d of %d\n", score, len(quiz.Problems))
}
```
We are almost there...

Let's take a look over what we added to the file.

First we defined a variable `score` of type integer. Because we didn't assign any value to it at this point, it will take the default integer value, which is 0.

Then, starting from line 14, we start looping over the questions in our quiz.

First thing first, we will print a message to the user asking him the question.

At line 18, we create a new Reader that get's the user's input from the console. It reads everything until it hits the "new line" character.

If you compare the user answer at this point with the correct one from the quiz struct, you will always get a "Wrong" result. This is because the reader also includes the "new line" sign into the input.

So at line 25, we trim those white spaces and new lines out of the user input.

Next, we campare the answer with the correct one and print a message to the player. if correct, we also increment the score.

Easy-peasy! 😃

If you build this now, you can actually play a basic version of our quiz game. 

But we are not done with it...

It's time to add the timeout feature. 😃

## Timeout feature and stoping the game

But first, let's explore how the Go routines work.

> Goroutines are just light weight threads.

You can use the go routines to defer resource intensive tasks or just tasks that you want to run in the background.

We will use them to run a background worker that will notify us when the timeout has expired for our game.

Go routines are very helpfull but not on their own.

We need a mechanism to communicate between our go routines (or threads).

In Go, this mechanism is called channel or go-channel.

We will add a new timer from the `time` standard package. This timer will use a separate go-routine to count down and notify us when the time has expired via a go-channel.

<img src="/static/svg/go_routines.svg" />

Let's implment this in our code.

To make this easier to read, I will replace the rest of the code that we built until now with placeholder comments so you can se where everything should go (no pun intended) 😃.

**main.go:**

```go
package main

import (
    // ...
    "time"
)

func main() {
    // ...
    // Logic to read the yaml file goes here

    // Create a new timer
    t := time.NewTimer(
        time.Duration(*limit)) * time.Second // Transform the limit pointer to seconds
    )

    // Game loop goes here
    var score int
	for i, problem := range quiz.Problems {
	    fmt.Printf("Question %d: %s?\n", i+1, problem.Question)

        select {
		case <-t.C:
			goto GameOver // Go to this tag if time expires
		default:
            // Reader reads the user input until he presses Enter (new line)
            reader := bufio.NewReader(os.Stdin)
            input, err := reader.ReadString('\n')
            if err != nil {
                log.Fatal(err)
            }

            // Remove any whitespace or new lines from the input string
            input = strings.TrimSpace(input)

			if input == problem.Answer {
				fmt.Println("Correct")
				score++
			} else {
				fmt.Println("Wrong")
			}
		}
    }

// Add tag to jump here if the timer has expired
GameOver:
    // Prints a message with the user score from the total possible
    fmt.Printf("Game over! Score is %d of %d\n", score, len(quiz.Problems))
}
```
At line 13, we created a timer using the `time` package and set it up to expire after the `limit` variable that we get from the user at the start of the game.

Next, at line 22 we added a select keywowrd to "listen" for event coming from the timer's `C` channel. The arrow `<-` shows the direction of the data, flowing from the timer go-routine to the main go-routine.

If the time expires, then we will "jump" to the GameOver tag, using the `goto` keyword.

This action, will skip the loop and display the score to the user.

If you run everything now with `go run . -limit=5` and wait for 5 seconds you will notice that the timer only seems to expires when you press enter.

This is happening because the block of code that reads the user input is actually blocking the main thread.

For this game to work better, we have to defer the user input to a different go-routine (thread).

We will build a go-channel to communicate with this new go-routine and get the input back and handle any errors.

> Try to handle the errors in the main thread as much as possible

Our final version of the main file will look like this:

**main.go:**

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
    "time"
    
    "gopkg.in/yaml.v2"
)

// Quiz type
type Quiz struct {
	Problems []Problem `yaml:"problems"`
}

// Problem type
type Problem struct {
	Question string `yaml:"question"`
	Answer   string `yaml:"answer"`
}

func main() {
	file := flag.String("file", "questions.yaml", "File for questions")
	limit := flag.Int("limit", 10, "Timeout limit in sec")
	flag.Parse()

	var quiz Quiz
	err = yaml.Unmarshal(b, &quiz)
	if err != nil {
		log.Fatal(err)
    }

	t := time.NewTimer(time.Duration(*limit) * time.Second)

	var score int
	for i, problem := range quiz.Problems {
        fmt.Printf("Question %d: %s?\n", i+1, problem.Question)
        
        // Create a type that accepts the input as string
        // And an error to be sent to the main thread
		type Answer struct {
			Input string
			Err   error
        }
        
        // Create a channel of type Answer
        answerChan := make(chan Answer)
        
        // Our go routine that will handle the user input
		go func() {
            // Move the reader inside the go routine
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
                // If error is not nil
                // Then create an Answer and pass it to the channel
                // Notice the direction of the arrow flowing towards the channel
				answerChan <- Answer{
					Err: err,
				}
            }
            
            // Also send the user input through the channel
			input = strings.TrimSpace(input)
			answerChan <- Answer{
				Input: input,
			}
		}()

		select {
		case <-t.C:
            goto GameOver
        // Change the default case with this case that accepts the content of the channel
        // This will get triggered only if a message is coming through the channel from the go routine
		case answer := <-answerChan:
            // handle the error if any from the channel
			if answer.Err != nil {
				log.Fatal(answer.Err)
            }
            
            // Same as before, handle the user input validation here
			if answer.Input == problem.Answer {
				fmt.Println("Correct")
				score++
			} else {
				fmt.Println("Wrong")
			}
		}
	}

GameOver:
	fmt.Printf("Game over! Score is %d of %d\n", score, len(quiz.Problems))
}
```
At line 45 we define a type that we will use to specify what data type can be passed through our channel.

The `Answer` type contains the user input as string and an error.

Line 51, we created a channel of type `Answer`

Then we used that channel in a new fancy go-groutine that we built from scratch as a closure.

Notice how at line 62 and 69 we passed the data through the channel.

Finally, at line 79, we have a new select case that accept and reads the data coming from the channel.

We use that data to handle the errors and read the user input.

This will be the final version of our go quiz game.

If you run this again with `go run .` you will notice that if you don't answer the questions in the default limit of 10 sec, then you will be kicked out of the game and get the final score.

This is the behaviour that we were expected to have.

## Structure application into packages

We want to create 3 separate packages:

- game - this will hold the the game loop and most of the game logic

- player - this will be a service to abstract the input and output to and from the player

- quiz - this will be a package that will hold all logic to read the yaml file and parse to structs

Our folder structure is going to look like this:

```
|-- go-quiz //root folder
    |-- game
        |-- logic.ga
        |-- service.go
        |-- logic_test.go
    |-- player
        |-- human.go
        |-- computer.go
    |-- quiz
        |-- logic.go
        |-- service.go
    |-- questions.yam
    |-- main.go
    |-- go.sum
    |-- go.mod
```

### Quiz package

The main responsability of the **quiz package** will be to fetch the data from the source (yaml file) and parse it to golang structures.

We will move the Quiz and Problem structs inside this package and expose them to be used by the other packages.

The New method will accept a file path as parameter and return a pointer to a Quiz struct and an error.

**quiz/model.go**

```go
package quiz

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Quiz _
type Quiz struct {
	Problems []Problem `yaml:"problems"`
}

// Problem _
type Problem struct {
	Question string `yaml:"question"`
	Answer   string `yaml:"answer"`
}

// New returns a new Quiz
func New(fileName string) (*Quiz, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return &Quiz{}, err
	}

	var quiz Quiz
	err = yaml.Unmarshal(b, &quiz)
	if err != nil {
		return &Quiz{}, err
	}

	return &quiz, nil
}
```

### Game package

We will move all the logic of the game, including the game loop, inside this package.

A service struct will contain the timer limit, the quiz struct and the player.

From the perspective of this package, the Player is just an interface that require two methods to be implemented:

- Print() to display a message to to player

- Input() to get information from the player

The Run() method on the service struct will run the game loop. Notice that we replaced the `fmt.Println()` with `service.player.Print()`.

Also, we replaced the logic that reads user input with `service.player.Input()`.

We did this, because we want to abstract the player of the game, and to allow different types of player to play. For example, a ComputerPlayer can implement the Player interface and play the game.

I told you that computers are going to take over the world... but maybe not yet.

**game/logic.go**

```go
package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/MihaiBlebea/go-quiz/quiz"
)

// Player interface
type Player interface {
	Print(string)
	Input() (string, error)
}

type service struct {
	limit  int
	quiz   quiz.Quiz
	player Player
}

// New _
func New(limit int, quiz quiz.Quiz, player Player) Service {
	return &service{limit, quiz, player}
}

func (s *service) Run() (int, error) {
	t := time.NewTimer(time.Duration(s.limit) * time.Second)

	var score int
	for i, prob := range s.quiz.Problems {
		s.player.Print(fmt.Sprintf("Problem %d: %s ?", i+1, prob.Question))

		type Answer struct {
			input string
			err   error
		}

		answerChan := make(chan Answer)

		go func() {
			input, err := s.player.Input()
			if err != nil {
				answerChan <- Answer{
					err: err,
				}
			}

			input = strings.TrimSpace(input)

			answerChan <- Answer{
				input: input,
			}
		}()

		select {
		case <-t.C:
			goto GameOver
		case answer := <-answerChan:
			if answer.err != nil {
				return score, answer.err
			}

			if answer.input == prob.Answer {
				s.player.Print("Correct\n")
				score++
			} else {
				s.player.Print("Wrong\n")
			}
		}
	}

GameOver:
	s.player.Print(fmt.Sprintf("Game over! Your score is %d from %d", score, len(s.quiz.Problems)))

	return score, nil
}
```

We also added a Service interface to abstract the implementation of the game.

**game/service.go**

```go
package game

// Service interface
type Service interface {
	Run() (int, error)
}
```

### Player package

The player package defines the two types of players that can play this game, Human and Computer.

Notice that both of these structs implement the interface Player from the game package so they can be easily passed as a player when creating a new game.

**player/human.go:**

```go
package player

import (
	"bufio"
	"fmt"
	"os"
)

// Human player
type Human struct {
}

// Print _
func (h *Human) Print(output string) {
	fmt.Println(output)
}

// Input _
func (h *Human) Input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
```

**player/computer.go**

```go
package player

import (
	"fmt"
)

// Computer player
type Computer struct {
	Answers []string
	counter int
}

// Print _
func (c *Computer) Print(output string) {
	fmt.Println(output)
}

// Input _
func (c *Computer) Input() (string, error) {
	defer c.incrementCounter()

	return c.Answers[c.counter], nil
}

func (c *Computer) incrementCounter() {
	c.counter++
}
```

Now that we structed our application into packages, we have a much more flexible solution.

We could easily replace the players of the game, maybe extract the quiz from a json file or even change how the game loop works without affecting the other parts of the application.

Also, now that we separated the logic based on responsability, we could add tests.

We will talk about Golang test in another article...

This is awesome! We have 3 new packages that we can push to github and require them  as dependencies to the main app if we really wanted to.

We are missing just one single file...

The main.go file.

As you noticed, we abstracted as much logic as possible and now, our main file should be just a couple of lines of code that instantiate some services and pass them around.

**main.go:**

```go
package main

import (
	"flag"
	"log"

	"github.com/MihaiBlebea/go-quiz/game"
	"github.com/MihaiBlebea/go-quiz/player"
	"github.com/MihaiBlebea/go-quiz/quiz"
)

func main() {
	fileName := flag.String("file", "problems.yaml", "The name of the file with the problems")
	limit := flag.Int("limit", 10, "The time limit for the quiz in seconds")
	flag.Parse()

	quiz, err := quiz.New(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	player := player.Human{}
	// player := player.Computer{Answers: []string{"London", "15"}}

	gameService := game.New(*limit, *quiz, &player)
	_, err = gameService.Run()
	if err != nil {
		log.Fatal(err)
	}
}
```

Notice how we are left with just 30 lines of code in our main file, and I am sure we could go even lower then that...

If you want to see the complete code, please check the <a href="https://github.com/MihaiBlebea/go-quiz" taget="_blank">Github repo link here</a>.

Let me know if you have any suggestions in the comments section below.