package main

import "fmt"

func main() {
	var JumlahBuku, HargaBuku, Diskon int
	var MemberShip string
	fmt.Scan(&JumlahBuku, &MemberShip)
	HargaBuku = JumlahBuku * 10000
	if JumlahBuku < 5 && MemberShip == "A" {
		Diskon = HargaBuku - int(float64(HargaBuku)*(0.1))
		fmt.Println("Rp ", Diskon)
	} else if JumlahBuku >= 5 && JumlahBuku <= 10 && MemberShip == "A" {
		Diskon = HargaBuku - int(float64(HargaBuku)*(0.2))
		fmt.Println("Rp ", Diskon)
	} else if JumlahBuku > 10 && MemberShip == "A" {
		Diskon = HargaBuku - int(float64(HargaBuku)*(0.3))
		fmt.Println("Rp ", Diskon)
	}
	if JumlahBuku < 5 && MemberShip == "B" {
		Diskon = HargaBuku - int(float64(HargaBuku)*(0.05))
		
		fmt.Println("Rp ", Diskon)
	} else if JumlahBuku >= 5 && JumlahBuku <= 10 && MemberShip == "B" {
		Diskon = HargaBuku - int(float64(HargaBuku)*(0.1))
		fmt.Println("Rp ", Diskon)
	} else if JumlahBuku > 10 && MemberShip == "B" {
		Diskon = HargaBuku - int(float64(HargaBuku)*(0.15))
		fmt.Println("Rp ", Diskon)
	}
	if JumlahBuku >= 5 && JumlahBuku <= 10 && MemberShip == "C" {
		Diskon = HargaBuku - int(float64(HargaBuku)*(0.05))
		fmt.Println("Rp ", Diskon)
	} else if JumlahBuku > 10 && MemberShip == "C" {
		Diskon = HargaBuku - int(float64(HargaBuku)*(0.10))
		fmt.Println("Rp ", Diskon)
	} else if MemberShip == "N" {
		fmt.Println("Rp ", HargaBuku)
	} else if JumlahBuku > 10 && MemberShip == "N" {
		Diskon = HargaBuku - int(float64(HargaBuku)*(0.05))
	}
}
