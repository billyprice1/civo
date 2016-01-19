package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/google/go-querystring/query"
	"github.com/jeffail/gabs"
)

type QuotaParams struct {
	Account           string `url:"-"`
	InstanceCount     string `url:"instance_count_limit"`
	CpuCore           string `url:"cpu_core_limit"`
	RamMB             string `url:"ram_mb_limit"`
	DiskGB            string `url:"disk_gb_limit"`
	DiskVolumeCount   string `url:"disk_volume_count_limit"`
	DiskSnapshotCount string `url:"disk_snapshot_count_limit"`
	PublicIPAddress   string `url:"public_ip_address_limit"`
	SubnetCount       string `url:"subnet_count_limit"`
	NetworkCount      string `url:"network_count_limit"`
	SecurityGroup     string `url:"security_group_limit"`
	SecurityGroupRule string `url:"security_group_rule_limit"`
	PortCount         string `url:"port_count_limit"`
}

func QuotaGet(account string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/quota?username="+account, HTTPGet, "")
}

func QuotaSet(params QuotaParams) (json *gabs.Container, err error) {
	v, _ := query.Values(params)
	return makeJSONCall(config.URL()+"/v1/quota/"+params.Account, HTTPPut, v.Encode())
}
