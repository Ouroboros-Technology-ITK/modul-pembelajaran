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
Buka PowerShell sebagai Administrator (tekan Windows > ketik "PowerShell" > klik kanan > pilih Run As Administrator) dan masukkan perintah ini: 

```
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
```

- <strong>Langkah 2</strong> - Aktifkan Virtual Machine Feature
<p align="justify">
Buka PowerShell sebagai Administrator dan jalankan: 

```
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```

Restart komputer untuk menyelesaikan penginstalan dan pembaruan ke WSL 2.


- <strong>Langkah 3</strong> - Download Linux kernel update package
  - Download latest package: <a href="https://learn.microsoft.com/en-us/windows/wsl/install-manual#step-4---download-the-linux-kernel-update-package" target="_blank"><strong>WSl 2 Linux Kernel Udpate Package</strong></a>
  - Run update package yang telah di download. (Double click untuk run, maka kita akan dimintai permissions, pilih 'Yes' untuk menyetujui penginstalan)

- <strong>Langkah 4</strong> - Set WSL 2 sebagai default version
<p align="justify">
Buka PowerShell dan jalankan perintah ini: 

```
wsl --set-default-version 2
```

Restart komputer untuk menyelesaikan penginstalan dan pembaruan ke WSL 2.


- <strong>Langkah 5</strong> - Install Linux distribution
  - Buka Microsoft Store dan search Linux distributin.
  - Kita akan membutuhkan <a href="https://www.microsoft.com/store/productId/9PN20MSR04DW" target="_blank"><strong>Ubuntu 22.04 LTS</strong></a>
  - Dari halaman distribution, klik tombol 'Get'.
    <p align="justify">
    Pertama kali menjalankan Linux distribution yang baru diinstall, kita akan diminta untuk menunggu satu atau dua menit agar fil mendekompres dan disimpan di PC kita. Kemudian kita perlu membuat username dan password untuk Linus distribution.<br>
  
##

<div align="justify">
    <!-- Prev Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering/1.%20Introduction/2.%20Day%20To%20Day" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/left%20(1).png" align="left" height="50" width="50"></a>
    <!-- Next Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/home%20(2).png" align="right" height="50" width="50"></a>
<div>
