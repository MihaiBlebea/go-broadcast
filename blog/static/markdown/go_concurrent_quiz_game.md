---
Title: Build a Go Quiz Game with concurrent jobs from scratch in 6 easy steps
Summary: Coding a game that involves concurrency must be sooooo complicated... right? Not really! It's quite easy by using Go routines system. Let me show you how to do it.
Image: /static/img/quiz_game.jpg
Tags:
    - game
    - go
    - goroutine
Slug: go-concurrent-quiz-game
# Published: "2020-08-27 10:04:05"
---
Coding a game must be sooooo complicated...

This is what I used to tell myself when just started to code. 

As time went by, I developed this habbit of always trying to code the same project when wanting to learn a new language.

You might have already notice that every coding course starts with a "Hello world" or "To do list" app.

I have chosen the Quiz game as it gives me the chance to compare more features across different languages.

For example, implementing a timeout timer in PHP might be more difficult then doing it in Go. But how much more difficult?

Let's give it a run...

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

So let's start by creating a folder `go-quiz` with a `main.go` file inside it:

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

The first iteration of the main function should look like this:

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

	"github.com/MihaiBlebea/go-test/quiz"
)

func main() {
	// Get the name of the questions file and the timeout limit in sec
	file := flag.String("file", "questions.yaml", "File with questions")
	limit := flag.Int("limit", 10, "Timeout limit in sec")
	flag.Parse()

    // Read the external file.
    // Noticed that we de-referenced the file variable before using it
    // file is just the pointer to the string, will returns a memory address
    // Use *file
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
If you run this, as it stands, you will get a log of the Quiz struct with the data from the yaml file.

In the next iteration, we will start building the game loop witch will contain all our logic.

To make this more readable, I have added a placeholder where the next code should go into in the main file.

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

If you build this now, you can actually play a basic version of our quiz game. But we are not done with it.

It's time to add the timeout feature. 😃

But first, let's explore how the Go routines work.

> Goroutines are just light weight threads.

You can use the go routines to defer resource intensive tasks or just tasks that you want to run in the background.

We will use them to run a background worker that will notify us when the timeout has expired for our game.

Go routines are very helpfull but not on their own.

We need a mechanism to communicate between our go routines (or threads).

In Go, this mechanism is called channel or go-channel.

We will add a new timer from the `time` standard package. This timer will use a separate go-routine to count down and notify us when the time has expired via a go-channel.

Let's implment this in our code.

To make this easier to read, I will replace the rest of the code that we built until now with placeholder comments so you can se where everything should go (no pun intended) 😃.

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
You notice that we used a select to "listen" for event coming from the timer `C` channel. The arrow `<-` shows the direction of the data, flowing from the timer go-routine to the main go-routine.

If the time expires, then we will break the select and the for loop and `goto` the GameOver tag at the bottom of the file.

This will display the score to the user.

If you run everything now with `go run . -limit=5` and wait for 5 seconds you will notice that the time only expires when you press enter.

This is happening because the block of code that reads the user input is actually blocking the main thread.

For this game to work better, we have to defer the user input to a different go-routine (thread).

We will build a go-channel to communicate with this new go-routine and get the input back and handle any errors.

> Try to handle the errors in the main thread as much as possible

Our final version of the main file will look like this:

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

	"github.com/MihaiBlebea/go-test/quiz"
)

func main() {
	file := flag.String("file", "questions.yaml", "File for questions")
	limit := flag.Int("limit", 10, "Timeout limit in sec")
	flag.Parse()

	fmt.Println(*file, *limit)

	q := quiz.Service{File: *file}
	quiz, err := q.Parse()
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
This will be the final version of our go quiz game.

If you run this again with `go run .` you will notice that if you don't answer the questions in the default limit of 10 sec, then you will be kicked out of the game and get the final score.

This is the behaviour that we were expected to have.



