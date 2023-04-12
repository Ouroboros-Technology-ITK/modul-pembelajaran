# Return

Perintah `return` pada fungsi akan memberikan balikan nilai yang bisa digunakan di luar fungsi tersebut.

Masih ingat dengan fungsi `addNumbers()` yang sudah ditambahkan parameter untuk menjumlahkan 2 angka berikut ini?
```go
func addNumbers(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("Sum:", sum)
}
```

Sebagai contoh, kita akan gunakan kembali fungsi `addNumbers()`, namun hasil dari penjumlahan tersebut dicetak di dalam fungsi itu sendiri.

Oleh karena itu, berikut cara kita membuat fungsi yang mengembalikan nilai:
```go
func addNumbers(n1 int, n2 int) int {
	sum := nl + n2
	return sum
}
```

Di sini, `int` sebelum tanda kurung kurawal pembuka `{` menunjukkan tipe data untuk balikan nilai dari fungsi. Dalam hal ini, `int` berarti fungsi akan mengembalikan nilai integer.

Dan `return sum` adalah kembalian dari nilai variabel `sum`.

Fungsi mengembalikan nilai ke tempat di mana dia dipanggil, sebaiknya kita menyimpan balikan nilai ke variabel.
```go
// function call
result := addNumbers(21, 13)
```

Di sini, kita menyimpan balikan nilai ke variabel result.

Berikut contoh lengkapnya:
```go
package main
import "fmt"
// function definition
func addNumbers(n1 int, n2 int) int {
	sum := nl + n2
	return sum
}

func main() {
	// function call
	result := addNumbers (21, 13)
	
	fmt.Println("Sum:", result)
}

```
```Output
Sum: 34
```

Dalam contoh di atas, menggunakan `return` kita mengembalikan nilai variabel `sum`. Balikan nilai ditampung ke variabel `result` di dalam fungsi `main`.

Berikut adalah cara kerja program:
![index](https://raw.githubusercontent.com/Ouroboros-Tech/modul-pembelajaran/Dasar-golang/image/go-function-return.jpeg.jpg)
(sumber: [programiz.com](https://www.programiz.com/sites/tutorial2program/files/go-function-return.png))

## Multiple return
Di Golang, sebuah `function` dapat:
• Tidak mengembalikan nilai
• Mengembalikan satu nilai
• Atau mengembalikan lebih dari satu nilai
Jika balikan nilai lebih dari 1, maka balikan nilai dituliskan di dalam tanda kurung `()`

Pada `function` yang memiliki `multiple return`, biasanya `return` yang terakhir berupa `error` yang digunakan untuk *error handling*.
```go
func GetData(key string) ([]string, error) {
	data, err := getDataFromDB(key)
	if err != nil {
		return nil, err
}
	return data, nil
}
```

## Named return
Di Golang memungkinkan kita untuk memberi nama variabel pada nilai `return`. Nama variabel pada nilai `return` ini dideklarasikan secara langsung ketika fungsi ini dieksekusi.

Contoh
```go
func GetDataFromMap (dataDict map[string]string, key string) (data string, err error) {
	if _, ok := dataDict[key]; !ok {
		return key, fmt.Errorf("Key %s not found", key)
	}
	data = dataDict[key]
	
	return data, err
}
```

Ketika menggunakan *named return*, kita bisa mengganti balikan nilai dengan *keyword* `return` saja.
```go
func GetDataFromMap (dataDict map[string]string, key string) (data string, err error) { 
	if, ok := dataDict[key]; !ok {
		return key, fmt.Errorf("Key %s not found", key)
	}
	data = dataDict[key]
	
	return
}
```

## Function helper
Manfaat lain dari fungsi adalah kita bisa membuat fungsi untuk membantu menyelesaikan *problem* pada fungsi yang kita buat. Ini biasa disebut dengan fungsi *helper*.

Contoh kasusnya, kita akan membuat fungsi untuk menyatakan apakah sebuah angka merupakan angka palindrome atau bukan? Untuk mengetahui apa itu *palindrome* bisa kita [lihat di sini]().

Kita akan membuat fungsi `reverse()` sebagai helper untuk membantu menyelesaikan fungsi `isPalindrome()` di bawah ini:
```go
package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// panggil fungsi reverse sebagai fungsi helper
	xReversed := reverse(x)

	return xReversed == x
}

func reverse(x int) int {
	var reversedDigit int
	for x != 0 {
		lastDigitx := x % 10
		reversedDigit = reversedDigit*10 + lastDigit

		x = x / 10
	}
	return reversedDigit
}

func main() {
	output := ispalindrome(121)
	fmt.Println(output)

	output = ispalindrome(12)
	fmt.Println(output)

	output = isPalindrome(-101)
	fmt.Println(output)
}
```
```Output
true
false
false
```

Terlihat bahwa fungsi `isPalindrome()` menggunakan fungsi `reverse()` untuk memecahkan masalah, sehingga kode sekarang terlihat lebih **modular** sesuai dengan fungsinya-masing masing.
