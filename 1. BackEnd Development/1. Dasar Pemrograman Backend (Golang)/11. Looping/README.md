# Looping
Pernahkah kita membayangkan jika harus menulis sebuah statement "Hello World" sebanyak seribu kali secara manual? Tentunya itu akan sangat melelahkan. Untungnya, Golang dapat mempermudah proses ini dengan menggunakan looping atau perulangan. Di halaman ini, kedua istilah tersebut menunjukkan hal yang sama.  

Loop berguna jika ingin menjalankan kode yang sama berulang-ulang. Setiap eksekusi loop disebut iterasi. 

![index](https://java2blog.com/wp-content/webpc-passthru.php?src=https://java2blog.com/wp-content/uploads/2017/04/whileLoopFlowDiagram.png&nocache=1)
(Sumber: [java2blog.com](https://java2blog.com/wp-content/webpc-passthru.php?src=https://java2blog.com/wp-content/uploads/2017/04/whileLoopFlowDiagram.png&nocache=1))

## `for`  
Golang menggunakan sintaks `for` untuk membuat sebuah `looping`. Terdapat 3 *statement* yang dapat diisi agar *looping* dapat berjalan dengan baik yang dipisah dengan `;` (titik koma).  
```go
for init_statement; condition_statement; post_statement {  
	// kode yang akan dieksekusi setiap perulangan.  
}  
```

## Init statement  
*Init statement* adalah nilai awal yang akan dijadikan sebagai penanda mulai dalam perulangan.  

Di Golang, *init statement* diinisiasi menggunakan variabel yang diisi dengan tipe data *number*.  

## Condition statement  
*Condition statement* adalah kondisi yang digunakan untuk mengecek / evaluasi setiap iterasi *loop*. Jika kondisi bernilai TRUE, maka looping akan terus berulang. Jika bernilai FALSE, *looping* berakhir. 

Dalam *condition* ini kita bisa menggunakan operasi perbandingan.

## Post statement  
Ini adalah *statement* yang digunakan untuk mengubah nilai *init statement* setiap terjadi perulangan kembali.  

Di Golang umumnya kita mengenal istilah "increment" untuk melakukan peningkatan nilai dan "decrement" untuk melakukan penurunan nilai.

## Basic for statement  
Kita ambil contoh membuat perulangan sederhana di Golang:  

```go
func main() {  
	for i = 0; i < 5 ; i++ {
		fmt.Println("Hello Go!", i + 1)  
	}  
}  
```
```Output
Hello Go! 1  
Hello Go! 2  
Hello Go! 3  
Hello Go! 4  
Hello Go! 5  
```

Contoh di atas adalah *looping* sebanyak 5 kali dengan memunculkan "Hello Go!" dan angka dari iterasi / perulangan ditambah `1`. 

*Init statement*-nya adalah `i := 0`, di mana perulangan ini di awali dengan membuat *short syntax* variable `i` bernilai `0`. 

*Condition statement* adalah `i < 5`, yang berarti *looping* akan terus berulang selama *variable* `i` kurang dari `5`.

*Post statement* akan mengubah nilai di *variable* `i` meningkat 1 nilai tiap perulangan.
![index](https://www.practical-go-lessons.com/img/syntax_for_loop_with_counter.2df74605.png)
(Sumber: [practical-go-lesson.com](https://www.practical-go-lessons.com/img/syntax_for_loop_with_counter.2df74605.png))

Contoh lainnya:  
```go
for i := 50; i > 0; i -= 10 {  
	fmt.Println(i)  
}  
```
```Output
50  
40  
30  
20  
10  
```

Init statement juga dapat diisi dengan variable yang sudah didefinisikan sebelumnya.  

```go
func main() {  
	var i int  
	for i = 0; i < 5; i++ {  
		fmt.Println("Hello Go!", i + 1)  
	}  
}  
```

Karena init statement berupa short syntax variable, maka kita juga dapat membuat variable sebanyak yang dibutuhkan.  

```go
func main() {  
	for i, total = 1, 0; i <= 5 ; i++ {  
		total += i // 1+ 2+ 3+ 4+ 5 = 15  
	}  
}  
```

Contoh di atas adalah membuat variable `i` dan `total` di *init statement* `for` . *variable* `i` dimulai  
dengan nilai `1` dan *variable* `total` akan dimulai dengan nilai `0` . Setiap perulangan akan menjumlahkan nilai `i` ke *variable* `total`.  

# break  
Terdapat *syntax* `break` yang bisa kita digunakan untuk memutus atau mengakhiri eksekusi perulangan.
Penggunaan `break` umumnya dipanggil di dalam kondisi.  
```go  
func main() {  
	for i = 0; i < 5; i++ {  
		if i == 3 {  
			break  
		}
		fmt.Println("Hello Go!", i+1)  
	}
}
```
```Output
Hello Go! 1  
Hello Go! 2  
Hello Go! 3  
```

Contoh di atas harusnya menghasilkan *output* iterasi sebanyak 5 kali. Namun, karena ada kondisi `break` ketika nilai `i` sama dengan `3`. Maka iterasi terhenti ketika nilai `i` sudah berisi nilai `3`, sehingga *output* yang dihasilkan hanya sebanyak 3 iterasi.  

## continue  
`continue` adalah *syntax* yang digunakan untuk melewati (*skip*) 1 atau lebih perulangan, kemudian diikuti dengan perulangan berikutnya.
```go
func main() {  
	for i = 1; i <= 5; i++ {  
		if i % 2 == 0 {  
			continue  
		}  
		
		fmt.Println("Hello Go!", i)  	
	}
}  
```
```go
Hello Go! 1  
Hello Go! 3  
Hello Go! 5  
```

Contoh di atas adalah penggunaan `continue` ketika nilai `i` adalah genap atau nilainya habis dibagi 2  
`(i % 2 == 0`), maka akan melewati perulangan tersebut dan lanjut di perulangan berikutnya. Sehingga output nya akan memunculkan iterasi dengan nilai `i` ganjil.  

## Condition-only in `for`  
`for` boleh hanya berisi *condition statement* saja, jika ingin membuat *init statement* di luar code perulangan dan *post statement* di dalam code perulangan.  
```go
func main() {  
	i := 0 // init statement  
		for i < 5 {  
			fmt.Println("Hello Go!", i+ 1)  
			i++ // post statement
	}  
}  
```
```Output
Hello Go! 1  
Hello Go! 2 Hello Go! 3  
Hello Go! 4  
Hello Go! 5  
```

## empty for statement  
Kita juga boleh tidak mengisi semua *statement* yang ada, sehingga hanya berisi kode `for` dan langsung membuat *block statement* `{}`. Namun, kita tetap harus membuat 3 *statement* looping (**init, condition, post**) di luar atau di dalam kode perulangan.  
```go
func main() {  
	var i = 0 // init statement  
	
	for {  
		if i >= 5 { // condition statement  
			break // menghentikan looping jika kondisi 'true'  
		}  
		fmt.Println("Hello Go!", i + 1) // code execute  
		
		i++ // post statement  
	}
}
```

## Infinite loop  
Perulangan ini akan berjalan terus-menerus (tidak akan pernah berhenti) hingga ada suatu interupsi dari sistem.  

Hal ini bisa terjadi karena kita lupa atau salah dalam membuat satu atau lebih dari 3 *statement looping*.

Misalnya adalah tidak membuat *condition statement* sama sekali, atau *membuat condition statement* selalu bernilai `true`.  

Contoh:  
```go
func main() {  
	for {  
		fmt.Println("Hello Go!")  
		
		time.Sleep(1 * time.Second) // ini adalah fungsi yang dibuat agar program delay selama 1 detik  
```
```Output
Hello Go! 
Hello Go! 
Hello Go!  
...  
```
Contoh di atas adalah bentuk `infinite loop` karena kita tidak mendefinisikan `condition statement` sama sekali. Sehingga kondisi di `for` akan selalu bernilai `true` dan melakukan perulangan terus menerus.  

## Nested looping  
*Nested looping* adalah perulangan yang **bersarang**. Maksudnya adalah dalam *looping* tersebut masih terdapat *looping* lagi bahkan dalam *looping* tersebut, masih memungkinkan untuk membuat perulangan lagi sehingga looping *tersebut* bersarang ke dalam. 

*Looping* akan diselesaikan dari tingkat terdalam. Ketika *loop* pada tingkat paling dalam telah selesai dilakukan, maka *looping* pada tingkat di atasnya akan dilakukan dan begitu seterusnya. Perulangan paling luar biasanya kita sebut dengan `outer loop`, sedangkan yang di dalam disebut `inner loop  `

```go
func main() {  
	for i = 0; i < 2; i++ {  
		for j = 0; j < 3; j++ {  
			fmt.Printf("Hello Go! i: %d, j: %d\n", i, j)  
		}
	}
}
```
```Output
Hello Go! i: 1 j: 1  
Hello Go! i: 1 j: 2  
Hello Go! i: 1 j: 3  
Hello Go! i: 2 j: 1  
Hello Go! i: 2 j: 2  
Hello Go! i: 2 j: 3  
```

Contoh di atas adalah bentuk *nested looping* dengan cara kerja yaitu `outer loop` tidak akan dilanjutkan dulu sebelum `inner loop` menyelesaikan perulangannya. Karena perulangan paling luar melakukan iterasi sebanyak 2 kali dan perulangan di dalam melakukan interasi 3 kali, maka hasilnya akan memunculkan *output* sebanyak 6 kali iterasi.
