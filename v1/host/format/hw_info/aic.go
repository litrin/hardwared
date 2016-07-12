package hw_info

import (
	"os/exec"
	"strconv"
	"strings"
)

func get_aic_sn() []pcieConnect {
	result, err := exec.Command("ipmitool", "raw", "0x3a", "0x16", "0x00", "0xc4", "0x10").Output()
	if err != nil {
		return nil
	}

	sn_list := strings.Split(string(result), " ")
	sn := ""
	buff := int64(0)
	for _, v := range sn_list {
		if len(v) == 0 {
			continue
		}
		buff, _ = strconv.ParseInt(v, 16, 0)

		if buff > 0x20 && buff < 0x7E {
			sn += string(buff)
		}
	}

	return []pcieConnect{pcieConnect{Status: STATUS_OK, Id: 0, Cable_id: sn}}
}
