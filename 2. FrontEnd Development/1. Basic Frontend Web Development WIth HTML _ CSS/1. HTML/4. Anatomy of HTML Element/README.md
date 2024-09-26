# Anatomy of Element HTML

<p align="justify">

<p align="center">
    (insert anotomy of element html image)
    <h5 align="center">Source: https://</h5>'
</p>

HTML merupakan dokumen yang disusun dari bebeberapa elemen. Sedangkan penyusun elemen terdiri dari <em>opening tag, content, dan closing tag</em>. Sedangkan <em>opening tag</em> sendiri memiliki atribut yang dapat diisi dengan nilai yang berbeda-beda.

### Opening Tag

*Opening tag* adalah pembuka dari sebuah elemen yang dibungkus dengan simbol `<` dan ditambah nama tag dan diakhiri dengan simbol `>`. Disini dimulainya content yang berada pada elemen.

Contoh dari penggunaan *opening tag* dibawah adalah dengan menggunakan `div` yang berisikan *content* elemen `h3`. 

```html
<div>
    <h3>Heading</h3>
</div>
```

### Closing & Self Closing Tag

##### Closing Tag

*Closing tag* adalah penanda akhirnya konten pada elemen dimana pada contoh sebelumnya dengan menggunakanan `h3` untuk mengakhiri konten yang berbentuk header. Untuk penulisanya sama dengan *opening tag* menggunakan simbol `<` dan `>` untuk membungkus nama tag sebuah elemen, yang membedakan adalah tambahan simbol `/` sebelum nama tag.

```html
</h3>
</p>
```

##### Self Closing Tag

Berbeda dengan *closing tag* fungsi *self closing tag* digunakan ketika elemen tidak memerlukan sebuah kontent di dalamnya, seperti pengunaan `<img/>`, `<br/>`, dan `<input/>`

### Content

*Content* adalah karakter atau elemen lain yang berada pada suatu elemen.

```html
<div>
    <h1>Heading</h1>
    <p>Lorem Ipsum</p>
</div>
```

Dari contoh kode diatas teks `Heading` dan `Lorem Ipsum` merupakan *content* yang berada didalam sebuah elemen. Sedangkan elemen `h1` dan `p` adalah sebuah elemen yang menjadi *content* didalam elemen `div`.

### Element

*Element* atau elemen adalah bagian utuh yang dimulai dari *opening tag* sampai *closing tag* dan konten di dalamanya.

```html
<p>Lorem Ipsum<p>
```

Kode diatas merupakan elemen yang menggunakan elemen `p` untuk membuat sebuah paragraf `Lorem Ipsum`.
