package g

import "fmt"

type Address struct {
	Server     string
	Namespace  string
	Repository string
}

func (a Address) HTTPSUrl() string {
	return fmt.Sprintf("https://%v/%v/%v.git", a.Server, a.Namespace, a.Repository)
}

func (a Address) SSHUrl() string {
	return fmt.Sprintf("git@%v:%v/%v.git", a.Server, a.Namespace, a.Repository)
}
