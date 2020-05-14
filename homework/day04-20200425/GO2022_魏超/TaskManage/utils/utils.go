package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/olekukonko/tablewriter"
)

// Input 常规信息交互
func Input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	return strings.TrimSpace(string(data))
}

// PWInput 密码信息交互
func PWInput(prompt string) string {
	fmt.Print(prompt)
	password, _ := gopass.GetPasswd()
	return strings.TrimSpace(string(password))
}

// TableFormat 格式化为表格模式输出
func TableFormat(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, d := range data {
		table.Append(d)
	}
	table.Render()
}

// PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
