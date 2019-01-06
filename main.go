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

				vAlarm, hasAlarm, errsAlarm := getValue(chip, feature, sensors.SubfeatureInAlarm)
				if errsAlarm != nil {
					errs = errsAlarm
					return
				}

				vMinAlarm, hasMinAlarm, errsMinAlarm := getValue(chip, feature, sensors.SubfeatureInMinAlarm)
				if errsMinAlarm != nil {
					errs = errsMinAlarm
					return
				}

				vMaxAlarm, hasMaxAlarm, errsMaxAlarm := getValue(chip, feature, sensors.SubfeatureInMaxAlarm)
				if errsMaxAlarm != nil {
					errs = errsMaxAlarm
					return
				}

				vLcritAlarm, hasLcritAlarm, errsLcritAlarm := getValue(chip, feature, sensors.SubfeatureInLcritAlarm)
				if errsLcritAlarm != nil {
					errs = errsLcritAlarm
					return
				}

				vCritAlarm, hasCritAlarm, errsCritAlarm := getValue(chip, feature, sensors.SubfeatureInCritAlarm)
				if errsCritAlarm != nil {
					errs = errsCritAlarm
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

				if hasAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "alarm"),
						Value: vAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasMinAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "min_alarm"),
						Value: vMinAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasMaxAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "max_alarm"),
						Value: vMaxAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasLcritAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "lcrit_alarm"),
						Value: vLcritAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasCritAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "crit_alarm"),
						Value: vCritAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
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
			case sensors.FeatureFan:
				vInput, hasInput, errsInput := getValue(chip, feature, sensors.SubfeatureFanInput)
				if errsInput != nil {
					errs = errsInput
					return
				}

				vAlarm, hasAlarm, errsAlarm := getValue(chip, feature, sensors.SubfeatureFanAlarm)
				if errsAlarm != nil {
					errs = errsAlarm
					return
				}

				vMinAlarm, hasMinAlarm, errsMinAlarm := getValue(chip, feature, sensors.SubfeatureFanMinAlarm)
				if errsMinAlarm != nil {
					errs = errsMinAlarm
					return
				}

				vMaxAlarm, hasMaxAlarm, errsMaxAlarm := getValue(chip, feature, sensors.SubfeatureFanMaxAlarm)
				if errsMaxAlarm != nil {
					errs = errsMaxAlarm
					return
				}

				if hasInput {
					vMin, errsMin := getOptionalValue(chip, feature, sensors.SubfeatureFanMin)
					if errsMin != nil {
						errs = errsMin
						return
					}

					vMax, errsMax := getOptionalValue(chip, feature, sensors.SubfeatureFanMax)
					if errsMax != nil {
						errs = errsMax
						return
					}

					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "input"),
						Value: vInput,
						Min:   vMin,
						Max:   vMax,
					})
				}

				if hasAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "alarm"),
						Value: vAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasMinAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "min_alarm"),
						Value: vMinAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasMaxAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "max_alarm"),
						Value: vMaxAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}
			case sensors.FeatureTemp:
				vInput, hasInput, errsInput := getValue(chip, feature, sensors.SubfeatureTempInput)
				if errsInput != nil {
					errs = errsInput
					return
				}

				vLowest, hasLowest, errsLowest := getValue(chip, feature, sensors.SubfeatureTempLowest)
				if errsLowest != nil {
					errs = errsLowest
					return
				}

				vHighest, hasHighest, errsHighest := getValue(chip, feature, sensors.SubfeatureTempHighest)
				if errsHighest != nil {
					errs = errsHighest
					return
				}

				vAlarm, hasAlarm, errsAlarm := getValue(chip, feature, sensors.SubfeatureTempAlarm)
				if errsAlarm != nil {
					errs = errsAlarm
					return
				}

				vMinAlarm, hasMinAlarm, errsMinAlarm := getValue(chip, feature, sensors.SubfeatureTempMinAlarm)
				if errsMinAlarm != nil {
					errs = errsMinAlarm
					return
				}

				vMaxAlarm, hasMaxAlarm, errsMaxAlarm := getValue(chip, feature, sensors.SubfeatureTempMaxAlarm)
				if errsMaxAlarm != nil {
					errs = errsMaxAlarm
					return
				}

				vLcritAlarm, hasLcritAlarm, errsLcritAlarm := getValue(chip, feature, sensors.SubfeatureTempLcritAlarm)
				if errsLcritAlarm != nil {
					errs = errsLcritAlarm
					return
				}

				vCritAlarm, hasCritAlarm, errsCritAlarm := getValue(chip, feature, sensors.SubfeatureTempCritAlarm)
				if errsCritAlarm != nil {
					errs = errsCritAlarm
					return
				}

				vEmergencyAlarm, hasEmergencyAlarm, errsEmergencyAlarm := getValue(chip, feature, sensors.SubfeatureTempEmergencyAlarm)
				if errsEmergencyAlarm != nil {
					errs = errsEmergencyAlarm
					return
				}

				if hasInput {
					vMin, errsMin := getOptionalValue(chip, feature, sensors.SubfeatureTempMin)
					if errsMin != nil {
						errs = errsMin
						return
					}

					vMax, errsMax := getOptionalValue(chip, feature, sensors.SubfeatureTempMax)
					if errsMax != nil {
						errs = errsMax
						return
					}

					vCrit, errsCrit := getOptionalThreshold(chip, feature, sensors.SubfeatureTempLcrit, sensors.SubfeatureTempCrit)
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

				if hasAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "alarm"),
						Value: vAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasMinAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "min_alarm"),
						Value: vMinAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasMaxAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "max_alarm"),
						Value: vMaxAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasLcritAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "lcrit_alarm"),
						Value: vLcritAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasCritAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "crit_alarm"),
						Value: vCritAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasEmergencyAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "emergency_alarm"),
						Value: vEmergencyAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}
			case sensors.FeatureCurr:
				vInput, hasInput, errsInput := getValue(chip, feature, sensors.SubfeatureCurrInput)
				if errsInput != nil {
					errs = errsInput
					return
				}

				vAverage, hasAverage, errsAverage := getValue(chip, feature, sensors.SubfeatureCurrAverage)
				if errsAverage != nil {
					errs = errsAverage
					return
				}

				vLowest, hasLowest, errsLowest := getValue(chip, feature, sensors.SubfeatureCurrLowest)
				if errsLowest != nil {
					errs = errsLowest
					return
				}

				vHighest, hasHighest, errsHighest := getValue(chip, feature, sensors.SubfeatureCurrHighest)
				if errsHighest != nil {
					errs = errsHighest
					return
				}

				vAlarm, hasAlarm, errsAlarm := getValue(chip, feature, sensors.SubfeatureCurrAlarm)
				if errsAlarm != nil {
					errs = errsAlarm
					return
				}

				vMinAlarm, hasMinAlarm, errsMinAlarm := getValue(chip, feature, sensors.SubfeatureCurrMinAlarm)
				if errsMinAlarm != nil {
					errs = errsMinAlarm
					return
				}

				vMaxAlarm, hasMaxAlarm, errsMaxAlarm := getValue(chip, feature, sensors.SubfeatureCurrMaxAlarm)
				if errsMaxAlarm != nil {
					errs = errsMaxAlarm
					return
				}

				vLcritAlarm, hasLcritAlarm, errsLcritAlarm := getValue(chip, feature, sensors.SubfeatureCurrLcritAlarm)
				if errsLcritAlarm != nil {
					errs = errsLcritAlarm
					return
				}

				vCritAlarm, hasCritAlarm, errsCritAlarm := getValue(chip, feature, sensors.SubfeatureCurrCritAlarm)
				if errsCritAlarm != nil {
					errs = errsCritAlarm
					return
				}

				if hasInput {
					vMin, errsMin := getOptionalValue(chip, feature, sensors.SubfeatureCurrMin)
					if errsMin != nil {
						errs = errsMin
						return
					}

					vMax, errsMax := getOptionalValue(chip, feature, sensors.SubfeatureCurrMax)
					if errsMax != nil {
						errs = errsMax
						return
					}

					vCrit, errsCrit := getOptionalThreshold(chip, feature, sensors.SubfeatureCurrLcrit, sensors.SubfeatureCurrCrit)
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

				if hasAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "alarm"),
						Value: vAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasMinAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "min_alarm"),
						Value: vMinAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasMaxAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "max_alarm"),
						Value: vMaxAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasLcritAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "lcrit_alarm"),
						Value: vLcritAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasCritAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "crit_alarm"),
						Value: vCritAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}
			case sensors.FeaturePower:
				{
					vAverage, hasAverage, errsAverage := getValue(chip, feature, sensors.SubfeaturePowerAverage)
					if errsAverage != nil {
						errs = errsAverage
						return
					}

					vLowest, hasLowest, errsLowest := getValue(chip, feature, sensors.SubfeaturePowerAverageLowest)
					if errsLowest != nil {
						errs = errsLowest
						return
					}

					vHighest, hasHighest, errsHighest := getValue(chip, feature, sensors.SubfeaturePowerAverageHighest)
					if errsHighest != nil {
						errs = errsHighest
						return
					}

					if hasAverage {
						perfdata = append(perfdata, Perfdata{
							Label: pdl(chipName, featureName, "average"),
							Value: vAverage,
						})
					}

					if hasLowest {
						perfdata = append(perfdata, Perfdata{
							Label: pdl(chipName, featureName, "average_lowest"),
							Value: vLowest,
						})
					}

					if hasHighest {
						perfdata = append(perfdata, Perfdata{
							Label: pdl(chipName, featureName, "average_highest"),
							Value: vHighest,
						})
					}
				}

				{
					vInput, hasInput, errsInput := getValue(chip, feature, sensors.SubfeaturePowerAverageInterval)
					if errsInput != nil {
						errs = errsInput
						return
					}

					if hasInput {
						perfdata = append(perfdata, Perfdata{
							Label: pdl(chipName, featureName, "average_interval"),
							UOM:   "s",
							Value: vInput,
						})
					}
				}

				{
					vInput, hasInput, errsInput := getValue(chip, feature, sensors.SubfeaturePowerInput)
					if errsInput != nil {
						errs = errsInput
						return
					}

					vLowest, hasLowest, errsLowest := getValue(chip, feature, sensors.SubfeaturePowerInputLowest)
					if errsLowest != nil {
						errs = errsLowest
						return
					}

					vHighest, hasHighest, errsHighest := getValue(chip, feature, sensors.SubfeaturePowerInputHighest)
					if errsHighest != nil {
						errs = errsHighest
						return
					}

					if hasInput {
						vMax, errsMax := getOptionalValue(chip, feature, sensors.SubfeaturePowerMax)
						if errsMax != nil {
							errs = errsMax
							return
						}

						vCrit, errsCrit := getOptionalThreshold(chip, feature, sensors.SubfeaturePowerCrit, sensors.SubfeaturePowerCrit)
						if errsCrit != nil {
							errs = errsCrit
							return
						}

						perfdata = append(perfdata, Perfdata{
							Label: pdl(chipName, featureName, "input"),
							Value: vInput,
							Crit:  vCrit,
							Max:   vMax,
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
				}

				vInput, hasInput, errsInput := getValue(chip, feature, sensors.SubfeaturePowerCap)
				if errsInput != nil {
					errs = errsInput
					return
				}

				if hasInput {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "cap"),
						Value: vInput,
					})
				}

				vAlarm, hasAlarm, errsAlarm := getValue(chip, feature, sensors.SubfeaturePowerAlarm)
				if errsAlarm != nil {
					errs = errsAlarm
					return
				}

				vCapAlarm, hasCapAlarm, errsCapAlarm := getValue(chip, feature, sensors.SubfeaturePowerCapAlarm)
				if errsCapAlarm != nil {
					errs = errsCapAlarm
					return
				}

				vMaxAlarm, hasMaxAlarm, errsMaxAlarm := getValue(chip, feature, sensors.SubfeaturePowerMaxAlarm)
				if errsMaxAlarm != nil {
					errs = errsMaxAlarm
					return
				}

				vCritAlarm, hasCritAlarm, errsCritAlarm := getValue(chip, feature, sensors.SubfeaturePowerCritAlarm)
				if errsCritAlarm != nil {
					errs = errsCritAlarm
					return
				}

				if hasAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "alarm"),
						Value: vAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasCapAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "cap_alarm"),
						Value: vCapAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasMaxAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "max_alarm"),
						Value: vMaxAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}

				if hasCritAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "crit_alarm"),
						Value: vCritAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
					})
				}
			case sensors.FeatureEnergy:
				vInput, hasInput, errsInput := getValue(chip, feature, sensors.SubfeatureEnergyInput)
				if errsInput != nil {
					errs = errsInput
					return
				}

				if hasInput {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "input"),
						Value: vInput,
					})
				}
			case sensors.FeatureHumidity:
				vInput, hasInput, errsInput := getValue(chip, feature, sensors.SubfeatureHumidityInput)
				if errsInput != nil {
					errs = errsInput
					return
				}

				if hasInput {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "input"),
						Value: vInput,
					})
				}
			case sensors.FeatureIntrusion:
				vAlarm, hasAlarm, errsAlarm := getValue(chip, feature, sensors.SubfeatureIntrusionAlarm)
				if errsAlarm != nil {
					errs = errsAlarm
					return
				}

				if hasAlarm {
					perfdata = append(perfdata, Perfdata{
						Label: pdl(chipName, featureName, "alarm"),
						Value: vAlarm,
						Crit:  OptionalThreshold{true, false, 0, 0},
						Min:   OptionalNumber{true, 0},
						Max:   OptionalNumber{true, 1},
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
	var vStart float64
	var hasStart bool
	var errsStart map[string]error

	if typeStart == typeEnd {
		hasStart = false
	} else {
		vStart, hasStart, errsStart = getValue(chip, feature, typeStart)
		if errsStart != nil {
			return OptionalThreshold{}, errsStart
		}
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
