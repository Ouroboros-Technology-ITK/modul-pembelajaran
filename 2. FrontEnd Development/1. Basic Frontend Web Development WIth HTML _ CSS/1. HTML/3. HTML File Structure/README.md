# HTML File Structure
<p align="justify">
Contoh dokumen HTML dari strukturnya untuk menyusun tampilan web.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Document</title>
</head>
<body>
    <h1>Heading Of Page</h1>
    ...
</body>
</html>
```

#### `<html>` Root element


Elemen `<html>` berfungsi untuk membungkus semua konten di halaman.

#### `<head>` Dokumen metadata element


Elemen `<head>` berfungi untuk menjadi wadah untuk keyword atau informasi tentang halaman yang dapat dibaca oleh mesin (<em>metadata</em>) tentang dokumen. 

Konten elemen <em>metadata</em> yang biasanya digunakan seperti <em> title, scripts, dan style sheets </em> penerapan element `<head>` diletakkan diantara `<html>` dan `<body>`.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Page's Title</title>
    <meta charset="utf-8">
    <link rel="stylesheet" href="styles.css">
    
</head>
<body>
    ...
</body>
</html>
```

Elemen-elemen yang dapat dimasukkan ke dalam `<head>` :

`<title>`:Digunakan untuk merubah title pada tab bar.

<p align="center">
    (insert image shows title and another component from the element)
    <h5 align="center">Source: https://</h5>'
  </p>

`<meta name="description" content="Lorem ipsum"/>`: memberikan ringkasan singkat dari halaman web pada hasil pencarian.

<p align="center">
    (insert image shows description of web from google search result)
    <h5 align="center">Source: https://</h5>'
  </p>


`<meta name="viewport" content="width=device-width, initial-scale=1" />` : adalah bagian dari sebuah halaman web yang bisa terlihat oleh pengguna atau menetapkan area mana saja yang dapat di lihat di halaman web dan menentukan skala konten nya.

Tampilan akan dapat berubah tergantung dari perangkat yang sedang digunakan,

<p align="center">
    (insert image shows description of web from google search result)
    <h5 align="center">Source: https://</h5>'
  </p>

Sangat penting untuk meningkatkan pengalaman pengguna dari berbagai macam perangkat. Agar pengguna betah saat membuka halaman web kita. Karena itu penting untuk mengembangkan web yang *mobile-friendly*.

#### `<body>` : Dokumen body element

Elemen `<body>` digunakan untuk membungkus atau mewakili konten dokumen HTML yang ingin ditampilkan ke user ketika mengunjungi suatu halaman. Elemen yang berada di body dapat berupa judul, paragraf, gambar, tabel, <em>hyperlink</em>, dan yang lainnya.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Page's Title</title>
    ...
</head>
<body>
    <h1>Heading 1</h1>
    <img src="img.jpg" width="500" height="500">

    <a href="https://google.com">Google</a>
    <p> Google search engine </p>
    ...        
</body>
</html>
```

# Family Tree of HTML Structure

Setiap dokumen HTML dapat di anggap sebagai <em>document tree</em>. Deskripsi dari tiap elemen pada <em>tree</em> bisa dijelaskan seperti mendeskirpsikan diagram pohon keluarga. Dimana terdapat <em>parent, child</em> dan <em>siblings</em>.

```html
<body>
    <div>
        <h1>Heading</h1>
        <p>Paragraf</p>
    </div>

    <div>
        <table>
            <tr>
              <th>Names</th>
              <th>Age</th>
            </tr>
            <tr>
              <td>Franks</td>
              <td>21</td>
            </tr>
            <tr>
              <td>John Doe</td>
              <td>55</td>
            </tr>
        </table>
    </div>
</body>
```

Jika struktur diatas digambarkan pada diagram akan menjadi seperti dibawah :

<p align="center">
    (insert image tree diagram)
    <h5 align="center">Source: https://</h5>'
</p>

#### Parent

<p align="center">
    (insert image tree diagram : colored the div with the child)
    <h5 align="center">Source: https://</h5>'
</p>

<em>Parent element</em> adalah elemen yang ada diatas dan terhubung dari elemen lain didalamnya. Untuk contoh dari diagram diatas `<div>` adalah <em>parent</em> dari `<h1>` dan `<p>`.

#### Child

<em>Child element</em> adalah elemen yang terhubung secara langsung dibawah <em>Parent element</em>. Pada diagram diatas `<h1>` dan `<p>` adalah <em>child elements</em> dari `<div>`.

#### Sibling

<p align="center">
    (insert image tree diagram : colored table to td)
    <h5 align="center">Source: https://</h5>'
</p>

<em>Sibling elements</em> merupakan elemen yang berada pada satu <em>parent element</em> yang sama. Dalam diagram diatas tiap tiga elemen `<td>` berbagi satu <em>parent</em> yang sama yaitu `<tr>`, dan tiga elemen `<tr>` merupakan <em>sibling elements</em> juga dengan <em>parent element</em> table.

#### Ancestor & Descendant

<strong><em>Ancestor</em> </strong>adalah sebuah elemen yang berada di paling atas <em>tree diagram</em>. Untuk contoh diagram dibawah  elemen `<body>` adalah ancestor atau nenek moyang dari diagramnya.

<p align="center">
    (insert image tree diagram : colored body as ascendant and all below element as descendant)
    <h5 align="center">Source: https://</h5>'
</p>

<strong><em>Descendant</em></strong> merupakan elemen-elemen yang terhubung tetapi lebih rendah dari elemen lain. Dari diagram diatas elemen `<table>` sampai `<td>` merupakan <em>descendant</em> dari `<div>`.
