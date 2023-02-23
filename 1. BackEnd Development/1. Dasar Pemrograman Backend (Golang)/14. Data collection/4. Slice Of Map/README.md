# Slice of map

## What is slice of map?
**Slice dan Map bisa dikombinasikan** dan ini juga sering digunakan pada banyak kasus, contohnya seperti data slice yang berisikan informasi siswa, dan banyak lainnya.

Cara menggunakannya cukup mudah, yaitu dibuat dengan pola seperti ini `[]map[string]int`, artinya kita membuat slice yang tipe dari setiap elemen-nya adalah `map[string]int` .

Agar lebih jelas, kita bisa praktekkan contoh di bawah ini:
```go
var students = []map[string]string{
	map[string]string{"name": "Aditira", "gender": "male"}, 
	map[string]string{"name": "Dito", "gender": "male"}, 
	map[string]string{"name": "Gina", "gender": "female"},
}
```

Jika kita menggunakan versi Golang terbaru, maka cara deklarasi slice-map bisa dipersingkat dengan tidak dituliskan tipe datanya di setiap elemen:
```go
var students = []map[string]string{
	{"name": "Aditira", "gender": "male"}, 
	{"name": "Dito", "gender": "male"},
	{"name": "Gina", "gender": "female"},
}
```

Dalam `[]map[string]string`, tiap elemen bisa saja memiliki `key` yang berbeda-beda, Contoh:
```go
var students = []map[string]string{
	{"name": "Aditira", "gender": "male", "hobby": "coding"},
	{"address": "jakarta", "id": "je02"},
	{"class": "camp2022"},
}
```

## Processing data

### Insert
Untuk menambahkan data pada slice, kita bisa memanfaatkan fungsi `append()` seperti pada slice namun diisi dengan tipe data *map*. Contoh:
```go
var students = []map[string]string{
	{"name": "Aditira", "gender": "male"}, 
	{"name": "Dito", "gender": "male"},
	{"name": "Gina", "gender": "female"},
}

var appendStudents = append(students, map[string]string{"name": "Ria", "gender": "female"})

for _, student range appendStudents {
	fmt.Println(student["gender"], student["name"])
}
```
Output:
```Output
male Aditira
male Dito
female Gina
female Ria
```

### Read

Untuk membaca data dari slice, kita cukup mengakses `index` pada *slice* dan mengakses `key` *map*-nya untuk mendapatkan nilai. Contoh:
```go
var students = []map[string]string{
	{"name": "Aditira", "gender": "male"},
	{"name": "Dito", "gender": "male"},
	{"name": "Gina", "gender": "female"},
}

fmt.Println(students[1]["name"])
fmt.Println(students[2]["gender"])
```
Output:
```Output
Dito
female
```

Kita juga bisa memanfaatkan *looping* `for` - `range` untuk membaca semua data pada *slice-map*:
```go
var students = []map[string]string{
	{"name": "Aditira", "gender": "male"}, 
	{"name": "Dito", "gender": "male"},
	{"name": "Gina", "gender": "female"},
}

for _, student := range students {
	fmt.Println(student ["gender"], student["name"])
}
```
Output:
```Output
male Aditira
male Dito
female Gina
```

Variabel `students` di atas berisikan informasi bertipe map`[string]string` dan tiap elemen-nya memiliki 2 *key* yang sama.

### Update
Untuk melakukan *update* nilai pada *slice-map* juga cukup mudah, kita bisa mengakses `index` dari *slice* diikuti dengan `key` dari *map* yang akan di-*update* dengan operator *assign* (` = `)
```go
var students = []map[string]string{
	{"name": "Aditira", "gender": "male"},
	{"name": "Dito", "gender": "male"},
	{"name": "Gina", "gender": "female"},
}

students[1]["name"] = "Deny"

fmt.Println(students[1]["name"])
```
Output:
```Output
Deny
```

Untuk update satu baris pada *slice* kita cukup akses pada satu `index` *target*, dan ubah dengan format *map*. Contoh:
```go
var students = []map[string]string{
	{"name": "Aditira", "gender": "male"}, 
	{"name": "Dito", "gender": "male"},
	{"name": "Gina", "gender": "female"},
}

students [2] = map[string]string{"name": "Tomi", "gender": "male"}

for _, student := range students {
	fmt.Println(student["gender"], student["name"])
}
```
Output:
```
male Aditira
male Dito
male Tomi
```


### Delete
Untuk melakukan *delete* pada *slice-map* kita bisa memanfaatkan fungsi `delete()` seperti pada *slice*, dengan *target* pada *slice* yang akan kita `delete`. Contoh:
```go
var students = []map[string]string{
	{"name": "Aditira", "gender": "male"},
	{"name": "Dito", "gender": "male"}, 
	{"name": "Gina", "gender": "female"},
}

delete(students[2], "gender")

for _, student := range students {
	fmt.Println(student["gender"], student["name"])
}
```
Output:
```
male Aditira
male Dito
Gina
```

Terlihat bahwa pada slice dengan index 2,  `gender` telah dihapus.
