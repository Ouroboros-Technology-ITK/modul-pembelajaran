
# Condition  

## Comparison operator  
Operasi perbandingan adalah operasi yang menghasilkan nilai *boolean* (benar atau salah). Jika hasil operasinya benar, maka nilainya adalah `true`. Sebaliknya, jika salah akan bernilai `false`.  

Berikut beberapa operator perbandingan di Golang :  

OPERATOR  |KETERANGAN  
--|--
\> |Lebih dari  
< | Kurang dari  
\>= | Lebih besar atau sama dengan  
<= | Lebih kecil atau sama dengan  
== | Sama dengan  
!= | Tidak sama dengan  

```go
func main() {  
	var namel, name2= "John", "John"  
	var result = name1 == name2  
	var result2 = 10 = 2
	  
	fmt.Println(result)  
	fmt.Println(result2)  
}  
```
```Output
true
false  
```

Di atas adalah contoh operasi perbandingan sederhana. `result` akan bernilai `true` karena nilai `name1` sama dengan nilai `name2`. `result2` akan bernilai `false` karena nilai `10` tidak lebih kecil dari nilai `2`.

## Boolean / logic operator  

Operasi ini adalah operasi yang hanya bisa digunakan untuk tipe data *boolean* dengan fungsi sebagai pembanding dua nilai *boolean* atau lebih.  

OPERATOR  |KETERANGAN
--|--
`&&`   |And / Dan
\|\| | Or / Atau
`!`|Not / Bukan / Kebalikan  

Berikut hasil operasi dengan menggunakan Operator `&&` :  
NILAI 1  |NILAI 2  | HASIL
--|--|--
true  | true|true
true  |false|false
false  |true|false
false  |false|false

Berikut hasil operasi dengan menggunakan Operator || :  
NILAI 1  |NILAI 2  |HASIL
--|--|--
true  |true|true
true  |false|true
false  |true|true
false  |false|false

Khusus untuk operasi `!`, kita bisa gunakan untuk membalikkan nilai boolean.  
NILAI  |HASIL
--|--
true  |false
false  |true

```go
func main() {  
	var score, attendance = 90, 70  
	var graduated score > 80 && attendance > 80  
	fmt.Println(graduated)  
}  
```
```Output
true  
```

Di atas adalah contoh operasi boolean dengan menggunakan Operator `&&`. `graduated` akan bernilai `true` karena nilai `score` lebih besar dari 80 dan nilai `attendance` lebih besar dari 70.

## if  
Golang memiliki sebuah *syntax* `if` yang bisa digunakan untuk menentukan apakah kondisi 
tertentu akan dijalankan atau tidak.  

![index](https://cdn.datamentor.io/wp-content/uploads/2017/11/r-if-statement.jpg)
(Sumber = [flowchartdesigns.blogspot.com](https://flowchartdesigns.blogspot.com/2018/11/flowchart-for-if-condition.html))

Caranya cukup menulis syntax `if` diikuti dengan *condition statement* kemudian diikuti dengan *block statement* `{}` yang akan dijalankan jika kondisi terpenuhi.  

```go
if condition {  
}    
```

Contoh:  
```go
func main() {  
	var name = "Budi"  
	if name == "Budi" {  
		fmt.Println("Halo, Budi")
	}
}
```
```Output
Halo, Budi  
```

Contoh di atas adalah kondisi jika nilai `name` sama dengan `Budi`, maka akan menjalankan fungsi `Println` dengan output `Halo`, `Budi`.  

## `if` with short statement  

*Condition* di Golang dapat menggunakan *short syntax* untuk mendefinisikan suatu variabel yang digunakan pada sebuah kondisi.  

Cara penulisannya, cukup memisahkan antara *short syntax* dan *condition statement* di dalam if dengan tanda titik koma `;` . 
```go
if short_syntax; condition_statement {  
	// code to be execute if condition is true  
}  
```

Contoh:  
```go
func main() {  
	if x := 75; x > 70 {  
		fmt.Println("Selamat, anda lulus")  
	}  
}  
```
```Output
Selamat, anda lulus  
```

Contoh di atas adalah bentuk short syntax untuk mendefinisikan variabel `x` dengan nilai `75`. Setelah itu, kita bisa menggunakan variabel `x` pada kondisi.

Kita dapat membuat sebanyak apapun variabel yang akan digunakan pada kondisi.
```go
func main() {  
	if x, y := 75, 80; x > 70 && y > 70 {  
		fmt.Println("Selamat, anda lulus")  
		
		fmt.Println("Nilai rata-rata: ", (x + y) / 2)  
	}  
}
```
```Output
Selamat, anda lulus  
Nilai rata-rata: 77  
```

## else  
`else` adalah program yang akan dijalankan jika kondisi di dalam `if` tidak terpenuhi (false).  
```go
if condition {  
	// code to be execute if condition is true  
} else {  
	// code to be execute if condition is false  
}
``` 

Contoh:  
```go
func main() {  
	var name = "Anonim"  
	if name == "Budi" {  
		fmt.Println("Hello Budi")  
	} else {  
		fmt.Println("Helo, boleh kenalan ?")  
	}  
}
```  
```Output
Helo, boleh kenalan ?  
```
Contoh di atas adalah kondisi jika nilai `name` tidak sama dengan `Budi`, maka akan menjalankan operasi di block statement `else` berupa fungsi `Println` dengan output `Helo, boleh kenalan ?`. 

## else if

Kadang di dalam *condition* kita ingin membuat beberapa kondisi dengan setiap kondisi memiliki perintah masing-masing yang akan dijalankan. Kita dapat membuat beberapa kondisi dengan menggunakan `else if` sebanyak yang diinginkan.  
```go
if condition1 {  
	// code to be executed if condition1 is true  
} else if condition2 {  
	// code to be executed if condition1 is false and condition2 is true  
} else {  
	// code to be executed if condition1 and condition2 are both false  
}  
```

Contoh:  
```go
func main() {  
	var score = 81  
	if score == 100 {  
		fmt.Println("Perfect")  
	} else if score >= 90 && score < 100 {  
		fmt.Println("Fantastic")  
	} else if score >= 80 && score < 90 {  
		fmt.Println("Good")  
	} else if score >= 70 && score < 80 {  
		fmt.Println("Graduated")  
	} else {  
		fmt.Println("Failed")  
	}  
}
```  
```Go
Good  
```

Contoh di atas adalah kondisi jika nilai `score` lebih besar atau sama dengan `80`, maka akan menjalankan operasi di *block statement* `else if` ke-2 berupa fungsi `Println` dengan output `Good`.  

*Nested condition*

*Condition* bisa berada di dalam *condition* lain, ini yang disebut dengan *nested condition*.  
```go
if condition1 {  
	// code to be executed if condition1 is true  
	if condition2 {  
		// code to be executed if both condition1 and condition2 are true  
	}  
}  
```

Contoh:  
```go
if num%2== 0 {  
	if num % 3 == 0 {  
		fmt.Println("Bilangan genap dan kelipatan 3")  
	} else {  
		fmt.Println("Bilangan genap bukan kelipatan 3")  
	}  
} else {  
	fmt.Println("Bilangan ganjil")  
}  
```

Beberapa contoh `if else` dapat dilihat di sini. 

## switch case  

`switch case` juga merupakan *syntax* yang digunakan untuk pengkondisian, umumnya digunakan untuk memilih salah satu dari banyak blok kode yang akan dieksekusi. 
```go
switch expression {  
case x:  
// code block  
case y:  
// code block  
case z:  
...  
default:  
// code block  
}  
```
Cara kerja dari `switch case` adalah membaca expression (biasanya adalah sebuah variabel berisi nilai yang ingin dibandingkan) yang diberikan, kemudian mencocokkan nilai dari salah satu `case` yang ada dan menjalankan perintah di dalamnya. Terdapat opsi `default` yang bisa digunakan untuk menjalankan kode jika tidak ada `case` yang cocok.  
```go
func main() {  
day := 4  
switch day {  
	case 1:  
		fmt.Println("Senin")  
	case 2:  
		fmt.Println("Selasa")  
	case 3:  
		fmt.Println("Rabu")  
	case 4:  
		fmt.Println("Kamis")  
	case 5:  
		fmt.Println("Jumat")  
	case 6:  
		fmt.Println("Sabtu")  
	case 7:  
		fmt.Println("Minggu")  
	default:  
		fmt.Println("Invalid day")
```
```Output
Kamis  
```

Di Golang, kita dapat mengecek pada beberapa kondisi di dalam 1 `case` dengan cara mendefinisikan  
nilai yang dipisah menggunakan `,` (koma).  

```go
func main() {  
	day := 4  
	switch day {  
		case 1, 2, 3, 4, 5:  
			fmt.Println("Weekday")  
		case 6, 7:  
			fmt.Println("Weekend")  
		default:  
			fmt.Println("Invalid day")
	}
} 

```


## Condition in case  
Kita juga dapat menggunakan condition di dalam `case` untuk menentukan apakah `case` tersebut akan dijalankan atau tidak.  
```go
switch {  
case condition:  
	// code block  
case condition2:  
	// code block  
default:  
	// code block  
}
```  


Contoh:  
```go
func main() {  
	number := 6  
	switch {  
		case number > 10:  
			fmt.Println("High number")  
		case number > 5:  
			fmt.Println("Middle number")  
		case number > 1:  
			fmt.Println("Low number")  
		default:  
			fmt.Println("Zero")  
	}
}
```
```Output 
Middle number  
```


`fallthrough`

Secara *default*, setelah satu *case* terpenuhi maka program akan otomatis keluar dari *switch*.  
Jika kita ingin mengecek pada kondisi yang ada di `case` selanjutnya, bisa menggunakan `fallthrough`.  

Contoh:  
```go
func main() {  
	number := 6  
	switch {  
		case number > 10:  
			fmt.Println("High number")  
			fallthrough // check to the next case  
		case number > 5:  
			fmt.Println("Middle number")  
			fallthrough // check to the next case  
		case number > 1:  
			fmt.Println("Low number")  
			fallthrough // check to the next case  
		default:  
			fmt.Println("Zero") 
	}
}
```
```Output
Middle number  
Low number  
Zero  
```

Beberapa contoh switch case dapat dilihat [di sini](a).
