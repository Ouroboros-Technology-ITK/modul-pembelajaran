# File
# What is I/O
I/O (Input/Output) adalah interaksi yang dilakukan antar sistem komputer dan lingkungan luar. Yang dimaksud "lingkungan luar" adalah manusia atau sistem komputer lainnya. I/O dapat dilakukan dengan berbagai media, seperti keyboard, mouse, printer, dan lain-lain.
Contoh sederhana I/O dengan manusia adalah ketika kita menggunakan aplikasi **Microsoft Word**. Ketika kita mengetikkan kata-kata di *keyboard* menggunakan tombol huruf, **Enter**, **Backspace**, dll. Proses ini disebut *input*. Kemudian komputer akan menampilkan kata-kata yang kita ketikkan di layar monitor. Proses ini disebut *output*.

Contoh I/O antar sistem komputer adalah CCTV. CCTV akan mengambil video dari lingkungan luar dan mengirimkannya ke *server*. Kemudian kita dapat mengakses data di *server* dan menampilkannya melalui tampilan video di komputer.

## `os` package
Golang memiliki beberapa *package* untuk melakukan interaksi I/O, salah satunya adalah *package* `os`. *Package* `os` memiliki fungsi-fungsi untuk membuat dan menghapus file, membaca dan menulis file, membuat direktori, dll.
Kita cukup mengimport package os untuk dapat menggunakan fungsi-fungsi yang ada di dalamnya.
```go
import "os"
```

## Create file and directory
Pembuatan *file* di Golang dapat menggunakan fungsi `os.Create()` dengan memasukkan path (lokasi file tersebut) sebagai parameter. Fungsi ini akan mengembalikan dua nilai, yaitu `*os.File` dan `error`. 
`*os.File` adalah tipe data yang merepresentasikan *file*. error` digunakan` untuk mengembalikan *error*. Jika tidak terjadi *error*, maka nilai *error* akan bernilai `nil`.

Perlu diperhatikan bahwa *file* yang baru terbuat statusnya adalah otomatis *open*, maka dari itu perlu untuk di-*close* menggunakan *method* `Close()` dari `*os.File` setelah file tidak digunakan lagi. Untuk menghindari terlalu banyak *file* yang statusnya *open* yang akan membuat komputer *no response* (macet) sehingga tidak bisa digunakan.
```go
import "os"

func main() {
	file, err := os.Create("test.txt") // membuat file 'test.txt'
	if err != nil {
		panic(err)
		return
	}
	defer file.Close() // close file setelah membuat file

	fmt.Printf("Success create file, %s", file.Name()) // file.Name() untuk mendapatkan nama file, yait
}
```
Output:
```
Success create file, test.txt
```

Jika berhasil, maka akan muncul satu *file* baru dengan nama test.txt di direktori yang sama dengan *file* `main.go`.

```
golang-test
test.txt
_main.go
```
Perlu diingat, fungsi tersebut memiliki **resiko**. Jika ternyata *file* yang akan dibuat sudah ada, maka akan ditimpa dan menghilangkan seluruh data di *file* tersebut.

Kita dapat memastikan terlebih dahulu apakah *file* tersebut sudah dibuat dengan menggunakan `os.Stat()` dengan memasukkan *path* sebagai *parameter*. Fungsi ini akan mengembalikan dua nilai, yaitu `os.FileInfo` dan `error`. Kita akan menggunakan kembalian `error` dari fungsi tersebut dengan memakai fungsi `os.IsNotExist()` untuk memastikan apakah *file* tersebut sudah dibuat atau belum. Jika belum, maka kita akan membuat *file* tersebut.
```go
import (
	"fmt"
	"os"
)

func main() {
	var _, err = os.Stat("test.txt") // mengecek apakah file sudah ada atau belum
	if os.IsNotExist(err) {          // jika file belum ada, maka buat file baru
		var file, err = os.Create("test.txt") // membuat file 'test.txt'
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Printf("Success create file, %s", file.Name())
		return
	} else {
		fmt.Println("File already exists")
	}
}
```
Output:
```
File already exists
```

Kita juga dapat membuat direktori dengan menggunakan fungsi `os.Mkdir()` dengan memasukkan 2 parameter yaitu *path* dan *permission* dari direktori tersebut, kita dapat menggunakan `os.ModePerm` untuk memberikan *permission* penuh.

>Note: Konsep *permission* digunakan di sistem operasi Linux sedangkan di Windows tidak digunakan. Silahkan memakai *default value* yang ada di contoh.

Fungsi ini akan mengembalikan satu nilai, yaitu `error`. Jika tidak terjadi `error`, maka nilai `error` akan bernilai `nil`.

Contoh, kita ingin membuat sebuah direktori bernama test di direktori yang sama dengan *file* `main.go`.
```go
import "os"
func main() {
	err := os.Mkdir("test", os.ModePerm) // membuat direktori 'test'
	if err != nil {
		panic(err)
	}

	fmt.Println("Success create directory")
}
```
Output:
```
Success create directory
```
Jika berhasil, maka akan muncul satu direktori baru dengan nama `test` di direktori yang sama dengan *file* `main.go`.
```
golang-test
|_ test
_main.go
```


## Open file
Pembukaan *file* di Golang dapat menggunakan fungsi `os.Open()`. Fungsi ini perlu diisi *parameter path* atau posisi dari file yang ingin kita buka. Sama halnya dengan `os.Create()`, fungsi ini akan mengembalikan dua nilai, yaitu `*os.File` dan `error`.
`*os.File` adalah tipe data yang merepresentasikan `file`. `error` digunakan untuk mengembalikan *error*. Jika tidak terjadi *error*, maka nilai `error` akan bernilai `nil`.

Contoh, terdapat *file* `test.txt` yang berisi data sebagai berikut:
```
testing file
```
Kita langsung saja menggunakan `os.Open()` dengan mengisi parameter berupa *path* atau posisi file dari `test.txt`.

```go
import "os"

func main() {
	// buka file
	file, err := os.Open("test.txt") // membuka file 'test.txt'
	if err != nil {
		panic(err)
		return
	}
	
	defer file.Close() // close file setelah membaca file
	// kita bisa mengolah code setelah file dibuka
	// ...
}
```
File yang sudah dibuka statusnya adalah *open*, maka dari itu perlu untuk di-*close* menggunakan *method* `Close()` dari `*os.File` setelah *file* tidak digunakan lagi.

Kita juga dapat menggunakan fungsi `os.OpenFile()` untuk membuka *file*. Fungsi ini memiliki parameter tambahan yaitu `flag` yang digunakan untuk menentukan *mode* pembukaan *file*. `flag` merupakan tipe data `int` yang memiliki beberapa nilai konstanta, yaitu:
* `os.0_RDONLY` : membuka *file* untuk dibaca saja
* `os.0_WRONLY` : membuka *file* untuk ditulis saja
* `os.O_RDWR` : membuka *file* untuk dibaca dan ditulis
* `os.0_APPEND` : membuka *file* untuk ditulis dan menambahkan data di akhir *file*
* `os.0_CREATE` : membuat *file* jika belum ada
* `os.0_TRUNC` : membuat *file* baru jika sudah ada

Kemudian, terdapat parameter terakhir berupa `perm` yang digunakan untuk menentukan *permission* dari *file* yang dibuat. `perm` merupakan tipe data `int`, untuk saat ini kita cukup tulis dengan `0644` saja.
```go
import "os"

func main() {
	// buka file
	file, err := os.OpenFile("test.txt", os.O_RDWR, 0644) // membuka file 'test.txt' dengan mode read a
	if err != nil {
		panic(err)
		return
	}

	defer file.Close() // close file setelah membuka file
	// kita bisa mengolah code setelah file dibuka
	// ...
}
```


## Read file
Untuk membaca file kita perlu membuka file terlebih dahulu seperti yang dijelaskan di atas. Kemudian kita akan menggunakan *package* `ioutil` untuk membaca *file* yang kita buka. Fungsi yang dipakai adalah `ioutil.ReadAll()` dengan parameter yang diisi dari `*os.File`. Fungsi ini mengembalikan isi *file* bertipe `byte` dan `error`. Kita perlu melakukan pengecekan apakah *error* ada atau tidak.
```go
import (
	"fmt"
	"os"
)

func main() {
// buka file
file, err := os.Open("test.txt")
if err != nil {
panic(err)
return
}
defer file.Close()
// membaca isi file
var content, err = ioutil.ReadAll(file)
if err != nil {
	panic(err)
return
}
fmt.Println(string (content)) // menampilkan isi file
}
```
Output:
```
testing file
```

Terdapat cara lain yang bisa digunakan untuk membaca file yaitu menggunakan fungsi `os.ReadFile()`. Fungsi ini membuka sekaligus membaca *file*. Kita cukup mengisi parameter dengan path dari posisi *file*, yang akan mengembalikan isi *file* bertipe `byte` dan `error`. Jangan lupa untuk mengecek *error* nya.
```go
import "os"

func main() {
	// membaca isi file
	var content, err = os.ReadFile("test.txt")
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(string(content)) // menampilkan isi file.
}
```
Output:
```
testing file
```


## Write and update file
Untuk menulis atau mengubah isi *file* di Golang, dapat menggunakan beberapa fungsi yang disediakan oleh *package* `os`. Pertama, kita dapat menggunakan fungsi `os.WriteFile()` dengan mengisi parameter *path* dari posisi * file*, data yang ingin ditulis, dan permission dari file dengan mengisi `0644` Fungsi ini akan mengembalikan `error`. Jangan lupa untuk mengecek *error* nya. 

Contoh, anggap kita ingin menulis di file bernama `write-data.txt` dengan data sebagai berikut:
```
ini hasil write data dari Golang
```
```go
import (
	"fmt"
	"os"
)

func main() {
	data = "ini hasil write data dari Golang"
	
	err := os.WriteFile("write-data.txt", []byte(data), 0644)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("success write data")
}
```
Output:
```
success write data
```

Hasil dari *write* data ini dapat kita lihat di *file* `write-data.txt` yang sudah kita buat.
Terdapat cara kedua, yaitu menggunakan fungsi dari `*os.File` bernama `WriteString()` dengan mengisi parameter data yang ingin ditulis. Fungsi ini akan mengembalikan `int` dan `error`. Nilai kembalian `int`, adalah representasi dari berapa banyak *bytes data* yang berhasil ditulis. Jangan lupa untuk mengecek *error* nya.
```go
import (
	"fmt"
	"os"
)

func main() {
	// create file
	file, err := os.Create("write-data.txt")
	if err != nil {
		panic(err)
		return
	}

	defer file.Close() // jangan lupa di close setelah selesai menulis file
	// menulis data ke file
	n, err = file.WriteString("ini hasil write data dari Golang")
	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("success wrote %d bytes data\n", n)
}
```
Output:
```
success wrote 32 bytes data
```

Hasil dari *write data* di variabel `n` menampilkan angka `32` yang merupakan jumlah *bytes* dari *data* yang berhasil kita tulis di *file* `write-data.txt`. Biasanya kita tidak perlu menangkap kembalian nilai dari `n` karena jarang digunakan. Sehingga kita bisa mengganti baris yang berisi `file.WriteString()` seperti
ini:
```go
// potongan kode di atas
// ...
_, err = file.WriteString("ini hasil write data dari Golang")
// ...
```

## Delete file and directory
Di Golang kita dapat dengan mudah menghapus sebuah *file* menggunakan fungsi `os.Remove()`. Kita cukup mengisi *path* / posisi dari file yang ingin kita hapus. Fungsi ini akan mengembalikan `error`. 

Contoh, kita ingin menghapus *file* `write-data.txt` yang sudah kita buat sebelumnya.\
```go
import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("write-data.txt")
	if err != nil {
		panic(err)
	}

	
	fmt.Println("success delete file")
}
```
Output:
```
success delete file
```
Untuk menghapus *directory* kita dapat menggunakan fungsi `os.RemoveAll()` dengan mengisi *path*/ posisi dari *directory* yang ingin kita hapus. Fungsi ini akan mengembalikan `error`. Fungsi ini juga akan menghapus semua *file* yang ada di dalam *directory* tersebut.

Contoh, kita ingin menghapus *directory* bernama `test`.
```go
import (
	"fmt"
	"os"
)
func main() {
	err := os.RemoveAll("test")
	if err != nil {
		panic (err)
	}
	fmt.Println("success delete directory")
}
```
Output:
```
success delete directory
```
