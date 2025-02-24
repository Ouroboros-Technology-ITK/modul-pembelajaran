# Package sort


## `sort`
Di Golang terdapat *package* `sort` yang berisi utilitas untuk kebutuhan *sorting* atau pengurutan data.

Algoritma sorting adalah algoritma yang digunakan untuk pengurutan sebuah kumpulan data dari terkecil ke terbesar atau sebaliknya. Algoritma ini biasanya digunakan untuk mengurutkan data dalam sebuah *array* atau *slice*.

Terdapat banyak logika algoritma yang bisa dipakai untuk melakukan *sorting* (bisa dilihat [di sini]()). Namun, di package `sort` mengimplementasikan algoritma *quick sort*.

Kita juga bisa melihat visualisasinya di website [Visualgo]().

## Basic sorting
### `sort.Ints()`
Fungsi ini bisa digunakan untuk mengurutkan tipe data `int` di Golang. Kita cukup mengisi parameter dari fungsi tersebut tanpa perlu menangkap hasilnya di sebuah variabel baru. Hasil dari pengurutan di fungsi ini otomatis akan mengubah urutan dari variabel yang *diinputkan*.
```go
import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 4, 3, 2, 1}
	sort.Ints(numbers) // tidak perlu menangkap hasilnya di variable, hasilnya akan otomatis mengurutka
	fmt.Println(numbers)
}
```
```Output
[1 2 3 4 5]
```

Contoh di atas adalah penggunaan fungsi `sort.Ints()` untuk mengurutkan data *integer*.

### `sort.Strings()`
Hampir sama seperti fungsi `Ints`, tapi fungsi ini digunakan khusus untuk mengurutkan data *string*.
```go
import (
	"fmt"
	"sort"
)

func main() {
	fruits := []string{"banana", "apple", "pear"}
	sort.Strings(fruits)
	fmt.Println(fruits)
}
```
Output
```Output
[apple banana pear]
```

### `sort.Float64s()`
Fungsi ini juga sama, tapi digunakan untuk mengurutkan data `float64`.
```go
import (
	"fmt"
	"sort"
)

func main() {
	prices := []float64(3.3, 5.5, 1.1, 2.2}
	sort.Float64s(prices)
	fmt.Println(prices)
}
```
Output:
```Output
[1.1 2.2 3.3 5.5]
```


## Custom comparator
Kita bisa menggunakan fungsi `sort.Slice()` untuk mengurutkan data dengan menggunakan fungsi *comparator* yang kita buat sendiri. Fungsi ini dapat digunakan untuk semua tipe data (`any`) di Golang. Dan dapat memudahkan kita ketika ingin pengurutan yang berbeda, seperti dari besar ke kecil.

Terdapat 2 parameter yang dibutuhkan untuk fungsi ini, yaitu *slice* dan parameter fungsi dengan *signature* `func (i, j int) bool`.
```go

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}
func main() {
	people := []Person{
		{"Bob", 31),
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26),
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age › people[j].Age // komparasi berdasarkan properti 'Age' di struct 'Person'
	})
	fmt.Println(people)
}
```
Output:
```Output
[{John 42} {Bob 31} {Jenny 26} {Michael 17}]
```

## Custom sorting
Di package `sort` terdapat *interface* `sort.Interface` yang bisa kita gunakan untuk melakukan kustomisasi *sorting*. Terdapat *method* yang harus diimplementasikan di interface ini, yaitu `Len()`, `Less()`, dan `Swap()`.
```go
// documentation from package sort

type Interface interface {
	// fungsi ini digunakan untuk mengembalikan jumlah data yang ada di slice
	Len() int
	
	// fungsi ini digunakan untuk melakukan komparasi antara 2 data di index 'i' dan 'j'
	// sebelumnya sudah dicontohkan di fungsi sort. Slice()"
	Less(i, j int) bool
	// fungsi ini digunakan untuk menukar data di index 'i' dan 'j'
	Swap(i, j int)
}
```

Jika kita ingin melakukan *sorting* untuk tipe data *struct* sebagai berikut, kita harus membuat *method* yang dibutuhkan untuk *interface* `sort.Interface`:
```go
type Person struct {
	Name string
	Age int
}
type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
```

Kita dapat menggunakan `sort.Sort` untuk melakukan *custom sorting*.
```go
func main() {
	persons := Persons{
		Person("John", 30),
		Person("Jane", 25),
		Person{"Jack", 40},
		Person("Jill", 45},
	}
	
	sort.Sort(persons)
	
	fmt.Println(persons)
}
```
Output:
```Output
[{Jane 25} {John 30} {Jack 40} {Jill 45}]
```

## Sort map by key or value
Kita bisa mengurutkan *map* berdasarkan *key* atau *value* dengan cara menampung terlebih dahulu ke dalam *array* atau *slice*.
```go
func main() {
	m := map[string]int{
		"Alice" : 1,
		"Cecil" : 2,
		"Bob" : 3,
	}

	key := make([]string, 0, len(m)) // membuat penampung untuk key dalam bentuk array

	for k := range m {
		key = append(key, k) // mengambil semua key dari map ke dalam array
	}

	sort.Strings(key) // menggunakan fungsi 'sort.Strings', karena key berupa tipe data string

	for _, k := range key {
		fmt.Println(k, m[k]) // menampilkan key dan value dari map setelah diurutkan
	}
}
```
Output:
```Output
Alice 1
Bob 3
Cecil 2
```

Contoh di atas akan memunculkan urutan data berdasarkan *key* yang sudah diurutkan.

>TL;DR: Untuk versi Golang 1.12 ke atas, package fmt akan menampilkan *output* tipe data map yang sudah terurut berdasarkan *key*.
