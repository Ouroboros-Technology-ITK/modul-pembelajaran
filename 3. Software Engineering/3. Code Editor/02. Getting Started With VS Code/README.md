# Getting Started With VS Code 
<p align="justify">
VS Code adalah salah satu code editor yang ringan, powerful dan dapat bejalan di berbagai macam sistem operasi seperti macOS, Linux, dan Windows.</p>

## Install VS Code
##### Linux Ubuntu
<p align="justify">
Cara termudah untuk mengintall Visual Studio Code untuk Debian/Ubuntu adalah dengan mengunduh dan menginstall package <code>.deb (64-bit)</code>, baik melalui <em>Graphical Software Center</em>, atau melalui <em>Command Line</em> berikut ini:</p>

```
sudo apt install ./<file>.deb

# Jika kita menggunakan Linux Distribution yang lebih lama, kita harus menjalankan perintah di bawah ini
# sudo dpkh -i <file>.deb
# sudo apt-get insttall -f # install dependencies
```

<p align="justify">
Selain itu kita juga dapat melakukan instalasi menggunakan <strong>Snap Store</strong> dengan cara:</p>

```
sudo snap install --classic code #or code-insiders
```

<p align="justify">
Perlu diperhatikan bahwa sebelumnya kita perlu melakukan instalasi <strong>Snap Store</strong> terlebih dahulu jika belum tersedia.</p><br>

##### macOS
<p align="justify">
Untuk melakukan instalasi pada macOS kita dapat mengikuti alur berikut:</p>
<ol style="list-style-type:circle;" style="text-align:justify">
  <li>Download <a href="https://code.visualstudio.com/download" target="_blank"><strong>Visual Studio Code</strong></a> untuk macOS.</li>
  <li>Buka daftar download browser dan cari aplikasi atau archive yang di-download.</li>
  <li>Jika file dalam bentuk archive, ekstrak isi archive. Gunakan double-click untuk beberapa browser atau pilih ikon 'kaca pembesar' dengan Safari.</li>
  <li>Drag <code>Visual Studio Code.app</code> ke folder aplikasi, sehingga membuatnya tersedia di Launchpad macOS.</li>
</ol><br>

##### Windows
<p align="justify">
Untuk melakukan instalasi pada Windows kita dapat mengikuti alur berikut:</p>
<ol style="list-style-type:circle;" style="text-align:justify">
  <li>Download <a href="https://code.visualstudio.com/download" target="_blank"><strong>Visual Studio Code</strong></a> untuk Windows.</li>
  <li>Setelah didownload, jalankan installer <code>VSCodeUserSetup-{version}.exe</code>.</li>
  <li>Secara default, VS Code akan diinstall pada path <code>C:\Users\{username}\AppData\Local\Programs\Microsoft VS Code</code>.</li>
</ol>

<p align="justify">
Untuk pengguna Windows, karena kedepannya kita akan sering menggunakan <em>command</em> Linux dalam setiap pembelajaran. Jadi disarankan agar kita menggunakan WSL sebagai terminal untuk menjalankan setiap beris perintah Linux pada VS Code. Caranya dengan menginstall <strong>Remote Development</strong> pada VS Code:</p>
<ol style="list-style-type:circle;" style="text-align:justify">
  <li>Arahkan kursor ke menu di paling kanan dari jendela VS Code.</li>
  <li>Pilih menu <strong>Extensions</strong> (Ctrl+Shift+x).</li>
  <li>Pada input pencarian, ketik "Remote Development", lalu pilih <strong>Extensions</strong> teratas dari hasil pencarian.</li>
  <li>Klik tombol <strong>Install</strong>.</li>
</ol>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/remote-dev.png"><br>

<p align="justify">
Untuk meilhat bagaimana cara menggunakannya kita bisa membuka dokumentasi <a href="https://code.visualstudio.com/docs/remote/wsl" target="_blank"><strong>Developing in WSL</strong></a>.<br><br>

## User Interface
<p align="justify">
Seperti banyak code editor lainnya. VS Code mengadopsi user interface yang umumnya digunakan dan layout dari explore terdapat di sebelah kiri, yang mana akan menampilkan semua file dan folder yang dapat akan diakses, dan editor di sebelah kanan yang akan menampilkan konten file yang telah dibuka.</p><br>

## Basic Layout
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/basic-layouy.png">

<p align="justify">
VS Code hadir dengan layout sederhana dan intuitif, memaksimalkan ruang yang disediakan untuk editor sambil menyisakan banyak ruang untuk menelusuri dan mengakses full context folder atau proyek yang dimiliki.</p>

<p align="justify">
UI VS Code dibagi menjadi lima area:</p>

- <strong>Editor</strong>
<p align="justify">
Area utama untuk mengedit file yang dimiliki. Kita dapat membuka editor sebanyak yang kita suka secara berdampingan secara vertikal dan horizontal.</p>

- <strong>Side Bar</strong>
<p align="justify">
Berisi tampilan berbeda seperti Explorer untuk membantu kita saat mengerjakan proyek.</p>

- <strong>Status Bar</strong>
<p align="justify">
Informasi tentang proyek yang dibuka dan file yang kita edit.</p>

- <strong>Ectivity Bar</strong>
<p align="justify">
Terletak di sisi paling kiri, activity bar menunjukkan untuk beralih di antara tampilan dan memberi context-specific indicators tambahan, seperti jumlah perubahan outgoing saat Git diaktifkan.</p>

- <strong>Panel</strong>
<p align="justify">
Kita dapat menampilkan panel yang berbeda di bawah wilayah editor untuk output atau informasi debug, errors dan warnings, atau intergrated terminal. Panel juga dapat dipindahkan ke kenan untuk ruang yang lebih vertikal.</p><br>

## Side By Side Editing
<p align="justify">
Kita dapat membuka editor sebanyak yang kita suka secara berdampingan, baik secara vertikal maupun horizontal. Jika kita sudah membuka satu editor, ada beberapa cara untuk membuka editor lain di samping editor yang sudah ada:</p>

<ul style="list-style-type:circle;" style="text-align:justify">
  <li><code>Alt</code> klik pada file di Explorer.</li>
  <li><code>Ctrl+\</code> untuk membagi editor aktif menjadi dua.</li>
  <li>Buka ke sisi <code>Ctrl+Enter</code> dari menu konteks Explorer pada file.</li>
  <li>Klik tombol <strong>Split Editor</strong> di kanan atas Editor.</li>
  <li>Drag dan drop file ke sisi mana pun dari wilayah editor.</li>
  <li><code>Ctrl+Enter</code> (macOS: <code>Cmd+Enter</code>) dalam daftar file <strong>Quick Open</strong> (<code>Ctrl+P</code>).</li>
</ul><br>

## Minimap
<p align="justify">
Minimap memberi kita high-level everview tentang source code yang kita miliki, yang berguna untu kquick navigation dan pemahaman kode. Minimap file ditampilkan di sisi kanan editor, kita dapat mengklik atau menyeret area yang diarsir untuk melompat dengan cepat ke berbagai bagian file kita.</p>
