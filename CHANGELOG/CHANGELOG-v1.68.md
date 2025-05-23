# Changelog v1.68

## [MALFORMED]


 - #12058 unknown section "ceph-csi"

## Know before update


 - All data sources created with the GrafanaAdditionalDatasource custom resource will have their UIDs changed.
 - Ingress-nginx pods will restart. The validation of the generated NGINX configuration during the validation of `Ingress` resources is disabled until finding a way of running validations safely.
 - NGINX Ingress controller pods will be restarted.

## Features


 - **[candi]** Add a new layout for HuaweiCloud provider with support for VPC peering connections. [#12224](https://github.com/deckhouse/deckhouse/pull/12224)
 - **[candi]** Added a way to manage node labels from local files. [#10905](https://github.com/deckhouse/deckhouse/pull/10905)
 - **[cert-manager]** Bump cert-manager version up to 1.17.1. This fixes acme confirmation for cloudflare provider. [#12808](https://github.com/deckhouse/deckhouse/pull/12808)
 - **[cloud-provider-aws]** Bumped the `aws-node-termination-handler` version to `1.22.1`. [#10861](https://github.com/deckhouse/deckhouse/pull/10861)
 - **[cloud-provider-aws]** Added an option to specify your IAM role. [#10439](https://github.com/deckhouse/deckhouse/pull/10439)
 - **[cni-cilium]** DVP-specific feature that allows having two pods in the same cluster with equal IP addresses. [#11138](https://github.com/deckhouse/deckhouse/pull/11138)
 - **[control-plane-manager]** Improved free space detection for etcd backups. [#11170](https://github.com/deckhouse/deckhouse/pull/11170)
 - **[dhctl]** Enabled parallel bootstrap for cloud permanent node groups. [#11031](https://github.com/deckhouse/deckhouse/pull/11031)
 - **[dhctl]** Added restrictions to the `dhctl` command according to environment variable `DHCTL_CLI_ALLOWED_COMMANDS`. [#11007](https://github.com/deckhouse/deckhouse/pull/11007)
 - **[dhctl]** Added a new `--tf-resource-management-timeout` flag and set TF default timeout to 10 minutes. [#10843](https://github.com/deckhouse/deckhouse/pull/10843)
 - **[docs]** Added documentation for integrating with the Huawei Cloud provider. [#11094](https://github.com/deckhouse/deckhouse/pull/11094)
 - **[extended-monitoring]** Added IAM authentication support for ECR in image-availability-exporter. [#10893](https://github.com/deckhouse/deckhouse/pull/10893)
 - **[extended-monitoring]** Namespace-level `threshold.extended-monitoring.flant.com` labels now override default threshold values for objects lacking the label. [#10880](https://github.com/deckhouse/deckhouse/pull/10880)
 - **[extended-monitoring]** Replaced `cert-exporter` with `x509-certificate-exporter`, adding support for custom Opaque certificates. [#10554](https://github.com/deckhouse/deckhouse/pull/10554)
 - **[istio]** Added automatic L7 protocol detection for public services in Federation for proper handling on remote clusters. [#11808](https://github.com/deckhouse/deckhouse/pull/11808)
 - **[istio]** Added `alliance.ingressGateway.advertise` MC option to override auto-discovered public ingress gateway addresses. [#11726](https://github.com/deckhouse/deckhouse/pull/11726)
 - **[istio]** Added Istio components built from source, with known security issues fixed. [#11189](https://github.com/deckhouse/deckhouse/pull/11189)
 - **[log-shipper]** Added support for referencing Kubernetes secrets in TLS configuration for ClusterLogDestination resources. [#11159](https://github.com/deckhouse/deckhouse/pull/11159)
 - **[log-shipper]** Added the `keepDeletedFilesOpenedFor` option to Vector to set the duration of keeping deleted files open when log storage is unavailable. [#10641](https://github.com/deckhouse/deckhouse/pull/10641)
 - **[multitenancy-manager]** Added the ability to view installed resources in the project status.
    Also, non-existent resources in the template are now automatically checked and skipped during project rendering. [#11898](https://github.com/deckhouse/deckhouse/pull/11898)
 - **[prometheus]** add grafana enable flag [#12460](https://github.com/deckhouse/deckhouse/pull/12460)
 - **[upmeter]** Added `backup.deckhouse.io/cluster-config` label to relevant module CRDs. [#10568](https://github.com/deckhouse/deckhouse/pull/10568)
 - **[user-authn]** Added `highAvailability` option to the DexAuthenticator custom resource. [#11049](https://github.com/deckhouse/deckhouse/pull/11049)
    All authenticators deployed in Deckhouse will inherit the `highAvailability` option from the corresponding module instead of using this option's value from the `user-authn` module. Therefore, for example, if the `prometheus` module runs in the non-HA mode, the DexAuthenticator instance will also be deployed without HA.

## Fixes


 - **[admission-policy-engine]** Fix the trivy provider hook to support anonymous registry cred secret. [#12109](https://github.com/deckhouse/deckhouse/pull/12109)
 - **[candi]** Reduce `request-timeout` in kubernetes-api to 60s. [#12191](https://github.com/deckhouse/deckhouse/pull/12191)
    kubernetes-api will be restarted.
 - **[candi]** Fix bashible step that disable firewalld. [#12126](https://github.com/deckhouse/deckhouse/pull/12126)
 - **[candi]** Escaped character `%` in systemd for work proxy in containerd configuration. [#11483](https://github.com/deckhouse/deckhouse/pull/11483)
 - **[candi]** Fixed security context settings for static pods and iptables containers. [#11309](https://github.com/deckhouse/deckhouse/pull/11309)
 - **[candi]** Set `dhctl` to ignore metadata changes in the `vcd_vapp_vm` resource. [#11201](https://github.com/deckhouse/deckhouse/pull/11201)
 - **[chrony]** Fixed a security vulnerability in `chrony-exporter`. [#11558](https://github.com/deckhouse/deckhouse/pull/11558)
 - **[chrony]** Fixed a port conflict. [#11449](https://github.com/deckhouse/deckhouse/pull/11449)
 - **[chrony]** Added the missing libcap library. [#11431](https://github.com/deckhouse/deckhouse/pull/11431)
 - **[cloud-provider-azure]** CVE fix, bump go, use images, src-artifact. [#11630](https://github.com/deckhouse/deckhouse/pull/11630)
 - **[cloud-provider-dynamix]** Fixed several minor bugs. [#11235](https://github.com/deckhouse/deckhouse/pull/11235)
 - **[cloud-provider-gcp]** Fix CVE, bump go, use images, src-artifact. [#11632](https://github.com/deckhouse/deckhouse/pull/11632)
 - **[cloud-provider-huaweicloud]** CVE fix, bump go to 1.23, use images, src-artifact. [#11631](https://github.com/deckhouse/deckhouse/pull/11631)
 - **[cloud-provider-openstack]** Bumped Go module dependencies to resolve known security issues. [#11634](https://github.com/deckhouse/deckhouse/pull/11634)
 - **[cloud-provider-vcd]** fix CVE, bump ccm, use go 1.23, use werf images and src-artifact [#11629](https://github.com/deckhouse/deckhouse/pull/11629)
 - **[cloud-provider-vcd]** Added `vAppTemplate` search by `vAppTemplate` identifier for Cluster API Provider Cloud Director. [#10902](https://github.com/deckhouse/deckhouse/pull/10902)
 - **[cloud-provider-vcd]** Backported a fix to `vcd-csi-plugin-legacy`. [#10776](https://github.com/deckhouse/deckhouse/pull/10776)
 - **[cloud-provider-vsphere]** CVE fix cloud-provider-vsphere. [#12125](https://github.com/deckhouse/deckhouse/pull/12125)
 - **[cloud-provider-vsphere]** Bump go, use images, src-artifact. [#11628](https://github.com/deckhouse/deckhouse/pull/11628)
 - **[cloud-provider-zvirt]** fix invalid zvirt csi template [#12961](https://github.com/deckhouse/deckhouse/pull/12961)
 - **[cloud-provider-zvirt]** fix zvirt-csi-driver patching (token refresh fix patch) [#12885](https://github.com/deckhouse/deckhouse/pull/12885)
 - **[cloud-provider-zvirt]** Removed `DisksMeta` from `cloud-data-discoverer` to prevent possible crashes. [#11511](https://github.com/deckhouse/deckhouse/pull/11511)
 - **[cloud-provider-zvirt]** Fixed a bug that prevented zVirt SDK API token from refreshing upon expiration. [#11040](https://github.com/deckhouse/deckhouse/pull/11040)
 - **[cni-cilium]** L2 neigh discovery (arp flood) disabled. [#12078](https://github.com/deckhouse/deckhouse/pull/12078)
 - **[common]** Fixed a security vulnerability in `task`. [#11572](https://github.com/deckhouse/deckhouse/pull/11572)
 - **[control-plane-manager]** Fixed a security vulnerability. [#11536](https://github.com/deckhouse/deckhouse/pull/11536)
 - **[deckhouse]** Fix checking automatic kubernetes version in managed clusters. [#12685](https://github.com/deckhouse/deckhouse/pull/12685)
 - **[deckhouse]** Explicitly set ports in d8-kube-dns deployment. [#12658](https://github.com/deckhouse/deckhouse/pull/12658)
 - **[deckhouse]** Set EnabledByModuleConfig condition by default for modules without this condition. [#12558](https://github.com/deckhouse/deckhouse/pull/12558)
 - **[deckhouse]** Add initial module config validation. [#12465](https://github.com/deckhouse/deckhouse/pull/12465)
 - **[deckhouse]** Add initial module config validation. [#12227](https://github.com/deckhouse/deckhouse/pull/12227)
 - **[deckhouse]** Fix module.yaml ignored when weight == 0. [#12254](https://github.com/deckhouse/deckhouse/pull/12254)
 - **[deckhouse]** Trigger the ModuleConfig handling when a module created. [#12039](https://github.com/deckhouse/deckhouse/pull/12039)
 - **[deckhouse]** Return `helm_lib_envs_for_proxy` template to deckhouse deployment. [#11921](https://github.com/deckhouse/deckhouse/pull/11921)
 - **[deckhouse-controller]** Fix DeckhouseRelease approval in the Manual update mode. [#12527](https://github.com/deckhouse/deckhouse/pull/12527)
 - **[deckhouse-controller]** Fixed a bug in the `kubectl` wrapper script. [#11494](https://github.com/deckhouse/deckhouse/pull/11494)
 - **[deckhouse-tools]** Fix tools links pointing to wrong path [#12346](https://github.com/deckhouse/deckhouse/pull/12346)
 - **[dhctl]** Do not create a local-converger lock for commanderMode [#13454](https://github.com/deckhouse/deckhouse/pull/13454)
 - **[dhctl]** Escape special shell characters in sudo preflight check. [#12533](https://github.com/deckhouse/deckhouse/pull/12533)
 - **[dhctl]** sudoPassword from the connection-config not used in dhctl CLI and dhctl server [#12419](https://github.com/deckhouse/deckhouse/pull/12419)
 - **[dhctl]** Fix broken connection-config in CLI and support sudoPassword in SSHConfig. [#12036](https://github.com/deckhouse/deckhouse/pull/12036)
 - **[dhctl]** Fixed a bug that prevented PersistentVolumes from being properly deleted by the `dhctl destroy` command in Stronghold. [#11878](https://github.com/deckhouse/deckhouse/pull/11878)
 - **[dhctl]** Fixed the behavior of static cluster destruction. [#11059](https://github.com/deckhouse/deckhouse/pull/11059)
 - **[docs]** Update Deckhouse CLI installation guide [#12573](https://github.com/deckhouse/deckhouse/pull/12573)
 - **[docs]** Fixed a security vulnerability in `docs-builder`. [#11583](https://github.com/deckhouse/deckhouse/pull/11583)
 - **[extended-monitoring]** Fixed security vulnerabilities in `events-exporter`. [#11442](https://github.com/deckhouse/deckhouse/pull/11442)
 - **[extended-monitoring]** Fixed security vulnerabilities in `extended-monitoring-exporter`. [#11425](https://github.com/deckhouse/deckhouse/pull/11425)
 - **[extended-monitoring]** Fixed security vulnerabilities in `image-availability-exporter`. [#11336](https://github.com/deckhouse/deckhouse/pull/11336)
 - **[ingress-nginx]** latest CVE fixes backported. [#12746](https://github.com/deckhouse/deckhouse/pull/12746)
    Ingress-nginx pods will restart. The validation of the generated NGINX configuration during the validation of `Ingress` resources is disabled until finding a way of running validations safely.
 - **[ingress-nginx]** Fixed graceful shutdown handling for Ingress controller pods. [#11656](https://github.com/deckhouse/deckhouse/pull/11656)
    NGINX Ingress controller pods will be restarted.
 - **[ingress-nginx]** Fixed a bug related to dynamic libraries in the NGINX Ingress controller pods. [#11644](https://github.com/deckhouse/deckhouse/pull/11644)
 - **[ingress-nginx]** Fixed security context settings for static pods and iptables containers. [#11309](https://github.com/deckhouse/deckhouse/pull/11309)
 - **[ingress-nginx]** Fixed an issue in the controller when handling default backends for `ExternalName` services. [#11028](https://github.com/deckhouse/deckhouse/pull/11028)
 - **[ingress-nginx]** Fixed missing TLSv1.3 support when `legacySSL` is enabled. [#10871](https://github.com/deckhouse/deckhouse/pull/10871)
 - **[istio]** Return missed `librt.so.1` and `libpthread.so.0` libraries. [#12061](https://github.com/deckhouse/deckhouse/pull/12061)
    IngressGateway pods in d8-istio namespace will be restarted, it is advisable to restart application pods with Istio sidecars.
 - **[local-path-provisioner]** Fixed security context settings for `helper-pod`. [#11322](https://github.com/deckhouse/deckhouse/pull/11322)
 - **[loki]** Fixed minor security vulnerabilities. [#11530](https://github.com/deckhouse/deckhouse/pull/11530)
 - **[loki]** Fixed the build parameters. [#11399](https://github.com/deckhouse/deckhouse/pull/11399)
 - **[metallb]** Add clearing of alert metric group D8MetallbObsoleteLayer2PoolsAreUsed [#12648](https://github.com/deckhouse/deckhouse/pull/12648)
 - **[metallb]** Fixed module startup error with empty Selector value in SVC. [#12328](https://github.com/deckhouse/deckhouse/pull/12328)
 - **[metallb]** Bumped Go module dependencies to fix security vulnerabilities. [#11650](https://github.com/deckhouse/deckhouse/pull/11650)
 - **[monitoring-applications]** Fix module dependencies. The `monitoring-applications` module was affected and disabled by the enabled script. [#12428](https://github.com/deckhouse/deckhouse/pull/12428)
 - **[monitoring-kubernetes]** Fixed security vulnerabilities. [#11591](https://github.com/deckhouse/deckhouse/pull/11591)
 - **[multitenancy-manager]** Set each Falco rule to work for a specific namespace only. [#11498](https://github.com/deckhouse/deckhouse/pull/11498)
 - **[network-gateway]** Fixed security context settings for static pods and iptables containers. [#11309](https://github.com/deckhouse/deckhouse/pull/11309)
 - **[node-local-dns]** Fix module dependencies. The `monitoring-applications` module was affected and disabled by the enabled script. [#12428](https://github.com/deckhouse/deckhouse/pull/12428)
 - **[node-local-dns]** Fixed security context settings for static pods and iptables containers. [#11309](https://github.com/deckhouse/deckhouse/pull/11309)
 - **[node-manager]** partial rollback of changes in images cluster-autoscaler [#12910](https://github.com/deckhouse/deckhouse/pull/12910)
 - **[node-manager]** partial rollback of changes in images cluster-autoscaler [#12854](https://github.com/deckhouse/deckhouse/pull/12854)
 - **[node-manager]** fix autoscaler RBAC for zvirt and dynamix [#11236](https://github.com/deckhouse/deckhouse/pull/11236)
 - **[node-manager]** Fixed CVE in images cluster-autoscaler and machine-controller-manager. [#12024](https://github.com/deckhouse/deckhouse/pull/12024)
 - **[node-manager]** Deleted `cert-manager.io/inject-ca-from` annotation from CAPI `ValidatingWebhookConfiguration`. [#11666](https://github.com/deckhouse/deckhouse/pull/11666)
 - **[node-manager]** Changes to `nodeGroup.spec.staticInstances.labelSelector` are now blocked to prevent unexpected behavior. [#11111](https://github.com/deckhouse/deckhouse/pull/11111)
 - **[okmeter]** drop service account fields from the okmeter daemonset manifest [#12303](https://github.com/deckhouse/deckhouse/pull/12303)
    okmeter daemonset pods will be rollout restarted
 - **[okmeter]** Removed RBAC from the `okmeter` module since the Kubernetes service account is currently unused and non-functional. [#10323](https://github.com/deckhouse/deckhouse/pull/10323)
    Okmeter agents are no longer able to access the Kubernetes API.
 - **[operator-trivy]** Added an init container to the `trivy-server` to set the correct user:group ownership for its volume. [#11671](https://github.com/deckhouse/deckhouse/pull/11671)
 - **[operator-trivy]** Fixed security vulnerabilities and updated the `report-updater` build. [#11597](https://github.com/deckhouse/deckhouse/pull/11597)
 - **[prometheus]** Enables WAL for Grafana SQLite database to prevent locking errors, thus fixing in-dashboard alerting. [#13068](https://github.com/deckhouse/deckhouse/pull/13068)
    Grafana deployment will be rollout restarted.
 - **[prometheus]** Changes the links in alert's template about alerts from url to relative links [#12028](https://github.com/deckhouse/deckhouse/pull/12028)
    The links in the alerts will be correct when the documentation module is enabled.
 - **[prometheus]** Fixed an issue in the Grafana datasource. [#11788](https://github.com/deckhouse/deckhouse/pull/11788)
    All data sources created with the GrafanaAdditionalDatasource custom resource will have their UIDs changed.
 - **[prometheus]** Improved Prometheus rolling updates. [#11694](https://github.com/deckhouse/deckhouse/pull/11694)
    The Prometheus StatefulSet will undergo a rollout restart.
 - **[prometheus]** Added the missing Monaco JS library to the `grafana-v10` image. [#11667](https://github.com/deckhouse/deckhouse/pull/11667)
    Grafana deployment pods will undergo a rollout restart.
 - **[prometheus]** Fixed security vulnerabilities in `alertmanager`. [#11459](https://github.com/deckhouse/deckhouse/pull/11459)
 - **[prometheus]** Fixed security vulnerabilities in the `grafana-v10` image. [#11417](https://github.com/deckhouse/deckhouse/pull/11417)
    Grafana in cluster will be restarted as the image has changed.
 - **[prometheus]** Fixed the `loki` build parameters. [#11399](https://github.com/deckhouse/deckhouse/pull/11399)
 - **[prometheus]** Fixed security vulnerabilities in `trickster`. [#11353](https://github.com/deckhouse/deckhouse/pull/11353)
 - **[prometheus]** Fixed security vulnerabilities in `memcached`. [#11296](https://github.com/deckhouse/deckhouse/pull/11296)
 - **[prometheus]** Fixed security vulnerabilities in `alerts-receiver`. [#11257](https://github.com/deckhouse/deckhouse/pull/11257)
 - **[prometheus-metrics-adapter]** Fixed security vulnerabilities in `prometheus-reverse-proxy`. [#11472](https://github.com/deckhouse/deckhouse/pull/11472)
 - **[prometheus-metrics-adapter]** Fixed security vulnerabilities in `k8s-prometheus-adapter`. [#11389](https://github.com/deckhouse/deckhouse/pull/11389)
 - **[prometheus-pushgateway]** CVE update. [#12116](https://github.com/deckhouse/deckhouse/pull/12116)
 - **[registrypackages]** CVE update ec2-describe-tags. [#12115](https://github.com/deckhouse/deckhouse/pull/12115)
 - **[registrypackages]** Bumped the `docker_auth` version in in `docker-registry`. [#11600](https://github.com/deckhouse/deckhouse/pull/11600)
 - **[registrypackages]** Fixed security vulnerabilities in `crictl`. [#11523](https://github.com/deckhouse/deckhouse/pull/11523)
 - **[runtime-audit-engine]** Fix CrashLoopBack due to invalid config. [#13144](https://github.com/deckhouse/deckhouse/pull/13144)
 - **[runtime-audit-engine]** copy falcosecurity_falcosidekick_falco_events_total metric to falco_events for backward compatibility [#12873](https://github.com/deckhouse/deckhouse/pull/12873)
 - **[runtime-audit-engine]** Fixed security vulnerabilities in `k8s-metacollector` and `rules-loader`. [#11581](https://github.com/deckhouse/deckhouse/pull/11581)
 - **[service-with-healthchecks]** Fixed several issues. [#12272](https://github.com/deckhouse/deckhouse/pull/12272)
 - **[terraform-manager]** CVE fix dynamix cloud provider and bump go version. [#12111](https://github.com/deckhouse/deckhouse/pull/12111)
 - **[terraform-manager]** CVE fix terraform-manager-gcp and bump go version. [#12068](https://github.com/deckhouse/deckhouse/pull/12068)
 - **[terraform-manager]** CVE fix terraform-manager-aws and bump go version. [#12065](https://github.com/deckhouse/deckhouse/pull/12065)
 - **[terraform-manager]** CVE fix terraform-manager-vcd and bump go version. [#12057](https://github.com/deckhouse/deckhouse/pull/12057)
 - **[terraform-manager]** CVE fix terraform-manager-huaweicloud and bump go version. [#12056](https://github.com/deckhouse/deckhouse/pull/12056)
 - **[terraform-manager]** CVE fix terraform-manager-zvirt and bump go version. [#11990](https://github.com/deckhouse/deckhouse/pull/11990)
 - **[terraform-manager]** CVE fix dynamix cloud provider and bump go version. [#11952](https://github.com/deckhouse/deckhouse/pull/11952)
 - **[terraform-manager]** Bumped the Go version to fix security invulnerabilities in `terraform-manager-yandex`. [#11924](https://github.com/deckhouse/deckhouse/pull/11924)
 - **[terraform-manager]** CVE fix terraform-manager-openstack and bump go version. [#11916](https://github.com/deckhouse/deckhouse/pull/11916)
 - **[terraform-manager]** Bumped the Go version to fix security invulnerabilities in `terraform-manager-vsphere`. [#11903](https://github.com/deckhouse/deckhouse/pull/11903)
 - **[terraform-manager]** Fixed the sizing policy change procedure in the `vcd_vapp_vm` resource. [#11202](https://github.com/deckhouse/deckhouse/pull/11202)
 - **[terraform-manager]** Removed the patch with `vcd_catalog_vapp_template` data source fixes for `terraform-provider-vcd`. [#10902](https://github.com/deckhouse/deckhouse/pull/10902)
 - **[upmeter]** Fixed metallb module probe. [#12231](https://github.com/deckhouse/deckhouse/pull/12231)
 - **[user-authn]** Fixed security vulnerabilities and updated the build. [#11588](https://github.com/deckhouse/deckhouse/pull/11588)
 - **[user-authn]** Bumped dependencies to fix security vulnerabilities. [#11332](https://github.com/deckhouse/deckhouse/pull/11332)

## Chore


 - **[admission-policy-engine]** Fixes policies schemas. Allow to exclude namespaces or use only object labels. [#12292](https://github.com/deckhouse/deckhouse/pull/12292)
 - **[admission-policy-engine]** Fixed security vulnerabilities and updated the `trivy` build. [#11568](https://github.com/deckhouse/deckhouse/pull/11568)
 - **[admission-policy-engine]** Bumped `gatekeeper` and refactored the build. [#11356](https://github.com/deckhouse/deckhouse/pull/11356)
 - **[candi]** Update Deckhouse CLI to 0.11.3 [#12691](https://github.com/deckhouse/deckhouse/pull/12691)
 - **[candi]** update d8-cli [#12104](https://github.com/deckhouse/deckhouse/pull/12104)
 - **[candi]** Bump patch versions of Kubernetes images: `v1.29.14`, `v1.30.1`, `v1.31.6` [#12081](https://github.com/deckhouse/deckhouse/pull/12081)
    Kubernetes control-plane components will restart, kubelet will restart.
 - **[candi]** Bump patch versions of Kubernetes images: `v1.29.13`, `v1.30.9`, `v1.31.5` [#12016](https://github.com/deckhouse/deckhouse/pull/12016)
    Kubernetes control-plane components will restart, kubelet will restart.
 - **[candi]** Bumped `kubernetes-cni` and `flannel-cni` binaries. [#11604](https://github.com/deckhouse/deckhouse/pull/11604)
 - **[candi]** Updated the Go version to `1.23.4` in base images. [#11241](https://github.com/deckhouse/deckhouse/pull/11241)
 - **[candi]** Updated the `log-shipper` version to `0.43.1`. [#11108](https://github.com/deckhouse/deckhouse/pull/11108)
 - **[cert-manager]** Bumped to version `1.16.2` and updated build parameters. [#11198](https://github.com/deckhouse/deckhouse/pull/11198)
 - **[chrony]** Updated the `chrony` version to `4.6.1`. [#11403](https://github.com/deckhouse/deckhouse/pull/11403)
 - **[chrony]** Added `chrony-exporter` templates. [#10976](https://github.com/deckhouse/deckhouse/pull/10976)
 - **[cilium-hubble]** Newly discovered CVEs have been addressed. [#12053](https://github.com/deckhouse/deckhouse/pull/12053)
 - **[cloud-provider-aws]** bump go, use images, src-artifact [#12155](https://github.com/deckhouse/deckhouse/pull/12155)
 - **[cloud-provider-aws]** bump go, use images, src-artifact [#11633](https://github.com/deckhouse/deckhouse/pull/11633)
 - **[cloud-provider-aws]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cloud-provider-aws]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cloud-provider-azure]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cloud-provider-azure]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cloud-provider-dynamix]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cloud-provider-dynamix]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cloud-provider-gcp]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cloud-provider-gcp]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cloud-provider-openstack]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cloud-provider-openstack]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cloud-provider-vcd]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cloud-provider-vcd]** Bumped `capcd-controller-manager` to version `1.3.2`. [#10823](https://github.com/deckhouse/deckhouse/pull/10823)
 - **[cloud-provider-vcd]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cloud-provider-vsphere]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cloud-provider-vsphere]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cloud-provider-yandex]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cloud-provider-yandex]** Fixed security vulnerabilities. [#11765](https://github.com/deckhouse/deckhouse/pull/11765)
 - **[cloud-provider-yandex]** Bumped Go module dependencies. [#11585](https://github.com/deckhouse/deckhouse/pull/11585)
 - **[cloud-provider-yandex]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cloud-provider-zvirt]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cloud-provider-zvirt]** Fixed the zVirt CSI build. [#11809](https://github.com/deckhouse/deckhouse/pull/11809)
 - **[cloud-provider-zvirt]** Bumped Go module dependencies to resolve known security issues. [#11510](https://github.com/deckhouse/deckhouse/pull/11510)
 - **[cloud-provider-zvirt]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cni-cilium]** Newly discovered CVEs have been addressed. [#12053](https://github.com/deckhouse/deckhouse/pull/12053)
 - **[cni-cilium]** A new parameter, `activeL2NeighborDiscoveryEnabled`, has been added to ModuleConfig of `cni-cilium. [#12203](https://github.com/deckhouse/deckhouse/pull/12203)
 - **[cni-cilium]** Improved the stability of the cilium-operator. [#12091](https://github.com/deckhouse/deckhouse/pull/12091)
 - **[cni-cilium]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cni-cilium]** Refactored the build and fixed security vulnerabilities. [#11382](https://github.com/deckhouse/deckhouse/pull/11382)
 - **[cni-cilium]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[cni-flannel]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[cni-flannel]** Bumped `kubernetes-cni` and `flannel-cni` binaries. [#11604](https://github.com/deckhouse/deckhouse/pull/11604)
 - **[cni-flannel]** Refactored the build and bumped to version `0.26.2`. [#11344](https://github.com/deckhouse/deckhouse/pull/11344)
 - **[cni-flannel]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[common]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[common]** Bumped `yq4` to version `3.45.1`. [#11574](https://github.com/deckhouse/deckhouse/pull/11574)
 - **[common]** Updated dependent libraries. [#11348](https://github.com/deckhouse/deckhouse/pull/11348)
 - **[common]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[control-plane-manager]** Replaced Logrus to implement Slog. [#10909](https://github.com/deckhouse/deckhouse/pull/10909)
    Components need to be restarted.
 - **[dashboard]** up log level  and version component [#12105](https://github.com/deckhouse/deckhouse/pull/12105)
 - **[deckhouse]** Refactored RBAC configuration. [#11605](https://github.com/deckhouse/deckhouse/pull/11605)
 - **[descheduler]** Updated the build and fixed security vulnerabilities. [#11221](https://github.com/deckhouse/deckhouse/pull/11221)
 - **[dhctl]** Add opportunity to bootstrap terranodes parallel with env DHCTL_PARALLEL_CLOUD_PERMANENT_NODES_BOOTSTRAP revert sequentially by default [#12560](https://github.com/deckhouse/deckhouse/pull/12560)
 - **[dhctl]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[dhctl]** Replaced Logrus to implement Slog and deleted the "simple" logger. [#10909](https://github.com/deckhouse/deckhouse/pull/10909)
 - **[dhctl]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[docs]** Update storage section in the documentation. [#12394](https://github.com/deckhouse/deckhouse/pull/12394)
 - **[docs]** Updated d8 mirror-related FAQ and Deckhouse CLI page in preparation for d8-cli v0.10 GA release [#12019](https://github.com/deckhouse/deckhouse/pull/12019)
 - **[documentation]** Fix bugs is navigation. [#12505](https://github.com/deckhouse/deckhouse/pull/12505)
 - **[documentation]** Bumped the `docs-builder` dependencies. [#11742](https://github.com/deckhouse/deckhouse/pull/11742)
 - **[global-hooks]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[global-hooks]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[go_lib]** Reverted CNI configuration to the approach used in Deckhouse 1.67, using the `d8-cni-configuration` secret. [#11935](https://github.com/deckhouse/deckhouse/pull/11935)
 - **[go_lib]** (Reverted) Prepared to migrate from using the `d8-cni-configuration` secret to ModuleConfigs for CNI configuration. [#10517](https://github.com/deckhouse/deckhouse/pull/10517)
 - **[ingress-nginx]** Refactored the build and updated dependent libraries to fix security vulnerabilities. [#11518](https://github.com/deckhouse/deckhouse/pull/11518)
 - **[ingress-nginx]** Disabled the VPA for the Istio sidecar. [#11084](https://github.com/deckhouse/deckhouse/pull/11084)
 - **[ingress-nginx]** Added a release requirement check for the `defaultControllerVersion` parameter. [#10941](https://github.com/deckhouse/deckhouse/pull/10941)
 - **[ingress-nginx]** Implemented filtering for `vhost` with multiple Ingress controllers in Grafana. [#10847](https://github.com/deckhouse/deckhouse/pull/10847)
 - **[istio]** Migrate Kiali images to werf and build from source [#12149](https://github.com/deckhouse/deckhouse/pull/12149)
 - **[istio]** Fix new CVE in Istio version 1.19.7. [#12037](https://github.com/deckhouse/deckhouse/pull/12037)
 - **[log-shipper]** Updated `vector` to version `0.44`. [#11595](https://github.com/deckhouse/deckhouse/pull/11595)
 - **[log-shipper]** Updated the `log-shipper` version to `0.43.1`. [#11108](https://github.com/deckhouse/deckhouse/pull/11108)
 - **[metallb]** Refactored the build and fixed security vulnerabilities. [#11345](https://github.com/deckhouse/deckhouse/pull/11345)
 - **[monitoring-kubernetes]** Fixed security vulnerabilities. [#11371](https://github.com/deckhouse/deckhouse/pull/11371)
 - **[monitoring-ping]** Fixed security vulnerabilities. [#11374](https://github.com/deckhouse/deckhouse/pull/11374)
 - **[node-local-dns]** Refactored the build and updated dependent libraries. [#11421](https://github.com/deckhouse/deckhouse/pull/11421)
 - **[node-local-dns]** Replaced Logrus to implement Slog. [#10909](https://github.com/deckhouse/deckhouse/pull/10909)
    Components need to be restarted.
 - **[openvpn]** Replaced Logrus to implement Slog. [#10909](https://github.com/deckhouse/deckhouse/pull/10909)
    Components need to be restarted.
 - **[operator-trivy]** Fixed security vulnerabilities and updated the build. [#11568](https://github.com/deckhouse/deckhouse/pull/11568)
 - **[prometheus]** update Grafana version to 10.4.15 [#11977](https://github.com/deckhouse/deckhouse/pull/11977)
    grafana deployment will be rollout restarted
 - **[prometheus]** Bumped FE Prometheus version to `2.55.1`. [#11651](https://github.com/deckhouse/deckhouse/pull/11651)
 - **[prometheus]** Implemented a static Python binary for the Grafana dashboard provisioner. [#11534](https://github.com/deckhouse/deckhouse/pull/11534)
    The Grafana deployment will undergo a rollout restart.
 - **[prometheus]** Deprecated support for Grafana v8. [#10359](https://github.com/deckhouse/deckhouse/pull/10359)
 - **[registry-packages-proxy]** Replaced Logrus to implement Slog. [#10909](https://github.com/deckhouse/deckhouse/pull/10909)
    Components need to be restarted.
 - **[registrypackages]** Added a static `resize2fs` version to the `e2fsprogs` package. [#11906](https://github.com/deckhouse/deckhouse/pull/11906)
 - **[registrypackages]** Bumped `kubernetes-cni` and `flannel-cni` binaries. [#11604](https://github.com/deckhouse/deckhouse/pull/11604)
 - **[registrypackages]** Bumped `jq` to version `1.7.1`. [#11370](https://github.com/deckhouse/deckhouse/pull/11370)
 - **[runtime-audit-engine]** Refactored the build and fixed security vulnerabilities. [#11290](https://github.com/deckhouse/deckhouse/pull/11290)
 - **[service-with-healthchecks]** Updated dependencies and fixed security vulnerabilities. [#11713](https://github.com/deckhouse/deckhouse/pull/11713)
    The `service-with-healthchecks` pods will be restarted.
 - **[snapshot-controller]** Updated the build and bumped the version to `8.2.0`. [#11279](https://github.com/deckhouse/deckhouse/pull/11279)
 - **[static-routing-manager]** Updated dependencies and fixed security vulnerabilities. [#11711](https://github.com/deckhouse/deckhouse/pull/11711)
    The `static-routing-manager` pods will be restarted.
 - **[terraform-manager]** Fix terraform CVEs. [#12096](https://github.com/deckhouse/deckhouse/pull/12096)
 - **[vertical-pod-autoscaler]** Updated the build and fixed security vulnerabilities. [#11219](https://github.com/deckhouse/deckhouse/pull/11219)

