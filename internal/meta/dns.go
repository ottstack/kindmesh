package meta

const SourceIP2SearchKeyPrefix = "SourceIP2Search_"
const AllDomainListKey = "AllDomainList"

type SourceIP2Search map[string][]string

type AllDomainList map[string]bool

func (a *AllDomainList) Clone() Cloneable {
	n := AllDomainList{}
	for k, v := range *a {
		n[k] = v
	}
	return &n
}

func (a *SourceIP2Search) Clone() Cloneable {
	n := SourceIP2Search{}
	for k, v := range *a {
		n[k] = v[:]
	}
	return &n
}
