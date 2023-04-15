# CSS Comment

Sama seperti membuat HTML, di dalam _syntax CSS_ kita juga bisa melakukan `comment`. Comment di CSS biasanya digunakan untuk memberikan penjelasan tambahan mengenai suatu _syntax CSS_.

CSS comment dimulai dengan `/*` dan di akhiri dengan `*/` .

```css
/* Merupakan css untuk Heading 1 */
h1 {
  color: red;
}
```

Kita juga bisa membuat _CSS Comment_ di dalam baris _syntax_ itu sendiri.

```css
h1 {
  color: red; /* Property yang merubah warna tulisan */
}
```

_CSS comment_ bisa juga membuat sebagian _syntax_ tidak _dirender_ oleh _web browser_.

```html
<!DOCTYPE html>
<html>
  <head>
    <style>
      h1 {
        color: red; /* Set text color to red */
        border-bottom: 1px solid blue;
      }
    </style>
  </head>
  <body>
    <h1>Hello, world!</h1>
  </body>
</html>
```

<p align=center>
    <img src="">
    <br> (Output code diatas)
</p>

Kemudia jika kita melakukan _comment_ seperti ini:

```html
<!DOCTYPE html>
<html>
  <head>
    <style>
      h1 {
        color: red; /* Set text color to red */
        /* border-bottom: 1px solid blue; */
      }
    </style>
  </head>
  <body>
    <h1>Hello, world!</h1>
  </body>
</html>
```

<p align=center>
    <img src="">
    <br> (Output code diatas)
</p>

`border-bottom` tidak akan di _render_ oleh _web browser_.
