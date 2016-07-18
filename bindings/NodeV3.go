package bindings

import "encoding/json"

type NodeV3 struct {
	*Schema
	/** IP */
	H2o string `json:"h2o,omitempty"`
	/** IP address and port in the form a.b.c.d:e */
	IpPort string `json:"ip_port,omitempty"`
	/** (now-last_ping)<HeartbeatThread.TIMEOUT */
	Healthy bool `json:"healthy,omitempty"`
	/** Time (in msec) of last ping */
	LastPing int64 `json:"last_ping,omitempty"`
	/** PID */
	Pid int32 `json:"pid,omitempty"`
	/** num_cpus */
	NumCpus int32 `json:"num_cpus,omitempty"`
	/** cpus_allowed */
	CpusAllowed int32 `json:"cpus_allowed,omitempty"`
	/** nthreads */
	Nthreads int32 `json:"nthreads,omitempty"`
	/** System load; average #runnables/#cores */
	SysLoad float32 `json:"sys_load,omitempty"`
	/** System CPU percentage used by this H2O process in last interval */
	MyCpuPct int32 `json:"my_cpu_pct,omitempty"`
	/** System CPU percentage used by everything in last interval */
	SysCpuPct int32 `json:"sys_cpu_pct,omitempty"`
	/** Data on Node memory */
	MemValueSize int64 `json:"mem_value_size,omitempty"`
	/** Temp (non Data) memory */
	PojoMem int64 `json:"pojo_mem,omitempty"`
	/** Free heap */
	FreeMem int64 `json:"free_mem,omitempty"`
	/** Maximum memory size for node */
	MaxMem int64 `json:"max_mem,omitempty"`
	/** Size of data on node's disk */
	SwapMem int64 `json:"swap_mem,omitempty"`
	/** #local keys */
	NumKeys int32 `json:"num_keys,omitempty"`
	/** Free disk */
	FreeDisk int64 `json:"free_disk,omitempty"`
	/** Max disk */
	MaxDisk int64 `json:"max_disk,omitempty"`
	/** Active Remote Procedure Calls */
	RpcsActive int32 `json:"rpcs_active,omitempty"`
	/** F/J Thread count, by priority */
	Fjthrds []int16 `json:"fjthrds,omitempty"`
	/** F/J Task count, by priority */
	Fjqueue []int16 `json:"fjqueue,omitempty"`
	/** Open TCP connections */
	TcpsActive int32 `json:"tcps_active,omitempty"`
	/** Open File Descripters */
	OpenFds int32 `json:"open_fds,omitempty"`
	/** Linpack GFlops */
	Gflops float64 `json:"gflops,omitempty"`
	/** Memory Bandwidth */
	MemBw float64 `json:"mem_bw,omitempty"`
}

func NewNodeV3() *NodeV3 {
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
		Schema:       &Schema{},
	}
}

// UnmarshalJSON to handle possible Infinity and NaN values
func (o *NodeV3) UnmarshalJSON(data []byte) error {
	type Alias NodeV3
	aux := &struct {
		SysLoad interface{} `json:"sys_load,omitempty"`
		Gflops  interface{} `json:"gflops,omitempty"`
		MemBw   interface{} `json:"mem_bw,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	o.SysLoad = jsonToFloat(aux.SysLoad)
	o.Gflops = jsonToDoubl(aux.Gflops)
	o.MemBw = jsonToDoubl(aux.MemBw)
	return nil
}
