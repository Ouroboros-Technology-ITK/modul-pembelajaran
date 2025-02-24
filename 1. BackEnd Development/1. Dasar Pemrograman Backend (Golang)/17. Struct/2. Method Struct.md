# Method Struct

## What is method?
*Method* adalah fungsi yang menempel pada *type* (bisa `struct` atau tipe data lainnya). *Method* bisa diakses lewat variabel objek.

Keunggulan *method* dibanding fungsi biasa adalah memiliki akses ke *property* `struct` hingga **level private** (dibahas lebih lanjut di materi `package & import`). Dan, dengan menggunakan *method* sebuah proses bisa [di-enkapsulasi]() dengan baik.

## Method in Go
Cara membuat method di Golang hampir sama seperti membuat fungsi (func), tetapi dalam deklarasinya harus ditentukan siapa pemilik (objek) dari method tersebut.
```go
type Hello struct {
}

func (h Hello) sayHello() {
	fmt.Println("Hello Go!")
}
```

Contoh di atas adalah bentuk *method* `sayHello()` yang dimiliki oleh *struct* `Hello`. Kita cukup menambah deklarasi objek di antara *syntax* `func` dan nama *methodnya*.

Kita menyebutnya `(h Hello)` dengan istilah **value receiver**. `h` menjadi `alias` dari *struct* `Hello` yang akan menjadi pemilik *method* tersebut. Penamaan alias bebas, umumnya huruf depan dari nama *struct*.

Cara mengakses *method* sama seperti pengaksesan *property* berupa variabel. Tinggal panggil saja *methodnya*.
```go
func main() {
	h := Hello{}
	h.sayHello() // memanggil method sayHello() dari struct 'Hello'
}
```
```
> go run main.go
Hello Go!
```

Kita juga dapat menggunakan atau memodifikasi *property* dari *struct* di dalam *method* menggunakan alias yang sudah dideklarasikan.
```go
type Person struct {
	Name string
	Age int
}
func (p Person) sayHello() {
	fmt.Println("Hello, my name is", p.Name)
	fmt.Println("My age is", p.Age)
}

func main() {
	P := Person{
		Name: "John",
		Age: 30,
	}
	
	p.sayHello()
}
```
```
> go run main.go
Hello, my name is John
My age is 30
```
Contoh di atas adalah bentuk method `sayHello()` yang dimiliki oleh *struct* `Person` dengan menggunakan *property* `Name` dan `Age` dari *struct* tersebut.

*Method* memiliki sifat yang sama persis dengan fungsi biasa. Bisa memiliki parameter, memiliki nilai balik, dll. Dari segi sintaks, pembedanya hanya ketika pengaksesan dan deklarasi.
```go
type Student struct {
	Name string
	Grade int
}

// method dapat berisi parameter dan dapat mengembalikan nilai
func (s Student) getNameAt(i int) string {
	return strings.Split(s.name," ")[i-1]
}
func main() {
	S := Student{
		Name: "Budi Setiawan",
		Grade: 5,
	}

	var nickname = s.getNameAt (0)
	fmt.Println(nickname)
}
```
```Output
> go run main.go
Budi
```

## Method pointer
*Method pointer* adalah *method* di mana *value receiver* berupa *pointer*.

Kelebihannya adalah, ketika kita melakukan manipulasi nilai pada property lain yang masih satu *struct*, nilai pada *property* tersebut akan bersifat *pass by reference*..

Contoh berikut:
```go
type Student struct {
	Name string
	Grade int
}

func (s Student) changeName1 (name string) { // value receiver dari 'Student')
	fmt.Println("on changeName1, name changed to:", name)
	s.Name = name
}

func (s *Student) changeName2(name string) { // receiver dengan tipe pointer dari 'Student'
	fmt.Println("on changeName2, name changed to:", name)
	s.Name name
}

func main() {
	var s = Student{"Isyana", 21}
	s.changeName1("Anya Geraldine")
	fmt.Println("after changeName:", s.Name)
	s.changeName2("Mimi Perih")
	fmt.Println("after changeName2: ", s.Name)
}
```
```Output
> go run main.go
on changeName1, name changed to: Anya Geraldine
after changeName1: Isyana
on changeName2, name changed to: Mimi Perih
after changeName2: Mimi Perih
```
Ketika menggunakan *method* `changeName1`, maka nilai pada *property* hanya berubah di dalam *method* tersebut. Tetapi tidak merubah nilai pada *property* di variabel `s`. Sehingga *property* `Name` tidak berubah.

Berbeda dengan *method* `changeName2`, karena menggunakan *pointer*, maka nilai pada property objeknya akan berubah. Nilai pada *property* `Name` yang ada di variabel `s` juga ikut berubah.

Kita menyebut `(s *Student)` dengan istilah **pointer receiver**.

Kapan method menggunakan *pointer receiver* dan *value receiver*?
* Jika *method* memerlukan untuk memodifikasi nilai dari *property*, gunakan *method pointer* atau *pointer receiver*.
* Jika *method* memerlukan untuk mengatasi `nil`, gunakan *method pointer* atau *pointer receiver*.
* Jika *method* tidak perlu memodifikasi nilai dari property, bisa menggunakan *method* biasa atau *value receiver*.

## Method Implementation
Kita dapat mengimplementasikan *method* dari sebuah *interface* sehingga dapat mempermudah pengaturan format *output* dari `struct` tersebut (`interface` akan di bahas lebih dalam di materi `interface`).

Salah satunya kita dapat mengimplementasikan *method* dari *interface* `fmt.Stringer` yang mengatur format *output* dari `struct` dengan membuat method `String()` yang mengembalikan tipe data `string`.
```Go
// di internal package 'fmt', file 'print.go'

type Stringer interface {
	String() string
}
```

Contoh, kita dapat mengimplementasikan method `String()` dari interface `fmt.Stringer` pada *struct* `Student` dengan membuat *method* `String()` yang mengembalikan tipe data `string`.
```go
type Student struct {
	Name string
	Grade int
}
func (s Student) String() string { //di sinilah kita mengimplementasikan method String() dari interface 
	return fmt.Sprintf("Hello, my name is %s, grade %d", s.Name, s.Grade)
}

type Person struct {
	Name string
	Age int
}

func main() {
	student := Student{ "Anya Geraldine", 22} // struct 'Student' mengimplementasi method String() dari 
	fmt.Println(student)
	
	person := Person("Isyana", 29} // struct 'Person' tidak mengimplementasikan method String() dari in
	fmt.Println(person)
}
```
```Output
Hello, my name is Anya Geraldine, grade 22
{Isyana 29} # format output hanya berupa value dari tiap property struct
```
Dengan mengimplementasikan *method* `String()` dari *interface* `fmt.Stringer`, kita dapat mengatur format *output* ketika menggunakan `fmt.Println()` atau `fmt.Printf()` dari *struct* tersebut. `fmt.Println()` akan menampilkan *output* dengan format yang ditentukan oleh *method* `String()` yang kita buat, yaitu `Hello, my name is Anya Geraldine, grade 22`.
