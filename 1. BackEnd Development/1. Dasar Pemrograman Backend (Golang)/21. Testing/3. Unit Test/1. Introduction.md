# Introduction  

## Unit testing in Go  
Golang menyediakan package testing yang berisi banyak sekali tools untuk keperluan unit testing. Dalam membuat unit testing di Golang, tedapat aturan yang harus diikuti. Pertama, kita harus membuat sebuah file dengan nama akhiran `_test.go` dan *package* nya harus sama. Jika *file* yang ingin kita tes bernama `main.go`, maka *file testing* harus bernama `main_test.go`.

Unit *testing* di Golang ditulis dalam bentuk fungsi dengan nama awalan `Test` diikuti dengan nama fungsi yang ingin kita tes. Fungsi di dalam unit *testing* harus memiliki parameter `*testing.T`. Fungsi ini akan dipanggil oleh package testing ketika kita menjalankan perintah `go test`.

Contoh, kita memiliki sebuah kode di `main.go` sebagai berikut:
```go
package main  

import "math"  

type Cube struct {  
	Side float64  
}  
func (c Cube) Area() float64 {  
	return 6 * math.Pow(c.Side, 2)  
}  

func (c Cube) Circumference() float64 {  
	return 12 * c.Side
}  
func (c Cube) Volume() float64 {  
	return math.Pow(c.Side, 3)
}
```

Terdapat struct `Cube` yang memiliki 3 *method* yaitu `Area`, `Circumference`, dan `Volume`. Untuk melakukan *unit testing* kita perlu membuat *file* `main_test.go` dengan *code testing* sebagai berikut:
* Melakukan *testing* untuk *method* `Area`  
```
test-golang  
|_main.go  
|_main_test.go  
```
```go
package main

import "testing"

func TestArea(t *testing.T) { // fungsi ini digunakan untuk mengetes method Area di struct 'Cube'
	var cube = Cube{Side: 2}    // menyiapkan data untuk diuji
	var expected float64 = 24.0 // nilai yang diharapkan

	if cube.Area() != expected { // jika hasil tidak sesuai dengan yang diharapkan, maka test dianggap
		t.Errorf("Expected %f, but got %f", expected, cube.Area())
	}
}
```
Contoh di atas adalah bentuk *unit test* untuk *method* `Area()` dari *struct* `Cube` yang ada di *package* `main`. Karena yang kita tes adalah *method* dari sebuah *struct*, maka kita perlu menyiapkan data terlebih dahulu dengan mengisi properti `Side` bernilai `2`. Kita juga dapat membuat *variable* `expected` yang berisi nilai yang diharapkan. Cara pengujiannya, dengan membuat sebuah kondisi yang membandingkan antara kembalian nilai dari *method* `Area()` dan nilai di *variable* `expected`. Jika hasilnya sama, maka dianggap `PASS`, jika tidak maka dianggap pengujian tersebut `FAIL`.

Untuk menjalankan tes tersebut, kita cukup menjalankan *command* `go test` di *terminal/console*. Pastikan kita berada di *active directory* dan sama dengan *file testing*.
```
> go test  
PASS ok  
test-golang 0.919s  
```

Ketika hasil yang dimunculkan di terminal adalah `PASS`, maka testing dianggap berhasil atau sesuai ekspektasi.  

Usahakan ketika membuat *unit testing*, semua kode yang dibuat harus di-*test*. Kita dapat melengkapi *unit testing* untuk *method* lain, sebagai berikut:
```go
package main

import "testing"

func TestArea(t *testing.T) {
	var cube = Cube{Side: 2}
	var expected float64 = 24.0

	if cube.Area() != expected {
		t.Errorf("Expected %f, but got %f", expected, cube.Area())
	}
}

func TestCircumference(t *testing.T) {
	var cube = Cube{Side: 2}
	var expected float64 = 24.0

	if cube.Circumference() != expected {
		t.Errorf("Expected %f, but got %f", expected, cube.Circumference())
	}
}

func TestVolume(t testing.T) {
	var cube = Cube{Side: 2}
	var expected float64 = 8.0

	if cube.Volume() != expected {
		t.Errorf("Expected %f, but got %f", expected, cube.Volume())
	}
}
```
terminal:
```terminal
> go test  
PASS  
ok test-golang 0.707s  
```

Kita dapat melihat secara detail informasi *testing* di *terminal* / *console* dengan menambah opsi `-v`.

```
go test -v  
---RUN  
TestArea  
--- PASS: TestArea (0.00s)  
=== RUN TestCircumference  
- PASS: TestCircumference (0.005)  
== RUN TestVolume  
---PASS: TestVolume (0.005)  
PASS 
ok  test-golang 0.8265  
```

Informasi yang didapat akan lebih rinci, dengan memberikan informasi fungsi *testing* apa saja yang dijalankan terlebih dahulu dan informasi hasil *testing* tersebut, apakah `PASS` atau `FAIL`. Karena kita membuat 3 fungsi *testing* yang menguji *method* `Area()`,` Curcumference()` dan `Volume()`, maka informasi yang didapat yaitu berupa menjalankan 3 fungsi testing dan 3 informasi `PASS`. Berarti semua fungsi yang dibuat sudah berhasil dan sesuai ekspektasi.

Di Golang, ketika kita menjalankan *command* `go test` maka semua *testing* yang ada di *package* yang sedang kita buat akan dijalankan.

Contoh, terdapat *package* `entity` yang berisi *struct* `person` dan `student`.  
```
entity  
|_ person.go  
|_ student.go  
```

```go
// file person.go  
package entity  

type Person struct {  
	Name string  
	DateBirth = time.Time  
}  
func (p Person) GetAge() int {  
	return time.Now().Year() - p.DateBirth.Year()  
}  
```
```go
// file student.go  
package entity

type Student struct {
	Person
	Batch string
	Grade int
}

func (s Student) GetPredicate() string {
	switch {
	case s.Grade == 100:
		return s.Name + "got Predicate: Perfect"
	case s.Grade < 100 && s.Grade >= 90:
		return s.Name + "got Predicate: Excellent"
	case s.Grade < 90 && s.Grade >= 80:
		return s.Name + "got Predicate: Good"
	case s.Grade < 80 && s.Grade >= 70:
		return s.Name + " got Predicate: Acceptable"
	default:
		return s.Name + " got Predicate: Bad"
	}
}
```
Kemudian kita dapat membuat *unit testing* di *file* `person_test.go` dan `student_test.go` sebagai  berikut:  
```
entity  
|_ person.go _person_test.go  
|_ student.go  
|_ student_test.go  
```
```go
// file person_test.go  
package entity

import (
	"testing"
	"time"
)

func TestGetAge(t testing.T) {
	var person = Person{
		Name:      "John Doe",
		DateBirth: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	var expected int = 42

	if person.GetAge() != expected {
		t.Errorf("Expected %d, but got %d", expected, person.GetAge())
	}
}
```
```go
// file student_test.go  
package entity

import (
	"testing"
	"time"
)

func TestGetPredicate(t *testing.T) {
	var student = Student{
		Person: Person{
			Name:      "John Doe",
			DateBirth: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Batch: "Batch 1",
		Grade: 90,
	}

	var expected string = "John Doe got Predicate: Excellent"
	if student.GetPredicate() != expected {
		t.Errorf("Expected %s, but got %s", expected, student.GetPredicate())
	}
}
```

Ketika kita coba menjalankan `go test` di *active directory* `entity`, maka akan menjalankan semua  *testing* yang ada di *package* `entity`.
```
> go test -v  
--- RUN TestGetAge  
--- PASS: TestGetAge (0.005)  
=== RUN TestGetPredicate  
--- PASS: TestGetPredicate (0.005)  
PASS  
ok 	test-golang/entity	0.0925  
```

## Specific Testing
Jika kita memiliki banyak *file testing* di dalam 1 *package*, kita dapat melakukan pengujian ke *file* secara spesifik dengan tambahan *command* berupa *file* apa yang mau di *test*.
```
go test [file] [file_testing] -v  
```

Contoh, ketika kita hanya ingin menguji semua kode yang ada di `main.go`, maka kita dapat menjalankan `command` sebagai berikut:  

```
> go test main.go main_test.go -v  
=== RUN   TestArea  
--- PASS: TestArea (0.005)  
=== RUN   TestCircumference  
--- PASS: TestCircumference (0.005)  
=== RUN   TestVolume  
--- PASS: TestVolume (0.005)  
PASS  
ok    command-line-arguments   0.434s  
```
Hasilnya tidak jauh beda dengan menjalankan *testing* biasa, namun yang di tes hanya spesifik di *file* `main.go` dengan *testing* berdasarkan *file* ``main_test.go``.

Di Golang, juga dapat melakukan *testing* lebih spesifik lagi, yaitu hanya menjalankan 1 fungsi *testing*. Caranya dengan menambah opsi `-run` dan diikuti dengan nama fungsinya.
```
go test run [function_name] -v  
```

Contoh, kita hanya ingin menjalankan test `TestArea` saja, maka kita dapat menjalankan *command* sebagai berikut:
```
> go test-run TestArea -v  
=== RUN TestArea  
--- PASS: TestArea (0.005)  
PASS  
ok    command-line-arguments   0.092s  
```
Hasilnya akan menjalankan 1 fungsi tes yaitu `TestArea` saja dengan hasil `PASS`.

## Method testing

Berikut beberapa method yang disediakan oleh package testing:  

METHOD|  DESCRIPTION
--|--
Log()|  Menampilkan log  
Logf()  |Menampilkan log dengan format  
Fail()  |Digunakan untuk menandakan proses testing gagal, tetapi proses testing akan tetap berlanjut
FailNow() |Digunakan untuk menandakan proses testing gagal, dan proses testing langsung berhenti 
Failed()  |Menampilkan laporan fail / gagal
SkipNow()|Melewati proses testing fungsi tertentu, melanjutkan ke testing fungsi setelahnya  
Skipped() |Menampilkan laporan skip  
Error() |Memanggil Log() diikuti dengan Fail()  
Errorf() |Memanggil Logf() diikuti dengan Fail() Fatal() Memanggil Log() diikuti dengan FailNow()  
Fatalf()  |Memanggil Logf() diikuti dengan FailNow()  
Skip()  |Memanggil Log() diikuti dengan skipNow()  
Skipf()  |Memanggil Logf() diikuti dengan SkipNow()  
Contoh penggunaan salah satu *method* di atas:  
```go
func TestArea(t *testing.T) {
	var cube = Cube{Side: 2}
	var expected float64 = 24.0

	t.Logf("Testing Area: %f", cube.Area()) // menampilkan log dengan format string
	if cube.Area() == 0 {
		t.Skip("Skip this test") // melewati proses testing
	}

	if cube.Area() != expected {
		t.Errorf("Expected %f, but got %f", expected, cube.Area()) // menampilkan Logf diikuti dengan F
	}
}
```

```
>go test -v  
=== RUN   TestArea  
	main_test.go:8: Testing Area: 24.000000  
--- PASS: TestArea (0.00s)  
PASS 
ok      test-golang   0.4375  
```

Terdapat *method* yang dapat digunakan untuk memberikan informasi dengan menggunakan `Run()`. *Method* ini memiliki 2 parameter, yaitu `string` dan fungsi. Fungsi di parameter ke 2 adalah fungsi yang akan dijalankan saat *testing* dengan 1 parameter berupa `*testing.T`.
```go
// ...  
t.Run("[Information]", func(t *testing.T) {  
	// code for testing  
})  
// ...  
```
Kita dapat menggunakan *method* ini untuk memisah beberapa *testing* jika dalam satu unit fungsi tersebut memiliki beberapa *behaviour*.

Contoh terdapat kode seperti berikut:
```go
package main
func SumPositiveNumber(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a and b must be positive")
	}

	return a + b, nil
}
```
Contoh di atas adalah fungsi yang memiliki 2 *behaviour*, yaitu penjumlahan 2 angka dan mengembalikan *error* jika salah satu dari 2 angka tersebut negatif.

Maka, kita dapat membuat *testing* sebagai berikut:
```go
package main  
import (
	"testing"
)

func TestSumPositiveNumber(t *testing.T) {
	t.Run("Return error when a or b is negative", func(t testing.T) {
		_, err := SumPositiveNumber(-1, 1)
		if err == nil {
			t.Error("Expected error, but got nil")
		}
	})
	t.Run("a and b is positive", func(t testing.T) {
		result, err := SumPositiveNumber(1, 1)
		if err != nil {
			t.Error("Expected nil, but got error")
		}

		if result != 2 {
			t.Errorf("Expected 2, but got %d", result)
		}
	})
}
```

```
> go test -v  
=== RUN   TestSumPositiveNumber  
=== RUN   TestSumPositive Number/Return_error_when_a_or_b_is_negative  
=== RUN   TestSumPositiveNumber/a_and_b_is_positive  
--- PASS: TestSumPositiveNumber (0.005)  
    --- PASS: TestSumPositiveNumber/Return_error_when_a_or_b_is_negative (0.00s)  
    --- PASS: TestSumPositiveNumber/a_and_b_is_positive (0.005)  
PASS  
ok      test-golang     0.4175  
```
Dengan menggunakan method `Run()`, kita bisa menguji *behaviour* yang berbeda-beda dari satu unit  fungsi.
