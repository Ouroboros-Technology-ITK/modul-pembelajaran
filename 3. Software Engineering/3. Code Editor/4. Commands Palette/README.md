# Commands Palette
<p align="justify">
Command memicu serangkaian action dalam VS Code. Jika kita pernah mengonfigurasi suatu keybinding, maka dapat dibilang kita telah menggunakan sebuag command. Command juga digunakan oelh extensions untuk mengekspos fungsionalitas kepada pengguna, bind action di UI VS Code, dan mengimplementasikan internal logic. Terdapat beberapa command yang biasa digunakan dan dapat meningkatkan produktivitas kita sebagai programmer.<br><br>

## Show All Commands
<p align="justify">
Dengan bantuan command pallete, kita dapat menampilkan sebuah commands dan mengakses commands kita sesuai dengan konteks kita saat ini. Kita hanya perlu mengetikkan keyword yang related dengan commands dan menemukannya. Disini kita juga dapat menemukan keyword yang relevan dengan setiap commands.<br>

<p align="justify">
Terdapat 2 cara dalam mengakses seluruh commands. Yang pertama yaitu dengan menggunakan <code>Help > Show All Commands</code>.<br>

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/command-pallete-1.png"><br>

<p align="justify">
Atau kita juga dapat menggunakan shortcut <code>Ctrl+Shift+P</code>, jika di mac menggunakan shorcut <code>Command+Shift+P</code>.<br><br>

## Clear Command Palette History
<p align="justify">
Kita dapat menggunakan command palette untuk menghapus command palette history dengan mengikuti petunjuk di bawah ini:<br>

<ol style="list-style-type:circle;" style="text-align:justify">
  <li>Buka VS Code dan tekan <code>Ctrl+Shift+P</code> untuk membuka command palette.</li>
  <li>Ketik <code>Clear</code> dan pilih <strong>Clear Command History</strong>. Maka dengan begitu command palette history yang kita miliki sekarang kosong.</li>
</ol><br>

## Change Command Paletta History Settings
<p align="justify">
VS Code memungkinan kita untuk membatasi berapa banyak recent command yang akan disimpan atau dinonaktifkan sepenuhnya:<br>

<ol style="list-style-type:circle;" style="text-align:justify">
  <li>Buka VS Code dan tekan <code>Ctrl+Shift+P</code> untuk membuka command palette.</li>
  <li>Ketik <code>settings</code> dan pilih <strong>Preferences: Open Settings (UI)</strong>.</li>
  <li>Cari <code>palette</code>. Setelah melihat <strong>Workbench > Command Palette: History</strong>, kita dapat mengaturnya ke angka lain dari command history yang ingin VS Code simpan. Atau kita dapat mengatur nilai ke 0 untuk menonaktifkan <strong>Command Palette History</strong>.</li>
    <li>Kita juga dapat mencentang kotak <strong>Workbench > Command Palette: Preserve Input</strong> di pengaturan untuk mengisi command palette dengan perintah yan gterakhri digunakan. Fitur ini menghemat waktu, terutama ketika ktia harus menggunakan satu perintah secara berulang-ulang.</li>
</ol><br>
