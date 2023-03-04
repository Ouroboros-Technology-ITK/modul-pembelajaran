# Interface

## What is interface?
*Interface* digunakan untuk mendefinisikan sekumpulan *method signature*.

*Method signature* adalah sebuah signature yang mengandung nama method, parameter dan kembalian nilai (`return`) yang diberikan.

Cara implementasi *interface* ar Golang adalah secara implisit. Jika suatu objek memiliki *method signature* yang sama dengan *interface*, maka objek tersebut mengimplementasikan *interface* terkait.

## Interface in Go
*Interface* di Golang merupakan tipe data. *Interface* digunakan untuk mendeskripsikan apa yang dapat dilakukan oleh objek (*function*).

Untuk mendefinisikan interface dapat menggunakan keyword `type` diikuti dengan nama interface dan kemudian menulis *keyword* `interface`.

Implementasi *interface* di Golang biasanya digunakan untuk membuat sebuah kontrak atau penampung tipe data objek (*struct*), yang nantinya dapat diisi oleh objek apapun yang memiliki *method* sama dengan *interface*.
![index](http://donofden.com/images/doc/golang-interface-1.png)

Struktur Interface (Sumber: [donofden](http://donofden.com/blog/2019/09/20/golang-interface-goroutine-channels))

Di dalam *block statement* `interface` dapat diisi dengan *method* yang ingin dideskripsikan.

Dengan memanfaatkan `interface` di Golang, kita dapat mendefinisikan sekumpulan *method signature* dalam sebuah wadah bertipe *interface* (*a value of interface type*).

Contoh :
```go
type Geometry interface {
	Area() float64
	Circumference() float64
}
```

Dapat dikatakan dengan bahasa manusia bahwa semua objek atau `struct` yang memiliki fungsi Area dan fungsi `Circumference` termasuk dalam *interface* `Geometry`. Sebagai contoh berikut:

* Persegi / rectangle
```go
// struct dan method signature dari 'Rectangle'
type Rectangle struct {
	Side float64
}

func (r Rectangle) Area() float64 {
	return r.Side * r.Side
}

func (r Rectangle) Circumference() float64 {
	return 4 * r.Side
}
```
* Lingkaran / circle
```go
// struct dan method signature dari 'Circle'
type Circle struct {
	Radius float64
}
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Circumference() float64 {
return 2 * 3.14 * c.Radius
}
```
Kemudian, kita bisa menggunakan *interface* `Geometry` untuk menampung semua `struct` memiliki *method signature* yang didefinisikan sebelumnya.

Contoh implementasi :
```go
func main() {

	var shape Geometry // variable 'shape' dapat diisi struct apapun selama memiliki method 'Area' dan
	shape Rectangle Side: 10 } // diisi dengan struct 'Rectangle'
	fmt.Println("Luas : ", shape.Area())
	fmt.Println("Keliling :", shape.Circumference())
	
	shape Circle{ Radius: 5 } // diisi dengan struct 'Circle'
	
	fmt.Println("Luas :", shape.Area())
	fmt.Println("Keliling :", shape.Circumference())
}
```
Output:
```Output
Luas 100
Keliling: 40
Luas 78.5
Keliling 31.4
```

Perhatikan kode di atas. *Variable* objek `shape` bertipe *interface* `Geometry`, digunakan untuk menampung objek konkrit yaitu *struct* `Rectangle` dan `Circle`.

Dari variabel tersebut, *method* `Area` dan `Circumference` diakses. Secara otomatis Golang akan mengarahkan pemanggilan *method* pada *interface* ke *method* asli milik `struct` yang bersangkutan. Ketika diisi dengan struct `Rectangle`, maka *method* `Area` mengarah ke `Rectangle.Area` dan `Circumference` mengarah ke `Rectangle.Circumference`. Demikian juga dengan `Circle`.

Perlu diketahui juga, jika ada *interface* yang menampung objek konkrit di mana struct-nya tidak memiliki salah satu *method* yang terdefinisi di *interface*, maka akan muncul **error**. Intinya kembali ke aturan awal, variabel *interface* hanya bisa menampung objek yang minimal memiliki semua *method* yang terdefinisi di interface-nya.

Contoh, terdapat *struct* `Triangle` yang hanya memiliki *method* `Area` :
```go
type Triangle struct {
	Width float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 1/2 t.Width * t.Height
}

func (t *Triangle) ChangeHeight(newHeight float64) {
	t.Height = newHeight
}

func (t *Triangle) ChangeWidth (newwidth float64) { 
	t.Width = newwidth
}

```
Ketika kita coba implementasi di `main`, maka akan terjadi *error*, karena struct `Triangle` tidak memiliki semua *method* yang didefinisikan di *interface* `Geometry`, yaitu tidak memiliki *method* `Circumference`.
```go
func main() {
	var geometry Geometry
	
	geometry = Triangle{ Width: 15, Height: 20}
	// error
	// struct 'Triangle' tidak memenuhi semua method yang dibutuhkan di interface 'Geometry'
}
```

## Embedded interface
*Interface* bisa di-*embed* ke *interface* lain, sama seperti `struct`. Cara penerapannya cukup dengan menuliskan nama *interface* yang ingin di-embed ke dalam *interface* tujuan.
```go
type Geometry interface {
	Area() float64
	Circumference() float64
}

type ThreeDimension interface {
	Geometry
	Volume() float64
}
```

Interface `Geometry` berisikan *method* `Area` dan `Circumference`, sedang `ThreeDimension` berisikan *method* `Volume` dan di-*embed* dengan interface `Geometry` sehingga memiliki kemampuan yang sama dengan `Geometry` ditambah dengan *method* miliknya sendiri.
```go
type Cube struct {
	Side float64
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}
func (c Cube) Area() float64 {
	return math.Pow(c.Side, 2) * 6
}
func (c Cube) Circumference() float64 {
	return c.Side * 12
}
func (c *Cube) ChangeSide (newSide float64) {
	c.Side = newSide
}
```
Objek hasil cetakan `struct` di atas, nantinya akan ditampung oleh objek *interface* `ThreeDimension` yang memiliki objek *interface* `Geometry`.
```go
func main() {
	var shape ThreeDimension

	shape = Cube Side: 10 }
	
	fmt.Println("Luas :", shape.Area())
	fmt.Println("Keliling :", shape.Circumference())
	fmt.Println("volume :", shape.Volume())
}
```
Output:
```Output
Luas 600
Keliling 120
Volume: 1000
```
