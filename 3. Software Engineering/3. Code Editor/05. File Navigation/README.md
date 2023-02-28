# File/Folder Navigation
## Introduction
<p align="justify">
Sanagt bagus untuk kita dapat melakukan navigasi antar file saat melakukan eksplorasi pada sebuah proyek yang sedang dikerjakan. Terkadang kita perlu untuk berpindah antar satu file/folder ke file/folder lainnya dengan cepat di antara kumpulan file/folder. Selain berpindah dengan menggunakan <code>explorer</code>, kita bisa juga <code>navigate</code> dengan shortcut sehingga kita bisa navigasi file/folder hanya menggunakan keyboard.</p><br>

## File Navigation
<p align="justify">
VS Code menyediakan dua perintah untuk kita dapat melakukan navigasi di dalam dan di seluruh file dengan key bindings yang mudah digunakan. Tahan <code>Ctrl</code> dan tekan <code>Tab</code> untuk melihat daftar semua file yang terbuka di grup editor. Untuk membuka salah satu file ini, gunakan <code>Tab</code> lagi untuk memilih file yang ingin kita navigasikan, lalu lepaskan <code>Ctrl</code> untuk membukanya.<br>

<p align="center">
<img height="150rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/quicknav.png"><br>

<p align="justify">
Kita dapat menggunakan <code>Ctrl + Alt + -</code> dan <code>Ctrl + Shift + -</code> untuk bernavigasi di antara file dan mengedit lokasi. Jika kita jump di antara baris yang berbeda dari file yang sama, shortcuts ini memungkinkan kita untuk menavigasi di antara lokasi tersebut dengan mudah.<br><br>

## Breadcrumbs
<p align="justify">
Editor memiliki navigation bar yang disebut <strong>Breadcrumbs</strong>. Ini menunjukkan lokasi saat ini dan memungkinkan kita dapat dengan cepat melakukan navigasi antara folder, file, dan simbol.<br>

<p align="center">
<img height="150rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/breadcrumbs.png"><br>

<p align="justify">
Memilih <strong>Breadcrumbs</strong> pada path akan menampilkan dropdown sehingga kita dapat dengan cepat menavigasi ke folder dan file lain.<br>

<p align="center">
<img height="150rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/breadcrumb-folder-dropdown.png"><br>

<p align="justify">
Kita dapat mematikan <strong>Breadcrumbs</strong> melalui <strong>View > Show Breadcrumbs</strong> atau dengan pengaturan <code>breadcrumbs.enable</code>.<br><br>

## Breadcrumbs Keyboard Navigation
<p align="justify">
Untuk berinteraksi dengan <strong>Breadcrumbs</strong>, gunakan perintah focus breadcrumbs atau tekan <code>Ctrl + Shift + .</code>.<br>

<p align="justify">
Kita juga dapat berinteraksi <strong>Breadcrumbs</strong> dropdown. Tekan <code>Ctrl + Shift + ;</code> untuk memfokuskan elemen terakhir, gunakan <code>left</code> dan<code>right</code> untuk menavigasi, dan gunakan <code>spasi</code> untuk membuka elemen di editor.<br><br>

## Go To Definition
<p align="justify">
JIka kita menekan <code>Ctrl</code> dan mengarahkan kursor ke simbol, preview deklarasi akan muncul:<br>

<p align="center">
<img height="100rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ctrlhover.png"><br>

> <strong>Tips: Kita juga dapat melompat dengan <code>Ctrl + Click</code> atau membuka definisi ke samping dengan <code>Ctrl + Alt + Click</code>.</strong>

<br>

## Go To Symbol
<p align="justify">
Kita dapat menavigasi simbol di dalam file dengan <code>Ctrl + Shift + O</code>. Dengan mengetik <code>:</code> simbol akan dikelompokkkan bedasarkan kategori. Tekan <code>Up</code> dan <code>Down</code> lalu arahkan ke tempat yang kita inginkan.<br>

<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/gotosymbol.png"><br><br>

## Open Symbol By Name
<p align="justify">
Beberapa bahasa mendukung jumping ke simbol di seluruh file dengan <code>Ctrl + T</code>. Ketik huruf pertama dari jenis yang ingin kita navigasikan, terlepas dari file mana yang memuatnya, dan tekan <code>Enter</code>.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/symbol.png"><br>
