package main

import (
	"fmt"
	"os"
)

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


var akun = make([]Akun, 0)
var session *Akun
// var dosen = make([]Dosen,0)
	// var mahasiswa = make([]Mahasiswa,0)
	// var kelas = make([]Kelas,0)
	// var mataKuliah = make([]MataKuliah,0)
	// var ruangan = make([]Ruangan,0)
	// var jadwal = make([]Jadwal,0)
func main(){

	

	// dummy data admin
	akun = append(akun,Akun{id: 1,email: "admin@telkomuniversity.ac.id",password: "admin@123",role: "admin"})
	akun = append(akun,Akun{id: 2,email: "andrianhidayat@student.telkomuniversity.ac.id",password: "andre@1",role: "mahasiswa"})
	akun = append(akun,Akun{id: 2,email: "dimas@telkomuniversity.ac.id",password: "dimas@1",role: "dosen"})

	loginView()
	routeHalaman(session.role)

}
func headerApp(){
	fmt.Println()
	fmt.Println("==============")
	fmt.Println("   JADWALKU   ")
	fmt.Println("==============")
}

func loginView(){
	headerApp()
	fmt.Println("Halaman Login")
	fmt.Println("--------------")
	fmt.Println()
	var inputEmail,password string
	fmt.Print("Email : ")
	fmt.Scan(&inputEmail)
	fmt.Print("Password : ")
	fmt.Scan(&password)
	loginService(inputEmail,password)
}

func loginService(email,password string){

	for _,value := range akun{
		
		if(value.email == email && value.password == password){
			session = &value
			return
		}
	}
	fmt.Println()
	fmt.Println("..............")
	fmt.Println("Email atau Password anda salah!!")
	fmt.Println("..............")
	fmt.Println()
	loginView()
}

func routeHalaman(role string){
	if role == "admin"{
		adminView()
	}

	if role == "mahasiswa"{
	fmt.Println("Halaman Mahasiswa")
	}

	if role == "dosen"{
	fmt.Println("Halaman Dosen")

	}
}

func adminView(){
	var input string
	headerApp()
	fmt.Println("Halaman Admin")
	fmt.Println("--------------")
	fmt.Println()
	fmt.Println("1. Master Akun")
	fmt.Println("2. Master Dosen")
	fmt.Println("3. Master Jadwal")
	fmt.Println("4. Master Mahasiswa")
	fmt.Println("5. Master Mata Kuliah")
	fmt.Println("6. Master Ruangan")
	fmt.Println("--------------")
	fmt.Println("L. Logout")
	fmt.Println("X. Tutup Aplikasi")
	fmt.Println("--------------")
	fmt.Println()
	fmt.Println("Contoh ( Pilih Halaman : 1 )")
	fmt.Print("Pilih Halaman : ")

	fmt.Scan(&input)

	switch input{
	case "1":
		fmt.Print("Akun")
		case "2":
		fmt.Print("Akun")
		case "3":
		fmt.Print("Akun")
		case "4":
		fmt.Print("Akun")
		case "5":
		fmt.Print("Akun")
		case "6":
		fmt.Print("Akun")
		case "L":
			loginView()
			session = nil
		case "X":
		os.Exit(0)
		default:
		adminView()
		fmt.Println("Inputan tidak valid")	
	}
}

