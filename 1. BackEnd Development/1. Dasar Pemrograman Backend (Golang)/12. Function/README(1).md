# Name and Parameter  
Mengapa kita memerlukan *function* (fungsi)? 

Bayangkan jika kita menulis bagian kode yang ingin kita gunakan ulang di baris kode lain. Bagaimana cara kita mengimplementasikan hal tersebut?🤔 Apakah kita harus **menulis ulang kode** yang ingin kita gunakan tadi?😨

Tenang, kita dapat membuat sebuah *fungsi* untuk bagian kode yang ingin kita gunakan ulang.☺️

*Fungsi* adalah sekumpulan blok kode yang dibungkus dengan nama tertentu. Jadi kita akan mendapatkan:
* Reusable code: yaitu tidak perlu melakukan copy-paste jika ingin menggunakan bagian kode yang sama.
* Modularity: yaitu mengelompokkan kode sesuai dengan fungsinya.  

Gambaran dari pembuatan fungsi di Golang adalah sebagai berikut:  
![index](https://media.geeksforgeeks.org/wp-content/uploads/20190725175724/Named-Return-Parameters-Golang1.jpg)
(Sumber: [geeksforgeeks.org](https://media.geeksforgeeks.org/wp-content/uploads/20190725175724/Named-Return-Parameters-Golang1.jpg))

## Simple function (without return and param)  
Fungsi `main` merupakan fungsi utama pada program Go. Namun kita bisa membuat fungsi lainnya dengan cara menuliskan *keyword* `func`, diikuti dengan nama fungsi, tanda kurung untuk mengisi parameter, dan tanda kurung kurawal untuk membungkus blok kode.  

Contohnya pada kode berikut terdapat fungsi bernama **addNumbers**.  
```go
func addNumbers() {  
	n1 := 12  
	n2 := 8  
	
	sum := nl + n2  
	fmt.Println("Sum 12+ 8:", sum)  
}  
```

Cara menggunakannya, kita bisa memanggilnya dari fungsi `main`.  
```go
package main  
import "fmt"  

// fungsi addNumbers  
func addNumbers() {  
	n1 := 12  
	n2 := 8  
	
	sum := nl + n2  
	fmt.Println("Sum 12+ 8:", sum)  
}  

func main() {  
	// memanggil fungsi addNumbers  
	addNumbers()  
}  
```
```Output
Sum 12 + 8: 20  
```
Berikut gambaran cara kerja fungsi `addNumbers()`:  
![index](https://www.programiz.com/sites/tutorial2program/files/go-function-working.png)
(Sumber: [programiz.com](https://www.programiz.com/sites/tutorial2program/files/go-function-working.png))

## Parameter  
Parameter adalah variabel yang di sisipkan pada saat pemanggilan fungsi, sehingga fungsi dapat menerima nilai *external* dan menggunakannya di dalam fungsi.  

Kenapa harus menggunakan parameter pada fungsi?🤔 Bayangkan jika kita membuat fungsi `sayHello()`, namun kita ingin menyapa berdasarkan nama yang berbeda pada setiap pemanggilannya. Jadi fungsi harus bekerja secara dinamis.

Untuk menambahkan parameter, kita bisa sisipkan parameter beserta tipe datanya di antara tanda kurung `()` setelah nama fungsi.  

Contoh:  
```go
// satu parameter  
func sayHello(name string) {
	fmt.Printf("Hello, %s", name)
}  

```
Sekarang, fungsi `sayHello()` menerima parameter `name` dengan tipe data string. Ketika memanggil fungsi ini, kita perlu memberikan nilai sesuai dengan tipe datanya.  
```go
// memanggil fungsi  
sayHello("Aditira")  
```
Berikut adalah contoh code lengkapnya:  
```go
package main  
import "fmt"  
// fungsi dengan satu parameter  
func sayHello(name string) {  
	fmt.Printf("Hello %s\n", name)  
}  
func main() {  
// mengirim parameter yang berbeda dalam fungsi yang sama  
	sayHello("Aditira")  
	sayHello("Dito")  
	sayHello("Dina")
}  
```
```Output
Hello Aditira  
Hello Dito  
Hello Dina  
```
Terlihat bahwa fungsi `sayHello()` mencetak nama yang berbeda-beda sesuai dengan parameter yang di kirim saat pemanggilan fungsi  

## Multi parameter  
Masih ingat dengan fungsi `addNumbers()` ? Fungsi ini menjumlahkan 2 angka di dalamnya yaitu 12 dan 8. Namun bagaimana jika ingin menjumlahkan angka yang berbeda-beda namun masih menggunakan fungsi yang sama? 🤔 Yup kita bisa menggunakan parameter. Namun di sini kita butuh 2 parameter yaitu **n1** dan **n2** 😧. Nah, jika parameter berjumlah lebih dari 1, pemisah antar parameter adalah koma `,` .
```go
func addNumbers(n1 int, n2 int) {  
	// code  
}  

```

Sekarang, fungsi `addNumbers()` menerima dua parameter: **n1** dan **n2** bertipe integer.  
```go
// memanggil fungsi dengan 2 parameter  
addNumbers(21, 13)  
```

Di sini, **21** dan **13** adalah argumen yang diteruskan ke fungsi `addNumbers()`.  

Contoh lengkapnya:  
```go
package main  
import "fmt"  

// fungsi menggunakan 2 parameter  
func addNumbers(n1 int, n2 int) {  
	sum := nl + n2  
	fmt.Println("Sum:", sum)  
}  
func main() {  
	// mengirim 2 parameter pada fungsi addNumbers  
	addNumbers(21, 13)  
}  
```

Berikut gambaran cara kerja program fungsi `addNumbers()` dengan 2 parameter:  
![index](https://www.programiz.com/sites/tutorial2program/files/go-function-parameter.png)
(Sumber: [programiz.com](https://www.programiz.com/sites/tutorial2program/files/go-function-parameter.png))

Dengan *multi parameter* kita juga bisa menambahkan parameter dengan tipe data yang berbeda-beda. 
Contoh:  
```go
func sayHello(name string, age int, present bool)  
```

## Variadic parameter  
*Variadic parameter* memungkinkan kita untuk memasukkan jumlah nilai parameter yang dinamis dengan tipe data yang sama. Deklarasinya dengan menambahkan tanda titik tiga kali (`...`) tepat setelah penulisan variabel (sebelum tipe data). Contoh:  
```go
package main  
import "fmt"  
func calcAvg(numbers ...int) float64 {  
	var total int = 0  
	for _, number := range numbers {  
		total += number
	}  
  
	var avg = float64(total) / float64(len(numbers))
	return avg  
}  

func main() {  
	var avg = calcAvg (2, 4, 3, 5, 4, 3, 3, 5, 5, 3)  
	var msg fmt.Sprintf("Rata-rata: %.2f", avg)  
	fmt.Println(msg)  
}
```
```Output
Rata-rata 3.70  
```

Perhatikan pada fungsi `calcAvg()`, parameter numbers di sisipkan tanda 3 titik (`...`), menandakan bahwa `numbers` adalah sebuah parameter variadic dengan tipe data `int`.  
```go
func calcAvg(numbers ...int) float64 {  
```

Pemanggilan fungsi dilakukan seperti biasa, hanya saja jumlah parameter yang di sisipkan bisa banyak.  
```go
var avg = calcAvg(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)  
```

Nilai tiap parameter bisa diakses metode looping `for` - `range`.  
```go
for _, number := range numbers {  
```

Jadi, kapan kita menggunakan fungsi *variadic*? Ketika input untuk parameter tidak diketahui atau tidak pasti jumlahnya, kita dapat gunakan *variadic parameter*.  

Contoh lainnya:  
```go
func Sum(numbers ...int) int {  
	total := 0
	for _, number := range numbers {  
		total += number
	}  
	return total  
}  

func main() {  
	fmt.Println(Sum(1, 2)) // 3  
	fmt.Println(Sum(1, 2, 3, 4, 5)) // 15  
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // 55  
}
```

## Fungsi Dengan Parameter Biasa & Variadic  
Parameter *variadic* bisa dikombinasikan dengan parameter biasa, **dengan syarat parameter variadic-nya harus diposisikan di akhir**. Contoh:  
```go
package main

import (
	"fmt"
	"strings"
)

func myHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func main() {
	myHobbies("Aditira", "Coding", "Gaming", "Jogging")
}
```
```Output
Hello, my name is: Aditira  
My hobbies are: Coding, Gaming, Jogging  
```
