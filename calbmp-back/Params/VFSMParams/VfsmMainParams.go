package VFSMParams

import (
	"calbmp-back/Params/ResultParams"
)

type VfsmMainParams struct {
	CreatedAt string                        `json:"created_at"`
	IkwParam  IkwParams                     `json:"ikw_param"`
	IgrParam  IgrParams                     `json:"igr_param"`
	IsoParam  IsoParams                     `json:"iso_param"`
	IrnParam  IrnParams                     `json:"irn_param"`
	IsdParam  IsdParams                     `json:"isd_param"`
	IroParam  IroParams                     `json:"iro_param"`
	IwqParam  IwqParams                     `json:"iwq_param"`
	TextRes   ResultParams.TextResultParams `json:"text_res"`
}
