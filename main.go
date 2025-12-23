package main

import (
	"bufio"
	"fmt"
	"minitask1/internals/bio"
	"minitask1/internals/circle"
	insertnumber "minitask1/internals/insertNumber"
	"minitask1/internals/triangle"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Hitung Luas Lingkaran")
		fmt.Println("2. Hitung Keliling Lingkaran")
		fmt.Println("3. Hitung Luas dan Keliling Lingkaran")
		fmt.Println("4. Buat Segitiga '*'")
		fmt.Println("5. Sisipkan Angka?")
		fmt.Println("6. Buat Profile")
		fmt.Println("0. Keluar Aplikasi")

		fmt.Print("pilih menu: ")
		scanner.Scan()
		choice := scanner.Text()
		choice = strings.ToLower(choice)

		switch choice {
		case "1":
			fmt.Println("Menghitung Luas Lingkaran ○")
			fmt.Print("masukan diameter: ")
			scanner.Scan()
			choice := scanner.Text()
			choice = strings.TrimSpace(choice)
			choiceConv, err := strconv.ParseFloat(choice, 64)
			if err != nil {
				fmt.Println("Input Tidak Valid", err)
				return
			}
			circle.LuasLingkaran(float32(choiceConv))
		case "2":
			fmt.Println("Menghitung Keliling Lingkaran ○")
			fmt.Print("masukan diameter: ")
			scanner.Scan()
			choice := scanner.Text()
			choice = strings.TrimSpace(choice)
			choiceConv, err := strconv.ParseFloat(choice, 64)
			if err != nil {
				fmt.Println("Input Tidak Valid", err)
				return
			}
			circle.KelilingLingkaran(float32(choiceConv))
		case "3":
			fmt.Println("Menghitung Luas dan Keliling Lingkaran ○")
			fmt.Print("masukan diameter: ")
			scanner.Scan()
			choice := scanner.Text()
			choice = strings.TrimSpace(choice)
			choiceConv, err := strconv.ParseFloat(choice, 64)
			if err != nil {
				fmt.Println("Input Tidak Valid", err)
				return
			}
			circle.LuasDanKelilingLingkaran(float32(choiceConv))
		case "4":
			fmt.Println("Membuat Segitiga △")
			fmt.Print("masukan tinggi segitiganya: ")
			scanner.Scan()
			choice := scanner.Text()
			choice = strings.TrimSpace(choice)
			choiceConv, err := strconv.ParseInt(choice, 10, 0)
			if err != nil {
				fmt.Println("Input Tidak Valid", err)
				return
			}
			triangle.Segitiga(int(choiceConv))
		case "5":
			fmt.Println("Index 0: 50")
			fmt.Println("Index 1: 75")
			fmt.Println("Index 2: 66")
			fmt.Println("Index 3: 👈 angka akan disisipkan disini")
			fmt.Println("Index 4: 20")
			fmt.Println("Index 5: 32")
			fmt.Println("Index 6: 90")
			fmt.Print("masukan angka yang ingin disisipkan: ")
			scanner.Scan()
			choice := scanner.Text()
			choice = strings.TrimSpace(choice)
			choiceConv, err := strconv.ParseInt(choice, 10, 0)
			if err != nil {
				fmt.Println("Input Tidak Valid", err)
				return
			}
			insertnumber.InsertNumber(int(choiceConv))
		case "6":
			fmt.Println("Membuat Profile")

			fmt.Print("Masukan nama: ")
			scanner.Scan()
			nama := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukan foto (URL/path): ")
			scanner.Scan()
			foto := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukan email: ")
			scanner.Scan()
			email := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukan umur: ")
			scanner.Scan()
			umurStr := strings.TrimSpace(scanner.Text())
			umur, err := strconv.Atoi(umurStr)
			if err != nil {
				fmt.Println("Error: Umur harus berupa angka")
				continue
			}

			fmt.Print("Masukan nomor telepon: ")
			scanner.Scan()
			noTelpStr := strings.TrimSpace(scanner.Text())
			noTelp, err := strconv.Atoi(noTelpStr)
			if err != nil {
				fmt.Println("Error: Nomor telepon harus berupa angka")
				continue
			}

			fmt.Print("Status Menikah? (true/false): ")
			scanner.Scan()
			statusStr := strings.TrimSpace(scanner.Text())
			status, err := strconv.ParseBool(statusStr)
			if err != nil {
				fmt.Println("Error: Status harus true atau false")
				continue
			}

			riwayat := make(map[string]string)
			fmt.Println("Masukan riwayat (ketik 'selesai' untuk berhenti):")
			for {
				fmt.Println("Masukan Kategori (Pendidikan/ketik 'Pekerjaan' jika sudah bekerja)")
				fmt.Print("Kategori: ")
				scanner.Scan()
				kategori := strings.TrimSpace(scanner.Text())

				if strings.ToLower(kategori) == "selesai" {
					break
				}

				if kategori == "" {
					continue
				}

				fmt.Print("Jurusan/Pekerjaan: ")
				scanner.Scan()
				deskripsi := strings.TrimSpace(scanner.Text())

				if deskripsi != "" {
					riwayat[kategori] = deskripsi
				}
			}

			profile, err := bio.CreateBio(nama, foto, email, umur, noTelp, status, riwayat)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			fmt.Println("\n=== PROFILE ANDA ===")
			fmt.Printf("Nama: %s\n", profile.Nama)
			fmt.Printf("Foto: %s\n", profile.Foto)
			fmt.Printf("Email: %s\n", profile.Email)
			fmt.Printf("Umur: %d\n", profile.Umur)
			fmt.Printf("No Telepon: %d\n", profile.NoTelp)
			fmt.Printf("Status: %v\n", profile.Status)

			if len(riwayat) > 0 {
				fmt.Println("Riwayat:")
				for key, value := range riwayat {
					fmt.Printf("  %s: %s\n", key, value)
				}
			}
		case "0":
			fmt.Println("Keluar dari aplikasi, sampai jumpa 🫡")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
