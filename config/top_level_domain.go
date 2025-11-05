// Copyright © 2025 Ping Identity Corporation
package config

type TopLevelDomain string

const (
	TopLevelDomainAPAC TopLevelDomain = "asia"
	TopLevelDomainAU   TopLevelDomain = "com.au"
	TopLevelDomainCA   TopLevelDomain = "ca"
	TopLevelDomainEU   TopLevelDomain = "eu"
	TopLevelDomainNA   TopLevelDomain = "com"
	TopLevelDomainSG   TopLevelDomain = "sg"
)

func IsValidTopLevelDomain(tld string) bool {
	switch TopLevelDomain(tld) {
	case TopLevelDomainAPAC, TopLevelDomainAU, TopLevelDomainCA, TopLevelDomainEU, TopLevelDomainNA, TopLevelDomainSG:
		return true
	}
	return false
}
