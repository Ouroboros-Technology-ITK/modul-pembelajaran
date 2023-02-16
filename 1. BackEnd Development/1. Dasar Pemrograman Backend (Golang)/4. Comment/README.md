# Comment
## What is comment?

comment atau komentar adalah sintaks yang digunakan untuk mengabaikan teks atau baris kode dieksekusi.

komentar bisa dimanfaatkan untuk menyisipkan catatan pada kode program, menulis penjelassan/deskripsi mengenai suatu blok kode, atau bisa juga diggunakan untuk me-remark kode(men-nonaktifkan kode).

untuk membuat  1 baris comment, gunakan `//`  pada awal baris sehingga text yang mengikutinya dalam satu baris menjadi pesan text biasa. jika ingin menulis dalam beberapa baris, cara efisien bisa menggunakan syntax `/*` untuk pembuka diawal komentar, dan `*/` diakhir komentar.

```go
func main() {	
	//komentar satu baris
	/*
		komentar multi baris
		fmt.Println("hello world");
	*/
	fmt.Println("hello Go!")
}
```
```shell
$ go run main.go
//output
Hello Go!
```
