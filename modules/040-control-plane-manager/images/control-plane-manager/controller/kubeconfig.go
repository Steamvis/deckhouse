/*
Copyright 2023 Flant JSC

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
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/deckhouse/deckhouse/pkg/log"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	configv1 "k8s.io/client-go/tools/clientcmd/api/v1"
	"sigs.k8s.io/yaml"
)

var (
	ErrClustersFieldEmpty        = errors.New("clusters field is empty")
	ErrUsersFieldEmpty           = errors.New("users field is empty")
	ErrCertDataFieldEmpty        = errors.New("client-certificate-data field is empty")
	ErrCertDecodeFailed          = errors.New("cannot PEM decode client-certificate-data")
	ErrCertParseFailed           = errors.New("cannot parse certificate from client-certificate-data")
	ErrServerAddressChanged      = errors.New("kubeconfig address field changed")
	ErrCertExpiringSoon          = errors.New("client certificate is expiring in less than 30 days")
	ErrCantReadOrUnmarshalConfig = errors.New("cannot read or unmarshal kubeconfig")
)

var shouldRecreateKubeConfigErrors = []error{
	ErrClustersFieldEmpty,
	ErrUsersFieldEmpty,
	ErrCertDataFieldEmpty,
	ErrCertDecodeFailed,
	ErrCertParseFailed,
	ErrServerAddressChanged,
	ErrCertExpiringSoon,
	ErrCantReadOrUnmarshalConfig,
}

func shouldRecreateKubeConfig(err error) bool {
	for _, e := range shouldRecreateKubeConfigErrors {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}

func renewKubeconfigs() error {
	log.Info("phase: renew kubeconfigs")

	kubeconfigs := []string{"admin", "controller-manager", "scheduler"}

	c, err := semver.NewConstraint(">= 1.29")
	if err != nil {
		return fmt.Errorf("constraint not being parsable: %s", err.Error())
	}
	v, err := semver.NewVersion(config.KubernetesVersion)
	if err != nil {
		return fmt.Errorf("version not being parsable: %s", err.Error())
	}
	// if KubernetesVersion >= 1.29
	if c.Check(v) {
		kubeconfigs = []string{"super-admin", "admin", "controller-manager", "scheduler"}
	}

	for _, v := range kubeconfigs {
		if err := renewKubeconfig(v); err != nil {
			return err
		}
	}
	return nil
}

func renewKubeconfig(componentName string) error {
	path := filepath.Join(kubernetesConfigPath, componentName+".conf")
	log.Infof("generate or renew %s kubeconfig", path)
	if _, err := os.Stat(path); err == nil && config.ConfigurationChecksum != config.LastAppliedConfigurationChecksum {
		tmpPath := filepath.Join(config.TmpPath, path)
		log.Infof("configuration has changed since last kubeconfig generation (last applied checksum %s, configuration checksum %s), verifying kubeconfig", config.LastAppliedConfigurationChecksum, config.ConfigurationChecksum)
		if err := prepareKubeconfig(componentName, true); err != nil {
			return err
		}

		err := validateKubeconfig(path, tmpPath)
		if err != nil {
			log.Error(err.Error())
		}
		if shouldRecreateKubeConfig(err) {
			removeFile(path)
		}
	}

	if _, err := os.Stat(path); err == nil {
		return nil
	}

	// regenerate kubeconfig
	log.Infof("generate new kubeconfig %s", path)
	return prepareKubeconfig(componentName, false)
}

func validateKubeconfig(path string, tmpPath string) error {
	currentKubeconfig, err := loadKubeconfig(path)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	tmpKubeconfig, err := loadKubeconfig(tmpPath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if len(currentKubeconfig.Clusters) == 0 {
		return fmt.Errorf("%w: %s", ErrClustersFieldEmpty, path)
	}

	if len(tmpKubeconfig.Clusters) == 0 {
		return fmt.Errorf("%w: %s", ErrClustersFieldEmpty, tmpPath)
	}

	if len(currentKubeconfig.AuthInfos) == 0 {
		return fmt.Errorf("%w: %s", ErrUsersFieldEmpty, path)
	}

	certData := currentKubeconfig.AuthInfos[0].AuthInfo.ClientCertificateData
	if len(certData) == 0 {
		return fmt.Errorf("%w: %s", ErrCertDataFieldEmpty, path)
	}

	block, _ := pem.Decode(certData)
	if len(block.Bytes) == 0 {
		return fmt.Errorf("%w: %s", ErrCertDecodeFailed, path)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("%w: %s: %v", ErrCertParseFailed, path, err)
	}

	if currentKubeconfig.Clusters[0].Cluster.Server != tmpKubeconfig.Clusters[0].Cluster.Server {
		return fmt.Errorf("%w: %s", ErrServerAddressChanged, path)
	}

	if certificateExpiresSoon(cert, 30*24*time.Hour) {
		return fmt.Errorf("%w: %s", ErrCertExpiringSoon, path)
	}

	return nil
}

func prepareKubeconfig(componentName string, isTemp bool) error {
	// kubeadm init phase kubeconfig apiserver --config /etc/kubernetes/deckhouse/kubeadm/config.yaml
	args := []string{"init", "phase", "kubeconfig", componentName, "--config", deckhousePath + "/kubeadm/config.yaml"}
	if isTemp {
		args = append(args, "--rootfs", config.TmpPath)
	}
	c := exec.Command(kubeadmPath, args...)
	out, err := c.CombinedOutput()
	for _, s := range strings.Split(string(out), "\n") {
		log.Infof("%s", s)
	}
	return err
}

func loadKubeconfig(path string) (*configv1.Config, error) {
	res := &configv1.Config{}
	r, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCantReadOrUnmarshalConfig, err)
	}

	err = yaml.Unmarshal(r, res)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCantReadOrUnmarshalConfig, err)
	}

	return res, nil
}

func updateRootKubeconfig() error {
	var path string
	if homeDir, hasHomeDir := os.LookupEnv("HOME"); hasHomeDir && homeDir != "/" {
		path = filepath.Join(homeDir, ".kube", "config")
	} else {
		path = "/root/.kube/config"
	}

	originalPath := filepath.Join(kubernetesConfigPath, "admin.conf")
	log.Infof("update root user kubeconfig (%s)", path)
	if _, err := os.Stat(path); err == nil {
		p, err := filepath.EvalSymlinks(path)
		if p == originalPath && err == nil {
			return nil
		}
		if err := os.Remove(path); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		return err
	}

	return os.Symlink(originalPath, path)
}

func checkKubeletConfig() error {
	kubeletPath := filepath.Join(kubernetesConfigPath, "kubelet.conf")
	log.Infof("phase: check kubelet config %s", kubeletPath)

	res, err := loadKubeconfig(kubeletPath)
	if err != nil {
		return err
	}

	if strings.HasPrefix(res.Clusters[0].Cluster.Server, "https://127.0.0.1:6445") {
		return nil
	}

	return fmt.Errorf("cannot find server: https://127.0.0.1:6445 in kubelet config %s, kubelet should be configured "+
		"to access apiserver via kube-api-proxy (through https://127.0.0.1:6445), probably node is not managed by node-manager", kubeletPath)
}

func installKubeadmConfig() error {
	log.Info("phase: install kubeadm configuration")
	if err := os.MkdirAll(filepath.Join(deckhousePath, "kubeadm", "patches"), 0700); err != nil {
		return err
	}

	if err := installFileIfChanged(filepath.Join(configPath, "kubeadm-config.yaml"), filepath.Join(deckhousePath, "kubeadm", "config.yaml"), 0600); err != nil {
		return err
	}
	for _, component := range []string{"etcd", "kube-apiserver", "kube-controller-manager", "kube-scheduler"} {
		if err := installFileIfChanged(filepath.Join(configPath, component+".yaml.tpl"), filepath.Join(deckhousePath, "kubeadm", "patches", component+".yaml"), 0600); err != nil {
			return err
		}
	}
	return nil
}

func waitImageHolderContainers() error {
	for {
		log.Info("phase: waiting for all image-holder containers will be ready")
		pod, err := config.K8sClient.CoreV1().Pods(namespace).Get(context.TODO(), config.MyPodName, metav1.GetOptions{})
		if err != nil {
			return err
		}
		isReady := true
		for _, container := range pod.Status.ContainerStatuses {
			if container.Name == "control-plane-manager" {
				continue
			}
			if !container.Ready {
				isReady = false
				break
			}
		}

		if isReady {
			return nil
		}
		time.Sleep(10 * time.Second)
	}
}

func checkEtcdManifest() error {
	etcdManifestPath := filepath.Join(manifestsPath, "etcd.yaml")
	if env := os.Getenv("D8_TESTS"); env == "yes" {
		etcdManifestPath = "testdata/etcd.yaml"
	}

	log.Info("phase: check etcd manifest", slog.String("path", etcdManifestPath))

	if _, err := os.Stat(etcdManifestPath); err != nil {
		log.Warn("etcd manifest absent", slog.String("path", etcdManifestPath))
		return nil
	}

	content, err := os.ReadFile(etcdManifestPath)
	if err != nil {
		return err
	}

	pod := &v1.Pod{}

	if err := yaml.Unmarshal(content, pod); err != nil {
		return err
	}

	foundAdvertiseParam := false
	foundNameParam := false
	foundDatadirParam := false
	for _, arg := range pod.Spec.Containers[0].Command {
		switch {
		case strings.HasPrefix(arg, "--advertise-client-urls=https://"):
			ip := strings.TrimPrefix(arg, "--advertise-client-urls=https://")
			ip = strings.TrimSuffix(strings.TrimPrefix(arg, "--advertise-client-urls=https://"), ":2379")
			if ip != config.MyIP {
				return fmt.Errorf("etcd is not supposed to change advertise address from %s to %s, please verify node's InternalIP", ip, config.MyIP)
			}
			foundAdvertiseParam = true
		case strings.HasPrefix(arg, "--name="):
			if name := strings.TrimPrefix(arg, "--name="); name != config.NodeName {
				return fmt.Errorf("etcd is not supposed to change its name from %s to %s, please verify node's hostname", name, config.NodeName)
			}
			foundNameParam = true
		case strings.HasPrefix(arg, "--data-dir="):
			if name := strings.TrimPrefix(arg, "--data-dir="); name != "/var/lib/etcd" {
				return fmt.Errorf("etcd is not supposed to change data-dir from %s to /var/lib/etcd, please verify current --data-dir", name)
			}
			foundDatadirParam = true
		}
	}

	if !foundAdvertiseParam {
		return fmt.Errorf("cannot find --advertise-client-urls submatch in etcd manifest %s", etcdManifestPath)
	}

	if !foundNameParam {
		return fmt.Errorf("cannot find --name submatch in etcd manifest %s", etcdManifestPath)
	}

	if !foundDatadirParam {
		return fmt.Errorf("cannot find --data-dir submatch in etcd manifest %s", etcdManifestPath)
	}

	return nil
}
