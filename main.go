package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Id     int
	Title  string
	IsDone bool
	Description string
}

var tasks []Task
var taskcounter int

func main() {
	reader := bufio.NewReader(os.Stdin)

	for	{
		fmt.Println("\nTo-Do List App")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		choice, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			AddTask(reader)
		case "2":
			ViewTasks()	
		case "3":
			MarkAsDone(reader)
		case "4":
			DeleteTasks(reader)
		case "5":
			fmt.Println("GoodBye 🥹✨")
		default:
			fmt.Println("Invalid choice. Please try again.")
		}


		
	}

}

func AddTask(reader *bufio.Reader) {
	fmt.Printf("<<<<<<-------Enter Title---->>>>>>>\n")
	title, _ := reader.ReadString('\n')	
	title = strings.TrimSpace(title)
	fmt.Printf("<<<<<<-------Enter Descrption---->>>>>>>\n")
	desc, _ := reader.ReadString('\n')	
	desc = strings.TrimSpace(desc)

	taskcounter++
	task := Task{
		Id: taskcounter,
		Title: title,
		IsDone: false,
		Description: desc,
	}

	tasks = append(tasks, task)
	fmt.Printf(">>>>>>>>>>>>>>Task Added Sucessfully<<<<<<<<<<<<\n")
	fmt.Println(task)
	fmt.Printf("--------------------------------------------------")
}

func ViewTasks()  {

	if len(tasks) <= 0 {
		return
	}
	fmt.Println("--------------Tasks--------------")
	for _,task := range tasks {
		fmt.Printf("Title is %v\n IsDone id %v", task.Id, task.IsDone )
	}
}

func MarkAsDone(reader *bufio.Reader) {
	ViewTasks()
	fmt.Println("Which One is done --->>>> ")

	taskId, _ := reader.ReadString('\n')
	taskId = strings.TrimSpace(taskId)

	taskid,_ := strconv.Atoi(taskId)

	for idx,task := range tasks {
		if task.Id == taskid {
			tasks[idx].IsDone = true
			return
		}
	}

	fmt.Println("-------------Task Marked As Done-------------")
}

func DeleteTasks(reader *bufio.Reader){
	ViewTasks()
	fmt.Println("Which One to Delete ->>>")
	taskID, _ := reader.ReadString('\n')
	taskID = strings.TrimSpace(taskID)

	taskid, _ := strconv.Atoi(taskID)
	
	for idx,task := range tasks {
		if task.Id == taskid {
			tasks = append(tasks[:idx], tasks[idx+1:]...)
		}
	}
}