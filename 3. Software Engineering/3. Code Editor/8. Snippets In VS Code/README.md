# Snippets In Visual Studio Code
## What Is Snippet?
<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br>

<p align="justify">
<strong>Code snippets</strong> adalah template yang memudahkan untuk memasukkan code patterns secara berulang, seperti <code>loop</code> atau <code>conditional-statements</code>. Dengan snippet, kita dapat membuat file boilerplate, dan menyisipkan blok teks yang biasa kita gunakan secara cepat.<br><br>

## How To Use Snippet In VS Code
<p align="justify">
Di VS Code ada banyak snippet yang sudah di implement secara default dan bisa langsung digunakan. VS Code juga memungkinkan kita membuat snippet versi kita sendiri yang akan bisa kita gunakan di semua project, atau membuat snippet yang hanya bisa digunakan untuk project tertentu serta dapat diakses oleh semua tim project, atau mengimpor snippet dari extensions.<br><br>

## Built-in Snippets
<p align="justify">
VS Code memiliki snippet bawaan untuk berbagai programming language, scripting language, ataupun markup language seperti: JavaScript, TypeScript, Markdown, HTML dan PHP.<br>

<p align="center">
<img height="150rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/builtin-javascript-snippets.png"><br>

<p align="justify">
Kita dapat melihat list snippet yang tersedia untuk suatu bahasa dengan menjalankan perintah <strong>Insert Snippet</strong> di Command Palette. Output-nya akan ditampilkan list snippet untuk bahasa dari file yang sedang dibuka saat ini.<br>

<p align="justify">
Namun, perlu diingat bahwa daftar ini juga menyertakan snippet yang telah kita defined, dan snippet apa pun yang disediakan oleh extensions yang telah kita pasang.<br>

## Install Snippets From The Marketplace
<p align="justify">
Banyak extensions di VS Code Marketplace yang merupakan kumpulan dari snippets. Kita dapat menelusuri extensions yang berisi snippets di tampilan extensions <code>Ctrl + Shift + X</code> menggunakan filter <code>@category: "snippets"</code>.<br>

<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/category-snippets.png"><br>

<p align="justify">
Jika kita menemukan extensions yang ingin kita gunakan, install, lalu restart VS Code dan snippets baru akan tersedia.<br><br>

## Create Your Own Snippets
<p align="justify">
Kita dapat dengan mudah membuat snippets versi kita sendiri. Untuk membuat atau mengedit snippets kita sendiri, pilih <strong>User Snippets</strong> pada menu <strong>File > Preferences</strong>, lalu pilih bahasa (berdasarkan pengenal bahasa) yang akan menampilkan snippets, atau opsi <strong>New Global Snippets file</strong> jika kita ingin snippets tersebut muncul untuk semua bahasa.<br>

<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/snippet-dropdown.png"><br>

<p align="justify">
File snippets ditulis dalam <code>JSON</code>, mendukung C-style comments, dan dapat men-define jumlah snippets yang tidak terbatas. Snippets mendukung sebagian besar TextMate syntax untuk behavior yang dinamis, memformat whitespace berdasarkan insertion context, dan memungkinkan pengeditan multiline yang mudah. TextMate syntax sendiri yaitu format daftar properti dalam bentuk XML di mana ekstensi file-nya adalah <code>.tmLanguage</code>.<br>

<p align="justify">
Di bawah ini adalah contoh <code>for loop</code> snippet untuk JavaScript:<br>

```
// in file 'Code/User/snippets/javascript.json'
{
  "For Loop": {
    "prefix": ["for", "for-const"],
    "body": ["for (const ${2:element} of ${1:array}) {", "\t$0", "}"],
    "description": "A for loop."
  }
}
```

<br>

<p align="justify">
Berikut di bawah ini penjelasan dari snippet di atas:<br>

- <code>For Loop</code> adalah nama dari snippet yang kita buat.
- <code>prefix</code> mendefinisikan satu atau beberapa kata yang me-trigger yang akan menampilkan snippet di IntelliSense. Substring matching dilakukan pada prefixes, jadi dalam kasus ini, <code>fc</code> bisa cocok dengan <code>for-const</code>.
- <code>body</code> adalah satu atau lebih baris dari sebuah konten, yang akan digabungkan sebagai beberapa baris saat di sisipkan. Baris baru dan tab yang disematkan akan di-format sesuai dengan konteks penyisipan snippet.
- <code>description</code> adalah deskripsi opsional dari snippet yang ditampilkan oleh IntelliSense.

<br><br>
