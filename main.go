package main

import (
	"bufio"
	"fmt"
	"minitask1/internals/bio"
	blackboardchan "minitask1/internals/blackboardChan"
	"minitask1/internals/checkout"
	"minitask1/internals/circle"
	insertnumber "minitask1/internals/insertNumber"
	"minitask1/internals/person"
	"minitask1/internals/reader"
	"minitask1/internals/triangle"
	workroutine "minitask1/internals/workRoutine"
	"os"
	"strconv"
	"strings"
	"sync"
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
		fmt.Println("7. Membaca Catatan")
		fmt.Println("8. Buat Akun Person")
		fmt.Println("9. Pembayaran")
		fmt.Println("10. Simulasi Bekerja")
		fmt.Println("11. BlackBoard")
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
		case "7":
			fmt.Println("Membaca File")
			fmt.Print("tekan 'enter' untuk membaca file: ")
			scanner.Scan()
			choice := scanner.Text()
			choice = strings.TrimSpace(choice)
			reader.ReaderFile()
		case "8":
			fmt.Println("Membuat Akun Person")

			fmt.Print("Masukan nama: ")
			scanner.Scan()
			nama := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukan Alamat: ")
			scanner.Scan()
			alamat := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukan Nomor Telpon: ")
			scanner.Scan()
			phone := strings.TrimSpace(scanner.Text())

			personData := person.NewPerson(nama, alamat, phone)
			fmt.Println("Akun Person Anda:\n", personData.PrintPerson())

			for {
				fmt.Println("1. Tampilkan Salam")
				fmt.Println("2. Ubah nama")
				fmt.Println("0. Kembali")

				fmt.Print("pilih menu: ")
				scanner.Scan()
				choice := scanner.Text()
				choice = strings.ToLower(choice)

				switch choice {
				case "1":
					fmt.Println(personData.GreetWithNamePerson())
				case "2":
					fmt.Print("Masukan nama baru: ")
					scanner.Scan()
					newName := strings.TrimSpace(scanner.Text())
					personData.UpdateNamePerson(newName)
				case "0":
					return
				default:
					fmt.Println("Pilihan tidak valid")
				}
			}
		case "9":
			fmt.Println("==Pembayaran==")

			var bills = []int{}

			for {
				fmt.Println("ketik 'selesai' untuk lanjut")
				fmt.Print("Masukan harga barang: ")
				scanner.Scan()
				choice := strings.TrimSpace(scanner.Text())

				if strings.ToLower(choice) == "selesai" {
					if len(bills) == 0 {
						fmt.Println("Belum ada barang yang ditambahkan")
						continue
					}
					break
				}

				if choice == "" {
					continue
				}

				choiceConv, err := strconv.Atoi(choice)
				if err != nil {
					fmt.Println("Input Tidak Valid. Harap masukkan angka.")
					continue
				}

				if choiceConv <= 0 {
					fmt.Println("Harga harus lebih dari 0")
					continue
				}

				bills = append(bills, choiceConv)
			}

			for {
				fmt.Println("Pilih metode pembayaran:")
				fmt.Println("1. Online Payment")
				fmt.Println("2. Bank Payment")
				fmt.Println("3. Fiktif Payment")
				fmt.Println("0. Keluar")

				fmt.Print("Pilih metode: ")
				scanner.Scan()
				choice := strings.TrimSpace(strings.ToLower(scanner.Text()))

				var paymentMethod checkout.ICheckout

				switch choice {
				case "1":
					paymentMethod = checkout.OnlinePayment{}
					paymentMethod.Payment(bills)
					return
				case "2":
					paymentMethod = checkout.BankPayment{}
					paymentMethod.Payment(bills)
					return
				case "3":
					paymentMethod = checkout.FictifPayment{}
					paymentMethod.Payment(bills)
					return
				case "0":
					return
				default:
					fmt.Println("Pilihan tidak valid")
					continue
				}
			}
		case "10":
			fmt.Println("Simulasi Berankat Kerja")
			fmt.Print("tekan 'enter' untuk memulai Simulasi: ")
			scanner.Scan()
			choice := scanner.Text()
			choice = strings.TrimSpace(choice)
			workroutine.WorkRoutine()
		case "11":
			fmt.Println("Black Board...")

			boxMessageChan := make(chan blackboardchan.Message)

			var messages []blackboardchan.Message

			fmt.Println("Masukan Pesan (ketik 'selesai' untuk berhenti):")
			for {
				fmt.Print("Masukan Nama Pengirim: ")
				scanner.Scan()
				sender := strings.TrimSpace(scanner.Text())

				if strings.ToLower(sender) == "selesai" {
					break
				}

				if sender == "" {
					continue
				}

				fmt.Print("Masukan Pesan: ")
				scanner.Scan()
				content := strings.TrimSpace(scanner.Text())

				if content != "" {
					messages = append(messages, blackboardchan.NewMessage(sender, content))
				}
			}
			fmt.Print("Masukan jumlah pengiriman: ")
			scanner.Scan()
			amountStr := strings.TrimSpace(scanner.Text())
			amount, err := strconv.Atoi(amountStr)
			if err != nil {
				fmt.Println("invalid input")
				continue
			}

			var wg sync.WaitGroup

			go func() {
				defer close(boxMessageChan)
				defer fmt.Println("Pengiriman Selesai")
				for i := 1; i <= amount; i++ {
					wg.Add(1)
					go blackboardchan.Blackboard(&wg, boxMessageChan)
					for _, msg := range messages {
						blackboardchan.MsgSender(boxMessageChan, msg)
					}
				}
			}()
			fmt.Println("Mulai mengirim pesan")
			wg.Wait()
		case "0":
			fmt.Println("Keluar dari aplikasi, sampai jumpa 🫡")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
