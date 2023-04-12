# Package time

## `time`
Golang menyediakan package `time` yang berisikan banyak sekali komponen yang bisa digunakan untuk keperluan penanggalan atau *date-time*.

### `time.Time`
Tipe time.Time merupakan representasi untuk objek date-time. Terdapat 2 cara yang bisa digunakan untuk membuatnya:
* Menggunakan fungsi `time.Now()` yang akan mengembalikan objek `date-time` yang berisi waktu saat ini.
* Membuatnya dengan fungsi `time.Date()` yang akan mengembalikan objek *date-time* yang berisi tanggal yang ditentukan dengan format:
```Output
time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)
```
```go
import "time"
func main() {
	var timel time.Time = time.Now()
	var time2 time.Time = time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC)
}
fmt.Println(time1)
fmt.Println(time2)
```
Output:
```Output
2022-07-24 14:37:20.733952 +0700 WIB m=+0.000058668
2020-06-01 00:00:00 +0000 UTC
```
Informasi *date-time* menggunakan `time.Now()` adalah relatif berdasarkan lokasi di mana aplikasi kita dijalankan. Karena kode di atas berjalan di lokasi Indonesia bagian barat, maka akan terdeteksi masuk dalam zona **GMT+7** atau **WIB**.

Sedangkan informasi dari fungsi `time.Date()` di atas sudah diatur saat deklarasi dengan zona **UTC**, maka informasinya menggunakan zona **UTC** yaitu **+0000 UTC**.

## Parsing date in Golang
### Parsing `string` to `time.Time`
Tipe data `string` dapat dikonversi menjadi `time.Time` dengan menggunakan fungsi `time.Parse()`.

Hal pertama yang perlu diperhatikan adalah menentukan *layout format* yang akan digunakan untuk *parsing*. Golang memiliki *layout format* yang cukup unik. Contohnya untuk *parsing* tahun bukan menggunakan `YYYY` tetapi menggunakan patokan tahun `2006`. Untuk parsing bulan tidak menggunakan `MM` tapi menggunakan `01`, dst.

Detailnya dapat dilihat di link berikut: [parsing time]()
```go
import "time"
func main() {
	layoutFormat := "2006-01-02 15:04:05"
	date, _ := time.Parse(layoutFormat, "2022-06-01 15:04:05")
	fmt.Println(date)
}
```
Output:
```
2022-06-01 15:04:05 +0000 UTC
```
Jika kita kebingungan dengan format *date-time* yang diberikan, maka kita bisa menggunakan *built-in layout time* yang sudah disediakan. kita menyebutnya dengan *predefined layout format date*
```go
import "time"

func main() {
	var date, = time.Parse(time. RFC822, "02 Sep 15 08:00 WIB")
	// time.RFC822 = "Mon, 02 Sep 15 08:00 WIB" -> format date-time bawaan dari golang

	fmt.Println(date)
}
```
Output:
```Output
2015-09-02 08:00:00 +0700 WIB
```

### Parsing `time.Time` to `string`
Kita dapat menggunakan fungsi `time.Format()` untuk mengkonversi `time.Time` ke *string*. Kita cukup mengisi parameter fungsi tersebut dengan format yang sudah di jelaskan di atas.
```go
func main() {
	var date = time.Now() // tipe data time.Time"
	
	var dateString = date.Format("Monday 02, January 2006 15:04 MST") // berubah menjadi tipe data 'str
	
	fmt.Println(dateString)
}
```
Output:
```Output
Monday 25, July 2022 07:50 WIB
```
Atau menggunakan format yang sudah disediakan (*predefined layout format date*).
```go
func main() {
	var date = time.Now()

	var dateString = date.Format(time.RFC3339)
	
	fmt.Println(dateString)
}
```
Output:
```Output
2022-07-25T07:53:14+07:00
```

## Handle Error Parsing `time.Time`
Ketika parsing `string` ke `time.Time`, sangat memungkinkan bisa terjadi *error* karena struktur data yang akan di-*parse* tidak sesuai dengan *layout* format yang digunakan. Kita bisa menangkap *error* dengan mengecek kembalian nilai ke 2 dari fungsi `time.Parse()`.
```go
func main() {
	var date, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
	
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	
	fmt.Println(date)
}
```
Output:
```Output
error parsing time "02 Sep 15 08:00 WIB": extra text: " 08:00 WIB"
```
Contoh di atas menghasilkan *error* karena format tidak sesuai dengan *layout* data yang akan *diparsing*.

## Timer and duration
Ada beberapa fungsi dalam package `time` yang bisa dimanfaatkan untuk **menunda** atau **mengatur** jadwal eksekusi sebuah proses dalam jeda waktu tertentu.

### `time.Sleep()`

Fungsi ini digunakan untuk menghentikan program sejenak. `time.Sleep()` bersifat **blocking**, statement di bawahnya tidak akan dieksekusi sampai pemberhentian selesai.

Contoh:
```go
import "time"
func main () {
	fmt.Println("start")
	time.Sleep(time.Second * 2)
	fmt.Println("after 2 seconds")
}
```
Output:
```Output
> go run main.go
start
after 2 seconds

```
Contoh di atas menghasilkan *output* `start`, kemudian setelah 2 detik muncul *output* `after 2 seconds`.

Kita dapat mengatur jeda waktu dengan mengisi durasi di parameter `time.Sleep()`
Golang menyediakan durasi yang bisa kita pilih mulai dari *millisecond* hingga *hour*.
```go
func main() {
	time.Sleep(time. Second) // jeda 1 detik
	time.Sleep(time. Minute) // jeda 1 menit

// jika ingin membuat jeda lebih dari 1 detik / menit, kita bisa meng-kali dengan angka yang diingi

	time.Sleep(time.Second * 2) // jeda 2 detik 
	time.Sleep(time.Minute * 4) // jeda 4 menit
}
```

## Scheduler
Selain untuk *blocking* proses, fungsi `time.Sleep()` bisa dimanfaatkan untuk membuat *scheduler* sederhana.
```go
func main() {
	for {
		fmt.Println("Hello Go!")
		time.Sleep(time.Second * 2)
	}
}
```
Output:
```Output
Hello Go!
Hello Go!
Hello Go!
...
```
Contohnya di atas merupakan *scheduler* untuk menampilkan *output* `Hello Go!` setiap 2 detik.

##  `time.Duration`
Di Golang juga terdapat `time.Duration` yang merepresentasikan durasi, contohnya seperti 1 menit, 2 jam 5 detik, dst.

Sebelumnya, sudah dibahas sedikit ketika membahas parameter di fungsi `time.Sleep()`.

Data dengan tipe ini bisa dihasilkan dari operasi pencarian *delta* atau selisih dari dua buah objek `time.Time`, atau bisa juga kita buat sendiri.

Tipe data durasi  `time.Duration`, yang sebenarnya merupakan tipe buatan baru dari `int64`.

Ada beberapa *predefined* konstanta durasi yang perlu kita ketahui:
* `time.Nanosecond` yang nilainya adalah 1
* `time.Microsecond` yang nilainya adalah 1000, atau 1000 x time. Nanosecond
* `time.Millisecond `yang nilainya adalah 1000000, atau 1000 x time.Microsecond
* `time.Second` yang nilainya adalah 1000000000, atau 1000 x time. Millisecond
* `time.Minute` yang nilainya adalah 60 x time. Second
* `time.Hour` yang nilainya adalah 60 x time. Minute
Dari *list* di atas bisa dicontohkan bahwa sebuah data dengan tipe `time.Duration` yang nilainya `1`, maka artinya durasi adalah 1 *nanosecond*.

### Calculate duration
Kita dapat menghitung durasi menggunakan fungsi `time.Since()`. Terdapat 1 parameter yang bisa diisi dengan `time.Time` yang akan menghitung durasi dari waktu tersebut.
```go
func main() {
	var start = time.Now()
	time.Sleep(time.Second * 2)
	
	var duration = time.Since(start)
	
	fmt.Println(duration)
}
```
Output:
```
2.005396666s
```
`time.Since` akan menghitung durasi antara waktu yang diberikan di *variable* `start` dengan jarak waktu ketika fungsi `time.Since` dipanggil.

Pada contoh di atas, karena ada *statement* `time.Sleep(time.Second * 2)` maka durasi waktu dari `time.Now()` sampai ke `time.Since()` akan berisi 2 detik (mungkin lebih sedikit, sekian mili/micro/nano-second, karena *eksekusi statement* juga butuh waktu).

### Calculate duration between two time
Hampir sama seperti `time.Since()`, namun di sini kita menggunakan fungsi `time.Sub()` yang akan menghitung durasi antara dua tipe `time.Time`.
```go
func main() {
	var start = time.Now()
	time.Sleep(time.Second 2)
	
	var end = time.Now()
	
	var duration = end.Sub(start) // menghitung durasi antara waktu yang diberikan di variable 'start'
	
	fmt.Println(duration)
}
```
Output:
```
2.005396666s
```

### Method `time.Duration`

Tipe `time.Duration` memiliki beberapa *method* yang sangat-sangat berguna untuk keperluan mengambil nilai durasinya dalam unit tertentu.

Misalnya, objek durasi tersebut ingin di-ambil nilainya dalam satuan unit detik, maka gunakan `.Seconds()`. Jika ingin dalam bentuk menit, maka gunakan .`Minutes()`, dst.

Caranya cukup akses saja *method*-nya, maka kita akan langsung dapat nilainya, tanpa perlu memikirkan operasi aritmatik konversinya.
```go
func main() {
	var start = time.Now()
	time.Sleep(time.Second * 2)

	var duration = time.Since(start)

	fmt.Println(duration.Milliseconds())
	fmt.Println(duration.Seconds())
	fmt.Println(duration.Minutes())
}

```
Output:
```
2001
2.00114475
0.0333524125
```
Contoh di atas, kita mengambil nilai durasi waktu dalam tiga bentuk, yaitu milidetik, detik, dan menit.
