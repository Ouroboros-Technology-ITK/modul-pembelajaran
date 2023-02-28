# DNS
## What Is DNS?
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Domain Name Server atau DNS adalah sebuah sistem yang menghubungkan Uniform Resource Locator (URL) dengan Internet Protocol Address (IP Address). Normalnya, untuk mengakses internet, kita perlu mengetik IP Address sebuah website. Cara ini cukup merepotkan, sebab kita perlu punya daftar lengkap IP Address website yang ingin dikunjungi dan memasukkannya secara manual. DNS akan memudahkan pekerjaan kita dengan cara cukup memasukkan nama domain / URL kemudian DNS akan menerjemahkannya ke dalam IP Address yang tepat.<br><br>

## Function
<p align="justify">
Dari penjelasan di atas, sebenarnya sudah jelas apa fungsi utama dari DNS. Supaya lebih jelas, berikut beberapa fungsi dari DNS:<br>

- Meminta informasi IP address sebuah website berdasarkan nama domain atau URL.
- Mencari server yang tepat untuk mengirimkan email.
- Melakukan pencarian di data cache untuk mempercepat identifikasi alamat dari domain yang dikunjungi sebelumnya.
<br><br>

## How Does It Work?
<p align="justify">
Prinsip dasar cara kerja DNS adalah dengan cara mencocokkan nama komponen URL dengan komponen IP address. Setiap URL dan IP Address memiliki bagian-bagian yang saling menjelaskan satu sama lain. Proses pencocokan ini disebut dengan query.<br>

<p align="justify">
Jika sulit dibayangkan, kita anggap seperti kegiatan mencari buku di perpustakaan. Ketika kita mencari buku di perpustakaan, biasanya kita akan diberi kode yang menjelaskan letak buku tersebut. Kode buku perpustakaan tersebut dinamai Dewey Decimal System (DDS). Biasanya terdiri atas kode topik buku, kode nama belakang penulis, dan kode tahun buku diterbitkan. Nah, konsep tersebut sama juga diterapkan dalam DNS. Untuk memahaminya lebih dalam, kita perlu mengetahui bagian-bagian URL yang tersusun dalam hierarki DNS. Sama seperti kode buku perpustakaan, setiap bagiannya menjelaskan bagian domain.<br>

<p align="justify">
Satu perbedaan yang paling mencolok adalah kode perpustakaan mulai dari depan, sedangkan DNS dimulai dari belakang. Maka dari itu, kita akan runut bagian-bagian DNS ini dari belakang. Berikut penjelasan lengkapnya:<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

- <strong>Root-Level Domain</strong> 
<p align="justify">
Merupakan bagian tertinggi dari hirarki DNS. Biasanya berwujud tanda titik <code>.</code> di bagian paling belakang sebuah URL.<br>

- <strong>Top-Level Domain</strong> 
<p align="justify">
Merupakan ekstensi yang berada di bagian depan root-level domain. Terdapat dua jenis TLD yang umumnya dipakai. Keduanya, yaitu Generic Top-Level Domain (GTLD) dan Country Code Top- Level Domain (CCLTD).<br>
  
  - <strong>GTLD</strong> menunjukkan sifat dari website tersebut, seperti <code>com</code> untuk komersial, <code>edu</code> untuk pendidikan dan edukasi, <code>gov</code> untuk pemerintahan, dan sebagainya.
  - <strong>CCLTD</strong> menunjukkan asal negara pemilik website, seperti <code>id</code> untuk website Indonesia, <code>uk</code> untuk Inggris, dan seterusnya.<br>
  
- <strong>Second-Level Domain</strong> 
<p align="justify">
Merupakan nama lain untuk domain itu sendiri. la sering digunakan sebagai identitas institusi atau branding. Dalam kasus URL <code>en.wikipedia.org</code>, yang dimaksud SLD adalah wikipedia.<br>

- <strong>Third-Level Domain atau subdomain</strong> 
<p align="justify">
Merupakan bagian dari domain utama yang berdiri sendiri. Apabila domain diibaratkan sebagai rumah, subdomain adalah salah satu ruang khusus di rumah itu sendiri.<br>

<p align="justify">
Setelah memahami cara kerja DNS di dalamnya, maka yang terjadi ketika kita menggunakan DNS adalah:<br>

- DNS server meminta informasi domain atau nama situs web yang akan dikunjungi.
- DNS server kemudian mencocokkan nama tersebut dengan angka berupa alamat IP.
- DNS lokal kemudian akan mencari alamat IP tersebut di cache lokal yang tersimpan dalam komputer. 
- Jika situs web pernah dibuka sebelumnya, maka data IP address dari DNS cache akan digunakan dan meneruskan request ke IP address tersebut.
- Jika alamat IP tidak ditemukan di dalam cache lokal komputer, maka DNS akan meminta data dari DNS rekursif dan jika alamat IP ditemukan maka komputer akan meneruskan request ke IP address tersebut. 
- Namun, jika alamat tidak ada di server DNS rekursif, maka DNS akan mencarinya dalam server DNS lainnya (protocol ini disebut DNS otoritatif).
- Setelah alamat IP ditemukan, situs web akan ditampilkan dan cache akan disimpan ke dalam DNS lokal.

