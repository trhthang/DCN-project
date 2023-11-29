	// fmt.Println("Đang đợi 60 giây...")
	// time.Sleep(60 * time.Second)

	// // apply cni:

	// // Tạo lệnh để áp dụng Calico CNI vào cluster
	// applyCniCmd := exec.Command("kubectl", "--kubeconfig="+kubeconfigPath, "apply", "-f", "https://raw.githubusercontent.com/projectcalico/calico/v3.26.1/manifests/calico.yaml")

	// // Chạy lệnh áp dụng Calico CNI
	// err2 := applyCniCmd.Run()
	// if err != nil {
	// 	fmt.Printf("Lỗi khi áp dụng Calico CNI: %v\n", err2)
	// 	return
	// }

	// fmt.Printf("Đã áp dụng Calico CNI cho cluster '%s'\n", clusterName)