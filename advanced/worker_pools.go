package advanced

import (
	"fmt"
	"time"
)

type TicketRequest struct {
	personID int
	numTickets int
	cost int
}

func processTickets(request <- chan TicketRequest, response chan <- int) {
	for req := range request {
		fmt.Printf("Processing %d tickets of person with ID: %d for a cost of %d\n", req.numTickets, req.personID, req.cost)
		// Simulating some work
		time.Sleep(1 * time.Second)
		response <- req.personID
	}

}

func worker(id int, tasks <-chan int, results chan <- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		// Simulating some work
		time.Sleep(time.Second)
		// Square the task and send it to results channel
		results <- task * task
	}
}

func WorkerPoolDemo() {
	numWorkers := 3
	numJobs := 10
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create workers
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	// Send values to the tasks channel
	for i := range numJobs {
		tasks <- i
	}
	close(tasks)

	// Collect the results
	for range numJobs {
		fmt.Println(<- results)
	}
	close(results)

	// This is an example of using multiple 'workers' to handle a certain amount of tasks
	// Since the workers are perfectly synchronized through the help of channels, this speeds up the processing time
	// We can use multiple synchronized workers to perform a large task in a small time frame

	// Below is an example of a real world use case of worker pools
	// We can process tickets of multiple people simultaneously using multiple workers
	numRequests := 5
	request := make(chan TicketRequest, numRequests)
	// Unbuffered channel will suffice for this use case
	response := make(chan int)

	for range numWorkers {
		go processTickets(request, response)
	}
	
	for i := range numRequests {
		request <- TicketRequest{personID: i + 1, numTickets: (i + 1) * 3 + i, cost: i * i + 3}
	}
	close(request)

	for range numRequests {
		fmt.Printf("Processed ticket of person with ID %d successfully\n", <- response)
	}
	close(response)
}