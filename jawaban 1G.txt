ERD yang di berikan di soal belum optimal karena ERD tersebut belum di normalisasi.
tabel `merchants` memiliki relasi many to many ke tabel `transactions` 
sama seperti tabel `outlets` dengan tabel `transactions`

dengan normalisasi maka akan menghasilkan tabel hasil relasi many to many tersebut
2. `outletstransactions` relasi tabel `outlets` dan `transactions`
3. `transactionsreport` relasi tabel `merchants` dan `outlets` dan `transactions`

pro :
1. merchant id di transactions akan memudahkan query report

cons :
1. data bisa jadi tidak konsisten karena akan ada kemungkinan data yang salah di transactions 
