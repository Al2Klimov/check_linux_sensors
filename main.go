//go:generate go run vendor/github.com/Al2Klimov/go-gen-source-repos/main.go github.com/Al2Klimov/check_linux_sensors

package main

import (
	"fmt"
	_ "github.com/Al2Klimov/go-gen-source-repos"
	sensors "github.com/Al2Klimov/go-linux-sensors"
	. "github.com/Al2Klimov/go-monplug-utils"
	"math"
	"os"
	"strings"
)

var posInf = math.Inf(1)
var negInf = math.Inf(-1)

func main() {
	os.Exit(ExecuteCheck(onTerminal, checkLinuxSensors))
}

func onTerminal() (output string) {
	return fmt.Sprintf(
		"For the terms of use, the source code and the authors\n"+
			"see the projects this program is assembled from:\n\n  %s\n",
		strings.Join(GithubcomAl2klimovGo_gen_source_repos, "\n  "),
	)
}

func checkLinuxSensors() (output string, perfdata PerfdataCollection, errs map[string]error) {
	sensors.Init(nil)
	defer sensors.Cleanup()

	for _, chip := range sensors.GetDetectedChips(nil) {
		chipNameRaw, errMB := chip.MarshalBinary()
		if errMB != nil {
			errs = map[string]error{"sensors_snprintf_chip_name()": errMB}
			return
		}

		chipName := string(chipNameRaw)

		for _, feature := range chip.GetFeatures() {
			featureName := feature.GetName()

			switch feature.GetType() {
			case sensors.FeatureIn:
				vInput, hasInput, errsInput := getValue(chip, feature, sensors.SubfeatureInInput)
				if errsInput != nil {
					errs = errsInput
					return
				}

				vAverage, hasAverage, errsAverage := getValue(chip, feature, sensors.SubfeatureInAverage)
				if errsAverage != nil {
					errs = errsAverage
					return
				}

				vLowest, hasLowest, errsLowest := getValue(chip, feature, sensors.SubfeatureInLowest)
				if errsLowest != nil {
					errs = errsLowest
					return
				}

				vHighest, hasHighest, errsHighest := getValue(chip, feature, sensors.SubfeatureInHighest)
				if errsHighest != nil {
					errs = errsHighest
					return
				}

				if hasInput {
					vMin, errsMin := getOptionalValue(chip, feature, sensors.SubfeatureInMin)
					if errsMin != nil {
						errs = errsMin
						return
					}

					vMax, errsMax := getOptionalValue(chip, feature, sensors.SubfeatureInMax)
					if errsMax != nil {
						errs = errsMax
						return
					}

					vCrit, errsCrit := getOptionalThreshold(chip, feature, sensors.SubfeatureInLcrit, sensors.SubfeatureInCrit)
					if errsCrit != nil {
						errs = errsCrit
						return
					}

					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "input"),
						Value: vInput,
						Crit:  vCrit,
						Min:   vMin,
						Max:   vMax,
					})
				}

				if hasAverage {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "average"),
						Value: vAverage,
					})
				}

				if hasLowest {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "lowest"),
						Value: vLowest,
					})
				}

				if hasHighest {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "highest"),
						Value: vHighest,
					})
				}
			case sensors.FeatureVid:
				vVid, hasVid, errsVid := getValue(chip, feature, sensors.SubfeatureVid)
				if errsVid != nil {
					errs = errsVid
					return
				}

				if hasVid {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "vid"),
						Value: vVid,
					})
				}
			}
		}
	}

	return
}

func getValue(chip *sensors.ChipName, feature sensors.Feature, typ sensors.SubfeatureType) (float64, bool, map[string]error) {
	if subfeature, hasSubfeature := chip.GetSubfeature(feature, typ); hasSubfeature {
		if value, errGV := chip.GetValue(subfeature.GetNumber()); errGV == nil {
			return value, true, nil
		} else {
			return 0, true, map[string]error{"sensors_get_value()": errGV}
		}
	} else {
		return 0, false, nil
	}
}

func getOptionalValue(chip *sensors.ChipName, feature sensors.Feature, typ sensors.SubfeatureType) (OptionalNumber, map[string]error) {
	if subfeature, hasSubfeature := chip.GetSubfeature(feature, typ); hasSubfeature {
		if value, errGV := chip.GetValue(subfeature.GetNumber()); errGV == nil {
			return OptionalNumber{true, value}, nil
		} else {
			return OptionalNumber{}, map[string]error{"sensors_get_value()": errGV}
		}
	} else {
		return OptionalNumber{}, nil
	}
}

func getOptionalThreshold(chip *sensors.ChipName, feature sensors.Feature, typeStart, typeEnd sensors.SubfeatureType) (OptionalThreshold, map[string]error) {
	vStart, hasStart, errsStart := getValue(chip, feature, typeStart)
	if errsStart != nil {
		return OptionalThreshold{}, errsStart
	}

	vEnd, hasEnd, errsEnd := getValue(chip, feature, typeEnd)
	if errsEnd != nil {
		return OptionalThreshold{}, errsEnd
	}

	threshold := OptionalThreshold{}

	if hasStart || hasEnd {
		threshold.IsSet = true

		if hasStart {
			threshold.Start = vStart
		} else {
			threshold.Start = negInf
		}

		if hasEnd {
			threshold.End = vEnd
		} else {
			threshold.End = posInf
		}
	}

	return threshold, nil
}

func pdl(perfdataComponents ...string) string {
	return strings.Join(perfdataComponents, "::")
}
