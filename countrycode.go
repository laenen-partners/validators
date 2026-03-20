package validators

import (
	"fmt"
	"strings"
)

var countriesAlpha2 = map[string]bool{
	"AD": true, "AE": true, "AF": true, "AG": true, "AI": true, "AL": true, "AM": true, "AO": true,
	"AQ": true, "AR": true, "AS": true, "AT": true, "AU": true, "AW": true, "AX": true, "AZ": true,
	"BA": true, "BB": true, "BD": true, "BE": true, "BF": true, "BG": true, "BH": true, "BI": true,
	"BJ": true, "BL": true, "BM": true, "BN": true, "BO": true, "BQ": true, "BR": true, "BS": true,
	"BT": true, "BV": true, "BW": true, "BY": true, "BZ": true, "CA": true, "CC": true, "CD": true,
	"CF": true, "CG": true, "CH": true, "CI": true, "CK": true, "CL": true, "CM": true, "CN": true,
	"CO": true, "CR": true, "CU": true, "CV": true, "CW": true, "CX": true, "CY": true, "CZ": true,
	"DE": true, "DJ": true, "DK": true, "DM": true, "DO": true, "DZ": true, "EC": true, "EE": true,
	"EG": true, "EH": true, "ER": true, "ES": true, "ET": true, "FI": true, "FJ": true, "FK": true,
	"FM": true, "FO": true, "FR": true, "GA": true, "GB": true, "GD": true, "GE": true, "GF": true,
	"GG": true, "GH": true, "GI": true, "GL": true, "GM": true, "GN": true, "GP": true, "GQ": true,
	"GR": true, "GS": true, "GT": true, "GU": true, "GW": true, "GY": true, "HK": true, "HM": true,
	"HN": true, "HR": true, "HT": true, "HU": true, "ID": true, "IE": true, "IL": true, "IM": true,
	"IN": true, "IO": true, "IQ": true, "IR": true, "IS": true, "IT": true, "JE": true, "JM": true,
	"JO": true, "JP": true, "KE": true, "KG": true, "KH": true, "KI": true, "KM": true, "KN": true,
	"KP": true, "KR": true, "KW": true, "KY": true, "KZ": true, "LA": true, "LB": true, "LC": true,
	"LI": true, "LK": true, "LR": true, "LS": true, "LT": true, "LU": true, "LV": true, "LY": true,
	"MA": true, "MC": true, "MD": true, "ME": true, "MF": true, "MG": true, "MH": true, "MK": true,
	"ML": true, "MM": true, "MN": true, "MO": true, "MP": true, "MQ": true, "MR": true, "MS": true,
	"MT": true, "MU": true, "MV": true, "MW": true, "MX": true, "MY": true, "MZ": true, "NA": true,
	"NC": true, "NE": true, "NF": true, "NG": true, "NI": true, "NL": true, "NO": true, "NP": true,
	"NR": true, "NU": true, "NZ": true, "OM": true, "PA": true, "PE": true, "PF": true, "PG": true,
	"PH": true, "PK": true, "PL": true, "PM": true, "PN": true, "PR": true, "PS": true, "PT": true,
	"PW": true, "PY": true, "QA": true, "RE": true, "RO": true, "RS": true, "RU": true, "RW": true,
	"SA": true, "SB": true, "SC": true, "SD": true, "SE": true, "SG": true, "SH": true, "SI": true,
	"SJ": true, "SK": true, "SL": true, "SM": true, "SN": true, "SO": true, "SR": true, "SS": true,
	"ST": true, "SV": true, "SX": true, "SY": true, "SZ": true, "TC": true, "TD": true, "TF": true,
	"TG": true, "TH": true, "TJ": true, "TK": true, "TL": true, "TM": true, "TN": true, "TO": true,
	"TR": true, "TT": true, "TV": true, "TW": true, "TZ": true, "UA": true, "UG": true, "UM": true,
	"US": true, "UY": true, "UZ": true, "VA": true, "VC": true, "VE": true, "VG": true, "VI": true,
	"VN": true, "VU": true, "WF": true, "WS": true, "YE": true, "YT": true, "ZA": true, "ZM": true,
	"ZW": true,
}

var countriesAlpha3 = map[string]bool{
	"ABW": true, "AFG": true, "AGO": true, "AIA": true, "ALA": true, "ALB": true, "AND": true, "ARE": true,
	"ARG": true, "ARM": true, "ASM": true, "ATA": true, "ATF": true, "ATG": true, "AUS": true, "AUT": true,
	"AZE": true, "BDI": true, "BEL": true, "BEN": true, "BES": true, "BFA": true, "BGD": true, "BGR": true,
	"BHR": true, "BHS": true, "BIH": true, "BLM": true, "BLR": true, "BLZ": true, "BMU": true, "BOL": true,
	"BRA": true, "BRB": true, "BRN": true, "BTN": true, "BVT": true, "BWA": true, "CAF": true, "CAN": true,
	"CCK": true, "CHE": true, "CHL": true, "CHN": true, "CIV": true, "CMR": true, "COD": true, "COG": true,
	"COK": true, "COL": true, "COM": true, "CPV": true, "CRI": true, "CUB": true, "CUW": true, "CXR": true,
	"CYM": true, "CYP": true, "CZE": true, "DEU": true, "DJI": true, "DMA": true, "DNK": true, "DOM": true,
	"DZA": true, "ECU": true, "EGY": true, "ERI": true, "ESH": true, "ESP": true, "EST": true, "ETH": true,
	"FIN": true, "FJI": true, "FLK": true, "FRA": true, "FRO": true, "FSM": true, "GAB": true, "GBR": true,
	"GEO": true, "GGY": true, "GHA": true, "GIB": true, "GIN": true, "GLP": true, "GMB": true, "GNB": true,
	"GNQ": true, "GRC": true, "GRD": true, "GRL": true, "GTM": true, "GUF": true, "GUM": true, "GUY": true,
	"HKG": true, "HMD": true, "HND": true, "HRV": true, "HTI": true, "HUN": true, "IDN": true, "IMN": true,
	"IND": true, "IOT": true, "IRL": true, "IRN": true, "IRQ": true, "ISL": true, "ISR": true, "ITA": true,
	"JAM": true, "JEY": true, "JOR": true, "JPN": true, "KAZ": true, "KEN": true, "KGZ": true, "KHM": true,
	"KIR": true, "KNA": true, "KOR": true, "KWT": true, "LAO": true, "LBN": true, "LBR": true, "LBY": true,
	"LCA": true, "LIE": true, "LKA": true, "LSO": true, "LTU": true, "LUX": true, "LVA": true, "MAC": true,
	"MAF": true, "MAR": true, "MCO": true, "MDA": true, "MDG": true, "MDV": true, "MEX": true, "MHL": true,
	"MKD": true, "MLI": true, "MLT": true, "MMR": true, "MNE": true, "MNG": true, "MNP": true, "MOZ": true,
	"MRT": true, "MSR": true, "MTQ": true, "MUS": true, "MWI": true, "MYS": true, "MYT": true, "NAM": true,
	"NCL": true, "NER": true, "NFK": true, "NGA": true, "NIC": true, "NIU": true, "NLD": true, "NOR": true,
	"NPL": true, "NRU": true, "NZL": true, "OMN": true, "PAK": true, "PAN": true, "PCN": true, "PER": true,
	"PHL": true, "PLW": true, "PNG": true, "POL": true, "PRI": true, "PRK": true, "PRT": true, "PRY": true,
	"PSE": true, "PYF": true, "QAT": true, "REU": true, "ROU": true, "RUS": true, "RWA": true, "SAU": true,
	"SDN": true, "SEN": true, "SGP": true, "SGS": true, "SHN": true, "SJM": true, "SLB": true, "SLE": true,
	"SLV": true, "SMR": true, "SOM": true, "SPM": true, "SRB": true, "SSD": true, "STP": true, "SUR": true,
	"SVK": true, "SVN": true, "SWE": true, "SWZ": true, "SXM": true, "SYC": true, "SYR": true, "TCA": true,
	"TCD": true, "TGO": true, "THA": true, "TJK": true, "TKL": true, "TKM": true, "TLS": true, "TON": true,
	"TTO": true, "TUN": true, "TUR": true, "TUV": true, "TWN": true, "TZA": true, "UGA": true, "UKR": true,
	"UMI": true, "URY": true, "USA": true, "UZB": true, "VAT": true, "VCT": true, "VEN": true, "VGB": true,
	"VIR": true, "VNM": true, "VUT": true, "WLF": true, "WSM": true, "YEM": true, "ZAF": true, "ZMB": true,
	"ZWE": true,
}

// CountryCode validates an ISO 3166-1 country code (alpha-2 or alpha-3).
func CountryCode(value string) Result {
	if value == "" {
		return valid(nil)
	}

	code := strings.ToUpper(strings.TrimSpace(value))

	switch len(code) {
	case 2:
		if !countriesAlpha2[code] {
			return invalid(ErrCountryCodeUnknown,
				fmt.Sprintf("Unknown country code: %s", code),
				"country_code", map[string]any{
					"value":  value,
					"format": "alpha-2",
				})
		}
		return valid(map[string]any{"format": "alpha-2", "code": code})
	case 3:
		if !countriesAlpha3[code] {
			return invalid(ErrCountryCodeUnknown,
				fmt.Sprintf("Unknown country code: %s", code),
				"country_code", map[string]any{
					"value":  value,
					"format": "alpha-3",
				})
		}
		return valid(map[string]any{"format": "alpha-3", "code": code})
	default:
		return invalid(ErrCountryCodeFormatInvalid,
			"Country code must be 2 (alpha-2) or 3 (alpha-3) letters",
			"country_code", map[string]any{
				"value":  value,
				"length": len(code),
			})
	}
}
