# Basic *input* argument  

Package `fmt` juga mendukung untuk menerima *input* dari terminal atau console. Berikut beberapa fungsi yang dapat digunakan untuk menerima *input*.  

## fmt.Scan  
Fungsi `Scan` menerima *input* dengan cara menangkap data atau nilai *input* nya ke dalam sebuah variable yang diisi di parameter fungsi. Kita menggunakan pointer untuk menangkap *input* nya dengan sintaks & (materi pointer akan dibahas di materi selanjutnya).  
```go 
func main() {  
var name string  
fmt.Print("Enter your name: ")  
fmt.Scan (&name) // set *input* to variable 'name'  
Variable  
fmt.Println("Hello", name)  
}  
```

Contoh di atas adalah penggunaan fungsi `Scan` untuk menyimpan data atau nilai *input* ke dalam variabel `name`. Kemudian akan memunculkan ke terminal dengan format: `Hello <name>`.  

`Scan` juga dapat menerima lebih dari satu parameter, kita cukup meng*input*nya di terminal yang dipisah dengan spasi.  

```go
func main() {  
var name, address string  
fmt.Print("Enter your name and address : ")  
fmt.Scan (&name, &address) // set *input* to variable 'name' and 'address'  
fmt.Println("Hello", name)  
fmt.Println("Your address", address)  
} 
```

```output
Enter your name and address: Budi Jakarta 
Hello Budi  
Your address Jakarta  
```

## fmt.Scanln  
`Scanln` hampir sama dengan fungsi `Scan`, namun setelah memberikan *input* di terminal atau console, kemudian akan menambah *new line* atau enter di akhir *input* nya.  

```go
func main() {  
	var name, address string  
	fmt.Print("Enter your name: ")  
	fmt. Scanln (&name)  
	fmt.Print("Enter your address : ")  
	fmt.Scanln(&address)  
	fmt.Println("Hello", name)  
	fmt.Println("Your address", address) 
}  
```
```output
go run main.go  
Enter your name: Budi  
Enter your address : Jakarta  
Hello Budi  
Your address Jakarta  
```

## fmt.Scanf  
Konsep fungsi dari `Scanf` hampir sama dengan `Printf`, yaitu untuk *formating*. Namun `Scanf` digunakan untuk *formating* *input* yang diberikan menggunakan format *string* sesuai dokumentasi. Dan fungsi Scanf tidak memberikan new line atau enter di akhir *input* nya.  

```go
func main() {  
var name string  
fmt.Print("Enter your name: ")  
fmt. Scanf("%s", &name) // set *input* with format string to variable name  
fmt.Println("Hello", name)  
}  
```
```output
Enter your name: Budi  
Hello Budi  
```

`Scanf` juga bisa menerima lebih dari 1 parameter seperti fungsi `Scan`. Namun kita harus menentukan format *string* yang akan diterima.  
```go
func main() {  
var name, address string  
fmt.Print("Enter your name and address: ")  
fmt.Scanf("%s %s", &name, &address) // menerima 2 *input* string yang dipisahkan oleh spasi, yang per 

fmt.Println("Hello", name)  
fmt.Println("Your address", address)  
}
```
```output
Enter your name and address: Budi Jakarta  
Hello Budi  
Your address Jakarta  
```
>NOTE: Tidak perlu membandingkan mana fungsi *input* terbaik atau terburuk. Karena setiap fungsi *input* digunakan sesuai kebutuhan dan kondisinya.
