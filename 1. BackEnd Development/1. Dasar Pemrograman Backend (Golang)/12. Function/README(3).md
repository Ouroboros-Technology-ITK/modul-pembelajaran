# Defer  
Terkadang program yang sedang dijalankan perlu untuk menjalankan resource tertentu, seperti input/output maupun *network resource* yang memerlukan *cleanup* setelah eksekusi berakhir.  

Untuk mengakhiri dan memastikan bahwa *resource* tersebut telah di-*cleanup*, kita dapat menggunakan `defer`.  

Contohnya untuk membuka file diperlukan *resource* dari **os**. `defer` digunakan untuk memastikan bahwa operasi `file.Close()` pasti dijalankan ketika serangkaian statement kode ini berakhir maupun jika terjadi *error*.  
```go
func readFile() {  
	file, err := os.Open("file.txt")  
	if err != nil {    
		log.Println("Cannot read file")  
	}
	defer file.Close()  
	//some other statements  
}
```

Cara kerja `defer` cukup unik. Meskipun di buat di awal, dia akan dieksekusi setelah semua kode dalam satu blok selesai berjalan. Contoh:  
```go
func main() {  
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
	fmt.Println("This will be printed second")
}  
```
```Output
This will be printed first  
This will be printed second  
This will be printed last 
```

Perlu diperhatikan bahwa `defer` dipanggil tepat sebelum `return` dijalankan, jadi jangan gunakan `defer` untuk mengakhiri program. Contohnya:  
```go
func main() {  
	fmt.Println(countByOne(2))  
}  
func countByOne(x int) int {  
	defer fmt.Println("Counting done")  
	
	fmt.Println("Counting start")  
	
	return x + 1  
}  
```
```Output
Counting start  
Counting done  
3  
```

Contoh di atas memunculkan `"Counting done"` terlebih dahulu karena menggunakan defer, setelah  itu `return` dijalankan dengan mengembalikan value `x + 1`.
