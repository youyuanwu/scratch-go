package wintls

import (
	"crypto/x509"
	"fmt"

	"github.com/google/certtostore"
	"golang.org/x/sys/windows"
)

func GetFromStore() (*windows.CertContext, *certtostore.Key, *x509.Certificate) {
	fmt.Println("open cert store")

	// Open the local cert store. Provider generally shouldn't matter, so use Software which is ubiquitous. See comments in getHostKey.
	store, err := certtostore.OpenWinCertStoreCurrentUser(certtostore.ProviderMSSoftware, "My", []string{"Youyuan Wu Test"}, nil, false)

	if err != nil {
		panic(err)
	}

	fmt.Println("get cert from cert store")
	// Obtain the first cert matching all of container/issuers/intermediates in the store.
	// This function is indifferent to the provider the store was opened with, as the store lists certs
	// from all providers.
	crt, context, err := store.CertWithContext()
	if err != nil {
		panic(err)
	}

	if crt == nil {
		panic("no cert")
	}

	fmt.Println("get key from cert")
	// Obtain the private key from the cert. This *should* work regardless of provider because
	// the key is directly linked to the certificate.
	key, err := store.CertKey(context)
	if err != nil {
		panic(fmt.Sprintf("private key not found in %s, %s", store.ProvName, err))
	}

	if key == nil {
		panic("no key")
	}

	fmt.Printf("find cert '%s' with private key in container '%s', algo '%s'\n", crt.Subject, key.Container, key.AlgorithmGroup)

	return context, key, crt
}
