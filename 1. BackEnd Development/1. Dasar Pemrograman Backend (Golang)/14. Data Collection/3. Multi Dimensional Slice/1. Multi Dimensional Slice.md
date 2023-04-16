# Slice multidimensional

## What is slice multidimensi?
Sama seperti *array* multidimensi, *slice* multidimensi adalah slice yang tiap elemennya juga berupa *slice* (dan bisa seterusnya, tergantung ke dalaman dimensinya). Contoh sederhananya yaitu 2 dimensi:

```go
var numbers = [][]int{{3, 2, 3}, {3, 4, 5}}
fmt.Println(numbers)
```
Output:
```Output
[[3 2 3] [345]]
```

Untuk contoh lebih lengkapnya, kita bisa perhatikan kode di bawah ini:
```go
// multi-dimensional slice
package main
import "fmt"
func main() {
	// Membuat multi-dimensional slice integer
	s1 := [][]int{
	{1, 2},
	{3, 4},
	{5, 6),
	{7, 8},
	}
	
	// Akses multi-dimensional slice integer
	fmt.Println("Slice 1: ", s1)
	// Membuat multi-dimensional slice
	string s2 := [][]string{
		[]string{"str1", "str2"},
		[]string{"str3", "str4"},
		[]string{"str5", "str6"},
	}
	
	// Akses multi-dimensional slice string
	fmt.Println("Slice 2: ", s2)
	
	// Cetak setiap isi slice di s2
	fmt.Println("MultiDimensional Slice 52:")
	for i = 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}
	
	// Cetak elemen pada slice sebagai matrix 
	fmt.Println("Slice 52 Like Matrix:")
	
	// number of rows in s2
	n := len(s2)
	
	// number of columns in $2
	m := len(s2[0])
	fmt.Println("rows: ", n, "cols: ", m)
	
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			fmt.Print(s2[i][j] + " ")
		}
		fmt.Print("\n")
	}
}	
```
Output:
```Output
Slice 1 :  [[1 2] [34] [56] [78]]
Slice 2 :  [[str1 str2] [str3 str4] [str5 str6]]
MultiDimensional Slice s2:
[str1 str2]
[str3 str4]
[str5 str6]
Slice 52 Like Matrix:
rows: 3 cols: 2
str1 str2
str3 str4
str5 str6
```

## Indexing in multidimension
Pada *slice* 2 Dimensi, *index* memiliki 2 tipe yang dikenal sebagai *index* `rows` dan `column`:
![index](https://adyork.github.io/python-oceanography-lesson/fig/netcdf_1D_2D_array.PNG)

Slice 2 Dimensi (Sumber: [adyork.github.io](https://adyork.github.io/python-oceanography-lesson/17-Intro-NetCDF/index.html))

Jadi, ketika kita mengakses index pada tanda kurung siku `[]` yang pertama, kita sedang mengakses *index* `row`, sedangkan ketika kita mengakses kurung siku `[][]` berikutnya, kita sedang mengakses index `col`.

Contoh:
```go
// Membuat slice dengan 2 rows dan 3 cols
var numbers = [][]int{
	{3, 2, 3},
	{3, 4, 5},
}

fmt.Println("row: ", numbers [1]) // Cetak nilai pada row
fmt.Println("column: ", numbers[1][1]) // Cetak nilai pada col
```
Output:
```Output
row: [345]
column: 4
```

Selain itu ada juga slice 3 Dimensi:

![index](https://polinema-programming.github.io/12/images/12-03.png)
Slice 3 Dimensi (Sumber: [polinema-programming.github.io](https://polinema-programming.github.io/12/jobsheet12.html))

Untuk implementasinya kita cukup menambahkan 3 kurung siku `[][][]` pada saat inisialisasi *slice*. 
Contoh:
```go
var numbers = [][][]int{
	{{311, 312, 313), (321, 322, 323}, {331, 332, 333}},
	{{211, 212, 213}, {221, 222, 223}, {231, 232, 233}}, 
	{{111, 112, 113}, {121, 122, 123}, {131, 132, 133}},
}
fmt.Println(numbers)
fmt.Println("Access numbers [2][0][2]: ", numbers[2][0][2]) // 113
fmt.Println("Access numbers[0][1] [3]: ", numbers[0][1][1]) // 322
fmt.Println("Access numbers [1][0][0]: ", numbers[1][0][0]) // 211
```
Output:
```Output
[[[311 312 313] [321 322 323] [331 332 333]] [[211 212 213] [221 222 223] [231 232 233]] [[111 112 113] [121, 122, 123] [131, 132, 133]
Access numbers[2][0][2]: 113
Access numbers[0][1][1]: 322
Access numbers[1][0][0]: 211
```

Sebenarnya untuk *slice* multidimensi ini tidak terbatas sampai 3 Dimensi, bisa lebih banyak lagi dimensinya dan tidak terbatas. Untuk mengakses nilainya tinggal disesuaikan dengan *index target*. Contoh:
```go
var numbers = [][][][][]int{{{{{1}}}}}
fmt.Println(numbers)
// Mengakses nilai
fmt.Println("Access numbers[0][0][0][0][0]: ", numbers[0][0][0][0][0])
```
Output:
```Output
[[[[[1]]]]]
Access numbers[0][0][0][0][0]: 1
```


## Update data
Cara melakukan *update* data pada *slice* multidimensi cukup mudah yaitu, kita akses index rows pada dimensi pertama, dan akses index `col` pada dimensi berikutnya sampai ketemu nilai yang ingin diubah, lalu ubah nilainya dengan operator assign (`=`). Contoh di bawah ini akan merubah angka 5 di akhir *slice* 2 dimensi menjadi 10:
```go
var numbers = [][]int{{3, 2, 3}, {3, 4, 5}}
fmt.Println("numbers: ", numbers)
numbers[1][2] = 10
fmt.Println("update numbers", numbers)
```
Output:
```
numbers: [[3 2 3] [345]]
update numbers [[3 2 3] [3 4 10]]
```
