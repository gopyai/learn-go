// main
package main

import (
	"crypto/x509"
	"io/ioutil"
)

func main() {
	privCA, cerCA, keyCA, cerCA_b := genCA()
	ioutil.WriteFile("ca.key", keyCA, 0777)
	ioutil.WriteFile("ca.cer", cerCA_b, 0777)

	privSvr, cerSvr, keySvr, cerSvr_b := genSvr(privCA, cerCA)
	ioutil.WriteFile("svr.key", keySvr, 0777)
	ioutil.WriteFile("svr.cer", cerSvr_b, 0777)

	_, _ = privSvr, cerSvr

	//

	cerCA2, e := x509.ParseCertificate(cerCA_b)
	isErr(e)
	cerSvr2, e := x509.ParseCertificate(cerSvr_b)
	isErr(e)

	e = cerSvr2.CheckSignatureFrom(cerCA2)
	isErr(e)
}

func isErr(e error) {
	if e != nil {
		panic(e)
	}
}
