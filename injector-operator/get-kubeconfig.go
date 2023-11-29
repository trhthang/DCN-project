package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Kiểm tra xem có đủ tham số không
	// if len(os.Args) < 2 {
	// 	fmt.Println("Sử dụng: go run get_kubeconfig.go <tên-cluster>")
	// 	return
	// }

	// Lấy tên cluster từ tham số dòng lệnh
	clusterName := "cluster-test"

	// Tạo đường dẫn đầy đủ tới tệp kubeconfig
	kubeconfigPath := filepath.Join("/home/test", fmt.Sprintf("%s.kubeconfig", clusterName))

	// Tạo lệnh để lấy kubeconfig của cluster
	getKubeconfigCmd := exec.Command("clusterctl", "get", "kubeconfig", clusterName)

	// Chạy lệnh để lấy kubeconfig
	kubeconfigOutput, err := getKubeconfigCmd.Output()
	if err != nil {
		fmt.Printf("Lỗi khi lấy kubeconfig: %v\n", err)
		return
	}

	// Lưu nội dung kubeconfig vào tệp
	err = os.WriteFile(kubeconfigPath, kubeconfigOutput, 0644)
	if err != nil {
		fmt.Printf("Lỗi khi lưu kubeconfig: %v\n", err)
		return
	}

	fmt.Printf("Đã lấy kubeconfig của cluster '%s' và lưu vào tệp '%s'\n", clusterName, kubeconfigPath)
}
