// 代码生成时间: 2025-08-10 23:05:33
package main

import (
    "encoding/json"
    "fmt"
    "github.com/tealeg/xlsx"
    "net/http"
    "os"
)

// ExcelGenerator 结构体，用于生成Excel表格
type ExcelGenerator struct {
    // 这里可以添加更多的字段来扩展ExcelGenerator的功能
}

// NewExcelGenerator 创建一个ExcelGenerator实例
func NewExcelGenerator() *ExcelGenerator {
    return &ExcelGenerator{}
}

// GenerateExcel 生成Excel文件
func (e *ExcelGenerator) GenerateExcel(data interface{}, filename string) error {
    // 创建一个新的Excel文件
    file := xlsx.NewFile()
    
    // 创建一个sheet
    sheet, err := file.AddSheet("Sheet1")
    if err != nil {
        return err
    }
    
    // 将数据转换为JSON，以便写入Excel
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }
    
    // 解析JSON数据，这里需要根据实际数据结构来解析
    // 假设数据是一个数组，每个元素是一个map，包含Excel的列信息
    var rows []interface{}
    err = json.Unmarshal(jsonData, &rows)
    if err != nil {
        return err
    }
    
    // 写入Excel标题行
    for _, row := range rows {
        r := sheet.AddRow()
        for key := range row.(map[string]interface{}) {
            cell := r.AddCell()
            cell.Value = key
        }
    }
    
    // 写入数据行
    for _, row := range rows {
        r := sheet.AddRow()
        for _, cellValue := range row.(map[string]interface{}) {
            cell := r.AddCell()
            cell.Value = fmt.Sprintf("%v", cellValue)
        }
    }
    
    // 保存Excel文件
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    
    if err := file.Save(f); err != nil {
        return err
    }
    
    return nil
}

// ExcelHandler 处理生成Excel的HTTP请求
func ExcelHandler(w http.ResponseWriter, r *http.Request) {
    // 检查请求方法
    if r.Method != http.MethodPost {
        http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
        return
    }
    
    // 解析请求体，这里假设请求体是JSON格式
    var data interface{}
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    // 创建ExcelGenerator实例
    generator := NewExcelGenerator()
    
    // 生成Excel文件
    err = generator.GenerateExcel(data, "example.xlsx")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // 设置响应头
    w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    w.Header().Set("Content-Disposition", "attachment; filename="example.xlsx"")
    
    // 读取并发送Excel文件
    http.ServeFile(w, r, "example.xlsx")
    
    // 删除临时文件
    os.Remove("example.xlsx")
}

func main() {
    // 设置路由
    http.HandleFunc("/generate", ExcelHandler)
    
    // 启动服务器
    fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}