# Map
## What is map?
Bila kita ingat pada **array** atau **slice**, untuk mengakses data, kita menggunakan **index** *number* dimulai dari **0**.  Untuk mengakses data, kita perlu tahu *index number*-nya.

Mudah diingat untuk data sederhana, seperti jumlah murid berdasarkan kelas di satu sekolah. Kita bisa menyimpan data jumlah murid tersebut dengan *index* 0 untuk kelas 1, index 1 untuk kelas 2, dst.

Namun untuk data yang lebih kompleks, sulit untuk mengingat *index number*-nya. Contohnya jumlah murid laki-laki dan murid perempuan di suatu sekolah berdasarkan kelas. Sulit bukan untuk mengingat index number mana yang menyimpan data murid laki-laki dari kelas 2 misalnya.

Untuk data yang lebih kompleks, kita bisa menggunakan *map*, di mana kita bisa menentukan jenis tipe data *index* yang akan kita gunakan.

Sederhananya, map adalah tipe data kumpulan `key-value` (kata kunci-nilai), di mana kata kuncinya bersifat unik, tidak boleh sama

Berbeda dengan *Array* dan *Slice*, jumlah data yang kita masukkan ke dalam *Map* boleh sebanyak- banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru.

Fungsi pada map yaitu:
OPERASI|KETERANGAN
--|--
`len (map)`|Untuk mendapatkan jumlah data di map  
`map[key]`|Mengambil data di map dengan key 
`map[key] = value`|Mengubah data di map dengan key
`make(map[TypeKey]TypeValue)`|Membuat map baru
`delete(map, key)`|Menghapus data di map dengan key

## Key dan value
Map adalah kumpulan pasangan key value yang tidak berurutan. Setiap value memiliki key yang artinya setiap key haruslah unik, namun value yang sama dapat digunakan berkali-kali.
![index](https://miro.medium.com/v2/resize:fit:828/format:webp/1*MF82DTElbNKV29T22_V8Ug.png)

Map Example (Sumber: [miro.medium.com](https://levelup.gitconnected.com/composite-data-types-in-golang-a829288b5553))

Map di Golang bisa dibuat dengan syntax sebagai berikut:
```go
a := map[KeyType]ValueType(key1: value1, key2: value2,...}
```

Sebagai contoh, jika kita ingin membuat map dengan key string dan value string, kita dapat membuatnya seperti ini:
```go
// cara pertama
var scoreMap map[string]string
// cara kedua
scoreMap := make(map[string]string)
// cara ketiga
scoreMap := map[string]string{
	"A": "100-86",
	"B": "75-86",
	"C": "50-74",
	"D": "30-49",
	"E": "0-29",
}
```

## Insert
*Zero value* dari *Map* adalah *nil*, maka tiap variabel bertipe *map* harus di-inisialisasi secara explisit nilai awalnya (agar tidak nil). Contoh:
```go
var data map[string]int
data["one"] = 1
// akan muncul error!
fmt.Println(data)
```
Output:
```go
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
		/home/aditira/demo/main.go:5 +0x2e
exit status 2
```

Agar tidak *error*, setelah *variable* `data` kita bisa menambahkan operator assign (`=`), dilanjutkan dengan tipe data map dan diikuti dengan tanda kurung kurawal `{}`, lalu setelahnya kita bisa menambahkan *key* dan *value* baru. Contoh:
```go
var data = map[string]int{}
data["one"] = 1
// tidak ada error
fmt.Println(data)
```
Output:
```Output
map[one:1]
```

Seperti pada *array* dan *slice*, *map* juga dapat di inisialisasi dengan cara *horizontal* maupun *vertical*:
```go
// cara horizontal
var months1 = map[string]int{"januari": 50, "februari": 40}
// cara vertical
var months2 = map[string]int{
	"januari": 50,
	"februari": 40,
}
```

`key` dan `value` dituliskan dengan pembatas tanda titik dua (` : `). Sedangkan tiap itemnya dituliskan dengan pembatas tanda koma (`, `). Khusus deklarasi dengan gaya vertikal, tanda koma perlu dituliskan setelah item terakhir.

`Variabel map` bisa di-inisialisasi dengan tanpa nilai awal, caranya menggunakan tanda kurung kurawal, contoh: `map[string]int{}`. Atau bisa juga dengan menggunakan *keyword* `make`. Contohnya:
```go
var months3 = map[string]int{}
var months4 = make(map[string]int)
```

## Read
Kita dapat mengakses elemen map dengan syntax. `value = map_name [key]`. Sebagai contoh pada kode berikut:
```go
scoreMap := map[string]string {
	"A": "100-86",
	"B": "75-86",
	"C": "50-74",
	"D": "30-49",
	"E": "0-29",
}

score := scoreMap["A"]
fmt.Println(score)
```
Output:
```Output
100-86
```

Namun, jika `key-value` tidak ada, maka *output* akan kosong:
```go
score := scoreMap["F"]
```
Output:
```
```

Cara lebih aman untuk mengakses `key-value` pada *map* adalah dengan mengecek apakah `key` pada *map* tersedia atau tidak:
```go
score, ok := scoreMap["F"]
if !ok {
	log.Println("Key tidak tersedia")
}
```


## Iterasi Item Map Menggunakan `for` - `range`
Cara penerapannya masih sama seperti pada *slice*, bedanya adalah data yang dikembalikan di tiap perulangan adalah `key` dan `value`, bukan `indeks` dan `elemen`. 
Contoh:
```go
var months = map[string]int{
	"januari": 50,
	"februari": 40,
	"maret": 34,
	"april": 67,
}

for key, val := range months {
	fmt.Println(key, " \t:", val)
}

```
Output:
```
januari			: 50
februari		: 40
maret			: 34
april			: 67
```

## Update
Untuk update nilai pada map, assign nilai baru ke `map[key]` menggunakan *assignment operator* (`=`).

Jika `key` yang diperbarui sudah ada, maka nilai tersebut akan diperbarui dengan nilai baru. *Syntax* nya adalah sebagai berikut untuk update `key` pada *Map* `x`:
```go
x[key] = new_value
```
Contoh:
```go
package main
import "fmt"
func main() {
	fruits: map[string]string{
		"a": "apple",
		"b": "banana",
	}

	fruits["a"] = "avocado"
	
	fmt.Print(fruits)
}
```
Output:
```go
map[a: avocado b:banana]
```

## Delete
Fungsi `delete()` digunakan untuk menghapus *item* dengan *key* tertentu pada *map*. Cara menggunakannya, dengan memasukan objek *map* dan *key* item yang ingin dihapus sebagai parameter. Contoh:
```go
var months = map[string]int("januari": 31, "februari": 28}
						  
fmt.Println(len(months)) // 2
fmt.Println(months)
						  
delete(months, "januari")
						  
fmt.Println(len(months)) // 1
fmt.Println(months)
```
*Item* yang memiliki `key` "januari" dalam variabel `months` akan dihapus.

Output:
```Output
2
map[februari:28 januari:31]
1
map[februari:28]
```
>Fungsi `len()` jika digunakan pada map akan mengembalikan jumlah item.
