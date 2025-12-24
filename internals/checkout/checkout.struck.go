package checkout

import "fmt"

type ICheckout interface {
	Payment(bill []int)
}

type BankPayment struct{}

func (b BankPayment) Payment(bill []int) {
	totalPayment := 0
	for _, num := range bill {
		totalPayment += num
	}
	fmt.Printf("berhasil melakukan pembayaran sebesar %d dengan Bank payment", totalPayment)
}

type OnlinePayment struct{}

func (b OnlinePayment) Payment(bill []int) {
	totalPayment := 0
	for _, num := range bill {
		totalPayment += num
	}
	fmt.Printf("berhasil melakukan pembayaran sebesar %d dengan Online payment", totalPayment)
}

type FictifPayment struct{}

func (b FictifPayment) Payment(bill []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recover from panic %v\n", r)
		}
	}()
	var Payment = []int{}
	for _, value := range bill {
		if value == 0 {
			panic("value payment must not be zero")
		}
		Payment = append(Payment, value)
	}
	fmt.Println("List Pembayaran")
	for _, value := range Payment {
		fmt.Printf("%v\n", value)
	}
}

var PaymentMethod ICheckout
