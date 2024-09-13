/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package support

import (
	"crypto/tls"
	"net/http"

	. "github.com/onsi/gomega"

	"k8s.io/client-go/transport"
)

func GetRayClusterClient(t Test, dashboardURL, bearerToken string) RayClusterClient {
	t.T().Helper()

	// Skip TLS check to work on clusters with insecure certificates too
	// Functionality intended just for testing purpose, DO NOT USE IN PRODUCTION
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyFromEnvironment,
	}
	client, err := NewRayClusterClient(RayClusterClientConfig{
		Address: dashboardURL,
		Client:  &http.Client{Transport: transport.NewBearerAuthRoundTripper(bearerToken, tr)},
	})
	t.Expect(err).NotTo(HaveOccurred())

	return client
}
