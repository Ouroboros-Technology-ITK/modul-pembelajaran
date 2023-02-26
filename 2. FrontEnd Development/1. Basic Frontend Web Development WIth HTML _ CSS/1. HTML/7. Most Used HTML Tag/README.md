# Most Used HTML Tag

<p align="justify">
Terdapat banyak tag HTML yang sering digunakan dalam pembuatan halaman web. 

### Button

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

### Button Interactivity

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
