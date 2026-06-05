package main

// TYPE
type Akun struct{
	id int
	email string
	password string
	role string 
	status string
}

type Dosen struct{
	nidn string
	namaDosen string
	jurusan string
	akun Akun
}

type Mahasiswa struct {
	nim string
	namaMahasiswa string
	jurusan string
	kelas Kelas
	akun Akun
}

type Kelas struct{
	id int 
	namaKelas string
	angkatan int 
	kapasitas int 
	dosen Dosen
}

type MataKuliah struct{
	kodeMk string
	namaMk string
	sks int
	semester int 
}
type Ruangan struct{
	id int
	namaRuangan string
	kapasitasRuangan string
	lokasiGedung string
}

type Jadwal struct{
	id int
	mataKuliah MataKuliah
	kelas Kelas
	dosen Dosen
	ruangan Ruangan
	hari string
	jamMulai string
	jamSelesai string
}


func main(){

}