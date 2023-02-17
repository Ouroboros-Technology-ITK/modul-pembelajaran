# Casting / Convert

## Conversion between numbering data type
Golang memiliki *built-in function* yang dapat mengubah antar tipe data *number*. Contohnya, dari `int` ke `float32` atau dari `int` ke `int64`.

Caranya cukup menggunakan conversion operator `data_type(data)`.
```go
func main() {
	var number int32 = 10
	
	var bigNum int64(number) // convert int32 to int64
	
	var floatNum = float32(number) // convert int32 to float32
}
```

Contoh di atas mengubah tipe data `int32` di *variable* `number` ke tipe data `int64` di *variable* `bigNum` dan mengubah tipe data `float32` di variable `floatNum`

## fmt.Sprintf
Fungsi dari `Sprintf` adalah untuk *formating* data yang akan ditulis ke *output*. Namun, kita dapat memanfaatkan fungsi ini untuk mengubah semua tipe data yang ada di Golang menjadi `string`.
Kita menggunakan *format verbs* (contoh: `%d`, `%t`, `%f`, dll) di Golang untuk mengubah setiap tipe data menjadi `string` yang bisa ditampung di variabel.
```go
func main() {
	var number int32 = 10
	var isMaried bool = false
	
	var numString = fmt. Sprintf("%d", number) // convert int32 ke string
	var isMariedString fmt.Sprintf("%t", isMaried) // convert bool ke string
	var Pi fmt.Sprintf("%f", math.Pi) // convert float64 dari variable 'Pi' di package 'math' ke string
	
	fmt.Println(numString)
	fmt.Println(isMariedString)
	fmt.Println(Pi)
	
	fmt.Println("type numString: ", reflect.Typeof(numString))
	fmt.Println("type isMarriedString: ", reflect.Typeof(isMariedString))
	fmt.Println("type Pi:", reflect.Typeof(Pi))
}
```
```Output
10
false
3.141592653589793
type numString: string
type isMarriedString: string
type Pi: string
```

Kita dapat menggunakan fungsi `reflect. TypeOf()` untuk mengecek tipe data dari sebuah *variabel*. Semua *output* dari contoh di atas akan menjadi `string` yang berisi nilai dari tipe data yang diubah.

## strconv.itoa
Terdapat fungsi khusus yang bisa kita gunakan untuk mengubah dari tipe data `int` ke `string` yaitu menggunakan fungsi `Itoa` dari *package* `strconv`.

Caranya cukup mengisi parameter di `Itoa` dengan nilai `int`, kemudian hasilnya akan mengembalikan tipe data `string` yang bisa ditampung di variabel.
```go
func main() {
	var str = strconv.Itoa (10) // convert string to int
	fmt.Println(str)
	fmt.Println("type:", reflect.Typeof(str))
}
```
```Output
10
type: string
```

## strconv.Atoi
Jika ingin sebaliknya, yaitu mengubah dari tipe data `string` ke tipe data `int`, kita bisa menggunakan fungsi `Atoi` dari *package* `strconv`.

Bedanya dengan fungsi `Itoa` adalah fungsi ini mengembalikan 2 nilai, yaitu hasil konversi dengan nilai `int` dan `error`. Error akan dibahas lebih lanjut di materi `error handling`.

```go
func main() {
	var num, _ = strconv.Atoi("10") // convert string to int
	
	fmt.Println(num)
	fmt.Println("type:", reflect.Typeof(num))
}
```
```Output
10
type: int
```

Kita hanya menangkap nilai `int` di *variable* `num`, tapi kita tidak menangkap nilai `error`, sehingga bisa menggunakan *keyword underscore*. Hasil dari konversi ini adalah mengubah `string` dengan nilai `"10"` menjadi `int` dengan nilai `10`.
