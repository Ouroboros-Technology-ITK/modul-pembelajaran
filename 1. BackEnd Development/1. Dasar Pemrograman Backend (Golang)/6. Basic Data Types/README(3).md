# Boolean

## What is boolean?  
Tipe data *boolean* adalah tipe data yang hanya berisi 2 nilai yaitu **TRUE** (benar) atau **FALSE** (salah). 

Tipe data *boolean* banyak digunakan untuk percabangan kode program (*condition*) untuk memutuskan apa yang mesti dijalankan ketika sebuah kondisi terjadi.  

## boolean di go  

Di Golang tipe data *boolean* direpresentasikan dengan kata kunci `bool`. *Boolean* di Golang hanya berisi `true` atau `false` (huruf kecil semua).  
```Go
func main() {  
fmt.Println("Benar =", true) fmt.Println("Salah =", false)  
}  
```
```Output
Benar = true  
Salah = false  
```
>Note: walaupun sangat sederhana, pastikan kita memahami konsepnya karena tipe data ini akan sering dipakai ketika nanti membuat sebuah program.
