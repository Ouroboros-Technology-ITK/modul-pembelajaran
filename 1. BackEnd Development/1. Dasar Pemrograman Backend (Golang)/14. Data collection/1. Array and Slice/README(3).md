# Slice

## What is Slice?
**Slice** adalah potongan dari duh Array (*reference elemen array*). Slice mirip dengan Array, yang membedakan adalah ukurannya bisa berubah. Slice dan Array selalu terkoneksi, di mana Slice bisa mengakses sebagian atau seluruh data Array.

Slice memiliki 3 tipe data yang berbeda, yaitu:
* **Pointer**: penunjuk data pertama di Array pada Slice
* **Length**: panjang dari Slice
* **Capacity**: kapasitas dari Slice, di mana *length* tidak boleh lebih dari *capacity*

Kita bisa membuat Slice dari Array dengan mengikuti pola berikut ini:
MEMBUAT SLICE | KETERANGAN
--|--
array[low:high] | Membuat Slice dari Array dimulai index low sampai index sebelum high
array[low:] | Membuat Slice dari Array dimulai index low sampai index akhir di Array
array[:high]|Membuat Slice dari Array dimulai index 0 sampai index sebelum high
array[:]|Membuat Slice dari Array dimulai index O sampal index akhir di Array FOODDS

Mari kita bahas semua dengan bantuan sebuah contoh:
```go
// Membuat array
arr := [7]string{"This", "is", "the", "tutorial", "of", "Go", "language"} // Menampilkan array
fmt.Println("Array: ", arr)

// Membuat slice
myslice = arr[1:6]

// Menampilkan slice
fmt.Println("Slice:", myslice)

// Menampilkan panjang slice
fmt.Println("Length Slice: %d", len(myslice))

// Menampilkan capacity slice
fmt.Println("\nCapacity Slice: %d", cap(myslice))
```
Output:
```Output

Array			: [This is the tutorial of Go language]
Slice			: [is the tutorial of Go]
Length Slice	: 5
Capacity Slice	: 6
```

Di sini *pointer* `index low` adalah 1 sehingga Slice mulai mengakses elemen dari indeks 1 sampai indeks 5, karena `index high` adalah 6 (*index sebelum 6 adalah 5*). Panjang Slice adalah 5, yang berarti jumlah total elemen yang ada dalam Slice adalah 5 dan kapasitas Slice 6 berarti dapat menyimpan maksimal 6 elemen.

![index](https://media.geeksforgeeks.org/wp-content/uploads/20190714183310/Untitled-Diagram38.jpg)

Slice of Array (sumber: [geeksforgeeks.org](https://www.geeksforgeeks.org/slices-in-golang

Ada beberapa fungsi pada Slice yang dapat kita gunakan yaitu:
FUNGSI|KETERANGAN
--|--
`make([]TypeData, length, capacity)`| Membuat Slice baru
`len(slice)`|Untuk mendapatkan panjang
`cap(slice)`|Untuk mendapatkan kapasitas
`append(slice, data)`| membuat slice baru dengan menambah data ke posisi terahir slice,
`copy (destination, source)`| Menyalin Slice dari *source* ke *destination*
 
Cara pembuatan Slice mirip seperti pembuatan Array, bedanya tidak perlu mendefinisikan jumlah elemen ketika awal deklarasi. Pengaksesan nilai elemen-nya juga sama. Contoh pembuatan Slice:
```go
var students = []string{"Aditira", "Gina", "Juno", "Dito"}
fmt.Println(fruits[0]) // "Aditira"
```
Kita juga dapat membuat Slice baru menggunakan fungsi m`ake([]type, length, capacity)`. Contoh:
```go
myslicel := make([]int, 5, 10)
fmt.Printf("myslice1 %v\n", myslicel) // myslicel = [0 0 0 0 0]
fmt.Printf("length = %d\n", len(myslice1)) // length = 5
fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 10
// with omitted capacity
myslice2 := make([]int, 5)
fmt.Printf("myslice2 = %v\n", myslice2) // myslice2 = [0 0 0 0 0]
fmt.Printf("length = %d\n", len(myslice2)) // length = 5
fmt.Printf("capacity = %d\n", cap(myslice2)) // capacity = 5
```
>Note: Jika parameter *capacity* tidak ditentukan, akan diset sama dengan *length*

## Append
Fungsi `append()` digunakan untuk menambahkan elemen pada Slice. Elemen baru tersebut diposisikan setelah indeks paling akhir. Nilai balik fungsi ini adalah Slice yang sudah ditambahkan nilai barunya. Contoh:
```go
var students = []string{"Aditira", "Dito"}
var appendStudents = append(students, "Ria")

fmt.Println(students) // ["Aditira", "Dito"]
fmt.Println(appendStudents) // ["Aditira", "Dito", "Ria"]
```

Ada 2 hal yang perlu diketahui dalam penggunaan fungsi ini, yaitu:
* Ketika jumlah elemen dan lebar kapasitas adalah sama `(len(students) == cap(students))`, maka elemen baru hasil `append()` merupakan referensi baru.
* Ketika jumlah elemen lebih kecil dibanding kapasitas`(len(students) < cap(students))`, elemen baru tersebut ditempatkan ke dalam cakupan kapasitas, menjadikan semua elemen slice lain yang referensi-nya sama akan berubah nilainya. Agar lebih jelas silakan perhatikan contoh berikut. 
```go
var students = []string{"Aditira", "Gina", "Dito"}
var bStudents students [0:2]

fmt.Println(cap (bStudents)) // 3
fmt.Println(len (bStudents)) // 2

fmt.Println(students) //["Aditira", "Gina", "Dito"]
fmt.Println(bStudents) // ["Aditira", "Gina"]

var cStudents = append(bStudents, "Juno")

fmt.Println(students) // ["Aditira", "Gina", "Juno"]
fmt.Println(bStudents) // ["Aditira", "Gina"] 
fmt.Println(cStudents) // ["Aditira", "Gina", "Juno"]
```

Terlihat bahwa elemen indeks ke-2 slice `students` nilainya berubah setelah ada penggunaan *keyword* `append()` pada `bStudents`. Slice `bStudents` kapasitasnya adalah **3** sedang jumlah datanya hanya **2**. Karena `len (bStudents) < cap (bStudents)`, maka elemen baru yang dihasilkan, terdeteksi sebagai perubahan nilai pada referensi yang lama (referensi elemen indeks ke-2 slice `students`), membuat elemen yang referensinya sama, nilainya berubah.

## Copy
Untuk meningkatkan kapasitas dari sebuah Slice kita harus membuat Slice baru yang lebih besar dan menyalin isi Slice asli ke dalamnya. Jadi kita bisa gunakan fungsi `copy(destination, source)` untuk menyalin Slice dari *source* ke *destination*:
```go
days := make([]string, 2, 2)
days[0] = "Senin"
days[1] = "Selasa"
copyDays := make([]string, len(days), (cap(days)+1)*2) // +1 seandainya cap (days) == 0
copy(copyDays, days) // Copy slice days ke copyDays
fmt.Println("days\t\t: ", days)
fmt.Println("cap\t\t: ", cap(days))
fmt.Println("copyDays\t: ", copyDays)
fmt.Println("cap\t\t: ", cap(copyDays))
copyDays [1]= "Rabu" // mengubah variable copyDays tidak akan berpengaruh ke variable days (copy by val
fmt.Println("-------------------")
fmt.Println("update days\t: ", days)
fmt.Println("update copyDays\t: ", copyDays)
4
```
Output:
```Output

days			: [Senin Selasa]
cap				: 2
copyDays		: [Senin Selasa]
cap				: 6
-------------------
update days		: [Senin Selasa]
update copyDays	: [Senin Rabu]
```

Fungsi `copy()` juga mengembalikan jumlah elemen yang berhasil di-*copy*.

```Go
copiedElemen := copy(copyDays, days)
fmt.Println(copiedElemen) // 2
```

>Note: Untuk melakukan *copy* pastikan ukuran Slice *destination* **sama** atau lebih **agar** data tidak terpotong.

## Concat
Untuk mengggabungkan Slice, kita bisa memanfaatkan fungsi `append()`. Contohnya bisa diperhatikan kode di bawah ini:

```go
students1 := []string{"Aditira", "Dito", "Afis", "Eddy"}
students2 := []string{"Imam", "Raam", "Zein"}
school := append(students1, students2...)
fmt.Println(school) // [Aditira Dito Afis Eddy Imam Raam Zein]
```

Operator tiga titik (`...`) setelah argumen kedua diperlukan karena `append()` merupakan fungsi variadik yang menerima argumen dalam jumlah tak terbatas.

```go
school := append(students1, students2...)
```

Baris di atas pada dasarnya adalah sintaks singkatan untuk memasukkan setiap elemen students2 secara manual ke dalam fungsi append() sebagai berikut:

```go
school := append(students1, "Imam", "Raam", "Zein")
```

Untuk menggabungkan Slice lagi kita bisa gunakan lagi fungsi append() dari penggabungan Slice sebelumnya, contoh:
```go
students1 := []string{"Aditira", "Dito", "Afis", "Eddy"}
students2 := []string{"Imam", "Raam", "Zein"}
school := append(students1, students2...)
fmt.Println(school)
students3 := []string{"Dion", "Ronal", "Tom"}
school2 := append(school, students3...) 
fmt.Println(schoo12)
```
Output:
```Output
[Aditira Dito Afis Eddy Imam Raam Zein]
[Aditira Dito Afis Eddy Imam Raam Zein Dion Ronal Tom]
```

## Slicing
Slicing adalah cara untuk mengakses bagian data array yang akan menghasilkan **length** dan **capacity** selain data Slicing itu sendiri, contohnya:
```go
var fruits = [4]string{"apple", "grape", "banana", "melon"}
fmt.Println("output: ", fruits) // output: [apple grape banana melon]
fmt.Println("len: ", len(fruits)) // len: 4
fmt.Println("cap: ", cap(fruits)) // cap: 4

var aFruits = fruits[0:3]
fmt.Println("output:", aFruits) // output: [apple grape banana]
fmt.Println("len: ", len(aFruits)) // len: 3
fmt.Println("cap: ", cap(aFruits)) // cap: 4

var bFruits fruits [1:4] fmt.Println("output: ", bFruits) // output: [grape banana melon]
fmt.Println("len: ", len(bFruits)) // len: 3
fmt.Println("cap: ", cap(bFruits)) // cap: 3
```

Variabel `fruits` disiapkan di awal dengan jumlah elemen 4, fungsi `len(fruits)` dan `cap(fruits)` pasti hasilnya 4.

Variabel `aFruits` dan `bFruits` merupakan Slice baru berisikan 3 buah elemen milik Slice `fruits`. Variabel `aFruits` mengambil elemen index 0, 1, 2 sedangkan `bFruits` 1, 2, 3.
Fungsi `len()` menghasilkan angka 3, karena jumlah elemen kedua slice ini adalah 3. Tetapi `cap(aFruits)` menghasilkan angka yang berbeda, yaitu 4 untuk `aFruits` dan 3 untuk `bFruits`. Kenapa?🤔 jawabannya bisa dilihat pada tabel berikut.

CODE|OUTPUT|`LEN()`|`CAP()`
--|--|--|--
fruits[0:4]|[ buah buah buah buah ]|4|4
aFruits[0:3] |[ buah buah buah ----]|3|4
bfruits[1:4]|---- [ buah buah buah ]|3|3

Kita analogikan Slicing 2 index menggunakan **x** dan **y**.
```Output
fruits[x:y]
```

**Slicing** yang dimulai dari indeks **0** hingga **y** akan mengembalikan elemen-elemen mulai indeks **0** hingga sebelum indeks **y**, dengan lebar kapasitas adalah sama dengan Slice aslinya.

>Sedangkan Slicing yang dimulai dari indeks x, yang di mana nilai x adalah lebih dari 0, membuat elemen ke-x slice yang diambil menjadi elemen ke-0 Slice baru. Hal inilah yang membuat kapasitas Slice berubah.

## Array vs slice

Salah satu perbedaan Slice dan Array bisa diketahui pada saat deklarasi variabel-nya, jika jumlah elemen tidak dituliskan, maka variabel tersebut adalah Slice.
```go
var studentsA = []string{"Aditira", "Gina"} // slice
var studentsB = [2]string{"Gina", "Juno"} // array
var studentsC = [...]string{"Juno", "Dito"} // array
```

Kalau perbedannya hanya di penentuan alokasi pada saat inisialisasi, kenapa tidak menggunakan satu istilah saja? 🤔 atau adakah perbedaan lainnya? 🤔


Sebenarnya Slice dan Array tidak bisa dibedakan karena merupakan sebuah kesatuan. Array adalah kumpulan nilai atau elemen, sedangkan Slice adalah referensi tiap elemen tersebut.

Slice bisa dibentuk dari Array yang sudah didefinisikan, caranya dengan memanfaatkan teknik 2 indeks untuk mengambil elemen-nya:
```go
var students = [4]string{"Aditira", "Dito", "Gina", "Eddy"}
var sliceStudents = fruits[0:2]

fmt.Println(sliceStudents) // ["Aditira", "Dito"]
```

Ketika mengakses elemen Array menggunakan satu buah indeks (seperti `data[2]`), nilai yang didapat merupakan hasil **copy** dari referensi aslinya. Berbeda dengan pengaksesan elemen menggunakan 2 indeks (seperti `data[0:2]`), nilai yang didapat adalah reference elemen atau Slice.

Kita bisa coba beberapa Slice pada tabel di bawah ini menggunakan variabel students di atas:

CODE | OUTPUT STUDENTS | 
--|--
`students [0:3]` [Aditira, Dito, Gina]|semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
`students [0:0]`[]| menghasilkan Slice kosong, karena tidak ada elemen sebelum ind
`students [4:4]` []| menghasilkan Slice kosong, karena tidak ada elemen yang dimulai
`students [4:0]`[]| error, pada penulisan students [a:b] nilai a harus lebih kecil atau :
`students[:]` [Aditira, Dito, Gina, Eddy] |semua elemen
`students[2:]` [Gina, Eddy]|semua elemen mulai indeks ke-2
`students[:2]` [Aditira, Dito]|semua elemen hingga sebelum indeks ke-2

Jadi, kapan kita menggunakan Array atau Slice? 🤔 Berikut adalah beberapa tips untuk memutuskannya:
* Jika entitas dideskripsikan dengan sekumpulan item dan panjangnya tetap, maka **gunakan Array**.
* Jika menjelaskan koleksi umum yang akan ditambahkan atau dihapus elemennya, maka **gunakan Slice**.
* Apakah kita akan memodifikasi koleksi dengan cara tertentu? Jika ya, maka **gunakan Slice**.
