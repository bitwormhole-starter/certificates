package certificates

import "strings"

////////////////////////////////////////////////////////////////////////////////

// Raw ...
type Raw struct {
	Data []byte
}

// Clone ...
func (inst *Raw) Clone() *Raw {
	data1 := inst.Data
	if data1 == nil {
		data1 = []byte{}
	}
	size := len(data1)
	data2 := make([]byte, size)
	copy(data2, data1)
	return &Raw{Data: data2}
}

////////////////////////////////////////////////////////////////////////////////

// Name ...
type Name struct {
	Country, Organization, OrganizationalUnit string
	Locality, Province                        string
	StreetAddress, PostalCode                 string
	SerialNumber, CommonName                  string
}

// Equals ...
func (inst *Name) Equals(other *Name) bool {
	if other == nil || inst == nil {
		return false
	}
	s1 := inst.String()
	s2 := other.String()
	return s1 == s2
}

func (inst *Name) String() string {

	list := make([]string, 0)

	list = append(list, inst.Country)
	list = append(list, inst.Organization)
	list = append(list, inst.OrganizationalUnit)

	list = append(list, inst.Locality)
	list = append(list, inst.Province)

	list = append(list, inst.StreetAddress)
	list = append(list, inst.PostalCode)

	list = append(list, inst.SerialNumber)
	list = append(list, inst.CommonName)

	builder := strings.Builder{}
	for _, str := range list {
		str = strings.TrimSpace(str)
		if str == "" {
			str = "-"
		}
		builder.WriteString(str)
		builder.WriteString("\n")
	}
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////
