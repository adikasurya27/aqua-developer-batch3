Introductin to Golang

- Apa itu Golang?
Go Language merupakan kepanjangan dari bahasa pemrograman satu ini. Seperti kita tahu namanya diawali dengan Go, ya benar dia merupakan bahasa pemrograman yang dikelola oleh Google. Google tidak bekerja sendirian, melainkan bekerja sama dengan 3 orang tokoh handal pada tahun 2009. Robert Griesemer, Rob Pike dan Ken Thompson merupakan ketiga tokoh tersebut.

- Type data

1. Numeric Non Decimal
Tipe data	Cakupan bilangan
uint8	    0 ↔ 255
uint16	    0 ↔ 65535
uint32	    0 ↔ 4294967295
uint64	    0 ↔ 18446744073709551615
uint	    sama dengan uint32 atau uint64 (tergantung nilai)
byte	    sama dengan uint8
int8	    -128 ↔ 127
int16	    -32768 ↔ 32767
int32	    -2147483648 ↔ 2147483647
int64	    -9223372036854775808 ↔ 9223372036854775807
int	        sama dengan int32 atau int64 (tergantung nilai)
rune	    sama dengan int32

var positiveNumber uint8 = 89
var negativeNumber = -1243423644

fmt.Printf("bilangan positif: %d\n", positiveNumber)
fmt.Printf("bilangan negatif: %d\n", negativeNumber)

Variabel positiveNumber bertipe uint8 dengan nilai awal 89. Sedangkan variabel negativeNumber dideklarasikan dengan nilai awal -1243423644. Compiler secara cerdas akan menentukan tipe data variabel tersebut sebagai int32 (karena angka tersebut masuk ke cakupan tipe data int32).

2. Numeric Decimal
Tipe data numerik desimal yang perlu diketahui ada 2, float32 dan float64. Perbedaan kedua tipe data tersebut berada di lebar cakupan nilai desimal yang bisa ditampung.

var decimalNumber = 2.62

fmt.Printf("bilangan desimal: %f\n", decimalNumber)
fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

Pada kode di atas, variabel decimalNumber akan memiliki tipe data float32, karena nilainya berada di cakupan tipe data tersebut. Ketika di run hasilnya:

bilangan desimal: 2.620000
bilangan desimal: 2.620

Template %f digunakan untuk memformat data numerik desimal menjadi string. Digit desimal yang akan dihasilkan adalah 6 digit. Pada contoh di atas, hasil format variabel decimalNumber adalah 2.620000. Jumlah digit yang muncul bisa dikontrol menggunakan %.nf, tinggal ganti n dengan angka yang diinginkan.

3. Boolean
Tipe data bool berisikan hanya 2 variansi nilai, true dan false. Tipe data ini biasa dimanfaatkan dalam seleksi kondisi dan perulangan.

var exist bool = true
fmt.Printf("exist? %t \n", exist)

Gunakan %t untuk memformat data bool menggunakan fungsi fmt.Printf().

4. String
Ciri khas dari tipe data string adalah nilainya di apit oleh tanda quote atau petik dua ("). Contoh penerapannya:

var message string = "Halo"
fmt.Printf("message: %s \n", message)

Selain menggunakan tanda quote, deklarasi string juga bisa dengan tanda grave accent/backticks (`), tanda ini terletak di sebelah kiri tombol 1. Keistimewaan string yang dideklarasikan menggunakan backtics adalah membuat semua karakter di dalamnya tidak di escape, termasuk \n, tanda petik dua dan tanda petik satu, baris baru, dan lainnya. Semua akan terdeteksi sebagai string.

- Variabel
1. Variabel beserta tipe data
package main

import "fmt"

func main() {
    var firstName string = "john"

    var lastName string
    lastName = "wick"

    fmt.Printf("halo %s %s!\n", firstName, lastName)
}
Keyword var di atas digunakan untuk deklarasi variabel, contohnya bisa dilihat pada firstName dan lastName.

2. Variabel menggunakan keyword var
Keyword var digunakan untuk membuat variabel baru.

var <nama-variabel> <tipe-data>
var <nama-variabel> <tipe-data> = <nilai>
contoh:
var lastName string
var firstName string = "john"

Menggunakan fungsi fmt.Printf()

fmt.Printf("halo john wick!\n")
fmt.Printf("halo %s %s!\n", firstName, lastName)
fmt.Println("halo", firstName, lastName + "!")

Fungsi fmt.Printf() tidak menghasilkan baris baru di akhir text, oleh karena itu digunakanlah literal newline yaitu \n, untuk memunculkan baris baru di akhir. Hal ini sangat berbeda jika dibandingkan dengan fungsi fmt.Println() yang secara otomatis menghasilkan new line (baris baru) di akhir.

3. Variabel tanpa tipe data
Metode jenis ini, keyword var dan tipe data tidak perlu ditulis. Variabel lastName dideklarasikan dengan menggunakan metode type inference. Penandanya tipe data tidak dituliskan pada saat deklarasi. Pada penggunaan metode ini, operand = harus diganti dengan := dan keyword var dihilangkan.
// menggunakan var, tanpa tipe data, menggunakan perantara "="
var firstName = "john"

// tanpa var, tanpa tipe data, menggunakan perantara ":="
lastName := "wick"

4. Multi Variabel
Go mendukung metode deklarasi banyak variabel secara bersamaan, caranya dengan menuliskan variabel-variabel-nya dengan pembatas tanda koma (,)
var first, second, third string
first, second, third = "satu", "dua", "tiga"

5. Inderscore
Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah.
_ = "belajar Golang"
_ = "Golang itu mudah"
name, _ := "john", "wick"

- Condition
1. If...Else
Pernyataan if....else memungkinkan Anda untuk mengeksekusi satu blok kode jika kondisi yang ditentukan bernilai benar dan blok kode lainnya jika bernilai salah.
if  condition { 
    // kode yang akan dieksekusi jika kondisinya benar
} else {
    // code to be executed if condition is false
}

2. Switch Case
Pernyataan switch digunakan untuk memilih salah satu dari banyak blok kode yang akan dieksekusi.
func main() {
	today := time.Now()
 
	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}

- Looping
1. for...range
Syntax :
for index, value := range mydatastructure {
        fmt.Println(value)
    }
example :
package main
import "fmt"

func main() {
  for i, ch := range "World" {
    fmt.Printf("%#U starts at byte position %d\n", ch, i)
  }
}

2. for...loop
for initialization; condition; update {
  statement(s)
}
example :
package main
import "fmt"

func main() {

  for i := 1; i <= 5; i++ {
    fmt.Println(i)
  }
}

- Functions

1. Function
package main
import "fmt"

// SimpleFunction prints a message
func SimpleFunction() {
	fmt.Println("Hello World")
}

func main() {
	SimpleFunction()
}

2. Consecutive named function parameters
// from
func add(x int, y int) int {
	return x + y
}

// to
func add(x, y int) int {
	return x + y
}

3. Multiple Return values
func swap(x, y string) (strig, string) {
	return y, x
}
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

- Struct
1. Struct
type student struct {
	name string
	grade int
}
func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)
}

2. Embedded struct
type person struct {
	name string
	age int
}

type student struct {
	grade int
	person
}
func main() {
	var s1 = student{
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("age :", s1.age)
	fmt.Println("age :", s1.person.age)
	fmt.Println("grade :", s1.grade)
	}
}

3. Combine struct with slice
type person struct {
	name string
	age int
}

var allStudents = []person{
	{name: "wick", age: 23},
	{name: "ethan", age: 23},
	{name: "bourne", age: 22},
}
for _, student := range allStudents {
	fmt.Println(student.name, "age is", student.age)
}

- Data Structure
1. Map
var chicken map[string]int
chicken = map[string]int{}

chicken["januari"] = 50
chicken["februari"] = 40

fmt.Prinln("januari", chicken["januaru"]) 
fmt.Prinln("mei", chicken["mei"])

2. Array
var names [4]string
name[0] = "trafalgar"
name[1] = "d"
name[2] = "water"
name[3] = "law"

fmt.Println(name[0], names[1], names[2], names[3])

- Slice
1. Slice
var fruits = []string{"apple", "grape", "banana", "melon"}
fmt.Println(fruits[0]) // "apple"

var fruitsA = []string{"apple", "grape"}      // slice
var fruitsB = [2]string{"banana", "melon"}    // array
var fruitsC = [...]string{"papaya", "grape"}  // array

Kode		Output									Penjelasan
fruits[0:2]	[apple, grape]					semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
fruits[0:4]	[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
fruits[0:0]	[]								menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
fruits[4:4]	[]								menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
fruits[4:0]	[]								error, pada penulisan fruits[a:b] nilai a harus lebih kecil atau sama dengan b
fruits[:]	[apple, grape, banana, melon]	semua elemen
fruits[2:]	[banana, melon]					semua elemen mulai indeks ke-2
fruits[:2]	[apple, grape]					semua elemen hingga sebelum indeks ke-2

2. Defer
package main

import "fmt"

func main() {
	defer fmt.Println("halo")
	fmt.Println("Selamat Datang")
}