note : valdasi kelas kosong pada jam tersebut

1. START
2. LOGIN -> email,password
3. CHECK ROLE -> admin,dosen,mahasiswa
4. 
if admin {
  
    1. Crud Jadwal
    2. Crud Matakuliah
    3. Crud Akun
    4. Crud Kelas 
}

if dosen {
    1. View Jadwal Mengajar Saya
       -> Menampilkan list jadwal kuliah yang diampu oleh dosen tersebut.
       -> Fitur: 
          - Urutkan Jadwal (Selection Sort / Insertion Sort) berdasarkan jam mulai.
          - Cari Jadwal (Sequential Search / Binary Search) berdasarkan nama Matakuliah.
          
    2. Statistik Mengajar (Spesifikasi e)
       -> Menampilkan total jam mengajar dosen dalam 1 / bulan (berdasarkan SKS/durasi).
       -> Menampilkan jumlah kelas yang tersisa/harus diajar pada hari berjalan.
       
    3. Lihat Daftar Mahasiswa per Kelas
       -> Menampilkan daftar mahasiswa yang mengambil kelas/matakuliah yang diampu dosen tersebut (memanfaatkan data "Crud Kelas/Mahasiswa" dari Admin).
       
    4. Profil Akun Dosen
       -> Menampilkan info pribadi (NIM, Nama , Jurusan) dan opsi untuk ubah password.
}
if mahasiswa {
    dashboard :
    1. View Jadwal Kuliah (Spesifikasi c & d)
       -> Menampilkan daftar seluruh jadwal kuliah atau jadwal khusus kelasnya.
       -> Fitur:
          - Urutkan Kelas (Selection Sort / Insertion Sort) berdasarkan jam mulai paling pagi.
          - Cari Kelas/Dosen (Sequential Search / Binary Search) berdasarkan nama Matakuliah atau nama Dosen.

    2. Cek Sisa Kuliah Hari Ini (Spesifikasi e)
       -> Menampilkan statistik jumlah mata kuliah yang tersisa yang harus dihadiri pada hari berjalan.

    3. Informasi Ruangan & Kapasitas
       -> Menampilkan daftar ruangan (misal: Lab 1, Ruang 302) untuk mengecek lokasi kelas dan melihat kapasitas ruangan agar tidak salah tempat.

    4. Lihat Teman Sekelas
       -> Menampilkan daftar mahasiswa lain yang berada di kelas atau angkatan yang sama (menggunakan data Kelas/Mahasiswa dari Admin).
}
