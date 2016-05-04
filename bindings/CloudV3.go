package bindings

type CloudV3 struct {
	*RequestSchema
	/** skip_ticks */
	SkipTicks bool `json:"skip_ticks"`

	/** version */
	Version string `json:"version"`

	/** branch_name */
	BranchName string `json:"branch_name"`

	/** build_number */
	BuildNumber string `json:"build_number"`

	/** Node index number cloud status is collected from (zero-based) */
	NodeIdx int32 `json:"node_idx"`

	/** cloud_name */
	CloudName string `json:"cloud_name"`

	/** cloud_size */
	CloudSize int32 `json:"cloud_size"`

	/** cloud_uptime_millis */
	CloudUptimeMillis int64 `json:"cloud_uptime_millis"`

	/** cloud_healthy */
	CloudHealthy bool `json:"cloud_healthy"`

	/** Nodes reporting unhealthy */
	BadNodes int32 `json:"bad_nodes"`

	/** Cloud voting is stable */
	Consensus bool `json:"consensus"`

	/** Cloud is accepting new members or not */
	Locked bool `json:"locked"`

	/** Cloud is in client mode. */
	IsClient bool `json:"is_client"`

	/** nodes */
	Nodes []NodeV3 `json:"nodes"`

	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string: "" `json:"_exclude_fields"`
	*/
}

func newCloudV3() *CloudV3 {
	return &CloudV3{
		SkipTicks:         false,
		Version:           "",
		BranchName:        "",
		BuildNumber:       "",
		NodeIdx:           0,
		CloudName:         "",
		CloudSize:         0,
		CloudUptimeMillis: 0,
		CloudHealthy:      false,
		BadNodes:          0,
		Consensus:         false,
		Locked:            false,
		IsClient:          false,
		Nodes:             nil,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
