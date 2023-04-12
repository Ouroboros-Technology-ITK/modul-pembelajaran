# Package strings

## strings
Golang menyediakan *package* `strings`, untuk keperluan pengolahan data `string`.

### `strings.ToLower()`
Mengubah semua huruf di *string* menjadi huruf kecil. Cara menggunakannya cukup memasukkan tipe data `string` ke parameter fungsinya.
```go
import "strings"
func main() {
	fmt.Println(strings.ToLower("HELLO GO!"))
	fmt.Println(strings.ToLower("Hello Go!"))
}
```
Output:
```Output
hello go!
hello go!
```

### `strings.ToUpper()`
Mengubah semua huruf kecil menjadi huruf besar. Cara menggunakannya sama seperti fungsi `ToLower()`.
```go
import "strings"

func main() {
	fmt.Println(strings.ToLower("hello go!"))
	fmt.Println(strings.ToLower("Hello Go!"))
}
```
Output:
```Output
HELLO GO!
HELLO GO!
```

### `strings.Replace()`
Fungsi ini digunakan untuk mengganti bagian dari *string* dengan *substring* tertentu. Jumlah *substring* yang diganti juga dapat ditentukan sesuai kebutuhan. Terdapat 4 parameter yang harus diisi:
* Parameter pertama adalah *string* yang akan digunakan sebagai string utama (base string).
* Parameter kedua, adalah *substring* yang akan diganti.
* Parameter ketiga, adalah *substring* penggantinya.
* Parameter keempat, adalah jumlah *substring* yang diganti.
```go
import "strings"

func main() {
	var newString = strings.Replace("banana", "a", "i", 1) 
	var newString2 = strings.Replace("banana", "a", "i", 2)
	
	fmt.Println(newString) // binana
	fmt.Println(newString2) // binina
}
```
Output:
```Output
binana
binina
```

Terdapat opsi jika ingin mengubah semua *substring* yang ada di *string*, dengan mengisi parameter ke 4 dengan `-1`.
```go
import "strings"

func main() {
	var newString = strings.Replace("banana", "a", "i", -1)
	fmt.Println(newString) // binini
}
```
Output:
```Output
binini
```

### `strings.Split()`
Digunakan untuk memisah `string` menjadi *array* / *slice* menggunakan tanda pemisah yang kita tentukan.

```go
import "strings"

func main() {
	var str = strings.Split("Welcome to Ruangguru", "")
	
	fmt.Println(str)
}
```
Output:
```Output
["Welcome", "to", "Ruangguru"]
```
Jika sudah menentukan pemisahnya tetapi tidak ditemukan di dalam string, maka *slice* akan tetap terbentuk dan berisi data *string* tanpa adanya pemisahan.

```go
import "strings"

func main() {
	var str = strings.Split("Hello Go!", "oo")
	
	fmt.Println(str)
}
```
Output:
```Output
["Hello Go!"]
```
Untuk memisah setiap *string* menjadi satu karakter di tiap *slice* nya, kita dapat menggunakan tanda pemisah *string* kosong `""`.
```go
import "strings"

func main() {
	var str = strings.Split("Hello", "")
	
	fmt.Println(str)
}
```
```
["H", "e", "1", "1", "o"]
```

### `strings.Join()`
Kebalikan dari `Split()`, yaitu menggabungkan *array* / *slice of string* menjadi satu *string* yang digabungkan dengan *substring* yang sudah ditentukan.
```go
import "strings"

func main() {
	var str = strings.Join([]string{"Hello", "Go!"}, "-")
	// parameter pertama adalah list array string
	// prameter kedua adalah substring yang akan digabungkan
	fmt.Println(str)
}
```
Output:
```Output
Hello-Go!
```

### `strings.Contains()`
Dipakai untuk mendeteksi apakah *string* (parameter kedua) merupakan bagian dari *string* lain (parameter pertama). Nilai kembaliannya berupa `bool`.
```go
import "strings"

func main() {
	fmt.Println(strings.Contains("Hello Go!", "Hello")) // true 
	fmt.Println(strings.Contains("Hello Go!", "Galaxy")) // false
}
```
Output:
```Output
true
false
```

`Println` pertama akan menampilkan `true`, karena *string* `Hello Go!` memiliki *string* `Hello`. Sedangkan `Println` kedua akan menampilkan `false`, karena *string* `Hello Go!` tidak memiliki string `Galaxy`.

### `strings.HasPrefix()`
Dipakai untuk mendeteksi apakah *string* (parameter kedua) merupakan awalan dari *string* lain (parameter pertama). Nilai kembaliannya berupa `bool`.
```go
import "strings"
func main() {
	fmt.Println(strings.HasPrefix("Hello Go!", "He")) // true
	fmt.Println(strings.HasPrefix("Hello Go!", "Go")) // false
}
```
Output:
```Output
true
false
```

### `strings.HasSuffix()`
Dipakai untuk mendeteksi apakah *string* di parameter pertama di akhiri dengan *string* tertentu (di parameter kedua).
```go
import "strings"
func main() {
	fmt.Println(strings.HasSuffix("Hello Go!", "He")) // false 
	fmt.Println(strings.HasSuffix("Hello Go!", "Go!")) // true
}
```
Output:
```Output
false
true
```

### `strings.Count()`
Dipakai untuk menghitung jumlah *substring* (parameter kedua) dalam *string* lain (parameter pertama).
```go
import "strings"
func main() {
	fmt.Println(strings.Count()"Hello Go!", "o")) // 2
}
```
Output:
```Output
2
```
Nilai kembalian bertipe data `int` dengan nilai `2` karena *string* `Hello Go!` memiliki 2 *substring* `o`.

### `strings.Index()`
Digunakan untuk mencari *index* dari *string* yang dicari (parameter kedua) dalam *string* lain (parameter pertama). Nilai kembalian bertipe data `int`.
```go
import "strings"
func main() {
	fmt.Println(strings.Index("Hello Go!", "ll")) // 2
}
```
Output:
```Output
2
```
*String* `11` berada pada *index* ke 2 dari *string* `Hello Go!` (*index* dimulai dari 0).

Jika terdapat lebih dari 1 *substring* yang sama, maka hanya akan mengembalikan *index* pertama.
```go
import "strings"

func main() {
	fmt.Println(strings.Index("Hello Go!", "o")) // 4
}
```
Output: 
```Output
4
```
String `o` pertama kali didapat di *index* ke 4, sehingga akan memunculkan *output* `4` walaupun terdapat juga substring `o` di index ke 7.

### `string.Repeat()`
Digunakan untuk mengulang *string* (parameter pertama) sebanyak yang ditentukan (parameter kedua).
```go
import "strings"
func main() {
	var str = strings.Repeat("Hello", 3)

	fmt.Println(str) // HelloHelloHello
}
```
Output:
```Output
HelloHelloHello
```
