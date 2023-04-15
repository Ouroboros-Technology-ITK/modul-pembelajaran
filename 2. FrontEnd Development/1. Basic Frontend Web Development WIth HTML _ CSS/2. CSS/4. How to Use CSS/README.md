# Penggunaan CSS

Untuk menggunakan CSS, pada umumnya terdapat 3 cara:

- Menggunakan _Inline Stylesheet_
- Menggunakan Tag HTML `<style>`
- Menyambungkan _external file_ CSS ke HTML dengan menggunakan _Tag_ HTML `<link>`

### Inline Stylesheet

Untuk menggunakan _**Inline Stylesheet/CSS**_, kita menyisipkan langsung pada _HTML Element_ tanpa menggunakan `CSS Selector`. Untuk menggunakan metode ini, cukup dengan menambahkan _attribute style_ diikuti dengan `CSS Declaration` sebagai `value` dari _attribute style_ tersebut.

```html
<!DOCTYPE html>
<html>
  <head>
    <title>Hello World</title>
  </head>
  <body>
    <!-- Inline Stylesheet -->
    <h1 style="color: magenta; ">Hello World</h1>
  </body>
</html>
```

Tag `h1` akan _dirender_ dengan perubahan warna teks menjadi `magenta`.

Output :

<p align=center>
    <img src="">
    <br> (Output code diatas)
</p>

### Tag `<style>`

disebut juga dengan **Internal CSS**, kita menambahkan CSS dengan cara menuliskan _syntax_ CSS langsung di dalam dokumen HTML. _Syntax CSS_ ini ditulis menggunakan tag `<style>`.

```html
<html>
  <head>
    <title>Hello World External CSS</title>
    <!-- Internal Stylesheet -->
    <style>
      h1 {
        color: green;
      }
    </style>
  </head>
  <body>
    <h1>Hello World Internal CSS</h1>
  </body>
</html>
```

Tag `h1` _dirender_ dengan perubahan warna menjadi `blue`.

Output :

<p align=center>
    <img src="">
    <br> (Output code diatas)
</p>

### External Stylesheet

Semua _syntax CSS_ akan dituliskan pada sebuah _file_ dengan format `.css`, kemudian disambungkan _file_ ini dengan menggunakan Tag HTML `<link>`.

Attribute href di `<link/>` berisi spesifik _location_ dari file yang akan disambungkan ke dalam document HTML. Value valid untuk attribute ini berupa absolute url dan relative url.

Untuk attribute rel dan type dari `<link/>` bisa diisi dengan `value stylesheet` dan `text/css`. Sebagai contoh, kita membuat 2 buah file: 1 file HTML dan 1 file CSS di dalam folder yang sama:

**File HTML : index.html**

```html
<!-- File HTML -->
<html>
  <head>
    <title>Hello World External CSS</title>
    <!-- External Stylesheet -->
    <link rel="stylesheet" href="style.css" />
  </head>
  <body>
    <h1>My Web Page</h1>
    <p>Hello World!</p>
  </body>
</html>
```

File HTML akan melakukan _render_ dengan menggunakan CSS yang ada pada file `style.css`.

**File CSS : style.css**

```css
/* File CSS */
body {
  background-color: pink;
}
h1 {
  color: blue;
}
P {
  color: black;
}
```

Output :

<p align=center>
    <img src="">
    <br> (Output code diatas)
</p>

_Tag_ `body` dirender dengan perubahan warna _background_ menjadi `pink`, tag `h1` dirender dengan perubahan warna menjadi `blue`, dan tag `p` dirender dengan perubahan warna menjadi `black`.

Dari ketiga cara menggunakan CSS di atas, tidak ada patokan wajib menggunakan yang mana. Kita bebas memilih menggunakan cara mana yang kita mau ketika ingin menggunakan **CSS**.
