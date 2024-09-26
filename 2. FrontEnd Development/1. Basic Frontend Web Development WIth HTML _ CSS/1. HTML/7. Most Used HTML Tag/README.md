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

### Image

Untuk menampilkan gambar, gunakan tag   `<img>` dan letakkan sumber file gambar di `src` *attribute*. Jika sumber gambar ada di URL lain, tulis *full pathnya* sebagai berikut:

```html
<img src="https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_92x30dp.png">
```
<br>
<center>

![google logo](https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_92x30dp.png)

</center>

Jika sumber gambar ada di URL yang sama dengan URL HTML, kita bisa menggunakan *relative path*:

```html
<img src="./google.png" alt="google icon">
```
Jika gambar tidak dapat ditampilkan, text yang ada di alt akan ditampilkan. alt attribute juga membantu tuna netra (accessibility feature) karena bisa dibaca oleh *screen reader*.

<p align="center">
    (tampilkan image saat tidak bisa dibaca)
    <h5 align="center">Source: https://</h5>
</p>

Untuk mengatur ukuran gambar, gunakan attribute width dan height

```html
<img src="./google.png" width="100" height="100">
```
Jika hanya salah satu attribute yang diisi, maka browser akan mengatur ukuran secara otomatis (*aspect ratio* tetap sama).



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


<hr>

### Link

####`<a>` 

Elemen HTML `<a>` (*elemen anchor*) digunakan untuk membuat hyperlink ke halaman web, file, alamat, email, atau lokasi di halaman yang sama. Elemen ini memiliki atribut yang sangat penting yaitu `href`.

```html
<a href="https://example.com">Website</a>
```

<br>
##### href attribute

Atribut `href` digunakan untuk menentukan URL yang akan dituju dari link tersebut. Jika elemen *anchor* tidak memiliki `href` di dalamnya maka elemen tersebut tidak akan menjadi sebuah *hyperlink*. Link yang dijadikan value dari href tidak terbatas pada URL berbasis HTTP, namun dapat berupa skema URL apa pun yang didukung oleh browser seperti di bawah ini:

- Email Adress:
 ```html
 <a href="mailto:someone@gmail.com">Send email</a>
 ```
- Nomor Hp :
```html
<a href="tel:+6282176589876">Call me</a>
```

<br>

##### target attribute

Atribut target digunakan untuk menentukan tempat untuk membuka dokumen yang dijadikan link.

```html
<a href="https://google.com" target="_blank"> New Tab</a>
```

Di bawah ini terdapat beberapa kata kunci yang memiliki arti khusus untuk tempat memuat URL:

- `_self`: Membuka link dalam frame yang sama dengan yang diklik (default)
- `_blank`: Membuka link di window atau tab baru
- `_parent`: Membuka link di parent frame


<hr>

### Table

Pembuatan table pada html menggunakan elemen `<table>`

```html
<table>
  <thead>
    <tr>
      <th>Nama</th>
      <th>Umur</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Andi</td>
      <td>25</td>
    </tr>
    <tr>
      <td>Budi</td>
      <td>30</td>
    </tr>
  </tbody>
  <tfoot>
    <tr>
      <td colspan="2">Jumlah Data: 2</td>
    </tr>
  </tfoot>
</table>
```

hasil yang ditampilkan dari kode HTML diatas akan menjadi seperti berikut:

<p align="center">
    (tampilkan hasil code diatas)
    <h5 align="center">Source: https://</h5>
</p>

####`<thead>`
Tag ini digunakan untuk mengelompokkan bagian header dari tabel. Biasanya digunakan untuk menampilkan judul atau keterangan pada bagian atas tabel. Tag `<thead>` harus ditempatkan sebelum tag `<tbody>` dan `<tfoot>`, jika ada.

```html
<thead>
  <tr>
    <th>Nama</th>
    <th>Umur</th>
    <th>Jenis Kelamin</th>
  </tr>
</thead>
```

####`<tbody>`
digunakan untuk mengelompokkan baris isi tabel. Biasanya terdiri dari satu atau beberapa baris yang berisi data atau informasi yang ingin ditampilkan di tabel.

```html
<tbody>
  <tr>
    <td>Andi</td>
    <td>25</td>
    <td>Laki-laki</td>
  </tr>
  <tr>
    <td>Budi</td>
    <td>30</td>
    <td>Laki-laki</td>
  </tr>
  <tr>
    <td>Citra</td>
    <td>28</td>
    <td>Perempuan</td>
  </tr>
</tbody>
```

####`<tfoot>`
digunakan untuk mengelompokkan baris footer tabel. Biasanya terdiri dari satu atau beberapa baris yang mengandung informasi tambahan, seperti jumlah data yang ada di tabel atau informasi statistik lainnya.

```html
<tfoot>
  <tr>
    <td colspan="3">Jumlah Data: 3</td>
  </tr>
</tfoot>
```

Pada `thead`, `tbody` dan `tfoot` setidaknya harus memiliki satu *tag* `<tr>` didalamnya.

#### `<tr>` 
mendefinisikan baris dalam tabel HTML. Elemen `<tr>` berisi satu atau lebih elemen `<th>` atau `<td>`

<p align="center">
    (tampilkan table dengan tanda yang mana itu tr)
    <h5 align="center">Source: https://</h5>
</p>

#### `<th>` 

Tag `<th>` mendefinisikan sel header dalam tabel HTML. Sel header berisi informasi header (dibuat dengan elemen `<th>`). Teks dalam elemen `<th>` dicetak tebal dan di tengah secara default.

<p align="center">
    (tampilkan table dengan tanda yang mana itu th)
    <h5 align="center">Source: https://</h5>
</p>

#### `<td>`
Tag `<td>` mendefinisikan sel data standar dalam tabel HTML. Teks dalam elemen `<td>` teratur dan rata kiri secara default.

<p align="center">
    (tampilkan table dengan tanda yang mana itu td)
    <h5 align="center">Source: https://</h5>
</p>

<hr>

### Form

Sebuah halaman web tidak hanya digunakan untuk menampilkan sebuah informasi, namun juga dapat digunakan untuk mengambil & mengirim data dari penggunanya. Salah satu cara untuk mendapatkan informasi tersebut adalah dengan menggunakan **form**.

#### `<form>`

Form di HTML dapat kita buat dengan tag `<form>` . Dalam sebuah form biasanya terdapat kotak input (control) dan elemen lainnya yang dapat diedit kemudian ditulis untuk di kirim ke sebuah server. Elemen `<form>` dapat berisi satu atau lebih elemen berikut:

- `<input>`
- `<textarea>`
- `<button>`
- `<select>`
- `<option>`
- `<optgroup>`
- `<fieldset>`
- `<label>`
- `<output>`

```html
<form action="./action.php">
  <label for="nama">Nama:</label>
  <input type="text" id="nama" name="nama"><br><br>

  <label for="email">Email:</label>
  <input type="email" id="email" name="email"><br><br>

  <label for="telepon">Telepon:</label>
  <input type="tel" id="telepon" name="telepon"><br><br>

  <label for="pesan">Pesan:</label>
  <textarea id="pesan" name="pesan"></textarea><br><br>

  <input type="submit" value="Kirim">
</form>
```

Hasil yang ditampilkan dari kode di atas akan seperti berikut: 

<p align="center">
    (tampilkan form diatas)
    <h5 align="center">Source: https://</h5>
</p>

##### `<label>`

Elemen HTML `<label>` mewakili sebuah judul atas isian tertentu. Untuk mengasosiasikan `<label>` dengan control tertentu bisa menggunakan 2 cara, yaitu:

1. Menggunakan atribut `for`, *value* dari atribut tersebut adalah nama `id` dari *control* yang ingin dituju.

```html
<label for="nama">Nama:</label>
<input type="text" id="nama" name="nama">
```

2. Menempatkan *control* didalam elemen.

```html
<label>Nama: <input type="text"></label>
```

##### `<input>`

Elemen HTML `<input>` adalah elemen form control yang paling banyak digunakan. Elemen ini digunakan untuk menunjukkan sebuah inputan (masukkan) dalam bentuk kotak dan sejenisnya yang dapat diedit/diketik untuk diisi data tertentu.

Elemen `<input>` adalah elemen yang tidak memiliki tag penutup (closing tag) dan merupakan elemen kosong yang tidak memiliki konten, hanya terdapat atribut saja. Elemen `<input>` dapat ditampilkan dalam banyak cara, bergantung pada atribut type. Beberapa di antaranya yaitu:

- Text 
```html
<label for="nama">Nama:</label>
<input type="text" id="nama" name="nama">

<!-- Teks input menjadi titik hitam karena bertipe password -->

<label for="password">Password:</label>
<input type="password" id="password" name="password">

```
<p align="center">
    (tampilkan form diatas)
    <h5 align="center">Source: https://</h5>
</p>
- Radio button
```html
<label>Jenis Kelamin:</label><br>

<input type="radio" id="pria" name="jk" value="pria">
<label for="pria">Pria</label><br>

<input type="radio" id="wanita" name="jk" value="wanita">
<label for="wanita">Wanita</label><br>
```
<p align="center">
    (tampilkan form diatas)
    <h5 align="center">Source: https://</h5>
</p>

- Checkbox
 ```html
 <label for="hobi">Hobi:</label><br>

<input type="checkbox" id="baca" name="hobi" value="baca">
<label for="baca">Membaca</label><br>

<input type="checkbox" id="olahraga" name="hobi" value="olahraga">
<label for="olahraga">Olahraga</label><br>

<input type="checkbox" id="traveling" name="hobi" value="traveling">
<label for="traveling">Traveling</label><br>
 ```
<p align="center">
    (tampilkan form diatas)
    <h5 align="center">Source: https://</h5>
</p>

- File
 ```html
<label for="foto">Foto:</label>
<input type="file" id="foto" name="foto">
 ```

 <p align="center">
    (tampilkan form diatas)
    <h5 align="center">Source: https://</h5>
</p>
<br>

##### `<textarea>`
Elemen HTML `<textarea>` merupakan elemen yang merepresentasikan input control yang memiliki teks lebih dari satu baris. Teks yang ditulis di dalam elemen `<textarea>` merupakan teks dasar dengan Karakter tanpa batas. Teks tersebut dapat di sisipkan secara langsung di antara open tag (`<textarea>`) dan closing tag (`</textarea>`) sebagai awal konten. Untuk mengubah ukuran elemen ini dapat
ditentukan oleh atribut `cols` dan `rows`.

```html
<form>
  <label for="nama">Nama:</label>
  <input type="text" id="nama" name="nama"><br><br>

  <label for="pesan">Pesan:</label>
  <textarea id="pesan" name="pesan">Tulis pesanmu sebanyak banyaknya disini.</textarea><br><br>

  <input type="submit" value="Kirim">
</form>
```

<p align="center">
    (tampilkan form diatas)
    <h5 align="center">Source: https://</h5>
</p>
<br>

##### `<select>`

Elemen HTML `<select>` merepresentasikan sebuah control yang digunakan untuk memilih deretan opsi yang diberikan.

dalam elemen `<select>` kita perlu menambahkan beberapa elemen `<option>`. `<option>` memuat pilihan yang akan diseleksi oleh pengguna.

```html
<form>
  <label for="kota">Kota:</label>
  
  <select id="kota" name="kota">
    <option value="">--Pilih Kota--</option>
    <option value="jakarta">Jakarta</option>
    <option value="bandung">Bandung</option>
    <option value="surabaya">Surabaya</option>
    <option value="yogyakarta">Yogyakarta</option>
  </select>
  
  <input type="submit" value="Submit">
</form>

```

<p align="center">
    (tampilkan form diatas)
    <h5 align="center">Source: https://</h5>
</p>
