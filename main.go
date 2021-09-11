package main

// func cetakNama(nama string) string {
// 	// fmt.Println(nama)
// 	// fmt.Println(gender)
// 	return nama
// }

// func cetakNama(nama string) (string, string) {
// 	// fmt.Println(nama)
// 	// fmt.Println(gender)
// 	return nama, "______"
// }

// Fungsi Variadic
// func cetakNama(nama ...string) {
// 	for _, value := range nama {
// 		fmt.Println(value)
// 	}
// }

// Fungsi sebagai return value
// func generatorBilanganGenap() func() int {
// 	i := 0
// 	return func() int {
// 		i += 2
// 		return i
// 	}
// }

// Struct
// type Person struct {
// 	firstname string
// 	lastname  string
// 	age       int
// }

// func (p *Person) Biodata() string {
// 	return fmt.Sprintf(`nama: '%s' '%s', age: %d`, p.firstname, p.lastname, p.age)
// }

// Hitung ...
// type Hitung interface {
// 	luas() float64
// 	keliling() float64
// }
// type lingkaran struct {
// 	diameter float64
// }
// type persegi struct {
// 	sisi float64
// }

// // Hitungan lingkaran
// func (l lingkaran) jariJari() float64 {
// 	return l.diameter / 2
// }
// func (l lingkaran) luas() float64 {
// 	return math.Phi * math.Pow(l.jariJari(), 2)
// }
// func (l lingkaran) keliling() float64 {
// 	return math.Phi * l.diameter
// }

// // Hitungan persegi
// func (p persegi) luas() float64 {
// 	return math.Pow(p.sisi, 2)
// }
// func (p persegi) keliling() float64 {
// 	return p.sisi * 4
// }

// Recover function
// func catch() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Error occured", r)
// 	} else {
// 		fmt.Println("Application running perfectly")
// 	}
// }

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

	// Fungsi part 1
	// cetakNama("Francesca Vania")
	// love, you := cetakNama("Francesca Vania")
	// fmt.Println(love, you)
	// cetakNama("Naruto", "Luffy", "Goku")

	// Fungsi part 2
	// i := 0
	// increment := func() int {
	// 	i += 2
	// 	return i
	// }
	// fmt.Println(increment())
	// Hasil run fungsi sbg return value
	// bilanganGenap := generatorBilanganGenap()
	// fmt.Println(bilanganGenap())
	// Hasil run fungsi sbg parameter
	// perulangan([]int{1, 3, 7, 9}, func(val int) {
	// 	fmt.Println(val)
	// })

	// Struct
	// p1 := Person{
	// 	firstname: "Francesca",
	// 	lastname:  "Vania",
	// 	age:       24,
	// }
	// fmt.Println(p1.Biodata())

	// Interface
	// var bangunDatar Hitung

	// bangunDatar = persegi{10.0}
	// fmt.Println("===> Persegi")
	// fmt.Println("luas		:", bangunDatar.luas())
	// fmt.Println("keliling	:", bangunDatar.keliling())

	// bangunDatar = lingkaran{14.0}
	// fmt.Println("===> Lingkaran")
	// fmt.Println("luas		:", bangunDatar.luas())
	// fmt.Println("keliling	:", bangunDatar.keliling())
	// fmt.Println("jari-jari	:", bangunDatar.(lingkaran).jariJari())
	// Interface Kosong
	// var bangunDatar interface{}
	// bangunDatar = "persegi"
	// fmt.Println(reflect.TypeOf(bangunDatar))
	// bangunDatar = 10
	// fmt.Println(reflect.TypeOf(bangunDatar))
	// bangunDatar = []string{"persegi"}
	// fmt.Println(reflect.TypeOf(bangunDatar))
	// bangunDatar = map[int]string{1: "persegi"}
	// fmt.Println(reflect.TypeOf(bangunDatar))

	// Defer
	// defer fmt.Println("halo 1")
	// fmt.Println("halo 2")
	// Jika defer lebih dari satu
	// defer fmt.Println("halo 1")
	// defer fmt.Println("halo 2")
	// fmt.Println("halo 3")
	// num := 2
	// if num == 2 {
	// 	func() {
	// 		defer fmt.Println("halo 1")
	// 		fmt.Println("halo 4")
	// 	}()
	// 	fmt.Println("halo 2")
	// }
	// fmt.Println("halo 3")

	// Error, Panic, Recover
	// input := "is"
	// number, err := strconv.Atoi(input)
	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	fmt.Println(err.Error())
	// 	fmt.Println(input, "is not number")
	// }
	// input := "is"
	// number, err := strconv.Atoi(input)
	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	panic(err.Error())
	// 	fmt.Println(input, "is not number")
	// } //Panic tidak harus digunakan saat error
	// input := "is"
	// number, err := strconv.Atoi(input)
	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	fmt.Println(err.Error())
	// 	fmt.Println(input, "is not number")
	// }
	// defer catch()
	// input := "ls"
	// number, err := strconv.Atoi(input)
	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	panic(err.Error())
	// 	fmt.Println(input, "is not number")
	// }
}

// Fungsi sebagai parameter
// func perulangan(angka []int, callback func(int)) {
// 	for _, value := range angka {
// 		callback(value)
// 	}
// }
