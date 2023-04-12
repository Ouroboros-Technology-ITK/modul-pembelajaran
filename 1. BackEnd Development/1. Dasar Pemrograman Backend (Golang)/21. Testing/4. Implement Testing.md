# Implement testing


## Multiple data testing
Kadang dalam membuat *unit testing*, kita perlu melakukan *testing* dengan data yang berbeda-beda untuk memastikan kode tersebut tidak ada kecacatan. Data yang perlu di *test* adalah data dengan karakteristik yang berbeda-beda.

Contoh, terdapat fungsi sederhana dengan menjumlahkan dua buah bilangan.
```go
package main

func SumTwoNumber(a int, b int) int {
	return a + b
}
```

Kita perlu mempertimbangkan beberapa data yang harus digunakan sebagai *input test*. Cara paling sederhana adalah memberikan 2 jenis bilangan yaitu negatif dan positif. Jika terdapat 2 parameter bilangan, maka akan muncul 4 *test*.

* Kedua parameter berisi bilangan positif
* Parameter pertama bilangan negatif, parameter kedua bilangan positif
* Parameter pertama bilangan positif, parameter kedua bilangan negatif
* Kedua parameter berisi bilangan negatif

Untuk membuat beberapa *test*, kita dapat membuat lebih dari 1 *test case* menggunakan perulangan sederhana dan membuat data yang perlu disiapkan terlebih dahulu.

```go
package main

import "testing"

type TestCase struct { // membuat struct untuk menyimpan data testing
	a        int
	b        int
	expected int
}

func TestSumTwoNumber(t *testing.T) {
	var listTest = []TestCase{ // menyiapkan semua data test case
		{a: 1, b: 2, expected: 3},
		{a: -2, b: 4, expected: 2},
		{a: 2, b: -7, expexted: -5},
		{a: -4, b: -5, expected: -9},
	}

	for _, test := range listTest { // melakukan perulangan untuk menjalankan semua test case
		result = SumTwoNumber(test.a, test.b)
		if result != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, result)
		}
	}
}
```

Kita dapat menjalankan *test* seperti biasa, menggunakan `go test`.
```
> go test -v
=== RUN   TestSumTwoNumber
--- PASS: TestSumTwoNumber (0.00s)
PASS 
ok      test-golang 0.0075
```

Jika fungsi `SumTwoNumber` adalah fungsi yang sesuai, maka harusnya semua *test case* akan `PASS`.

## Setup and Clear
Kita dapat melakukan *setup* sebelum menjalankan *test case*. *Setup* ini adalah untuk melakukan inisialisasi sebelum menjalankan *test case*. Misalnya, kita ingin membuat sebuah *database* terlebih dahulu, atau perlu mempersiapkan data *struct* dan menjalankan beberapa *method* terlebih dahulu.

Golang tidak memiliki *keyword* untuk melakukan *setup* sebelum menjalankan *test case*. Namun, kita dapat membuat sebuah fungsi `SetupTest` yang diisi untuk kebutuhan setup. Fungsi ini dapat dibuat dengan 1 parameter yaitu `t *testing.T`. Fungsi ini akan dijalankan sebelum menjalankan *test case*.
```go
func SetupTest(t *testing.T) {
// melakukan setup
}
```

Contoh, kita ingin melakukan *test case* di `main.go` dengan kode sebagai berikut:
```go
type Person struct {
	Name string
	Age int
	BaseSalary float64
	Wealth float64
}

func NewPerson(name string, age int, baseSalary float64) *Person {
	return &Person{
		Name: name,
		Age: age,
		BaseSalary: baseSalary,
	}
}

func (p *Person) GetWealth() float64 {
	return p.Wealth
}

func (p *Person) SetBonus (monthOfWork int) {
	p.Wealth += (float64(monthOfWork) / 12) * p.BaseSalary * 0.5
}

func (p *Person) AddIncome (income float64) {
	p.Wealth income
}

func (p *Person) AddCost (cost float64) {
	p.Wealth cost
}

func (p *Person) Clearwealth() {
	p.Wealth = 0
}
```
Dengan membuat setup, maka kita dapat mempersiapkan data yang di perlukan terlebih dahulu tanpa harus menyiapkan data di setiap test case.

```go
// di main_test.go
var person *Person = NewPerson ("John Doe", 20, 100_000)
// setup test untuk mempersiapkan data struct 'Person'
func SetupTest(t *testing.T) {
t.Log("Setup test case")
// menjalankan SetBonus terlebih dahulu sebelum menjalankan test case person.SetBonus (12)
}
```
Kita tinggal membuat *test case* untuk semua *method* yang dibuat.
```go
var person *Person = NewPerson("John", 20, 100000)

func SetupTest(t *testing.T) {
	t.Log("Setup test case")

	person.SetBonus(12)
}
func TestAddIncome(t *testing.T) {
	// menjalankan setup test
	SetupTest(t)
	
	person.AddIncome(100000)
	if person.GetWealth() != 150000.0 {
		t.Errorf("Expected %f, but got %f", 150000.0, person.GetWealth())
	}
}
func TestAddCost(t *testing.T) {
	// menjalankan setup test
	SetupTest(t)
	
	person.AddCost(10000)
	if person.GetWealth() != 190000.0 {
		t.Errorf("Expected %f, but got %f", 190000.0, person.GetWealth())
	}
}
```
Contoh di atas, kita dapat melihat bahwa kita tidak perlu menjalankan `SetBonus()` di setiap *test case*. Kita hanya perlu membuat data `person`, kemudian menjalankan `SetBonus()` di fungsi `SetupTest` dan menjalankan fungsi `SetupTest` di setiap *test case*.
```
> go test -v
=== RUN TestAddIncome
	main_test.go:8: Setup test case
--- PASS: TestAdd Income (0.005)
=== RUN TestAddCost
	main_test.go:8: Setup test case
--- PASS: TestAddCost (0.00s) 
PASS
ok test-golang 0.385s
```
Kita juga dapat melakukan *clear* atau *teardown* setelah menjalankan *test case*. Misalnya, kita ingin menghapus data *database* yang telah kita buat, atau me-*reset* data yang dibutuhkan. Hal ini menjadi penting saat *integration testing*.

Sama seperti *setup*, *package* `testing` tidak menyediakan *keyword* untuk melakukan *clear* atau *teardown*. Namun, kita dapat memodifikasi fungsi ``SetupTest`` dengan menambah *return* berupa fungsi yang berisi 1 parameter `*testing.T`. Fungsi ini akan dijalankan setelah menjalankan *test case*.
```go
func SetupTest(t *testing.T) func(*testing.T) {
	// melakukan setup
	return func(t *testing.T) {
		// melakukan clear
	}
}
```

Contoh, kita ingin melakukan *clear* setelah menjalankan *test case* dengan memanggil *method* `ClearWealth()`.
```go
var person *Person = NewPerson("John", 20, 100000)
func SetupTest(t *testing.T) func(*testing.T) {
	t.Log("Setup test case")
	person.SetBonus (12)
	return func(t *testing.T) {
		t.Log("Clear test case")
		// melakukan clear dengan memanggil method 'ClearWealth()' 
		person.ClearWealth()
	}
}
```

Sisanya, kita dapat melakukan *clear* setelah menjalankan *test case* dengan menampung fungsi `SetupTest` di *variable* dan memanggil fungsi tersebut setelah menjalankan *test case* menggunakan `defer`.

```go
var person *Person = NewPerson("John", 20, 100000)

func SetupTest(t *testing.T) func(*testing.T) {
	t.Log("Setup test case")
	person.SetBonus(12)
	return func(t *testing.T) {
		t.Log("Clear data test case set wealth to 0")
		person.ClearWealth()
	}
}

func TestAddIncome(t *testing.T) {
	// menampung fungsi 'SetupTest' di variable 'clear'
	clear := SetupTest(t)

	// melakukan clear setelah menjalankan test case
	// kita dapat memastikan bahwa clear ini dijalankan setelah test case menggunakan 'defer'
	defer clear(t)

	person.AddIncome(100000)
	if person.GetWealth() != 150000.0 {
		t.Errorf("Expected %f, but got %f", 150000.0, person.GetWealth())
	}
}

func TestAddCost(t *testing.T) {
	clear := SetupTest(t)

	defer clear(t)

	person.AddCost(10000)
	if person.GetWealth() != 40000.0 {
		t.Errorf("Expected %f, but got %f", 40000.0, person.GetWealth())
	}
}
```

Output:
```go
> go test -v
=== RUN	  TestAddIncome
	main_test.go:8: Setup test case
	main_test.go:13: Clear data test case set wealth to e
--- PASS: TestAddIncome (0.00s)
=== RUN   TestAddCost
	main_test.go:8: Setup test case
	main_test.go:13: Clear data test case set wealth to e
--- PASS: TestAddCost (0.00s)
PASS
ok      test-golang      0.453s
```

Contoh di atas adalah contoh sederhana untuk melakukan *setup* dan *clear* di *test case*. Sebelum menjalankan *test case*, kita memanggil fungsi `SetupTest` untuk menjalankan *SetBonus()*. Namun, setelah menjalankan *test case*, kita memanggil fungsi `clear` yang menjalankan `ClearWealth()` untuk mengatur ulang data di properti `Wealth` menjadi `0`.
