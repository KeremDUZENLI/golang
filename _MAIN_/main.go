package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	Print()
}

func Print() {
	PL := fmt.Println

	liste := []interface{}{
		// CreateList(),
		// AppendNumber(5),
		// EachCons([]int{1, 2, 3, 4, 5}, 3),
		// GetInput(),

		// Digitize(1089),
		// ToAlternatingCase("Hello WOrld"),
		// Contamination("abcd", "z"),
		// Divisors(12),
		// CountBy(5, 2),
		// HighAndLow("1 2 3 4 5"),
		// TwoSort([]string{"bitcoin", "take", "Over", "the", "world", "maybe", "who", "knows", "perhaps"}),
		// ExpressionMatter(5, 1, 3),
		// TwiceAsOld(42, 22),
		// ToJadenCase("most trees are blue"),
		// Greet("Kerem"),
		// GetGrade1(10, 80, 100),
		// GetGrade2(10, 80, 100),
		// NbDig(10, 3),
		// Evaporator(10, 10, 5),
		// DNAtoRNA("ATGC"),
		// DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry"),
		// AbbrevName("sam harris"),
		// Arithmetic(8, 2, "add"),
		// SliceString("/aaa"),

		// ArrayDelete(),
		// RowSumOddNumbers(4),

		// TypeFounder(),
		// Between(1, 4),
		// TwoToOne("loopingisfunbutdangerousa", "lessdangerousthancodinga"),
		// EndsWidth("abc", "ab"),
		// FindOdd([]int{10, 20, 30, 56, 67, 90, 10, 20}),
		// Count([]string{"a", "b", "c", "c"}, "c"),

		// Grow([]int{1, 2, 6}),
		// SquareSum([]int{1, 2, 3}),

		// StockList([]string{"BBAR 100", "CDXE 200", "BKWR 300", "BTSQ 400", "DRTY 500"}, []string{"A", "B", "C", "D"}),
	}

	for _, v := range liste {
		PL(v)
	}
}

/*
QUESTIONS ------------------------------------------------------------------------------------------------------------------
*/

func CreateList() []int {
	arr := make([]int, 5)
	for i := 1; i < 6; i++ {
		arr[len(arr)-i] = i
	}
	return arr
}

func AppendNumber(n int) []int {
	arr2 := []int{}
	for i := n; i > 0; i-- {
		arr2 = append(arr2, i)
	}
	return arr2
}

func EachCons(arr3 []int, n int) [][]int {
	arrList := [][]int{}

	for l := 0; l < (len(arr3) - n + 1); l++ {
		arrInt1 := []int{}

		for i := 0; i < n; i++ {
			arrInt1 = append(arrInt1, arr3[i+l])
		}
		arrList = append(arrList, arrInt1)
	}

	return arrList
}

func GetInput() string {
	var inp string

	fmt.Print("INPUT: ")
	fmt.Scanln(&inp)

	return inp
}

func Digitize(n int) []int {

	l := len(strconv.Itoa(n))

	var result []int

	for i := 0; i < l; i++ {

		d := (math.Pow(10, float64(i)))

		num := (n / int(d)) % 10
		result = append(result, num)
		// fmt.Println(num)
	}

	return result

}

func ToAlternatingCase(str string) string {
	splt := strings.Split(str, "")
	le := len(splt)

	for i := 0; i < le; i++ {
		switch splt[i] {
		case strings.ToUpper(splt[i]):
			splt[i] = strings.ToLower(splt[i])

		case strings.ToLower(splt[i]):
			splt[i] = strings.ToUpper(splt[i])
		}
	}

	spltJoin := strings.Join(splt, "")
	return spltJoin
}

func Contamination(text, char string) string {
	splt := strings.Split(text, "")
	le := len(splt)
	var emptyList []string

	for i := 0; i < le; i++ {
		emptyList = append(emptyList, char)
	}

	emptyListJoin := strings.Join(emptyList, "")
	return emptyListJoin
}

func Divisors(n int) int {
	liste := []int{}

	for i := 1; i < n; i++ {
		if n%i == 0 {
			liste = append(liste, i)
		}
	}

	return len(liste)
}

func CountBy(x, n int) []int {
	liste := []int{}

	for i := 0; i < n; i++ {
		liste = append(liste, (x + x*i))
	}

	return liste
}

func HighAndLow(in string) string {

	splt := strings.Split(in, " ")
	emptyList1 := []int{}
	emptyList2 := []int{}
	finalList := []string{}

	for i := 0; i < len(splt); i++ {
		intVar, _ := strconv.Atoi(splt[i])
		emptyList1 = append(emptyList1, intVar)
		emptyList2 = append(emptyList2, intVar)
	}

	for l := 0; l < len(emptyList1); l++ {
		if emptyList1[0] < emptyList1[l] {
			emptyList1[0] = emptyList1[l]
		}
	}

	for n := 0; n < len(emptyList2); n++ {
		if emptyList2[0] > emptyList2[n] {
			emptyList2[0] = emptyList2[n]
		}
	}

	finalList = append(finalList, strconv.Itoa(emptyList1[0]), strconv.Itoa(emptyList2[0]))
	finalJoin := strings.Join(finalList, " ")
	return finalJoin
}

func TwoSort(arr []string) string {
	sort.Strings(arr)

	first := arr[0]
	splt := strings.Split(first, "")

	final := ""

	for i := 0; i < len(splt)-1; i++ {
		final += splt[i] + "***"
	}

	final += splt[len(splt)-1]

	return final
}

func ExpressionMatter(a int, b int, c int) int {
	x1 := (a + b) + c
	x2 := (a + b) * c

	x3 := (a * b) + c
	x4 := (a * b) * c

	x5 := (b + c) * a

	x6 := (b * c) + a

	liste := []int{x1, x2, x3, x4, x5, x6}
	fmt.Println(liste)

	for i := 0; i < 6; i++ {
		if liste[0] < liste[i] {
			liste[0] = liste[i]
		}
	}

	return liste[0]
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	x := dadYearsOld - (sonYearsOld * 2)

	y := math.Abs(float64(x))

	return int(y)
}

func ToJadenCase(str string) string {
	splt := strings.Split(str, " ")
	var liste []string

	for i := 0; i < len(splt); i++ {
		// liste = append(liste, strings.Title(splt[i]))
		liste = append(liste, cases.Title(language.Und, cases.NoLower).String(splt[i]))
	}

	emptyListJoin := strings.Join(liste, " ")
	return emptyListJoin
}

func Greet(name string) string {
	// fill in solution here
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

func GetGrade1(a, b, c int) rune {
	score := (a + b + c) / 3

	if 90 <= score && score <= 100 {
		return 'A'
	} else if 80 <= score && score < 90 {
		return 'B'
	} else if 70 <= score && score < 80 {
		return 'C'
	} else if 60 <= score && score < 70 {
		return 'D'
	} else {
		return 'F'
	}
}

func GetGrade2(a, b, c int) rune {

	score := (a + b + c) / 3

	switch {
	case 90 <= score && score <= 100:
		return 'A'
	case 80 <= score && score < 90:
		return 'B'
	case 70 <= score && score < 80:
		return 'C'
	case 60 <= score && score < 70:
		return 'D'
	default:
		return 'F'
	}
}

func NbDig(n int, d int) int {
	liste := []int{}
	// var liste []int

	count := 0
	// var count int

	sign := strconv.Itoa(d)

	for i := 0; i <= n; i++ {
		liste = append(liste, i*i)
	}

	for l := 0; l <= n; l++ {
		c := strconv.Itoa(liste[l])
		splt := strings.Split(c, "")

		for m := 0; m < len(splt); m++ {
			if splt[m] == sign {
				count++
			}
		}
	}

	return count
}

func Evaporator(content float64, evapPerDay int, threshold int) int {
	day := 0
	limit := content * float64(threshold) / 100

	for i := 0; ; i++ {
		content -= content * (float64(evapPerDay) / 100)
		day += 1

		fmt.Println(content)

		if content <= limit {
			break
		}
	}

	return day
}

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {

	if firstAttacker == fighter1.Name {
		fighter2.Health -= fighter1.DamagePerAttack
		for i := 0; ; i++ {
			if fighter2.Health > 0 {
				fighter1.Health -= fighter2.DamagePerAttack
				if fighter1.Health > 0 {
					fighter2.Health -= fighter1.DamagePerAttack
				} else {
					return fighter2.Name
				}
			} else {
				return fighter1.Name
			}
		}
	} else {
		fighter1.Health -= fighter2.DamagePerAttack
		for l := 0; ; l++ {
			if fighter1.Health > 0 {
				fighter2.Health -= fighter1.DamagePerAttack
				if fighter2.Health > 0 {
					fighter1.Health -= fighter2.DamagePerAttack
				} else {
					return fighter1.Name
				}
			} else {
				return fighter2.Name
			}
		}
	}
}

func AbbrevName(name string) string {
	splt := strings.Split(name, " ")
	var liste []string

	for i := 0; i < len(splt); i++ {
		splt[i] = strings.ToUpper(splt[i])
		spltEach := strings.Split(splt[i], "")
		liste = append(liste, spltEach[0])
	}

	return strings.Join(liste, ".")
}

func Arithmetic(a int, b int, operator string) int {
	if operator == "add" {
		return a + b
	} else if operator == "subtract" {
		return a - b
	} else if operator == "multiply" {
		return a * b
	} else {
		return a / b
	}
}

func SliceString(a string) string {
	return a[1:]
}

type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

func ArrayDelete() []Article {
	x := []Article{{"Title_1", "Description_1", "Content_1"}, {"Title_2", "Description_2", "Content_2"}}
	// y := []Article{{}}

	fmt.Printf("%T\n%v\n", x, x)
	fmt.Println(x[0].Title)

	x = append(x[:1], x[1+1:]...)

	return x
	// fmt.Println(y)
}

func RowSumOddNumbers(n int) int {
	var liste_ext [][]int
	var x int = -1

	for i := 0; i <= n; i++ {
		var liste_int []int

		for l := 0; l < i; l++ {
			x += 2
			liste_int = append(liste_int, x)
			fmt.Printf("%v.Item = %v\n", i, x)
		}

		fmt.Printf("%v\n\n", liste_int)
		liste_ext = append(liste_ext, liste_int)
	}

	fmt.Println(liste_ext)
	fmt.Println(liste_ext[n:])

	// TOTAL
	var total int = 0

	for j := 0; j < len(liste_ext[n:][0]); j++ {
		total += (liste_ext[n:][0][j])
	}

	return total
}

func TypeFounder() any {
	type Log struct {
		Description string      `json:"description"`
		Field       []string    `json:"field"`
		Value       interface{} `json:"value"`
	}

	/* ### BSON ### */
	type Struct_BSON struct {
		ID     int
		LOGS   Log
		STATUS string
		SCORE  int
	}

	var jsonEngineNaturalPerson = bson.D{
		{Key: "_id", Value: "12345678910"},
		{Key: "logs", Value: Log{}},
		{Key: "status", Value: ""},
		{Key: "score", Value: 1000},
	}

	doc, _ := bson.Marshal(jsonEngineNaturalPerson)
	var struct_BSON_New Struct_BSON
	bson.Unmarshal(doc, &struct_BSON_New)

	typ := reflect.TypeOf(struct_BSON_New.LOGS)
	return typ
}

func Between(a, b int) []int {
	var liste []int

	for i := a; i <= b; i++ {
		liste = append(liste, i)
	}

	return liste
}

func TwoToOne(s1 string, s2 string) string {
	var liste1 []string = strings.Split(s1, "")
	var liste2 []string = strings.Split(s2, "")
	var liste0 []string = append(liste1, liste2...)

	sort.Strings(liste0)
	fmt.Println(liste0)

	for i := 0; i < len(liste0)-1; i++ {
		if liste0[i] == liste0[i+1] {
			liste0 = append(liste0[:i], liste0[i+1:]...)
			i = 0
		}
	}
	fmt.Println(liste0)

	/// SPECIAL CONDITION ///
	if liste0[0] == liste0[1] {
		liste0 = append(liste0[:0], liste0[1:]...)
	}

	fmt.Println(liste0)
	/// SPECIAL CONDITION ///

	stringRes := strings.Join(liste0, "")
	return stringRes
}

func EndsWidth(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func FindOdd(seq []int) []int {
	sort.Ints(seq)
	dict := make(map[int]int)
	// var finalListe []int
	finalListe := []int{}

	for _, num := range seq {
		dict[num] = dict[num] + 1
	}
	fmt.Println(dict)

	for number, count := range dict {
		// fmt.Println(number, count)
		if count%2 == 1 {
			finalListe = append(finalListe, number)
		}
	}
	return finalListe
}

func Count(liste []string, findVar string) int {
	dict := make(map[string]int)

	for _, str := range liste {
		dict[str] += 1
	}

	return dict[findVar]
}

func Grow(arr []int) int {
	sum := 1
	for _, i := range arr {
		sum *= i
	}

	return sum
}

func SquareSum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += (i * i)
	}

	return sum
}

func StockList(listArt []string, listCat []string) string {
	listeNam := []string{}
	listeNum := []string{}
	dict := make(map[string]int)

	listeLet := []string{}
	listeAmo := []int{}

	emptStr := ""

	// LIST CATEGORY and LIST NUMBER
	for _, i := range listArt {
		splt1 := strings.Split(i, " ")
		listeNam = append(listeNam, splt1[0])
		listeNum = append(listeNum, splt1[1])
	}

	// DICT LETTER:AMOUNT
	/*
		for _, word := range listeNam {
			letter := string(word[0])
			_, exist := dict[letter]

			if exist {
				dict[letter]++
			} else {
				dict[letter] = 1
			}
		}
	*/

	// DICT LETTER:NUMBER
	for index, word := range listeNam {
		letter := string(word[0])

		number, _ := strconv.Atoi(listeNum[index])
		dict[letter] += number
	}

	// LIST LETTER and LIST AMOUNT
	for _, letterInGivingList := range listCat {
		amount := dict[letterInGivingList]
		listeLet = append(listeLet, letterInGivingList)
		listeAmo = append(listeAmo, amount)
	}

	// STRING
	for i := 0; i < len(listeLet); i++ {
		eStr := fmt.Sprintf("(%s : %v) - ", listeLet[i], listeAmo[i])
		emptStr += eStr
	}

	l := len(emptStr) - 3
	// return listeNam, listeNum, dict, listeLet, listeAmo, emptStr[:l]
	return emptStr[:l]
}
