# What Is CSS ?

<p align=justify>
CSS merupakan singkatan dari Cascading Style Sheet,
berguna untuk mempercantik tampilan dari sebuah dokumen yang 
berformat markup language seperti HTML ataupun XML.
Ketika dirender pada web browser, CSS ini berperan penting dalam membuat hasil render HTML
ataupun XML tersebut menjadi bagus dan enak untuk dilihat.
</p>

<p align=justify>
Banyak hal bisa dilakukan dengan menggunakan CSS,  kita dapat mengubah atau memodifikasi tampilan dari elemen-elemen HTML seperti warna, font, ukuran, layout, dan lain sebagainya.

Penggunaan CSS dapat untuk mengatur bagaimana elemen-elemen di dalam dokumen HTML seharusnya tampil pada layar ketika dirender oleh web browser, seperti berapa ukuran font dari judul artikel, berapa lebar dan tinggi suatu gambar, berapa tebal garis pada suatu tabel, dan masih banyak lagi.

</p>

<p align=center>
    <img src="">
    <br> (Struktur HTML vs CSS)
    <br> Source : https://
</p>

<p align="justify">
    Bahasa yang lebih <i>simple</i> nya CSS itu adalah sebuah syntax yang digunakan untuk mempercantik HTML yang sudah kita buat atau menambahkan design ke suatu halaman HTML, seperti menjadi baju pada tubuh manusia.
</p>

Pahami 2 buah _keyword_ ini untuk mendapatkan pemahaman menyeluruh tentang CSS, yaitu **Cascading & Style Sheet**

### Style Sheet

Sebuah syntax yang akan digunakan untuk melakukan pengaturan bagaimana bentuk tampilan dan layouting dari sebuah dokumen yang berformat markup language.

```css
button {
  background-color: #4caf50;
  border: none;
  color: white;
  padding: 10px 20px;
  font-size: 16px;
}

p {
  font-size: 18px;
  line-height: 1.5;
  margin: 10px 0;
}
```

Diatas adalah contoh penggunaan _style sheet_ didalam sebuah dokumen CSS:

- `button` dan `p` disebut sebagai **CSS Selector**
- `background-color`, `border`, `font-size`, `padding` dan yang lainnya yang berada didalam CSS Selector disebut dengan **CSS Property**
- `#4CAF50`, `white`, `10px` dan semua yang berada disamping CSS Property merupakan **Value dari sebuah CSS Property**

### Cascade

Cascading berasal dari kata Cascade yaitu sebuah algoritma yang akan mendefinisikan bagaimana urutan dari sebuah style akan diimplementasikan kedalam markup language.

Untuk menentukan CSS Property mana yang akan diimplementasikan ketika terjadinya conflict atau kondisi tumpang tindih CSS Selector di dalam sebuah dokumen HTML.

```html
<button class="btn">Click me</button>
```

<p align=center>
    <img src="">
    <br> (button image from code above)
    <br> Source : https://
</p>

```css

```

Jika ada 2 buah CSS Selector yang mencoba menggunakan CSS Property yang sama seperti contoh di atas, maka urutan yang paling bawah lebih dahulu.

<p align=center>
    <img src="">
    <br> (button image from code above after put stylesheet)
    <br> Source : https://
</p>

Lalu bagaimana dengan penggunaan property yang berbeda di 2 selector yang sama seperti dibawah makan output nya akan menggabungkan 2 property nya dan dijalankan bersamaan dengan urutan atas ke bawah.

Konsep cascade inilah yang akan mengatur tampilan seperti apa yang akan dirender oleh web browser jika ada CSS Selector yang tumpang tindih atau conflict seperti contoh di atas.

konsep cascade ini sendiri akan menggunakan sistem scoring terhadap CSS selector. Sistem scoring ini bernama css _specificity_, berguna untuk menentukan CSS _rules_ mana yang akan diterapkan kedalam suatu HTML element (elemen HTML) ketika terjadi _conflict_.

Contohnya `id` mempunyai _score_ yang lebih tinggi dibandingkan dengan _selector_ `class`, maka CSS yang seperti ini:

```html
<!DOCTYPE html>
<html>
  <head>
    <style>
      #title {
        color: blue;
      }
      .heading {
        color: red;
      }
    </style>
  </head>
  <body>
    <h1 class="heading" id="title">Hello World</h1>
  </body>
</html>
```

Heading 1 akan berwarna biru karena selector id(#title) punya score _specifity_ lebih besar dibandingankan dengan selector class (.heading).

# Cara Kerja CSS

Prinsip utama dari cara kerja CSS Ada 2 part ketika CSS akan bekerja, yaitu:

1. Memilih `HTML Element` dengan menggunakan `CSS Selector`
2. Mengimplementasikan `css Property` dan `css Property Value` terhadap `HTML Element` yang sudah dipilih tadi.

### Part 1 : Pilih `HTML Element`

nantinya akan diimplementasikan `rule` CSS dari `CSS Property` & `Value`. Pemilihan HTML Element ini terlihat mudah pada awalnya, namun semakin rumit struktur sebuah documen HTML, maka untuk memilih `HTML Element` juga semakin rumit.

Beberapa tipe CSS _selector_ bisa kita gunakan untuk memilih _Element HTML_ yang tepat.

- **Simple Selector** : `selector by tag name`, `selector by id`, dan `selector by class name`.

- **Advance Selector** : `Combinator Selectors`, `Pseudo Class Selectors`, `Pseudo Element Selectors`, dan `Attribute Selectors`.

```css
p {
  /* akan digunakan ke semua tag <p> */
}

.btn {
  /* hanya akan digunakan di <button name="btn"></button> */
  /* tidak akan terpanggil di <button></button>*/
}

#title {
  /* akan digunakan oleh tag yang memanggilnya dengan
  * <h1 id="title"></h1>
  */
}
```

### Part 2 : Implementasi `CSS Property` & `Value`

kita mengimplementasikan `CSS Property` beserta `Value` -nya kepada `HTML Element` yang sudah dipilih pada part sebelumnya. Ada banyak sekali `CSS Property` yang bisa digunakan. Masing-masing `CSS Property` ini juga memiliki Value valid yang berbeda-beda juga.

Contohnya kita dapat menggunakan `font-size` untuk mengatur ukuran dari suatu _font_ salah satu value valid nya berupa _fixed size_ dalam satuan unit px dan cm.

```css
p {
  font-size: 16px;
  color: #333;
}
```

#### Property Values

Ini adalah tabel yang menampilkan daftar nilai untuk properti font-size dalam CSS.

| Nilai      | Deskripsi                                                                 |
| ---------- | ------------------------------------------------------------------------- |
| medium     | Mengatur ukuran font menjadi ukuran medium (default)                      |
| xx-small   | Mengatur ukuran font menjadi ukuran xx-small                              |
| x-small    | Mengatur ukuran font menjadi ukuran extra small (x-small)                 |
| small      | Mengatur ukuran font menjadi ukuran small                                 |
| large      | Mengatur ukuran font menjadi ukuran large                                 |
| x-large    | Mengatur ukuran font menjadi ukuran extra large (x-large)                 |
| xx-large   | Mengatur ukuran font menjadi ukuran xx-large                              |
| smaller    | Mengatur ukuran font menjadi ukuran yang lebih kecil dari elemen induknya |
| larger     | Mengatur ukuran font menjadi ukuran yang lebih besar dari elemen induknya |
| length / % | Mengatur ukuran font menjadi ukuran tetap dalam satuan px, cm, dll.       |

```css
.btn {
  font-size: 20px;
}
```

Begitu juga dengan `CSS Property` lainnya, ada banyak value _valid_ untuk masing masin `CSS Property`, namun tidak perlu untuk menghapal semuanya. Kita cukup hanya mmembaca *list*nya di awal belajar, sehingga kita tahu ada `CSS Property` apa saja yang bisa digunakan dan _value valid_ nya.

Banyak sumber untuk mencari `CSS Property` lainnya, seperti yang dapat dilihat pada link berikut : [CSS Properties W3S](https://www.w3schools.com/cssref/index.php).

# Tools untuk Membuat CSS

Setelah memahami dan mempelajari tentang kegunaan dari CSS, yang untuk memperbagus suatu tampilan _document_ HTML ketika _dirender_ oleh _web browser_. Proses _development_ CSS akan berdampingan dengan proses _development_ HTML. _Tools_ yang digunakan pun akan sama untuk pembuatan dokumen CSS & HTML seperti **code editor** dan **web browser**.

### Code Editor

Sama seperti HTML, untuk membuat suatu CSS kita perlu sebuah _code editor_. Pada dasarnya CSS sama seperti HTML, berupa plain text yang disimpan ke dalam file document HTML atau suatu file khusus dengan ekstensi `.css`. Tidak ada batasan text editor apa yang bisa digunakan untuk membuat CSS. Namun pembuatan nya sering menggunakan **IDE VS Code** dan yang akan digunakan pada pembelajaran di modul ini.

### Web Browser

Untuk melihat hasil _render_ sebuah dokumen HTML, kita memerlukan web browser. Untuk melihat hasil render CSS juga sama, cukup gunakan web browser untuk membuka document HTML yang sudah ada penerapan CSS, maka hasil _render_ document HTML tersebut akan mengikut sertakan render CSS.

<p align=center>
    <img src="https://devopedia.org/images/article/282/9041.1597665465.jpg">
    <br> CSS in HTML Render
    <br> (Source: <a href="https://devopedia.org/document-object-model">https://devopedia.org/</a>)
</p>

Untuk web browser-nya, bisa menggunakan sebagian besar web browser yang tersedia secara umum pada saat ini.

<p align=center>
    <img src="https://visionlab.nz/wp-content/uploads/VL-web-browsers.jpg">
    <br> Most Popular Web Browser's
    <br> (Source: <a href="https://visionlab.nz/should-i-save-my-passwords-in-browsers-that-have-an-in-built-password-manager/">https://visionlab.nz/</a>)
</p>
