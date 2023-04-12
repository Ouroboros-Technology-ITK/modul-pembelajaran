# JSON

## What is JSON?
*Javascript Object Notation* (JSON) adalah format pertukaran atau berbagi data yang ringan, mudah dibaca, dapat ditulis oleh manusia, serta mudah diterjemahkan dan dibuat oleh komputer. *JSON* diturunkan dari bahasa pemrograman Javascript, akan tetapi dapat digunakan untuk bebagai bahasa pemrograman termasuk Golang. JSON menggunakan ekstensi `.json`.

JSON memiliki struktur seperti `map` atau *dictionary*, dengan memiliki *key* (kunci) dan *value* (nilai) yang dibungkus oleh tanda kurung kurawal `{}`. *Value* di JSON mengikuti tipe data yang ada di bahasa pemrograman Javascript, yaitu `string`, `number`, `boolean`, `null`, `array`, dan `object`.

Contoh objek JSON:
```json
{
	"first_name": "John", // tipe data 'string'
	"last_name": "Doe",
	"age": 30, // tipe data 'number'
	"is_single": true, // tipe data 'boolean'
	"childres": null, // tipe data 'null'
	"hobbies": ["coding", "sports", "reading"], // tipe data 'array'
	// tipe data 'object'
	"address": {
		"street": "Jl. Kebon Kacang",
		"city": "Jakarta",
		"province": "DKI Jakarta",
		"postal_code": "12345"
	}
}
```
Meskipun di dalam *file* `.json` kita sering melihat format yang ditulis dalam beberapa baris seperti contoh di atas, JSON juga dapat ditulis disatu baris saja.

```json
{ "first_name": "Sammy", "last_name": "Shark", "online": true }
```
Referensi: [json.org](https://www.json.org/json-en.html).

JSON ini sangat penting dipelajari karena merupakan format pertukaran data yang paling banyak digunakan saat ini. JSON juga sering digunakan untuk pertukaran data antar aplikasi, antar aplikasi dengan *server*, dan antar *server* dengan *server*. Kita sebagai *backend developer* nantinya akan sering berhadapan dengan JSON.

## `encoding/json` package
Di Golang kita dapat menggunakan *package* `encoding/json` untuk mengolah data JSON ke `struct` dan sebaliknya.

Anggap terdapat *string* JSON sebagai berikut:
```json
{ "first_name" : "Sammy", "last_name": "Shark", "online" : true, }
```
Kita perlu membuat `struct` yang dapat merepresentasikan data JSON tersebut.
```go
type User struct {
	FirstName string
	LastName string
	Online bool
}
```
Setelah itu, kita perlu membuat *tag* `json` untuk mengatur bagaimana tiap *value* JSON tersebut akan ditransformasikan ke tiap properti `struct`. Cara menulisnya adalah dengan menambahkan sebuah string yang dibungkus <code>`</code> (*backtick*) dan referensi *key* JSON.

```go
type User struct {
	FirstName string `json: "first_name"` // properti 'FirstName' akan ditransformasikan ke key 'first_name'
	LastName string `json: "last_name"` // properti 'LastName' akan ditransformasikan ke key 'last_name'
	Online bool `json: "online"` // properti 'Online' akan ditransformasikan ke key 'online' di JSON
}
```

*Struct* tersebut sekarang dapat digunakan untuk menampung hasil *decode* JSON. Proses decode dapat dilakukan dengan fungsi `json.Unmarshal` dengan 2 parameter, parameter pertama adalah *string* JSON yang diubah menjadi `byte`, dan parameter kedua adalah `struct` yang akan diisi dengan data JSON menggunakan *pointer*. Fungsi ini mengembalikan *error* yang akan berisi `nil` jika berhasil melakukan *decode*, dan akan mengembalikan `error` jika terjadi kegagalan *decode*.
```go
package main

import (
	"encoding/json"
	"fmt"
)

var jsonData = `{ "first_name" : "Sammy", "last_name": "Shark", "online": true, }`

type User struct {
	FirstName string `"json: "first_name"`
	LastName  string `"json:"last_name"`
	Online    bool   `json: "online"`
}

func main() {
	var user User

	var err = json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
```
Output:
```
{Sammy Shark true}
```
Contoh di atas adalah bentuk *decode* JSON ke `struct`. Tiap nilai di data JSON akan di isi ke setiap properti `struct` sesuai dengan *tag* `json` yang telah kita buat. Nilai dari `first_name` di JSON akan diisi ke properti `FirstName` di` struct`, nilai dari `last_name` akan diisi ke properti `LastName`, dan nilai dari `online` akan diisi ke properti `Online`.

Kita juga dapat melakukan *decode* JSON ke `map` dengan mengubah parameter kedua adalah `map` yang akan diisi dengan data JSON menggunakan *pointer*. Namun, untuk mengatasi masalah nilai di JSON yang dinamis, kita harus menggunakan tipe data `interface{}`.
```go
package main

import (
	"encoding/json"
	"fmt"
)

var jsonData = `{ "first_name" : "Sammy", "last_name": "Shark", "online": true, }`

func main() {
	var user map[string]interface{}
	var err = json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
```
Output:
```
map[first_name: Sammy last_name: Shark online:true]
```
Setiap nilai di JSON akan diisi ke `map` dengan *key* yang namanya sesuai dengan *key* di JSON, dan nilainya juga.

Sebaliknya, kita juga dapat melakukan *encode* `struct` atau `map` ke JSON dengan menggunakan fungsi `json.Marshal`. Fungsi ini mengembalikan `byte` dan `error`.

`byte` yang dikembalikan adalah data JSON yang telah di *encode* dari `struct` atau `map` yang kita berikan. Kita dapat melihatnya dengan mengubah `byte` tersebut menjadi `string`.

Kembalian *error* akan berisi `nil` jika berhasil melakukan *encode* menjadi JSON, dan akan mengembalikan `error` jika terjadi kegagalan.
```go
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Online    bool   `json: "online"`
}

func main() {
	var user = User{
		FirstName: "Sammy",
		LastName:  "Shark",
		Online:    true,
	}

	var jsonData, err = json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData)) // casting ke string
}
```
Output:
```
"first_name":"Sammy", "last_name":"Shark", "online":true}
```

Berikut contoh jika menggunakan `map`:
```go
func main() {
	var user = map[string]interface{}{
		"first_name": "Sammy",
		"last_name":  "Shark",
		"online":     true,
	}

	var jsonData, err = json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
```
Output:
```
{"first_name":"Sammy", "last_name": "Shark", "online":true}
```


## Read file JSON
Kita dapat membaca *file* JSON dengan menggunakan fungsi `ioutil.ReadFile` dengan 1 parameter yaitu *path* atau posisi dari *file* JSON. Fungsi ini mengembalikan `byte` dan `error`. `byte` yang dikembalikan adalah data JSON yang telah dibaca dari file JSON yang kita berikan. Kita dapat melihatnya dengan mengubah `byte` tersebut menjadi `string`.

Contoh, terdapat *file* JSON bernama `user.json` dengan isi sebagai berikut:

```
{
	"first_name": "Sammy",
	"last_name": "Shark",
	"online": true
}
```
Kita dapat membaca *file* tersebut dengan menggunakan fungsi `ioutil.ReadFile` dengan 1 parameter yaitu nama *file* JSON.
```go
func main() {

	var jsonData, err = ioutil.ReadFile("user.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
```
Output:
```
{"first_name":"Sammy", "last_name": "Shark", "online":true}
```
Setelah itu, kita dapat melakukan *decode* JSON seperti yang telah kita pelajari sebelumnya. Karena fungsi `ioutil.ReadFile` mengembalikan `byte`, maka kita tidak perlu menggunakan *casting* `[]byte()` ke parameter pertama `json.Unmarshal()`.

```go
type User struct {
	FirstName string `json: "first_name"`
	LastName  string `json:"last_name"`
	Online    bool   `json: "online"`
}

func main() {
	// read file JSON
	var jsonData, err = ioutil.ReadFile("user.json")
	if err != nil {
		panic(err)
	}

	// decode JSON ke struct
	var user User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
```
Output:
```
{Sammy Shark true}
```


## Write, Update and Replace file JSON
Kita dapat menulis, memperbarui dan mengubah *file* JSON dengan menggunakan fungsi `ioutil.WriteFile` dengan 3 parameter yaitu:
1. *path* atau posisi dari *file* JSON
2. data JSON yang akan ditulis bertipe `byte`
3. *permission* dari *file* JSON.

Fungsi ini mengembalikan `error`. `error` yang dikembalikan akan berisi `nil` jika berhasil menulis, memperbarui dan mengganti *file* JSON, dan akan mengembalikan `error` jika terjadi kegagalan.

Namun, sebelumnya kita perlu melakukan *encode* `struct` atau `map` ke JSON dengan menggunakan fungsi `json.Marshal` yang telah kita pelajari sebelumnya.

Contoh, kita akan menulis JSON di file `user.json` dengan isi dari `struct` sebagai berikut:
```go
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Online    bool   `json: "online"`
}

func main() {
	var user = User{
		FirstName: "Sammy",
		LastName:  "Shark",
		Online:    true,
	}

	// encode dari struct ke JSON
	var jsonData, err = json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// write file JSON ke file 'user.json'
	err = ioutil.WriteFile("user.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success write file JSON")
}
```
Output:
```
Success write file JSON
```
Contoh di atas akan menulis data JSON di file `user.json`. `0644` adalah *permission* dari *file* `user.json`. Jika *file* `user.json` sudah ada, maka data JSON yang ada di dalamnya akan ditimpa dengan data JSON yang baru.
