# Most Used HTML Tag

<p align="justify">
Terdapat banyak tag HTML yang sering digunakan dalam pembuatan halaman web. 

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
    <h5 align="center">Source: https://</h5>'
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
    <h5 align="center">Source: https://</h5>'
</p>

Setelah ditekan :

<p align="center">
    (insert all page after pop up showed)
    <h5 align="center">Source: https://</h5>'
</p>

