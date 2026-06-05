# 🏛️ Sistem Informasi Manajemen Perkuliahan

Sistem ini dirancang untuk mengelola penjadwalan kuliah, data akademis, serta memberikan informasi statistik perkuliahan bagi **Admin**, **Dosen**, dan **Mahasiswa**. Sistem ini mengimplementasikan algoritma pengurutan (*Sorting*) dan pencarian (*Searching*) untuk penyajian datanya.

---

## 🧭 Alur Kerja Sistem (Flowchart Logic)

1. **START**
2. **LOGIN** -> Input `email` & `password`.
3. **CHECK ROLE** -> Sistem memvalidasi hak akses ke salah satu role berikut:
   
### 👨‍💼 [IF ROLE = ADMIN]
* **1. CRUD Jadwal** * Mengelola pembuatan, pembaruan, dan penghapusan jadwal kuliah.
    * ⚠️ **VALIDASI UTAMA:** Sistem wajib mengecek **validasi ruangan & kelas kosong pada jam tersebut**. Jadwal tidak akan tersimpan jika terjadi bentrok waktu, ruangan, atau kelas di hari yang sama.
* **2. CRUD Mata Kuliah** * Mengelola daftar mata kuliah, bobot SKS, dan penempatan semester.
* **3. CRUD Akun** * Mengelola kredensial login (`email`, `password`, `role`, dan `status`) untuk seluruh pengguna.
* **4. CRUD Kelas** * Mengatur data rombongan belajar (nama kelas, angkatan, kapasitas) beserta plot Dosen Wali.

### 👨‍🏫 [IF ROLE = DOSEN]
* **1. View Jadwal Mengajar Saya** * Menampilkan daftar jadwal kuliah yang diampu oleh dosen yang bersangkutan.
    * **Fitur:**
        * *Urutkan Jadwal:* Menggunakan **Selection Sort / Insertion Sort** berdasarkan jam mulai kuliah.
        * *Cari Jadwal:* Menggunakan **Sequential Search / Binary Search** berdasarkan nama Mata Kuliah.
* **2. Statistik Mengajar** * Menampilkan total jam mengajar dosen dalam 1 bulan (berdasarkan bobot SKS/durasi).
    * Menampilkan jumlah kelas yang tersisa/harus diajar pada hari berjalan.
* **3. Lihat Daftar Mahasiswa per Kelas** * Menampilkan daftar mahasiswa yang mengambil kelas/mata kuliah yang diampu oleh dosen tersebut.
* **4. Profil Akun Dosen** * Menampilkan informasi pribadi (NIDN, Nama, dan Jurusan) serta opsi untuk memperbarui password.

### 👨‍🎓 [IF ROLE = MAHASISWA]
* **1. View Jadwal Kuliah (Dashboard)** * Menampilkan daftar seluruh jadwal kuliah khusus untuk kelas mahasiswa tersebut.
    * **Fitur:**
        * *Urutkan Kelas:* Menggunakan **Selection Sort / Insertion Sort** berdasarkan jam mulai paling pagi.
        * *Cari Kelas/Dosen:* Menggunakan **Sequential Search / Binary Search** berdasarkan nama Mata Kuliah atau nama Dosen.
* **2. Cek Sisa Kuliah Hari Ini** * Menampilkan informasi jumlah mata kuliah tersisa yang harus dihadiri pada hari berjalan.
* **3. Informasi Ruangan & Kapasitas** * Menampilkan daftar ruangan (misal: Lab 1, Ruang 302) untuk mengecek lokasi kelas serta kapasitas ruangan agar tidak salah tempat.
* **4. Lihat Teman Sekelas** * Menampilkan daftar mahasiswa lain yang berada di dalam satu kelas yang sama.

---

## 📊 Entity-Relationship Diagram (ERD)

```mermaid
erDiagram
    AKUN ||--o| DOSEN : "memiliki"
    AKUN ||--o| MAHASISWA : "memiliki"
    DOSEN ||--o{ KELAS : "dosen wali"
    KELAS ||--o{ MAHASISWA : "berisi"
    MATAKULIAH ||--o{ JADWAL : "diplot ke"
    KELAS ||--o{ JADWAL : "dijadwalkan"
    DOSEN ||--o{ JADWAL : "mengajar"
    RUANGAN ||--o{ JADWAL : "digunakan di"

    AKUN {
        int id_akun PK
        string email UK
        string password
        enum role "admin, dosen, mahasiswa"
        enum status "aktif, nonaktif"
    }

    DOSEN {
        string nidn PK
        string nama_dosen
        string jurusan
        int id_akun FK
    }

    KELAS {
        int id_kelas PK
        string nama_kelas
        int angkatan
        int kapasitas
        string nidn_dosen_wali FK
    }

    MAHASISWA {
        string nim PK
        string nama_mahasiswa
        string jurusan
        int id_kelas FK
        int id_akun FK
    }

    MATAKULIAH {
        string kode_mk PK
        string nama_mk
        int sks
        int semester
    }

    RUANGAN {
        int id_ruangan PK
        string nama_ruangan UK
        int kapasitas_ruangan
        string lokasi_gedung
    }

    JADWAL {
        int id_jadwal PK
        string kode_mk FK
        int id_kelas FK
        string nidn_dosen FK
        int id_ruangan FK
        enum hari "Senin, Selasa, Rabu, Kamis, Jumat, Sabtu"
        time jam_mulai
        time jam_selesai
    }