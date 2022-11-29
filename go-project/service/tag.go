package service

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
	"go-gin/models"
	"go-gin/pkg/export"
	file2 "go-gin/pkg/file"
	"io"
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

	err = file2.IsNotExistMkDir(export.GetExcelFullPath())
	if err != nil {
		return "", err
	}

	err = file.Save(fullPath)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func Import(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}
	rows := xlsx.GetRows("标签信息")
	for idxRow, row := range rows {
		if idxRow < 1 {
			continue
		}
		var data []string
		for _, cell := range row {
			data = append(data, cell)
		}
		models.AddTag(data[1], 1, data[2])
	}
	return nil
}
