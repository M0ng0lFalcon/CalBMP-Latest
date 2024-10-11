package VFSMParams

type SpecifyLengthParam struct {
	VfsmId      int    `json:"vfsm_id"`
	CreatedTime string `json:"created_time"`
	Upper       int    `json:"upper"`
	Lower       int    `json:"lower"`
	Increment   int    `json:"increment"`
}
