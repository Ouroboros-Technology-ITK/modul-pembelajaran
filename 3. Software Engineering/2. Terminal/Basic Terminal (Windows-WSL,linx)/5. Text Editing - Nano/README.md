# Text Editing - Nano
<p align="justify">
Saat bekerja menggunakan command line, aktivitas membuat atau mengedit file teks merupakan sesuatu yang akan sering dilakukan. Bagi kamu yang membutuhkan text editor terminal yang sederhana dan mudah untuk digunajna, nano adalah solusinya..<br><br>

## What Is Nano Text Editor ?
<p align="justify">
GNU nano adalah teks editor terminal yang mudah digunakan untuk OS Unix dan Linux. Nano mencakup semua fungsi dasar yang diharapkan dari teks editor biasa seperti penyorotan sintaks, multiple buffer, search dan replace dengan dukungan reguler ecpression, periksa ejaan, UTH-8 encoding, dll.<br><br>

## How To Use Nano
<p align="justify">
Untuk mengecek ketersediaan dan versi Nano di sistem operasi yang dimiliki, jalankan command atau perintah di bawah ini:

```
nano --version
```

Biasanya akan muncul output seperti ini:

```
GNU nano, version 2.9.3
(C) 1999-2011, 2013-2018 Free Software Foundation, Inc.
(C) 2014-2018 the contributors to nano
Email: nano@nano-editor.org Web: https://nano-editor.org/
```

Jika versi Nano tidak muncul, maka kita bisa mengikuti langkah-langkah install berikut ini:</p>
- Install Nano di Ubuntu dan Debian

```
sudo apt-get install nano
```

- Install Nano di CentOS dan Fedora

```
sudo yum install nano
```

<p align="justify">
Untuk membuka file yang ada atau membuat file baru, ketik <code>nano</code> di terminal Linux diikuti dengan nama file:

```
nano filename
```

Misalnya kita ingin membuka file dengan nama <code>nano demo.txt</code>. Maka terminal commandnya akan terliaht seperti ini:

```
nano demo.txt
```

Kita bisa membuka berbagai jenis file, seperti <code>.txt</code>, <code>.php</code>, <code>.html</code>, dll.</p>

<p align="justify">
Untuk membuka file <code>nano demo.txt</code> yang tersimpan di folder <code>/home/folder</code>. Kita bisa menjalankan:

```
nano /home/folder/demo.txt
```

Atau dapat juga dengan masuk melalui direktori tersebut dahulu:

```
cd /home/folder
nano demo.txt
```

<p align="justify">
Jika file tidak ditemukan, maka Nano akan membuatkan file tersebut. Jika command <code>nano</code> dijalankan tanpa menambahkan nama file yang spesifik (hanya mengetik <code>nano</code>) maka Nano akan menampilkan editor file yang kosong dan tidak bernama. Ketika akan keluar dari editor kita akan diminta untuk memberi nama pada file tersebut.<br>

<p align="justify">
Berikut adalah tampilan interface dari Nano text editor setelah command diatas dijalankan. Gunakan tombol bertanda anak panah di keyboard untuk menggerakkan cursor pada Nano.

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/nano.png"><br>

<p align="justify">
Di bagian bawah jendela terdapat shortcut yang dapat digunakan pada Nano text editor. Tanda "<code>^</code>" (caret) menunjukkan untuk menekan tombol <code>ctrl</code> (Windows) atau <code>control</code> (macOS). Berikut beberapa contohnya:</p>
<ul style="list-style-type:circle;">
  <li>Tekan tombol <code>ctrl + O</code> untuk menyimpan perubahan yang dibuat di file dan melanjutkkan proses editing.</li>
  <li>Tekan tombol <code>ctrl + X</code> untuk keluar dari editor.</li>
</ul><br>

## Cara Mencari dan Mengganti Teks
<p align="justify">
Tekan tombol <code>ctrl + W</code> untuk mencari teks. Masukkan value dan tekan enter. Jika ingin mencari teks yang sama, tekan tombol <code>alt + W</code>.<br>

<p align="justify">
Untuk mencari dan mengganti teks, tekan tombol <code>ctrl + W</code>. Setelah itu, tekan tombol <code>ctrl + R</code> untuk menambahkan teks yang akan dicari dan teks yang akan menggantikannya. Lalu kita akan diarahkan ke instance pertama dari teks yang dituju. Tekan tombol <code>Y</code> untuk mengganti satu teks atau <code>A</code> untuk mengganti semua instance.<br>

<p align="justify">
Jika ingin kembali ke keadaan awal setelah mengetikkan shortcut, gunakan <code>ctrl + c</code> untuk membatalkan proses saat ini.<br><br>

## Cara Edit Teks
<p align="justify">
Berikut ini adalah shortcut yang sering digunakan pada saat mengetik teks di Nano:</p>
<ul style="list-style-type:circle;">
  <li>Untuk memilih teks, arahkan kursor ke depan teks yang diinginkan dan tekan tombol <code>alt + A</code>. Shortcut ini akan menandai teks yang dikehendaki. Gerakkan kursor di sekitar teks dengan menggunakan tombol bertanda anak panah.</li>
  <li>Tekan tombol <code>alt + 6</code> untuk menyalin dan meletakkan teks yang dipilih ke clipboard.</li>
  <li>Tekan tombol <code>ctrl + K</code> untuk memotong teks yang dimaksud.</li>
  <li>Tekan tombol <code>ctrl + U</code> untuk meletekkan (paste) teks, navigasi ke baris yang dituju.</li>
</ul><br>

## Shortcut Nano
<p align="justify">
Berikut adalah tabel yang berisikan command yang sering digunakna saat menggunakan Nano teks editor:

| COMMAND  | PENJELASAN                                                                     |
|----------|--------------------------------------------------------------------------------|
| ctrl + A | Pindah ke awal baris.                                                          |
| ctrl + E | Pindah ke akhir baris.                                                         |
| ctrl + Y | Scroll ke bawah halaman.                                                       |
| ctrl + v | Scroll ke atas halaman.                                                        |
| ctrl + G | Jendela Bantuan akan muncul dan menampilkan semua command yang bisa digunakan. |
| ctrl + J | Merapikan paragraf yang ada.                                                   |
| ctrl + C | Menampilkan posisi kursor di teks.                                             |
| ctrl + X | Keluar dari Teks Editor                                                        |
| ctrl + \ | Mengganti string atau regular expression.                                      |
| ctrl + T | Mengaktifkan tool pemeriksa ejaan, jika tersedia.                              |
| ctrl + _ | Pindah ke baris dan nomor kolom yang spesifik.                                 |

<br><br>

## Why Use Nano
<p align="justify">
Jika kita berbicara soal penyuntingan atau pengeditan teks via command line, Nano adalah salah satu tool yang dapat diandalkan. Karena kemudahan penggunaannya, Nano sangat populer sebagi basic teks editor terutama unutk para pengguna Linux.<br>

<p align="justify">
MEksipun terkesan simple, Nano memiliki beberapa keunggulan diantaranya yaitu memperbolehkan kita untuk menyalin, meletakkan, memilih, dan mencari teks. Juga ada bar di bagian bawah editor yang menampilkan berbagai shortcut fungsional dari Nano. Nano dapat digunakan baik oleh user pemula maupun yang sudah berpengalaman.<br>

<p align="justify">
Dokumentasi selengkapnya mengenai Nano Teks Editor dapat dilihat di <a href="https://www.nano-editor.org/docs.php" target="_blank"><strong>nano-editor.org</strong></a>.<br><br>

<br><br>
<div align="justify">
    <!-- Prev Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering/1.%20Introduction/2.%20Day%20To%20Day" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/left%20(1).png" align="left" height="50" width="50"></a>
    <!-- Next Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/home%20(2).png" align="right" height="50" width="50"></a>
<div>
