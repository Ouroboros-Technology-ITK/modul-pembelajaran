# Pseudocode
## What Is Pseudocode?
<p align="justify">
Pseudocode adalah cara penulisan algoritma yang hampir menyerupai bahasa pemrograman, namun pseudocode ditulis dengan lebih sederhana menggunakan bahasa baku yang mudah dipahami oleh manusia. Kita dapat menggunakan bahasa apapun dalam membuat pseudocode, namun karena kita nantinya akan menggunakan bahasa pemrograman yang umumnya berbahasa Inggris, maka di materi ini akan menggunakan pseudocode <strong>dalam bahasa Inggris</strong>.<br>

<p align="justify">
Contoh bentuk pseudocode sederhana : <strong>Menghitung luas persegi panjang</strong>.<br>

```
PROGRAM AreaRectangle

READ AND WRITE "width" with number
READ AND WRITE "height" with number
READ AND WRITE "total" with number
STORE "total" with CALCULATE "width" multiple by "height"
PRINT "total"
```

> <strong>NOTE: Jangan terpaku dengan contoh, karena dalam pseudocode tidak harus menggunakan kata tertentu. Selama pseudocode dapat dipahami oleh pembaca, berarti sudah benar.</strong>

## Variable And Data Type
<p align="justify">
Kita tidak asing dengan istilah variabel dan data, apalagi saat pelajaran matematika. Namun apa sih kedua hal tersebut?<br>

<p align="justify">
Mari kita analogikan dengan contoh sehari hari. Jika kita memegang sebuah minuman es jeruk, maka penampung atau gelas air jeruk tersebut adalah sebuah <strong>variabel</strong>, sedangkan <strong>tipe data</strong> adalah air jeruknya.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/jus-jeruk-1.png"><br>

<p align="justify">
Kita bisa menggunakan analogi lain untuk mecontohkan data type.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/jus-jeruk-2.png"><br>

<p align="justify">
Bentuk gelasnya sama tapi di gelas A diisi dengan kopi, gelas B diisi dengan air jeruk, dan gelas C diisi dengan susu. Hal ini menunjukkan bahwa variabel dapat diisi dengan tipe data berbeda tapi dalam jenis yang sama (air).<br><br>

## Variable
<p align="justify">
<strong>Variabel</strong> adalah suatu pengenal (identifier) yang digunakan untuk mewakili suatu nilai tertentu di dalam program dan dapat diisi nilai yang bisa diubah sesuai kebutuhan. Idealnya dalam penamaan sebuah variabel dapat merepresentasikan kegunaan dari variabel tersebut.<br>

<p align="justify">
Kita dapat menggunakan kata "READ" untuk membaca nilai dari sebuah variabel, dan kata "WRITE" untuk menulis nilai dari sebuah variabel.<br>

## Data Type
<p align="justify">
<strong>Data Type</strong> atau tipe data adalah jenis data atau nilai yang dapat ditampung dan diolah oleh sebuah variabel. Kita dapat menggunakan "STORE" atau "SET" untuk mengisi variabel dengan tipe data tertentu.<br>

<p align="justify">
Tipe data dibagi menjadi 2 kelompok besar yaitu:<br>

1. <strong>Tipe data primitif dikelompokkan sebagai berikut:</strong>

- <strong>Tipe Data Bilangan (number)</strong>
<p align="justify">
Merupakan tipe data yang berisi bilangan bulat, desimal, atau bilangan pecahan. Contohnya adalah 10, 10, 1.5, 200.000.000.<br>

- <strong>Tipe Data String</strong>
<p align="justify">
Merupakan tipe data yang berisi kumpulan karakter karakter, biasanya dibungkus menggunakan tanda petik (") atau petik satu ('). Contohnya adalah "kamu", "kenangan indah bersamanya", "i love you 3000", "c3m4ng4t".<br>

- <strong>Tipe Data Boolean</strong>
<p align="justify">
Merupakan tipe data yang hanya berisi nilai true atau false. Contohnya adalah true, false.<br><br>

2. <strong>Tipe data terstruktur dikelompokkan sebagai berikut:</strong>

- <strong>Tipe Data Array/larik</strong>
<p align="justify">
Merupakan kumpulan tipe data yang sama dan memiliki index yang unik. Contohnya adalah [1, 2, 3, 4, 5]<br>

- <strong>Tipe Data Record</strong>
<p align="justify">
Merupakan kumpulan tipe data yang berisi data berbeda. Contohnya adalah {name: "John", age: 30}.<br>

- <strong>Tipe Data Set</strong>
<p align="justify">
Merupakan kumpulan tipe data yang tidak memiliki index tapi harus unik (tidak boleh ada yang duplikat satu sama lain). Contohnya adalah {1, 2, 3, 4, 5}.<br>

<p align="justify">
Contoh pseudocode menggunakan variabel dan data type:<br>

```
READ AND WRITE "name" with string 
READ AND WRITE "age" with number 
STORE "name" with "ruangguru"
STORE "age" with 25
PRINT "name"
PRINT "age"
```

<br>

## If Condition
<p align="justify">
Dalam aktivitas sehari hari, kita sering dihadapkan sebuah kondisi untuk memilih beberapa pilihan dan keputusan tersebut akan menghasilkan hasil yang berbeda dengan keputusan yang lain.<br>

<p align="justify">
Sebagai contoh, Pak Budi sedang lapar dan ingin makan ke warung, terdapat 2 menu yaitu makanan A seharga Rp. 25.000 dan makanan B seharga Rp. 19.000. Saat itu Pak Budi hanya membawa uang Rp. 20.000 dan merasa makanan A sangat menggiurkan. Maka Pak Budi akan mengecek apakah uang tersebut cukup untuk membeli makanan A atau tidak, jika cukup maka Pak Budi akan membeli makanan A, jika tidak maka Pak Budi akan membeli makanan B.<br>

```
STORE "money" with number 20000
STORE "price_food_A" with number 25000

IF "money" more than or equal "price_food_A"
    PRINT "Membeli makanan A"

ELSE
    PRINT "Membeli makanan B"
```

<p align="justify">
Contoh di atas adalah bentuk condition, kita bisa menggunakan kata "IF" untuk menentukan apakah kondisi tersebut benar dan kita bisa menggunakan "ELSE" untuk kondisi yang salah.<br><br>

## Looping
<p align="justify">
Umumnya kita sering melakukan aktivitas berulang (aktivitas yang sama berkali-kali. Contohnya seperti membuat kopi. Kita pasti akan mengaduk kopi yang sudah dituang dengan air panas berulang kali sampai larut. Aktivitas ini yang kita sebut dengan perulangan / Looping.<br>

```
WHILE "hungry"
    DO "eat"
```

<p align="justify">
Pseudocode di atas cukup simple dan mencontohkan kita proses paling sederhana dalam looping. WHILE adalah standard keyword untuk menunjukkan sebuah perulangan. Dari contoh di atas dapat diartikan bahwa selama masih "hungry" maka lakukan "eat".<br>

<p align="justify">
Pseudocode di atas cukup simple dan mencontohkan kita proses paling sederhana dalam looping. WHILE adalah standard keyword untuk menunjukkan sebuah perulangan. Dari contoh di atas dapat diartikan bahwa selama masih "hungry" maka lakukan "eat".<br>

```
STORE "full level" with e

WHILE "full level" < 5
    ADD "full level" by 1

DISPLAY "I'm full!"
```

<p align="justify">
Pseudocode di atas menggambarkan kita mulai dari level kenyang kita dari 0, berarti kita saat ini sangat lapar. Setiap kali kita melakukan proses makan, tingkat kenyang kita akan bertambah 1. Karena kita hanya kuat makan hingga 5 kali, maka kondisinya adalah "full level" < 5. Contoh ini hampir mendekati pemahaman bahasa pemrograman. Terdapat awalan sebelum looping dilakukan dan terdapat kondisi untuk menghentikan looping.<br><br>

## Procedure
<p align="justify">
Procedure atau prosedur adalah program kecil atau subprogram yang berada di program utama untuk menyelesaikan masalah khusus dengan menggunakan atau tidak menggunakan parameter.<br>

<p align="justify">
Parameter adalah data masukan untuk subprogram yang nantinya akan diproses lebih lanjut di dalam prosedur.<br>

<p align="justify">
Kita menggunakan kata "PROCEDURE" diikuti dengan nama prosedur dan parameter untuk membuat sebuah prosedur, dan kita tutup alur prosedur dengan "END PROCEDURE". Untuk memanggil prosedur yang sudah dibuat dapat menggunakan keyword "CALL".<br>

```
PROCEDURE AddTwoNumber (number1, number2) 
    STORE "result" with number1 plus number2 
    PRINT "result"
END PROCEDURE

CALL AddTwoNumber(5, 10)
```

<p align="justify">
Pseudocode di atas merupakan bentuk prosedur dengan nama "AddTwoNumber" atau menambahkan dua buah angka. Terdapat 2 parameter yaitu "number1" dan "number2" yang nantikan dapat kita isi dengan nilai yang ingin kita tambahkan dan di dalamnya terdapat proses penambahan kemudian menampilkan hasil penambahan tersebut.<br><br>

## Function
<p align="justify">
Sedikit berbeda dengan procedure, function atau fungsi juga digunakan untuk menyelesaikan masalah khusus, namun dapat mengembalikan nilai atau data yang dapat ditampung oleh variabel atau digunakan dalam program utama. Fungsi dapat dipanggil di dalam prosedur.<br>

<p align="justify">
Kita bisa menggunakan kata "FUNCTION" diikuti dengan nama fungsi dan parameter, kemudian ditutup dengan "END FUNCTION". Untuk mengembalikan nilai atau hasil dari dalam fungsi, kita bisa menggunakan keyword "RETURN".<br>

```
FUNCTION AreaRectabgle (side)
    STORE "result" with "side" times "side"
    RETURN "result"
END FUNCTION

STORE "area" with AreaRectabgle(5)
PRINT "area"
```

<p align="justify">
Pseudocode di atas adalah contoh penerapan fungsi dengan nama "AreaRectangle" atau luar persegi. Terdapat 1 parameter yang bernama "side". Fungsi ini akan menghitung luas persegi dengan perhitungan "side" kali "side" dan hasilnya akan dikembalikan, sehingga kita dapat menangkap hasilnya di variable "area".<br><br>

## Sorting Concept
<p align="justify">
Sebenarnya kita sangat familiar dengan konsep ini, seperti ketika dikelas kita pasti diabsen oleh guru dengan urutan nama dari "A" ke "Z", atau menyebutkan angka dari 1 sampai 10. Jika kalian mahasiswa, pasti kita sering mengurutkan data acak di sheet dengan menggunakan rumus sehingga datanya terurut dari terkecil ke terbesar. Contoh tersebut adalah bentuk konsep sorting atau pengurutan. Namun karena kita akan membuat sebuah program dan berkomunikasi dengan komputer, maka kita perlu memahami bagaimana sorting bekerja.<br><br>

## Swapping 
<p align="justify">
Pertama kita perlu paham konsep penukaran atau swapping. Konsep ini adalah bentuk menukar posisi urutan dari 2 posisi ke posisi yang saling berlawanan.<br>

<p align="justify">
Kita ambil contoh sebagai berikut, bagaimana caranya kita memindahkan isi dari Gelas A yang berisi kopi dengan gelas B yang berisi air jeruk. Kalau kita coba langsung mengisi ke salah satu gelas, maka isinya akan tercampur. Maka kita dapat menggunakan gelas penampung untuk memindahkan salah satu gelas agar kosong dan kemudian isinya dipindahkan. Maka kita berhasil memindahkan isi dari kedua gelas.<br><br>

## Sorting
<p align="justify">
Alur sorting sebenarnya sangat sederhana, kita cukup membandingkan antara deretan data yang ingin diurutkan dari awal ke akhir kemudian menukar posisi data yang tidak sesuai.<br>

<p align="justify">
Contoh ada urutan angka <code>3, 2, 1, 4</code> dengan rincian, angka 3 di urutan 1, angka 2 di urutan 2, angka 1 di urutan 3, dan angka 4 di urutan terakhir.<br>

<p align="justify">
Alur pertama yang dilakukan adalah membandingkan angka pertama dan ke-2. Karena 3 lebih besar dari 2, maka kita akan menukar posisi dari 3 dan 2. Hasil akhirnya akan menjadi:<br>

<p align="justify">
Kita lakukan lagi dengan membandingkan angka ke-2 dan 3. Karena 3 lebih besar dari 1, maka kita akan menukar posisi dari 3 dan 1. Hasil akhirnya akan menjadi:<br>

<p align="justify">
Proses terakhir adalah membandingkan angka ke-3 dan terakhir. Karena 3 tidak lebih besar dari 4, maka tidak ada penukaran posisi. Hasil akhir:<br>

<p align="justify">
Karena data masih belum urut, kita akan mengulangi proses pengecekan dari urutan awal lagi dan menukar posisi angka yang kurang tepat hingga hasilnya urut.<br>

<p align="justify">
Proses sorting sederhana ini disebut dengan istilah bubble sort, kita bisa melihat visualisasinya di sini. Pseudocode dapat dibuat sebagai berikut:<br>

```
PROCEDURE Sorting (data)
    STORE "swapped" WITH TRUE
    WHILE "swapped"
        DO
        STORE "swapped" WITH FALSE
        FOR "i" from 1 in "data" // variabel 'i' sebagai posisi dari setiap data dengan dimulai pada po
            IF "data"["1"]> "data" ["i+1"]
                STORE "temp" WITH "data"["1"]
                STORE "data"["i"] WITH "data"["i+1"]
                STORE "data" ["i+1"] WITH "temp"
                STORE "swapped" WITH TRUE
        END FOR
    PRINT data
END PROCEDURE
```

<br><br>
