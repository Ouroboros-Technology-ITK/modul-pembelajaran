# Extension - Most Common Use VS Code Extension
Seperti yang kita ketahui bahwa VS Code memiliki marketplace-nya sendiri. Ada beberapa ekstensi yang sering digunakan dan disarankan untuk menginstal demi membantu proses pembelajaran kedepannya. Berikut beberapa ekstensinya, antara lain:<br>

- <strong>Git Lens</strong>
<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br>

<p align="justify">
GitLens adalah sebuah ektensi pada VS Code yang dapat kita gunakan untuk visualisasi dan melihat history/riwayat git commit pada kode dalam sebuah proyek. Keuntungan jika menggunakan GitLens akan memudahkan kita dan tim untuk melihat siapa yang merubah kode setiap baris dalam proyek, misal kita menemukan bugs, kita dan tim akan segera bisa menangani dan memperbaiki bugs yang ada karena ter tracking dengan jelas siapa yang terakhir mengubah kode tersebut.<br>

<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br><br>

- <strong>Git History</strong>
<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br>

<p align="justify">
Menggunakan Git History kita dapat mengakses git log dengan grafik dan detail. Seperti namanya, kita dapat melihat dan mencari history dan kita juga dapat membandingkan branches, commits, dan file di seluruh commits bersama dengan banyak fitur lain-lain.<br>

<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br>

<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br><br>

- <strong>Prettier</strong>
<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br>

<p align="justify">
Dengan extension ini, kita dapat merapikan kode secara cepat tanpa perlu diubah satu per satu. Prettier akan merapikan kode kita berdasarkan setelan default-nya atau kita dapat mengatur sendiri panjang baris maksimum atau membungkus kode bila perlu. Extension ini mendukung bahasa pemrograman JavaScript, TypeScript, Flow, JSX, JSON, CSS, SCSS, Less, HTML, Vue, Angular, GraphQL, Markdown, dan YAML.<br>

<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br><br>

- <strong>REST Client</strong>
<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br>

<p align="justify">
REST Client memiliki nama yang paling jelas untuk tool dan VS Code marketplace description meringkas persis apa yang dilakukannya: "REST Client memungkinkan kita untuk mengirim permintaan HTTP dan melihat respons dalam Visual Studio Code secara langsung." Sesederhana itu, pada dasarnya, REST Client adalah HTTP tools yang ada di dalam VS Code.<br><br>

- <strong>Live Server</strong>
<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br>

<p align="justify">
Live Server adalah salah satu plugin di VS Code yang dapat digunakan untuk melihat secara langsung hasil dari halaman HTML selama kita melakukan development.<br>

<p align="justify">
Ekstensi satu ini memungkinkan kita untuk melakukan live reload aplikasi pada saat proses development. Ketika kita melakukan perubahan pada IDE maka secara otomatis kita bisa melihat perubahannya tanpa perlu melakukan reload manual.<br>

<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br><br>

- <strong>Live Share</strong>
<p align="center">
<img height="200rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/ajax-snippet.gif"><br>

## What Is Live Share?
<p align="justify">
Live share adalah plugin di dalam VS Code yang memungkinkan kita membuat suatu proyek ramai-ramai di VS Code. Misal kita mengalami sedikit masalah dengan code yang sedang kita buat dan membutuhkan bantuan dari rekan satu tim untuk membantu kita menyeleseikan masalah yang dihadapi langsung pada VS Code kita secara online.<br>

## Why We Need Live Share?
Ekstensi VSCode Live Share memungkinkan kita dan orang lain mengedit file yang sama, men-debug aplikasi yang sama, dan bahkan menjalankan perintah di console yang sama secara bersamaan.<br>

Kita dan rekan tim kita harus menginstal ekstensi Live Share terlebih dahulu untuk mulai berkolaborasi. Setelah itu, kita harus melakukan Sign In. Kita dapat memilih untuk melakukan Sign In dengan akun Microsoft atau GitHub. Untuk melakukan Sign In, gunakan tombol Sign In yang terdapat pada status bar bawah dengan ikon orang.<br>

## How To Open A Session
<p align="justify">
Di host machine, mulai sesi kolaborasi dengan cara berikut:<br>

- Tekan CTRL+SHIFT+P di VSCode untuk membuka command palette.
- Masuk ke Start Collaboration Session dan pilih liveshare.start dari daftar.
- Klik OK ketika kita mendapatkan pesan pop-up untuk menerima port yang tidak diblokir di firewall kita. Dengan melakukan hal tersebut, memungkinkan VSCode untuk berinteraksi dengan guest machine secara real-time.

<p align="justify">
Secara default, kedua participants dalam sesi kolaborasi akan memiliki hak editor dan dapat membuat perubahan pada file yang sama.<br>

- Selanjutnya, atur izin akses kolaborator kita dan generate invitation link ke sesi kolaborasi dengan cara: 
- Klik Make read-only mode di kanan bawah, untuk membatasi kolaborator kita dengan akses read- only.
- Klik Copy again untuk menyalin invitation link ke clipboard kita.
- Kirim invitation link ke guest. Kita dapat menggunakan email atau chat tool seperti Slack atau Discord.

## How To Join A Session
<p align="justify">
Setelah kita mendapatkan invitation link kita dapat membuka link tersebut dari web browser untuk bergabung dalam sesi kolaborasi. Atau kita juga dapat melakukannya dengan klik File > Join Collaboration Session. lalu paste di invitation link yang kita dapat dan tekan Enter untuk mengonfirmasi.<br>
