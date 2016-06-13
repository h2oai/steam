package az

type Principal struct {
	Id          int64
	WorkgroupId int64
}

func NewPrincipal(id, workgroupId int64) *Principal {
	return &Principal{
		id,
		workgroupId,
	}
}
