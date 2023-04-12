# Empty Interface


## What is empty interface?
*Interface* di Golang selain sebagai implementasi method signature, sebenarnya merupakan tipe data yang dapat diisi dengan tipe data apapun. Kita menyebutnya dengan "empty interface" atau *interface* kosong yang ditulis dengan *syntax* `interface{}`.

## Empty interface in Go
interface{} merupakan tipe data, sehingga cara penggunaannya sama seperti tipe data lainnya, hanya saja nilai yang diisikan bisa apa saja.
```go
var anything interface{}

anything = "Hello Go!"
anything = 100
anything = false
anything = []string{"Aku", "Ganteng"}
```
Namun, konsep dari `interface{}` di Golang tidak sepenuhnya seperti *dynamic programming* yang ada di bahasa pemrograman lain.

Tipe data `interface{}` merupakan **tipe data tersendiri** yang dapat diisi tipe data apapun. Namun, kita tidak dapat mengubah sifatnya ataupun menggunakan operasi dan fungsi bawaan dari tipe data yang diisi.

Contoh :
```go
func main() {
}
	var anything interface{} = 10 // diisi dengan tipe data number
	
	fmt.Println(anything + 11)
	// error:
	// tidak bisa menggunakan operasi matematik karena interface bukan merupakan tipe data number
	anything = true // diisi dengan tipe data boolean
	fmt.Println(anything && false)
	// error:
	// tidak dapat menggunakan operasi boolean karena interface bukan merupakan tipe 
```

## Type alias any
Khusus untuk Golang versi 1.18, kita dapat menulis *interface* kosong atau tipe data yang dapat menerima tipe data apapun menggunakan *syntax* `any`.
```go
var anything any
anything = "Hello Go!"
anything = 100
anything = false
```

## Type assertion
Untuk membuat *interface* kosong menjadi tipe data yang memiliki sifat asli sesuai dengan tipe data yang diisi, kita dapat mengubahnya menggunakan *syntax* `variable. (type)`.
```go
func main() {
	var anything interface{} = 10 // diisi dengan tipe data number
	var num = anything.(int) + 11 // mengubah tipe data menjadi tipe data number
	fmt.Println(num)
	anything = true // diisi dengan tipe data boolean
	fmt.Println(anything.(bool) && true) // mengubah tipe data menjadi tipe data boolean
}
```
Output:
```Output
21
true
```

Kita juga dapat mengubah tipe data *pointer* dari sebuah `struct`. Caranya adalah menggunakan *type assertion* dengan menambahkan tanda `*` sebelum tipe data yang ingin diubah.
```go
type Person struct {
	Name string
	Age int
}

func main() {
	var person interface{} &Person{
		Name: "Aku",
		Age: 10,
	}

	var name = person.(*Person).Name // diubah menjadi struct 'Person', kemudian mengambil nilai dari  
	fmt.Println(name)
}
```
Output:
```Output
Aku
```
Contoh di atas adalah contoh penggunaan *type assertion* untuk mengubah tipe * dari sebuah struct.

Variabel `person` berupa `interface{}` di assign dengan `&Person{}` yang merupakan *pointer* ke *struct* `Person`.

Kita dapat mengubah variabel `person` yang bertipe `interface{}` menjadi sifat dari *struct* `Person` dengan melakukan assertion menggunakan syntax `person. (*Person).` Kemudian langsung mengambil nilai dari properti `Name` yang bertipe `string` dari *struct* `Person`, sehingga *variable* `name` menjadi tipe `string` dan bernilai `Aku`.

## `nil`
`nil` bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya `nil` berarti memiliki nilai kosong.

`nil` berbeda dengan *default value* yang sudah dijelaskan di materi tipe data. Nil adalah nilai kosong, benar-benar kosong. Ketika sebuah `interface{}` tidak di-*assign* dengan nilai apapun, maka akan berisi `nil` atau kekosongan hakiki.
```go
func main() {
	var anything interface{}
	fmt.Println(anything)
}
```
Output:
```Output
<nil>
```
`nil` bukan hanya terjadi kepada tipe data `interface{}` yang tidak diisi nilai, tetapi ada beberapa tipe data yang bisa diset dengan `nil`:
* pointer
* slice
* `map`
* `channel` (akan kita bahas di materi concurrency)

## Empty interface in slice
Interface kosong dapat digunakan untuk mengisi sebuah slice.
```go
var listData []interface{}

// untuk Golang versi 1.18 atau seterusnya, kita dapat menggunakan tipe data any`
var listData []any
```
Dengan memanfaatkan slice dan `interface{}`, kita bisa membuat data slice yang isinya bisa apa saja.
```go

func main() {
	var listData []interface{}
	listData = append(listData, "Hello Go!") // di assign dengan string
	listData = append(listData, 100) // di assign dengan int
	listData = append(listData, false) // di assign dengan boolean
	listData = append(listData, []string{"Aku", "Ganteng"}) // di assign dengan slice of string
}
```

Empty interface in `map`

Kita juga dapat menggunakan *interface* kosong untuk mengisi sebuah `map`. Namun, tipe data `interface{}` hanya dapat digunakan untuk *value data type* di dalam `map`.
```go
var mapData map[string] interface{}
// untuk Golang versi 1.18 atau seterusnya, kita dapat menggunakan tipe data any
var mapData map[string]any

func main() {
	mapData = map[string]interface{}
	mapData["name"] = "Aku" // di assign dengan string
	mapData["age"] = 10 // di assign dengan int
	mapData["is_true"] = false // di assign dengan boolean
	mapData["list_data"] = []string{"Aku", "Ganteng"} // di assign dengan slice of string
}
```
