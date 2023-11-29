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
	Clusters []Cluster `json:"clusters"`
}

type Cluster struct {
	Name     string          `json:"name"`
	Metadata ClusterMetadata `json:"metadata"`
	Spec     ClusterSpec     `json:"spec"`
}

type ClusterMetadata struct {
	Labels map[string]string `json:"labels"`
}

type ClusterSpec struct {
	Packages []PackageVariant `json:"packages"`
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
								"name": "nginx-package",
								"repo": "lkass-packages",
								"package": "pkg-example-nginx",
								"revision": "v5"
							}
						]
					}
				},
				{
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
								"name": "prometheus-package",
								"repo": "lkass-packages",
								"package": "pkg-example-prometheus",
								"revision": "v5"
							}
						]
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

	for _, cluster := range lc.Spec.Clusters {
		wCluster := WorkloadCluster{
			APIVersion: "example/v1alpha1",
			Kind:       "WorkloadCluster",
			Metadata: Metadata{
				Name: lc.Metadata.Name,
			},
			Spec: WorkloadSpec{
				ClusterName: cluster.Name,
				Labels:      cluster.Metadata.Labels,
			},
		}

		pVariant := PackageVariant{
			Name:     cluster.Spec.Packages[0].Name,
			Repo:     cluster.Spec.Packages[0].Repo,
			Package:  cluster.Spec.Packages[0].Package,
			Revision: cluster.Spec.Packages[0].Revision,
		}

		pOutput := generatePackageVariantCRD(pVariant, cluster.Name)

		wOutput, _ := json.MarshalIndent(wCluster, "", "  ")
		fmt.Println(string(wOutput))
		fmt.Println(pOutput)
	}
}

func generatePackageVariantCRD(p PackageVariant, clusterName string, lc LogicalCluster) string {
	crd := fmt.Sprintf(`
apiVersion: config.porch.kpt.dev/v1alpha2
kind: PackageVariant
metadata:
  name: %s
spec:
  upstream:
    repo: %s
    package: %s
    revision: %s
  downstream:
    package: example-configsync
    repo: %s
  pipeline:
    mutators:
    - image: gcr.io/kpt-fn/set-annotations:v0.1.4
      configMap:
        nephio.org/cluster-name: %s
`, p.Name, p.Repo, p.Package, p.Revision, clusterName, lc.Metadata.Name)

	return crd
}

type WorkloadCluster struct {
	APIVersion string       `json:"apiVersion"`
	Kind       string       `json:"kind"`
	Metadata   Metadata     `json:"metadata"`
	Spec       WorkloadSpec `json:"spec"`
}

type WorkloadSpec struct {
	ClusterName string            `json:"clusterName"`
	Labels      map[string]string `json:"labels"`
}
