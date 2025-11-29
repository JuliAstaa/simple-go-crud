package main

import (
	"fmt"

	"simple-crud/core/user"
)

func main() {
	var choices int

	for {
		endApp := false

		fmt.Println("Selamat datang di aplikasi CRUD sederhana")
		fmt.Println("1. Semua Data")
		fmt.Println("2. Data Berdasarkan ID")
		fmt.Println("3. Edit Data")
		fmt.Println("4. Delete Data")
		fmt.Println("5. Create Data")
		fmt.Println("6. Exit")
		fmt.Print("Silahkan pilih menu [1 2 3 4 5 6]: ")
		fmt.Scan(&choices)

		switch choices {
		case 1:
			fmt.Println(user.GetAllUsers())

		case 2:
			var id int
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&id)

			result, err := user.GetUserById(id)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(result)
			}

		case 3:
			var id int
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&id)
			var username, email string
			fmt.Print("Masukkan username: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan email: ")
			fmt.Scan(&email)

			err := user.EditUser(username, email, id)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("User dengan ID ", id, "berhasil diedit")
			}

		case 4:
			var id int
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&id)

			result, err := user.DeleteUser(id)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("User dengan ID ", result.ID, "berhasil dihapus")
			}

		case 5:
			var username, email string
			fmt.Print("Masukkan username: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan email: ")
			fmt.Scan(&email)

			newUser := user.User{
				Username: username,
				Email:    email,
			}

			_, err := user.RegisterUser(newUser)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Berhasil tambah data")
			}

		case 6:
			endApp = true

		default:
			endApp = true
		}

		if endApp {
			break
		}

	}
}
