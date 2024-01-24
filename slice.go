package main

import "fmt"

func main() {
	names := [...]string{"andrew", "wiliam", "budi", "wawan", "abraham", "dafa"}

	slice1 := names[3:5]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	dayslice := days[5:]
	fmt.Println(dayslice)

	dayslice[0] = "sabtu baru"
	dayslice[1] = "minggu baru"
	fmt.Println(dayslice)

	dayslice2 := append(dayslice, "Libur Baru")

	fmt.Println(dayslice)
	fmt.Println(dayslice2)
	fmt.Println(days)
}
