# Terminal
## <code>ping</code>
<p align="justify">
Package Internal Grapher atau PING, adalah command yang digunakan untuk memeriksa status konektivitas antar jaringan. Dengan command ping, kita dapat menghitung durasi waktu yang dibutuhkan ketika client mengirim pesan ke host dan client menerima respon dari host dalam sebuah jaringan. PING bekerja dengan cara client mengirim pesan Internet Control Message Protocol (ICMP) ke target host dan client akan menerima response Echo dari target host kemudian dihitung selisih waktunya.<br><br>


## Install <code>ping</code>
<p align="justify">
Kita dapat mengecek ketersediaan <code>ping</code> dengan command:<br>

```
ping -h
```

<p align="justify">
atau<br>

```
ping -v
```

<p align="justify">
Jika tidak ada ping yang terinstal, bisa menginstal dengan menggunakan:<br>

```
apt-get update && apt-get install -y iputils-ping
```

<br>

## Function
<p align="justify">
Seperti yang dijelaskan sebelumnya, ping digunakan untuk memeriksa status konektivitas antar jaringan. Kita bisa mencoba langsung dengan mengecek konektivitas ke <code>google.com</code>.<br>

```
ping www.google.com
```

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Jika konektivitas jaringan kita normal, maka akan muncul pesan berulang seperti gambar di atas.<br><br>

## <code>whois</code> domain
<p align="justify">
Whois adalah layanan internet yang memberikan informasi tentang sebuah domain. Informasi ini sering disebut <code>whois record</code> atau <code>data whois</code>. Informasi itu meliputi nama lengkap pemilik, alamat, nomor telepon, email, dan sebagainya.<br>

<p align="justify">
Beberapa informasi yang diberikan adalah:<br>
<ol style="list-style-type:circle;" style="text-align:justify">
  <li>Informasi kontak, berupa nama pendaftar /admin/teknisi, nama organisasi, alamat rumah, nomor telepon, email, dan fax.</li>
  <li>Informasi registrasi, berisi nama Whois servernya, URL, registrar, IANA ID, Abuse contact phone, dan Abuse contact email. Kontak Billing, yang berisi tentang tagihan dan pembayaran yang dilakukan selama pembelian domain.</li>
  <li>Kontak Billing, yang berisi tentang tagihan dan pembayaran yang dilakukan selama pembelian domain.</li>
  <li>Status, berupa informasi apakah domain sudah diperbarui (updated), prohibited, atau expired.</li>
  <li>Tanggal penting, berisi informasi seputar tanggal update, tanggal domain dibuat, tanggal domain didaftarkan</li>
  <li>Nama server, ini hanya berisi URL server domain</li>
</ol>

<p align="justify">
Kita bisa mencobanya dengan command <code>whois</code> ke domain <code>www.ruangguru.com</code>.<br><br>

## <code>dig</code> domain
<p align="justify">
<strong>Domain Information Groper</strong> atau dig adalah command yang digunakan untuk mengambil informasi tentang domain name server atau DNS.<br>

## Install <code>dig</code>
<p align="justify">
Pastikan dulu apakah dig sudah terinstal dengan menggunakan command:<br>

```
dig -v
```

<p align="justify">
Outputnya akan berupa informasi versi dari dig yang terinstall.<br>

<p align="justify">
Jika tidak ada ping yang terinstal, bisa menginstal dengan menggunakan:<br>

```
sudo apt install dnsutils
```

<br>

## Lookup
<p align="justify">
Penggunaan yang paling dasar adalah melihat informasi DNS dari domain yang kita inginkan. Kita coba langsung dengan mengecek ke domain <code>ruangguru.com</code>.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Hasilnya kita akan mendapatkan daftar informasi seperti gambar di atas. Kita cukup fokus ke bagian <code>ANSWER SECTION</code>, karena bagian itu adalah hasil dari query lookup yang dilakukan oleh <code>dig</code>.<br>

<p align="justify">
Output yang diperoleh dapat diseleksi dengan hanya memunculkan bagian <code>ANSWER SECTION</code> saja dengan menambah option <code>+noall</code> dan <code>+answer</code>.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Kita dapat menyederhanakan hasil output dengan menambah opsi <code>+short</code> di belakangnya.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Informasi yang didapat hanya berupa IP server dari domain yang kita masukkan.<br><br>

## Specific Quary Name Server
<p align="justify">
Secara default jika tidak ada name server yang ditentukan, command dig akan menggunakan server yang terdaftar di internal file <code>/etc/resolv.conf</code>.<br>

<p align="justify">
Untuk menentukan name server yang akan dijalankan query, dapat menggunakan simbol <code>@</code> (at) dan diikuti dengan alamat IP name server atau hostname.<br>

<p align="justify">
Misal kita mencari informasi DNS dari domain <code>ruangguru.com</code> dengan menggunakan name server google (8.8.8.8).<br>

<p align="justify">
Sekilas tidak ada yang berbeda, tetapi di bagian <code>SERVER</code> sudah menggunakan name server google (8.8.8.8).<br><br>

## dig -x host
<p align="justify">
Terdapat option <code>-x</code> yang digunakan untuk mengambil informasi menggunakan host name. Kita dapat langsung mencobanya dengan t name berikut <code>172.217.14</code> (google.com).<br>
  
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br><br>

## <code>wget</code> file
<p align="justify">
Wget adalah command yang digunakan untuk mengunduh file dari internet. Nama wget adalah gabungan kata dari WWW (World Wide Web) dan get sehingga tool ini disupport menggunakan protokol FTP, SFTP, HTTP dan HTTPS.<br><br>

## Install <code>wget</code>
<p align="justify">
Kita dapat memastikan <code>wget</code> sudah terinstal atau tidak dengan command:<br>

```
wget --help
```

<p align="justify">
Perintah atau command tersebut untuk melihat seluruh informasi command yang tersedia sekaligus memastikan apakah <code>wget</code> sudah terinstall atau belum. Jika belum terinstal dapat menggunakan command:<br>

```
apt-get install wget
```

<br>

## Basic Usage
<p align="justify">
Kita bisa langsung mencoba untuk mendownload file dengan command <code>wget</code> kemudian diikuti dengan URL file yang ingin kita unduh/download.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Contoh diatas adalah hasil download file gambar dari URL <code>http://www.google.com/images/srpr/logo3w.png</code>. File tersebut akan disimpan ke direktori aktif di Terminal.<br>

<p align="justify">
Agar dapat mengatur posisi file ke direktori yang kita inginkan, kita dapat menggunakan opsi <code>-P</code> dan diikuti dengan nama direktori.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Hasilnya, file akan disimpan ke direktori <code>Downloads</code>.<br>

<p align="justify">
Atau dapat juga mengubah nama file yang dihasilkan dengan opsi <code>-0</code> dan diikuti dengan nama file yang kita inginkan.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Nama yang awalnya adalah <code>logo3w.png</code> akan berubah menjadi <code>new-logo.png</code> ketika sudah berhasil di download ke direktori kita.<br><br>

## <code>wget -c</code> file
<p align="justify">
Seringkali kita menghadapi kondisi koneksi internet tidak stabil atau tiba-tiba listrik padam. Masalah ini akan menyebabkan proses download file menjadi terhenti. Kita dapat menggunakan command <code>wget -c</code> untuk melanjutkan file yang sudah didownload sebelumnya.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Jika kita tidak menggunakan option <code>-C</code>, maka file akan di download ulang. Akibatnya akan terdapat file baru dengan ditambah nama <code>.1</code> di belakangnya karena sebelumnya file dengan nama yang sama sudah ada.<br><br>

## <code>curl</code>
<p align="justify">
CURL atau singkatan dari Client URL adalah tool yang digunakan untuk menangani berbagai hal yang berhubungan dengan URL (uniform resource locator). CURL mendukung banyak protokol yang umumnya adalah HTTP, FTP, SSH, dan lain-lain. Kita dapat mengecek versi dari curl menggunakan command <code>curl –version</code>.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Selain informasi versi, kita juga dapat mengetahui daftar protokol yang didukung oleh curl yang terpasang.<br><br>

## Syntax Curl Commadn Dasar
<p align="justify">
Sintaks dasar dari curl adalah sebagai berikut:<br>

```
curl [OPTION] [URL]
```

<p align="justify">
Kegunaan <code>curl</code>yang paling sederhana adalah untuk menampilkan konten dari sebuah halaman web. Kita bisa coba melihat konten dari <code>https://www.google.com</code>.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Karena kita mengakses konten dari URL, maka outputnya akan berupa text saja atau lebih tepatnya HTML. Kita akan belajar HTML di materi terpisah.<br><br>

## Opsi File Curl Command
<p align="justify">
Curl dapat mengunduh file dari remote location menggunakan opsi <code>-O</code> (O besar).<br>

```
curl -0 http://www.google.com/image/srpr/logo3w.png
```

<p align="justify">
Hasil download akan disimpan ke direktori aktif dari terminal dengan nama file yang sama. Kita dapat mengubah nama file yang ingin didownload dengan opsi <code>-o</code> (o kecil).<br>

```
curl -o new-logo.png http://www.google.com/image/srpr/logo3w.png
```

<p align="justify">
Terdapat opsi juga yang bisa digunakan untuk melanjutkan unduhan yang terganggu menggunakan opsi <code>-C -</code> .<br>

```
curl -C-0 http://www.google.com/image/srpr/logo3w.png
```

<p align="justify">
Terakhir, jika ingin mengunduh beberapa file langsung dapat dilakukan satu command <code>curl</code>, seperti contoh berikut:<br>

```
curl -0 http://www.google.com/image/srpr/logo3w.png 0 http://www.google.com/image/srpr/logo-new.png
```

<br>

## Curl Command Untuk HTTP
<p align="justify">
CURL juga dapat digunakan ketika ada sebuah proxy server. Kita dapat menggunakan opsi <code>-x</code> untuk menentukan proxy server.<br>

```
curl -x proxy.example.com:8080 -U username:password -0 http://www.google.com/image/srpr/logo3w.png
```

<p align="justify">
Contoh di atas adalah bentuk akses ke proxy. Kita dapat menggunakan opsi <code>-U</code> untuk menentukan username dan password untuk proxy server.<br><br>

## Header   
<p align="justify">
Permintaan ke HTTP akan selalu berisi header. Header HTTP mengirim informasi tambahan tentang server web jarak jauh bersama dengan permintaannya. Sementara melalui browser development tool, kita dapat memeriksa informasi header.<br>

<p align="justify">
Kita dapat mengambil informasi header dari sebuah web menggunakan opsi <code>-I</code>.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

## Method GET and POST
<p align="justify">
Secara default <code>curl</code> menggunakan method GET HTTP untuk mengakses sebuah URL. Materi tentang method HTTP akan dijelaskan secara terpisah, tapi command <code>curl</code> dapat menggunakan method POST untuk mengirim data ke server.<br>

```
curl -data "text-Hello" https://myDomain.com/firstPage.jsp
```

<p align="justify">
atau<br>

```
curl -X POST -data "text-Hello" https://myDomain.com/firstPage.jsp
```

<p align="justify">
Perintah atau command di atas berisi permintaan POST yang diikuti oleh permintaan GET ke URL <code>https://myDomain.com/firstPage.jsp</code>.<br>

<p align="justify">
Kita juga dapat menentukan beberapa metode HTTP dalam satu curl command. Lakukan ini dengan menggunakan opsi <code>--next</code>, seperti ini:<br>

```
curl -data "text-Hello" https://myDomain.com/firstPage.jsp --next https://myDomain.com/displayResult.jsp
```

<p align="justify">
Contoh di atas adalah mengirimkan data ke URL <code>https://myDomain.com/firstPage.jsp</code> dan kemudian mengambil data dari URL <code>https://myDomain.com/displayResult.jsp</code>.<br>
