#  Pointer  
https://raw.githubusercontent.com/Ouroboros-Tech/modul-pembelajaran/blob/Dasar-golang/image/nilai%20memori.jpg
https://raw.githubusercontent.com/Ouroboros-Tech/modul-pembelajaran/blob/Dasar-golang/image/blok%20memori.jpg
https://raw.githubusercontent.com/Ouroboros-Tech/modul-pembelajaran/Dasar-golang/image/alamat%20memori.jpg

Nilai atau variabel yang dideklarasikan dalam program akan disimpan di dalam komputer. Setiap nilai yang disimpan di dalam komputer memiliki alamat memori. **Pointer** adalah variabel yang menyimpan alamat memori dari suatu nilai.

Sebagai contoh, ketika kita mendeklarasikan variabel number dengan nilai 10, nilai tersebut akan disimpan dalam komputer di alamat memori tertentu.  
```go
var number = 10  
```
![index](https://raw.githubusercontent.com/Ouroboros-Tech/modul-pembelajaran/blob/Dasar-golang/image/nilai%20memori.jpg)
Nilai `10` dari deklarasi variabel number disimpan pada alamat memori `0xc000018030`.  

Lalu apakah maksud dari alamat memori? Bayangkan sebuah kompleks perumahan, di setiap rumah pasti memiliki nomor rumah dan alamat.  

![index](https://raw.githubusercontent.com/Ouroboros-Tech/modul-pembelajaran/blob/Dasar-golang/image/blok%20memori.jpg)
Seperti kompleks perumahan, memori juga memiliki alamat untuk setiap blok memori. Pada ilustrasi di bawah, terdapat 4 blok memori dengan alamat memori `0xc000018030`, `0xc000018038`, `0xc000018040`, dan `0xc000018048`. Di setiap alamat memori tersebut terdapat nilai yang disimpan, sebagai contoh pada alamat memori `0xc000018030` menyimpan nilai `100`.  
![index](https://raw.githubusercontent.com/Ouroboros-Tech/modul-pembelajaran/blob/Dasar-golang/image/nilai%20memori.jpg)
Dengan adanya pointer, kita dapat membuat kode program yang kompleks tetapi tetap menghemat memori. Dan dapat langsung melakukan modifikasi nilai yang disimpan di memori. 

Pointer di Golang memiliki dua operator, address operator `&` (symbol and) dan indirect operator `*` (asterisk).  

## Syntax `&`  

Operator ini digunakan untuk mendapatkan alamat memori dari suatu nilai. Caranya cukup memberikan tanda `&` sebelum nama variabel nya.  
```go
func main() {  
	var number = 10  
	fmt.Println(&number)  
}  
```
```Output  
0xc000018030  
```

Ketika kode di atas dieksekusi maka output yang dihasilkan adalah `0xc000018030` yang merupakan alamat memori di mana nilai dari variabel `number` disimpan.  

Syntax `*`  
Operator ini digunakan untuk mendapatkan nilai dari variabel yang ditunjuk / direferensikan. Istilah lain dari indirect operator adalah **dereferencing**. Caranya cukup memberikan tanda `*` sebelum nama variabel nya.  
```go
func main() {
	var number = 10  
	var pointerNumber = &number // pointer Number berisi alamat memori dari variable 'number'  
	fmt.Println(pointerNumber)  
	fmt.Println(*pointerNumber) // mendapatkan nilai dari variable 'number'
}    
```
```Output
0xc000018030  
10  
```

Operator `*` bisa menjadi tipe data yang dapat diisi dengan pointer dari tipe data sejenis.  
```go
func main() {  
	var (  
		number int  
		pointerNumber *int // hanya dapat diisi dengan pointer dari tipe data 'int'   
		
	number = 10  
	pointerNumber = &number // 'pointerNumber' di assign dengan alamat memori dari 'number' bertipe dat  
}  
```

Contoh penggunaan:  
```go
func AddByOne (number *int) { // hanya menerima parameter pointer dari tipe data 'int'  
	*number *number + 1  
}  
func main() {  
	var number = 10  
	AddByOne(&number)  
	fmt.Println(number)  
}  
	
```
```Output
11  
```

## Pass by value  
Secara *default* di Golang, kita melakukan *passing by value* ketika mengirimkan nilai dari variabel ke fungsi atau dari variabel ke variabel lain. Artinya, jika kita mengirim sebuah nilai ke variabel atau parameter fungsi, maka sebenarnya yang di kirim adalah **duplikasi nilai** nya.  
```go
func AddByOne (number int) {  
	number = number + 1  
	fmt.Println("in function AddByOne:", number)  
}  
	
func main() {  
	var number int = 10  
	AddByOne(number)  
	fmt.Println("in function main:", number)  
}  
```
```Output
in function AddByOne: 11  
in function main: 10  
```
Contoh di atas adalah contoh *passing by value*, di mana nilai yang di kirim dari *variable* `number` ke fungsi `AddByOne` adalah **duplikasi nilai** `10`. Hasilnya, nilai yang ada di *variable* `number` tidak berubah walaupun terdapat operasi peningkatan 1 angka di fungsi `AddByOne`.  

## Pass by reference  
Berbeda dengan *passing by value*, *passing by reference* mengirimkan nilai dari variabel ke fungsi atau dari variabel ke variabel lain. Artinya, jika kita mengirim sebuah nilai ke variabel atau parameter fungsi, maka sebenarnya yang di kirim adalah **reference** atau **alamat memori** ke nilai tersebut.  
```go
func AddByOne (number *int) {  
	*number = *number + 1
	
	fmt.Println("in function AddByOne:", *number)  
}  

func main() {  
		var number int 10  
	AddByOne(&number)  
	fmt.Println("in function main:", number)  
}  
```
```Output
in function AddByOne: 11  
in function main: 11 
```
Jika kita menggunakan *passing by reference*, maka nilai yang di kirim dari variabel `number` ke fungsi `AddByOne` adalah **reference** ke nilai `10` yang ada di variabel `number`. Hasilnya, nilai yang ada di variabel `number` akan berubah juga ketika terdapat operasi peningkatan 1 angka di fungsi `AddByOne`. 

Pass by reference vs pass by value (Sumber:  [twitter/suksr](https://twitter.com/suksr/status/738130336270422017))

## Point zero value  
Pointer memiliki nilai *zero value* yaitu `nil` atau kosong (materi ini akan dibahas di `empty interface`). Hal ini perlu kita perhatikan karena ketika kita membuat `nil` pointer dan mengakses atribut dari pointer tersebut akan menyebabkan program *panic* (materi ini akan dibahas di `error handling` ). Pesan *error panic* yang dikeluarkan adalah `panic: runtime error: invalid memory address or nil pointer dereference`.  
```go
func AddByOne (number *int) {  
	*number = *number + 1  
	// panic error  
	// karena referensi parameter 'number' tidak berisi alamat memori apapun (nil)  
}  

func main() {  
	var number *int // variable pointer yang tidak diisi referensi memori, akan menghasilkan zero value  
	
	AddByOne(number)  
}  
```
```Output
panic: runtime error: invalid memory address or nil pointer dereference  
[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x1007387f4]  

goroutine 1 [running]:  
main. AddByOne(...) main.go:15 +0x1f
```
