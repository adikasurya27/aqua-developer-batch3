Relational Database

Relational Database adalah kumpulan item data dengan hubungan yang telah ditentukan sebelumnya. Berbagai item ini disusun menjadi satu set tabel dengan kolom dan baris. Tabel digunakan untuk menyimpan informasi tentang objek yang akan direpresentasikan dalam database. Tiap kolom pada tabel memuat jenis data tertentu dan bidang menyimpan nilai aktual atribut. Baris pada tabel merepresentasikan kumpulan nilai terkait dari satu objek atau entitas. Tiap baris pada tabel dapat ditandai dengan pengidentifikasi unik yang disebut kunci utama, dan baris di antara beberapa tabel dapat dibuat saling terkait menggunakan kunci asing. Data ini dapat diakses dengan berbagai cara tanpa menyusun ulang tabel basis data itu sendiri.

- DDL (Data Defiiton Language) : untuk membuat dan memodifikasi objek database seperti tabel, indeks, dan pengguna. 
1. Create
Query   : create <database_name>;
fungsi  : perintah create merupakan sebuah perintah yang bisa digunakan ketika membuat sebuah database yang baru, baik itu berupa tabel baru atau sebuah kolom baru.

2. Alter
Query   : alter table <table_name> rename to <table_name_change>;
Fungsi  : untuk melakukan perubahan struktur tabel yang telah dibuat.

3. Drop
Query   : drop table <table_name>;
Fungsi  : digunakan dalam menghapus baik itu berupa database, table maupun kolom hingga index.

- DML (Data Manipulation Language) : bahasa pemrograman komputer yang digunakan untuk menambah, menghapus, dan memodifikasi data dalam database.
1. Insert
Query   : insert into public.<database_name> (name, age) values ('dika', 21);
Fungsi  : Perintah insert digunakan untuk memasukkan sebuah record baru di dalam sebuah tabel database.

2. Update
Query   : update public.<table_name> set age = 25 where id = 1;
Fungsi  : Update digunakan untuk mengubah data, atau memodifikasi data yang terdapat didalam tabel.

3. Delete
Query   : delete from public.<table_name> where id = 1;
Fungsi  : Perintah delete digunakan untuk menghapus data atau record di dalam table.

- DCL (Definition Control Language) :  Untuk memberi akses database ke user tertentu.

- Join : perintah dalam SQL yang berfungsi untuk menggabungkan informasi dari dua tabel. Hasilnya adalah seperangkat kolom yang digabung dari kedua tabel. Fungsi JOIN sendiri memiliki terdiri dari beberapa jenis perintah seperti INNER JOIN dan OUTER JOIN (Left Join, Right Join dan Full Join).
1. Left Join    : membuat sebuah parameter dari table sebelah kiri, dan jika ada data atau record yang kosong atau tidak berelasi maka akan berisi NULL di sebelah kanan.
2. Right Join   : membuat sebuah parameter pada sebelah kanan jika data pada table terdapat data atau record yang kosong atau tidak berelasi maka akan berisi NULL.
3. Full Join    : membuat sebuah parameter pada kanan dan kiri.

- Aggregation
"Database Aggregation" adalah istilah yang cukup umum yang dapat merujuk pada beberapa hal, tetapi umumnya, ini adalah proses di mana data dikumpulkan dan diekspresikan dalam bentuk ringkasan.
Agregasi dapat mengambil banyak bentuk; misalnya, data mentah dapat dikumpulkan selama periode waktu tertentu untuk memberikan statistik seperti rata-rata, minimum, maksimum, jumlah, dan hitungan. Data gabungan berguna karena Anda dapat menganalisisnya untuk mendapatkan wawasan tentang sumber daya tertentu atau bahkan kelompok sumber daya.
Contoh fungsi aggregation meliputi:
1. Count()
2. Sum()
3. Avg()
4. Min()
5. Max()

- Subquery
Fungsinya untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil.
Query   : update products set stock = subquery.stock - 2 from(select id, stock from products where id = 1) as subquery where products.id = 1

- Function
Suatu sub program yang diperuntukan untuk mengerjakan suatu perintah tertentu sesuai dengan fungsi method itu sendiri. Dengan menggunakan fungsi tentu saja pekerjaan untuk mengelolah dan analisis data menjadi lebih mudah dan akurat. Penggunaanya pun tidaklah sulit bagi pemula perlu untuk menguasai function di mysql ini, Sebagai dasar untuk mengembangkan aplikasi web menggunakan database mysql.