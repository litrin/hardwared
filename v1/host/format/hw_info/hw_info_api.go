package hw_info

type pcieConnect struct {
	Id       int    `json:"id"`
	Status   string `json:"status"`
	Cable_id string `json:"cableId"`
}

type nic struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	IP     string `json:"ip"`
	MAC    string `json:"mac"`
	MTU    int    `json:"mtu"`
	Name   string `json:"name"`
}

type memory struct {
	Id        int    `json:"id"`
	Status    string `json:"status"`
	Capacity  int    `json:"capacity"`
	Frequency string `json:"frequency"`
}

type CPUObject struct {
	Id        int    `json:"id"`
	Status    string `json:"status"`
	Mode      string `json:"mode"`
	CoreNum   int    `json:"coreNum"`
	Frequency string `json:"frequency"`
}

type Hosts struct {
	Id           int           `json:"id"`
	Status       string        `json:"status"`
	HostName     string        `json:"hostName"`
	Cpus         []CPUObject   `json:"cpus"`
	Memorys      []memory      `json:"memorys"`
	NICs         []nic         `json:"nics"`
	Strorage     []Storage     `json:"storage"`
	PcieConnects []pcieConnect `json:"pcieConnections"`
}

type _data struct {
	Data []Hosts `json:"hosts"`
}

type _page_framework struct {
	Status     string `json:"status"`
	ErrMessage string `json:"errMsg"`
	Data       _data  `json:"data"`
}

type Page_framework struct {
	Result _page_framework `json:"result"`
}

type Storage struct {
	Id       int    `json:"id"`
	Capacity int    `json:"capacity"`
	ISN      string `json:"isn"`
	Type     string `json:"type"`
	Status   string `json:"status"`
	Path     string `json:"driverPath"`
}
