package main

import "fmt"

/*
Karena di Go Lang Pass By Value,
maka semua data yg di assign ke variable baru
dan jika ada perubahan di Value yg lama
maka variable yang baru tidak berubah
hal ini karena dia membuat memory baru

untuk menjadikannya Pass By Reference
kita bisa pakai pointer

*/

type Address struct {
	City, Province, Country string
}

func main() {
	// Pass By Value
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := address1

	address2.City = "Malang"
	address2.Province = "Jawa Timur"

	fmt.Println(address1)
	fmt.Println(address2)

	// Pass By Reference
	address3 := address1
	address4 := &address3
	var address5 *Address = address4

	address3.City = "Jakarta Utara"

	fmt.Println(address3) // ini akan ikut berubah, karena dia 1 memory dengan address 4
	fmt.Println(address4)
	fmt.Println(address5)

	// operator *
	// berfungsi untuk memaksa reference ikut berubah memory nya ke inheritance nya

	addressA := Address{"1", "2", "3"}
	addressB := &addressA

	addressB = &Address{"4", "5", "6"}

	fmt.Println(addressA) // di sini Address A jadi punya Value yang beda dengan Address B
	fmt.Println(addressB)

	// kita paksa dengan cara seperti ini
	addressA = Address{"1", "2", "3"}
	addressB = &addressA
	*addressB = Address{"4", "5", "6"}

	fmt.Println(addressA)
	fmt.Println(addressB)

	// operator new()
	// untuk membuat pointer kosongan
	addressC := new(Address)
	addressD := &addressC

	fmt.Println(addressC)
	fmt.Println(addressD)

	// pointer pada function
	addressX := Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "",
	}
	fmt.Println(addressX)
	mockCountryFailed(addressX)
	fmt.Println(addressX)
	mockCountrySuccess(&addressX) // pastikan data yg dikirim merupakan pointer
	fmt.Println(addressX)

}

func mockCountryFailed(address Address) {
	address.Country = "Indonesia"
}

func mockCountrySuccess(address *Address) {
	address.Country = "Indonesia"
}
