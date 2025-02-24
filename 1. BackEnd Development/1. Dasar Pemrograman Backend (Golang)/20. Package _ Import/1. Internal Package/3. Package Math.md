# Package math

## `math`
Golang menyediakan package `math` yang berisikan konstanta dasar dan keperluan operasi matematis. Namun di sini, kita tidak akan membahas semua fungsi dari package `math`, hanya fungsi-fungsi umum saja.

### `math.floor()`
Fungsi ini mengubah bilangan desimal menjadi bilangan bulat terdekat ke bawah. Cara menggunakannya cukup memasukkan tipe data float64 ke parameter fungsinya.
```go
import "math"

func main() {
	fmt.Println(math.Floor(3.7)) // 3
	fmt.Println(math.Floor(10 / 2.7)) // 3
}
```
Output:
```Output
3
3
```
Dari contoh di atas, karena `3.7` merupakan bilangan desimal, maka akan dibulatkan ke bawah menjadi `3`. Sama halnya, jika `10 / 2.7` merupakan bilangan desimal dengah hasil `3.7037`, maka akan dibulatkan ke bawah menjadi `3`.

### `math.Ceil()`
Fungsi ini mengubah bilangan desimal menjadi bilangan bulat terdekat ke atas. Cara menggunakannya sama dengan `math.Floor()`, yaitu cukup memasukkan tipe data `float64` ke parameter fungsinya.
```go
import "math"

func main() {
	fmt.Println(math.Ceil(3.7)) // 4
	fmt.Println(math.Ceil(10 / 2.7)) // 4
}
```
Output:
```Output
4
4
```
Contoh di atas adalah bentuk pembulatan ke atas dari `3.7` dan `10 / 2.7`.

### `math.Abs()`
Fungsi ini mengubah bilangan menjadi absolut. Jika bilangan yang dimasukkan adalah negatif, maka akan menjadi positif. Kalau bilangannya adalah positif, maka akan tetap.
```go
import "math"

func main() {
	fmt.Println(math.Abs(-3.7))
	fmt.Println(math.Abs(10 / 2.7))
}
```
Output:
```Output
3.7
3.7037037037037037
```
Contoh di atas adalah mengubah bilangan menjadi absolut, tetapi tidak akan dibulatkan seperti fungsi yang dijelaskan sebelumnya.

### `math.pow()`
Fungsi ini digunakan untuk menghitung perpangkatan. Terdapat 2 parameter yang harus diisi.
```go
import "math"

func main() {
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Pow(2, -3))
}
```
Output:
```Output
8
0.125
```
Beberapa case lainnya bisa dilihat [di sini]().

## Random
Di Golang terdapat sebuah *package* yaitu `math`/`rand` yang isinya banyak sekali fungsi untuk keperluan penciptaan angka *random*. *Package* ini mengadopsi **PRNG** atau *pseudo-random number generator*. Deret angka random yang dihasilkan sangat tergantung dengan angka `seed` yang digunakan.

*Random Number Generator* (RNG) merupakan sebuah perangkat (bisa *software*, bisa *hardware*) yang menghasilkan data deret/urutan angka yang sifatnya acak.

RNG bisa berupa *hardware* yang murni bisa menghasilkan data angka acak, atau bisa saja sebuah *pseudo-random* yang menghasilkan deret angka-angka yang terlihat acak tetapi sebenarnya tidak benar-benar acak, yang deret angka tersebut sebenarnya merupakan hasil kalkulasi algoritma deterministik dan probabilitas. Jadi untuk *pseudo-random* ini, asalkan kita tahu *state*-nya maka kita akan bisa menebak hasil deret angka *random*-nya.

Dalam konsep *random* terdapat istilah `seed` atau titik mulai (*starting point*). *Seed* ini digunakan oleh RNG dalam pembuatan angka random di tiap urutannya.

Cara menggunakan random di Golang cukup *import* `math/rand`, lalu set seed-nya, lalu panggil fungsi untuk *generate* angka *random*-nya.
```go
import "math/rand"

func main() {
	rand.Seed(10) // starting point dengan seed '10'
	
	fmt.Println("random 1:", rand.Int())
	fmt.Println("random 2:", rand.Int())
	fmt.Println("random 3:", rand.Int())
}
```
Output:
```Output
random 1: 5221277731205826435
random 2: 3852159813000522384
random 3: 8532807521486154107
```

Fungsi `rand.Seed()` digunakan untuk penentuan nilai *seed*. Fungsi `rand.Int()` digunakan untuk *generate* angka random dalam bentuk numerik bertipe `int`. Fungsi `rand.Int()` ini setiap kali dipanggil akan menghasilkan angka berbeda, tapi pasti hasilnya akan selalu tetap jika mengacu ke deret.
Anehnya, ketika kita coba eksekusi code tersebut berkali-kali, hasil *generate*-nya akan sama terus menerus. Hal ini karena kita mengatur seed nya selalu sama yaitu 10.
Output 1:
```
random 1: 5221277731205826435
random 2: 3852159813000522384 
random 3: 8532807521486154107 
```
Output 2:
```
random 1: 5221277731205826435
random 2: 3852159813000522384
random 3: 8532807521486154107
```

Terus bagaimana cara agar angka yang dihasilkan selalu berbeda setiap kali di-eksekusi?

Caranya adalah menggunakan angka yang *unique* atau unik sebagai `seed`, contohnya seperti angka `unix nano` dari waktu sekarang.
Kita bisa memodifikasi fungsi `rand.Seed()` sebagai berikut
```go
import "math/rand"
func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // set seed dengan 'unix nano' waktu sekarang
	
	fmt.Println("random 1:", rand.Int())
	fmt.Println("random 2:", rand.Int())
	fmt.Println("random 3:", rand.Int())
}
```

Karena waktu terus berubah, maka nilai dari seed akan juga berubah. Sehingga menghasilkan angka *random* yang benar-benar *random*.

Output 1:
```Output
random 1: 7570338795182245453 
random 2: 4162021830936518477 
random 3: 4020747620041710952 
```
Output 2: 
```Output
random 1: 6528598712357027683 
random 2: 7882030262508925758 
random 3: 5937383941118082047
```

### Random with custom range
Kita dapat mengatur *upper range random* angka, seperti dari angka 0 sampai 20 atau dari 0 sampai 100. Cukup menggunakan fungsi `rand.Intn()` yang akan menghasilkan angka *random* antara 0 hingga *upper range* yang kita tentukan.

```go
import ("math/rand", "time")

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random 0 - 20:", rand.Intn(20))
	fmt.Println("random 0 - 100:", rand.Intn(100))
}
```
Output:
```Output
random 0 - 20: 19
random 0 - 100: 71
```


### Random with string data type
Untuk menghasilkan data *random string*, ada banyak cara yang bisa digunakan, salah satunya adalah dengan memanfaatkan alfabet dan hasil *random* numerik dari *package* `math/rand`.

Codenya seperti berikut:
```go
import "math/rand"
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b = make([]rune, length)
	for i = range b {
		b[i]= letters[rand.Intn(len(letters))]
	}
	return string(b)
}
```

Dengan fungsi di atas kita bisa dengan mudah melakukan *generate string random* dengan panjang karakter yang sudah ditentukan, misal `randomString(10)` akan menghasilkan *random string* 10 karakter.
```go
func main() {
	fmt.Println("random string: ", randomString(10))
}
```
Output:
```
random string: EVuzySQOHW
```
