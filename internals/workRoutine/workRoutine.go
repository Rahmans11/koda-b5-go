package workroutine

import (
	"fmt"
	"sync"
	"time"
)

func WorkRoutine() {
	var wg sync.WaitGroup
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	wg.Go(func() {
		func() {
			fmt.Println("Mulai Mandi")
			time.Sleep(time.Millisecond * 3000)
			fmt.Println("Selesai Mandi")
			time.Sleep(time.Millisecond * 3000)
		}()
	})
	wg.Wait()
	wg.Go(func() {
		func() {
			fmt.Println("Buat Kopi")
			time.Sleep(time.Millisecond * 3000)
			fmt.Println("Selesai Buat Kopi")
			time.Sleep(time.Millisecond * 3000)
		}()
	})
	wg.Wait()
	wg.Go(func() {
		func() {
			fmt.Println("Mulai Menyiapkan Sarapan")
			time.Sleep(time.Millisecond * 3000)
			fmt.Println("Selesai Menyiapkan Sarapan")
			time.Sleep(time.Millisecond * 3000)
		}()
	})
	wg.Wait()
	wg.Go(func() {
		func() {
			fmt.Println("Mulai Merapikan Kamar")
			time.Sleep(time.Millisecond * 3000)
			fmt.Println("Selesai Merapikan Kamar")
			time.Sleep(time.Millisecond * 3000)
		}()
	})
	wg.Wait()
	fmt.Println("Berangkat Kerja")
	fmt.Println("Meluncur!!!")
}
