# URL Parsing


## `net/url` package
*Parsing* URL adalah proses untuk mengurai atau mendapatkan informasi dari setiap anatomi URL, seperti *schema*, *hostname*, *path*, dll. Golang dapat melakukan *parsing* URL dengan mudah menggunakan *package* `net/url`. *Package* ini memiliki fungsi `Parse` yang menerima *string* URL dan mengembalikan sebuah *struct* `*URL` dan `error`. `*URL` memiliki beberapa atribut yang dapat digunakan untuk mengakses informasi URL. Dan `error` akan bernilai `nil` jika URL yang dimasukkan *valid*.
```go
import "net/url"
```

Contoh, kita memiliki URL `https://www.google.com/search?q=golang` dan ingin mengakses informasi dari URL tersebut.
```go
import "net/url"

var urlString = "https://www.google.com/search?q=golang"

func main() {
	var uri, err = url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	fmt.Println(uri.Scheme)   // mendapatkan informasi schema dari URL
	fmt.Println(uri.Host)     // mendapatkan informasi hostname dari URL
	fmt.Println(uri.Path)     // mendapatkan informasi path dari URL
	fmt.Println(uri.RawQuery) // mendapatkan informasi query dari URL
}
```

```
> go run main.go
https
www.google.com
/search
q=golang
```

Kita bisa mengakses informasi URL dengan menggunakan setiap atribut URL dari struct `*URL`, seperti:
* `Scheme` dengan mengembalikan *scheme* atau protokol dari URL
* `Host` dengan mengembalikan *host* dari URL
* `Path` dengan mengembalikan *path* dari URL
* `RawQuery` dengan mengembalikan *data* query dari URL
Kita juga dapat mengakses tiap *query* dari URL dengan menggunakan fungsi `Query` dari `*URL`. Fungsi `Query` mengembalikan `map[string][]string` yang berisi *key* dan *value* dari *query*.
```go
import "net/url"

var urlString = "https://www.google.com/search?q=golang&hl=id"

func main() {
	var uri, err = url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	fmt.Println(uri.Query()["q"][0])  // mendapatkan informasi query 'q'
	fmt.Println(uri.Query()["h1"][0]) // mendapatkan informasi query 'hl'
}
```

```
> go run main.go
golang 
id
```

Contoh di atas adalah mendapatkan *value* dari *query* `q` dan `hl`.
