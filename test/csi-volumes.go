/*
Copyright 2019 The Kubernetes Authors.
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

package test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/testfiles"
	"k8s.io/kubernetes/test/e2e/storage/testsuites"
	"k8s.io/kubernetes/test/e2e/storage/utils"
	"path"
)

var CSITestSuites = []func() testsuites.TestSuite{
	testsuites.InitVolumesTestSuite,
	testsuites.InitVolumeIOTestSuite,
	testsuites.InitVolumeModeTestSuite,
	testsuites.InitSubPathTestSuite,
	testsuites.InitProvisioningTestSuite,
}

// This executes testSuites for csi volumes.
var _ = utils.SIGDescribe("CSI Volumes", func() {
	testfiles.AddFileSource(testfiles.RootFileSource{Root: path.Join(framework.TestContext.RepoRoot, "../../deploy/kubernetes/")})

	defaultFramework := framework.NewDefaultFramework("default-framework")
	driverName := "csi-nfsplugin"
	manifests := []string{driverName,
		"csi-attacher-nfsplugin.yaml",
		"csi-attacher-rbac.yaml",
		"csi-nodeplugin-nfsplugin.yaml",
		"csi-nodeplugin-rbac.yaml"}
	cleanup, err := defaultFramework.CreateFromManifests(nil, manifests...)

	if err != nil {
		framework.Failf("deploying %s driver: %v", driverName, err)
	}

	curDriver := NFSdriver()
	Context(testsuites.GetDriverNameWithFeatureTags(curDriver), func() {
		testsuites.DefineTestSuite(curDriver, CSITestSuites)
		By(fmt.Sprintf("uninstalling %s driver", driverName))
		defer cleanup()
	})

})
