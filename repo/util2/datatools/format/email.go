package format

import (
	"strings"
)

func MakeSet(list []string) (retval []string) {

    set := make(map[string]bool, len(list))

	for _, item := range list {
		if _, ok := set[item]; !ok {
            set[item] = true
            retval = append(retval, item)
		}
	}
	return
}

func GetDomainNames(s string) (retval string) {

	if len(s) < 3 {
		return s
	}

	list := strings.Fields(s)
	sb := strings.Builder{}

	for _, item := range list {
        if index := strings.Index(item, "@"); index > 1 {
            item = item[index+1:]
        }
		// for strings.Contains(item, "@") {
		// 	item = item[strings.Index(item, "@")+1:]
		// }
		sb.WriteString(" ")
		sb.WriteString(item)
	}

	return sb.String()[1:] // there is always a leading space from the loop
}

// FromGmailFilterNames removes the " OR " separators between elements
// in the gmail filters sender list
func FromGmailFilterNames(s string) string {
    return strings.Join(strings.Split(s," OR "), " ")
}

// ToGmailFilterNames removes the " OR " separators between elements
// in the gmail filters sender list
func ToGmailFilterNames(s string) string {
    return strings.Join(strings.Split(s," "), " OR ")
}

func GetTopLevelDomains(s string) (retval string) {

	if len(s) < 3 {
		return s
	}

    // removes the xxxxxx@ prefix
	list := strings.Fields(GetDomainNames(s))

	tmp := make([]string, 0, len(list))



	sb := strings.Builder{}

	for _, s := range list {
		parts := strings.Split(s, ".")

		if len(parts) > 1 {
			parts = parts[len(parts)-2:]
		}

		tmp = append(tmp, strings.Join(parts, "."))
	}

	sb.WriteString(" ")
	sb.WriteString(strings.Join(MakeSet(tmp), "."))

	return strings.TrimSpace(sb.String())
    // return strings.Join(MakeSet(list)," ")

}
