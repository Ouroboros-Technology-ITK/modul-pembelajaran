# Get 3rd party package  

## Package management  
Di Golang terdapat beberapa *module* / *library project* lain yang dapat membantu kita dalam pengembangan sistem. Kita bisa mendapatkannya dari pihak ketiga yang ada di internet secara gratis.

### `go get`  
Kita dapat menggunakan *command* `go get` dari Golang untuk mendapatkan *module / library* yang ada. Cukup menulis `go get` diikuti dengan nama *module* yang diinginkan di *root directory project* kita.
```
go get <module-name>  
```

Sebagai contoh, jika kita ingin menggunakan *module* [currency](https://pkg.go.dev/golang.org/x/text/currency) untuk kebutuhan konversi dan *formatting* keuangan. Langkah-langkah untuk menggunakannya sebagai berikut:

Pertama, *install module/package* dengan `go get github.com/bojanz/currency` 
```
> go get github.com/bojanz/currency  
```

Kita bisa memastikan bahwa *module/package* sudah terinstall dengan cara melihat *file* `go.mod` yang ada di *root directory project* kita.
```go
// file go.mod  
module test-golang  

go 1.18  

require (  
	github.com/bojanz/currency v1.0.4 // indirect 
	github.com/cockroachdb/apd/v3 v3.1.0 // indirect 
	github.com/pkg/errors v0.9.1 // indirect  
)
```
Jika berhasil, maka terdapat tambahan nama *module* di dalam `required` pada *file* `go.mod` yang sudah kita *install*.

Kedua, kita `import`* module* di dalam *project* kita.

Bonus: Jika menggunakan *editor* VSCode, terdapat fitur *autocomplete*, saat kita memanggil `field` dari *module/library* pihak ketiga, maka akan otomatis di-*import* oleh VSCode.
```go
import (  
	"github.com/bojanz/currency" //baris ini otomatis akan ditambahkan oleh editor VSCode  
)
func main() {  
	var dollar, err = currency.NewAmount ("100", "USD") // menggunakan fungsi 'NewAmount' dari package 'currency' 
	if err != nil {  
		panic(err)  
	}  
	
	formater := currency. NewFormatter (currency.NewLocale("en_US")) // menggunakan fungsi 'NewFormatter' 
	fmt.Println(formater.Format(dollar))  
	dollarToIdr, err := dollar.Convert("IDR", "14000") // menggunakan fungsi 'Convert' dari package 'cu 
	if err != nil {  
		panic(err)
	}  
	fmt.Println(dollarToIdr)  
}  
```
Output:
```
$100.00  
1400000 IDR  
```

### `go mod tidy`
Seringkali kita kadang salah memasukkan *library* yang kita inginkan ke dalam `go.mod` atau sudah tidak ingin menggunakan *library* tersebut. Kita dapat menggunakan command `go mod tidy` di *terminal* atau *console*. Perintah ini akan menghapus semua *library* yang tidak digunakan dari `go.mod`.

Pastikan direktori yang ada di *terminal* atau *console* merupakan *root directory* dari *project* kita yang sejajar dengan *file* `go.mod`.  
```
> go mod tidy  
```
Setelah menjalankan perintah ini, maka *file* `go.mod` akan berubah dengan 
menghilangkan semua library yang tidak digunakan di `require`.

## Golang awesome  
Kadang kita kebingungan mencari sekumpulan *library* yang dapat digunakan untuk kebutuhan program kita. Biasanya komunitas dari tiap bahasa pemrograman memiliki list informasi *library* yang sudah dibuat dan dapat kita gunakan secara gratis.

Di Golang juga terdapat *list library* yang sudah dikumpulkan oleh komunitas menjadi satu. Berikut linknya : [Golang Awesome](https://go.libhunt.com/).  

Kita dapat menggunakan secara gratis, semua *library* sesuai kebutuhan kita.  
