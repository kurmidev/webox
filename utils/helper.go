package utils

import "encoding/json"

func GetRoles(roleid int) string {
	switch roleid {
	case SADMIN:
		return "SuperAdmin"
	case USER:
		return "User"
	case STAFF:
		return "Staff"
	case ADMIN:
		return "Administrator"
	case SUBSCRIBER:
		return "Subscriber"
	default:
		return ""
	}
}

func GetCustomerType(i int) string {
	if i == 1 {
		return "Residential"
	} else {
		return "Commercial"
	}
}

func GetSubscriberBoquueStatus(i int) string {
	switch i {
	case -6:
		return "Manual Suspended"
	case -4:
		return "No Bouquet"
	case 1:
		return "Active"
	case -2:
		return "Expired"
	default:
		return "Suspended"
	}

}

func BouqueTypeslbl(i int) string {
	switch i {
	case 1:
		return "Base"
	case 2:
		return "Addons"
	case 3:
		return "Alacarte"
	default:
		return "Base"
	}
}

func BoxTypeLbl(i int) string {
	switch i {
	case 1:
		return "High Definition(HD)"
	case 2:
		return "Hybrid(HY)"
	default:
		return "Standard Definition(SD)"
	}
}

func IsApp(i int) string {
	if i > 0 {
		return "Yes"
	}

	return "No"
}

func FormatJson(s string) map[string]string {
	var r = make(map[string]string)
	err := json.Unmarshal([]byte(s), &r)
	if err != nil {
		return map[string]string{}
	}
	return r
}
