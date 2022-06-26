package pflag

import "strconv"

// -- count Value
type countValue int

func newCountValue(val int, p *int) *countValue {
	*p = val
	return (*countValue)(p)
}

func (i *countValue) Set(s string) error {
	// "+1" means that no specific value was passed, so increment
	if s == "+1" {
		*i = countValue(*i + 1)
		return nil
	}
	v, err := strconv.ParseInt(s, 0, 0)
	*i = countValue(v)
	return err
}

func (i *countValue) Type() string {
	return "count"
}

func (i *countValue) String() string { return strconv.Itoa(int(*i)) }

func countConv(sval string) (interface{}, error) {
	i, err := strconv.Atoi(sval)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// GetCount return the int value of a flag with the given name
func (f *FlagSet) GetCount(name string) (int, error) {
	val, err := f.getFlagType(name, "count", countConv)
	if err != nil {
		return 0, err
	}
	return val.(int), nil
}

// CountVar defines a count flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
// A count flag will add 1 to its value every time it is found on the command line
func (f *FlagSet) CountVar(p *int, name string, usage string) {
	f.CountVarP(p, name, "", usage)
}

// CountVarP is like CountVar only take a shorthand for the flag name.
func (f *FlagSet) CountVarP(p *int, name, shorthand string, usage string) {
	flag := f.VarPF(newCountValue(0, p), name, shorthand, usage)
	flag.NoOptDefVal = "+1"
}

// CountVar like CountVar only the flag is placed on the CommandLine instead of a given flag set
func CountVar(p *int, name string, usage string) {
	CommandLine.CountVar(p, name, usage)
}

// NewCountVar like NewCountVar only the flag is placed on the CommandLine instead of a given flag set
func NewCountVar(p *int, name string, usage string) *Flag {
	return NewCountVarP(p, name, "", usage)
}

// CountVarP is like CountVar only take a shorthand for the flag name.
func CountVarP(p *int, name, shorthand string, usage string) {
	CommandLine.CountVarP(p, name, shorthand, usage)
}

// NewCountVarP is like NewCountVar only take a shorthand for the flag name.
func NewCountVarP(p *int, name, shorthand string, usage string) *Flag {
	flag := NewVarPF(newCountValue(0, p), name, shorthand, usage)
	flag.NoOptDefVal = "+1"
	return flag
}

// Count defines a count flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
// A count flag will add 1 to its value every time it is found on the command line
func (f *FlagSet) Count(name string, usage string) *int {
	p := new(int)
	f.CountVarP(p, name, "", usage)
	return p
}

// CountP is like Count only takes a shorthand for the flag name.
func (f *FlagSet) CountP(name, shorthand string, usage string) *int {
	p := new(int)
	f.CountVarP(p, name, shorthand, usage)
	return p
}

// Count defines a count flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
// A count flag will add 1 to its value evey time it is found on the command line
func Count(name string, usage string) *int {
	return CommandLine.CountP(name, "", usage)
}

// NewCount defines a count flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
// A count flag will add 1 to its value evey time it is found on the command line
func NewCount(name string, usage string) (*int, *Flag) {
	return NewCountP(name, "", usage)
}

// CountP is like Count only takes a shorthand for the flag name.
func CountP(name, shorthand string, usage string) *int {
	return CommandLine.CountP(name, shorthand, usage)
}

// NewCountP is like NewCount only takes a shorthand for the flag name.
func NewCountP(name, shorthand string, usage string) (*int, *Flag) {
	p := new(int)
	return p, NewCountVarP(p, name, shorthand, usage)
}
