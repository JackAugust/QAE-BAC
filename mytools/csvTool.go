package mytools

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCSV(path string, data [][]string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	// file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatalf("写入记录失败: %v", err)
		return
	}
}
func AppendCSV(path string, data [][]string) {
	// file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatalf("写入记录失败: %v", err)
		return
	}
}

// // WriteMapToCSV 将 map[string]float64 写入指定的 CSV 文件
// func WriteMapToCSV(path string, data map[string][]float64) {
// 	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
// 	// file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
// 	// file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatalf("failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	// 创建 CSV 写入器
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// 遍历 map 并写入数据行
// 	for key, value := range data {
// 		err := writer.Write([]string{key, fmt.Sprintf("%.5f", value[0]), fmt.Sprintf("%.5f", value[1])})
// 		if err != nil {
// 			log.Fatalf("写入记录失败: %v, %v", key, err)
// 			return
// 		}
// 	}

// }
func ReadCSV(path string) [][]string {
	file, err := os.Open(path) // 打开CSV文件
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()
	// 创建 CSV 读取器
	reader := csv.NewReader(file)
	// 不要求每行一致，默认是0，即每行列数必须一致
	reader.FieldsPerRecord = -1
	// 读取所有的记录
	records, err2 := reader.ReadAll()
	if err != nil {
		// fmt.Println(err2)
		log.Fatal(err2)
	}

	return records
}
