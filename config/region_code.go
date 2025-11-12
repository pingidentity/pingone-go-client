// Copyright © 2025 Ping Identity Corporation
package config

// RegionCode represents a PingOne region code.
type RegionCode string

const (
	// RegionCodeNA represents the North America region
	RegionCodeNA RegionCode = "NA"
	// RegionCodeEU represents the Europe region
	RegionCodeEU RegionCode = "EU"
	// RegionCodeAP represents the Asia-Pacific region
	RegionCodeAP RegionCode = "AP"
	// RegionCodeAPAC is an alias for Asia-Pacific region
	RegionCodeAPAC RegionCode = "APAC"
	// RegionCodeCA represents the Canada region
	RegionCodeCA RegionCode = "CA"
	// RegionCodeAU represents the Australia region
	RegionCodeAU RegionCode = "AU"
	// RegionCodeSG represents the Singapore region
	RegionCodeSG RegionCode = "SG"
)
