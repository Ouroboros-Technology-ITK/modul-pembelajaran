# Ginkgo


## What is Ginkgo?
Ginkgo adalah *framework* (kerangka kerja) yang dibuat untuk membuat testing yang mudah dibaca dan ditulis di bahasa pemrograman Golang. Ginkgo dibuat berdasarkan konsep "*Behaviour Driven Development*" ([BDD]()) dan dapat digunakan untuk segala jenis *testing*, seperti *unit testing*, *integration testing*, *system testing*, *acceptance testing*, dll.

Ginkgo membuat kita dapat fokus pada *testing* itu sendiri, bukan bagaimana kita membuat *testing*. Ginkgo juga dapat digunakan untuk membuat *testing* yang dapat dijalankan secara paralel dan berulang-ulang, sehingga kita dapat menghemat waktu ketika menjalankan *testing* dan dapat memastikan *testing* yang dibuat benar-benar berjalan dengan baik.

Berikut referensi *source code* di Github terkait [Ginkgo]() dan juga dokumentasi [Ginkgo]().


## Installation and Bootstrapping
*Bootstrapping* adalah proses untuk melakukan inisialisasi *testing* yang dapat dijalankan di project kita menggunakan Ginkgo. Karena Ginkgo adalah sebuah *module*, maka kita perlu melakukan inisiasi *project* terlebih dahulu dan memastikan sudah terdapat file `go.mod`.

Contoh, kita ingin membuat sebuah *project* kalkulator sederhana yang proses kerjanya hampir sama dengan kalkulator pada umumnya. Kalkulator yang kita buat saat ini hanya dapat melakukan 4 operasi dasar, yaitu penjumlahan, pengurangan, perkalian, dan pembagian.

Pastikan kita sudah melakukan `go mod init` di *project* kita. Kemudian kita membuat *package* `calculator` dengan berisi *file* `calculator.go`.
```
my-project/
|_calculator
	|_calculator.go
|_go.mod
|_go.sum
```

Setelah itu, kita perlu melakukan instalasi Ginkgo di *project* kita dengan menjalankan 3 *command* berikut.

```
> go install -mod-mod github.com/onsi/ginkgo/v2/ginkgo
> go get -u github.com/onsi/ginkgo/v2
> go get -u github.com/onsi/gomega/...
```
Command di atas digunakan untuk melakukan instalasi Ginkgo dan memasang *module* Ginkgo ke *project* kita.

Kita dapat memastikan bahwa Ginkgo sudah terinstall dengan menjalankan *command* `ginkgo version`.
```
> ginkgo version
Ginkgo Version 2.1.4
```

Jika instalasi sudah berhasil, kita dapat mengisi kode di *package* `calculator`, *file* `calculator.go` seperti berikut:
```go
// You can edit this code!
// Click here and start typing.
package calculator

type Calculator struct{
	Base float64
}

func InitCalculator (base float64) *Calculator { // inisiasi kalkulator
	return &Calculator{
		Base: base,
	}
}

func (c *Calculator) Add(number float64) { // method untuk penjumlahan
	c.Base += number
}

func (c *Calculator) Subtract(number float64) { // method untuk pengurangan
	c.Base -= number
}

func (c *Calculator) Multiply(number float64) { // method untuk perkalian 
	c.Base *= number
}

func (c *Calculator) Divide(number float64) error { // method untuk pembagian 
	if number == 0 {
		return errors.New("cannot divide by zero")
	}

	c.Base /= number
	return nil
}

func (c *Calculator) Result() float64 { // method untuk menampilkan hasil
	return c.Base
}

func (c *Calculator) Reset() { // method untuk mereset hasil
	c.Base = 0
}
```
Saatnya kita melakukan *bootstrapping* untuk membuat *testing* yang dapat dijalankan di *project* kita. Untuk melakukan *bootstrapping*, kita perlu menjalankan *command* `ginkgo bootstrap` di dalam *package* `calculator`.
```
> cd calculator
> ginkgo bootstrap
Generating ginkgo test suite bootstrap for main in:
	calculator_suite_test.go
```
Maka akan muncul *file* `calculator_suite_test.go` yang berisi *code unit testing* dari Golang.
```
my-project/
|_calculator
|	|_calculator_suite_test.go
|	|_calculator.go
|_go.mod
|_go.sum
```
```go
package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCalculator (t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}
```

Setelah berhasil melakukan *bootstrapping*, kita dapat membuat jenis *testing* yang dibutuhkan untuk *project* kita. Karena kita ingin membuat *unit test*, maka jalankan *command* `ginkgo generate` dan diikuti nama file yang ingin di *test* di dalam direktori `calculator`. Karena *file* yang ingin kita *test* adalah `calculator.go`, maka kita akan menjalankan *command* `ginkgo generate calculator`.
```
> ginkgo generate main
Generating ginkgo test for Main in:
	calculator_test.go
```
Maka akan muncul *file* `calculator_test.go` yang menjadi *base* pembuatan *code unit test* untuk *file* `calculator.go`.

```
my-project
|_calculator
|	|_calculator_suite_test.go
|	|_calculator_test.go 
|	|_calculator.go
|_go.mod
|_go.sum
```
Isi dari *file* `calculator_test.go`:
```go
package main_test
import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"calculator"
)

var _ = Describe("Calculator", func() {
})
```
Setelah berhasil membuat *file* `calculator_test.go`, kita dapat menulis *code unit testing* yang dibutuhkan untuk project kita.

## Unit test with ginkgo
#### Introduction
Ginkgo menggunakan bahasa yang lebih *human readable* dan *natural* untuk menulis *testing*. Terdapat macam-macam blok testing yang dapat digunakan di Ginkgo.

Istilah *spec* digunakan di Ginkgo untuk menjelaskan hasil akhir dari kode (*behaviour of your code*) secara terstruktur.
1. Blok penampungan (*container block*) Blok ini digunakan untuk menyusun *spec* yang akan kita tulis secara bertingkat. Kita dapat menggunakan 3 blok penampung, yaitu `Describe`, `Context`, dan `When`.
   * `Describe` digunakan untuk menampung *spec* yang berhubungan dengan subject yang sama.
   * `Context` digunakan untuk menampung *spec* yang berhubungan dengan kondisi yang sama.
   * `When` digunakan untuk menampung *spec* yang berhubungan dengan action yang sama.

2. Blok pengaturan (*setup block*) Blok ini digunakan untuk melakukan pengaturan sebelum dan/atau setelah menjalankan *spec* yang akan kita tulis. Kita dapat menggunakan 2 blok pengaturan, yaitu `BeforeEach` dan `AfterEach`.
   * BeforeEach digunakan untuk melakukan pengaturan sebelum menjalankan spec.
   * AfterEach digunakan untuk melakukan pengaturan setelah menjalankan spec.

3. Blok pengujian (*assertion block*) Blok ini digunakan untuk melakukan pengujian terhadap *subject* yang akan kita tulis. Kita dapat menggunakan 2 blok pengujian, yaitu `It` atau `Specify`.

Cara menulis blok dengan Ginkgo yaitu cukup menulis nama blok dengan diisi 2 parameter, yaitu *spec* dan fungsi yang akan dijalankan.
```go
var _ = Describe("Calculator", func() {
	// bisa diisi dengan blok atau pengujian
})
```

Gingko menggunakan *method* dari Gomega untuk melakukan pengujian terhadap *subject*. Kita dapat menggunakan 2 macam *method*, yaitu `Expect` dan `Eventually`.

Cara menggunakannya adalah sebagai berikut:
```go
var _ = Describe("Core", func() {
	Describe("Example context to test core", func() {
		It("Should have an assertion", func() {
			// pengujian menggunakan 'Expect' dari Gomega
			Expect(2 + 3).To(Equal(5))
		})
	})
})
```

Terdapat banyak *method* pengujian yang dapat digunakan.
* `ToNot` digunakan untuk mengecek apakah hasil tidak sesuai dengan *expected*.
* `Equal` digunakan untuk melakukan pengujian apakah hasil sama dengan *expected*.
* `BeNil` digunakan untuk melakukan pengujian apakah hasil yang didapat bernilai `nil`.
* `BeTrue` digunakan untuk melakukan pengujian apakah hasil bernilai `true`
* `BeFalse` digunakan untuk melakukan pengujian apakah hasil bernilai `false`
* `ContainSubstring` digunakan untuk melakukan pengujian apakah hasil yang didapat memiliki `substring` dari *expected* (ini untuk kasus hasil berupa tipe data `string`).
* `ContainElement` digunakan untuk melakukan pengujian apakah hasil yang didapat berada di dalam list *expected* (biasanya dalam bentuk *array* / *slice*).
Dan masih banyak lagi, kita dapat melihat di [dokumentasi Gomega]().

Berikut contohnya:
```go
var _ = Describe("Core", func() {
	Describe("Example context to test core", func() {
		It("Should have an assertion", func() {
			// pengujian dengan 'Equal'
			Expect(2 + 3).To(Equal(5))

			// pengujian dengan 'ToNot'
			Expect(2 + 3).ToNot(Equal(4))

			// pengujian dengan 'Not' dan 'Equal'
			Expect(2 + 3).To(Not(Equal(4)))

			// pengujian dengan 'BeNumerically' yang di uji dengan kondisi '' atau 'sama dengan 5
			Expect(2 + 3).To(BeNumerically("==", 5))
		})
	})
})
```


## Create and running unit test
Mari kita lanjutkan dengan membuat *unit test* untuk *project* kita sebelumnya di *file* `calculator_test.go`.

Pertama, kita dapat membuat *unit test* untuk *method* `Add` dari *struct* `Calculator`.

```go
var _ = Describe("Calculator", func() {
	Describe("Add", func() {
		// blok pengujian
		It("Should add the base number with adder number", func() {
			// arrange
			calc := calculator.Calculator{Base: 3}

			// act
			calc.Add(2)

			// assertion
			Expect(calc.Result()).To(Equal(float64(5)))
		})
	})
})
```
Untuk menjalankan *unit test* yang sudah dibuat, kita cukup menjalankan *command* `ginkgo` atau `go test` di *package* `calculator`.
```
> ginkgo
Running Suite: Calculator Suite - /Users/ruangguru/my-project/calculator
=========================================================================
Random Seed: 1661763280

Will run 1 of 1 specs
·
Ran 1 of 1 Specs in 0.000 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.229118292s
Test Suite Passed
```
Hasil dari *command* tersebut akan memunculkan *output* yang sedikit berbeda dengan *output* yang dapat kita lihat jika menggunakan *package* `testing`. *Output* dari ginkgo akan memunculkan output yang lebih mudah dipahami dan memberikan informasi yang lebih detail.

Kita juga dapat menambah opsi -v untuk mendapatkan output yang lebih detail.
```
> gingko -v
Running Suite: Calculator Suite - /Users/ruangguru/my-project/calculator
Random Seed: 1661763561

Will run 1 of 1 specs
------------------------------
Calculator Add Adding base number 3 with number 2
	Should return 5
	/Users/ruangguru/my-project/calculator/calculator_test.go:13
●

Ran 1 of 1 Specs in 0.000 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.184130542s
Test Suite Passed
```
Jika ternyata *test* yang kita buat tidak berhasil, kita dapat melihat *output* yang lebih detail di mana posisi yang gagal.

Anggap saja terdapat kekeliruan di *method* `Add` yang kita buat, dengan kode berikut:
```go
// ...

func (c *Calculator) Add(number float64) { // method untuk penjumlahan
	c.Base += number + 1 // kekeliruan, tidak sengaja menambah code '+1'
}
// ...
```
Dan kita coba jalankan *unit test* yang sudah dibuat.

```
> gingko
Running Suite: Calculator Suite - /Users/ruangguru/my-project/calculator
========================================================================
Random Seed: 1661763824

Will run 1 of 1 specs
● [FAILED] [0.000 seconds]
Calculator
/Users/ruangguru/my-project/calculator/calculator_test.go:10
	Add
	/Users/ruangguru/my-project/calculator/calculator_test.go:11
		Adding base number 3 with number 2
		/Users/ruangguru/my-project/calculator/calculator_test.go:12
			[It] Should return 5
			/Users/ruangguru/my-project/calculator/calculator_test.go:13

Expected
	<float64>: 6
to equal
	<float64>: 5
In [It] at: /Users/ruangguru/my-project/calculator/calculator_test.go:16
----------------------------

Ran 1 of 1 Specs in 0.001 seconds
FAIL! 0 Passed | 1 Failed | 0 Pending | 0 Skipped
--- FAIL: TestCalculator (0.005)
FAIL

Ginkgo ran 1 suite in 1.101791375s
Test Suite Failed
```
Hasil *output* yang dihasilkan akan lebih detail dan mudah dipahami, bahwa terdapat kekeliruan di *method* `Add` yang kita buat, harusnya menghasilkan `5` tapi di *method* yang kita buat malah menghasilkan `6`. Kita dapat melengkapi *unit test* yang sudah dibuat dengan menambahkan blok `It` yang berfungsi sebagai pengujian untuk setiap *method* yang kita buat.
```go
var _ = Describe("Calculator", func() {
	Describe("Add", func() {
		It("Should add the base number with adder number", func() {
			calc := calculator.InitCalculator(3)
			calc.Add(2)
			Expect(calc.Result()).To(Equal(5.0))
		})
	})

	Describe("Subtract", func() {
		It("Should subtract the base number with subtractor number", func() {
			calc := calculator.InitCalculator(10)
			calc.Subtract(5)
			Expect(calc.Result()).To(Equal(5.0))
		})
	})

	Describe("Multiply", func() {
		It("Should multiply the base number with multiplier number", func() {
			calc := calculator.InitCalculator(2)
			calc.Multiply(2)
			Expect(calc.Result()).To(Equal(4.0))
		})
	})

	Describe("Divide", func() {
		It("Should return error when the divisor is 0", func() {
			calc := calculator.InitCalculator(2)
			err := calc.Divide(0)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("cannot divide by zero"))
		})

		It("Should divide the base number with divisor number", func() {
			calc := calculator.InitCalculator(10)
			calc.Divide(2)
			Expect(calc.Result()).To(Equal(5.0))
		})
	})
})
```

Hasil dari *unit test* yang sudah kita buat akan menjadi seperti ini.
```
> ginkgo
Running Suite: Calculator Suite - /Users/ruangguru/my-project/calculator
========================================================================
Random Seed: 1661766188

Will run 5 of 5 specs
●●●●●

Ran 5 of 5 Specs in 0.000 seconds
SUCCESS! -- 5 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 519.209ms
Test Suite Passed
```
Hasil tersebut akan memberikan informasi bahwa terdapat 5 *test case* (*spec*) yang berhasil dijalankan dan semuanya berhasil. Kita dapat menambah opsi `-v` untuk mendapatkan *output* yang lebih detail.


## Multiple testing
Kita juga dapat membuat *multiple unit test* dengan menggunakan blok `Context` yang berfungsi untuk membuat beberapa *test case* yang sejenis. Anggap saja kita ingin membuat *test case* untuk *method* `Add` dengan beberapa kasus angka positif atau negatif.
```GO
type TestData struct {
	Spec   string // deskripsi dari test case
	Base   float64
	Input  float64
	Expect float64
}

var _ = Describe("Calculator", func() {
	Describe("Add", func() {
		var tests = []TestData{
			{"Adding base positive number with positive number", 10, 10, 20},
			{"Adding base negative number with positive number", -5, 10, 5},
			{"Adding base positive number with negative number", 10, -20, -10},
			{"Adding base negative number with negative number", -10, -10, -20},
		}

		for _, test := range tests {
			Context(test.Spec, func() {
				It("Should return the correct result", func() {
					calc := calculator.InitCalculator(test.Base)
					calc.Add(test.Input)
					Expect(calc.Result()).To(Equal(test.Expect))
				})
			})
		}
	})
})
```

Hasil dari *unit test* yang sudah kita buat akan menghasilkan 4 *test case* yang berbeda dan semuanya berhasil.
```
> ginkgo
Running Suite: Calculator Suite - /Users/ruangguru/my-project/calculator
========================================================================
Random Seed: 1661766938

Will run 4 of 4 specs
●●●●

Ran 4 of 4 Specs in 0.000 seconds
SUCCESS!-- 4 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.139358667s
Test Suite Passed
```

Pastikan kita membuat *test case* dengan semua kemungkinan yang ada agar tidak terjadi kecacatan di kode yang kita buat.


## Setup and teardown
Ginkgo dapat memudahkan kita untuk melakukan *setup* dan teardown pada *test case* yang kita buat. Kita dapat menggunakan blok `BeforeEach` dan `AfterEach` untuk melakukan *setup* dan *teardown* pada *test case* yang kita buat.

Kita dapat merubah *unit test* yang sudah kita buat menjadi seperti ini.
```go
var _ = Describe("Calculator", func() {
	// arrange
	var calc calculator.Calculator

	BeforeEach(func() {
		calc = calculator.InitCalculator(5) // setiap kali pengujian, 'Base' di calculator akan diisi 5
	})

	Describe("Add", func() {
		It("Should add the base number with adder number", func() {
			calc.Add(4)
			Expect(calc.Result()).To(Equal(9.0))
		})
	})

	Describe("Subtract", func() {
		It("Should subtract the base number with subtractor number", func() {
			calc.Subtract(2)
			Expect(calc.Result()).To(Equal(3.0))
		})
	})

	Describe("Multiple", func() {
		It("Should multiply the base number with multiplier number", func() {
			calc.Multiply(10)
			Expect(calc.Result()).To(Equal(50.0))
		})
	})

	Describe("divide", func() {
		It("Should divide the base number with divider number", func() {
			calc.Divide(5)
			Expect(calc.Result()).To(Equal(1.0))
		})
	})

	AfterEach(func() {
		calc.Reset() // setiap kali selesai pengujian, 'Base' akan di reset dengan nilai 0
	})
})
```

Hasil dari unit test yang sudah kita buat akan menjadi seperti ini.
```
> ginkgo
Running Suite: Calculator Suite - /Users/ruangguru/my-project/calculator
========================================================================
Random Seed: 1661772787

Will run 4 of 4 specs
●●●●

Ran 4 of 4 Specs in 0.000 seconds
SUCCESS! -- 4 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.087910542s
Test Suite Passed
```

Dari contoh di atas, `BeforeEach` akan dijalankan sebelum *test case* yang kita buat dengan melakukan inisiasi *struct* `Calculator` dan nilai properti `Base` akan diisi `5`. Dan setiap kali selesai pengujian, akan memanggil `AfterEach` yang akan mengosongkan `Base` dengan nilai `0`.

Dengan menggunakan `BeforeEach` dan `AfterEach`, kita dapat mempersingkat beberapa pengaturan / persiapan data di *code unit test* dan dapat berfokus pada *test case* yang kita buat.

## Parallel running test
Ginkgo mendukung pengujian secara parallel dengan menggunakan opsi `-p` atau `--parallel`. Hal ini dapat membantu kita untuk menghemat waktu dalam proses pengujian, apalagi jika kita memiliki banyak sekali *test case* yang akan dijalankan.

Jika kita mencoba menjalankan pengujian dengan opsi `-p` atau `--parallel`, di *project* kita, maka akan menghasilkan output seperti ini:
```
> ginkgo -p
Running Suite: Calculator Suite - /Users/ruangguru/my-project/calculator
========================================================================
Random Seed: 1661773194

Will run 4 of 4 specs
Running in parallel across 7 processes
●●●●

Ran 4 of 4 Specs in 0.004 seconds
SUCCESS! -- 4 Passed | 0 Failed | 0 Pending | 0 Skipped

Ginkgo ran 1 suite in 1.118454s
Test Suite Passed
```

Terdapat informasi tambahan, yaitu `Running in parallel across 7 processes`, yang menandakan bahwa kita menjalankan pengujian secara *parallel* dengan jumlah proses yang ada adalah `7`. Sehingga seluruh proses pengujian memakan waktu `1.118454` detik.

## Integration testing
Ginkgo dirancang dari awal untuk mendukung *integration testing*. Salah satu masalah yang kadang muncul dari *integration testing* adalah terdapat *test case* yang kadang berhasil, tetapi kadang gagal atau disebut "angin-anginan". Hal ini muncul karena kita menggunakan *test case* yang berbeda-beda dan saling terintegrasi antar tiap *test case*, dan memungkinkan terjadi kasus yang tidak diharapkan. 

Ginkgo memberikan opsi untuk mengatasi masalah ini dengan menggunakan opsi `--repeat`. Kita dapat menguji berkali-kali *test case* yang kita buat dengan menggunakan opsi ini dan perlu menjalankan *command* berkali-kali. Kita cukup mengisi opsi `--repeat` dengan jumlah percobaan yang kita inginkan.
```
> ginkgo --repeat-N # mengulang test sebanyak N kali
```

Kita juga dapat menggunakan opsi `--flake-attempts` untuk menguji *test case* yang gagal dan memastikan bahwa *test case* tersebut benar-benar gagal, bukan *test case* yang kadang gagal kadang berhasil. Cara penggunaannya hampir sama seperti opsi `--repeat`, yaitu melakukan pengujian sebanyak `N` kali.
```
> ginkgo --flake-attempts-N # mengulang test yang gagal sebanyak N kali
```

Terakhir, terdapat opsi yang dapat digunakan untuk menghentikan pengujian yang tidak kunjung selesai. Biasanya kasus ini terjadi karena kita melakukan pengujian untuk integrasi yang prosesnya *blocking* atau tidak berjalan sebagaimana mestinya. Kita dapat menghentikan pengujian dengan menggunakan opsi `--timeout` dan juga dapat memastikan bahwa proses integrasi ini tidak melebihi waktu yang sudah ditentukan.
```
> ginkgo --timeout-duration # menghentikan pengujian jika prosesnya melebihi N detik
```

Durasi untuk opsi `--timeout` adalah format *duration*, kita dapat menentukan format waktu yang berbeda dengan menambah informasi:
* `ms` untuk milidetik.
* `s` untuk detik
* `m` untuk menit
* `h` untuk jam
```
> ginkgo --timeout-100ms # menghentikan pengujian jika prosesnya melebihi 100 milidetik > ginkgo --timeout=10s # menghentikan pengujian jika prosesnya melebihi 10 detik
> ginkgo --timeout=1m # menghentikan pengujian jika prosesnya melebihi 1 menit
> ginkgo --timeout=1m30s # menghentikan pengujian jika prosesnya melebihi 1 menit 30 detik
```
