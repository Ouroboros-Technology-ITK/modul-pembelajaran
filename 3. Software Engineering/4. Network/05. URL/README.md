# URL 
## What Is URL?
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
URL adalah singkatan dari Uniform Resource Locator. URL merupakan sebuah alamat yang digunakan untuk mengidentifikasi sebuah resource di internet. Secara teori, URL akan mengidentifikasi sebuah resource atau sumber daya di internet, namun dalam praktiknya, URL juga digunakan untuk mengidentifikasi sebuah resource di dalam sebuah jaringan komputer.<br>

<p align="justify">
Sumber daya tersebut biasanya berupa sebuah halaman web, sebuah service atau bisa juga sebuah file seperti dokumen, gambar, dll.<br>

<p align="justify">
Berikut contoh URL:<br>

```
https://www.google.com/search?q=google+search
https://www.youtube.com/?hl=en&gl=ID
https://developer.mozilla.org/en-US/docs/Learn/
```

<p align="justify">
Secara anatomi, URL terdiri dari beberapa bagian. Terdapat bagian yang wajib dan bagian yang opsional.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

## Schema/Protokol
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Bagian pertama dari URL adalah protokol yang digunakan untuk mengakses resource / sumber daya tersebut. Protokol ini digunakan untuk menentukan cara mengakses sumber daya tersebut. Protokol yang sering digunakan adalah <code>http</code> dan <code>https</code>.<br>

<p align="justify">
Namun, terdapat juga protokol lain yang sering dipakai seperti <code>mailto</code> untuk mengirim email, <code>ftp</code> untuk mengakses file, dll.<br>

```
mailto:example@gmail.com
```

<br>

## Authority
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Authority adalah bagian yang mengidentifikasi resource / sumber daya tersebut. Bagian ini terdiri dari dua bagian, yaitu domain dan port.<br>

<p align="justify">
Pemisah antara Schema dan Authority adalah <code>://</code>.<br>

```
https://www.google.com
https://www.youtube.com 
```

<br>

## Domain
<p align="justify">
Domain adalah sebuah alamat yang digunakan untuk mengidentifikasi sebuah resource/sumber daya. Domain ini bisa berupa sebuah nama domain atau sebuah IP address.<br>

```
www.google.com
www.youtube.com
```

<p align="justify">
Terdapat beberapa jenis domain:<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

- Top Level Domain (TLD): TLD juga dikenal sebagai ekstensi domain. Dalam contoh URL di atas <code>.org</code> berperan sebagai TLD.
- Second Level Domains (SLD) : SLD merupakan nama domain yang kita daftarkan yaitu <code>example</code>.
- Third Level Domain (3LD): Third Level Domain merupakan bagian dari nama domain atau alamat situs web yang muncul sebelum SLD. Bagian ini juga dikenal sebagai subdomain. Pada contoh URL di atas, maka <code>photos</code> adalah Third Level Domain atau subdomain website kita.<br><br>

## Subdomain
<p align="justify">
Subdomain merupakan bagian dari domain utama yang berdiri sendiri. Bisa dibilang subdomain ini seperti "anak" dari domain utama. Ia memang bagian yang terpisah dari domain, tetapi takkan bisa ada tanpa domain utama itu sendiri.<br>

<p align="justify">
Umumnya subdomain digunakan untuk dua hal. Pertama, subdomain dipakai untuk melakukan staging dan testing sebuah website. Langkah ini akan memudahkan Anda dalam tracking dan menyimpan perubahan dalam website.<br>

<p align="justify">
Kedua, subdomain juga sering digunakan untuk memisahkan antara website yang berbeda. Misalnya, Anda memiliki sebuah website yang berisi konten yang berbeda dengan website lainnya. Anda bisa memisahkan konten tersebut dengan menggunakan subdomain.<br>

<p align="justify">
Subdomain ditulis di depan domain utama, dipisahkan dengan tanda titik <code>.</code> dan diikuti dengan nama subdomain.<br>

<p align="justify">
Contoh subdomain:<br>

```
https://support.google.com
https://maps.google.com
```

<p align="justify">
Contoh di atas adalah subdomain dari domain <code>google.com</code> yaitu <code>support</code> dan <code>maps</code>.<br><br>

## Port
<p align="justify">
Port sudah dijelaskan sebelumnya. Port ini biasanya dihilangkan dari URL karena umumnya menggunakan port HTTP (80) dan HTTPS (433), namun jika port tidak dihilangkan, maka port akan dituliskan setelah host dengan menggunakan tanda titik dua <code>:</code> sebagai pemisah.<br>

```
www.google.com:80
www.ruangguru.com:443
```

<br>

## Path
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Bagian ini juga merupakan identifikasi resource / sumber daya. Bedanya dengan Authority, bagian ini mengidentifikasi resource / sumber daya secara spesifik. Bagian ini juga sering disebut sebagai path.<br>

<p align="justify">
Path mengacu pada lokasi yang tepat dari halaman, posting, file, atau aset lainnya. Path berada setelah nama host dan dipisahkan oleh garis miring <code>/</code>. Path juga terdiri dari ekstensi file asset, seperti gambar (.jpg atau .png, dll.), dokumen (.pdf atau .docx), dan banyak lagi.<br>

<p align="justify">
Contoh path:<br>

```
https://www.example.com/path/to/file.html 
https://www.ruangguru.com/beasiswa/bkk
```

<br>

## Parameters
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Bagian ini digunakan untuk mengirimkan data ke resource/sumber daya. Data yang di kirimkan biasanya berupa parameter berisi key dan value yang digunakan untuk melakukan query ke resource/sumber daya tersebut. Biasa diawali dengan tanda tanya <code>?</code> dan setiap parameter dipisahkan dengan tanda ampersand <code>&</code>.<br>

<p align="justify">
Parameter juga dapat diatur secara dinamis di sebuah path sebagai value yang dipisahkan oleh garis miring <code>/</code> dan karakter lain (tergantung pada sistem yang digunakan dan bagaimana penerapannya). Parameter biasanya digunakan untuk tracking dan analytics serta encoding specific information untuk digunakan dalam situs web dan aplikasi.<br>

```
https://www.google.com/search?q=google+search&hl=id
```

<p align="justify">
Contoh URL di atas berisi query <code>q</code> dengan nilai <code>google search</code> dan query <code>h1</code> dengan nilai <code>id</code>. <br>

<p align="justify">
Contoh lain URL dengan parameter:<br>

```
https://www.example.com/solutions?user=123&color=blue 
https://www.example.com/solutions/user/123/color/blue
```

<br>

## Anchor
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Anchor digunakan di situs web untuk menerapkan "bookmark" dan elemen navigasi halaman internal. Ini dapat digunakan untuk menyediakan links ke lokasi tertentu dalam halaman. Anchor akan ditempatkan segera setelah file/path URL (bila ada).<br><br>

## How To Access URL
<p align="justify">
Kita dapat mengakses URL dengan memasukkan URL tersebut ke dalam browser. Kita cukup mengetik URL tersebut di dalam address bar browser dan tekan tombol Enter.<br>

<p align="justify">
Kita juga dapat mengakses URL dengan menggunakan command line menggunakan perintah <code>curl</code> dan diikuti alamat URL.<br>

```
> curl https://www.google.com/search?q=google+search&hl=id
```

<br>

## Localhost
```
http://localhost
```

<p align="justify">
Localhost adalah sebuah domain spesial yang digunakan untuk mengakses resource / sumber daya yang berada di lokal komputer kita. Localhost juga dapat kita akses menggunakan alamat <code>IP 127.0.0.1</code>.<br>

<p align="justify">
Keterbatasan localhost adalah URL ini memerlukan sebuah server yang sudah kita persiapkan sebelumnya. Selain itu domain localhost hanya bisa diakses dari komputer kita sendiri, tidak bisa diakses oleh komputer lain meskipun terkoneksi ke jaringan yang sama.<br><br>

## What Are The Function Of Localhost?
<p align="justify">
Bagi seorang developer, localhost berperan sangat penting dalam pembuatan dan pengembangan aplikasi berbasis internet maupun sebuah website.<br>

<p align="justify">
Di mana localhost digunakan untuk mengakses website maupun aplikasi secara lokal atau offline. Sehingga developer bisa melakukan pengujian maupun melihat langsung tampilan dari website atau aplikasi sebelum dipublikasikan di internet.<br>

<p align="justify">
Setelah aplikasi berjalan dengan baik kita dapat mempublikasikan aplikasi / website tersebut ke internet menggunakan sebuah hosting/server dan domain yang kita miliki.<br><br>
