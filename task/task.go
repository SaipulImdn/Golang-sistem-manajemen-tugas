package task

import "time"

// Struct untuk merepresentasikan tugas
type Task struct {
	Title       string
	Description string
	DueDate     time.Time
	Status      string
}

// Fungsi untuk menambahkan tugas baru ke dalam slice tugas
func AddTask(tasks []Task, title string, description string, dueDate time.Time) []Task {
	newTask := Task{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      "Belum Selesai",
	}
	tasks = append(tasks, newTask)
	return tasks
}

// Fungsi untuk menghapus tugas dari slice tugas berdasarkan indeks
func RemoveTask(tasks []Task, index int) []Task {
	tasks = append(tasks[:index], tasks[index+1:]...)
	return tasks
}

// Fungsi untuk mengupdate status tugas berdasarkan indeks
func UpdateTaskStatus(tasks []Task, index int, status string) []Task {
	tasks[index].Status = status
	return tasks
}
