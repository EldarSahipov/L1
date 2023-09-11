package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println(method1(str))
	fmt.Println(method2(str))
	fmt.Println(method3(str))
}

/*
 Если он находит два одинаковых символа
 (кроме тех, которые находятся на одной и той же позиции), \
 то он устанавливает флаг isUnique в false
*/
// O(n^2)
func method1(str string) bool {
	strLower := strings.ToLower(str)

	isUnique := true

	for i, v1 := range strLower { //
		for j, v2 := range strLower {
			if i != j && v1 == v2 {
				isUnique = false
			}
		}
	}
	return isUnique
}

// O(n)
/*
Этот метод использует карту mapRune, чтобы отслеживать уникальные символы в строке.
*/
func method2(str string) bool {
	strLower := strings.ToLower(str)
	mapRune := make(map[rune]bool)

	for _, char := range strLower {
		if mapRune[char] { // Если символ уже встречался, функция возвращает false, иначе символ добавляется в карту.
			return false
		}
		mapRune[char] = true
	}
	return true
}

/*
Этот метод использует сортировку для упорядочивания символов в строке.
Затем он проверяет, что нет соседних символов, которые идентичны.
Если он находит такие соседние символы, он возвращает false.
Этот метод требует O(n * log(n)) времени из-за сортировки
*/
// O(n * log(n))
func method3(str string) bool {
	strLower := strings.ToLower(str)

	sliceRunes := strings.Split(strLower, "")
	sort.Strings(sliceRunes)

	for i := 1; i < len(sliceRunes); i++ {
		if sliceRunes[i-1] == sliceRunes[i] {
			return false
		}
	}

	return true
}

/*
Method2 (использующий карту) является наиболее эффективным вариантом с точки зрения временной сложности,
так как его сложность O(n). Method3 (использующий сортировку) может быть полезным,
если требуется дополнительная обработка после сортировки символов.
Method1, хотя он работает, менее эффективен из-за квадратичной сложности времени.
*/
