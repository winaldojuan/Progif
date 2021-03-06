# Progif

Dokumentasi Penggunaan atau pengujian dari webservice pengaturan data buku perpustakaan:
1. Konek ke jaringan ITB
2. Buka browser, dan masukkan alamat http://167.205.67.224:8181/home/
3. Setelah masuk, akan terlihat 4 tombol, yaitu GET, POST, UPDATE, dan DELETE
4. Pengujian akan dilakukan menurut modul yang telah dibuat

Pengujian GET Request:
1. Klik tombol "GET"
2. Akan terlihat 6 tombol, yaitu getall, search by judul, search by penulis, search by tahun terbit, search by penerbit, dan home

*NB : untuk setiap tombol, diklik, jangan dengan "enter"

**Modul Get All**
1. Pertama, klik tombol Get All. Setelah itu akan masuk ke halaman yang memiliki 1 tombol
2. Klik tombol yang ada di halaman tersebut
3. Webservice akan memberi response dari GET Request yang dikirim saat tombol tersebut ditekan
4. Responnya adalah data buku dalam json di tabel "buku" yang memiliki 1000 data didalamnya
5. Klik tombol "back" untuk kembali ke halaman get

**Modul Get by Judul**
1. Klik tombol "Search by Judul".
2. Setelah itu akan masuk ke halaman yang memiliki 1 field dan masukkan kata "this" sebagai pengujian
3. Klik tombol search
4. Webservice akan memberi response dari GET Request yang dikirim saat tombol tersebut ditekan
5. Responnya adalah data buku dalam json di tabel "buku" yang mengandung kata "this"
6. Bisa dicoba dengan keyword lainnya
7. Klik tombol "back" untuk kembali ke halaman get

**Modul Get by Penulis**
1. Klik tombol "Search by Penulis".
2. Setelah itu akan masuk ke halaman yang memiliki 1 field dan masukkan kata "Mark" sebagai pengujian
3. Klik tombol search
4. Webservice akan memberi response dari GET Request yang dikirim saat tombol tersebut ditekan
5. Responnya adalah data buku dalam json di tabel "buku" yang mengandung kata "Mark"
6. Bisa dicoba dengan keyword lainnya
7. Klik tombol "back" untuk kembali ke halaman get

**Modul Get by Tahun Terbit**
1. Klik tombol "Search by Tahun Terbit".
2. Setelah itu akan masuk ke halaman yang memiliki 1 field dan masukkan angka 2001 sebagai pengujian
3. Klik tombol search
4. Webservice akan memberi response dari GET Request yang dikirim saat tombol tersebut ditekan
5. Responnya adalah data buku dalam json di tabel "buku" yang diterbitkan pada tahun 2001
6. Bisa dicoba dengan keyword lainnya
7. Klik tombol "back" untuk kembali ke halaman get

**Modul Get by Penerbit**
1. Klik tombol "Search by Penerbit".
2. Setelah itu akan masuk ke halaman yang memiliki 1 field dan masukkan kata "Random" sebagai pengujian
3. Klik tombol search
4. Webservice akan memberi response dari GET Request yang dikirim saat tombol tersebut ditekan
5. Responnya adalah data buku dalam json di tabel "buku" yang mengandung kata "Random"
6. Bisa dicoba dengan keyword lainnya
7. Klik tombol "back" untuk kembali ke halaman get
8. Klik tombol "home" untuk kembali ke halaman utama

**Modul POST Request**
1. Klik tombol "POST" di halaman utama
2. Akan terlihat form dengan field ID, Judul, Penulis, Tahun Terbit, dan Penerbit
3. Masukkan data yang ingin dimasukkan, mulai dari ID = 1
4. Masukkan string bebas ke judul, penulis dan penerbit.
5. Masukkan integer bebas ke Tahun Penerbit
6. Klik tombol "POST"
7. Ketika "POST" diklik, POST Request dikirimkan beserta file json, yaitu hasil encode dari data yang dimasukkan
8. Webservice akan memasukkan hasil json yang diterima ke tabel "test"
9. Tabel yang dipakai berbeda untuk memudahkan pengujian
10. Masukkan data lain, misal ID = 2 beserta data lainnya
11. Data di tabel "test" akan ditampilkan ketika tombol "POST" diklik
12. Klik tombol "home" untuk kembali ke halaman utama

**Modul POST Request**
1. Klik tombol "POST" di halaman utama
2. Akan terlihat form dengan field ID, Judul, Penulis, Tahun Terbit, dan Penerbit
3. Masukkan data yang ingin dimasukkan, mulai dari ID = 1
4. Masukkan string bebas ke judul, penulis dan penerbit.
5. Masukkan integer bebas ke Tahun Penerbit
6. Klik tombol "POST"
7. Ketika "POST" diklik, POST Request dikirimkan beserta file json, yaitu hasil encode dari data yang dimasukkan
8. Webservice akan memasukkan hasil json yang diterima ke tabel "test"
9. Tabel yang dipakai berbeda untuk memudahkan pengujian
10. Masukkan data lain, misal ID = 2 beserta data lainnya (disarankan agar delete dapat diuji nantinya)
11. Data di tabel "test" akan ditampilkan ketika tombol "POST" diklik
12. Klik tombol "home" untuk kembali ke halaman utama

**Modul PUT Request**
1. Klik tombol "UPDATE" di halaman utama
2. Akan terlihat form dengan field ID, masukkan ID yang dimasukkan sebelumnya pada tahap POST Request sebagai pengujian
3. Klik tombol pada halaman tersebut
4. Akan terlihat form dengan field ID, Judul, Penulis, Tahun Terbit, dan Penerbit
5. Seluruh data dapat diubah, termasuk ID.
6. Data yang diubah adalah data buku yang ID nya dimasukkan pada form sebelumnya
7. Masukkan string bebas ke judul, penulis dan penerbit.
8. Masukkan integer bebas ke Tahun Penerbit
9. Klik tombol "POST"
10. Ketika "POST" diklik, PUT Request dikirimkan beserta file json, yaitu hasil encode dari data yang dimasukkan
11. Webservice akan memasukkan hasil json yang diterima ke tabel "test"
12. Tabel yang dipakai berbeda untuk memudahkan pengujian
13. Coba untuk mengganti data lainnya
14. Data di tabel "test" akan ditampilkan ketika tombol "POST" diklik
15. Klik tombol "home" untuk kembali ke halaman utama

**Modul DELETE Request**
1. Klik tombol "DELETE" di halaman utama
2. Akan terlihat form dengan field ID
3. Masukkan ID data buku yang ingin dihapus, misalnya 1
4. Klik tombol "delete"
5. Ketika "delete" diklik, DELETE Request dikirimkan ke webservice
6. Webservice akan delete data
7. Data di tabel "test" akan ditampilkan ketika tombol "delete" diklik
8. Apabila ketika delete diklik tidak menampilkan apapun, artinya data di tabel hanya 1 dan telah didelete sesuai permintaan pengguna
9. Klik tombol "home" untuk kembali ke halaman utama
