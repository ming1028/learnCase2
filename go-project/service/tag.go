package service

import (
	"github.com/tealeg/xlsx"
	"go-gin/pkg/export"
	"strconv"
	"time"
)

func Export() (string, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("标签信息")
	if err != nil {
		return "", err
	}
	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow() // 添加一行
	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}
	// 添加数据 循环拼接数据切片，增加AddRow, AddCell

	time := strconv.Itoa(int(time.Now().Unix()))
	fileName := "tags-" + time + ".xlsx"

	fullPath := export.GetExcelFullPath() + fileName
	err = file.Save(fullPath)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
