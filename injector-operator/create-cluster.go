package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	// Kiểm tra xem có đủ tham số không
	// if len(os.Args) < 2 {
	// 	fmt.Println("Sử dụng: go run create_and_apply_cluster.go <tên-cluster>")
	// 	return
	// }

	// Lấy tên cluster từ tham số dòng lệnh
	// clusterName := os.Args[1]
	clusterName := "cluster-test"

	// Tạo tên tệp YAML từ tên cluster
	yamlFileName := fmt.Sprintf("/home/test/%s.yaml", clusterName)

	// Tạo lệnh để tạo cluster
	createCmd := exec.Command("clusterctl", "generate", "cluster", clusterName,
		"--flavor", "development",
		"--kubernetes-version", "v1.28.0",
		"--control-plane-machine-count=1",
		"--worker-machine-count=1")

	// Mở tệp để lưu đầu ra
	yamlFile, err := os.Create(yamlFileName)
	if err != nil {
		fmt.Printf("Lỗi khi mở tệp YAML: %v\n", err)
		return
	}
	defer yamlFile.Close()

	// Đặt đầu ra của lệnh để viết vào tệp YAML
	createCmd.Stdout = yamlFile
	createCmd.Stderr = os.Stderr

	// Chạy lệnh tạo cluster
	err = createCmd.Run()
	if err != nil {
		fmt.Printf("Lỗi khi tạo cluster: %v\n", err)
		return
	}

	fmt.Printf("Đã tạo cluster '%s' và lưu vào tệp YAML '%s'\n", clusterName, yamlFileName)

	// Tạo lệnh để áp dụng cluster
	applyCmd := exec.Command("kubectl", "apply", "-f", yamlFileName)

	// Chạy lệnh áp dụng cluster
	err = applyCmd.Run()
	if err != nil {
		fmt.Printf("Lỗi khi áp dụng cluster: %v\n", err)
		return
	}

	fmt.Printf("Đã áp dụng cluster '%s'\n", clusterName)

	// Đợi 20 giây
	fmt.Println("Đang đợi 20 giây...")
	time.Sleep(20 * time.Second)

	// ------------- lưu kubeconfig

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

	// Đợi 60 giây
	fmt.Println("Đang đợi 60 giây...")
	time.Sleep(60 * time.Second)

	// apply cni:

	// Tạo lệnh để áp dụng Calico CNI vào cluster
	applyCniCmd := exec.Command("kubectl", "--kubeconfig="+kubeconfigPath, "apply", "-f", "https://raw.githubusercontent.com/projectcalico/calico/v3.26.1/manifests/calico.yaml")

	// Chạy lệnh áp dụng Calico CNI
	err2 := applyCniCmd.Run()
	if err != nil {
		fmt.Printf("Lỗi khi áp dụng Calico CNI: %v\n", err2)
		return
	}

	fmt.Printf("Đã áp dụng Calico CNI cho cluster '%s'\n", clusterName)

}
