# Introduction
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Database.jpeg"> <h5 align="center">Source: https://www.gramedia.com/best-seller/apa-itu-database/</h5><br>

<p align="justify">
Pada pembahasan materi terminal akan dibutuhkan sistem operasi berbasis Linux. Tapi jangan khawatir, untuk pengguna Windows kita bisa menginstall WSL (<em>Windows Subsystem for Linux</em>) agar kita bisa menggunakan terminal berbasis Linux meskipun menggunakan sistem operasi Windows.<br>

## What Is WSL ?
<p align="justify">
WSL adalah fitur opsional Windows yang memungkinkan program Linux beejalan secara native di Windows. Hal yang dapat kita lakukan dengan menggunakna WSL adalah sebagai berikut: </p>
<ul style="list-style-type:circle;" style="text-align:justify">
  <li>Memilih distribusi GNU/Linux dari Microsoft Store.</li>
  <li>Menjalankan command line tools yang umum seperti <code>ls -la</code>, <code>touch</code>, <code>grap</code> atau perintah Linux lainnya.</li>
  <li>Menjalankan aplikasi terminal base GNU/Linux termasuk tools text editor: <code>vim</code> dan <code>nano</code>.</li>
  <li>Menginstall package manager tambahan menggunakan distribusi GNU/Linux.</li>
</ul>

<p align="center">
    <a href="https://youtu.be/48k317kOxqg" target="_blank"><img src="https://img.youtube.com/vi/48k317kOxqg/0.jpg"></a> 
    <h5 align="center">Source: What Can I Do With WSL? | One Dev Question</h5>
<p><br>

## What Is WSL 2?
<p align="justify">
WSL 2 adalah versi terbaru dari WSL yang bertujuan untuk menambahkan kompabilitas pemanggilan sistem. Perbandingannya dengan WSL 1 dapat kita lihat pada halaman resminya di <a href="https://learn.microsoft.com/en-us/windows/wsl/" target="_blank"><strong>learn.microsoft.com.</strong></a>

<p align="center">
    <a href="https://www.youtube.com/watch?v=MrZolfGm8Zk" target="_blank"><img src="https://img.youtube.com/vi/MrZolfGm8Zk/0.jpg"></a> 
    <h5 align="center">Source: WSL2: Code Faster On The Windows Subsystem For Linux! | Tabs vs Space</h5>
<p><br>

## Prerequisites
<p align="justify">
Untuk update ke WSL 2, kita harus menggunakan <strong>Windows 10</strong>. </p>
<ul style="list-style-type:circle;" style="text-align:justify">
  <li>Untuk sistem x64: Versi 1903 atau yang lebih baru, dengan Build 18362 atau yang lebih bagus.</li>
  <li>Untuk sistem ARM64: Versi 2004 atau yang lebih baru, dengan Build 19041 atau yang lebih baru.</li>
</ul>
<p align="justify">
atau menggunakan <strong>Windows 11</strong>.</p>

> <strong>info: Build yang lebih rendah dari Build 18362 tidak mendukung WSL 2.</strong>

<p align="justify">
Untuk memeriksa versi dan nomor Build kita, tekan tombol Windows + R, lalu ketik <code>winver</code>, pilih OK.

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Database.jpeg"> <h5 align="center">Source: https://www.gramedia.com/best-seller/apa-itu-database/</h5><br>

## Step By Step Installing WSL
<p align="justify">
UNtuk menginstall WSL secara simple, sebenarnya bisa menggunakan command <code>wsl --install</code>. Namun disini kita akan mencoba menginstall secara manual agar kita dapat mengetahui apa saja yang harus dilakukan untuk menjalankan WSL.<br>

- <strong>Langkah 1</strong> - Aktifkan Windows Subsystem for Linux
<p align="justify">
Buka PowerShell sebagai Administrator (tekan Windows > ketik "PowerShell" > klik kanan > pilih Run As Administrator) dan masukkan perintah ini. 

```
dism.exe /online /enable-featur /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
```<br>
Untuk menjadi seorang programmer, maka paling tidak harus menguasai dalam 1 buah bahasa pemrograman. Programmer sudah mulai mempelajari dan memahami secara dalam bagaimana cara kerja dari suatu syntax di dalam bahasa pemrograman.


<div align="justify">
    <!-- Prev Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering/1.%20Introduction/2.%20Day%20To%20Day" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/left%20(1).png" align="left" height="50" width="50"></a>
    <!-- Next Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/home%20(2).png" align="right" height="50" width="50"></a>
<div>
