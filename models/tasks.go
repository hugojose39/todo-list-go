package models

import (
	"time"
	"todo-list/db"
)

type Task struct {
	Id          int
	Description string
	Date        string
	Done        bool
}

func parseDate(dateString string) time.Time {
    layout := "2006-01-02T15:04:05Z"
    parsedDate, err := time.Parse(layout, dateString)

    if err != nil {
        panic(err.Error())
    }

    return parsedDate
}

func formatDate(date time.Time) string {
    return date.Format("02/01/2006")
}

func GetAllTasks() []Task {
	db := db.ConectWithDataBase()

	selectAllTasks, err := db.Query("select * from tasks order by id asc")

	if err != nil {
		panic(err.Error())
	}

	t := Task{}
	tasks := []Task{}

	for selectAllTasks.Next() {
		var id int
		var description, date string
		var done bool

		err = selectAllTasks.Scan(&id, &description, &date, &done)

		if err != nil {
			panic(err.Error())
		}

		t.Id = id
		t.Description = description
		t.Date = formatDate(parseDate(date))
		t.Done = done

		tasks = append(tasks, t)
	}

	defer db.Close()
	return tasks
}

func CreateTask(description string, date string, done bool) {
	db := db.ConectWithDataBase()

	insertDataBase, err := db.Prepare(
		"insert into tasks (description, date, done) values ($1, $2, $3)",
	)

	if err != nil {
		panic(err.Error())
	}

	insertDataBase.Exec(description, date, done)

	defer db.Close()
}

func DeleteTask(id string) {
	db := db.ConectWithDataBase()

	deleteTask, err := db.Prepare("delete from tasks where id = $1")

	if err != nil {
		panic(err.Error())
	}

	deleteTask.Exec(id)

	defer db.Close()
}

func EditTask(id string) Task {
	db := db.ConectWithDataBase()

	task, err := db.Query("select * from tasks where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	taskEdit := Task{}

	for task.Next() {
		var id int
		var description, date string
		var done bool

		err = task.Scan(&id, &description, &date, &done)

		if err != nil {
			panic(err.Error())
		}

		taskEdit.Id = id
		taskEdit.Description = description
		taskEdit.Date = date
		taskEdit.Done = done
	}

	defer db.Close()
	return taskEdit
}

func UpdateTask(id int, description string, date string, done bool) {
	db := db.ConectWithDataBase()

	updateTask, err := db.Prepare(
		"update tasks set description=$1, date=$2, done=$3",
	)

	if err != nil {
		panic(err.Error())
	}

	updateTask.Exec(description, date, done)
	defer db.Close()
}
