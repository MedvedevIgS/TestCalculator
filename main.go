package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func seach_sim(str string, chars []string) bool {
	for i := range chars {
		if strings.Contains(str, chars[i]) {
			return true
		}
	}
	return false
}

func seach_byte(b byte, chars []byte) bool {
	for i := range chars {
		if b == chars[i] {
			return true
		}
	}
	return false
}

func check_byte(b string, chars []byte) bool {
	var found bool
	for i := 0; i < len(b); i++ {
		found = false
		j := 0
		for j < len(chars) && found == false {
			if b[i] == chars[j] {
				found = true
			}
			j++
		}
		if !found {
			return false
		}
	}
	return true
}

func seach_abs(str string, chars []string) bool {
	var flag bool = true
	for i := range chars {
		if !strings.Contains(str, chars[i]) {
			flag = false
		}
	}
	return flag
}

func calculation(a int, b int, c string) int {
	var res int
	switch c {
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "+":
		res = a + b
	case "-":
		res = a - b
	}
	return res
}

func CheckRome(a byte, b byte) bool {
	RomeSim := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	var err = false
	switch a {
	case 'I':
		if !seach_byte(b, RomeSim[0:3]) {
			err = true
		}
	case 'V':
		if !seach_byte(b, []byte{RomeSim[0]}) {
			err = true
		}
	case 'X':
		if !seach_byte(b, RomeSim[0:5]) {
			err = true
		}
	case 'L':
		if !seach_byte(b, RomeSim[0:3]) {
			err = true
		}
	case 'C':
		if !seach_byte(b, RomeSim) {
			err = true
		}
	case 'D':
		if !seach_byte(b, RomeSim[0:5]) {
			err = true
		}
	case 'M':
		if !seach_byte(b, RomeSim) {
			err = true
		}
	}
	if err {
		fmt.Println("Ошибка: некоректная запись операнда римскими цифрами")
		return false
	} else {
		return true
	}
}

func RTA(a string) (int, bool) {
	checkOK := true
	for i := 0; i < len(a)-1; i++ {
		if !CheckRome(a[i], a[i+1]) {
			checkOK = false
		}
	}
	if checkOK {
		Dubl := []string{"IIII", "XXXX", "CCCC",
			"IIIV", "IIV", "IIX", "IIIX",
			"XXL", "XXXL", "XXC", "XXXC", "XXD", "XXXD",
			"CCD", "CCCD", "CCM", "CCCM"}
		checkOK = !seach_sim(a, Dubl)
	}
	if checkOK {
		switch a {
		case "I":
			return 1, true
		case "II":
			return 2, true
		case "III":
			return 3, true
		case "IV":
			return 4, true
		case "V":
			return 5, true
		case "VI":
			return 6, true
		case "VII":
			return 7, true
		case "VIII":
			return 8, true
		case "IX":
			return 9, true
		case "X":
			return 10, true
		default:
			fmt.Println("Ошибка: операнд больше 10 (X)")
			return 0, false
		}
	} else {
		fmt.Println("Некоректная запись операнда римскими цифрами")
		return 0, false
	}
}

func ATR(b int) string {
	strint := strconv.Itoa(b)
	res := ""
	if len(strint) == 3 {
		res += "C"
		strint = strint[1:]
	}
	if len(strint) == 2 {
		switch strint[0] {
		case '1':
			res += "X"
			strint = strint[1:]
		case '2':
			res += "XX"
			strint = strint[1:]
		case '3':
			res += "XXX"
			strint = strint[1:]
		case '4':
			res += "XL"
			strint = strint[1:]
		case '5':
			res += "L"
			strint = strint[1:]
		case '6':
			res += "LX"
			strint = strint[1:]
		case '7':
			res += "LXX"
			strint = strint[1:]
		case '8':
			res += "LXXX"
			strint = strint[1:]
		case '9':
			res += "XC"
			strint = strint[1:]
		default:
			strint = strint[1:]
		}
	}

	switch strint[0] {
	case '1':
		res += "I"
		return res
	case '2':
		res += "II"
		return res
	case '3':
		res += "III"
		return res
	case '4':
		res += "IV"
		return res
	case '5':
		res += "V"
		return res
	case '6':
		res += "VI"
		return res
	case '7':
		res += "VII"
		return res
	case '8':
		res += "VIII"
		return res
	case '9':
		res += "IX"
		return res
	default:
		return res
	}
}

func RomeOper(chars []string) (int, bool) {
	var check bool
	var oper1, oper2 int
	oper1, check = RTA(chars[0])
	if check {
		oper2, check = RTA(chars[2])
	} else {
		return 0, false
	}
	if !check {
		return 0, false
	}
	var res int = calculation(oper1, oper2, chars[1])
	if res <= 1 {
		fmt.Println("Ошибка: результат получился меньше I")
		return 0, false
	}
	return res, true
}

func ArabOper(chars []string) (int, bool) {
	oper1, _ := strconv.Atoi(chars[0])
	oper2, _ := strconv.Atoi(chars[2])
	if oper1 > 10 || oper2 > 10 || oper1 < 1 || oper2 < 1 {
		fmt.Println("Ошибка: операнд должен быть в диапозоне от 1 до 10")
		return 0, false
	} else {
		res := calculation(oper1, oper2, chars[1])
		return res, true
	}
}

func Check(chars []string) (bool, bool) {
	massOperat := []string{"+", "-", "*", "/"}

	if len(chars) < 3 {
		fmt.Println("Ошибка: каждый операнд должен быть записан через пробел")
		return false, false
	}
	if len(chars) > 3 {
		fmt.Println("Ошибка: формат математической операции не удовлетворяет условию два операнда и один оператор (+, -, /, *)")
		return false, false
	}

	if !seach_sim(chars[1], massOperat) {
		fmt.Println("Ошибка: второй знак не оператор")
		return false, false
	}

	var Rome1, Arab1, Rome2, Arab2 bool

	Rome1, Arab1 = Check_summand(chars[0])
	if !Rome1 && !Arab1 {
		fmt.Println("Ошибка: некоректная запись первого операнда")
		return false, false
	}

	Rome2, Arab2 = Check_summand(chars[2])
	if !Rome2 && !Arab2 {
		fmt.Println("Ошибка: некоректная запись второго операнда")
		return false, false
	}

	if Rome1 == !Rome2 {
		fmt.Println("Ошибка: используются одновременно разные системы счисления.")
		return false, false
	}

	return true, Rome1
}

func Check_summand(operand string) (bool, bool) {
	var Rome, Arab bool
	romesim := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	arabsim := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-'}
	Rome = check_byte(operand, romesim)
	if Rome {
		Arab = false
	} else {
		Arab = check_byte(operand, arabsim)
	}

	return Rome, Arab
}

func calculationSTR(str string) bool {
	var resstr string = ""
	str_ := strings.Replace(str, "  ", " ", -1)
	for str_ != str {
		str = str_
		str_ = strings.Replace(str_, "  ", " ", -1)
	}
	str = strings.Trim(str_, " ")
	str = strings.Trim(str, "\r\n")
	operands := strings.Split(str, " ")
	check, Rome := Check(operands)
	if check {
		if Rome {
			res, rescheck := RomeOper(operands)
			if rescheck {
				resstr = ATR(res)
			}
		} else {
			res, rescheck := ArabOper(operands)
			if rescheck {
				resstr = strconv.Itoa(res)
			}

		}
	}
	if resstr != "" {
		fmt.Println(resstr)
		return true
	} else {
		return false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var check bool = true
	for check {
		fmt.Println("Введите выражение:")
		strarr, _ := reader.ReadString('\n')
		check = calculationSTR(strarr)
	}
}
