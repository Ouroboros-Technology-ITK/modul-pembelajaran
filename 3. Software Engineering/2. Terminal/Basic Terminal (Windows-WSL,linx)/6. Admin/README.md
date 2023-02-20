# Admin

## What Is Administration Role?
<p align="justify">
Admin berdasarkan <em>Cambridge Dictionary</em> adalah kependekkan dari administrasi atau sebuh kegiatan yang terlibat dalam mengelola atau mengatur bisnis atau organisasi lain. Dalam Linux sendiri, Administrator Linux memiliki peran aktif dalam patching, compiling, dan troubleshooting Linux. Administrator Linux dapat melakukan pembaruan sistem dan melakukan konfigurasi sistem.<br><br>

## Sudo
<p align="justify">
Perintah <code>sudo</code> memungkinkan kita untuk menjalankan program sebagai pengguna lain, secara default penggunaan <code>root</code>. Jika kita sering dan terbiasa menggunakan command line, <code>sudo</code> adalah salah satu perintah yang akan sering kita gunakan.<br>

<p align="justify">
Menggunakan akun <code>sudo</code> lebih aman daripada <code>root</code>, karena kita dapat memberikan hak administratif terbatas kepada pengguna individu tanpa mereka tahu kata sandi <code>root</code>.<br>

<p align="justify">
Package <code>sudo</code> biasanya sudah diinstall pada sebagian besar Linux Distribution.<br>

<p align="justify">
Untuk memeriksa apakah package <code>sudo</code> diinstall pada sistem, buka terminal Linux, kemudian ketik <code>sudo</code>, dan tekan <code>Enter</code>. Jika <code>sudo</code> sudah terinstall di sistem, maka terminal akan menampilkan output berupa pesan bantuan singkat.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/what-admin.png"><br>

<p align="justify">
Jika belum terinstall, kita akan melihat output seperti <code>sudo command not found</code>. Kita dapat dengan mudah menginstall package <code>sudo</code> menggunakan distro package manager dengan menjalan perintah berikut:

```
apt install sudo
```
<br>

## Syntax
<p align="justify">
Syntax untuk perintah <code>sudo</code> adalah sebagai berikut:

```
sudo OPTION.. COMMAND
```

<p align="justify">
Perintah <code>sudo</code> memiliki banyak opsi yang mengontrol behavior-nya, tetapi biasanya <code>sudo</code> digunakan dalam bentuknya yang paling basic, tanpa opsi apapun.

```
sudo command
```

<p align="justify">
Dimana <code>command</code> adalah perintah yang ingin kita lakukan dengan hak <code>sudo</code>.<br>

<p align="justify">
<code>sudo</code> akan membaca file <code>/etc/sudoers</code> dan memeriksa apakah user yang meminta hak sudo ada di file sudoers. Pertama kali kita menggunakan <code>sudo</code>, kita akan diminta memasukkan kata sandi pengguna dan perintah itu akan dieksekusi sebagai <code>root</code>.<br>

<p align="justify">
Sebagi contoh, untuk mendaftar semua file di direktori <code>/root</code>, kita dapat menggunakan perintah berikut:

```
sudo ls /root

[sudo] password for linuxid:
.   ..  .bashrc     .cache      .config     .local      .profile
```
<br>

## Examples Of Sudo in Linux
<p align="justify">
Contoh basic penggunaan <code>sudo</code> adalah ketika kita ingin melakukan udpate package repository dengan menjalankan perintah dibwah ini:

```
apt-get update
```

<p align="justify">
Kita akan melihat error message dikarenakan kita tidak memiliki izin yang diperlukan untuk menjalankan perintah.

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/apt-get-updt.png"><br>

<p align="justify">
Sekarang kita akan mencoba kembali menjalankan perintah yang sama dengan <code>sudo</code>:

```
sudo apt-get udpate
```

<p align="justify">
Masukkan password saat diminta ijin dan sistem akan menjalankan perintah dan memperbarui package repository.

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/sude-update.png"><br><br>

<br><br>
<div align="justify">
    <!-- Prev Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering/1.%20Introduction/2.%20Day%20To%20Day" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/left%20(1).png" align="left" height="50" width="50"></a>
    <!-- Next Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/home%20(2).png" align="right" height="50" width="50"></a>
<div>
