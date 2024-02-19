# Golden Ratio Project
![Golden Ratio Visualization](https://github.com/CryptoPendora/GoldenRatio/blob/main/png.png?raw=true)

# Golden Ratio Calculation in Go

This project demonstrates how to calculate the Golden Ratio using high precision arithmetic in Go. The Golden Ratio, often symbolized as φ (phi), is a well-known mathematical constant with the value $`\frac{1+\sqrt{5}}2`$


## Implementation

The main program calculates the Golden Ratio with a precision of 10,000 bits, ensuring an extremely accurate result. It utilizes the `math/big` package for operations with big floating-point numbers.

### Code Explanation

```go
package main

import (
    "fmt"
    "math/big"
)

func main() {
    // Define precision for the calculation
    prec := uint(10000) // For example, 10000 bits

    // Create big.Float numbers for sqrt(5) with defined precision
    sqrt5 := new(big.Float).SetPrec(prec).Sqrt(big.NewFloat(5))

    // Create big.Float numbers for 1 and 2 with the same precision
    one := new(big.Float).SetPrec(prec).SetFloat64(1)
    two := new(big.Float).SetPrec(prec).SetFloat64(2)

    // Calculate the Golden Ratio: (1 + sqrt(5)) / 2
    phi := new(big.Float).SetPrec(prec).Quo(new(big.Float).Add(one, sqrt5), two)

    // Print the Golden Ratio with the defined precision
    fmt.Printf("Golden Ratio: %.*g\n", prec, phi)
}
```

## Overview
This project is dedicated to exploring and implementing the Golden Ratio in various fields of mathematics and design. The Golden Ratio, approximately equal to 1.618033988749895, is a fascinating number often found in nature, art, architecture, and even financial markets.

## Visualization
Included in this repository is an image that visualizes the Golden Ratio in a creative and insightful manner. This image aims to provide a visual understanding of how the Golden Ratio can be applied and appreciated in different contexts.

## Contribution
If you find this project interesting and would like to support our efforts, we welcome contributions of all kinds, not limited to code or financial donations. Your support enables us to continue our research and create more engaging content.

### Donate Bitcoin (BTC)
To support us with a Bitcoin donation, please send your contributions to the following Bitcoin address:

16HXFqHMwFEkubWAmhmiT8W1jDWunV8Ajk
###
![BTC](https://github.com/CryptoPendora/GoldenRatio/blob/main/CP.png?raw=true)


Your donations are greatly appreciated and will be used to fund further development of the project.

## Contact
For more information, suggestions, or inquiries, feel free to reach out to us. We are always looking for collaboration opportunities and feedback on our work.

Thank you for your interest in our project and your support!

คำนวณด้วยความแม่นยำสูง (High Precision Calculation):

ส่วนนี้ของโค้ดตั้งค่าความแม่นยำสำหรับการคำนวณเป็น 10000 บิต เพื่อให้ได้ค่า Golden Ratio ที่แม่นยำมากขึ้น
sqrt5 คือ การคำนวณรากที่สองของ 5 โดยใช้ประเภทข้อมูล big.Float ซึ่งอนุญาตให้มีความแม่นยำสูงในการคำนวณ
จากนั้นใช้ one (ค่า 1) และ two (ค่า 2) ที่เป็น big.Float พร้อมความแม่นยำเดียวกันในการคำนวณสมการ $`\frac{1+\sqrt{5}}2`$
ผลลัพธ์คือค่า phi ซึ่งเป็นการคำนวณค่า Golden Ratio ที่มีความแม่นยำสูง
คำนวณด้วยความแม่นยำปกติ (Standard Precision Calculation):

ในส่วนความคิดเห็นของโค้ด, มีการแสดงวิธีการคำนวณค่า Golden Ratio โดยใช้ความแม่นยำปกติ (ไม่ได้ระบุความแม่นยำเฉพาะเช่น 10000 บิต)
ใช้ big.Float สำหรับการคำนวณ sqrt5 (รากที่สองของ 5) และคำนวณสมการเดียวกัน 
$`\frac{1+\sqrt{5}}2`$
​
 )/2 โดยใช้ค่า one และ two ที่ไม่ได้ตั้งค่าความแม่นยำเฉพาะ
ผลลัพธ์จะได้ค่า phi คือค่า Golden Ratio ที่คำนวณด้วยความแม่นยำปกติของประเภทข้อมูล

