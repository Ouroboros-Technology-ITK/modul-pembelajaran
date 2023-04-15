# Struktur CSS

dengan mempelajari CSS structure ini, kita bisa menyederhanakan pemahaman bagaimana cara menggunakan CSS, dan nantinya mampu untuk menggunakan CSS dengan lebih advanced lagi.

<p align=center>
    <img src="">
    <br> (Struktur Penulisan CSS)
    <br> Source : https://
</p>

CSS Rule terdiri dari dua yaitu **CSS Selector** & **CSS Declaration**.

- Satu `CSS Rule` terdapat satu buah `selector`
- `declaration` minimal memiliki satu atau lebih pasangan `CSS Property` & `value`

### Selector

merupakan bagian yang terletak sebelum kurung pembuka **CSS Declaration**. Berisikan kombinasi dari _syntax_ untuk memilih _HTML element_ mana yang akan diterapkan **CSS rule**.

```css
p {
  /* Tidak ada declaration yang ditambahkan di sini */
}
```

Selector ini memilih semua _element_ `<p>` pada halaman web.

### Declaration

bagian yang terletak di dalam kurung _curly_ (0) setelah `CSS Selector`. Berisikan minimal 1 pasang `CSS Property` dan `Value`. Jika ada lebih dari 1 pasang `CSS Property` dan `Value`, maka diperlukan tanda; (titik koma / semi collon) sebagai pemisah dari tiap pasang `CSS Property` dan `Value`.

```css
p {
  color: red;
  font-size: 16px;
  font-style: italic;
}
```

#### `CSS Property`

Bagian yang ada di sebelah kiri dari tanda (titik dua) di dalam kurung _curly_ ({}). Merupakan _human-readable_ _identifiers_ yang akan menunjukkan perubahan _styling_ seperti apa yang akan diterapkan terhadap sebuah HTML Element.

Pada contoh diatas yang merupakan sebuah _properties_ adalah `color`, `font-size` & `font-style`.

#### `Value`

Bagian yang ada di sebelah kanan dari tanda (titik dua) di dalam kurung curly (0). Merupakan nilai perubahan _styling_ yang akan diterapkan dari sebuah `CSS Property` .

`CSS Property` dan `Value` akan dituliskan secara berpasangan. Jadi kita tidak bisa menuliskan hanya salah satu saja.

```css
p {
    /*Salah*/
    color;
}

p {
    /*Benar*/
    color: blue;
}
```

Ada banyak sekali CSS Property yang bisa digunakan. List lengkap bisa kita lihat pada link ini : [CSS Property List](https://www.w3schools.com/cssref/index.php).

Tidak perlu menghapal semua `CSS Property` tersebut. Kita hanya perlu membaca semua list yang ada ketika baru belajar agar kita bisa mengetahui ada `CSS Property` apa saja yang bisa kita gunakan. Ketika sudah tau dan untuk memastikan _syntax_ nya benar, kita dapat melihat pada link referensi untuk memastikan apakah benar.
