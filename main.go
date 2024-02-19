//donate 16HXFqHMwFEkubWAmhmiT8W1jDWunV8Ajk

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// กำหนดความแม่นยำสำหรับการคำนวณ
	prec := uint(10000) // ตัวอย่างเช่น 10000 บิต

	// สร้างจำนวนทศนิยมขนาดใหญ่สำหรับ sqrt(5) พร้อมกำหนดความแม่นยำ
	sqrt5 := new(big.Float).SetPrec(prec).Sqrt(big.NewFloat(5))

	// สร้างจำนวนทศนิยมขนาดใหญ่สำหรับ 1 และ 2 พร้อมกำหนดความแม่นยำ
	one := new(big.Float).SetPrec(prec).SetFloat64(1)
	two := new(big.Float).SetPrec(prec).SetFloat64(2)

	// คำนวณค่า (1 + sqrt(5)) / 2
	phi := new(big.Float).SetPrec(prec).Quo(new(big.Float).Add(one, sqrt5), two)

	// พิมพ์ค่า Golden Ratio พร้อมกับความแม่นยำที่กำหนด
	fmt.Printf("Golden Ratio: %.*g\n", prec, phi)
}

/*
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// สร้างและคำนวณ sqrt(5)
	sqrt5 := new(big.Float).Sqrt(big.NewFloat(5))

	// สร้างจำนวนทศนิยมขนาดใหญ่สำหรับ 1 และ 2
	one := big.NewFloat(1)
	two := big.NewFloat(2)

	// คำนวณค่า (1 + sqrt(5)) / 2
	phi := new(big.Float).Quo(new(big.Float).Add(one, sqrt5), two)

	// พิมพ์ค่า Golden Ratio
	fmt.Println("Golden Ratio:", phi)
}
*/
