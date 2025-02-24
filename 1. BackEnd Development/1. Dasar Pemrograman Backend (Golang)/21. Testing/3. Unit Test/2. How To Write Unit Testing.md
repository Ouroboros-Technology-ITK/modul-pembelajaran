# How to write unit testing

Cara membuat *unit test* berbeda-beda antara bahasa pemrograman. Beberapa teknik berikut dapat digunakan untuk membuat *unit test* dan diterapkan di semua bahasa pemrograman.

## AAA
AAA atau 3A pattern merupakan singkatan dari *Arrange*, *Act*, *Assert*. AAA adalah teknik yang digunakan untuk membuat *unit test*. AAA ini digunakan untuk memisahkan antara *Arrange*, *Act*, dan *Assert*.

*Arrange* adalah bagian di mana kita membuat data yang akan digunakan untuk *testing*.

*Act* adalah bagian di mana kita menjalankan fungsi yang akan kita *test*.

*Assert* adalah bagian di mana kita membandingkan hasil dari fungsi yang kita *test* dengan hasil yang diharapkan.

Contoh, terdapat kode di *package* `lamp` dengan file `lamp.go` sebagai berikut:
```go
package lamp
type Lamp struct {
	State bool
}

func (1 *Lamp) Turnon() {
	1.State = true
}

func (1 *Lamp) TurnOff() {
	1.State = false
}
```

Kita dapat menggunakan teknik AAA di *file* `lamp_test.go` sebagai berikut:
```go
package lamp
func TestTurnOn(t *testing.T) {
	// arrange
	lamp := Lamp{State: false} // lamp is off

	//act
	lamp.TurnOn() // turn on the lamp

	//assert
	if lamp.State != true {
		t.Errorf("lamp state should be true, got %v", lamp.State)
	}

	// act 2
	lamp.TurnOff() // turn off the lamp

	// assert 2
	if lamp.State != false {
		t.Errorf("lamp state should be false, got %v", lamp.State)
	}
}
```
Berdasarkan contoh di atas, kita dapat memisahkan antara *arrange*, *act*, dan *assert*. *Arrange* digunakan untuk menyiapkan dengan membuat *variable* `lamp` yang berisi *struct* `Lamp`.

*Act* digunakan untuk menjalankan *method* `TurnOn` dan `TurnOff`.

*Assert* digunakan untuk membandingkan hasil setelah `act` dijalankan, ketika sudah menjalankan *method* `TurnOn` maka nilai properti `State` harusnya jadi `true` dan setelah menjalankan *method* `TurnOff` harusnya jadi `False`.

## Avoid repeated AAA
Ketika terdapat *act* yang dipisahkan oleh *assert* atau *arrange*, pada contoh testing di atas terdapat tes untuk beberapa *behavior*. Sebagai contoh, pada kode di atas terdapat dua *behavior* yang dilakukan dalam satu tes. Ketika terdapat lebih dari satu *behavior* dalam *test*, lakukan refactor sehingga masing-masing tes hanya tes satu *behavior* saja.
```go
func TestTurnOn (t *testing.T) {
	//arrange
	lamp := Lamp[State: false}
	
	//act
	lamp.TurnOn()
	
	//assert
	if lamp.State != true {
		t.Errorf("lamp state should be true, got %v", lamp.State)
	}
}	
func TestTurnOff (t *testing. T) {
	//arrange
	lamp := Lamp[State: true}
	
	//act
	lamp.TurnOff()
	
	//assert
	if lamp.State != false {
		t.Errorf("lamp state should be false, got %v", lamp. State)
	}
}
```
## Given when then
Secara garis besar teknik ini sama dengan teknik AAA. Given berkaitan dengan *arrange*, when berkaitan dengan *act*, dan *then* berkaitan dengan *assert*.

Teknik ini dapat digunakan ketika terdapat *stakeholder non-technical* dari kode yang ditulis agar lebih mudah dibaca.

## Unit test naming
Memberi nama yang tepat pada *unit test* dapat meningkatkan *readability* pada kode. Salah satu cara penamaan yang dapat digunakan adalah sebagai berikut:
• Gunakan bahasa yang ekspresif dan mudah dipahami.
• Beri nama fungsi seolah mendeskripsikan skenario unit *test* kepada orang *non-technical*.
• Memisahkan antar kata dengan menggunakan `_` (underscores). Hal ini dapat meningkatkan *readability* ketika nama fungsi *unit test* relatif panjang.

Contoh:
```go
func TestWithdrawal_With_Enough_Balance_Should_Return_Amount (t *testing.T) {
// your code here
}
func TestWithdrawal_With_Not_Enough_Balance_Should_Return_Zero(t *testing.T) {
// your code here
}
```
