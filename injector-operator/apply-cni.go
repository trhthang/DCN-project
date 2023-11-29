package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func main() {
	// Kiểm tra xem có đủ tham số không
	// if len(os.Args) < 2 {
	// 	fmt.Println("Sử dụng: go run apply_cni.go <tên-cluster>")
	// 	return
	// }

	// Lấy tên cluster từ tham số dòng lệnh
	clusterName := "cluster-test"

	// Tạo đường dẫn đầy đủ tới tệp kubeconfig
	kubeconfigPath := filepath.Join("/home/test", fmt.Sprintf("%s.kubeconfig", clusterName))

	// Tạo lệnh để áp dụng Calico CNI vào cluster
	applyCniCmd := exec.Command("kubectl", "--kubeconfig="+kubeconfigPath, "apply", "-f", "https://raw.githubusercontent.com/projectcalico/calico/v3.26.1/manifests/calico.yaml")

	// Chạy lệnh áp dụng Calico CNI
	err := applyCniCmd.Run()
	if err != nil {
		fmt.Printf("Lỗi khi áp dụng Calico CNI: %v\n", err)
		return
	}

	fmt.Printf("Đã áp dụng Calico CNI cho cluster '%s'\n", clusterName)
}
