# Package encryption

## Decode encode Base64
*encode* adalah proses penempatan urutan karakter (huruf, angka, dan symbol) tertentu ke dalam format khusus sehingga menjadi sebuah sandi. Sedangkan *decode* adalah proses sebaliknya yaitu konversi dari format yang disandikan, kembali pada karakter asli.
![index](https://raw.githubusercontent.com/Ouroboros-Tech/modul-pembelajaran/Dasar-golang/image/plainTexttoBase64.png) Go memiliki *package* `encoding/base64`, berisikan fungsi-fungsi untuk kebutuhan *encode* dan *decode* data ke *base64* dan sebaliknya. Data yang akan di-*encode* harus bertipe `[]byte`, perlu dilakukan casting untuk data-data yang belum sesuai tipenya.

## EncodeString & DecodeString
Fungsi `EncodeToString()` digunakan untuk *encode* data dari bentuk *string* ke *base64*. Fungsi `DecodeString()` melakukan kebalikan dari `EncodeToString()`.
Parameter di `EncodeToString()` harus di-*casting* terlebih dahulu ke dalam bentuk `[]byte` sebelum di-*encode* menggunakan fungsi tersebut.

Hasil *encode* adalah data *base64* bertipe `string`.

Sedangkan pada fungsi *decode* `base64.StdEncoding.DecodeString()`, data *base64* bertipe *string* di-*decode* kembali ke *string* aslinya, tapi bertipe `[]byte`. Kita bisa melakukan *casting* ke *string* dengan `string(data)`
```go
import "encoding/base64"
func main() {
	var data = "Budi Setiawan"

	// encode
	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded: ", encodedString)

	// decode
	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("Decoded: ", decodedString)
}
```
Output:
```Output
Encoded: VG9kIGVzdGF0aW9u
Decoded: Budi Setiawan
```

## Encode decode URL
Khusus *encode* data *string* yang isinya merupakan URL, lebih efektif menggunakan `URLEncoding` dibandingkan `StdEncoding`.
Dengan menggunakan *encode* di URL, data yang akan di kirimkan tidak akan mengandung karakter khusus di URL, seperti `/`, `?` dan `&` .
```go
func main() {
	var data = "Budi Setiawan"

	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded: ", encodedString)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("Decoded: ", decodedString)
}
```
Output:
```Output
Encoded: QnVkaSBTZXRpYXdhbg==
Decoded: Budi Setiawan
```

## Hash sha1
*Hash* adalah algoritma enkripsi untuk mengubah text menjadi deretan karakter acak. Jumlah karakter hasil hash selalu sama. *Hash* termasuk *one-way encryption*, membuat hasil dari hash tidak bisa dikembalikan ke text asli.
**SHA1** atau *Secure Hash Algorithm 1* merupakan salah satu algoritma *hashing* yang sering digunakan untuk enkripsi data. Hasil dari sha1 adalah data dengan lebar `20 byte` atau `160 bit`, biasa ditampilkan dalam bentuk bilangan heksadesimal `40 digit`.

Golang menyediakan *package* `crypto/sha1`, berisikan *library* untuk keperluan *hashing*.
```go
import "crypto/sha1"

// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	var text = "this is secret"
	var sha = sha1.New() // initialize sha1
	sha.Write([]byte(text)) // method 'Write' akan menampung data ke dalam sha1 yang nantinya akan di hashing
	var hash = sha.Sum(nil) // method 'Sum' akan melakukan hashing dari data yang sebelumnya diisi
	var hashedString = fmt.Sprintf("%x", hash) // casting hasil hashing bertipe []byte ke string
	fmt.Println(hashedString)
}

```
Output:
```Output
f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8
```
Kita perlu menggunakan fungsi `New()` dari *package* `sha1` untuk menginisialisasi *hash* tersebut. Hasil fungsi tersebut akan menjadi objek `hash.Hash`, yang memiliki 2 method yaitu:
* Method `write()`, digunakan untuk menge-*set* data yang akan di-*hash*. Data harus bertipe `[]byte`.
* Mehtod `Sum()`, digunakan untuk eksekusi proses *hashing* yang menghasilkan data *hash* bertipe `[]byte`. Method ini memerlukan parameter `[]byte` untuk saat ini bisa diisi dengan `nil`.

Sisanya, kita bisa menggunakan *formatting* dari `fmt.Sprintf()` untuk menampilkan hasil *hash* bertipe `[]byte` menjadi `string` atau melakukan *convert* dengan fungsi `string(hash)`.

## Salting method
*Salt* dalam konteks kriptografi adalah data acak yang digabungkan pada data asli sebelum proses *hash* dilakukan.

Ketika setiap kali melakukan *hash* untuk data yang sama, maka hasilnya pasti sama. Di sinilah kegunaan salt, teknik ini berguna untuk mencegah serangan menggunakan metode pencocokan data-data yang hasil *hash*-nya adalah sama (*dictionary attack*).
```go
func main() {
	var text = "this is secret"
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("%s %s", text, salt) // menggabungkan text dengan salt, menjadi: "this
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var hash = sha.Sum(nil)
	var hashedString = fmt.Sprintf("%x", hash)
	fmt.Println(hashedString)
}
```
Output 1: `66599fcac53b1b3c4b63b70c47b2c3ea8f96c68b`
Output 2: `a5b3224e47bcf43090957e74bbb593ea45959e85`
Output 3: `b7995374b92f523a84f16360a0a7315ccada952c`

*Salt* yang digunakan adalah hasil dari ekspresi `time.Now().UnixNano()`. Hasilnya akan selalu unik setiap nanodetiknya,

## `bcrypt`
Bcrypt merupakan fungsi *hashing* kata sandi yang dirancang berdasarkan *cipher* Blowfish.

Metode ini sering dipakai untuk enkripsi *password user*. *Salt* dan data hasil *hash* harus disimpan pada *database*, karena digunakan dalam pencocokan *password* setiap *user* ketika melakukan *login*.

Di Golang, kita bisa menggunakan *package* `bcrypt` dengan melakukan *import* dari *package* `golang.org/x/crypto/bcrypt`.

Terdapat 2 fungsi dasar yang bisa kita gunakan. Pertama adalah `bcrypt.GenerateFromPassword()`. Fungsi ini digunakan untuk meng-*enkripsi password* dengan mengisi 2 parameter. parameter 1 diisi data yang akan dienkripsi yang di-*casting* menjadi `[]byte`. Parameter 2 adalah *salt* yang kita pakai. 
Kedua, adalah `bcrypt.CompareHashAndPassword()`. Fungsi ini digunakan untuk mencocokkan *password* yang di-*inputkan* dengan *password* yang sudah di *hash*. Fungsi ini mengembalikan nilai *error*. Jika *password* yang di-*inputkan* sama dengan *password* yang di *hash*, maka nilai *error* akan `nil`.


```go
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	var password = "ini-rahasia-loh"
	
	// GenerateFromPassword => hash
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(hashPassword))

	// CompareHashAndPassword => match or not
	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte("ini-rahasia-loh"))

	if err != nil { // error akan terjadi (tidak berisi 'nil') jika password yang diinputkan tidak sama
		fmt.Println("Password salah")
	} else {
		fmt.Println("Password benar")
	}
}

```
Output:
```Output
$2a$10$ca7Cd/z0wS41KkcZn1ufMet7SPJxa8Xwx1gMaoQzm65S85Np7BCgC
Password benar
```
