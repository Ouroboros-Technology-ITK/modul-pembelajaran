# Numeric  
## What is numeric?  
*Numeric* adalah tipe data yang berisi karakter angka (0-9). Golang memiliki 2 jenis tipe data *number* yaitu `integer` (bilangan bulat) dan `floating point` (bilangan desimal).  

Secara lebih detail tipe data *number* di Golang dibagi menjadi beberapa tipe data berdasarkan kegunaan dan alokasi memorinya. Berikut tipe data & yang dapat digunakan:  
 
### int
Tipe data *integer* dibagi menjadi 4 tipe data:  
Console I/O  

TIPE DATA  | NILAI MINIMUM | NILAI MAKSIMUM
--|--|--
int8 | -128 | 127   
int16 | -32768  | 32767
int32 | -2147483648 | 2147483647
int64 | -9223372036854775808 | 9223372036854775807

Setiap tipe data mewakili cakupan angka dan alokasi memorinya. Jika membutuhkan data dengan cakupan hanya 0-100, kita cukup menggunakan `int8` sebagai tipe datanya. Sebaliknya, jika membutuhkan data dengan cakupannya lebih dari 10 miliar, kita bisa menggunakan `int64` sebagai tipe datanya karena cakupannya yang lebih luas. 

Dengan memilih tipe data yang sesuai, kita dapat memaksimalkan efisiensi alokasi memori dari program kita.  
```go
func main() {
	fmt.Println(-30)  
	fmt.Println(10_000_000_000)  
}
```

Kita juga dapat membuat tipe data integer dengan *keyword* `int` saja. Tipe data ini minimal sama dengan `int32` berdasarkan pengecekan sistem operasi yang kita gunakan. Jika menggunakan 32-bit, maka int ini akan sama dengan `int32`. Jika menggunakan 64-bit, maka int ini akan sama dengan `int64`.  
Terdapat juga tipe data `rune` yang merupakan alias dari tipe data `int32`.

### uint  
Tipe data ini adalah kepanjangan dari unsigned integer, yaitu tipe data bilangan bulat yang positif dan tidak mengandung nilai negatif.  
TIPE DATA  | NILAI MINIMUM | NILAI MAKSIMUM  
--|--|--
uint8  | 0 | 255
uint16  | 0 | 65535
uint32  | 0 | 4294967295
uint64  | 0 | 18446744073709551615 
  

Karena tipe data ini cakupannya di awali dengan 0, maka cakupan maksimumnya menjadi lebih besar daripada *signed integer* pasangannya (e.g. nilai maksimum dari `uint8` lebih besar daripada `int8`).  

Kita juga dapat menggunakan *unsigned integer* dengan *keyword* `uint`. Tipe data ini bernilai minimal `uint32`. Jika sistem operasi kita menggunakan menggunakan 64-bit, maka `uint` ini akan sama dengan `uint64`.  

Terdapat juga tipe data `byte` yang merupakan alias dari tipe data `uint8` (range 0 - 255).  

### float  
Ini merupakan tipe data yang bisa kita gunakan untuk menyimpan bilangan desimal. Berikut adalah tipe data *float* yang dapat digunakan:  
TIPE DATA  |NILAI MINIMUM|NILAI MAKSIMUM
--|--|--
float32  |-3.4E+38|3.4E+38
float64  | -1.7E+308| 1.7E+308

```go
func main() {
	fmt.Println(3.14)
	fmt.Println(-2.232321431412)
}
```
>Note: -3.4E+38 artinya minus angka 3.4 dikuti dengan 38 nol di belakangnya. Sebaliknya 3.4E+38 artinya plus angka 3.4 dikuti dengan 38 nol di belakangnya.  

##  Operation  

Kita bisa menggunakan tipe data number dengan operasi matematika sebagai berikut:  

OPERATOR/SYNTAX  |KETERANGAN
--|--
+| Penambahan
-|Pengurangan
\*|Perkalian
/ |Pembagian
%  |Mendapatkan sisa pembagian
Contoh :  
```go
func main() {  
	fmt.Println(1 + 1)
	fmt.Println(2-1)
	fmt.Println(2* 2)
	fmt.Println(4 / 2)
	fmt.Println(10 % 3) // sisa pembagian 10 dengan 3 adalah 1  
}
```
```Output
2  
1  
4  
2  
1  
```

## Augmented assignment  
Ini merupakan operasi yang digunakan untuk mempersingkat operasi matematika.  
OPERATOR/SYNTAX  |OPERASI MATEMATIKA |KETERANGAN
--|--|--
a += 10  | a = a + 10|Penambahan
a -= 10  | a = a - 10  |Pengurangan
a \*= 10  |a = a * 10  |Perkalian
a /= 10  |a = a/ 10  |Pembagian
a %=10  | a = a % 10  |Mendapatkan sisa pembagian
  
Contoh:  
```go
func main() {  
	var a= 10  
	a + 10 // a 10+10 20  
	fmt.Println(a)  
}
```  
```Output
20  
```

## Unary operator  
Operasi ini juga digunakan untuk mempersingkat operasi matematika khusus.  
OPERATOR/SYNTAX|OPERASI MATEMATIKA|KETERANGAN
--|--|--
a++  |a = a + 1  |increment (naik 1 angka) 
a--  | a = a - 1  |decrement (turun 1 angka)  
-a  | a = -a  |mengubah jadi negatif (negasi) 
+a  | a = +a |mengubah jadi positif  

Contoh:  
```go
func main() {  
	var a = 10  
	a++ // a = 10 + 1 = 11  
	fmt.Println(a)
	
	var b = 10  
	fmt.Println(-b)  
}  
```
```output
11  
-10
```
