package main

import (
	"encoding/json"
	"fmt"
)

type LogicalCluster struct {
	APIVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   Metadata    `json:"metadata"`
	Spec       LogicalSpec `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type LogicalSpec struct {
	Clusters []ClusterWrapper `json:"clusters"`
}

type ClusterWrapper struct {
	Cluster Cluster `json:"cluster"`
}

type Cluster struct {
	ClusterID int             `json:"id"`
	Name      string          `json:"name"`
	Metadata  ClusterMetadata `json:"metadata"`
	Spec      ClusterSpec     `json:"spec"`
}

type ClusterMetadata struct {
	Labels map[string]string `json:"labels"`
}

type ClusterSpec struct {
	Packages []PackageWrapper `json:"packages"`
}

type PackageWrapper struct {
	PackageVariant PackageVariant `json:"packageVariant"`
}

type PackageVariant struct {
	Name     string `json:"name"`
	Repo     string `json:"repo"`
	Package  string `json:"package"`
	Revision string `json:"revision"`
}

func main() {
	input := `{
		"apiVersion": "example/v1alpha1",
		"kind": "LogicalCluster",
		"metadata": {
			"name": "mec-cluster"
		},
		"spec": {
			"clusters": [
				{
					"cluster": {
						"id": 1,
						"name": "RegionalCluster",
						"metadata": {
							"labels": {
								"type": "regional",
								"region": "us-west1"
							}
						},
						"spec": {
							"packages": [
								{
									"packageVariant": {
										"name": "nginx-package",
										"repo": "lkass-packages",
										"package": "pkg-example-nginx",
										"revision": "v5"
									}
								}
							]
						}
					}
				},
				{
					"cluster": {
						"id": 2,
						"name": "Edge01",
						"metadata": {
							"labels": {
								"type": "edge",
								"region": "us-west1"
							}
						},
						"spec": {
							"packages": [
								{
									"packageVariant": {
										"name": "prometheus-package",
										"repo": "lkass-packages",
										"package": "pkg-example-prometheus",
										"revision": "v5"
									}
								}
							]
						}
					}
				}
			]
		}
	}`

	var lc LogicalCluster
	if err := json.Unmarshal([]byte(input), &lc); err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, clusterWrapper := range lc.Spec.Clusters {
		cluster := clusterWrapper.Cluster
		fmt.Println("Cluster ID:", cluster.ClusterID)
		fmt.Println("Cluster Name:", cluster.Name)
		fmt.Println("Cluster Labels:", cluster.Metadata.Labels)
		fmt.Println("Cluster Spec:")
		fmt.Printf("%+v\n", cluster.Spec)
		fmt.Println("------------------")
	}
}
