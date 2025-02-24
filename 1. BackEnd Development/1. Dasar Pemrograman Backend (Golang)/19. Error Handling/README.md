
# Error handling

## Basic error handling
Salah satu hal yang sering terjadi ketika mengembangkan suatu sistem adalah terjadinya *error*. *Error* yang sering terjadi biasanya terdiri dari dua macam, *compile error* dan *runtime error*.

### Compile error
Compile error merupakan error yang terjadi ketika program di-*compile*. Sebagai contoh kesalahan syntax atau kesalahan assign nilai di tipe data yang berbeda.
```go
var number int = "not a number"
fmt.Println(number)
```

Jika menggunakan *code editor*, *compile error* di Golang akan langsung ditampilkan dengan warna merah, bahkan sebelum kode tersebut dijalankan.
Jika kode tersebut dijalankan, maka akan terjadi *error* seperti pada hasil keluaran berikut.
```Output
cannot use "not a number" (type untyped string) as type int32 in assignment
```
Jika terdapat *compile error*, kita dapat melakukan perbaikan pada kode 
sebelum kode tersebut dijalankan.

### Runtime error
*Runtime error* adalah error yang terjadi ketika program atau kode telah berjalan.
```go
func div(first, second int) int {
return first/second
}

func main() {
	div(10, 0) // runtime error: integer divide by zero
}
```

Sebagai contoh, fungsi `div`  di atas merupakan fungsi yang melakukan pembagian. Jika kita mengeksekusi kode, maka kode akan berjalan. Jika terdapat pembagian dengan bilangan 0, maka kode akan berhenti dan mengeluarkan *error* `division by zero`.
Atau contoh lain, seperti berikut:
```go
func main() {
	var anything interface{} = 123
	fmt.Println(anything. (string) + "is a number")
}
```

Ketika melakukan konversi tipe data dari *interface* kosong ke *string* yang diisi dengan nilai integer, maka akan terjadi **runtime error**.


## `error`
Golang memiliki cara khusus untuk menangani *error*. Terdapat *built-in error handling* yang dapat digunakan untuk menangani *error*, yaitu sebuah *interface* `error`. *Interface* `error` memiliki satu method `Error() string` yang akan mengembalikan informasi tentang `error` tersebut. Dan error dapat berisi `nil`.
```go
var err error
```

Di Golang, banyak sekali fungsi yang mengembalikan nilai balik lebih dari satu. Biasanya, salah satu kembalian adalah bertipe `error`. Contohnya seperti pada fungsi `strconv.Atoi()`. Fungsi tersebut digunakan untuk konversi data `string` menjadi numerik.

Fungsi ini mengembalikan 2 nilai balik. Nilai balik pertama adalah hasil konversi, dan nilai balik kedua adalah `error`. Ketika konversi berjalan mulus, nilai balik kedua akan bernilai `nil`. Sedangkan ketika konversi gagal, penyebabnya bisa langsung diketahui dari *error* yang dikembalikan.
```go
func main() {
	var number int
	var err error

	var test = "123"
	number, err = strconv.Atoi (test) // return number = 123, err = nil
	test = "abc"
	number, err = strconv.Atoi (test) // return number = 0, err = error (invalid syntax)
}
```
Untuk membuat *error*, kita dapat menggunakan dua pendekatan yaitu menggunakan `errors.New` dan `fmt.Errorf`.

Pendekatan pertama, yaitu menggunakan *package* `errors` yang merupakan *package* bawaan dari Golang dengan menggunakan fungsi `New`. Fungsi ini akan mengembalikan *interface* `error`.
```go
var err error = errors.New("Ini adalah error")
```

Pendekatan kedua, menggunakan fmt.Errorf yang merupakan fungsi yang dapat digunakan untuk membuat error. Bedanya dengan pendekatan pertama adalah kita dapat melakukan formatting pada error yang akan dikembalikan.
```go
var err error = fmt.Errorf("ini adalah error %d", 1)
```
Karena `error` merupakan *interface*, maka kita juga dapat membuat *custom error* mengimplementasi *interface* `error`. Syaratnya harus memiliki *method* `Error()` yang akan mengembalikan *string* sebagai informasi *error*-nya.
```go
type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}
func main() {
	var err error = CustomError("Ini adalah error"}
	
	fmt.Println(err)
}
```
Output:
```
Ini adalah error
```


## Error handling
Cara sederhana untuk melakukan *error handling* di Golang adalah dengan mengecek (pengkondisian) apakah nilai sebuah *error* adalah `nil` atau tidak. Jika tidak, maka bisa dikatakan bahwa *error* tersebut terjadi.
```go
if err != nil {
// do something
}
```
Contoh seperti berikut:
```go
func main() {
	var number int
	var err error

	test = "abc"
	number, err = strconv.Atoi (test)
	if err != nil { // error handling
		// terjadi error
		fmt.Println(err)
		return
	}
	
	// tidak terjadi error
	fmt.Println(number)
}
```
output:
```Output
strconv.Atoi: parsing "abc": invalid syntax
```

Kita dapat mengatasi *error* dari fungsi `strconv.Atoi` dengan menggunakan *error handling* berupa kondisi yang dapat digunakan untuk menangani *error*.

## Wrapping error
Ketika kita mendapatkan *error*, kita mungkin menginginkan informasi atau konteks tambahan pada *error* tersebut. Informasi ini dapat berupa nama *function* yang menyebabkan *error* atau informasi lain.

Golang memungkinkan kita untuk membungkus suatu *error* ke dalam *error* yang lain. Umumnya digunakan untuk memberikan hierarki dan informasi tambahan pada *error*. Caranya cukup membungkus *error* adalah menggunakan `fmt.Errorf` dengan *verb* `%w`.

Sebagai contoh, pada kode berikut terdapat function `IsFileExists` yang digunakan untuk mengecek apakah *file* terdapat di dalam direktori di komputer.
```go
func IsFileExists(filename string) (bool, error) {
	file, err := os.Open(filename) // open file
	defer file.Close()
	if err != nil {
		return false, fmt.Errorf("Error in IsFileExists, err: %w", err) // wrapping error
	}
	return true, nil
}
func main() {
	isExist, err := IsFileExists("test.txt")
	if err != nil {
		fmt.Println(err)
		return
}
```
Output:
```Output
Error in IsFileExists, err: open test.txt: no such file or directory
```
Pada fungsi `IsFileExists`, ketika terdapat *error* maka fungsi tersebut mengembalikan *error* dengan informasi tambahan berupa nama *function* dan juga *error* dari pemanggilan *method* `os.Open()`. Untuk mendapatkan *error* yang berada di dalam *error* yang dibungkus (dalam kasus di atas adalah *error* dari pemanggilan fungsi `os.Open()`, kita dapat menggunakan fungsi `Unwrap` dari *package* `errors`.
```go
func main() {
	isExist, err := IsFileExists("file.txt")
	if err != nil {	
		if errWrap := errors.Unwrap (err); errwrap != nil {
			fmt.Println(errWrap)
		}
	}
}
```
Output:
```Output
open test.txt: no such file or directory
```

*Error* yang didapat dari pemanggilan fungsi `Unwrap` hanya berisi *error* dari fungsi `os.Open()`.

## `Error.Is`

Ketika kita membungkus *error* ke dalam *error* lain, kita tidak dapat menggunakan `==` untuk melakukan pengecekan jenis *error*. Cara yang dapat digunakan untuk mengecek jenis *error* pada kasus ini adalah dengan menggunakan `errors.Is`.
```go
func main() {
	isFileExists, err := IsFileExists("test.txt")
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File doesn't exists")
		}
	}
}
```


## Panic
*Panic* adalah kondisi di mana *runtime* Golang tidak dapat memproses suatu baris kode. Hal ini dapat terjadi karena kesalahan dalam menulis kode (seperti pembagian dengan nol) atau masalah lingkungan pengembangan (seperti kekurangan memori). Ketika *panic* terjadi, program akan segera berhenti.
```go
func main() {
fmt.Println(10/0)
}
```
Output:
```Output
./main.go:8:16 division by zero
exit status 2
```

Golang menggunakan *panic* untuk menangani kesalahan yang tidak dapat ditangani oleh program. Kita dapat membuat *panic* dengan menggunakan keyword `panic`. Hal ini dapat kita gunakan ketika ingin menghentikan program jika terjadi *error*.
```go
func main() {
	number, err := strconv.Atoi("abc")
	if err != nil {
		panic(err)
	}
	// tidak melanjutkan baris kode ini jika terjadi error
	count = number + 1
	fmt.Println(count)
}
```
Output
```Output
panic: strconv.Atoi: parsing "abc": invalid syntax

goroutine 1 [running]:
./main.go:96 +0x85
exit status 2
```


## Recover
Kadang kita menghadapi *panic* yang tidak kita inginkan atau berharap kode terus berjalan di kondisi tertentu. Golang memungkinkan kita untuk menangkap *panic* dan mengatasi sesuai dengan keinginan kita menggunakan `recover`. Fungsi `recover` dieksekusi di dalam `defer` untuk mengecek apakah terjadi *panic* atau tidak.

```go
func division (first, second int) int {
	defer func() {
		if v := recover(); v != nil { // menangkap 'panic' yang terjadi 
		fmt.Println(v)
		}
	}())
	
	return first/second // terjadi panic di sini
}

func main() {
	div := division (10, 0)

	// kode akan terus berjalan karena sudah dilakukan recover
	fmt.Println(div)
	fmt.Println("Program Selesai")
}
```
Output:
```Output
runtime error: integer divide by zero
0
Program Selesai
```
Ketika kita menggunakan *recover*, ketika terjadi *panic*, maka program akan terus berjalan.
