# HTML Attribute

<p align="justify">
HTML <em>Atttributes</em> berfungsi untuk merubah perilaku atau tampilan pada elemen yang dituliskan didalam *opening tag*.
</p>

```html
<a href="https://www.google.com/">Google Search Engine</a>
<p style="color:blue">This is Lorem Ipsum</p>
```

Kode elemen `<a>`  di atas menggunakan *attribute* yang digunakan untuk merubah perilaku elemen saat berinteraksi dengan teks `"Google Search Engine"` maka akan mengunjungi halaman pencarian google. Sedangkan elemen `<p>` menggunakan *attribute* nya untuk merubah tampilan menjadi teks berwarna biru.

### Global Attributes

*Global Attributes* adalah atribut yang bisa digunakan untuk semua elemen HTML. Untuk contoh *global attributes* yang sering digunakan adalah `class`, `id`, dan `style`.

### **`class`**

Atribut **`class`** digunakan untuk menjadi pengenal pada elemen HTML.

```html
<body>
    <div class="keranjang-buah">
        <ol>
            <li>Apel 1</li>
            <li>Apel 2</li>
        </ol>   
    </div>

    <div class="keranjang-buah">
        <ol>
            <li>Mangga 1</li>
        </ol>   
    </div>

    <div class="kotak">
        <ol>
            <li>Novel</li>
        </ol>   
    </div>
</body>
```

Dari kode diatas mengumpakan, dari 3 tempat penyimpanan, `class="keranjang-buah"` yang berisikan buah-buah an digunakan agar lebih mudah dikenali kalau isi dari penyimpanan nya adalah buah.

### **`id`**

Atribut **`id`** berfungsi untuk menjadi pengenal dari sebauh elemen, Sehingga diharuskan `id` tidak memiliki nama yang sama pada elemen lain di satu dokumen HTML.

Analogi dari `id` bisa menggunakan 3 keranjang buah, agar membantu untuk mengenali bisa diberikan `id` seperti `keranjang-apel`, `keranjang-pisang`, dan `keranjang-mangga` di tiap keranjang nya.

```html
<body>
    <div id="keranjang-apel">
        <ol>
            <li>Apel 1</li>
            <li>Apel 2</li>
        </ol>
    </div>

    <div id="keranjang-pisang">
        <ol>
            <li>Pisang 1</li>
            <li>Pisang 2</li>
            <li>Pisang 3</li>
        </ol>
    </div>

    <div id="keranjang-mangga">
        <ol>
            <li>Mangga 1</li>
        </ol>
    </div>
</body>
```

### **`style`**

Atribut **`style`** digunakan untuk merubah tampilan pada elemen yang diinginkan dengan menggunakan CSS dalam satu baris yang sama.

Atribut ini biasa nya tidak di rekomendasikan karena *developer* lebih sering menggunakan CSS *file* yang terpisah.

```html
<p style="color:blue">This is Lorem Ipsum</p>
```

Contoh di atas menggunakan color pada CSS didalam kode baris yang sama pada value `style` untuk merubah warna teks dari *content* paragraf.

Untuk daftar **attributes global** yang lebih lengkap bisa dilihat pada  [Dokumentasi W3schools]([https://](https://www.w3schools.com/tags/ref_standardattributes.asp)).

### Spesific Attributes

Di beberapa elemen HTML ada yang memiliki atribut yang hanya bisa dipakai pada elemen nya sendiri atau ke beberapa elemen saja. 

Di bawah ini beberapa daftar dari *spesific tag attributes* yang penggunaan nya hanya bisa di elemen yang spesific.

### **`href`**

Atribut `href` digunakan untuk elemen `<a>` sebagai *hyperlink* untuk berpindah ke *link* yang di tempatkan di *value* `href`, dengan berinteraksi dengan teks yang ada di *content* element `<a>`.

```html
<a href="https://www.google.com/">Go to Google Seatch Engine</a>
```

elemen diatas akan memindahkan halaman web yang sedang dikunjungi ke halaman pencarian google.

Tag yang menggunakan atribut `href` selain `<a>` tidak akan menjadi *hyperlink*.

### **`selected`**

Atribut `selected` merupakan sebuah atribut **boolean** yang digunakan untuk elemen `<option>`, atribut ini menentukan opsi yang terpilih sebelum halaman selesai dimuat. Opsi yang sudah di tentukan tadi akan dimuat terlebih dahulu pada daftar *drop-down* yang menggunakan elemen `<select>`.

```html
<body>
	<label for="buah">Pilih Buah:</label>

	<select id="buah">
  		<option value="ap">Apel</option>
  		<option value="saab">Pisang</option>
  		<option value="vw" selected>Durian</option>
	</select>
</body>
```

Untuk perbedaan `<option>` yang menggunakan atribut `selected` bisa di lihat dari gambar dibawah :

#### Menggunakan `selected` attribute

<p align="center">
    (insert option use selected web preview)
    <h5 align="center">Source: https://</h5>'
</p>

*Drop-down* menu akan langsung menampilkan durian karena menggunakan atribut `selected`.

#### Tidak Menggunakan `selected` attribute

<p align="center">
    (insert option use selected web preview)
    <h5 align="center">Source: https://</h5>'
</p>

*Drop-down*  menampilkan opsi pertama yaitu apel karena tidak ada elemen `<option>` yang menggunakan atribut`selected`.
