  
# Module  

## Initialize project  
Ketika membuat sebuah *project* yang kompleks, kita harus membuat sebuah *module* agar kita tahu di mana *root directory project* kita. Di Golang terdapat *command* yang bisa digunakan, yaitu menggunakan `go mod init`, kemudian diikuti dengan nama *module / project* kita.
```
go mod init <module-name>  
```

Umumnya kita menamai *module* kita sama dengan nama direktori *project* kita, atau sesuai dengan nama *repository* dari *project* kita.

Contoh:  
```
> go mod init my-project  
> go mod init github.com/ruang-guru/my-project  
```

Ketika kita sudah mendefinisikan *module* di project kita, maka kita bisa membuat beberapa `package` atau *directory* (direktori) yang dibutuhkan.

Contoh:  
Kita dapat membuat terlebih dahulu direktori yang akan di jadikan sebagai *root directory* dari project kita. 

Di contoh ini, kita membuat direktori bernama `test-golang`. Kemudian gunakan *terminal* dan menjadikannya *active directory*. Setelah itu tinggal jalankan `go mod init`.
```
> go mod init test-golang  
```

Setelah *command* tersebut dijalankan, maka akan muncul 1 file bernama `go.mod` sebagai penanda *root directory* yang berisi informasi nama *module* dan versi Golang. Di dalam *file* tersebut juga dapat berisi *library* dari *project* lain yang dapat kita *import* sehingga bisa digunakan ke dalam *project* kita.
```
module test-golang  

go 1.18  
```

Karena kita belum menggunakan *library* apapun di dalam *project* kita, maka yang ada di file `go.mod` hanya berisi nama *module* dan versi Golang kita. Dan seperti biasa, kita dapat membuat file `main.go` sebagai *root file* dari project kita.
```
> touch main.go  
```

Maka di dalam direktori *project* kita akan berisi 2 *file*:

```
test-golang  
|_ go.mod  
|_ main.go  
```

## Package in directory  
Ketika membuat *project* yang kompleks, sangat disarankan untuk memisah kode menjadi beberapa *package*, atau direktori **berdasarkan logic nya atau "kesamaan fungsi/behaviour"**.

Karena kita sudah membuat *module*, maka kita bisa membuat *package* atau direktori di dalamnya. dan tidak perlu memjalankan `go mod init` di setiap *package* tersebut.

Contoh, kita ingin membuat *package* bernama `calculator` yang berisi fungsi matematis.
```
test-golang  
|__calculator  
	|__calculator.go  
|__go.mod  
|__main.go  
```

Di *file* `calculator.go`, kita perlu mendeklarasikan nama *package*. Umumnya nama *package* yang dideklarasikan di dalam *file* tersebut sama dengan nama direktorinya.

Berarti contoh di sini akan menjadi *package* `calculator` :
```go
// file 'calculator.go"  

package calculator  

func Add(a, b int) int {  
	return a + b  
}  

func Subtract (a, b int) int {  
	return a - b  
}  

func Multiply(a, b int) int {  
	return a * b  
}  

func Divide(a, b int) int {  
	return a / b  
}  
```

Berikut beberapa tips untuk penamaan *package* di Golang.
* Umumnya menggunakan kata benda.  
* Menggunakan terminologi yang singkat dan jelas.
* Menggunakan lower case. Tidak menggunakan snake case atau camel case.  

## Basic import  
Kita dapat menggunakan beberapa *package* yang sudah dibuat di *project* kita agar dapat digunakan di *root directory* atau *package* lain, kita menyebutnya dengan `import`.

Sebelumnya sudah dibahas tentang *import internal package* dari Golang dengan cara memanggil nama *package* nya menggunakan *keyword* `import`.

Namun, di sini ada sedikit perbedaan. Jika ingin melakukan *import* dari *package* yang kita buat sendiri di project kita, kita harus meng-*import* nya dengan format nama *module* kemudian nama *package* kita.
```go
"<module-name>/<package-name>"  
```

Contoh:  
```go
import "my-project/mypackage"  

import "github.com/ruang-guru/my-project/mypackage"  
```
Setelah itu, kita dapat menggunakan semua *field* yang ada di dalam *package* tersebut.  

*Field* adalah semua *identifier* yang dapat diakses di dalam semua *file* yang berada *package* tersebut (`var`, `const`, `struct`, fungsi, `interface`, dll). Cara mengaksesnya cukup menulis dengan di awali nama *package* tersebut kemudian diikuti oleh nama *field*.

Di contoh kita sebelumnya, kita dapat *mengimport package* `calculator` di *file* `main.go` dengan cara melakukan *import* dan dapat menggunakan semua *field* yang ada.
```go
// file 'main.go'

package main

import (
	"fmt"
	"test-golang/calculator"
) // mengimport package 'calculator' di main

func main() {
	fmt.Println(calculator.Add(1, 2))      // mengakses fungsi 'Add' dari package 'calculator'
	fmt.Println(calculator.Subtract(1, 2)) // mengakses fungsi 'Subtract' dari package 'calculator'
	fmt.Println(calculator.Multiply(1, 2)) // mengakses fungsi 'Multiply' dari package 'calculator'
	fmt.Println(calculator.Divide(1, 2))   // mengakses fungsi 'Divide' dari package 'calculator'
} 
```
Output: 
```
3  
-1  
2  
0.5  
```

Contoh lain :  
- Terdapat *file* `person.go` di *package* `person`
```
person  
|__person.go  
main.go  
go.mod  
```
```go
// file 'person.go'

package person

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() string {
	return "Helo," + p.Name
}
```
```go
// file 'main.go'  
package main  

import "test-golang/person" // mengimport package 'person' di main  

func main() {  
	p := person.Person{ // membuat objek 'p' dari struct 'Person' di package 'person' 
		Name: "John",  
		Age: 30,  
	} 
	
	fmt.Println(p.SayHello()) // mengakses fungsi 'SayHello' dari package 'person'  
}
```
```
Helo, John  
```

Jika terdapat nama *package* yang sama, kita dapat memberikan alias pada salah satu package untuk menghindari terjadinya "*conflict*".
```go
import (  
	"crypto/rand"  
	mrand "math/rand"  
}
```

Sebagai contoh, terdapat dua nama *package* `rand` yang sama yaitu dari `crypto/rand` dan `math/rand`. Untuk menghindari "*conflict*" karena nama *package* yang sama, kita dapat memberi alias `mrand` pada *package* `math/rand`.  

Ketika ingin menggunakan *package* `math/rand`, kita bisa memanggilnya dengan nama alias.
```go
import (
	"crypto/rand" // dipanggil dengan cara di awali dengan 'rand'
	mrand "math/rand" // dipanggil dengan alias 'mrand'
)

func main() {
	cryptoRand, err := rand.Read(make([]byte, 10)) // mengakses fungsi 'Read' dari package 'crypto/rand
	if err != nil {
		panic(err)
	}

	fmt.Println(cryptoRand)

	mrand.Seed(10) // menggunakan fungsi 'Seed' dari package 'math/rand' dengan alias 'mrand'
	mathRand := mrand.Int()

	fmt.Println(mathRand)
}
```
Output:
```
10  
5221277731205826435  
```

## Public fields  
Di Golang terdapat *rules* yang unik. Kita dapat meng-*export* semua *field* di *package* tersebut dengan nama yang di awali dengan huruf besar / kapital (*public field*).

Contoh, anggap kita memiliki sebuah *package* bernama `public` yang berisi sebuah file `public.go`.
```
public  
|__public.go
main.go  
go.mod
```

* file public.go  
```go
package public

// semua field yang ada di dalam package ini adalah public, karena di awali dengan huruf besar / kapita
var Name = "Budi Setiawan"
const Pi = 3.14
  
type Person struct {
Name string
Age int
}
  
func Hello(name string) string {
return "Hello" + name
}  
```

* *file* `main.go`  

```go
package main  
import ( 
	"test-golang/public"  
)

func main() {  
	fmt.Println(public.Name)  
	fmt.Println(public.Pi)  
	fmt.Println(public. Person(Name: "Budi", Age: 20})  
	fmt.Println(public. Hello ("Budi"))  
}
```
Output
```
Budi Setiawan  
3.14  
{Budi 20)  
Hello Budi
```
Kita dapat menggunakan semua *field* yang ada di *package* `public`, karena nama *field*-nya di awali dengan huruf besar / kapital.  

## Private fields  
Sebaliknya, kita dapat membuat *field* yang ada di dalam *package* tersebut menjadi *private* (tidak dapat diakses) dengan menggunakan nama yang di awali dengan huruf kecil / lower case.

* *file* `public.go`
```go
package private  
// semua field yang ada di dalam package ini adalah private, karena di awali dengan huruf kecil / lower  
type student struct {  
	name string  
	batch int  
}  
func hello(name string) string {  
	return "Hello" + name  
}  
```
Konsep ini bisa digunakan untuk menghindari kesalahan atau penggunaan *field* yang tidak diinginkan.  

Bisa juga untuk meng-enkapsulasi *field* yang hanya dapat diakses sesuai aturan tertentu.
```go
package private

type student struct { // private field
name string
batch int
}
  
func NewStudent (name string, batch int) *student {
return &student{
name: name,
batch: batch,
}
}
func (s *student) GetName() string {
return "Name: " + s.name
}
```

```go
// file 'main.go  
package main  

import "test-golang/private"  

func main() {  
	s := private.NewStudent("Budi", 20) // menggunakan fungsi 'NewStudent' dan mendapatkan struct 'student' 
	fmt.Println(s.GetName()) // menggunakan fungsi 'GetName' untuk mengakses property 'name' di struct 'student'
}
```
Output:
```
Name: Budi  
```

## Import cycle error  
Kita dapat menemukan *error* ketika lebih dari satu *package* saling bergantung satu sama lain. Istilah ini juga disebut dengan "*Circular dependency*".

Di Golang, *error* ini terjadi ketika kita mengimport *package* yang bergantung satu sama lain.

Sebagai contoh, terdapat dua *package* yaitu `person` dan `car`. Di *package* `person` terdapat *struct* `Person`. Di *struct* ini terdapat atribut `Cars` yang memiliki tipe *slice* dari objek `Car` dari *package* `car`.

```
car  
|_car.go  
person  
|_person.go  
main.go  
go.mod  
```
```go
// di file 'person.go  
package person  

import "test-golang/car"  

type Person struct {  
	Name string  
	Cars []car.Car  
}  
```

Di sisi lain, pada *package* car terdapat *struct* Car yang memiliki atribut Owner. Atribut ini memiliki tipe  Person dari *package* person.
```go
package car  

import "test-golang/person"  

type Car struct {  
	Model string  
	Owner person.Person  
}
```

Dari contoh tersebut, antara `Person` dan `Car` memiliki dependensi satu sama lain. Hal ini merupakan **circular dependency**.

Ketika terdapat *circular dependency*, Golang akan mengeluarkan *error* `import cycle not allowed`. *Error* ini berupa *compile error*, sehingga bisa langsung muncul ketika kita membuatnya di *code editor*.

## Import cycle error solution  
Ketika mendapatkan *error circular dependency*, kita harus memperbaikinya agar program kita dapat berjalan. Caranya adalah menghapus salah satu *field* yang bergantung satu sama lain.

Jika dilihat dari contoh sebelumnya, kita dapat menghapus salah satu *property*, yaitu menghapus *property* `Cars `di *struct* `Person`, atau menghapus *property* `Owner` di *struct* `Car`.
```go
// di file 'car.go'  
type Car struct { // menghapus property 'Cars'  
	Model string  
}  

// di file 'person.go'  
package person  
import "test-golang/car"  
type Person struct {  
	Name string  
	Cars []car.Car  
}  
```

Atau jika ingin tetap mempertahankan atributnya, kita dapat mengatasinya dengan menggabungkan kedua *struct* dalam satu *package*. Kita dapat membuat *package* baru bernama `entity` yang berisi *struct* `Person` dan Car di dalam file terpisah. 
```
entity  
|_car.go  
|_person.go  
main.go  
go.mod  
```
```go
// di file 'car.go'  
package entity  

type Car struct {  
	Model string  
	Owner *Person  
}  
```
```go
// di file 'person.go'  
package entity  
type Person struct {  
	Name string  
	Cars []Car  
}  
```
