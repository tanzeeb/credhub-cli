package credhub_test

import (
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry-incubator/credhub-cli/credhub"
)

var _ = Describe("Client()", func() {
	It("should return a simple http.Client", func() {
		ch := credhubFromConfig(Config{ApiUrl: "http://example.com"})
		client := ch.Client()

		Expect(client).ToNot(BeNil())
	})

	Context("With CaCerts", func() {
		It("should return a http.Client with tls.Config with RootCAs", func() {
			fixturePath := "./fixtures/"
			caCertFiles := []string{
				"auth-tls-ca.pem",
				"server-tls-ca.pem",
				"extra-ca.pem",
			}
			var caCerts []string
			expectedRootCAs := x509.NewCertPool()
			for _, caCertFile := range caCertFiles {
				caCertBytes, err := ioutil.ReadFile(fixturePath + caCertFile)
				if err != nil {
					Fail("Couldn't read certificate " + caCertFile + ": " + err.Error())
				}

				caCerts = append(caCerts, string(caCertBytes))
				expectedRootCAs.AppendCertsFromPEM(caCertBytes)
			}

			ch := credhubFromConfig(Config{ApiUrl: "https://example.com", CaCerts: caCerts})

			client := ch.Client()

			transport := client.Transport.(*http.Transport)
			tlsConfig := transport.TLSClientConfig

			Expect(client.Timeout).To(Equal(45 * time.Second))

			Expect(tlsConfig.InsecureSkipVerify).To(BeFalse())
			Expect(tlsConfig.PreferServerCipherSuites).To(BeTrue())
			Expect(tlsConfig.RootCAs.Subjects()).To(ConsistOf(expectedRootCAs.Subjects()))
		})
	})

	Context("With InsecureSkipVerify", func() {
		It("should return a http.Client with tls.Config without RootCAs", func() {
			ch := credhubFromConfig(Config{ApiUrl: "https://example.com", InsecureSkipVerify: true})
			client := ch.Client()

			transport := client.Transport.(*http.Transport)
			tlsConfig := transport.TLSClientConfig

			Expect(client.Timeout).To(Equal(45 * time.Second))

			Expect(tlsConfig.InsecureSkipVerify).To(BeTrue())
			Expect(tlsConfig.PreferServerCipherSuites).To(BeTrue())
		})
	})
})