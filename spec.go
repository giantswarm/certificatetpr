package certificatetpr

import "github.com/giantswarm/certificatetpr/spec"

type Searcher interface {
	SearchCluster(clusterID string) (Cluster, error)
	SearchMonitoring(clusterID string) (Monitoring, error)
}

type Spec struct {
	AllowBareDomains bool               `json:"allowBareDomains" yaml:"allowBareDomains"`
	AltNames         []string           `json:"altNames" yaml:"altNames"`
	ClusterComponent string             `json:"clusterComponent" yaml:"clusterComponent"`
	ClusterID        string             `json:"clusterID" yaml:"clusterID"`
	CommonName       string             `json:"commonName" yaml:"commonName"`
	IPSANs           []string           `json:"ipSans" yaml:"ipSans"`
	Organizations    []string           `json:"organizations" yaml:"organizations"`
	TTL              string             `json:"ttl" yaml:"ttl"`
	VersionBundle    spec.VersionBundle `json:"versionBundle" yaml:"versionBundle"`
}
