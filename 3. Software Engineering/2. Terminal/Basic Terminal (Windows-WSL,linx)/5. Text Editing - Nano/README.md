# Text Editing - Nano
<p align="justify">
Saat bekerja menggunakan command line, aktivitas membuat atau mengedit file teks merupakan sesuatu yang akan sering dilakukan. Bagi kamu yang membutuhkan text editor terminal yang sederhana dan mudah untuk digunajna, nano adalah solusinya..<br><br>

## What Is Nano Text Editor ?
<p align="justify">
GNU nano adalah teks editor terminal yang mudah digunakan untuk OS Unix dan Linux. Nano mencakup semua fungsi dasar yang diharapkan dari teks editor biasa seperti penyorotan sintaks, multiple buffer, search dan replace dengan dukungan reguler ecpression, periksa ejaan, UTH-8 encoding, dll.<br><br>

## How To User Nano
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
  <li>Tekan tombol <code>ctrl + o</code> untuk menyimpan perubahan yang dibuat di file dan melanjutkkan proses editing.</li>
  <li>Tekan tombol <code>ctrl + x</code> untuk keluar dari editor.</li>
</ul><br>

## Cara Mencari dan Mengganti Teks
<p align="justify">
Command <code>w</code> menunjukkan siapa yang logged on dan apa yang mereka lakukan termasuk waktu aktif.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/w.png"><br><br>

## <code>whoami</code>
<p align="justify">
Command <code>whoami</code> digunakan untuk menampilkan username pengguna saat ini.<br><br>

## <code>uname -a</code>
<p align="justify">
Untuk mengetahui seluruh informasi terkait sistem yang kita miliki, kita dapat menjalankan command <code>uname -a</code> yang akan menampilkan informasi terkait system name, hostname, kernel-version, kernel release, dan machine hardware name kita.<br><br>

## <code>cat /proc/cpuinfo</code>
<p align="justify">
<code>cat /proc/cpuinfo</code> menampilkan jenis prosesor yang dijalankan sistem kita, termasuk jumlah CPU yang ada.<br><br>

## <code>cat /proc/meminfo</code>
<p align="justify">
<code>cat /proc/meminfo</code> digunakan untuk menampilkan jumlah memori yang kosong dan yang digunakan (baik fisik maupun swap) pada sistem serta shared memory dan buffer yang digunakan oleh kernel.<br><br>



<br><br>
<div align="justify">
    <!-- Prev Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering/1.%20Introduction/2.%20Day%20To%20Day" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/left%20(1).png" align="left" height="50" width="50"></a>
    <!-- Next Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/home%20(2).png" align="right" height="50" width="50"></a>
<div>
