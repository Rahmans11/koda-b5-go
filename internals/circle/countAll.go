package circle

import "fmt"

func LuasDanKelilingLingkaran(d float32) {
	keliling := d * 22 / 7
	luas := d * d * 1 / 4 * 22 / 7
	fmt.Printf("Keliling Lingkaran %.2f\nLuas Lingkaran %.2f", keliling, luas)
}
