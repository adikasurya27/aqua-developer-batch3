Version Control System (GIT)

-   Commands
git init <options> : digunakan untuk membuat repository di file lokal.
git add <directory> : perintah yang digunakan untuk menambahkan file baru di repository yang dipilih.
git commit <options> : digunakan untuk menyimpan perubahan yang sudah dilakukan, namun tidak ada perubahan yang terjadi pada remote repository.
git push <options> : digunakan dalam mengirimkan perubahan file yang dilakukan setelah di commit ke remote repository.
git pull <options> <repository> :  mengambil commit terbaru dari remote repository.
git branch <branch> : menambahkan fitur baru atau memperbaiki bug.
git merge <branch_name> : perintah yang digunakan untuk menggabungkan cabang aktif serta cabang yang dipilih.
git status : digunakan untuk mengetahui sebuah status dari sebuah repository lokal.
git fetch  <repository Url> : mengambil commit terbaru dari remote repository
git remote : untuk melihat perbedaan perubahan di revisi.
git stash <command> : untuk menyimpan semua progress yang sudah kamu lakukan sejak commit terakhir tanpa membuat sebuah commit untuk state itu sendiri.

-   Commit Convention
Spesifikasi Komit Konvensional adalah konvensi ringan di atas pesan komit. Ini menyediakan seperangkat aturan yang mudah untuk membuat riwayat komit eksplisit; yang membuatnya lebih mudah untuk menulis alat otomatis di atasnya. Konvensi ini sesuai dengan SemVer, dengan menjelaskan fitur, perbaikan, dan pemutusan perubahan yang dibuat dalam pesan komit.

build : Membangun perubahan terkait (misalnya: npm terkait/menambahkan dependensi eksternal)
chore : Perubahan kode yang tidak akan dilihat oleh pengguna eksternal (misalnya: ubah ke file .gitignore atau file .prettierrc)
feat : Fitur baru
fix : Perbaikan bug
docs : Perubahan terkait dokumentasi
refactor : Kode yang tidak memperbaiki bug atau menambahkan fitur. (misalnya: Anda dapat menggunakan ini ketika ada perubahan semantik seperti mengganti nama variabel/nama fungsi)
perf : Kode yang meningkatkan kinerja
style : Kode yang terkait dengan gaya
test : Menambahkan tes baru atau membuat perubahan pada tes yang ada

-   Semantic Versioning
Semantic Version ini adalah sistem penomoran versi sebuah produk/software.
Format atau cara penulisan Semantic Version ini terdiri atas 3 bagian yang dipisahkan dengan tanda titik: MAJOR.MINOR.PATCH.
Commands v<major>.<minor>.<patch>
Major: Jika terjadi perubahan angka pada Major, berarti ada perubahan yang sangat drastis, ditambahkannya fitur baru termasuk akan ada beberapa API/fitur yang tidak lagi cocok dengan versi sebelumnya.
Minor: Perubahan angka pada Minor berarti ada fitur baru tanpa menghilangkan fitur pada versi sebelumnya.
Patch: Penambahan angka pada bagian patch artinya ada perbaikan bug tanpa menghilangkan  fitur pada versi sebelumnya.

-   Git Manajemen
Trunk Based Development adalah pendorong utama  Continuous Integration dan dengan perluasan Continuous Delivery. Ketika individu dalam tim melakukan perubahan mereka ke trunk beberapa kali sehari, menjadi mudah untuk memenuhi persyaratan inti Continuous Integration bahwa semua anggota tim berkomitmen ke trunk setidaknya sekali setiap 24 jam. Ini memastikan basis kode selalu dapat dirilis sesuai permintaan dan membantu mewujudkan Continuous Delivery.