# Basic Golang Command

## go build

`go build` - digunakan untuk mengkompilasi (compile) file program dengan output berupa binary file yang dapat dijalankan sesuai dengan OS. Jika menggunakan Windows maka outputnya memiliki extensi `.exe` & tidak berekstensi jika menggunakan Linux/Mac, bisa juga ditambah opsi `-o` diikuti nama baru sehingga nama file terkompilasi akan berubah

Jika kita ingin menjalankan file bernama `greet.go`, perintahnya 
`go build greet.go

Meskipun tidak ada yang terlihat jelas setelah kita menjalankan perintah, Go telah membuat file eksekusi program kita. Jika kita mengetikkan perintah ls, kita akan melihat program Go asli dan file eksekusinya.

pada terminal kita, kita ketik go build diikuti dengan nama file kita dan tekan Enter. 

`ls`
`greet     greet.go`

Untuk mengeksekusi file tersebut, kita panggil:

`./greet`

file binary dapat dijalankan dengan langsung menyebutkan nama file beserta lokasinya, dimana `./` artinya current directory 
contoh
```shell
$ go build main.go
$ ls
main // jika windows maka main.exe

$ ./hello world
Hello Go!

```
> file bisa dijalankan secara native pada OS tanpa menggunakan Go Command Line Tools ataupun IDE.  sehingga file dapat didistrusi ke sistem operasi lain 


## go run

apa yang terjadi jika kita ingin mengubah program kita? Kita harus mengkompilasi file lain yang dapat dieksekusi dan kemudian menjalankan file tersebut. Bayangkan jika kita harus melakukan hal tersebut setiap saat!

Untungnya, kita memiliki perintah go run yang diikuti dengan nama file Go. Kita dapat menggunakan perintah go run seperti ini:
go run greet.go

Perintah go run mengkompilasi kode (seperti go build) dan mengeksekusinya. Tidak seperti go build, go run TIDAK akan membuat file yang dapat dieksekusi di folder kita saat ini.

`go run` digunakan untuk mengeksekusi file berekstensi `<filename>.go` di terminal. misal menggunakan file `main.go` berisikan kode berikut:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")
}
```

lalu ketikkan perintah `go run` pada terminal diikuti dengan nama file tersebut.

```shell
go run main.go


// output
Hello Go!
```




