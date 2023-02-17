# String  
## What is string?  
Tipe data *string* adalah tipe data yang berisi kumpulan karakter. *String* di Golang direpresentasikan dengan keyword `string`. Nilai data *string* di Golang di awali dengan karakter `"` (petik dua) dan di akhiri dengan" (petik dua).  

```go
"Hello World"  
"2134 !_+@&#*()"  
```
Golang juga mendukung tipe data string yang dapat dibuat lebih dari 1 baris (multi line) yang di awali dengan karakter \` (backtick) dan di akhiri juga dengan \` (backtick).  
```go
`
This is string  
With multi-line  
Hello Go!  
`
```

## Basic function in String  
#### `len`  
`len` adalah fungsi yang ada di Golang untuk menghitung jumlah karakter string. Cara penggunaannya cukup sederhana, yaitu memanggil fungsi `len` dengan parameter yang diisi dengan *string*. Fungsi ini akan mengembalikan angka yang berisi total karakter dari string yang diberikan.  
```go
func main() {  
}  
fmt.Println(len("Hello World"))  
```
```output
11 # total character  
```

#### Get character  
Kita juga dapat mengambil karakter tertentu di string dengan cara mengambil posisi index dari karakter yang kita inginkan. Caranya adalah menggunakan sintaks `[<index>]` di belakang string. Urutan index dimulai dari 0.  
![[th 1.webp]]

Adapun Urutan menurun menggunakan index negatif dimulai dari -1 pada akhir string 

![[th 3.webp]]

Hasilnya akan berupa data `byte` yang berisi ASCII code yang mewakili karakter tersebut. ASCII code adalah representasi angka dari karakter alfabet.  


```go
func main() {  
}  
fmt.Println("Golang"[0])  
fmt.Println("Golang" [2])  
```
```Output
71 # output is ASCII code of "G"  
108 # output is ASCII code of "1"  
```

Contoh di atas adalah output dari *ASCII* code karakter "G" (71) dan "I" (108) yang ada di *string* "Golang".  

Jika ingin menampilkannya kembali dalam bentuk *string*, kita bisa menggunakan fungsi `string` dengan parameter yang diisi `byte` data.  

```go
func main() {  
fmt.Println(string("Golang "[0]))  
}  
```
```Output
G
``` 

Pastikan bahwa ketika ingin mendapatkan sebagian karakter, tidak melebihi dari total karakter yang ada. Jika melebihi, maka akan terjadi error.  

```go
func main() {  
}  
fmt.Println("Golang" [10])  
```
```Output  
# command-line-arguments  
./main.go: 6:21: invalid string index 10 (out of bounds for 6-byte string)  
```

Contoh di atas adalah bentuk *error* karena kita mencoba mengambil karakter di index ke-10 sedangkan *string* "Golang" hanya memiliki 6 karakter.  
