# Variable

Variabel adalah sebuah tempat untuk menyimpan data atau nilai.  

Golang mengadopsi dua jenis penulisan variabel, yaitu yang dituliskan tipe data-nya dan juga yang tidak. Kedua cara tersebut valid dan tujuannya sama, pembedanya hanya cara penulisannya saja.  

## Basic declaration  

Penulisan variabel yang paling dasar di Golang adalah menggunakan syntax var diikuti dengan nama variabel dan tipe datanya.  
```go
var hello string  
var age int  
var pi float64  
```

Kita dapat mengisi nilainya dengan cara di assign ke variabel yang sudah dideklarasikan menggunakan syntax `=`.
```go
var hello string  
hello "Hello Go!" // assign new value  
```

Atau dapat langsung mengisi nilainya pada saat deklarasi variabel.  
```go
var hello string = "Hello Go!"  
```

Terdapat opsi lain jika ingin langsung mengisi nilai variabel, yaitu **tidak perlu** mendeklarasikan tipe datanya. Golang akan otomatis menentukan tipe data yang tepat berdasarkan nilai yang diisi.  
```go
var hello= "Hello Go!"  
// sama saja dengan  
// var hello string "Hello Go!"  
```

Ketika sudah mendeklarasikan variabel beserta tipe datanya maka tidak dapat diisi dengan tipe data lain.  
```go
var hello "Hello Go!"  
hello= true // error: string cannot assign with bool data type  
```

## Default value for variable  

Jika deklarasi dari nilai variabel tidak ditentukan maka variabel akan diatur ke nilai *default* dari tipenya :  
TIPE DATA  | DEFAULT VALUE  
--|--
string  | ""  
int  |0
bool  |false  


```go
func main() {  
	var (
		name string  
		age int  
		isMarried bool  
	)
	fmt.Println(name) // ""  
	fmt.Println(age) // 0  
	fmt.Println(isMarried) // false  
}
```

## Short variable declaration  

Cara kedua, yaitu menggunakan *short syntax* dengan menuliskan nama variabel diikuti dengan *syntax* `:=` dan nilai variabelnya. Penulisan ini tidak memerlukan deklarasi tipe data. Variabel tersebut akan mendapatkan tipe data berdasarkan nilai yang diisi. 
```go
hello := "Hello World"  
// sama saja dengan  
// var hello string "Hello World" 
``` 

Sama dengan penulisan yang biasa, ketika sudah dideklarasikan maka tidak dapat di-*assign* dengan tipe data yang berbeda.  

Terdapat beberapa perbedaan antara basic *syntax* dan *short syntax*, yaitu:  
VAR (BASIC SYNTAX)  |:= (SHORT SYNTAX)  
--|--
Bisa digunakan *di dalam* dan *di luar* functions  |Hanya bisa digunakan di dalam function.  
Deklarasi variabel dan assign nilai dapat dilakukan secara terpisah | Deklarasi variabel dan assign nilai harus dilakukan tanpa terpisah

## Constanta (const)  

Jika suatu variabel harus memiliki nilai tetap yang tidak dapat diubah, Anda dapat menggunakan *syntax* `const`.  

Type `const` mendeklarasikan variabel sebagai "konstan", yang berarti variabel tersebut tidak dapat diubah dan hanya bisa dibaca.  

Nilai atau data konstanta harus langsung diisi saat awal deklarasi.  
```go
const pi float64 = 3.14159265359  
// jika kita melakukan assign nilai baru, maka error  
// pi 2.34 --> error  
```

## Multi var and const  
Di Golang, kita dapat mendeklarasikan beberapa variabel atau constant dalam satu baris sehingga lebih sederhana.  
```go
var name string  
var address string  
var status string  
// can be simplified to  
var name, address, status string  
```

Contoh di atas adalah variable `name`, `address`, `status` dengan tipe data `string`.  
Atau jika ingin membuat beberapa *variable* dengan tipe data yang berbeda, kita cukup membuat satu *syntax* `var` dengan dibungkus `(` dan `)`.  
```go
var ( 
	name string  
	age int  
	isMarried bool
)
```

Golang juga dapat meng-*declare* beberapa constanta dengan cara yang sama, hanya saja dengan tipe data `const`.  

```go
const (  
	E float64  
	Pi float64 3.14 2.718  
)
```
Kita juga dapat membuat beberapa *short syntax* variable dengan lebih sederhana.  

```go
name, age, isMarried: "Budi", 30, true  
```

## Variabel Underscore  ` _ `

Golang memiliki aturan unik, yaitu tidak boleh ada satupun variabel yang menganggur (tidak dipakai). Artinya, semua variabel yang dideklarasikan harus digunakan. Jika ada variabel yang tidak dipakai, *error* akan muncul pada saat kompilasi dan program tidak akan bisa di jalankan di *terminal* atau *console*.  
```go
func main() {  
var hello "Hello Go!"  
}  
```
```Output
# command-line-arguments  
./main.go:6 hello declared and not used  
```

*Underscore* `_` adalah *reserved variable* yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah.  

Biasanya variabel *underscore* sering dimanfaatkan untuk menampung nilai balik fungsi yang tidak digunakan.  
```go
hello, _ := "Hello", "Go!"  
_, _ = fmt.Scanf("%s", &hello)  
```

Pada contoh di atas, variabel `hello` akan berisikan teks `"Hello"`, sedang nilai `"Go!"` ditampung oleh variabel *underscore*, menandakan bahwa nilai tersebut tidak akan digunakan.  

>Note: Perlu diketahui, bahwa isi variabel underscore tidak dapat ditampilkan. Data yang sudah masuk variabel tersebut akan hilang  

## Variable naming rule  

Terdapat beberapa aturan penamaan variabel di Golang, yaitu:
* Nama variabel harus dimulai dengan huruf atau karakter garis bawah (), tidak boleh di awali dengan angka.
* Nama variabel hanya boleh berisi karakter alfanumerik dan garis bawah (a-z, A-Z, 0-9, dan _).
* Nama variabel ini case-sensitive (age, Age and AGE adalah variable yang berbeda).
* Nama variabel tidak boleh mengandung spasi.
* Nama variabel tidak boleh sama persis dengan keyword Golang (var, const, int, bool, dll).  

```go
var 1name string // error: invalid variable name  
var var string // error: keyword var cannot be used as variable name  
var hello world string // error: var name with space  
```

## Multi-Word Variable Names  

Nama variabel dengan lebih dari satu kata bisa jadi sulit dibaca. Convention (aturan umum tidak tertulis)  
di Golang adalah menggunakan **PascalCase** atau **camelCase**.  

• Camel case: Setiap kata kecuali yang pertama, dimulai dengan huruf kapital.  
• Pascal case: Setiap kata dimulai dengan huruf kapital.  
• Snake case: Setiap kata dipisahkan dengan garis bawah.  
```go
var myVariableName = "Hello Go!" // camel case  
var MyVariableName = "Hello Go!" // pascal case  
var my_variable_name "Hello Go!" // snake case.  
```
> Note: Snake Case tidak disarankan digunakan di Golang, karena jarang ada yang menggunakan convention tersebut. 
