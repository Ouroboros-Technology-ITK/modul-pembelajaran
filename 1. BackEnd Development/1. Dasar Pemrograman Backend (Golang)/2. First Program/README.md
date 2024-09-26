## hello world
---
untuk memulai pemrograman Go, pertama kita perlu membuat file dengan esktensi `.go` seperti `main.go` . untuk pembelajaran ini buatlah folder kosong bernama `basic Golang`, gunakan untuk menyimpan file `.go` yang akan dibuat pada pembelajaran ini.

Program pertama kita akan mencetak pesan klasik "hello world" dengan file bernama `hello-world.go`. Berikut adalah kode sumber lengkapnya.

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```


## Compiling File
---
mari kita pelajari cara menggunakan Kompiler Go untuk mengkompilasi sebuah file menjadi sebuah file yang dapat dieksekusi.

Pada terminal kita, ketik `go build` diikuti dengan nama file kita dan tekan Enter. Jika kita ingin menjalankan file bernama `hello-world.go`, perintahnya akan terlihat seperti:

`$ go build hello-world.go`

Meskipun tidak ada yang terlihat pada terminal setelah perintah dijalankan, Go telah membuat file eksekusi program kita. Jika kita mengetikkan perintah `ls`, kita akan melihat program Go asli dan file eksekusinya.
`$ ls`
`// output` 
`hello-world hello-world.go`
Untuk mengeksekusi file tersebut, kita panggil:

./hello-world
Untuk menjalankan program, letakkan kode di hello-world.go dan gunakan go run:
- `$ go run hello-world.go`
`//output hello world


## Package
---
Setiap program Go terbuat dari paket-paket ("package"). setidaknya terdapat satu file dengan nama package `main` (Utama). file yang ber-package `main` akan dieksekusi pertama kali ketika program dijalankan

cara  menentukan package dengan menggunakan keyword `package`

```go
// package <nama-package>
// contoh:
package main
```

## Import
---
setelah package ditentukan, untuk mengimpor isi dari package atau menggunakan package lain & memanfaatkannya maka digunakan keyword `import`.

sebagai contoh kode dibawah ini menggunakan paket `main` untuk meng-import `"fmt"` (format) dan `"math/rand"` (random).

```go
// deklarasi main package
package main

// impor isi package
import (
	"fmt"
	"math/rand"
)

// Main function
func main() {
	fmt.Println("Bilangan kesukaan saya adalah", rand.Intn(10))
}
// output 
//Bilangan kesukaan saya adalah <0-9>
```

>jika dilihat kembali pada fungsi `import`, paket `"math/rand"` memiliki tanda "/" ditengahnya, artinya program mengimpor paket `rand` yang berada didalam paket `math` . 


## main() & init() function
---
Bahasa Go memiliki dua fungsi untuk tujuan khusus: main() dan init(). Fungsi main() adalah titik masuk program yang dapat dieksekusi dan tidak menerima argumen atau mengembalikan apapun. Fungsi ini dideklarasikan secara implisit dan dipanggil ketika paket diinisialisasi. Fungsi init() ada di setiap paket dan dipanggil sebelum pemanggilan fungsi main(). Fungsi ini digunakan untuk menginisialisasi variabel global yang tidak dapat diinisialisasi dalam konteks global.

```go
// deklarasi main package
package main
  
// impor isi package
import "fmt"
  
// beberapa fungsi init() 
func init() {
    fmt.Println("Welcome to init() function")
}
  
func init() {
    fmt.Println("Hello! init() function")
}
  
// Main function
func main() {
    fmt.Println("Welcome to main() function")
}

/* Output
Welcome to init() function
Hello! init() function
Welcome to main() function
*/
```
