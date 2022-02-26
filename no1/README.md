## Nomor 1

### Cara menjalankan server
Perhatikan bahwa database yang dganakan adalah MySQL. 

Sebelum menjalankan server kita perlu setup environement mulai dari database, jwt, port dengan membuat file pada forlder __no1__ dengan nama __.env__ dengan skema yang mirip pada file __no1/env.example__.

Untuk menjalankan server dapat di run command berikut:
```
go run main.go
```

### Postman Documentation
Dokumentasi postman ada di folder __no1/documentation__.

### Tugas yang selesai
Tanda ✔ bererti sudah selesai:

A. Membuat fungsi login (5 point) ✔

B. Untuk authorization pada point A gunakan JWT (6 point) ✔

C. Laporan nama merchant, omzet per hari dalam pada bulan november mulai tanggal 1
sampai dengan tanggal 30 dengan pagination. Apabila tidak ada transaksi pada tanggal itu
omzet akan bernilai 0 (6 point) ✔

D. API untuk menampilkan laporan nama merchant, nama outlet, omzet per hari pada bulan
november mulai tanggal 1 sampai dengan tanggal 30 dengan pagination. Apabila tidak ada
transaksi pada tanggal itu omzet akan bernilai 0 (6 point) ✔

E. Pada poin C pastikan user tidak bisa melakukan akses pada merchant_id yang bukan
miliknya (10 point) ✔

F. Pada poin D pastikan user tidak bisa melakukan akses laporan pada outlet_id yang bukan
miliknya (5 point) ✔

G. Dari test case pada point C dan point D, apakah struktur ERD yang dibentuk sudah optimal ✔
? berikan penjelasannya (9 point)
- Untuk menjawab pertanyaan di atas menurut saya harus melihat PRD secara utuh. jik ahanya dari soal saja menurut saya pernu dibuat suatu table yang berisi tanggal sehingga bisa mudah untuk query disetiap tanggalnya.

H. Dokumen teknis Data Manipulation. ✔
- Dokumen ada pada folder __no1/documentation__.