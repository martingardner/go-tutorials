package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	urlRaw := "postgres://user:pass@host.com:5432/path?k=va#f"

	//Parse returns two variables to account for, the actual value and if an error
	u, err := url.Parse(urlRaw)
	if err != nil {
		panic(err)
	}

	fmt.Println("u.Scheme - ", u.Scheme)
	fmt.Println("u.User - ", u.User)
	fmt.Println("u.User.Username() - ", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("p u.User.Password() - ", p)

	fmt.Println("u.Host - ", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host - ", host)
	fmt.Println("port - ", port)

	fmt.Println("u.Path - ", u.Path)
	fmt.Println("u.Fragment - ", u.Fragment)
	fmt.Println("u.RawQuery - ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("m rawQuery - ", m)
	//0 is always fixed
	fmt.Println("m[k][0] - ", m["k"][0])
}
