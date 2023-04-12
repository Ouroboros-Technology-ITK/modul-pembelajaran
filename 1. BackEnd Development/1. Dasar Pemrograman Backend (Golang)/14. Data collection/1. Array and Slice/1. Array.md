# Array

## What is array?
Dalam membuat program, terkadang kita perlu menyimpan kumpulan data yang sejenis. Contohnya adalah daftar mahasiswa, dari pada kita membuat satu-satu seperti ini:
```go
var student1 string = "Aditira"
var student2 string = "Dito"
var student3 string = "Afis"
```

Lebih baik kita menggunakan Array:
```go
var students = [3]string{"Aditira", "Dito", "Afis"}
```

Jadi apa itu Array?🤔 Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama dan disimpan dalam sebuah variabel.

## Array in go
Dalam bahasa Go, Array dibuat dengan dua cara:

* Array dibuat menggunakan keyword `var` diikuti dengan nama array, panjang array dan tipe datanya.
  Syntaxnya adalah:
```go  
var array_name [length]data_type
```
Contoh:
```go
var myarr [3]string
```
![index](https://media.geeksforgeeks.org/wp-content/uploads/20190704101725/Untitled-Diagram34.jpg)

Array syntax (sumber: [media geeksforgeeks.org](https://www.geeksforgeeks.org/arrays-in-go/))

* Array juga dapat dideklarasikan menggunakan `:=` . Ini lebih fleksibel dari pada deklarasi di atas. *Syntaxnya* adalah:
```go
array_name := [length]data_type{item1, item2, item3,...itenN}
```

* Contoh:
```go
myarr := [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}
```
![index](https://media.geeksforgeeks.org/wp-content/uploads/20190704102210/Untitled-Diagram35.jpg)

Array syntax steno (sumber: [geekforgeeks.org](https://www.geeksforgeeks.org/arrays-in-go/))

Pada Golang, Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatan, menjadi elemen/data yang disimpan jumlahnya tidak bisa melebihi dari yang sudah dialokasikan.

Default nilai tiap elemen array pada awalnya tergantung dari tipe datanya. Jika `int` maka tiap element *Zero value*-nya adalah `0` ,jika `bool` maka `false` ,dan seterusnya.

Contoh:
```go
arrInteger := [4]int{10, 16}
fmt.Println(arrInteger)

arrBoolean := [4]bool{false, true}

fmt.Println(arrBoolean)
```
```
[10 16 0 0]
[false true false false]
```

Perhatikan bahwa nilai dari 2 index array terakhir berisi default value.

## Indexing
![index](https://media.geeksforgeeks.org/wp-content/uploads/20190704100840/Untitled-Diagram33.jpg)

Array syntax Indexing (sumber: [geekforgeeks.org](https://www.geeksforgeeks.org/arrays-in-go/))

Setiap elemen array memiliki indeks berupa angka yang merepresentasikan posisi urutan elemen tersebut. Indeks array dimulai dari `0` . Ini dapat digambarkan dengan tabel berikut:


INDEX |DATA
--|--
0 |Aditira
1| Dito
2 |Tegar
3 |Afis

Contoh di Golang:
```go
func main() {
	var names [4]string
	names[0] = "Aditira"
	names[1] = "Dito"
	names[2] = "Tegar
	names[3] = "Afis"
}
```
Variabel `names` dideklarasikan sebagai *array string* dengan alokasi elemen 4 slot. Cara mengisi slot elemen kita bisa langsung mengakses elemen menggunakan indeks, lalu mengisinya.

Pengisian elemen array pada indeks yang tidak sesuai dengan alokasi akan menghasilkan error. Contoh: jika array memiliki 2 slot, maka pengisian nilai slot 3 dan seterusnya akan tidak valid.
```go
func main() {
	var names [2]string
	names[0] = "Aditira"
	names[1] = "Dito"
	names[3] = "Tegar"
}
```
Output:
```Output
/main.go:7:8: invalid argument: array index 3 out of bounds [0:2]
```


## Insert
Ada beberapa cara untuk inisialisasi nilai awal pada Array yaitu
* Horizontal
* Vertical
* Tanpa Jumlah Elemen
* Multidimensi
* Menggunakan Keyword `make`

### Horizontal
Caranya dengan menuliskan data elemen dalam kurung kurawal setelah tipe data, dengan pembatas antar elemen adalah tanda koma (,). Contoh:
```go
var fruits = [4]string{"apel", "anggur", "pisang", "melon"}
fat.PrintIn("Jumlah element \t\t:", len(fruits))
fmt.Println("Isi semua element \t:", fruits)
```
Output:
```Output
Jumlah element		: 4 
Isi semua element	: [apel anggur pisang melon]
```

### Vertical
Bisa juga ditulis dalam bentuk *vertical* seperti contoh di bawah ini:

```go
var fruits = [4])string{
	"apel",
	"anggur",
	"pisang",
	"melon",
}
```
Khusus pada vertical, tanda koma waijib dituliskan setelah elemen, **termasuk elemen terakhir**. Jika tidak, maka akan *error*.

### Tanpa Jumlah Elemen
Membuat Array diperbolehkan untuk tidak menuliskan jumlah lebar array-nya, cukup ganti dengan tanda 3 titik (...). Jumlah elemen akan di kalkulasi secara otomatis menyesuaikan data elemen yang diisikan.
```go
var numbers = [...]int{2, 3, 2, 4, 3}
fmt.Println("data array \t:", numbers)
fmt.Println("jumlah elemen \t:", len(numbers))
```
Output:
```go
data array	: [123243]
jumlah elemen	: 5
```

Variabel `numbers` akan secara otomatis memiliki jumlah elemen 5, karena pada saat deklarasi disiapkan 5 buah elemen.

### Multidimensi

Array multidimensi adalah array yang tiap elemennya juga berupa array (dan bisa seterusnya, tergantung ke dalaman dimensinya).

Cara deklarasi array multidimensi secara umum sama dengan cara deklarasi array biasa, dengan cara menuliskan data array dimensi selanjutnya sebagai elemen array dimensi sebelumnya.

Khusus untuk array yang merupakan sub dimensi atau elemen, boleh tidak dituliskan jumlah datanya.

Contohnya:
```go
var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
var numbers2 = [2][3]int{{3, 2, 3}, {3, 4 5}}
  
fmt.Println("numbers1", numbers1)
fmt.Println("numbers2", numbers2)
```
Output:
```Output
numbers1: [[3 2 3] [3 4 5]]
numbers2:  [[3 2 3] [3 4 5]]
```

Kedua array di atas memiliki elemen yang sama:
* Variabel `numbers1` menggunakan cara deklarasi array biasa.
* Variabel `numbers2` menggunakan cara khusus, di mana kita tidak mendefinisikan jumlah data (*length*) dari array sub elemen,

### Menggunakan Keyword `make`

Deklarasi sekaligus alokasi data array juga bisa dilakukan lewat keyword `make` . *Syntaxnya* adalah:
```go
make ([]data_type, length)
```

Contoh:
```go
var fruits = make([]string, 2)
fruits[0] = "apple"
fruits[1] = "manggo"
fmt.Println(fruits) // [apple mango]
```

Parameter pertama keyword `make` diisi dengan tipe data elemen array yang diinginkan, parameter kedua adalah jumlah elemennya. Pada kode di atas, variabel `fruits` tercetak sebagai *array string*
dengan alokasi 2 slot.

## Read

Untuk membaca semua data pada Array, kita bisa akses/print variabel yang mewakili array tersebut,
Contohnya:
```go
var fruits = [4]string{}"spel", "anggur", "pisang", "melon"}
fmt.Println(fruits)
```

Penggunaan fungsi `fmt.Println()` pada data array tanpa mengakses indeks tertentu, akan menghasilkan *output* dari semua array yang ada.

Kita juga bisa mengakses salah satu elemen Array dengan menunjuk posisi *index*-nya dengan pola `array[index]`. Contoh:
```go
var fruits = [4]string{}"apel", "anggur", "pisang", "melon"}
fmt.Println(fruits(1]) // anggur
fmt.Println(fruits(3]) // melon
```

*Outputnya* adalah nilai dari elemen yang ditunjuk *index*-nya.

## Update
Untuk melakukan modifikasi pada Array, kita bisa mengakses index tertentu lalu mengubahnya dengan data yang kita inginkan dengan pola `array[index] = update_value`.Contoh:
```go
var fruits = [4]string{"apel", "durian", "pisang", "melon"}

fruits[2] = "anggur" //mengubah pisang menjadi anggur

fmt.Printin(fruits)
```
```Output
apel durian anggur melon
```

## Array Functions

Beberapa *function* yang sudah kita bahas di atas disebut sebagai Array Functions.

OPERASI | KETERANGAN
--|--
`Len(array)` |Untuk mendapatkan panjang Array
`array[index]` | Mendapat data di posisi index
`array[index]=value` |Mengubah data di posisi index
`make([]data_type, length)` |Membuat array baru

## Looping
*Keyword* `for` dan array memiliki hubungan yang sangat erat. Dengan memanfaatkan perulangan menggunakan *keyword* ini, kita bisa mengakses elemen-elemen dalam array.

### Perulangan Elemen Array Menggunakan Keyword for

Kita bisa memanfaatkan iterasi perulangan untuk mengakses elemen berdasarkan indeks-nya. Contohnya:
```go
for fruits = [4]string{"apel", "durian", "pisang", "melon"}

for i := 0; i < len(fruits); i++ {
	fmt.Printf("elemen %d : %s \n", i, fruits[i])
}
```
```Output
elemen 0 : apel
elemen 1 : durian
elemen 2 : pisang
elemen 3 : melon
```
Perulangan di atas dijalankan sebanyak jumlah elemen array `fruits` (bisa diketahui dari kondisi `i < len(fruits)` . Di tiap perulangan, elemen array diakses lewat variabel iterasi `i`.

### Perulangan Elemen Array Menggunakan Keyword `for` - `range`
Ada cara yang lebih sederhana me-looping data array, dengan menggunakan *keyword* `for` - `range`. 
Contohnya:
```go
var fruits = [4]string{"apel", "durian", "pisang", "melon"}

for i, fruit := range fruits {
	fmt.Printf("elemen %d : %s\n", i, fruit)
}
```
Output:
```
elemen 0 : apel
elemen 1 : durian
elemen 2 : pisang
elemen 3 : melon
```
Array `fruits` diambil elemen-nya secara berurutan. i tiap elemen ditampung oleh variabel `fruit` (tanpa huruf s), sedangkan indeks ditampung variabel `i`.

*Output* program di atas, sama dengan *output* program sebelumnya, hanya cara yang digunakan berbeda.

### Penggunaan Variabel Underscore `_` Dalam `for` - `range`

Kadang kala ketika looping menggunakan `for` - `range`, ada kemungkinan di mana data yang dibutuhkan adalah elemen-nya saja, indeks-nya tidak. Sedangkan kode di atas, `range` mengembalikan 2 data, yaitu indeks dan elemen.

Seperti yang sudah diketahui, bahwa di Go tidak memperbolehkan adanya variabel yang menganggur atau tidak dipakai. Contoh code di bawah akan *error*:
```go
var fruits = [4]string{"apel", "durian", "pisang", "melon"}

for i, fruit := range fruits {
	fmt.Printf("nama buah : %s\n", fruit)
}
```
Output:
```Output
./Main.go:8:6: i declared but not used
```

Di sinilah salah satu kegunaan variabel pengangguran, atau underscore ( `_` ). Tampung saja nilai yang tidak ingin digunakan ke *underscore*.
```go
var fruits = [4]string{"apel", "durian", "pisang", "melon"}

for _, fruit := range fruits {
	fmt.Printf("nama buah : %s\n", fruit)
}
```
Output:
```Output
elemen 0 : apel
elenen 1 : durian
elenen 2 : pisang
elenen 3 : melon
```
Pada kode di atas, yang sebelumnya adalah variabel `i` diganti dengan `_`, karena variabel i tidak digunakan.

Jika yang dibutuhkan hanya indeks elemen-nya saja, bisa gunakan 1 buah variabel setelah *keyword* `for`.
```go
for i,  _ := range fruits { }
// atau
for i := range fruits { }
```
