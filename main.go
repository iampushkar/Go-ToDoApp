package main

import (
	"fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
	fmt.Println("Welcome to the TODO App in Golang")
  var indexInput int
  var taskInput, newTaskInput string 

  fmt.Println("MENU : ")
  fmt.Println("1. Add Task")
  fmt.Println("2. List Tasks")
  fmt.Println("3. Mark as Complete")
  fmt.Println("4. Update Task")
  fmt.Println("5. Delete Task")
  fmt.Println("6. Exit")

  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("Enter your choice (1, 2, 3, 4, 5, 6) : ")
    scanner.Scan()
    input := scanner.Text()

    choice, err := strconv.Atoi(input)
    if err != nil {
      fmt.Println("Invalid Choice")
      continue 
    }

    switch choice {
      case 1:
      fmt.Print("Enter the task : ")
      scanner.Scan()
      taskInput = scanner.Text()
      addTask(taskInput)
      case 2:
      listTasks()
      case 3:
      fmt.Print("Enter index: ")
      scanner.Scan()
      indexInput, _ = strconv.Atoi(scanner.Text())
      markCompleted(indexInput)
      case 4:
      fmt.Print("Enter index: ")
      scanner.Scan()
      indexInput, _ = strconv.Atoi(scanner.Text())
      fmt.Print("Enter new task: ")
      scanner.Scan()
      newTaskInput = scanner.Text()
      updateTasks(indexInput, newTaskInput)
      case 5:
      fmt.Print("Enter index: ")
      scanner.Scan()
      indexInput, _ = strconv.Atoi(scanner.Text())
      deleteTask(indexInput)
      case 6:
      os.Exit(0)
      default:
      fmt.Println("Invalid Choice")
    }
  }
}

type Task struct {
  TaskName string
  completed bool
}

var tasks []Task 

func addTask(task string) {
  newTask := Task{TaskName: task, completed: false}
  tasks = append(tasks, newTask)
  fmt.Println("Task Added")
}

func listTasks() {
  for i, task := range tasks {
    var status string = "n"
    if task.completed {
      status = "d"
    }

    fmt.Printf("%d. %s [%s] \n", i+1, task.TaskName, status)
  }
}

func markCompleted(index int) {
  if index >=1 && index <= len(tasks) {
    tasks[index - 1].completed = true
    fmt.Println("Tasks marked as completed")
  } else {
    fmt.Println("Invalid Index")
  }
}

func updateTasks(index int, newTaskName string) {
  if index >= 1 && index <= len(tasks) {
    tasks[index - 1].TaskName = newTaskName
    fmt.Println("TaskName updated")
  } else {
    fmt.Println("Invalid Index")
  }
}

func deleteTask(index int) {
  if index >= 1 && index < len(tasks) {
    tasks = append(tasks[:index-1], tasks[index:]...)
    fmt.Println("Task Deleted")
  } else {
    fmt.Println("Invalid Index")
  }
}

