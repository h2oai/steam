package bindings

type NodeV3 struct {
	*Schema
	/** IP */
	H2o string `json:"h2o"`

	/** IP address and port in the form a.b.c.d:e */
	IpPort string `json:"ip_port"`

	/** (now-last_ping)<HeartbeatThread.TIMEOUT */
	Healthy bool `json:"healthy"`

	/** Time (in msec) of last ping */
	LastPing int64 `json:"last_ping"`

	/** PID */
	Pid int32 `json:"pid"`

	/** num_cpus */
	NumCpus int32 `json:"num_cpus"`

	/** cpus_allowed */
	CpusAllowed int32 `json:"cpus_allowed"`

	/** nthreads */
	Nthreads int32 `json:"nthreads"`

	/** System load; average #runnables/#cores */
	SysLoad float32 `json:"sys_load"`

	/** System CPU percentage used by this H2O process in last interval */
	MyCpuPct int32 `json:"my_cpu_pct"`

	/** System CPU percentage used by everything in last interval */
	SysCpuPct int32 `json:"sys_cpu_pct"`

	/** Data on Node memory */
	MemValueSize int64 `json:"mem_value_size"`

	/** Temp (non Data) memory */
	PojoMem int64 `json:"pojo_mem"`

	/** Free heap */
	FreeMem int64 `json:"free_mem"`

	/** Maximum memory size for node */
	MaxMem int64 `json:"max_mem"`

	/** Size of data on node's disk */
	SwapMem int64 `json:"swap_mem"`

	/** #local keys */
	NumKeys int32 `json:"num_keys"`

	/** Free disk */
	FreeDisk int64 `json:"free_disk"`

	/** Max disk */
	MaxDisk int64 `json:"max_disk"`

	/** Active Remote Procedure Calls */
	RpcsActive int32 `json:"rpcs_active"`

	/** F/J Thread count, by priority */
	Fjthrds []int16 `json:"fjthrds"`

	/** F/J Task count, by priority */
	Fjqueue []int16 `json:"fjqueue"`

	/** Open TCP connections */
	TcpsActive int32 `json:"tcps_active"`

	/** Open File Descripters */
	OpenFds int32 `json:"open_fds"`

	/** Linpack GFlops */
	Gflops float64 `json:"gflops"`

	/** Memory Bandwidth */
	MemBw float64 `json:"mem_bw"`
}

func newNodeV3() *NodeV3 {
	return &NodeV3{
		H2o:          "",
		IpPort:       "",
		Healthy:      false,
		LastPing:     0,
		Pid:          0,
		NumCpus:      0,
		CpusAllowed:  0,
		Nthreads:     0,
		SysLoad:      0.0,
		MyCpuPct:     0,
		SysCpuPct:    0,
		MemValueSize: 0,
		PojoMem:      0,
		FreeMem:      0,
		MaxMem:       0,
		SwapMem:      0,
		NumKeys:      0,
		FreeDisk:     0,
		MaxDisk:      0,
		RpcsActive:   0,
		Fjthrds:      nil,
		Fjqueue:      nil,
		TcpsActive:   0,
		OpenFds:      0,
		Gflops:       0.0,
		MemBw:        0.0,
	}
}
