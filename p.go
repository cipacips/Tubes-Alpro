package main

import (
	"fmt"
	"os"
	"os/exec"
)

type pasien struct {
	nama, email, alamat, status string
	umur                        int
	gejala                      [12]string
	point                       int
}

var pasienList [20]pasien

var jumlahPasien int

func menu() {

	fmt.Println("Menu:")
	fmt.Println("1. Daftar Pasien")
	fmt.Println("2. Daftar Gejala Covid-19")
	fmt.Println("3. Tampilkan Daftar Pasien")
	fmt.Println("4. Cari Nama Pasien Berdasarkan Status")
	fmt.Println("5. Edit Data Pasien")
	fmt.Println("6. Hapus Data Pasien")
	fmt.Println("9. Clear Terminal")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih (1/2/3/4/5/6/9/0)? : ")
}

func main() {
	var pilih string
	fmt.Println("Selamat Datang")
	fmt.Println("")
	for {
		menu()
		fmt.Scan(&pilih)
		switch pilih {
		case "1":
			daftar()
		case "2":
			listGejalaCOVID19()

		case "3":
			sortingpointurgensi()
		case "4":
			search()
		case "5":
			edit_pasien()
		case "6":
			delete_pasien()
		case "9":
			clearscreen()

		case "0":
			fmt.Println("")
			fmt.Println("Terima kasih telah menggunakan sistem screening COVID-19.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")

		}
	}

}

func daftar() {
	fmt.Println("Masukkan data pasien:")
	fmt.Print("Nama: ")
	fmt.Scan(&pasienList[jumlahPasien].nama)
	fmt.Print("Umur: ")
	fmt.Scan(&pasienList[jumlahPasien].umur)
	fmt.Print("Alamat: ")
	fmt.Scan(&pasienList[jumlahPasien].alamat)
	fmt.Print("Email: ")
	fmt.Scan(&pasienList[jumlahPasien].email)
	pointPenyakit(jumlahPasien)

	jumlahPasien = jumlahPasien + 1
	fmt.Println("")
	fmt.Println("Data pasien berhasil dimasukkan.")
	fmt.Println("")

}

func listGejalaCOVID19() {
	fmt.Println("")
	fmt.Println("-----------------------")
	fmt.Println("DAFTAR GEJALA COVID")
	fmt.Println("-----------------------")
	fmt.Println("1. Mata Merah")
	fmt.Println("2. Hilangnya Indra Penciuman")
	fmt.Println("3. Hilangnya kemampuan mengecap rasa")
	fmt.Println("4. Sakit Tenggorokan")
	fmt.Println("5. Sakit Kepala")
	fmt.Println("6. Nyeri Otot dan Sendi")
	fmt.Println("7. Batuk Kering")
	fmt.Println("8. Kelelahan")
	fmt.Println("9. Demam")
	fmt.Println("10. Diare")
	fmt.Println("11. Ruam Kulit atau Perubahan Warna pada Jari Tangan atau Kaki")
	fmt.Println("12. Sesak Napas")
	fmt.Println("-----------------------")
}

func pointPenyakit(i int) {
	var text, text2, text3, text4, confirm string
	text = "Apakah Kamu Merasakan :"
	text2 = "Ketik Y / N"
	text3 = "Mohon Masukkan Input Yang Benar :"
	text4 = "Gejala Apa Yang Kamu Rasakan..?"

	fmt.Println("")
	fmt.Println(text4)
	fmt.Println(text)
	fmt.Println("Mata Merah?")
	fmt.Println("Y untuk Konfirmasi / Iya")
	fmt.Println("N Jika Tidak Memiliki Gejala")
	fmt.Println("")
	fmt.Println(text2)
	fmt.Println("")
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 1

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Hilangnya Indra Penciuman?")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 2

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Hilangnya Kemampuan Mengecap Rasa?")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 2

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Sakit Tenggorokan")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 1

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Sakit Kepala")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 1

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Nyeri Otot dan Sendi?")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 1

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Batuk Kering")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 1

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Kelelahan")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 1

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Demam")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 5

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Diare")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 2

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Ruam Kulit atau Perubahan Warna pada Jari Tangan atau Kaki")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 5

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("Sesak Napas")
	fmt.Println(text2)
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		pasienList[i].point = pasienList[i].point + 7

	} else if confirm == "N" || confirm == "n" {

	} else {
		fmt.Println(text3)
		fmt.Scan(&confirm)
	}

	if pasienList[i].point > 12 {
		pasienList[i].status = "Terkonfirmasi"
	} else if pasienList[i].point > 5 {
		pasienList[i].status = "Suspect"
	} else if pasienList[i].point >= 1 {
		pasienList[i].status = "Probable"
	} else if pasienList[i].point == 0 {
		pasienList[i].status = "BebasCovid"
	}

}

func sortingpointurgensi() {
	for i := 0; i < jumlahPasien-1; i++ {
		for j := 0; j < jumlahPasien-i-1; j++ {
			if pasienList[j].point < pasienList[j+1].point {
				pasienList[j], pasienList[j+1] = pasienList[j+1], pasienList[j]
			}
		}
	}
	if jumlahPasien > 0 {
		fmt.Println("Pasien telah diurutkan berdasarkan point urgensi.")
		fmt.Println("")
		tampillistpasien()
	} else {
		fmt.Println("Maaf belum ada pasien yang terdaftar, Silahkan input pasien")
	}

}

func search() {
	var i, found int
	var cari string

	fmt.Println("")
	fmt.Println("Masukkan Status Pasien Yang Ingin Ditampikan :")
	fmt.Scan(&cari)
	fmt.Println("")

	found = -1

	i = 0
	for i < jumlahPasien {
		if pasienList[i].status == cari {
			fmt.Println("Berikut Data Pasien yang ditemukan :")
			fmt.Println("")
			fmt.Printf("%d. Nama: %s\n   Umur: %d\n   email : %s\n   alamat : %s\n   Status : %s\n", i+1, pasienList[i].nama, pasienList[i].umur, pasienList[i].email, pasienList[i].alamat, pasienList[i].status)
			found = i
		}
		i = i + 1
	}
	if found == -1 {
		fmt.Println("Maaf Data Yang kamu cari tidak ditemukan")
		fmt.Println("")
	}

	fmt.Print("")
}

func tampillistpasien() {

	fmt.Println("Daftar Pasien:")
	for i := 0; i < jumlahPasien; i++ {
		fmt.Println("")
		fmt.Printf("%d. Nama: %s\n   Umur: %d\n   email : %s\n   alamat : %s\n   Status: %s\n", i+1, pasienList[i].nama, pasienList[i].umur, pasienList[i].email, pasienList[i].alamat, pasienList[i].status)
	}

	fmt.Print("")
}

func clearscreen() {
	fmt.Println("Apakah ingin Menghapus Terminal?")
	fmt.Println("")
	fmt.Println("Tekan Y jika ingin menghapus")

	var x string

	fmt.Scan(&x)
	if x == "y" || x == "Y" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		menu()
	}
}

func edit_pasien() {
	var i, found int
	var cari, konfirmasi string
	fmt.Println("Masukkan nama pasien yang ingin diedit :")
	fmt.Scan(&cari)
	i = 0
	found = -1

	for i < jumlahPasien {
		if pasienList[i].nama == cari {
			fmt.Println("Berikut Data Pasien yang ditemukan :")
			fmt.Println("")
			fmt.Printf("%d. Nama: %s\n   Umur: %d\n   email : %s\n   alamat : %s\n   Status : %s\n", i+1, pasienList[i].nama, pasienList[i].umur, pasienList[i].email, pasienList[i].alamat, pasienList[i].status)
			fmt.Println("Apakah anda yakin ingin mengedit nama pasien ini?")
			fmt.Println("Masukkan Y jika iya atau N jika tidak")
			fmt.Scan(&konfirmasi)
			if konfirmasi == "N" {
				break
			}
			found = i

			fmt.Println("Masukkan data pasien:")
			fmt.Print("Nama: ")
			fmt.Scan(&pasienList[found].nama)
			fmt.Print("Umur: ")
			fmt.Scan(&pasienList[found].umur)
			fmt.Print("Alamat: ")
			fmt.Scan(&pasienList[found].alamat)
			fmt.Print("Email: ")
			fmt.Scan(&pasienList[found].email)
			pointPenyakit(jumlahPasien)
		}
		i = i + 1
	}
	if found == -1 {
		fmt.Println("Maaf Data Yang kamu cari tidak ditemukan")
		fmt.Println("")
		return
	}

}

func delete_pasien() {
	var i, found int
	var cari, konfirmasi string
	fmt.Println("Masukkan nama pasien yang ingin dihapus :")
	fmt.Scan(&cari)
	i = 0
	found = -1

	for i < jumlahPasien {
		if pasienList[i].nama == cari {
			fmt.Println("Berikut Data Pasien yang ditemukan :")
			fmt.Println("")
			fmt.Printf("%d. Nama: %s\n   Umur: %d\n   email : %s\n   alamat : %s\n   Status : %s\n", i+1, pasienList[i].nama, pasienList[i].umur, pasienList[i].email, pasienList[i].alamat, pasienList[i].status)
			fmt.Println("Apakah anda yakin ingin menghapus nama pasien ini?")
			fmt.Println("Masukkan Y jika iya atau N jika tidak")
			fmt.Scan(&konfirmasi)
			if konfirmasi == "N" {
				break
			}

			found = i
		}
		i = i + 1
	}
	if found == -1 {
		fmt.Println("Maaf Data Yang kamu cari tidak ditemukan")
		fmt.Println("")
		return
	}

	for i = found; i < jumlahPasien-1; i++ {
		pasienList[i] = pasienList[i+1]
	}
	jumlahPasien = jumlahPasien - 1
}
