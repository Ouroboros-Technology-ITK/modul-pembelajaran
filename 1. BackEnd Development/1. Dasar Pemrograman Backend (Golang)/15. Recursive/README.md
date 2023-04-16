# Recursive

## What is recursive?
*Recursive* adalah proses **suatu fungsi memanggil dirinya sendiri** secara langsung ataupun tidak langsung. Hal ini dapat memecahkan masalah tertentu dengan memanggil salinan dirinya sendiri dan menyelesaikannya menjadi sub-masalah yang lebih kecil dari masalah sebelumnya. Contoh *case*-nya yaitu [Towers of Hanoi (TOH), Inorder/Preorder/Postorder Tree Traversals](), dll.

## Why we learn recursive?
Dengan mempelajari recursive, kita dapat mengurangi panjang kode kita dan membuatnya lebih mudah untuk dibaca dan ditulis. Dalam pemrograman, terkadang kita menemukan masalah yang lebih mudah dipecahkan dengan teknik *recursive*. Ini biasanya digunakan untuk menyelesaikan permasalahan yang memiliki keteraturan pola dalam prosesnya. Contohnya adalah *factorial*:

![index](https://prepinsta.com/wp-content/uploads/2022/04/Factorial-of-a-Number-using-Recursion-in-Python.webp)

Factorial pattern (Sumber: [ebookreading.net](https://prepinsta.com/python-program/find-the-factorial-of-the-number-using-recursion/))

Pola di atas bisa diselesaikan dengan mudah dengan proses *recursive*.

Selain itu, *recursive* merupakan salah satu kemampuan fundamental yang harus dikuasai oleh kita yang ingin terjun ke dunia pemrograman. Maka tidak heran jika beberapa pertanyaan terkait *recursive* muncul ketika melakukan coding interview di perusahaan teknologi.


## Recursive in Go
Di Golang kita bisa membuat fungsi recursive dengan pola di bawah ini:
![index](https://raw.githubusercontent.com/Ouroboros-Tech/modul-pembelajaran/Dasar-golang/image/swift-recursion.jpg)

Syntax Recursion (Sumber: [programiz.com](https://www.programiz.com/swift-programming/recursion))

Sebagai contoh, mari kita buat sebuah fungsi penghitung mundur:
```go
package main
import "fmt"

func countDown(number int) {
	// menampilkan parameter "number"
	fmt.Println(number)
	// pemanggilan recursive dengan mengurangi nilai 'number"
	countDown(number - 1)
}
func main() {
	fmt.Println("Countdown Starts: ")
	// pemanggilan fungsi 'countDown()"
	countDown(3)
}
```
Output:
```Output
Countdown Starts:
3
2
1
0
-1
-2
-3
```

Perhatikan fungsi `countDown(number - 1)`, ini adalah pemanggilan fungsi *recursive* dan kita mengurangi satu angka dalam setiap panggilan. Mungkin kita berasumsi bahwa fungsi akan berhenti di angka **0**. Namun kenyataannya, fungsi terus berjalan hingga resource komputer kita habis dan berhenti. Ini kita namakan **infinite loop**.

Ada dua aspek yang harus dimiliki oleh *recursive*, yaitu:
1. Mengetahui kapan harus memanggil dirinya kembali
2. Mengetahui kapan harus berhenti

Fungsi `countDown()` sudah memenuhi aspek pertama namun belum memenuhi aspek kedua. Oleh karena itu kita bisa menggunakan **condition** dan hanya memanggil fungsi jika kondisinya terpenuhi. Berikut adalah contohnya:
```go
package main

import "fmt"

func countDown (number int) {
	if number> 0 {
		fmt.Println(number)
		countDown(number - 1)
	} else {
		fmt.Println("Countdown Stops")
	}
}

func main() {
	fmt.Println("Countdown Starts")
	countDown(3)
}
```
Output:
```Output
Countdown Starts
3
2
1
Countdown Stops
```

Kita memanggil fungsi hanya jika angkanya lebih besar dari **0**. Jika sudah mencapai angka **O** maka *recursive* berakhir.

NUMBER > 0|PRINT|RECURSIVE CALL
--|--|--
true|3|countDown(2)
true|2|countDown(1)
true|1|countDown(0)
false|Countdown stops|function execution stops
Berikut adalah gambaran proses *recursive* di atas:

Recursion (Sumber: programiz.com)

Jadi pada contoh di atas, kita menggunakan `if...else` statement untuk **mencegah infinite loop**.

## Factorial of a number using Go Recursion

*Factorial* adalah case yang sering dipecahkan dengan teknik *recursive*. Untuk mengetahui lebih lanjut tentang *factorial*, kita bisa lihat [di sini]().

Setelah kita mengetahui tentang *factorial*, kita bisa lihat contohnya dengan Golang.
```go
package main

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func main() {
	var number = 3
	var result = factorial(number)
	fmt.Println("The factorial of 3 is", result)
}
```
Output:
```Output
The factorial of 3 is 6
```

Dalam contoh *factorial* di atas, kita membuat fungsi *recursive* bernama `factorial()` yang memanggil dirinya sendiri jika nilai number tidak sama dengan **0**.

```Output
return num factorial (num - 1)
```

Dalam fungsi `factorial()` di parameter setiap pemanggilan fungsi berikutnya, kita mengurangi nilai number sebanyak **1**.

Factorial Recursion Go (Sumber: programiz.com)

Contoh visual proses program recursive mencari faktorial suatu bilangan bisa dilihat [di sini]().

>Note: Masukkan angka di kiri atas untuk menjalankan visual factorial.


## Difference between Recursion and Iteration
Di sini kita akan membahas perbedaan antara *Recursion* dengan *Iteration* dengan contoh kasus **Factorial**:

• Sebuah program disebut **recursive** ketika sebuah fungsi memanggil dirinya sendiri.
• Sebuah program disebut **iteratif** ketika ada loop (atau pengulangan).

Contoh dalam Golang bisa dilihat di bawah ini:
```go
package main
import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func factorialLoop (value int) int {
	result := 1
	for i = value; i > 0; i-- {
		result *= i // result result i
	}
	return result
}
func main() {
	loop := factorialLoop(20)
	fmt.Println(loop)
	recursive := factorialRecursive(20)
	fmt.Println(recursive)
}
```
Output:
```Output
3628800
3628800
```

Untuk penjelasan lebih detail, silahkan kunjungi [di sini]().
