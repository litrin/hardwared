package hw_info

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type diskInfo [4]string

func isdct_get_value(field string) string {

	reg, _ := regexp.Compile(`(\S+)\s?$`)
	matched := reg.FindAllString(field, 2)

	return matched[0]
}

func get_isdc_info() []string {
	result, _ := exec.Command("/usr/bin/isdct", "show", "-intelssd").Output()
	//
	// Default results from cli like this:
	//
	//- Intel SSD DC P3600 Series CVMD408500711P6IGN -
	//
	//Bootloader : 8B1B0125
	//DevicePath : /dev/nvme0n1
	//DeviceStatus : Healthy
	//Firmware : 8DV10043
	//FirmwareUpdateAvailable : Your Intel SSD has pre-production firmware. Please contact Intel Customer Support for further assistance at the following website: http://www.intel.com/go/ssdsupport.
	//Index : 0
	//ModelNumber : INTEL SSDPE2ME016T4
	//ProductFamily : Intel SSD DC P3600 Series
	//SerialNumber : CVMD408500711P6IGN
	//`

	return strings.Split(string(result), "\n")

}

func serialnumber_to_size(sn string) int {
	size, err := strconv.Atoi(sn[8:11])
	if err != nil {
		return 0
	}

	if sn[12] == 52 {
		size *= 100
	}

	return size << 10
}
func get_nvme() []Storage {
	isdct_info := get_isdc_info()
	buff := new(diskInfo)

	storage_list := []Storage{}
	if len(isdct_info) < 12 {
		return storage_list
	}

	disk_id := 0
	for i := 1; i < len(isdct_info); i++ {
		row := i % 12
		switch row {
		case 0:
			status := STATUS_FAIL
			if buff[1] == "Healthy" {
				status = STATUS_OK
			}
			size := serialnumber_to_size(buff[2])
			storage_list = append(storage_list, Storage{Capacity: size, Path: buff[0],
				Status: status, Type: DISK_TYPE_NVME, ISN: buff[3], Id: disk_id})

			buff = new(diskInfo)
			disk_id++

			break
		case 4:
			buff[0] = isdct_get_value(isdct_info[i])
			break
		case 5:
			buff[1] = isdct_get_value(isdct_info[i])
			break
		case 9:
			buff[2] = isdct_get_value(isdct_info[i])
			break
		case 11:

			buff[3] = isdct_get_value(isdct_info[i])
			break

		}
	}

	return storage_list
}
