# Most Used HTML Tag

<p align="justify">
Terdapat banyak tag HTML yang sering digunakan dalam pembuatan halaman web. 

<hr>


### Container

#### `<header>`

Elemen HTML ini berbeda dengan `<head>`. Penggunaan nya untuk mengelompokkan elemen yang berkaitan dengan bagian atas halaman dari web. seperti judul,logo, navigasi yang terdiri dari deretan link atau tautan social media dan elemen lainnya.
</p>

```html
<header>
    <h1>Halaman Saya</h1>
    <p>Ini adalah deskripsi header.</p>
    <nav>
        <a href="https://www.facebook.com/contoh" target="_blank">Facebook</a>
        <a href="https://www.instagram.com/contoh" target="_blank">Instagram</a>
        <a href="https://www.twitter.com/contoh" target="_blank">Twitter</a>
    </nav>
</header>
```

#### `<article>`

Tag `<article>` adalah salah satu tag struktural HTML yang penggunaanya untuk mengelompokkan konten suatu artikel, berita atau postingan yang dapat berdiri sendiri untuk memisahkan dari artikel lainnya pada halaman web.

```html
<article>
<header>
<h2>Judul Artikel</h2>
    <p>Oleh: John Doe</p>
    <p>Tanggal Publikasi: 10 Januari 2022</p>
</header>
    <p>Ini adalah paragraf pertama pada artikel.</p>
    <p>Ini adalah paragraf kedua pada artikel.</p>
    <p>Ini adalah paragraf ketiga pada artikel.</p>
</article>
```

#### `<div>`

<p align="justify">
Digunakan untuk mengelompokkan elemen-elemen HTML menjadi sebuah bagian atau divisi pada halaman. Penggunaan nya sering digunakan dalam membuat tata letak halaman dengan menggunakan CSS. Tetapi jika menggunakan HTML saja `<div>` tidak berpengaruh pada konten atau tata letak.

```html
<div class="container">
    <h1>Ini adalah judul</h1>
    <img src="gambar.jpg" alt="Ini adalah gambar">
    <ul>
        <li>Item daftar 1</li>
        <li>Item daftar 2</li>
        <li>Item daftar 3</li>
    </ul>
</div>
```

Contoh diatas, tag `<div>` digunakan untuk menampung beberapa elemen HTML. Kita memberikan `class="container"` di tag `<div>`, yang kemudian container bisa diberikan styling dengan menggunakan CSS.

#### `<Footer>`

Tag `<footer>` adalah salah satu tag HTML yang digunakan untuk menampilkan informasi berkaitan dengan akhir dari suatu halaman web. Informasi yang ada pada *footer* biasanya tentang hak cipta, tautan ke halaman terkait, dan informasi kontak.
</p>

```html
<footer>
    <p>&copy; 2022 Contoh Perusahaan. All rights reserved.</p>
    <nav>
        <ul>
            <li><a href="#">Tentang Kami</a></li>
            <li><a href="#">Kebijakan Privasi</a></li>
            <li><a href="#">Hubungi Kami</a></li>
        </ul>
    </nav>
</footer>
```

Contoh di atas, tag `<footer>` berfungsi menampilkan informasi yang berikaitan dengan akhir dari sebuah halaman web, seperti hak cipta dan tautan ke halaman terkait.

`<footer>` biasanya sering digunakan untuk memberikan informasi tambahan yang relevan dengan halaman web, agar pengunjung dapat mengetahui lebih lanjut tentang halaman web dan organisasi di baliknya.

<hr>

### Text

#### `<h1>` hingga `<h6>`

Tag `<h1>` adalah tag yang digunakan untuk menampilkan teks judul, dimana `<h1>` adalah judul teratas sedangkan `<h6>` adalah judul terbawah. Semakin besar nomor pada tag heading, maka akan semakin rendah tingkat judul nya.

```html
<h1>Ini adalah heading level 1</h1>
<h2>Ini adalah heading level 2</h2>
<h3>Ini adalah heading level 3</h3>
<h4>Ini adalah heading level 4</h4>
<h5>Ini adalah heading level 5</h5>
<h6>Ini adalah heading level 6</h6>
```

Contoh diatas menggunakan `h1` hingga `h6` untuk menampilkan teks judul. Tag `h1` menjadi judul halaman utama, sementara `h2` hingga `h6` menjadi sub-bagian halaman.

<p align="center">
    (tampilkan hasil dari kode diatas)
    <h5 align="center">Source: https://</h5>
</p>

Penting juga untuk memerhatikan tingkatan dari *tag header*. Penggunaan dari *Heading* yang tepat dapat menaikkan *ranking* **SEO**.

```html
<!-- Mulai dengan tag <h1> -->
<h2>Ini adalah heading level 2</h2>

<!-- 
    Gunakan tingkatan 1 level 
    kebawah dari tingkat sebelumnya
    seharusnya <h3>
-->
<h5>Ini adalah heading level 5</h5>
```

#### `<p>`

*Tag* `p` digunakan untuk menampilkan paragraf pada halaman web.

```html
<p>Ini adalah contoh teks paragraf.</p>
<p>Paragraf baru.</p>
```
Baris pemisah akan secara *default* muncul untuk menjadi pemisah di antara paragraf.

<p align="center">
    (tampilkan hasil dari kode diatas)
    <h5 align="center">Source: https://</h5>
</p>

Kita dapat menambahkan teks sebanyak yang kita inginkan dalam satu *tag* `p`. Sehingga teks nya akan ditampilkan dalam satu blok paragraf.

#### `<span>`

*Tag* `span` digunakan menandai sebagian kecil dari sebuah teks atau elemen dalam dokumen yang ingin diberikan atribut tertentu. Letak *tag* `span` dituliskan secara *inline* sehingga tidak merubah struktur barisnya dan akan tetap ditampilkan dalam baris yang sama.

```html
<p>Saya suka makan nasi <span style="background-color: yellow;">kuning</span> dengan piring <span style="color:red">merah</span></p>
```

<p align="center">
    (tampilkan hasil dari kode diatas)
    <h5 align="center">Source: https://</h5>
</p>


Perlu diketahui bahwa tag `span` tidak memiliki arti atau pengaruh struktural pada dokumen seperti tag heading atau tag paragraf.

<hr>

### List

*Tag list* didalam HTML merupakan *Tag* yang sering digunakan dalam pembuatan daftar pada halaman web.

#### `<ul>` : Unordered List

*Tag* `ul` digunakan untuk membuat list yang tidak memerlukan urutan tertentu. *Item* dari penggunaan *tag* `ul` akan ditandai dengan tanda titik,lingkaran atau kotak.

```html
<ul>
  <li>Apel</li>
  <li>Jeruk</li>
  <li>Pisang</li>
</ul>
```

<p align="center">
    (tampilkan hasil dari kode diatas)
    <h5 align="center">Source: https://</h5>
</p>

#### `<ol>` : Ordered List

*Tag* `ol` digunakan untuk membuat daftar dengan urutuan tertentu atau terstruktur. Penanda dari tiap *item* akan berbentuk angka, huruf atau angka romawi.

```html
<ol>
  <li>Apple</li>
  <li>Orange</li>
  <li>Banana</li>
</ol>
```

<p align="center">
    (tampilkan hasil dari kode diatas)
    <h5 align="center">Source: https://</h5>
</p>

#### Nested List

HTML juga mendukung *nested list* atau daftar bertingkat. Penggunaan nya untuk membuat list dalam list atau daftar yang memiliki tingkatan yang berbeda.

```html
<ol>
  <li> Menu Makanan
    <ul>
        <li>Ayam Goreng</li>
        <li>Ikan Bakar</li>
        <li>Sate Kambing</li>
    </ul>
  </li>
  <li>Menu Minuman
    <ul>
        <li>Es Teh</li>
        <li>Jus Jeruk</li>
    </ul>
  </li>
</ol>
```

Dalam penulisan *nested list* `ul` untuk "menu makanan" dan "menu minuman" menjadi konten dari `li` atau *item list* dari `ol` yang menjadi parent paling atas.

<p align="center">
    (tampilkan hasil dari kode diatas)
    <h5 align="center">Source: https://</h5>
</p>

<hr>

### Button

<p align="justify">

*Button* biasanya digunakan untuk membuat sebuah tombol di dalam halaman web agar lebih interaktif seperti menerima *user input*. Penggunaan `<form>`.

Konten *button* ditempatkan di dalam tag HTML nya.


```html
<button>Click Here!</button>
```

<p align="center">
    (insert the button above after rendedered)
    <h5 align="center">Source: https://</h5>'
</p>

Tidak hanya teks yang bisa dimasukkan menjadi konten *Button* tetapi juga seperti tag `<img>`.

```html
<button>
    <button>
        <img 
            src="gambar.png" 
            alt="Gambar tombol">
    </button>
</button>
```

<p align="center">
    (insert the button above after rendedered)
    <h5 align="center">Source: https://</h5>
</p>

##### Button Interactivity

*Button* yang tidak bergantung dengan *form* tidak bisa melakukan interaksi banyak. Biar dapat bisa interaktif, perlu menambahkan `onClick` sebagai attribute dari *button* dengan javascript yang akan dipelajari di modul Javascript sendiri.

```html
<button 
    onclick="alert('Tombol diklik!')">Klik di sini
</button>
```

Contoh diatas `onClick` digunakan untuk mengeluarkan pop-up dengan teks "Tombol diklik!" dengan fungsi javascript `alert()`. 

<p align="center">
    (insert the button above after rendedered)
    <h5 align="center">Source: https://</h5>
</p>

Setelah ditekan :

<p align="center">
    (insert all page after pop up showed)
    <h5 align="center">Source: https://</h5>
</p>

