package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Thay đổi giá trị này thành tên của instance bạn muốn lấy thông tin
	instanceName := "capi-quickstart"

	// Chạy lệnh kubectl để lấy thông tin chi tiết của một instance cụ thể
	cmd := exec.Command("kubectl", "describe", "cluster", instanceName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running kubectl describe cluster %s: %v\n", instanceName, err)
		return
	}

	instanceOutput := string(output)
	// Chuỗi cần tìm kiếm
	searchString := "Phase:"

	// Tìm vị trí của chuỗi cần tìm kiếm
	index := strings.Index(instanceOutput, searchString)
	if index == -1 {
		fmt.Println("Không tìm thấy chuỗi cần tìm kiếm")
		return
	}

	// Cắt chuỗi từ vị trí của chuỗi cần tìm kiếm đến hết
	substring := instanceOutput[index+len(searchString):]

	// Xóa các khoảng trắng thừa
	substring = strings.TrimSpace(substring)

	// In ra giá trị
	fmt.Println("Giá trị của status.conditions.phase:", substring)
}
