# File Management

## <code>ls</code>
<p align="justify">
Command <code>ls</code> adalah command yang digunakan untuk menampilkan konten daftar file atau folder (direktori) yang ada dalam direktori aktif saat ini. Misalnya, untuk menampilkan daftar semua konten di direktori home, kita akan menjalankan command <code>ls</code> dari direktori home.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ls.png"><br><br> 

## <code>ls -a</code>
<p align="justify">
Ketika menampilkan daftar file dan kita ingin menyertakan hidden file (file tersembunyi) di dalam daftar tersebut, maka bisa kita tambahkan opsi command <code>-a</code>> Opsi ini akan menampilkan hidden files yang dimulai dengan tanda titik (.) seperti yang ditunjukkan di bawah ini.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ls-a.png"><br><br>

## <code>ls -la</code>
<p align="justify">
Selain menampilkan hidden file dengan opsi <code>-a</code>, kita bisa menggunakan opsi <code>-l</code> untuk menampilkan daftar file dalam bentuk list. Opsi <code>-l</code> adalah singkatan dari listing dan mencetak informasi tambahan seperti file permissions, user, group, file size, dan date of creation.<br>

<p align="justify">
Opsi ini bisa digabung dengan opsi <code>-a</code>, sehingga bisa kita ketik dengan <code>la</code>. Jika dijalankan command <code>ls -la</code> maka akan tampil sebuah list detail beserta hidden files.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ls-la.png"><br><br>

## <code>cd</code>
<p align="justify">
Command <code>cd</code> merupakan singkatand dari change directory, merupakan suatu command yang digunakan untuk berpindah atau menavigasi direktori aktif sekarang ke direktori lainnya. Misalnya, untuk menavigasi dari direktori <code>c/Documents</code>, dijalankan command <code>cd Documents</code>.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/cd.png"><br><br>

## <code>touch</code>
<p align="justify">
Command <code>touch</code> digunakan untuk membuat file pada sistem Unix/Linux. Syntax command-nya seperti berikut: <code>touch filename</code>.<br>

<p align="justify">
Misalnya, untuk membuat file <code>file1.txt</code>, jalankan command <code>touch file1.txt</code>. Untuk mengecek apakah file berhasil dibuat, kita bisa jalankan command <code>ls</code> dan lihat apakah file yang baru kita buat ada di dalam list file.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/touch.png"><br><br>

## <code>clear</code>
<p align="justify">
Mari kita asumsikan terminal kita pernuh dengan perintah/command dan output seperti yang ditunjukkan di bawah ini.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/clear.png"><br>

<p align="justify">
Jalankan command <code>clear</code>, maka layar terminal akan dibersihkan seperti yang ditunjukkan di bawah ini.<br><br>

## <code>mv</code>
<p align="justify">
Command <code>mv</code> adalah command yang cukup sebaguna. Bergantung pada cara penggunaannya, dimana command ini dapat mengganti nama file atau memindahkannya dari satu lokasi ke lokasi lainnya. Untuk memindahkan file, gunakan command <code>mv</code> diikuti dengan nama file yang akan dipindahkan kemudian lokasi tempat file dipindahkan <code>mv filename /path/to/destination</code>. Contoh, kita akan memindahkan <code>file1.txt</code> ke folder <code>/Videos</code>, jalankan command <code>mv file1.txt ../Videos</code>. Untuk mengkonfirmasi pemindahan file, jalankan command <code>ls ../Videos</code>.<br><br>

<p align="justify">
Untuk mengganti nama file, gunakan command <code>mv filename1 filename2</code>. Command ini menghapus nama file asli dan menetapkan argumen kedua sebagai nama file baru.<br><br>

> <strong>Note: Perhatikan penggunaan <code>../</code> di depan sebuah folder, ini berarti kita akan mundur satu direktori.</strong>
<br><br>

## <code>cp</code>
<p align="justify">
Command <code>cp</code>, kependekan cari copy, menyalin file dari satu lokasi ke lokasi lainnya. Berbeda dengan command <code>mv</code>, command <code>cp</code> mempertahankan file asli di lokasinya saat ini dan membuat salinan/duplikat di direktori yang berbeda. Misalnya, untuk menyalin file <code>file2.txt</code> dari direktori saat ini ke direktori <code>/Documents</code>, kita bisa jalankan perintah <code>cp file2.txt ../Documents</code>.<br><br>

## <code>rm</code>
<p align="justify">
Command <code>rm</code>, kependekan dari remove, digunakan untuk menghapus file. Syntaxnya cukup mudah <code>rm filename</code>. Misalnya, untuk menghapus <code>file2.txt</code>, jalankan perintah <code>rm file2.txt</code>.<br>

<p align="justify">
Kita juga dapat menggunakan command <code>rm -r</code> untuk menghapus sebuah file, yang membedakan dengan command <code>rm</code> adalah jika argumen yang diberikan kepada command <code>rm -r</code> adalah sebuah direktori, maka sistem akan menghapus direktori tersebut beserta semua file dalam direktori dan juga subdirektori jika ada.<br><br>

## <code>rmdir</code>
<p align="justify">
Command <code>rmdir</code> untuk menghapus sebuah direktori. MIsalnya, untuk menghapus direktori <code>folder1</code>, kita bisa jalankan command <code>rmdir folder1</code>.<br><br>

## <code>open</code>
<p align="justify">
Untuk membuka file apapun dari command line dengan aplikasi default, cukup dengan mengetik <code>open</code> diikuti dengan nama file/path. Misalnya, kita ingin membuka <code>file1.txt</code>, maka jalankan command <code>open file1.txt</code>.<br><br>

<br><br>
<div align="justify">
    <!-- Prev Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering/1.%20Introduction/2.%20Day%20To%20Day" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/left%20(1).png" align="left" height="50" width="50"></a>
    <!-- Next Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/home%20(2).png" align="right" height="50" width="50"></a>
<div>
