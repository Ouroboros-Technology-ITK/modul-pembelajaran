# Type of testing

## Manual testing
Sesuai dengan namanya, manual, berarti proses pengujian dilakukan dengan cara manual. Proses pengujian ini dilakukan oleh manusia dengan cara memasukkan *input* dan memeriksa *output* secara manual. Proses ini dilakukan secara berulang-ulang sampai mendapatkan hasil yang diinginkan sesuai dengan kebutuhan dari prasyarat.

*Testing* ini masih sering digunakan oleh *developer* dengan melakukan *testing* manual kode program dalam skala kecil. testing ini memiliki kelemahan, yaitu tidak dapat dilakukan secara otomatis. Jika terdapat perubahan pada kode program, maka proses *testing* harus dilakukan ulang.

*Manual testing* akan menjadi sangat sulit untuk pengujian dalam skala *enterprise*. Karena memerlukan banyak waktu dan biaya untuk melakukan *testing*.

## Unit testing
*Unit testing* adalah pengujian yang dilakukan secara otomatis dengan berfokus pada pengujian unit yang terkecil pada desain aplikasi. Karena dalam sebuah aplikasi banyak memiliki unit-unit kecil, maka untuk mengujinya biasanya dibuat program kecil atau main program untuk menguji unit-unit aplikasi.

Unit-unit kecil ini dapat berupa fitur atau fungsi dan pengujian unit biasanya dilakukan saat kode program dibuat.

*Unit testing* umumnya dilakukan oleh *developer*, dengan membuat *code testing* dalam bahasa pemrograman yang sama dengan bahasa pemrograman di aplikasi.


## Integration testing
*Integration testing* juga merupakan pengujian yang dilakukan secara otomatis dengan pengujian secara integrasi. Integrasi di sini adalah pengujian dari hasil penggabungan unit-unit sebagai suatu kombinasi, bukan lagi sebagai suatu unit yang individual.

*Integration testing* sebaiknya dilakukan secara bertahap untuk menghindari kesulitan penelusuran jika terjadi kesalahan *error* / *bug*.

*Testing* ini bisa dilakukan oleh *developer* atau *tester* dengan membuat *code testing* (pada tingkat modul / *package*) dari sekumpulan unit-unit yang saling terintegrasi (*testing* pada tingkat modul / *package* / *library*).


## End to end testing
*End to end testing*, sering disingkat E2E, adalah pengujian yang dilakukan untuk memvalidasi proses kerja aplikasi dari sudut pandang pengguna / user.
*Testing* ini dibuat secara otomatis dengan membuat *code testing* yang dapat mensimulasikan tingkah laku atau tindakan-tindakan dari pengguna, hampir mirip seperti bot.
E2E testing bisa dilakukan olah *Developer* atau *tester*, tapi lebih sering dilakukan oleh tim *testing* khusus yang disebut *software engineer in test* (SEIT).

Contoh, ketika kita membuat aplikasi pendaftaran pengguna, maka kita bisa membuat code E2E *testing* yang dapat mewakili tingkah laku pendaftaran pengguna di aplikasi.

## Acceptance testing
*Acceptance Testing* adalah tahap *final* dari proses pengujian aplikasi / *software* sebelum akhirnya dirilis atau diperbarui secara komersil. *Testing* ini hampir sama dengan E2E *testing*, namun dilakukan oleh pengguna akhir.

Proses ini dapat menentukan apakah suatu fitur dapat diterima atau harus diperbaiki oleh developer / software engineer sebelum dirilis.

Jika terdapat fitur yang tidak lolos pada tahap acceptance testing, maka artinya aplikasi / *software* tersebut tidak memenuhi spesifikasi meskipun banyak beberapa fitur lain yang sudah memenuhi spesifikasi karena akan mengganggu proses aplikasi jika nanti dirilis.


## Software testing pyramid
*Testing pyramid* adalah konsep untuk membangun aplikasi / *software* yang berkualitas dengan cara meminimalkan waktu yang dibutuhkan selama pengembangan dan meminimalkan kerusakan dengan membuat alur *testing* yang dibutuhkan.

![index](https://www.compass-testservices.com/wp-content/uploads/2020/11/Test-pyramid-1-min.png)
Berdasarkan gambar di atas, kita paham bahwa semakin ke atas, maka tes yang dibuat akan semakin mahal dan semakin susah dibuat. Semakin ke bawah, tes akan dibuat semakin cepat dan handal, namun testing yang dibuat harus semakin banyak.
Dengan memperbanyak tes otomatis dalam skala kecil, maka akan menghasilkan aplikasi / *software* berkualitas tinggi secara efisien.
