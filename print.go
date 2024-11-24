package gobtphelper

import (
	"fmt"
	"strings"
)

// TableData 表示表格的数据
type TableData struct {
	Headers []string
	Rows    [][]string
}

// PrintTable 打印表格
func PrintTable(data TableData) {
	// 计算列宽
	colWidths := make([]int, len(data.Headers))
	for i, header := range data.Headers {
		colWidths[i] = len(header)
	}
	for _, row := range data.Rows {
		for i, cell := range row {
			if len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	// 打印表头
	PrintRow(data.Headers, colWidths)
	//PrintSeparator(colWidths)

	// 打印数据行
	for _, row := range data.Rows {
		PrintRow(row, colWidths)
	}

	// 打印底部边框
	PrintSeparator(colWidths)
}

// PrintRow 打印一行
func PrintRow(row []string, colWidths []int) {
	line := "+"
	for _, width := range colWidths {
		line += strings.Repeat("-", width+2) + "+"
	}
	fmt.Println(line)

	line = "|"
	for i, cell := range row {
		line += fmt.Sprintf(" %-"+fmt.Sprintf("%d", colWidths[i])+"s |", cell)
	}
	fmt.Println(line)
}

// PrintSeparator 打印分隔符
func PrintSeparator(colWidths []int) {
	line := "+"
	for _, width := range colWidths {
		line += strings.Repeat("-", width+2) + "+"
	}
	fmt.Println(line)
}
