package main

func main() {
	// Operator Aritmatik
	// Penjumlahan
	// a, b := 10, 4
	// c := a + b
	// fmt.Println(a + b)
	// fmt.Println(c)

	// var a int8 = 10
	// var b int8 = 5
	// c := a + int8(b)
	// fmt.Println(c)

	// Pembagian
	// var a float32 = 10
	// var b float32 = 3
	// c := a / b
	// fmt.Println(c)

	// Modulus
	// var a int = 10
	// var b int = 3
	// c := a % b
	// fmt.Println(c)

	// Operator Logika
	// var a bool = true
	// var b bool = false
	// fmt.Println(a && b)
	// fmt.Println(a || b)
	// fmt.Println(!a)
	// fmt.Println(!b)

	// Operator Perbandingan
	// var a int = 10
	// var b int8 = 8
	// fmt.Println(a == int(b))
	// fmt.Println(a >= int(b))

	// Perulangan if-else part 1
	// warna := "biru"

	// if warna == "merah" {
	// 	fmt.Println("berhenti")
	// } else if warna == "hijau" {
	// 	fmt.Println("jalan")
	// } else {
	// 	fmt.Println("lainnya")
	// }

	// perulangan if-else temporary variable part 2
	// if warnaLampu := "merah"; warnaLampu == "merah" {
	// 	fmt.Println("Berhenti")
	// } else if warnaLampu == "kuning" {
	// 	fmt.Println("Hati-hati")
	// } else {
	// 	fmt.Println("Silahkan melaju")
	// }

	// Switch Case
	// var warnaLampu = "kuning"

	// switch warnaLampu {
	// case "merah":
	// 	{
	// 		fmt.Println("Berhenti")
	// 	}
	// case "kuning":
	// 	fmt.Println("Hati-hati")
	// default:
	// 	fmt.Println("Silahkan jalan")
	// }

	// Switch Case part 2
	// var warnaLampu = "merah"

	// switch {
	// case warnaLampu == "merah":
	// 	{
	// 		fmt.Println("Berhenti")
	// 	}
	// case warnaLampu == "kuning":
	// 	fmt.Println("Hati-hati")
	// case warnaLampu == "hijau":
	// 	fmt.Println("Silahkan melaju")
	// 	fallthrough
	// default:
	// 	fmt.Println("Tidak tahu")
	// }

	// For Loop
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("Berhitung: ", i)
	// }

	// For Range
	// names := []string{"John", "Doe", "Amalia"}
	// for _, name := range names {
	// 	fmt.Println(name)
	// }

	// While Loop
	// var i = 0
	// for i <= 10 {
	// 	fmt.Println("Berhitung", i)
	// 	i += 2
	// }

	// Do While Loop
	// var i = 0
	// for {
	// 	fmt.Println("Berhitung", i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("Angka Genap", i)
	// }

	// Looping Label
	// loopTerluar:
	// for i := 1; i < 5; i++ {
	// 	for k := 1; k < 5; k++ {
	// 		fmt.Println("[", i, "]", "[", k, "]")
	// 		if i == 2 {
	// 			break loopTerluar // ini untuk menghentikan perulangan for berdua
	// 		}
	// 		// kondisi berjalan di for K
	// 	}
	// }

	// Array Slice Map
	// var names = [...]string{"Jean", "Pierre", "Polnaref"}
	// var angka = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// // names[0] = "Jean"
	// // names[1] = "Pierre"
	// // names[2] = "Polnaref"
	// // fmt.Println(names[0], names[1], names[2])
	// // fmt.Println(len(names))
	// // fmt.Println("Nama: ", names, " terdiri dari ", len(names), " kata")
	// fmt.Println("Angka: ", angka, " terdiri dari ", len(angka), " angka")

	// Array Slice Map part 2
	// var names = [5]int{10, 20, 30, 40, 50}
	// // For Oldschool
	// // for i := 0; i < len(names); i++ {
	// // 	fmt.Println("Berhitung", names[i])
	// // }
	// // For range
	// // for index, value := range names {
	// // 	fmt.Println("index", index, "value", value)
	// // }
	// for _, value := range names {
	// 	fmt.Println("value", value)
	// }

	// // Slice Map part 1
	// var names1 = []int{1, 2, 3, 4, 5}
	// // var names2 = names1[0:4]

	// var names2 = make([]int, 2, 3)
	// // names2[0] = "Ironman"
	// // names2[1] = "Antman"

	// names2[0] = 10
	// names2[1] = 20
	// // names2 = append(names2, "Black Panther")
	// // names2 = append(names2, "Captain America")
	// names2 = append(names2, names1...)

	// fmt.Println(names1, names2)
	// // fmt.Println(len(names2), cap(names2))

	// Slice Map part 2
	// var names1 = map[string]string{
	// 	"Januari":  "jan",
	// 	"Februari": "feb",
	// 	"Maret":    "mar",
	// }
	// var names1 = make(map[string]string)
	// names1["Januari"] = "jan"
	// names1["Februari"] = "feb"
	// names1["Maret"] = "mar"
	// for key, value := range names1 {
	// 	fmt.Println("Key", key, "Value", value)
	// }
	// fmt.Println(names1)
	// fmt.Println(len(names1))
}
