package mytools

import (
	"fmt"
	"log"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func In(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}
func Index(target string, list []string) int {
	index := -1
	for i, s := range list {
		if s == target {
			index = i
			break
		}
	}

	if index >= 0 {
		return index
	}
	return -1
}

func SavePic(data map[float64]float64, path string) {
	// 准备散点图的数据
	var pts plotter.XYs
	for x, y := range data {
		pts = append(pts, plotter.XY{X: x, Y: y})
	}

	// 创建一个新的图表
	p := plot.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 添加散点图
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal("plotter.NewScatter失败: ", err)
	}
	p.Add(s)

	// 设置标题
	p.Title.Text = "Fruit Prices"
	p.X.Label.Text = "Index"
	p.Y.Label.Text = "Price"

	// 保存为 PNG 文件
	if err := p.Save(5*vg.Inch, 5*vg.Inch, path); err != nil {
		log.Fatal("保存失败：", err)
	}

	fmt.Println("散点图已保存")
}
