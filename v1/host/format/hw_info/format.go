package hw_info

import (
	"encoding/json"
	hwcpu "github.com/mickep76/hwinfo/cpu"
	hwnet "github.com/mickep76/hwinfo/interfaces"
	hwmem "github.com/mickep76/hwinfo/memory"
	"os/exec"
)

const (
	STATUS_OK      = "healthy"
	STATUS_FAIL    = "Fail"
	DISK_TYPE_NVME = "Nvme"
	UNKNOWN        = "unknown"
)

func Test() string {
	b, _ := json.Marshal(All())
	return string(b)
}

func get_cpu() []CPUObject {
	c := hwcpu.New()
	c.Update()

	cpu_info := c.GetData()

	frg := cpu_info.Model[len(cpu_info.Model)-8:]

	i := 0
	cpu_count := cpu_info.Sockets
	cpu_list := make([]CPUObject, cpu_count)

	for i = 0; i < cpu_count; i++ {
		cpu_list[i] =
			CPUObject{Id: i, Frequency: frg, Status: STATUS_OK, CoreNum: cpu_info.Physical / cpu_count, Mode: cpu_info.Model}
	}

	return cpu_list
}

func get_mem() []memory {
	c := hwmem.New()
	c.Update()

	mem_info := c.GetData()

	m := memory{Id: 0, Status: STATUS_OK, Capacity: mem_info.TotalKB >> 10, Frequency: UNKNOWN}

	return []memory{m}

}

func get_nic() []nic {

	c := hwnet.New()
	c.Update()

	net_info := c.GetData()
	nics := make([]nic, len(net_info))

	ip := ""
	for i, v := range net_info {

		if len(v.IPAddr) > 0 {
			ip = v.IPAddr[0]
		} else {
			ip = UNKNOWN
		}
		nics[i] = nic{Id: i, Status: STATUS_OK, MAC: v.HWAddr, MTU: v.MTU, Name: v.Name, IP: ip}

	}

	return nics
}

func get_hostname() string {
	s := exec.Command("hostname")
	d, _ := s.Output()
	sD := string(d)
	sD = sD[:len(sD)-1]
	return sD
}

func All() *Page_framework {

	host_info := []Hosts{{Id: 0, Status: STATUS_OK, HostName: get_hostname(),
		Cpus: get_cpu(), Memorys: get_mem(), NICs: get_nic(), PcieConnects: get_aic_sn(),
		Strorage: get_nvme()}}

	return NewPageFramework(host_info, "")

}

func Combind(channel chan Page_framework, length int) *Page_framework {

	hosts := make([]Hosts, length)

	var content Page_framework
	for index := 0; index < length; index++ {
		content = <-channel

		hosts[index] = content.Result.Data.Data[0]
		hosts[index].Id = index
	}

	return NewPageFramework(hosts, "")
}

func NewPageFramework(hosts []Hosts, err string) *Page_framework {
	if err == "" {
		return &Page_framework{Result: _page_framework{Status: STATUS_OK, Data: _data{hosts}, ErrMessage: ""}}
	}

	return &Page_framework{Result: _page_framework{Status: STATUS_FAIL, Data: _data{hosts}, ErrMessage: err}}
}
