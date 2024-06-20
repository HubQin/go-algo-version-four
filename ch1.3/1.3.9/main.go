package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// 判断字符是否是运算符
func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("输入表达式: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	ops := []string{}  // 操作数和操作符栈
	expr := []string{} // 表达式栈

	for _, char := range input {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			expr = append(expr, string(char))
		} else if isOperator(char) {
			ops = append(ops, string(char))
		} else if char == ')' {
			// 从操作符栈中弹出运算符
			operator := ops[len(ops)-1]
			ops = ops[:len(ops)-1]

			// 从表达式栈中弹出两个操作数
			operand2 := expr[len(expr)-1]
			expr = expr[:len(expr)-1]
			operand1 := expr[len(expr)-1]
			expr = expr[:len(expr)-1]

			// 将完整的子表达式推入表达式栈
			subExpr := "( " + operand1 + " " + operator + " " + operand2 + " )"
			expr = append(expr, subExpr)
		}
	}

	// 最后表达式栈中应该只剩下一个完整的表达式
	// 输入: 1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )
	// 输出: ( ( 1 + 2 ) * ( ( 3 - 4 ) * ( 5 - 6 ) ) )
	fmt.Println(expr[0])
}
