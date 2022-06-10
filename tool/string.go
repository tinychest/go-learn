package tool

import "strings"

func JoinRepeat(elem string, times int, seps ...string) string {
	switch times {
	case 0:
		return ""
	case 1:
		return elem
	}

	const defaultSep = ","
	var (
		sep  = append(seps, defaultSep)[0]
		elen = len(elem)
		slen = len(sep)
	)

	var b strings.Builder
	b.Grow(times*elen + (times-1)*slen)

	b.WriteString(elem)
	for i := 1; i < times; i++ {
		b.WriteString(defaultSep)
		b.WriteString(elem)
	}
	return b.String()
}
