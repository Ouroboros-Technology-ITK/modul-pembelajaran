# HTTP Client


## Package `net/http`

*Package* `net/http` bisa dibagi menjadi 2 pembahasan, *package* `net` dan *package* `http`.

## Package `net`
Package `net` merupakan implementasi dari *network/transport* layer *I/O*, termasuk TCP/IP, UDP dan *domain name resolution*. Selengkapnya kita bisa lihat [di sini]()

Berikut adalah `Dial` *function* untuk *connect* ke server.

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// membuat koneksi ke server golang.org melalui tcp network
	conn, err := net.Dial("tcp", "golang.org:80")
	// periksa jika ada error saat koneksi ke server
	// akan keluar dari program 'log.Fatal(...)`
	if err != nil {
		log.Fatal("The following error occured", err)
	}
	// koneksi tersambung, kirim GET request ke server
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// kirim request "enter" untuk mengakhiri koneksi
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("Send request error ", err)
	}
	
	fmt.Println("Response: ", status)
}
```
Output:
```
Response: HTTP/1.0 200 OK
```

Dari contoh *code* di atas, kita mengirim *request* GET menggunakan protokol HTTP, melalui `tcp` *network* ke *server domain* `golang.org` *port* `80`.


## Package `http`

Kalau *package* `net` untuk **network/transport** *layer*, maka *package* `http` adalah implementasi dari **application layer HTTP** protokol yang berjalan di atas *layer network* tersebut.

*Package* `http` menyediakan HTTP (Hypertext Transfer Protocol) *client* dan *server implementation*. Mari kita lihat kode implementasi dari *request* http **GET**, **POST** dan **PostForm** di bawah ini:

```go
resp, err := http.Get("http://example.com/")
...
resp, err := http. Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http. PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
```

Kita bisa mencoba implementasi HTTP tersebut menggunakan [Postman Echo]() yaitu *service* yang dapat digunakan untuk menguji **REST client** dan membuat contoh panggilan API.


## HttpClient
HttpClient adalah sebuah cara untuk mengirimkan HTTP *request* dan menerima respon dari server yang dituju. Jadi, aplikasi *client* akan melakukan *http request* ke aplikasi server, pada *endpoint* sesuai dengan spesifikasi yang dijelaskan di bawah.


## NewRequest
*Package* `net/http`, selain berisikan *tools* untuk keperluan pembuatan web, juga berisi fungsi-fungsi untuk melakukan *http request*. Salah satunya adalah `http.NewRequest()`.

Fungsi `http.NewRequest()` digunakan untuk membuat *request* baru. Fungsi tersebut memiliki 3 parameter yang wajib diisi.
* Parameter pertama, berisikan tipe *request* **POST** atau **GET** atau lainnya
* Parameter kedua, adalah **URL** tujuan *request*
* Parameter ketiga, form data *request* (jika ada)

Fungsi tersebut menghasilkan *instance* bertipe `http.Request`. *Object* tersebut nantinya di sisipkan pada saat eksekusi *request*.

Kita bisa gunakan *3rd Party API* untuk melakukan *request*, contohnya di sini: [pokeapi.co]()

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Statement yang menghasilkan instance http.Client, diperlukan untuk eksekusi request
	var client = &http.Client{}

	// http.NewRequest() digunakan untuk membuat request baru
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/ability/?limit=1", nil)
	if err != nil {
		panic(err)
	}

	// Request dengan method Do() pada instance http.Client yang sudah dibuat
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Cetak status code request
	fmt.Println("Status: ", resp.Status)

	// Kita bisa membaca response body menggunakan package ioutil.
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Convert body type menjadi string dan cetak
	fmt.Println(string(responseData))
}
```
Output:
```
Status: 200 OK
{"count":327,"next": "https://pokeapi.co/api/v2/ability/?offset=1&limit=1","previous":null,"results":[{"
```

## GET
*Method* **GET** akan menampilkan data/nilai pada URL yang di *request*, contohnya kita bisa gunakan 3rd Party API untuk melakukan request, salah satunya menggunakan: [postman.com]()
```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// http.Get Implementation:
	resp, err := http.Get("https://postman-echo.com/get?fool-bar1&foo2=bar2")
	if err != nil {
		log.Fatal(err)
	}

	// Kita bisa membaca response body menggunakan package ioutil.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Convert body type to string
	sb := string(body)
	log.Printf(sb)
}
```
Output:

```
2022/03/21 16:41:00 ("args":{"fool": "bar1","foo2": "bar2"}, "headers":{...}
```

Dari contoh di atas, `http.Get` mengembalikan 2 *variable* yaitu `resp` dan `err`. Di mana `resp` adalah *respond* yang akan didapat setelah melakukan *request*. kita bisa membaca *body* dari *responds* menggunakan fungsi `ReadAll` dari `ioutil`.

Kita juga dapat menerjemahkan *response* menjadi `struct` dengan cara sebagai berikut:
* Pertama siapkan dulu struct untuk mengambil beberapa field yang dibutuhkan
```go
type headers struct {
	XForwardedProto string `json: "x-forwarded-proto"`
	XForwardedPort  string `json: "x-forwarded-port"`
	Host            string `json: "host"`
}

type JsonResponse struct {
	Headers headers
}
```

* Pada fungsi `main()` lakukan *decode response body* menggunakan fungsi `json. Unmarshal()` di bawah ini:
```go
var jsonResponse JsonResponse
err = json.Unmarshal(body, &jsonResponse)

if err != nil {
	panic(err)
}

fmt.Println(jsonResponse)
```
Output:
```
> {{https 443 postman-echo.com}}
```

## POST
*Method* POST akan mengirimkan data atau nilai langsung ke *body* untuk ditampung, tanpa menampilkan pada URL, contoh:
```go
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Encode data:
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Toby",
		"email": "Toby@example.com",
	})
	responseBody = bytes.NewBuffer(postBody)
	// http.Post Implementation:
	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
```

Dari contoh di atas `http.Post` hampir sama dengan `http.Get` bedanya dia memiliki 3 parameter yaitu: `(url string, contentType string, body io.Reader)`
* **Url** adalah alamat dari sebuah server
* **ContentType adalah** tipe data yang akan di kirim, pada contoh di atas berupa tipe data JSON
* **Body** adalah data yang akan dibaca oleh *server*

Sehingga *output* yang akan dihasilkan adalah:
```
2022/03/21 16:50:52 {"args":{},"data":{"email":"Toby@example.com", "name": "Toby"},"files":{},"form":{},"
```


## Secure HTTP request
Komunikasi antara aplikasi *client* dan aplikasi *server* dapat terjadi dengan 2 cara, yaitu *insecure request* (http) dan *secure request* (https).

*Secure request* adalah bentuk *request* yang datanya ter-enkripsi, dan saat ini sudah menjadi *default* untuk *http request* antar aplikasi *server* dan *client*.

*Request* jenis ini pada sisi *client* atau *consumer* membutuhkan konfigurasi di mana *file certificate* diperlukan.


## Insecure HTTP request
Pada beberapa kasus tertentu, bisa jadi *client* **tidak memiliki** *certificate* namun komunikasi ingin tetap dilakukan, masih bisa dilakukan (dengan catatan server meng-*allow* kapabilitas ini), yaitu menggunakan teknik **insecure request**.

>Dalam *insecure request*, komunikasi ( pertukaran data) terjalin tanpa ada proses enkripsi data.

Dari sisi *client*, cara membuat *insecure request* sangat mudah, cukup aktifkan atribut *insecure* pada *request*. Misalnya menggunakan `curl`, maka cukup tambahkan *flag* `--insecure` pada command:
```shell
curl -X POST https://<domain>/data \
	--insecure \
	 -H 'Content-Type: application/json' \
	 -d '{"Name": "Aditira Jamhuri"}'
```
Untuk pembahasan lengkap dari *Insecure Request* dan *Secure Request*, kita bisa kunjungi tutorialnya [di sini]()
