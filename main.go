package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"manajemen-tugas/task"
	"manajemen-tugas/taskmanager"
)

func main() {
	tasks := make([]task.Task, 0)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("==== Sistem Manajemen Tugas ====")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Tampilkan Daftar Tugas")
		fmt.Println("3. Perbarui Status Tugas")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Urutkan Daftar Tugas")
		fmt.Println("6. Keluar")
		fmt.Print("Pilihan Anda: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSuffix(choice, "\n")

		switch choice {
		case "1":
			fmt.Print("Judul: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSuffix(title, "\n")

			fmt.Print("Deskripsi: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSuffix(description, "\n")

			fmt.Print("Tanggal Jatuh Tempo (dd-mm-yyyy): ")
			dateStr, _ := reader.ReadString('\n')
			dateStr = strings.TrimSuffix(dateStr, "\n")

			dueDate, err := time.Parse("02-01-2006", dateStr)
			if err != nil {
				fmt.Println("Format tanggal salah. Tugas tidak ditambahkan.")
				continue
			}

			tasks = task.AddTask(tasks, title, description, dueDate)
			fmt.Println("Tugas berhasil ditambahkan.")

		case "2":
			taskmanager.ShowTasks(tasks)

		case "3":
			fmt.Print("Masukkan indeks tugas yang akan diperbarui statusnya: ")
			indexStr, _ := reader.ReadString('\n')
			indexStr = strings.TrimSuffix(indexStr, "\n")

			index, err := strconv.Atoi(indexStr)
			if err != nil || index < 0 || index >= len(tasks) {
				fmt.Println("Indeks tugas tidak valid.")
				continue
			}

			fmt.Print("Masukkan status baru (belum selesai/selesai): ")
			status, _ := reader.ReadString('\n')
			status = strings.TrimSuffix(status, "\n")

			tasks = task.UpdateTaskStatus(tasks, index, status)
			fmt.Println("Status tugas berhasil diperbarui.")

		case "4":
			fmt.Print("Masukkan indeks tugas yang akan dihapus: ")
			indexStr, _ := reader.ReadString('\n')
			indexStr = strings.TrimSuffix(indexStr, "\n")

			index, err := strconv.Atoi(indexStr)
			if err != nil || index < 0 || index >= len(tasks) {
				fmt.Println("Indeks Tugas tidak valid.")
				continue
			}

			tasks = task.RemoveTask(tasks, index)
			fmt.Println("Tugas berhasil dihapus.")

		case "5":
			fmt.Println("Pilih kriteria pengurutan:")
			fmt.Println("1. Tanggal jatuh tempo")
			fmt.Println("2. Status")
			fmt.Print("Pilihan anda: ")

			sortBy, _ := reader.ReadString('\n')
			sortBy = strings.TrimSuffix(sortBy, "\n")

			sortedTasks := taskmanager.SortTasks(tasks, sortBy)
			taskmanager.ShowTasks(sortedTasks)

		case "6":
			fmt.Println("Terima Kasih!")
			os.Exit(0)

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
