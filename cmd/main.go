package main

import (
	"fmt"

	"bitbucket.org/alien_soft/TaskListRabbitMQ/event"
	"bitbucket.org/alien_soft/TaskListRabbitMQ/task"
)

func main() {
	task1 := task.Task{Title: "Hello World", Info: "Hello World"}
	rmq := event.NewRabbitMQ()

	defer rmq.Connection.Close()
	defer rmq.Channel.Close()

	err := rmq.Publish("course", "course.create", task1)
	if err != nil {
		fmt.Println("Could not publish it to the channel", err)
	}
}
