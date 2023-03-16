# Package regexp


**Regex** atau **regexp** atau *regular expression* adalah suatu teknik yang digunakan untuk pencocokan string dengan pola tertentu. *Regex* biasa dimanfaatkan untuk pencarian dan pengubahan data *string*.

Go mengadopsi standar *regex* RE2, untuk melihat sintaks yang di-support *engine* Golang, bisa langsung merujuk ke dokumentasinya [di sini](). Atau mencoba menggunakan regex [di sini]().

## `regexp`
Golang memiliki *package* `regexp` yang bisa digunakan untuk memanfaatkan konsep *regular expression* di program kita.

## `regexp.Compile()`
Fungsi `regexp.Compile()` digunakan untuk mengkompilasi ekspresi *regex*. Fungsi tersebut mengembalikan objek bertipe `regexp.*Regexp`.

Kita bisa memasukkan *pattern* yang ingin dicocokkan ke dalam parameter fungsi ini. Fungsi ini mengembalikan 2 nilai, yaitu objek bertipe `regexp.*Regexp` dan *error*. Jangan lupa untuk mengecek dan *menghandle* jika terjadi error.
```go
func main() {
	var regex, err = regexp.Compile(`[a-zA-Z]+`)
	
	if err != nil {
		panic(err)
	}
}
```

Ekspresi `[a-zA-Z]+` maksudnya adalah, semua *string* yang merupakan *alphabet* baik itu hurufnya kecil (a-z) ataupun besar (A-Z). Ekspresi tersebut di-*compile* oleh `regexp.Compile()` lalu disimpan ke variabel `regex`.

Di dalam *object* `regexp.*Regexp` terdapat beberapa *method* yang bisa digunakan untuk mencocokkan *string* dengan ekspresi *regex* yang sudah *dicompile*.

## `regexp.MatchString()`
*Method* ini digunakan untuk mendeteksi apakah *string* memenuhi sebuah pola atau *pattern regexp*. *Output*-nya adalah `bool`.
```go
func main() {
	var regex, err = regexp.Compile('[a-zA-Z]+')
	if err != nil {
		panic(err)
	}

	var isMatch = regex.MatchString("Hello Go!`)
	
	fmt.Println(isMatch) // true
}
```
Pada contoh di atas `isMatch` bernilai `true` karena *string* `"Hello Go!"` memenuhi pola *regex* `[a- zA-Z]+`.

## `regexp.FindString()`
*Method* ini digunakan untuk mencari *string* yang memenuhi sebuah pola atau *pattern regexp*. *Output*-nya adalah *substring*.
```go
func main() {
	var regex, err = regexp.Compile(^[a-zA-Z]+")
	if err != nil {
		panic(err)
	}
	var find = regex.FindString("Hello Go!`)
	
	fmt.Println(find)
}
```
Output:
```Output
Hello
```
Fungsi ini hanya mengembalikan 1 buah hasil saja. Jika ada banyak *substring* yang sesuai dengan ekspresi *regexp*, akan dikembalikan yang pertama saja.

Contoh di atas, kita mendapat hasil `Hello` karena *string* `"Hello Go!"` memenuhi pola *regex* `[a-zA- Z]+`.

## `regexp.FindAllString()`
Sama dengan `regexp.FindString()` hanya saja fungsi ini mengembalikan semua *substring* yang sesuai dengan ekspresi *regexp*.

Terdapat 2 parameter, parameter pertama adalah *string* yang akan dicari, parameter kedua adalah jumlah *substring* yang akan dikembalikan. Hasilnya akan mengembalikan *array of string*.
```go
func main() {
	var regex, err = regexp.Compile('[a-zA-Z]+')
	if err != nil {
		panic(err)
	}
	var findOne = regex.FindString("Hello Go!", 1)
	var findAll = regex.FindString("Hello Go!", -1)
	
	fmt.Println(findOne)
	fmt.Println(findAll)
}
```
Output
```Output
[Hello]
[Hello Go]
```
Jika parameter kedua diisi dengan `-1`, maka akan mengembalikan semua data.

## `regexp.ReplaceAllString()`
Berguna untuk me-*replace* semua *string* yang memenuhi kriteria *regexp*, dengan *string* lain.
```go
}
func main() {
var regex, err = regexp.Compile(^[a-zA-Z]+)
if err != nil {
}
panic(err)
var replace regex. ReplaceAllString("Hello Go!`, `Hi!")
fmt.Println(replace)
```
```go
Hi! Hi!!
```
Pada contoh di atas, karena *substring* `Hello` dan `Go` memenuhi kriteria *regexp*, maka semua akan di-*replace* dengan *string* `Hi!`.
