package InputFileGenerateUtil

import (
	"calbmp-back/model"
	"calbmp-back/util/StringUtil"
	"math"
)

// GetRecord3BasicValues : get uslek, usles, uslep, IREG, slope, slopeLength
func GetRecord3BasicValues(
	Chorizon model.ChorizonModel,
	RainDistributionIreg model.RainDistributionIreg,
	slopeLengthFlag bool,
	knowSlope bool,
	kSlope float64,
) (
	uslek string,
	usles float64,
	uslep float64,
	IREG string,
	slope float64,
	slopeLength float64,
) {
	// get uslek
	uslek = "0" + Chorizon.Kwfact

	// get slope length
	slopeLength = StringUtil.Convert2Float(Chorizon.Slopelength_m)
	// slopeLengthFlag = true  -> change slopeLength
	// slopeLengthFlag = false -> continue
	// if slope length == 0, init it
	if slopeLength == 0.0 {
		slopeLength = 300 * 0.3048
	} else if slopeLengthFlag == true && slopeLength < 2.3 {
		slopeLength *= 0.3048
	}

	// get slope
	slope = StringUtil.Convert2Float(Chorizon.Slope_per)
	if knowSlope {
		slope = kSlope
	}

	// usles have a formula
	theta := math.Atan(slope / 100.0)
	M := 0.2
	if slope > 5 {
		M = 0.5
	} else if slope <= 5 && slope > 3 {
		M = 0.4
	} else if slope <= 3 && slope > 1 {
		M = 0.3
	} else if slope <= 1 {
		M = 0.2
	}
	// get usles by formula
	usles = math.Pow(slopeLength/72.6, M) * (65.41*math.Sin(theta*theta) + 4.56*math.Sin(theta) + 0.065)

	// get uslep
	uslep = 1.0 // default value

	// get ireg
	IREG = RainDistributionIreg.IREG

	// return values
	return uslek, usles, uslep, IREG, slope, slopeLength
}
