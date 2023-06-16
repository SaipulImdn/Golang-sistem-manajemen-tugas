package taskmanager

import (
	"fmt"
	"sort"

	"manajemen-tugas/task"
)

// Fungsi untuk menampilkan daftar tugas
func ShowTasks(tasks []task.Task) {
	if len(tasks) == 0 {
		fmt.Println("Belum ada tugas yang ditambahkan.")
		return
	}

	fmt.Println("Daftar Tugas:")
	fmt.Println("-------------------------------------------")
	for i, t := range tasks {
		fmt.Printf("Indeks: %d\n", i)
		fmt.Printf("Judul: %s\n", t.Title)
		fmt.Printf("Deskripsi: %s\n", t.Description)
		fmt.Printf("Jatuh Tempo: %s\n", t.DueDate.Format("02 January 2006"))
		fmt.Printf("Status: %s\n", t.Status)
		fmt.Println("-------------------------------------------")
	}
}

// Fungsi untuk mengurutkan daftar tugas berdasarkan tanggal jatuh tempo atau status
func SortTasks(tasks []task.Task, sortBy string) []task.Task {
	sortedTasks := make([]task.Task, len(tasks))
	copy(sortedTasks, tasks)

	switch sortBy {
	case "tanggal":
		sort.Slice(sortedTasks, func(i, j int) bool {
			return sortedTasks[i].DueDate.Before(sortedTasks[j].DueDate)
		})
	case "status":
		sort.Slice(sortedTasks, func(i, j int) bool {
			return sortedTasks[i].Status < sortedTasks[j].Status
		})
	}

	return sortedTasks
}
