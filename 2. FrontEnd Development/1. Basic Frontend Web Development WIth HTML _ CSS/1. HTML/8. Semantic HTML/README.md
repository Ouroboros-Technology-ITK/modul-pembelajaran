# Semantic HTML

Semantic HTML adalah penggunaan tag HTML yang memiliki makna (*meaningful*) sesuai dengan fungsinya pada halaman web. Tujuan utama dari penggunaan semantic HTML adalah untuk meningkatkan aksesibilitas, SEO (Search Engine *Optimization*), dan pemeliharaan kode yang lebih mudah.

Pada awalnya HTML tidak memiliki semantic tag (elemen semantik) di dalamnya, jadi umumnya orang-orang membuat layout dengan menggunakan tag yang bukan sebagai kegunaannya. 

Seperti menggunakan tag `<table>` dan ada juga dengan tag `<div>`. Sebenarnya menggunakan kedua tag tersebut tidak sepenuhnya salah, karena membuat layout dengan kedua tag itu benar-benar bisa.

Perhatikan *output* dari HTML dibawah ini:

<p align="center">
    (tampilkan hasil code dibawah)
    <h5 align="center">Source: https://</h5>
</p>

```html
<h1>Hello World</h1>
<span
 style="font-weight: bold; font-size= 32px; margin: 21px 0;"> 
 Hello World
</span>
```

Kedua *output* di atas dibuat dengan *tag* `<h1>` dan `<span>`, dari *visual* terlihat sama, tapi coba perhatikan : 

- Kode yang pertama secara semantik benar (Tujuannya membuat heading atau judul)

- Kode yang kedua.menyisipkan style visual dalam sebuah tag HTML yang tidak memiliki makna (`<span>` dan `<div>` adalah tag yang tidak memiliki makna).

Jadi gunakanlah tag yang sesuai dengan peruntukannya karena hampir semua tag itu memiliki makna di dalamnya.

##### HTML 4 vs HTML 5

<p align="center">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/Html4-vs-html5-struttura.png/800px-Html4-vs-html5-struttura.png?20210128181712" alt="html4 vs html5 structure">
    <h5 align="center">Source: <a href="https://commons.wikimedia.org/wiki/File:Html4-vs-html5-struttura.png" target="_blank">commons.wikimedia.org</a></h5>
</p>

Berbeda dari penggunaan HTML4 yang biasanya menggunakan tag <div> sebagai pembentuk dari halaman aplikasi web yang kita buat. HTML5, elemen HTML sudah dibuatkan kategori dari penggunaannya.

Pada HTML4 biasanya kita membagi tiap-tiap bagian dari halaman dengan menggunakan tag `<div>` karena memang `<div>` adalah elemen yang tidak memiliki makna jadi dia hanya digunakan sebagai pembungkus atau tempat kita menggabungkan beberapa tag atau elemen HTML.

##### Kenapa perlu semantic HTML

Alasan utama kita harus menulis semantic HTML adalah:

- **Accessibility**: agar HTML kita mudah dibaca oleh user.
**- Maintainability**: agar HTML kita mudah dibaca oleh developer.
- **SEO**: agar HTML kita mudah dibaca mesin pencari.

Jadi, HTML harus ditulis dengan merepresentasikan data yang akan ditampilkan, bukan berdasarkan bagaimana dia ditampilkan. Setiap data yang ditulis harus dibungkus dengan sebuah tag yang kita sudah tahu maknanya. Membuat paragraph menggunakan `<p>` , membuat tombol menggunakan `<button>` .

Keuntungan menggunakan semantic element adalah:

- Memudahkan aksesibilitas, terutama saat menggunakan *Screen Reader.*
- Lebih mudah mencari dan mengelola kode yang memiliki makna.
- Berpengaruh pada SEO, karena semantic HTML dianggap sebagai kata kunci yang penting.

**HTML 4 Structure**

```html
<div id="header">
  <div id="logo">
    <img src="logo.png" alt="Logo Perusahaan">
  </div>
  <div id="nav">
    <ul>
      <li><a href="#">Beranda</a></li>
      <li><a href="#">Tentang Kami</a></li>
      <li><a href="#">Kontak</a></li>
    </ul>
  </div>
</div>
<div id="content">
  <div id="article">
    <h2>Judul Artikel</h2>
    <p>Isi artikel...</p>
  </div>
  <div id="sidebar">
    <h3>Artikel Terbaru</h3>
    <ul>
      <li><a href="#">Judul Artikel Terbaru 1</a></li>
      <li><a href="#">Judul Artikel Terbaru 2untuk membuat elemen yang menyatakan waktu.</a></li>
    </ul>
  </div>
</div>
<div id="footer">
  <p>&copy; 2023</p>
</div>
```

Itu adalah contoh layout yang dibuat menggunakan tag `<div>`. Tag ini
membungkus elemen HTML. Mungkin akan mudah paham dengan membaca nama-nama class yang diberikan pada elemen `div`. Seperti tidak ada masalah , Tapi nanti kalau elemen `<div>` nya sudah banyak, kita akan kesulitan membacanya.

**HTML 5 Structure**

```html
<header>
  <div id="logo">
    <img src="logo.png" alt="Logo Perusahaan">
  </div>
  <nav>
    <ul>
      <li><a href="#">Beranda</a></li>
      <li><a href="#">Tentang Kami</a></li>
      <li><a href="#">Kontak</a></li>
    </ul>
  </nav>
</header>
<main>
  <article>
    <h2>Judul Artikel</h2>
    <p>Isi artikel...</p>
  </article>
  <aside>
    <h3>Artikel Terbaru</h3>
    <ul>
      <li><a href="#">Judul Artikel Terbaru 1</a></li>
      <li><a href="#">Judul Artikel Terbaru 2</a></li>
    </ul>
  </aside>
</main>
<footer>
  <p>&copy; 2023</p>
</footer>
```

Kode ini akan lebih mudah dibaca dibandingkan kode sebelumnya.

##### Tag Semantic HTML
Berikut adalah beberapa tag HTML yang sering digunakan dalam semantic HTML:

- `<header>` : digunakan untuk bagian atas halaman web yang biasanya berisi logo, navigasi, dan judul halaman.
- `<nav>` :  digunakan untuk menandai bagian navigasi pada halaman web.
- `<main>` : digunakan untuk menandai konten utama pada halaman web.
- `<section>` : digunakan untuk membagi bagian dari halaman web menjadi beberapa bagian yang terkait dengan topik tertentu.
- `<article>` : digunakan untuk menandai konten artikel yang berdiri sendiri dan dapat dibagikan.
- `<aside>` : digunakan untuk menandai konten yang terkait atau pendukung dari konten utama pada halaman web.
- `<footer>` : digunakan untuk bagian bawah halaman web yang biasanya berisi informasi hak cipta, tautan legal, dan informasi kontak
- `<details>` : untuk membuat elemen detail atau spoiler.
- `<figcaption>` :  untuk membuat teks caption pada figure.
- `<figure>` : untuk membuat figure atau gambar pada artikel.
- `<mark>` : untuk menandai teks.
- `<summary>` : untuk membuat ringkasan artikel atau isi spoiler.
- `<time>` : untuk membuat elemen yang menyatakan waktu.

#### Membuat Layout dengan Semantik HTML

Buatlah file baru dengan nama `layout.html`, dan tambahkan kode berikut:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Contoh Layout</title>
</head>

<body>
    <header>
        <h1>Belajar Semantik HTML</h1>
    </header>
    
    <nav>
        <a href="#">Home</a>
        <a href="#">About</a>
    </nav>

    <article>
        <h1>Belajar TAG Semantik</h1>
        <p>
            Sematik adalah elemen yang memiliki makna dan tujuan. Tujuannya
            agar kode HTML mudah dibaca dan tidak ada penyalahgunaan tag.
        </p>
    </article>

    <footer>Copyright $copy; 2023</footer>
</body>
</html>
```

Buka dokumen HTML nya dengan menggunakan browser, maka akan terlihat hasilnya seperti ini :

<p align="center">
    (tampilkan hasil code dibawah)
    <h5 align="center">Source: https://</h5>
</p>

#### Testing Semantik HTML

Bagaimana jika kita ingin menambahkan *semantic element* pada elemen yang tidak memiliki *semantic*, karena ada elemen yang tidak punya *semantic* khususnya, tapi kita ingin membuatnya agar elemen tersebut tetap baik ketika dibaca oleh *Screen Reader* khususnya untuk ***accessibility***.


Caranya kita bisa menambahkan sesuatu yang disebut WAI-ARIA **(Web Accessibility Initiative - Accessible Rich Internet Applications).**

Buatlah file baru dengan nama `test-layout.html` dan isi dengan kode berikut:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Welcom to Website</h1>
    <div id="navigation">
        <h2>Navigation Title</h2>
        <ul>
            <li><a href="#">Link 1</a></li>
            <li><a href="#">Link 2</a></li>
            <li><a href="#">Link 3</a></li>
        </ul>
    </div>
</body>
</html>
```

Sebelum mejalankan *file* nya, pasang terlebih dahulu *Extension* <a href="https://chrome.google.com/webstore/detail/screen-reader/kgejglhpjiefppelpmljglcjbhoiplfn?hl=en">Screen Reader Chrome</a>.

Selanjutnya buka dokumen html yang sudah dibuat tadi. Maka akan terdengar bahwa *tag* <div id="navigation"> tidak terbaca oleh *Screen Reader*.

Lalu coba rubah dokumen `test-layout.html` dengan kode *semantic element* berikut, dan jalankan kembali dan gunakan *Screen Reader*:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Welcom to Website</h1>
    <nav>
        <h2>Navigation Title</h2>
        <ul>
            <li><a href="#">Link 1</a></li>
            <li><a href="#">Link 2</a></li>
            <li><a href="#">Link 3</a></li>
        </ul>
    </nav>
</body>
</html>
```
Maka tag `nav` akan terbaca oleh *Screen Reader*. kode di atas adalah semantic element, bagaimana jika kita ingin menambahkan semantic element pada elemen yang tidak memiliki semantic? Kita ubah file `test-semantic.html` dengan kode berikut, lalu jalankan kembali dan gunakan *Screen Reader*:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Welcom to Website</h1>
    <div id="navigation" role="navigation" aria-label="my-main-navigation">
        <h2>Navigation Title</h2>
        <ul>
            <li><a href="#">Link 1</a></li>
            <li><a href="#">Link 2</a></li>
            <li><a href="#">Link 3</a></li>
        </ul>
    </div>
</body>
</html>
```

Maka akan terdengar bahwa tag `div id="navigation" role="navigation" aria-labe:
main-navigation”>` terbaca sebagai element semantic oleh *Screen Reader*.



