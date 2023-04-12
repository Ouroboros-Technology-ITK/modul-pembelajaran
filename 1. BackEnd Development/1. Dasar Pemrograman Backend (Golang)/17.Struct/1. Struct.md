# Struct

## What is struct?
Struct adalah kumpulan dari beberapa variabel dengan beragam tipe data. Struct biasanya dikenal dengan istilah records

## Why we use struct?
Bayangkan, misalnya kita ingin menyimpan data mahasiswa. Kita bisa saja membuatnya seperti ini:
```go
var name string = "Aditira"
var address string = "Cianjur"
var age int = 21
```

Rasanya terlihat kurang bagus! 🤔 Agar lebih terstruktur, maka kita gunakan struct:
```go
type student struct {
	name string
	address string
	age int
}
```
Kita bisa membuat beberapa variabel (`field`) dalam satu struct yang memiliki tipe data sesuai dengan kebutuhan. Dengan membuat struct, *grouping data* akan lebih rapi dan mudah untuk di-*maintain*.

## Basic struct declaration
*Struct* dapat kita buat dengan *keyword* `type` diikuti dengan nama *struct* dan *keyword* `struct` kemudian isi *struct* kita buat di antara tanda `{` dan `}`. Berikut adalah *syntax* lengkap untuk mendeklarasikan sebuah *struct*.
```go
type nama_struct struct {
	field1 tipe_data
	field2 tipe_data
	.
	.
	.
	field3 tipe_data
}
```

## Tagging
Saat membuat *struct* di Golang, kita juga bisa menambahkan *tagging* setelah tipe data pada setiap *field*. *Tagging* itu sendiri merupakan informasi opsional yang bisa ditambahkan pada masing-masing *field* di *struct*.
```go
type person struct {
	name string `tag1`
	age	 int	`tag2`
}
```

Tagging biasanya dimanfaatkan untuk keperluan `encode/decode` data JSON. Kita akan bahas ini lebih dalam pada materi **JSON**.

## Struct implementation
Cara mengimplementasikan *struct* cukup mudah, sebagai contoh kita akan buat *struct* bernama `person` lalu *struct* tersebut di-*assign* pada sebuah variabel dan kita isi dengan nilai yang sesuai dengan *field* pada *struct* yang di-*assign*:
```go
package main
import "fmt"

// struct person
type person struct {
	name string
	age int
}

func main() {
	var student person // struct di assign ke variable student
	// struct diisi
	student.name = "Aditira"
	student.age = 25
	// struct dicetak
	fmt.Println("name  :", student.name)
	fmt.Println("grade :", student.age)
}
```
Output:
```Output
name : Aditira
grade 25
```

Terlihat bahwa *struct* diisi melalui variabel `student`, maka ini dinamakan sebagai **Indirect Implementation**.

Semua *field* pada *struct* awalnya memiliki *zero value* sesuai dengan tipe datanya.

*Field* dari *struct* bisa diakses nilainya menggunakan notasi titik, contohnya `student.name` . Nilai *field* nya juga bisa diubah, contohnya `student.age = 28`.

Deklarasi *struct* bisa juga bisa dituliskan secara *horizontal*, contohnya:
```go
type person struct { name string; age int}
```
Tanda titik koma ( ;) digunakan sebagai pembatas *field* yang dituliskan secara *horizontal*.

>Note: Namun pada *text editor* yang terinstal *plugin* Go di dalamnya, cara ini tidak akan bisa dilakukan, karena setiap kali *file* di-*save*, kode akan dirapikan otomatis.

Untuk inisialisasi nilai pada struct bisa dilakukan seperti ini:
```go
var student = person{name: "Aditira", age: 28}
```

Berikut adalah contoh lengkapnya:
```go
package main
import "fmt"
// struct person
type person struct { name string; age int}

func main() {
// struct di assign ke variable student dan langsung diisi dengan nilai:
	var student person{name: "Aditira", age: 28}
	// struct dicetak
	fmt.Println("name: ", student.name)
	fmt.Println("grade :", student.age)
}
```
Terlihat bahwa *struct* diisi pada saat *struct* di-*assign* ke variabel `student`, maka ini dinamakan sebagai **Direct Implementation**.

## Embed struct
**Embedded** *struct* adalah mekanisme untuk menempelkan sebuah *struct* sebagai *field* dari *struct* lain. Agar lebih mudah dipahami, kita lihat contoh berikut:
```go
package main
import "fmt"
// struct person
type person struct {
	name string
	age int
}

// struct student
type student struct {
	grade int
	person // struct person sebagai field dari struct student
}

func main() {
	var r1 student()
	r1.name "Aditira"
	r1.age 21
	r1.grade = 2
	
	fmt.Println("name  :", r1.name)
	fmt.Println("age   :", r1.age)
	fmt.Println("age   :", r1.person.age)
	fmt.Println("grade :", r1.grade)
}
```
Output:
```Output
name : Aditira
age : 21
age : 21
grade: 2
```
Dari kode di atas, perhatikan bahwa *struct* `person` di-embed ke dalam *struct* `student`. Jadi terlihat bahwa caranya cukup mudah, cukup dengan menuliskan nama *struct* yang ingin di-*embed* ke dalam *body struct target*.

*Embedded struct* adalah *mutable*, jadi nilai *field*-nya nya bisa diubah.

Khusus untuk field yang bukan asli (*field* turunan dari *struct* lain), bisa diakses dengan cara mengakses *struct parent*-nya terlebih dahulu, contohnya `r1.person. age`. Nilai yang ditampilkan sama dengan `r1.age`.

## Embedded Struct With Same Property Name
Jika salah satu nama *field* sebuah *struct* memiliki kesamaan dengan field milik *struct* lain yang di-embed, maka pengaksesan *field*-nya harus dilakukan secara eksplisit atau jelas. Contoh:
```go
package main
import "fmt"

type person struct {
	name string
	age int
}

type student struct {
	person
	age    int
	grade  int
}

func main() {
var r1
	student{}
	r1.name = "Aditira"
	r1.age = 21 // age dari struct student
	r1.person.age = 22 // age dari struct person
	fmt.Println(r1.name)
	fmt.Println(r1.age)
	fmt.Println(r1.person.age)
}
```
Output:
```Output
Aditira
21
22
```

Jika terjadi kesamaan seperti yang terlihat pada *field* `age`, maka cara mengakses *property* `age` milik *struct* `person` harus melewati *field struct* `student`, yaitu dengan menuliskan nama *struct* yg di-*embed* diikuti nama *property*-nya, contohnya: `r1.person.age = 22`

.
## Nested struct
*Nested struct* adalah anonymous struct yang di-*embed* ke sebuah *struct*. Deklarasinya langsung di dalam struct peng-*embed*. Contoh:
```go
type student struct {
	person struct {
		name string
		age int
	}
	grade int
	hobbies []string
}
```
Teknik ini biasa digunakan ketika decoding data JSON yang struktur datanya cukup kompleks.

## Slice of struct
Kita sudah belajar bahwa *slice* adalah tipe data yang bisa menyimpan lebih dari satu data yang sejenis. Data tersebut bisa berupa tipe data apapun, termasuk *struct*. Kita cukup membuat *slice* dengan tipe data struct yang kita inginkan. Contohnya:
```go
type student struct {
	name string
	age int
}

func main() {
	var Students []student
}
```
Atau dapat menggunakan *keyword* `make` untuk membuat *slice*.
```go
func main() {
	var students = make([] student, 0)
}
```
Contoh di atas membuat sebuah *slice* dengan tipe data `student`. Kita bisa menambahkan data ke *slice* tersebut dengan cara yang sama seperti menambahkan data ke *slice*.
```go
func main() {
	var students []student
	students = append(students, student{"Aditira", 21})
	students = append(students, student("Rahmat", 22})
	fmt.Println(students)
}
```
Output:
```Output
> go run main.go
[{Aditira 21} {Rahmat 22}]
```

Atau dapat langsung menambahkan data ke *slice* saat deklarasi.
```go
func main() {
	var students = []student{
		{"Aditira", 21},
		{"Rahmat", 22}
	}
}
```

Kita juga dapat mengakses properti *struct* yang ada di dalam *slice* dengan cara yang sama seperti mengakses properti *struct* biasa, namun kita perlu mendapatkan *index* dari data yang ingin kita akses terlebih dahulu. Contohnya:
```go
func main() {
	var students = []student{
		student{"Aditira", 21},
		student{"Rahmat", 22},
	}
	// mendapatkan value dari field name di data pertama (index ke 0)
	fmt.Println(students[0].name)
	
	// mendapatkan value dari field age di data kedua (index ke 1)
	fmt.Println(students[0].age)
}
```
```Output
> go run main.go
Aditira
21
```

Contoh di atas adalah mengakses *property / field* `name` dan `age` dari data pertama di dalam *slice* `students`.
