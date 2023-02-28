# Port
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Port adalah mekanisme yang memungkinkan komputer terhubung dengan beberapa sesi koneksi dengan komputer dan program lainnya dalam jaringan. Port dapat mengidentifikasikan aplikasi dan layanan yang menggunakan koneksi di dalam jaringan TCP/IP. Sehingga, port juga mengidentifikasikan sebuah proses tertentu di mana sebuah server dapat memberikan sebuah layanan kepada klien atau bagaimana sebuah klien dapat mengakses sebuah layanan yang ada dalam server.<br><br>

## Logical dan Physical Port
<p align="justify">
Jika dilihat dari terminologinya, jenis port dibedakan menjadi dua, yaitu logical port (port non-fisik) dan physical port (port fisik).<br>

<p align="justify">
Logical port adalah jalur yang digunakan oleh aplikasi untuk menghubungkan dengan komputer lain melalui jaringan TCP/IP. Salah satu contohnya adalah mengkoneksikan komputer dengan internet. Port ini berperan penting dalam jaringan komputer.<br>

<p align="justify">
Physical port adalah soket, slot atau jack koneksi yang memungkinkan kabel bisa dihubungkan dengan komputer, router, modem, USB, dan perangkat lainnya. Masing-masing soket memiliki bentuk fisik dan fungsi yang berbeda pula. Beberapa port jaringan fisik terdapat pada perangkat keras jaringan komputer.<br><br>

## Port Numbering
<p align="justify">
Dilihat dari penomorannya, logical port terbagi menjadi tiga jenis. Ada jenis port yang terdaftar di <a href="https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml" target="_blank"><strong>Internet Assigned Numbers Authority (IANA)</strong></a>, dan ada yang tidak, berikut pembagiannya:<br>

- <strong>Well-known Port</strong>
<p align="justify">
Berkisar dari 0-1023. Ini merupakan port yang dikenali atau port sistem. Port ini selalu merepresentasikan layanan jaringan yang sama dan ditetapkan oleh IANA.<br>

- <strong>Registered Port</strong>
<p align="justify">
Berkisar dari 1024-49151. Port ini diketahui dan terdaftar di IANA tetapi tidak dialokasikan secara permanen, sehingga dapat menggunakan port number yang sama.<br>

- <strong>Dynamically Assigned Port</strong>
<p align="justify">
Berkisar dari 49152-65535. Port ini ditetapkan oleh sistem operasi atau aplikasi yang digunakan untuk melayani request dari pengguna sesuai dengan kebutuhan.<br>

<p align="justify">
Berikut ini beberapa contoh logical port yang sering digunakan beserta fungsinya:<br>

- <strong>Port 20 & 21 (FTP)</strong>
<p align="justify">
Port 20 dan 21 merupakan port untuk FTP atau File Transfer Protocol. FTP adalah protokol yang berguna dalam mentransfer data di dalam suatu jaringan.<br>

- <strong>Port 22 (SSH)</strong>
<p align="justify">
Port 22 adalah port standar untuk SSH (Secure Shell). Port ini berfungsi mengirimkan data melalui jaringan dalam bentuk terenkripsi. Dapat digunakan untuk menjalankan fungsi atau tugas yang bisa diakses dari jarak jauh, misalnya menghubungkan ke host atau server.<br>

- <strong>Port 23 (TELNET)</strong>
<p align="justify">
Port 23 TELNET adalah port untuk menghubungkan komputer dan server jarak jauh. Fungsinya mirip dengan SSH, hanya saja port 23 TELNET tidak menggunakan enkripsi pada koneksinya.<br>

- <strong>Port 25 (SMTP)</strong>
<p align="justify">
Port 25 berfungsi untuk SMTP (Simple Mail Transfer Protocol) yang berfungsi memastikan pengiriman email melalui jaringan dikomunikasikan dengan aman antara sesama SMTP server.<br>

- <strong>Port 53 (DNS)</strong>
<p align="justify">
Port 53 adalah jenis port untuk DNS yang berfungsi sebagai penerjemah nama domain menjadi alamat IP pada setiap host. Port ini mencocokkan nama domain yang dapat dibaca manusia dengan alamat IP yang dapat dibaca mesin. Jadi, kamu tak perlu mengetik alamat IP dalam bentuk angka ketika hendak mengunjungi website (lebih lengkapnya akan dibahas pada materi DNS).<br>

- <strong>Port 80 (HTTP/Web Server)</strong>
<p align="justify">
Port 80 berfungsi untuk HTTP, yakni memungkinkan browser terhubung ke halaman web. Port ini akan menerima permintaan koneksi dari klien, kemudian setelah koneksi berhasil dibuat, kamu akan mendapat akses ke berbagai halaman web di internet. HTTP/ web server juga memiliki port alternatif, yaitu port 8080 dan 80.<br>

- <strong>Port 443 (HTTPS)</strong>
<p align="justify">
Hampir mirip dengan port 80, port 443 berguna untuk menghubungkan klien ke internet, namun dengan fitur keamanan tambahan yang tidak dimiliki port HTTP 80. Port 443 mengenkripsi paket jaringan sebelum mentransfernya.<br>

- <strong>Port 1433 (MS SQL)</strong>
<p align="justify">
Port 1433 adalah port yang digunakan untuk menghubungkan komputer ke database Microsoft SQL Server. Lebih lengkapnya tentang Microsoft SQL dapat dilihat <a href="https://www.microsoft.com/en-us/sql-server/sql-server-downloads" target="_blank"><strong>disini</strong></a>.<br>

- <strong>Port 3306 (MySQL)</strong>
<p align="justify">
Port 3306 adalah port yang digunakan untuk menghubungkan komputer ke database MySQL. Lebih lengkapnya tentang MySQL dapat dilihat <a href="https://dev.mysql.com/doc/" target="_blank"><strong>disini</strong></a>.<br>

- <strong>Port 5432 (PostgreSQL)</strong>
<p align="justify">
Port 5432 adalah port yang digunakan untuk menghubungkan komputer ke database PostgreSQL. Lebih lengkapnya tentang PostgreSQL dapat dilihat <a href="https://www.postgresql.org/docs/" target="_blank"><strong>disini</strong></a>.<br>

- <strong>Port 8000* 8080</strong>
<p align="justify">
Port 8000 atau 8080 adalah port yang sering digunakan untuk web Cache port Server proxy web. Atau bisa juga digunakan untuk menjalankan aplikasi yang ada di dalam server / localhost. Lebih lengkapnya tentang proxy web dapat dilihat <a href="https://www.speedguide.net/port.php?port=8000" target="_blank"><strong>disini</strong></a>.<br>

<p align="justify">
Dan masih banyak lagi, kita bisa melihat seluruh list logical port di internet atau <a href="https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers" target="_blank"><strong>disini</strong></a>.<br>
