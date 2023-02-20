# Introduction
<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/Linux-Imager.jpeg"> <h5 align="center">Source: https://id.wikipedia.org/wiki/Kernel_Linux</h5><br>

<p align="justify">
Pada pembahasan materi terminal akan dibutuhkan sistem operasi berbasis Linux. Tapi jangan khawatir, untuk pengguna Windows kita bisa menginstall WSL (<em>Windows Subsystem for Linux</em>) agar kita bisa menggunakan terminal berbasis Linux meskipun menggunakan sistem operasi Windows.<br>

## What Is WSL ?
<p align="justify">
WSL adalah fitur opsional Windows yang memungkinkan program Linux beejalan secara native di Windows. Hal yang dapat kita lakukan dengan menggunakna WSL adalah sebagai berikut: </p>
<ul style="list-style-type:circle;" style="text-align:justify">
  <li>Memilih distribusi GNU/Linux dari Microsoft Store.</li>
  <li>Menjalankan command line tools yang umum seperti <code>ls -la</code>, <code>touch</code>, <code>grap</code> atau perintah Linux lainnya.</li>
  <li>Menjalankan aplikasi terminal base GNU/Linux termasuk tools text editor: <code>vim</code> dan <code>nano</code>.</li>
  <li>Menginstall package manager tambahan menggunakan distribusi GNU/Linux.</li>
</ul>

<p align="center">
    <a href="https://youtu.be/48k317kOxqg" target="_blank"><img src="https://img.youtube.com/vi/48k317kOxqg/0.jpg"></a> 
    <h5 align="center">Source: What Can I Do With WSL? | One Dev Question</h5>
<p><br>

## What Is WSL 2?
<p align="justify">
WSL 2 adalah versi terbaru dari WSL yang bertujuan untuk menambahkan kompabilitas pemanggilan sistem. Perbandingannya dengan WSL 1 dapat kita lihat pada halaman resminya di <a href="https://learn.microsoft.com/en-us/windows/wsl/" target="_blank"><strong>learn.microsoft.com.</strong></a>

<p align="center">
    <a href="https://www.youtube.com/watch?v=MrZolfGm8Zk" target="_blank"><img src="https://img.youtube.com/vi/MrZolfGm8Zk/0.jpg"></a> 
    <h5 align="center">Source: WSL2: Code Faster On The Windows Subsystem For Linux! | Tabs vs Space</h5>
<p><br>

## Prerequisites
<p align="justify">
Untuk update ke WSL 2, kita harus menggunakan <strong>Windows 10</strong>. </p>
<ul style="list-style-type:circle;" style="text-align:justify">
  <li>Untuk sistem x64: Versi 1903 atau yang lebih baru, dengan Build 18362 atau yang lebih bagus.</li>
  <li>Untuk sistem ARM64: Versi 2004 atau yang lebih baru, dengan Build 19041 atau yang lebih baru.</li>
</ul>
<p align="justify">
atau menggunakan <strong>Windows 11</strong>.</p>

> <strong>info: Build yang lebih rendah dari Build 18362 tidak mendukung WSL 2.</strong>

<p align="justify">
Untuk memeriksa versi dan nomor Build kita, tekan tombol Windows + R, lalu ketik <code>winver</code>, pilih OK.

<p align="center">
<img height="300rm" align="center" src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/winver.png"><br>

## Step By Step Installing WSL
<p align="justify">
UNtuk menginstall WSL secara simple, sebenarnya bisa menggunakan command <code>wsl --install</code>. Namun disini kita akan mencoba menginstall secara manual agar kita dapat mengetahui apa saja yang harus dilakukan untuk menjalankan WSL.<br>

- <strong>Langkah 1</strong> - Aktifkan Windows Subsystem for Linux
<p align="justify">
Buka PowerShell sebagai Administrator (tekan Windows > ketik "PowerShell" > klik kanan > pilih Run As Administrator) dan masukkan perintah ini: 

```
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
```
<br>

- <strong>Langkah 2</strong> - Aktifkan Virtual Machine Feature
<p align="justify">
Buka PowerShell sebagai Administrator dan jalankan: 

```
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```

Restart komputer untuk menyelesaikan penginstalan dan pembaruan ke WSL 2.
<br>

- <strong>Langkah 3</strong> - Download Linux kernel update package
  - Download latest package: <a href="https://learn.microsoft.com/en-us/windows/wsl/install-manual#step-4---download-the-linux-kernel-update-package" target="_blank"><strong>WSl 2 Linux Kernel Update Package</strong></a>
  - Run update package yang telah di download. (Double click untuk run, maka kita akan dimintai permissions, pilih Yes untuk menyetujui penginstalan)
<br><br>

- <strong>Langkah 4</strong> - Set WSL 2 sebagai default version
<p align="justify">
Buka PowerShell dan jalankan perintah ini: 

```
wsl --set-default-version 2
```
<br>

- <strong>Langkah 5</strong> - Install Linux distribution
  - Buka Microsoft Store dan search Linux distributin.
  - Kita akan membutuhkan <a href="https://www.microsoft.com/store/productId/9PN20MSR04DW" target="_blank"><strong>Ubuntu 22.04 LTS</strong></a>.
  - Dari halaman distribution, klik tombol Get. <br>
    <p align="justify">
    Pertama kali menjalankan Linux distribution yang baru diinstall, kita akan diminta untuk menunggu satu atau dua menit agar fil mendekompres dan disimpan di PC kita. Kemudian kita perlu membuat username dan password untuk Linus distribution.<br>
  
## How To Use WSL 2
<p align="justify">
Pada distributin Ubuntu yang baru saja diinstall, kita akan disambut dengan terminal Linux. COba jalankan beberapa perintah Linux:</p>

- <strong>Where Am I?</strong>
<p align="justify">
Gunakan perintah dasar Linux <code>pwd</code> untuk menemukan path direktori kerja kita saat ini. <code>pwd</code> adalah kependekan dari "<em>print working directory</em>". Kita hanya perlu memasukkan <code>pwd</code>, path saat ini akan ditampilkan penuh, yaitu path semua direktori yang diawali dengan garis miring (/). Misalnya, /home/username.<br>

  ```
  $ pwd
  /Users/onyxfish
  ```
<br>

- <strong>Going There</strong>
<p align="justify">
Gunakan perintah dasar Linux <code>cd</code> untuk menjelajahi file dan direktori Linux. Perintah Linux ini memerlukan path penuh atau nama direktori, tergantung pada direktori kerja kita saat ini. <code>cd</code> adalah kependekan dari "<em>change directory</em>". Misalnya saat ini kita sedang ingin beralih ke direktori baru, misalnya /Users/onyxfish/Documents, maka kita perlu mengetikkan cd diikuti path absolut direktori tersebut:

  ```
  $ cd /Users/onyxfish/Documents
  ```
<br>

- <strong>What's Here?</strong>
<p align="justify">
Gunakan perintah dasar Linux <code>ls</code> untuk melihat file dan direktori pada sistem. Menjalankannya tanpa flag atau parameter akan menampilkan konten direktori kerja saat ini.</p/>

  - Untuk melihat file apa yang ada di dalam direktori kerja saat ini, gunakan perintah <code>ls</code>: 

```
$ ls
Applications			bin
Confidential			confidential.py
Desktop				    confidential2.tc
Documents			    external_lbrs
Downloads			    feed.json
Dropbox				    feed.py
Library				    floobits
Movies				    gitconfig
Music				    google_analytics_auth.dat
Pictures			    pgadmin.log
Public				    s3cfg
SpiderOak Hive			src
VirtualBox VMs			syria
adt-bundle-mac			tmp
android-sdk-macosx
```

  - Untuk melihat file apa yang ada di direktori yang berbeda:

```
$ ls Documents
Arduino				android-workspace
Aspyr				eagle
GoPro Projects		gopromote.prm
Hero Lab			nicar_expenses.pdf
Klei				nicar_receipts.pdf
Library				orchard bank travel number.txt
Logan.game			pedagogy.txt
Logans.band			piano 1.band
MapBox				test.txt
NACIS notes.txt		tyler-env.txt
Roblox
```
  
  - Secara default, <code>ls</code> tidak akan menampilkan file yang namanya diberi tanda titik. Ini disebut file tersembunyi biasanya terkait dengan program konfigurasi. Namun, sangat berguna untuk melihat file-file ini, yang dapat kalian lakukan dengan flag <code>-a</code>:
  
```
$ ls -a
.				        .pypirc
..				        .python-eggs
.CFUserTextEncoding		.qgis
.DS_Store			    .qgis2
.MacOSX			    	.relay.conf
.NERDTreeBookmarks		.rnd
.Platformer		       	.rstudio-desktop
.Rhistory			    .s3cfg
.Rube Goldberg			.spyder2
.Trash			    	.sqlite_history
.Xauthority		    	.ssh
.android		    	.subversion
.anyconnect		    	.teamocil
.bash_history			.tilemill
.bash_profile			.tmux.conf
.build			    	.tox
.cache			    	.uibtedbn
.clan_auth.dat			.vagrant.d
.clan_secrets.json		.vim
.config			    	.viminfo
.cordova		    	.viminfo.tmp
.cups			    	.viminfz.tmp
.distlib		    	.vimrc
.dropbox		    	.virtualenvs
.dropbox-master			.wireshark
.eaglerc		    	.wireshark-etc
.ec2			    	.ypp_42
.floorc		    		.zcompdump
.fontconfig		    	.zcompdump-nomad-5.0.2
.gem			    	.zprofile
.gitconfig		    	.zsh-update
.gitsh_history			.zsh_history
.gnome2			    	.zshrc
.gnupg			    	Applications
.goaccessrc		    	Confidential
.godot			    	Desktop
.haxelib		    	Documents
.heroku			    	Downloads
.hkzftsorc		    	Dropbox
.hxcpp_config.xml		Library
.hxcpp_config.xml.bak	Movies
.ievms			      	Music
.infinit		    	Pictures
.inkscape-etc			Public
.ipython		    	SpiderOak Hive
.keybase		    	VirtualBox VMs
.keybase-installer		adt-bundle-mac
.lesshst		    	android-sdk-macosx
.lighttable		    	bin
.local			    	confidential.py
.matplotlib		    	confidential2.tc
.mongorc.js		    	external_lbrs
.netrc			    	feed.json
.ngrok			    	feed.py
.node-gyp		    	floobits
.npm			    	gitconfig
.oh-my-zsh		    	google_analytics_auth.dat
.pgadmin_histoqueries	pgadmin.log
.pgpass			    	s3cfg
.pip			    	src
.psql_history			syria
.pylint.d		    	tmp
```
<br>

- <strong>Touch Command</strong>4
<p align="justify">
Gunakan perintah dasar Linux <code>touch</code> untuk perintah dasar Linux yang memungkinkan kalian membuat file baru yang kosong dan mengubah timestamp di command line Linux. Contohnya, masukkan perintah berikut untuk membuat file HTML bernama Web di direktori Documents:</p/>

```
$ touch /home/username/Documents/Web.html
```

- <strong>Update</strong>

```
$ sudo apt update
```

<div align="justify">
    <!-- Prev Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering/1.%20Introduction/2.%20Day%20To%20Day" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/left%20(1).png" align="left" height="50" width="50"></a>
    <!-- Next Page -->
    <a href="https://github.com/Ouroboros-Tech/modul-pembelajaran/tree/main/3.%20Software%20Engineering" target="_blank"><img src="https://github.com/Ouroboros-Tech/modul-pembelajaran/blob/main/image/home%20(2).png" align="right" height="50" width="50"></a>
<div>
