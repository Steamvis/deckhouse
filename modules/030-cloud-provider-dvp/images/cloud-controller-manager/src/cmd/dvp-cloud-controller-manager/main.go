/*
Copyright 2025 Flant JSC

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

package main

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/util/wait"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/cloud-provider/app"
	cloudcontrollerconfig "k8s.io/cloud-provider/app/config"
	"k8s.io/cloud-provider/names"
	"k8s.io/cloud-provider/options"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"

	_ "dvp-cloud-controller-manager/pkg/cloudprovider/dvp"

	_ "k8s.io/component-base/metrics/prometheus/clientgo" // load all the prometheus client-go plugins
	_ "k8s.io/component-base/metrics/prometheus/version"  // for version metric registration
	"k8s.io/klog/v2"
)

const (
	dvpCloudProviderName = "dvp"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	opts, err := options.NewCloudControllerManagerOptions()
	if err != nil {
		klog.Fatalf("unable to initialize command options: %v", err)
	}

	controllerInitializers := app.DefaultInitFuncConstructors
	fss := cliflag.NamedFlagSets{}

	controllerAliases := names.CCMControllerAliases()

	command := app.NewCloudControllerManagerCommand(
		opts,
		cloudInitializer,
		controllerInitializers,
		controllerAliases,
		fss,
		wait.NeverStop,
	)

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func cloudInitializer(config *cloudcontrollerconfig.CompletedConfig) cloudprovider.Interface {
	cloudConfig := config.ComponentConfig.KubeCloudShared.CloudProvider
	providerName := cloudConfig.Name

	if providerName == "" {
		providerName = dvpCloudProviderName
	}

	// initialize cloud provider with the cloud provider name and config file provided
	cloud, err := cloudprovider.InitCloudProvider(providerName, cloudConfig.CloudConfigFile)
	if err != nil {
		klog.Fatalf("Cloud provider could not be initialized: %v", err)
	}
	if cloud == nil {
		klog.Fatalf("Cloud provider is nil")

		return nil
	}

	if !cloud.HasClusterID() {
		if config.ComponentConfig.KubeCloudShared.AllowUntaggedCloud {
			klog.Warning("detected a cluster without a ClusterID. A ClusterID will be required in the future. Please tag your cluster to avoid any future issues")
		} else {
			klog.Fatalf("no ClusterID found. A ClusterID is required for the cloud provider to function properly. This check can be bypassed by setting the allow-untagged-cloud option")
		}
	}

	return cloud
}
