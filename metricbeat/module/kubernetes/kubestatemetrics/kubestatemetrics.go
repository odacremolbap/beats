// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package controllermanager

import (
	"github.com/elastic/beats/metricbeat/helper/prometheus"
	"github.com/elastic/beats/metricbeat/mb"
)

func init() {

	mapping := &prometheus.MetricsMapping{
		Metrics: map[string]prometheus.MetricMap{
			"process_cpu_seconds_total":                   prometheus.Metric("process.cpu.sec"),
			"process_resident_memory_bytes":               prometheus.Metric("process.memory.resident.bytes"),
			"process_virtual_memory_bytes":                prometheus.Metric("process.memory.virtual.bytes"),
			"process_open_fds":                            prometheus.Metric("process.fds.open.count"),
			"process_start_time_seconds":                  prometheus.Metric("process.started.sec"),
			"http_request_duration_microseconds":          prometheus.Metric("http.request.duration.us"),
			"http_request_size_bytes":                     prometheus.Metric("http.request.size.bytes"),
			"http_response_size_bytes":                    prometheus.Metric("http.response.size.bytes"),
			"http_requests_total":                         prometheus.Metric("http.request.count"),
			"rest_client_requests_total":                  prometheus.Metric("client.request.count"),
			"workqueue_longest_running_processor_seconds": prometheus.Metric("workqueue.longestrunning.sec"),
			"workqueue_unfinished_work_seconds":           prometheus.Metric("workqueue.unfinished.sec"),
			"workqueue_adds_total":                        prometheus.Metric("workqueue.adds.count"),
			"workqueue_depth":                             prometheus.Metric("workqueue.depth.count"),
			"workqueue_retries_total":                     prometheus.Metric("workqueue.retries.count"),
			"node_collector_evictions_number":             prometheus.Metric("node.collector.eviction.count"),
			"node_collector_unhealthy_nodes_in_zone":      prometheus.Metric("node.collector.unhealthy.count"),
			"node_collector_zone_size":                    prometheus.Metric("node.collector.count"),
			"node_collector_zone_health":                  prometheus.Metric("node.collector.health.pct"),
			"leader_election_master_status":               prometheus.BooleanMetric("leader.is_master"),
		},

		Labels: map[string]prometheus.LabelMap{
			"handler": prometheus.KeyLabel("handler"),
			"code":    prometheus.KeyLabel("code"),
			"method":  prometheus.KeyLabel("method"),
			"host":    prometheus.KeyLabel("host"),
			"name":    prometheus.KeyLabel("name"),
			"zone":    prometheus.KeyLabel("zone"),
		},
	}

	mb.Registry.MustAddMetricSet("kubernetes", "kubestatemetrics",
		prometheus.MetricSetBuilder(mapping),
		mb.WithHostParser(prometheus.HostParser))
}

func certificateMappigs() map[string]prometheus.MetricMap {
	/*
		# HELP kube_certificatesigningrequest_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_certificatesigningrequest_labels gauge
		kube_certificatesigningrequest_labels{certificatesigningrequest="csr-7snpd"} 1
		# HELP kube_certificatesigningrequest_created Unix creation timestamp
		# TYPE kube_certificatesigningrequest_created gauge
		kube_certificatesigningrequest_created{certificatesigningrequest="csr-7snpd"} 1.563532766e+09
		# HELP kube_certificatesigningrequest_condition The number of each certificatesigningrequest condition
		# TYPE kube_certificatesigningrequest_condition gauge
		kube_certificatesigningrequest_condition{certificatesigningrequest="csr-7snpd",condition="approved"} 1
		kube_certificatesigningrequest_condition{certificatesigningrequest="csr-7snpd",condition="denied"} 0
		# HELP kube_certificatesigningrequest_cert_length Length of the issued cert
		# TYPE kube_certificatesigningrequest_cert_length gauge
		kube_certificatesigningrequest_cert_length{certificatesigningrequest="csr-7snpd"} 899
		# HELP kube_certificatesigningrequest_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_certificatesigningrequest_annotations gauge
		kube_certificatesigningrequest_annotations{certificatesigningrequest="csr-7snpd"} 1
	*/
	return nil
}

func configMapMappigs() map[string]prometheus.MetricMap {
	/*
		# HELP kube_configmap_info Information about configmap.
		# TYPE kube_configmap_info gauge
		kube_configmap_info{namespace="kube-system",configmap="coredns"} 1
		kube_configmap_info{namespace="kube-system",configmap="kube-proxy"} 1
		kube_configmap_info{namespace="kube-system",configmap="kindnet-cfg"} 1
		kube_configmap_info{namespace="kube-system",configmap="ip-masq-agent"} 1
		kube_configmap_info{namespace="kube-system",configmap="extension-apiserver-authentication"} 1
		kube_configmap_info{namespace="kube-system",configmap="kubeadm-config"} 1
		kube_configmap_info{namespace="kube-system",configmap="kubelet-config-1.14"} 1
		kube_configmap_info{namespace="kube-public",configmap="cluster-info"} 1
		# HELP kube_configmap_created Unix creation timestamp
		# TYPE kube_configmap_created gauge
		kube_configmap_created{namespace="kube-system",configmap="kube-proxy"} 1.563532771e+09
		kube_configmap_created{namespace="kube-system",configmap="kindnet-cfg"} 1.563532772e+09
		kube_configmap_created{namespace="kube-system",configmap="ip-masq-agent"} 1.563532772e+09
		kube_configmap_created{namespace="kube-system",configmap="extension-apiserver-authentication"} 1.563532767e+09
		kube_configmap_created{namespace="kube-system",configmap="kubeadm-config"} 1.56353277e+09
		kube_configmap_created{namespace="kube-system",configmap="kubelet-config-1.14"} 1.56353277e+09
		kube_configmap_created{namespace="kube-public",configmap="cluster-info"} 1.563532771e+09
		kube_configmap_created{namespace="kube-system",configmap="coredns"} 1.563532771e+09
		# HELP kube_configmap_metadata_resource_version Resource version representing a specific version of the configmap.
		# TYPE kube_configmap_metadata_resource_version gauge
		kube_configmap_metadata_resource_version{namespace="kube-system",configmap="kubelet-config-1.14",resource_version="162"} 1
		kube_configmap_metadata_resource_version{namespace="kube-public",configmap="cluster-info",resource_version="365"} 1
		kube_configmap_metadata_resource_version{namespace="kube-system",configmap="coredns",resource_version="198"} 1
		kube_configmap_metadata_resource_version{namespace="kube-system",configmap="kube-proxy",resource_version="219"} 1
		kube_configmap_metadata_resource_version{namespace="kube-system",configmap="kindnet-cfg",resource_version="249"} 1
		kube_configmap_metadata_resource_version{namespace="kube-system",configmap="ip-masq-agent",resource_version="259"} 1
		kube_configmap_metadata_resource_version{namespace="kube-system",configmap="extension-apiserver-authentication",resource_version="43"} 1
		kube_configmap_metadata_resource_version{namespace="kube-system",configmap="kubeadm-config",resource_version="158"} 1
		# HELP kube_configmap_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_configmap_annotations gauge
		kube_configmap_annotations{namespace="kube-system",configmap="kubeadm-config"} 1
		kube_configmap_annotations{namespace="kube-system",configmap="kubelet-config-1.14"} 1
		kube_configmap_annotations{namespace="kube-public",configmap="cluster-info"} 1
		kube_configmap_annotations{namespace="kube-system",configmap="coredns"} 1
		kube_configmap_annotations{namespace="kube-system",configmap="kube-proxy"} 1
		kube_configmap_annotations{namespace="kube-system",configmap="kindnet-cfg"} 1
		kube_configmap_annotations{namespace="kube-system",configmap="ip-masq-agent"} 1
		kube_configmap_annotations{namespace="kube-system",configmap="extension-apiserver-authentication"} 1
	*/
	return nil
}

func cronJobMappigs() *prometheus.MetricsMapping {

	mapping := &prometheus.MetricsMapping{
		Metrics: map[string]prometheus.MetricMap{
			"process_cpu_seconds_total":                   prometheus.Metric("process.cpu.sec"),
			"process_resident_memory_bytes":               prometheus.Metric("process.memory.resident.bytes"),
			"process_virtual_memory_bytes":                prometheus.Metric("process.memory.virtual.bytes"),
			"process_open_fds":                            prometheus.Metric("process.fds.open.count"),
			"process_start_time_seconds":                  prometheus.Metric("process.started.sec"),
			"http_request_duration_microseconds":          prometheus.Metric("http.request.duration.us"),
			"http_request_size_bytes":                     prometheus.Metric("http.request.size.bytes"),
			"http_response_size_bytes":                    prometheus.Metric("http.response.size.bytes"),
			"http_requests_total":                         prometheus.Metric("http.request.count"),
			"rest_client_requests_total":                  prometheus.Metric("client.request.count"),
			"workqueue_longest_running_processor_seconds": prometheus.Metric("workqueue.longestrunning.sec"),
			"workqueue_unfinished_work_seconds":           prometheus.Metric("workqueue.unfinished.sec"),
			"workqueue_adds_total":                        prometheus.Metric("workqueue.adds.count"),
			"workqueue_depth":                             prometheus.Metric("workqueue.depth.count"),
			"workqueue_retries_total":                     prometheus.Metric("workqueue.retries.count"),
			"node_collector_evictions_number":             prometheus.Metric("node.collector.eviction.count"),
			"node_collector_unhealthy_nodes_in_zone":      prometheus.Metric("node.collector.unhealthy.count"),
			"node_collector_zone_size":                    prometheus.Metric("node.collector.count"),
			"node_collector_zone_health":                  prometheus.Metric("node.collector.health.pct"),
			"leader_election_master_status":               prometheus.BooleanMetric("leader.is_master"),
		},

		Labels: map[string]prometheus.LabelMap{
			"handler": prometheus.KeyLabel("handler"),
			"code":    prometheus.KeyLabel("code"),
			"method":  prometheus.KeyLabel("method"),
			"host":    prometheus.KeyLabel("host"),
			"name":    prometheus.KeyLabel("name"),
			"zone":    prometheus.KeyLabel("zone"),
		},
	}

	/*
	   # HELP kube_cronjob_labels Kubernetes labels converted to Prometheus labels.
	   # TYPE kube_cronjob_labels gauge
	   # HELP kube_cronjob_info Info about cronjob.
	   # TYPE kube_cronjob_info gauge
	   # HELP kube_cronjob_created Unix creation timestamp
	   # TYPE kube_cronjob_created gauge
	   # HELP kube_cronjob_status_active Active holds pointers to currently running jobs.
	   # TYPE kube_cronjob_status_active gauge
	   # HELP kube_cronjob_status_last_schedule_time LastScheduleTime keeps information of when was the last time the job was successfully scheduled.
	   # TYPE kube_cronjob_status_last_schedule_time gauge
	   # HELP kube_cronjob_spec_suspend Suspend flag tells the controller to suspend subsequent executions.
	   # TYPE kube_cronjob_spec_suspend gauge
	   # HELP kube_cronjob_spec_starting_deadline_seconds Deadline in seconds for starting the job if it misses scheduled time for any reason.
	   # TYPE kube_cronjob_spec_starting_deadline_seconds gauge
	   # HELP kube_cronjob_next_schedule_time Next time the cronjob should be scheduled. The time after lastScheduleTime, or after the cron job's creation time if it's never been scheduled. Use this to determine if the job is delayed.
	   # TYPE kube_cronjob_next_schedule_time gauge
	   # HELP kube_cronjob_annotations Kubernetes annotations converted to Prometheus labels.
	   # TYPE kube_cronjob_annotations gauge
	*/
	return nil
}

func daemonSetMappings() map[string]prometheus.MetricMap {

	/*
	   # HELP kube_daemonset_created Unix creation timestamp
	   # TYPE kube_daemonset_created gauge
	   kube_daemonset_created{namespace="kube-system",daemonset="kube-proxy"} 1.563532771e+09
	   kube_daemonset_created{namespace="kube-system",daemonset="kindnet"} 1.563532772e+09
	   kube_daemonset_created{namespace="kube-system",daemonset="ip-masq-agent"} 1.563532772e+09
	   # HELP kube_daemonset_status_current_number_scheduled The number of nodes running at least one daemon pod and are supposed to.
	   # TYPE kube_daemonset_status_current_number_scheduled gauge
	   kube_daemonset_status_current_number_scheduled{namespace="kube-system",daemonset="kube-proxy"} 1
	   kube_daemonset_status_current_number_scheduled{namespace="kube-system",daemonset="kindnet"} 1
	   kube_daemonset_status_current_number_scheduled{namespace="kube-system",daemonset="ip-masq-agent"} 1
	   # HELP kube_daemonset_status_desired_number_scheduled The number of nodes that should be running the daemon pod.
	   # TYPE kube_daemonset_status_desired_number_scheduled gauge
	   kube_daemonset_status_desired_number_scheduled{namespace="kube-system",daemonset="kube-proxy"} 1
	   kube_daemonset_status_desired_number_scheduled{namespace="kube-system",daemonset="kindnet"} 1
	   kube_daemonset_status_desired_number_scheduled{namespace="kube-system",daemonset="ip-masq-agent"} 1
	   # HELP kube_daemonset_status_number_available The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available
	   # TYPE kube_daemonset_status_number_available gauge
	   kube_daemonset_status_number_available{namespace="kube-system",daemonset="kube-proxy"} 1
	   kube_daemonset_status_number_available{namespace="kube-system",daemonset="kindnet"} 1
	   kube_daemonset_status_number_available{namespace="kube-system",daemonset="ip-masq-agent"} 1
	   # HELP kube_daemonset_status_number_misscheduled The number of nodes running a daemon pod but are not supposed to.
	   # TYPE kube_daemonset_status_number_misscheduled gauge
	   kube_daemonset_status_number_misscheduled{namespace="kube-system",daemonset="kindnet"} 0
	   kube_daemonset_status_number_misscheduled{namespace="kube-system",daemonset="ip-masq-agent"} 0
	   kube_daemonset_status_number_misscheduled{namespace="kube-system",daemonset="kube-proxy"} 0
	   # HELP kube_daemonset_status_number_ready The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
	   # TYPE kube_daemonset_status_number_ready gauge
	   kube_daemonset_status_number_ready{namespace="kube-system",daemonset="kube-proxy"} 1
	   kube_daemonset_status_number_ready{namespace="kube-system",daemonset="kindnet"} 1
	   kube_daemonset_status_number_ready{namespace="kube-system",daemonset="ip-masq-agent"} 1
	   # HELP kube_daemonset_status_number_unavailable The number of nodes that should be running the daemon pod and have none of the daemon pod running and available
	   # TYPE kube_daemonset_status_number_unavailable gauge
	   kube_daemonset_status_number_unavailable{namespace="kube-system",daemonset="kube-proxy"} 0
	   kube_daemonset_status_number_unavailable{namespace="kube-system",daemonset="kindnet"} 0
	   kube_daemonset_status_number_unavailable{namespace="kube-system",daemonset="ip-masq-agent"} 0
	   # HELP kube_daemonset_updated_number_scheduled The total number of nodes that are running updated daemon pod
	   # TYPE kube_daemonset_updated_number_scheduled gauge
	   kube_daemonset_updated_number_scheduled{namespace="kube-system",daemonset="kube-proxy"} 1
	   kube_daemonset_updated_number_scheduled{namespace="kube-system",daemonset="kindnet"} 1
	   kube_daemonset_updated_number_scheduled{namespace="kube-system",daemonset="ip-masq-agent"} 1
	   # HELP kube_daemonset_metadata_generation Sequence number representing a specific generation of the desired state.
	   # TYPE kube_daemonset_metadata_generation gauge
	   kube_daemonset_metadata_generation{namespace="kube-system",daemonset="kube-proxy"} 1
	   kube_daemonset_metadata_generation{namespace="kube-system",daemonset="kindnet"} 1
	   kube_daemonset_metadata_generation{namespace="kube-system",daemonset="ip-masq-agent"} 1
	   # HELP kube_daemonset_labels Kubernetes labels converted to Prometheus labels.
	   # TYPE kube_daemonset_labels gauge
	   kube_daemonset_labels{namespace="kube-system",daemonset="kube-proxy",label_k8s_app="kube-proxy"} 1
	   kube_daemonset_labels{namespace="kube-system",daemonset="kindnet",label_tier="node",label_app="kindnet",label_k8s_app="kindnet"} 1
	   kube_daemonset_labels{namespace="kube-system",daemonset="ip-masq-agent",label_app="ip-masq-agent",label_k8s_app="ip-masq-agent",label_tier="node"} 1
	   # HELP kube_daemonset_annotations Kubernetes annotations converted to Prometheus labels.
	   # TYPE kube_daemonset_annotations gauge
	   kube_daemonset_annotations{namespace="kube-system",daemonset="ip-masq-agent",annotation_deprecated_daemonset_template_generation="1"} 1
	   kube_daemonset_annotations{namespace="kube-system",daemonset="kube-proxy",annotation_deprecated_daemonset_template_generation="1"} 1
	   kube_daemonset_annotations{namespace="kube-system",daemonset="kindnet",annotation_deprecated_daemonset_template_generation="1"} 1
	*/

	return nil
}

func deploymentMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_deployment_created Unix creation timestamp
		# TYPE kube_deployment_created gauge
		kube_deployment_created{namespace="kube-system",deployment="coredns"} 1.563532771e+09
		kube_deployment_created{namespace="kube-system",deployment="kube-state-metrics"} 1.563532907e+09
		# HELP kube_deployment_status_replicas The number of replicas per deployment.
		# TYPE kube_deployment_status_replicas gauge
		kube_deployment_status_replicas{namespace="kube-system",deployment="coredns"} 2
		kube_deployment_status_replicas{namespace="kube-system",deployment="kube-state-metrics"} 1
		# HELP kube_deployment_status_replicas_available The number of available replicas per deployment.
		# TYPE kube_deployment_status_replicas_available gauge
		kube_deployment_status_replicas_available{namespace="kube-system",deployment="coredns"} 2
		kube_deployment_status_replicas_available{namespace="kube-system",deployment="kube-state-metrics"} 1
		# HELP kube_deployment_status_replicas_unavailable The number of unavailable replicas per deployment.
		# TYPE kube_deployment_status_replicas_unavailable gauge
		kube_deployment_status_replicas_unavailable{namespace="kube-system",deployment="coredns"} 0
		kube_deployment_status_replicas_unavailable{namespace="kube-system",deployment="kube-state-metrics"} 0
		# HELP kube_deployment_status_replicas_updated The number of updated replicas per deployment.
		# TYPE kube_deployment_status_replicas_updated gauge
		kube_deployment_status_replicas_updated{namespace="kube-system",deployment="coredns"} 2
		kube_deployment_status_replicas_updated{namespace="kube-system",deployment="kube-state-metrics"} 1
		# HELP kube_deployment_status_observed_generation The generation observed by the deployment controller.
		# TYPE kube_deployment_status_observed_generation gauge
		kube_deployment_status_observed_generation{namespace="kube-system",deployment="coredns"} 1
		kube_deployment_status_observed_generation{namespace="kube-system",deployment="kube-state-metrics"} 1
		# HELP kube_deployment_spec_replicas Number of desired pods for a deployment.
		# TYPE kube_deployment_spec_replicas gauge
		kube_deployment_spec_replicas{namespace="kube-system",deployment="kube-state-metrics"} 1
		kube_deployment_spec_replicas{namespace="kube-system",deployment="coredns"} 2
		# HELP kube_deployment_spec_paused Whether the deployment is paused and will not be processed by the deployment controller.
		# TYPE kube_deployment_spec_paused gauge
		kube_deployment_spec_paused{namespace="kube-system",deployment="coredns"} 0
		kube_deployment_spec_paused{namespace="kube-system",deployment="kube-state-metrics"} 0
		# HELP kube_deployment_spec_strategy_rollingupdate_max_unavailable Maximum number of unavailable replicas during a rolling update of a deployment.
		# TYPE kube_deployment_spec_strategy_rollingupdate_max_unavailable gauge
		kube_deployment_spec_strategy_rollingupdate_max_unavailable{namespace="kube-system",deployment="coredns"} 1
		kube_deployment_spec_strategy_rollingupdate_max_unavailable{namespace="kube-system",deployment="kube-state-metrics"} 1
		# HELP kube_deployment_spec_strategy_rollingupdate_max_surge Maximum number of replicas that can be scheduled above the desired number of replicas during a rolling update of a deployment.
		# TYPE kube_deployment_spec_strategy_rollingupdate_max_surge gauge
		kube_deployment_spec_strategy_rollingupdate_max_surge{namespace="kube-system",deployment="coredns"} 1
		kube_deployment_spec_strategy_rollingupdate_max_surge{namespace="kube-system",deployment="kube-state-metrics"} 1
		# HELP kube_deployment_metadata_generation Sequence number representing a specific generation of the desired state.
		# TYPE kube_deployment_metadata_generation gauge
		kube_deployment_metadata_generation{namespace="kube-system",deployment="coredns"} 1
		kube_deployment_metadata_generation{namespace="kube-system",deployment="kube-state-metrics"} 1
		# HELP kube_deployment_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_deployment_labels gauge
		kube_deployment_labels{namespace="kube-system",deployment="coredns",label_k8s_app="kube-dns"} 1
		kube_deployment_labels{namespace="kube-system",deployment="kube-state-metrics",label_k8s_app="kube-state-metrics"} 1
		# HELP kube_deployment_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_deployment_annotations gauge
		kube_deployment_annotations{namespace="kube-system",deployment="coredns",annotation_deployment_kubernetes_io_revision="1"} 1
		kube_deployment_annotations{namespace="kube-system",deployment="kube-state-metrics",annotation_deployment_kubernetes_io_revision="1"} 1
	*/
	return nil
}

func enpointMappings() map[string]prometheus.MetricMap {

	/*
		# HELP kube_endpoint_info Information about endpoint.
		# TYPE kube_endpoint_info gauge
		kube_endpoint_info{namespace="kube-system",endpoint="kube-state-metrics"} 1
		kube_endpoint_info{namespace="kube-system",endpoint="kube-scheduler"} 1
		kube_endpoint_info{namespace="kube-system",endpoint="kube-controller-manager"} 1
		kube_endpoint_info{namespace="default",endpoint="kubernetes"} 1
		kube_endpoint_info{namespace="kube-system",endpoint="kube-dns"} 1
		# HELP kube_endpoint_created Unix creation timestamp
		# TYPE kube_endpoint_created gauge
		kube_endpoint_created{namespace="kube-system",endpoint="kube-dns"} 1.563532786e+09
		kube_endpoint_created{namespace="kube-system",endpoint="kube-state-metrics"} 1.563532907e+09
		kube_endpoint_created{namespace="kube-system",endpoint="kube-scheduler"} 1.563532769e+09
		kube_endpoint_created{namespace="kube-system",endpoint="kube-controller-manager"} 1.563532769e+09
		kube_endpoint_created{namespace="default",endpoint="kubernetes"} 1.563532769e+09
		# HELP kube_endpoint_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_endpoint_labels gauge
		kube_endpoint_labels{namespace="kube-system",endpoint="kube-state-metrics",label_k8s_app="kube-state-metrics"} 1
		kube_endpoint_labels{namespace="kube-system",endpoint="kube-scheduler"} 1
		kube_endpoint_labels{namespace="kube-system",endpoint="kube-controller-manager"} 1
		kube_endpoint_labels{namespace="default",endpoint="kubernetes"} 1
		kube_endpoint_labels{namespace="kube-system",endpoint="kube-dns",label_k8s_app="kube-dns",label_kubernetes_io_cluster_service="true",label_kubernetes_io_name="KubeDNS"} 1
		# HELP kube_endpoint_address_available Number of addresses available in endpoint.
		# TYPE kube_endpoint_address_available gauge
		kube_endpoint_address_available{namespace="kube-system",endpoint="kube-state-metrics"} 2
		kube_endpoint_address_available{namespace="kube-system",endpoint="kube-scheduler"} 0
		kube_endpoint_address_available{namespace="kube-system",endpoint="kube-controller-manager"} 0
		kube_endpoint_address_available{namespace="default",endpoint="kubernetes"} 1
		kube_endpoint_address_available{namespace="kube-system",endpoint="kube-dns"} 6
		# HELP kube_endpoint_address_not_ready Number of addresses not ready in endpoint
		# TYPE kube_endpoint_address_not_ready gauge
		kube_endpoint_address_not_ready{namespace="kube-system",endpoint="kube-state-metrics"} 0
		kube_endpoint_address_not_ready{namespace="kube-system",endpoint="kube-scheduler"} 0
		kube_endpoint_address_not_ready{namespace="kube-system",endpoint="kube-controller-manager"} 0
		kube_endpoint_address_not_ready{namespace="default",endpoint="kubernetes"} 0
		kube_endpoint_address_not_ready{namespace="kube-system",endpoint="kube-dns"} 0
		# HELP kube_endpoint_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_endpoint_annotations gauge
		kube_endpoint_annotations{namespace="kube-system",endpoint="kube-state-metrics",annotation_endpoints_kubernetes_io_last_change_trigger_time="2019-07-19T10:42:05Z"} 1
		kube_endpoint_annotations{namespace="kube-system",endpoint="kube-scheduler",annotation_control_plane_alpha_kubernetes_io_leader="{\"holderIdentity\":\"kind-control-plane_79101009-aa11-11e9-b654-0242ac110002\",\"leaseDurationSeconds\":15,\"acquireTime\":\"2019-07-19T10:39:29Z\",\"renewTime\":\"2019-07-19T10:43:46Z\",\"leaderTransitions\":0}"} 1
		kube_endpoint_annotations{namespace="kube-system",endpoint="kube-controller-manager",annotation_control_plane_alpha_kubernetes_io_leader="{\"holderIdentity\":\"kind-control-plane_790c5a19-aa11-11e9-af27-0242ac110002\",\"leaseDurationSeconds\":15,\"acquireTime\":\"2019-07-19T10:39:29Z\",\"renewTime\":\"2019-07-19T10:43:46Z\",\"leaderTransitions\":0}"} 1
		kube_endpoint_annotations{namespace="default",endpoint="kubernetes"} 1
		kube_endpoint_annotations{namespace="kube-system",endpoint="kube-dns",annotation_endpoints_kubernetes_io_last_change_trigger_time="2019-07-19T10:40:35Z"} 1
	*/
	return nil
}

func hpaMappings() map[string]prometheus.MetricMap {

	/*
	   # HELP kube_hpa_metadata_generation The generation observed by the HorizontalPodAutoscaler controller.
	   # TYPE kube_hpa_metadata_generation gauge
	   # HELP kube_hpa_spec_max_replicas Upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
	   # TYPE kube_hpa_spec_max_replicas gauge
	   # HELP kube_hpa_spec_min_replicas Lower limit for the number of pods that can be set by the autoscaler, default 1.
	   # TYPE kube_hpa_spec_min_replicas gauge
	   # HELP kube_hpa_status_current_replicas Current number of replicas of pods managed by this autoscaler.
	   # TYPE kube_hpa_status_current_replicas gauge
	   # HELP kube_hpa_status_desired_replicas Desired number of replicas of pods managed by this autoscaler.
	   # TYPE kube_hpa_status_desired_replicas gauge
	   # HELP kube_hpa_labels Kubernetes labels converted to Prometheus labels.
	   # TYPE kube_hpa_labels gauge
	   # HELP kube_hpa_status_condition The condition of this autoscaler.
	   # TYPE kube_hpa_status_condition gauge
	   # HELP kube_hpa_annotations Kubernetes annotations converted to Prometheus labels.
	   # TYPE kube_hpa_annotations gauge
	*/

	return nil
}

func ingressMappings() map[string]prometheus.MetricMap {

	/*
		# HELP kube_ingress_info Information about ingress.
		# TYPE kube_ingress_info gauge
		# HELP kube_ingress_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_ingress_labels gauge
		# HELP kube_ingress_created Unix creation timestamp
		# TYPE kube_ingress_created gauge
		# HELP kube_ingress_metadata_resource_version Resource version representing a specific version of ingress.
		# TYPE kube_ingress_metadata_resource_version gauge
		# HELP kube_ingress_path Ingress host, paths and backend service information.
		# TYPE kube_ingress_path gauge
		# HELP kube_ingress_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_ingress_annotations gauge
	*/

	return nil
}

func jobMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_job_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_job_labels gauge
		# HELP kube_job_info Information about job.
		# TYPE kube_job_info gauge
		# HELP kube_job_created Unix creation timestamp
		# TYPE kube_job_created gauge
		# HELP kube_job_spec_parallelism The maximum desired number of pods the job should run at any given time.
		# TYPE kube_job_spec_parallelism gauge
		# HELP kube_job_spec_completions The desired number of successfully finished pods the job should be run with.
		# TYPE kube_job_spec_completions gauge
		# HELP kube_job_spec_active_deadline_seconds The duration in seconds relative to the startTime that the job may be active before the system tries to terminate it.
		# TYPE kube_job_spec_active_deadline_seconds gauge
		# HELP kube_job_status_succeeded The number of pods which reached Phase Succeeded.
		# TYPE kube_job_status_succeeded gauge
		# HELP kube_job_status_failed The number of pods which reached Phase Failed.
		# TYPE kube_job_status_failed gauge
		# HELP kube_job_status_active The number of actively running pods.
		# TYPE kube_job_status_active gauge
		# HELP kube_job_complete The job has completed its execution.
		# TYPE kube_job_complete gauge
		# HELP kube_job_failed The job has failed its execution.
		# TYPE kube_job_failed gauge
		# HELP kube_job_status_start_time StartTime represents time when the job was acknowledged by the Job Manager.
		# TYPE kube_job_status_start_time gauge
		# HELP kube_job_status_completion_time CompletionTime represents time when the job was completed.
		# TYPE kube_job_status_completion_time gauge
		# HELP kube_job_owner Information about the Job's owner.
		# TYPE kube_job_owner gauge
		# HELP kube_job_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_job_annotations gauge
	*/

	return nil
}

func limitRangeMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_limitrange Information about limit range.
		# TYPE kube_limitrange gauge
		# HELP kube_limitrange_created Unix creation timestamp
		# TYPE kube_limitrange_created gauge
		# HELP kube_limitrange_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_limitrange_annotations gauge
	*/

	return nil
}

func namespaceMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_namespace_created Unix creation timestamp
		# TYPE kube_namespace_created gauge
		kube_namespace_created{namespace="kube-system"} 1.563532766e+09
		kube_namespace_created{namespace="kube-public"} 1.563532766e+09
		kube_namespace_created{namespace="kube-node-lease"} 1.563532766e+09
		kube_namespace_created{namespace="default"} 1.563532769e+09
		# HELP kube_namespace_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_namespace_labels gauge
		kube_namespace_labels{namespace="default"} 1
		kube_namespace_labels{namespace="kube-system"} 1
		kube_namespace_labels{namespace="kube-public"} 1
		kube_namespace_labels{namespace="kube-node-lease"} 1
		# HELP kube_namespace_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_namespace_annotations gauge
		kube_namespace_annotations{namespace="kube-public"} 1
		kube_namespace_annotations{namespace="kube-node-lease"} 1
		kube_namespace_annotations{namespace="default"} 1
		kube_namespace_annotations{namespace="kube-system"} 1
		# HELP kube_namespace_status_phase kubernetes namespace status phase.
		# TYPE kube_namespace_status_phase gauge
		kube_namespace_status_phase{namespace="kube-node-lease",phase="Active"} 1
		kube_namespace_status_phase{namespace="kube-node-lease",phase="Terminating"} 0
		kube_namespace_status_phase{namespace="default",phase="Active"} 1
		kube_namespace_status_phase{namespace="default",phase="Terminating"} 0
		kube_namespace_status_phase{namespace="kube-system",phase="Active"} 1
		kube_namespace_status_phase{namespace="kube-system",phase="Terminating"} 0
		kube_namespace_status_phase{namespace="kube-public",phase="Active"} 1
		kube_namespace_status_phase{namespace="kube-public",phase="Terminating"} 0
	*/

	return nil
}

func nodeMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_node_info Information about a cluster node.
		# TYPE kube_node_info gauge
		kube_node_info{node="kind-control-plane",kernel_version="4.15.0-54-generic",os_image="Ubuntu Disco Dingo (development branch)",container_runtime_version="containerd://1.2.6-0ubuntu1",kubelet_version="v1.14.2",kubeproxy_version="v1.14.2",provider_id=""} 1
		# HELP kube_node_created Unix creation timestamp
		# TYPE kube_node_created gauge
		kube_node_created{node="kind-control-plane"} 1.563532766e+09
		# HELP kube_node_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_node_labels gauge
		kube_node_labels{node="kind-control-plane",label_kubernetes_io_hostname="kind-control-plane",label_kubernetes_io_os="linux",label_node_role_kubernetes_io_master="",label_beta_kubernetes_io_arch="amd64",label_beta_kubernetes_io_os="linux",label_kubernetes_io_arch="amd64"} 1
		# HELP kube_node_spec_unschedulable Whether a node can schedule new pods.
		# TYPE kube_node_spec_unschedulable gauge
		kube_node_spec_unschedulable{node="kind-control-plane"} 0
		# HELP kube_node_spec_taint The taint of a cluster node.
		# TYPE kube_node_spec_taint gauge
		# HELP kube_node_status_condition The condition of a cluster node.
		# TYPE kube_node_status_condition gauge
		kube_node_status_condition{node="kind-control-plane",condition="MemoryPressure",status="true"} 0
		kube_node_status_condition{node="kind-control-plane",condition="MemoryPressure",status="false"} 1
		kube_node_status_condition{node="kind-control-plane",condition="MemoryPressure",status="unknown"} 0
		kube_node_status_condition{node="kind-control-plane",condition="DiskPressure",status="true"} 0
		kube_node_status_condition{node="kind-control-plane",condition="DiskPressure",status="false"} 1
		kube_node_status_condition{node="kind-control-plane",condition="DiskPressure",status="unknown"} 0
		kube_node_status_condition{node="kind-control-plane",condition="PIDPressure",status="true"} 0
		kube_node_status_condition{node="kind-control-plane",condition="PIDPressure",status="false"} 1
		kube_node_status_condition{node="kind-control-plane",condition="PIDPressure",status="unknown"} 0
		kube_node_status_condition{node="kind-control-plane",condition="Ready",status="true"} 1
		kube_node_status_condition{node="kind-control-plane",condition="Ready",status="false"} 0
		kube_node_status_condition{node="kind-control-plane",condition="Ready",status="unknown"} 0
		# HELP kube_node_status_phase The phase the node is currently in.
		# TYPE kube_node_status_phase gauge
		# HELP kube_node_status_capacity The capacity for different resources of a node.
		# TYPE kube_node_status_capacity gauge
		kube_node_status_capacity{node="kind-control-plane",resource="ephemeral_storage",unit="byte"} 1.006530654208e+12
		kube_node_status_capacity{node="kind-control-plane",resource="hugepages_1Gi",unit="byte"} 0
		kube_node_status_capacity{node="kind-control-plane",resource="hugepages_2Mi",unit="byte"} 0
		kube_node_status_capacity{node="kind-control-plane",resource="memory",unit="byte"} 3.3571418112e+10
		kube_node_status_capacity{node="kind-control-plane",resource="pods",unit="integer"} 110
		kube_node_status_capacity{node="kind-control-plane",resource="cpu",unit="core"} 8
		# HELP kube_node_status_capacity_pods The total pod resources of the node.
		# TYPE kube_node_status_capacity_pods gauge
		kube_node_status_capacity_pods{node="kind-control-plane"} 110
		# HELP kube_node_status_capacity_cpu_cores The total CPU resources of the node.
		# TYPE kube_node_status_capacity_cpu_cores gauge
		kube_node_status_capacity_cpu_cores{node="kind-control-plane"} 8
		# HELP kube_node_status_capacity_memory_bytes The total memory resources of the node.
		# TYPE kube_node_status_capacity_memory_bytes gauge
		kube_node_status_capacity_memory_bytes{node="kind-control-plane"} 3.3571418112e+10
		# HELP kube_node_status_allocatable The allocatable for different resources of a node that are available for scheduling.
		# TYPE kube_node_status_allocatable gauge
		kube_node_status_allocatable{node="kind-control-plane",resource="ephemeral_storage",unit="byte"} 1.006530654208e+12
		kube_node_status_allocatable{node="kind-control-plane",resource="hugepages_1Gi",unit="byte"} 0
		kube_node_status_allocatable{node="kind-control-plane",resource="hugepages_2Mi",unit="byte"} 0
		kube_node_status_allocatable{node="kind-control-plane",resource="memory",unit="byte"} 3.3571418112e+10
		kube_node_status_allocatable{node="kind-control-plane",resource="pods",unit="integer"} 110
		kube_node_status_allocatable{node="kind-control-plane",resource="cpu",unit="core"} 8
		# HELP kube_node_status_allocatable_pods The pod resources of a node that are available for scheduling.
		# TYPE kube_node_status_allocatable_pods gauge
		kube_node_status_allocatable_pods{node="kind-control-plane"} 110
		# HELP kube_node_status_allocatable_cpu_cores The CPU resources of a node that are available for scheduling.
		# TYPE kube_node_status_allocatable_cpu_cores gauge
		kube_node_status_allocatable_cpu_cores{node="kind-control-plane"} 8
		# HELP kube_node_status_allocatable_memory_bytes The memory resources of a node that are available for scheduling.
		# TYPE kube_node_status_allocatable_memory_bytes gauge
		kube_node_status_allocatable_memory_bytes{node="kind-control-plane"} 3.3571418112e+10
		# HELP kube_node_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_node_annotations gauge
		kube_node_annotations{node="kind-control-plane",annotation_kubeadm_alpha_kubernetes_io_cri_socket="/run/containerd/containerd.sock",annotation_node_alpha_kubernetes_io_ttl="0",annotation_volumes_kubernetes_io_controller_managed_attach_detach="true"} 1
	*/

	return nil
}

func persistentVolumeMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_persistentvolumeclaim_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_persistentvolumeclaim_labels gauge
		# HELP kube_persistentvolumeclaim_info Information about persistent volume claim.
		# TYPE kube_persistentvolumeclaim_info gauge
		# HELP kube_persistentvolumeclaim_status_phase The phase the persistent volume claim is currently in.
		# TYPE kube_persistentvolumeclaim_status_phase gauge
		# HELP kube_persistentvolumeclaim_resource_requests_storage_bytes The capacity of storage requested by the persistent volume claim.
		# TYPE kube_persistentvolumeclaim_resource_requests_storage_bytes gauge
		# HELP kube_persistentvolumeclaim_access_mode The access mode(s) specified by the persistent volume claim.
		# TYPE kube_persistentvolumeclaim_access_mode gauge
		# HELP kube_persistentvolumeclaim_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_persistentvolumeclaim_annotations gauge
		# HELP kube_persistentvolume_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_persistentvolume_labels gauge
		# HELP kube_persistentvolume_status_phase The phase indicates if a volume is available, bound to a claim, or released by a claim.
		# TYPE kube_persistentvolume_status_phase gauge
		# HELP kube_persistentvolume_info Information about persistentvolume.
		# TYPE kube_persistentvolume_info gauge
		# HELP kube_persistentvolume_capacity_bytes Persistentvolume capacity in bytes.
		# TYPE kube_persistentvolume_capacity_bytes gauge
		# HELP kube_persistentvolume_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_persistentvolume_annotations gauge
	*/

	return nil
}

func podDisruptionBudgetMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_poddisruptionbudget_created Unix creation timestamp
		# TYPE kube_poddisruptionbudget_created gauge
		# HELP kube_poddisruptionbudget_status_current_healthy Current number of healthy pods
		# TYPE kube_poddisruptionbudget_status_current_healthy gauge
		# HELP kube_poddisruptionbudget_status_desired_healthy Minimum desired number of healthy pods
		# TYPE kube_poddisruptionbudget_status_desired_healthy gauge
		# HELP kube_poddisruptionbudget_status_pod_disruptions_allowed Number of pod disruptions that are currently allowed
		# TYPE kube_poddisruptionbudget_status_pod_disruptions_allowed gauge
		# HELP kube_poddisruptionbudget_status_expected_pods Total number of pods counted by this disruption budget
		# TYPE kube_poddisruptionbudget_status_expected_pods gauge
		# HELP kube_poddisruptionbudget_status_observed_generation Most recent generation observed when updating this PDB status
		# TYPE kube_poddisruptionbudget_status_observed_generation gauge
		# HELP kube_poddisruptionbudget_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_poddisruptionbudget_annotations gauge
	*/

	return nil
}

func podMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_pod_info Information about pod.
		# TYPE kube_pod_info gauge
		kube_pod_info{namespace="kube-system",pod="kube-proxy-w26cb",host_ip="172.17.0.2",pod_ip="172.17.0.2",uid="86c2c109-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="DaemonSet",created_by_name="kube-proxy",priority_class="system-node-critical"} 1
		kube_pod_info{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",host_ip="172.17.0.2",pod_ip="10.244.0.3",uid="86795bd7-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="ReplicaSet",created_by_name="coredns-fb8b8dccf",priority_class="system-cluster-critical"} 1
		kube_pod_info{namespace="kube-system",pod="ip-masq-agent-lhrqv",host_ip="172.17.0.2",pod_ip="172.17.0.2",uid="86c86f8c-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="DaemonSet",created_by_name="ip-masq-agent",priority_class=""} 1
		kube_pod_info{namespace="kube-system",pod="ubuntu",host_ip="172.17.0.2",pod_ip="172.17.0.2",uid="e0ec86b6-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="<none>",created_by_name="<none>",priority_class=""} 1
		kube_pod_info{namespace="kube-system",pod="etcd-kind-control-plane",host_ip="172.17.0.2",pod_ip="172.17.0.2",uid="acf7cd75-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="<none>",created_by_name="<none>",priority_class="system-cluster-critical"} 1
		kube_pod_info{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",host_ip="172.17.0.2",pod_ip="10.244.0.4",uid="cf8e98e4-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="ReplicaSet",created_by_name="kube-state-metrics-6696d9cc47",priority_class=""} 1
		kube_pod_info{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",host_ip="172.17.0.2",pod_ip="10.244.0.2",uid="867bd2cf-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="ReplicaSet",created_by_name="coredns-fb8b8dccf",priority_class="system-cluster-critical"} 1
		kube_pod_info{namespace="kube-system",pod="kube-scheduler-kind-control-plane",host_ip="172.17.0.2",pod_ip="172.17.0.2",uid="abc63ff1-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="<none>",created_by_name="<none>",priority_class="system-cluster-critical"} 1
		kube_pod_info{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",host_ip="172.17.0.2",pod_ip="172.17.0.2",uid="ac5ed02c-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="<none>",created_by_name="<none>",priority_class="system-cluster-critical"} 1
		kube_pod_info{namespace="kube-system",pod="kindnet-wzfk8",host_ip="172.17.0.2",pod_ip="172.17.0.2",uid="86c2d0da-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="DaemonSet",created_by_name="kindnet",priority_class=""} 1
		kube_pod_info{namespace="kube-system",pod="kube-apiserver-kind-control-plane",host_ip="172.17.0.2",pod_ip="172.17.0.2",uid="a2d706cd-aa11-11e9-9d2e-0242ac110002",node="kind-control-plane",created_by_kind="<none>",created_by_name="<none>",priority_class="system-cluster-critical"} 1
		# HELP kube_pod_start_time Start time in unix timestamp for a pod.
		# TYPE kube_pod_start_time gauge
		kube_pod_start_time{namespace="kube-system",pod="kube-controller-manager-kind-control-plane"} 1.56353276e+09
		kube_pod_start_time{namespace="kube-system",pod="kindnet-wzfk8"} 1.563532786e+09
		kube_pod_start_time{namespace="kube-system",pod="kube-apiserver-kind-control-plane"} 1.56353276e+09
		kube_pod_start_time{namespace="kube-system",pod="kube-scheduler-kind-control-plane"} 1.56353276e+09
		kube_pod_start_time{namespace="kube-system",pod="etcd-kind-control-plane"} 1.56353276e+09
		kube_pod_start_time{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc"} 1.563532908e+09
		kube_pod_start_time{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45"} 1.563532828e+09
		kube_pod_start_time{namespace="kube-system",pod="kube-proxy-w26cb"} 1.563532786e+09
		kube_pod_start_time{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7"} 1.563532828e+09
		kube_pod_start_time{namespace="kube-system",pod="ip-masq-agent-lhrqv"} 1.563532786e+09
		kube_pod_start_time{namespace="kube-system",pod="ubuntu"} 1.563532937e+09
		# HELP kube_pod_completion_time Completion time in unix timestamp for a pod.
		# TYPE kube_pod_completion_time gauge
		# HELP kube_pod_owner Information about the Pod's owner.
		# TYPE kube_pod_owner gauge
		kube_pod_owner{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",owner_kind="<none>",owner_name="<none>",owner_is_controller="<none>"} 1
		kube_pod_owner{namespace="kube-system",pod="kindnet-wzfk8",owner_kind="DaemonSet",owner_name="kindnet",owner_is_controller="true"} 1
		kube_pod_owner{namespace="kube-system",pod="kube-apiserver-kind-control-plane",owner_kind="<none>",owner_name="<none>",owner_is_controller="<none>"} 1
		kube_pod_owner{namespace="kube-system",pod="kube-scheduler-kind-control-plane",owner_kind="<none>",owner_name="<none>",owner_is_controller="<none>"} 1
		kube_pod_owner{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",owner_kind="ReplicaSet",owner_name="coredns-fb8b8dccf",owner_is_controller="true"} 1
		kube_pod_owner{namespace="kube-system",pod="ip-masq-agent-lhrqv",owner_kind="DaemonSet",owner_name="ip-masq-agent",owner_is_controller="true"} 1
		kube_pod_owner{namespace="kube-system",pod="ubuntu",owner_kind="<none>",owner_name="<none>",owner_is_controller="<none>"} 1
		kube_pod_owner{namespace="kube-system",pod="etcd-kind-control-plane",owner_kind="<none>",owner_name="<none>",owner_is_controller="<none>"} 1
		kube_pod_owner{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",owner_kind="ReplicaSet",owner_name="kube-state-metrics-6696d9cc47",owner_is_controller="true"} 1
		kube_pod_owner{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",owner_kind="ReplicaSet",owner_name="coredns-fb8b8dccf",owner_is_controller="true"} 1
		kube_pod_owner{namespace="kube-system",pod="kube-proxy-w26cb",owner_kind="DaemonSet",owner_name="kube-proxy",owner_is_controller="true"} 1
		# HELP kube_pod_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_pod_labels gauge
		kube_pod_labels{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",label_pod_template_hash="6696d9cc47",label_k8s_app="kube-state-metrics"} 1
		kube_pod_labels{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",label_k8s_app="kube-dns",label_pod_template_hash="fb8b8dccf"} 1
		kube_pod_labels{namespace="kube-system",pod="kube-proxy-w26cb",label_controller_revision_hash="6c5f45759d",label_k8s_app="kube-proxy",label_pod_template_generation="1"} 1
		kube_pod_labels{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",label_k8s_app="kube-dns",label_pod_template_hash="fb8b8dccf"} 1
		kube_pod_labels{namespace="kube-system",pod="ip-masq-agent-lhrqv",label_k8s_app="ip-masq-agent",label_pod_template_generation="1",label_tier="node",label_app="ip-masq-agent",label_controller_revision_hash="699c8f6fcc"} 1
		kube_pod_labels{namespace="kube-system",pod="ubuntu"} 1
		kube_pod_labels{namespace="kube-system",pod="etcd-kind-control-plane",label_component="etcd",label_tier="control-plane"} 1
		kube_pod_labels{namespace="kube-system",pod="kindnet-wzfk8",label_app="kindnet",label_controller_revision_hash="6744c84f84",label_k8s_app="kindnet",label_pod_template_generation="1",label_tier="node"} 1
		kube_pod_labels{namespace="kube-system",pod="kube-apiserver-kind-control-plane",label_component="kube-apiserver",label_tier="control-plane"} 1
		kube_pod_labels{namespace="kube-system",pod="kube-scheduler-kind-control-plane",label_component="kube-scheduler",label_tier="control-plane"} 1
		kube_pod_labels{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",label_component="kube-controller-manager",label_tier="control-plane"} 1
		# HELP kube_pod_created Unix creation timestamp
		# TYPE kube_pod_created gauge
		kube_pod_created{namespace="kube-system",pod="kube-controller-manager-kind-control-plane"} 1.563532849e+09
		kube_pod_created{namespace="kube-system",pod="kindnet-wzfk8"} 1.563532786e+09
		kube_pod_created{namespace="kube-system",pod="kube-apiserver-kind-control-plane"} 1.563532833e+09
		kube_pod_created{namespace="kube-system",pod="kube-scheduler-kind-control-plane"} 1.563532848e+09
		kube_pod_created{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7"} 1.563532785e+09
		kube_pod_created{namespace="kube-system",pod="ip-masq-agent-lhrqv"} 1.563532786e+09
		kube_pod_created{namespace="kube-system",pod="ubuntu"} 1.563532937e+09
		kube_pod_created{namespace="kube-system",pod="etcd-kind-control-plane"} 1.56353285e+09
		kube_pod_created{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc"} 1.563532908e+09
		kube_pod_created{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45"} 1.563532785e+09
		kube_pod_created{namespace="kube-system",pod="kube-proxy-w26cb"} 1.563532786e+09
		# HELP kube_pod_status_scheduled_time Unix timestamp when pod moved into scheduled status
		# TYPE kube_pod_status_scheduled_time gauge
		kube_pod_status_scheduled_time{namespace="kube-system",pod="ubuntu"} 1.563532937e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="etcd-kind-control-plane"} 1.56353276e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc"} 1.563532908e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45"} 1.563532828e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="kube-proxy-w26cb"} 1.563532786e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7"} 1.563532828e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="ip-masq-agent-lhrqv"} 1.563532786e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="kube-controller-manager-kind-control-plane"} 1.56353276e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="kindnet-wzfk8"} 1.563532786e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="kube-apiserver-kind-control-plane"} 1.56353276e+09
		kube_pod_status_scheduled_time{namespace="kube-system",pod="kube-scheduler-kind-control-plane"} 1.56353276e+09
		# HELP kube_pod_status_phase The pods current phase.
		# TYPE kube_pod_status_phase gauge
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-proxy-w26cb",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-proxy-w26cb",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-proxy-w26cb",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-proxy-w26cb",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="kube-proxy-w26cb",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="ip-masq-agent-lhrqv",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="ip-masq-agent-lhrqv",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="ip-masq-agent-lhrqv",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="ip-masq-agent-lhrqv",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="ip-masq-agent-lhrqv",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="ubuntu",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="ubuntu",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="ubuntu",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="ubuntu",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="ubuntu",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="etcd-kind-control-plane",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="etcd-kind-control-plane",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="etcd-kind-control-plane",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="etcd-kind-control-plane",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="etcd-kind-control-plane",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-apiserver-kind-control-plane",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-apiserver-kind-control-plane",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-apiserver-kind-control-plane",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-apiserver-kind-control-plane",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="kube-apiserver-kind-control-plane",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-scheduler-kind-control-plane",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-scheduler-kind-control-plane",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-scheduler-kind-control-plane",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-scheduler-kind-control-plane",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="kube-scheduler-kind-control-plane",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",phase="Unknown"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kindnet-wzfk8",phase="Pending"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kindnet-wzfk8",phase="Succeeded"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kindnet-wzfk8",phase="Failed"} 0
		kube_pod_status_phase{namespace="kube-system",pod="kindnet-wzfk8",phase="Running"} 1
		kube_pod_status_phase{namespace="kube-system",pod="kindnet-wzfk8",phase="Unknown"} 0
		# HELP kube_pod_status_ready Describes whether the pod is ready to serve requests.
		# TYPE kube_pod_status_ready gauge
		kube_pod_status_ready{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kindnet-wzfk8",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="kindnet-wzfk8",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kindnet-wzfk8",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kube-apiserver-kind-control-plane",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="kube-apiserver-kind-control-plane",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kube-apiserver-kind-control-plane",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kube-scheduler-kind-control-plane",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="kube-scheduler-kind-control-plane",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kube-scheduler-kind-control-plane",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="ip-masq-agent-lhrqv",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="ip-masq-agent-lhrqv",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="ip-masq-agent-lhrqv",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="ubuntu",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="ubuntu",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="ubuntu",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="etcd-kind-control-plane",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="etcd-kind-control-plane",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="etcd-kind-control-plane",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kube-proxy-w26cb",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="kube-proxy-w26cb",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="kube-proxy-w26cb",condition="unknown"} 0
		kube_pod_status_ready{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",condition="true"} 1
		kube_pod_status_ready{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",condition="false"} 0
		kube_pod_status_ready{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",condition="unknown"} 0
		# HELP kube_pod_status_scheduled Describes the status of the scheduling process for the pod.
		# TYPE kube_pod_status_scheduled gauge
		kube_pod_status_scheduled{namespace="kube-system",pod="kindnet-wzfk8",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="kindnet-wzfk8",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kindnet-wzfk8",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-apiserver-kind-control-plane",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-apiserver-kind-control-plane",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-apiserver-kind-control-plane",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-scheduler-kind-control-plane",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-scheduler-kind-control-plane",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-scheduler-kind-control-plane",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-proxy-w26cb",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-proxy-w26cb",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="kube-proxy-w26cb",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="ip-masq-agent-lhrqv",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="ip-masq-agent-lhrqv",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="ip-masq-agent-lhrqv",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="ubuntu",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="ubuntu",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="ubuntu",condition="unknown"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="etcd-kind-control-plane",condition="true"} 1
		kube_pod_status_scheduled{namespace="kube-system",pod="etcd-kind-control-plane",condition="false"} 0
		kube_pod_status_scheduled{namespace="kube-system",pod="etcd-kind-control-plane",condition="unknown"} 0
		# HELP kube_pod_container_info Information about a container in a pod.
		# TYPE kube_pod_container_info gauge
		kube_pod_container_info{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",image="k8s.gcr.io/kube-scheduler:v1.14.2",image_id="sha256:1c93cc1335f8df0a96db1a773bb2851920fb574e1c9386f3960674279d5b978b",container_id="containerd://5c166408211121af0ee6653e3f8c580db54a10638ce21bd76c8d32f1834ac70a"} 1
		kube_pod_container_info{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",image="k8s.gcr.io/kube-controller-manager:v1.14.2",image_id="sha256:58f6abb9fb1b336348d3bb9dd80b5ecbc8dc963a3c1c20e778a0c20d3ed25344",container_id="containerd://34bb96f6ed95ec986c867d137b7c78abd9b5ffef9213de9bb9f2fc3f7931dd9f"} 1
		kube_pod_container_info{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",image="docker.io/kindest/kindnetd:0.1.0",image_id="sha256:f227066bdc5f9aa2f8a9bb54854e5b7a23c6db8fce0f927e5c4feef8a9e74d46",container_id="containerd://5ab56f5b198a2451f9c54fec42df486bb898ad61fbc1d40822da95de62406d09"} 1
		kube_pod_container_info{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",image="k8s.gcr.io/kube-apiserver:v1.14.2",image_id="sha256:e455634c173b0060e537f229155cbb1649d96945d8de54f3321eebd092d66a0c",container_id="containerd://0651d875670471c90b8fc4287432312f7a926f04e1abf30493d3a77876b7459c"} 1
		kube_pod_container_info{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",image="k8s.gcr.io/kube-proxy:v1.14.2",image_id="sha256:7387c4b88e2df50ccca4f6f8167992605cfe50d0075a647b5ab5187378ac2bd8",container_id="containerd://ffac605555495c462d424755b5ad36925795b5bf305b90324c15df2673fc5b65"} 1
		kube_pod_container_info{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",image="k8s.gcr.io/coredns:1.3.1",image_id="sha256:eb516548c180f8a6e0235034ccee2428027896af16a509786da13022fe95fe8c",container_id="containerd://278d72088ccafc9b6b9368f48177bf1cb80d9e94022f96eab94c5dd80189404b"} 1
		kube_pod_container_info{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",image="k8s.gcr.io/ip-masq-agent:v2.4.1",image_id="sha256:19bb968f77bba3a5b5f56b5c033d71f699c22bdc8bbe9412f0bfaf7f674a64cc",container_id="containerd://e814ba3d3cad671c04eb17e4f01a148a9c91a7647425a8c1cbc9e52bda628675"} 1
		kube_pod_container_info{namespace="kube-system",pod="ubuntu",container="ubuntu",image="docker.io/library/ubuntu:latest",image_id="docker.io/library/ubuntu@sha256:9b1702dcfe32c873a770a32cfd306dd7fc1c4fd134adfb783db68defc8894b3c",container_id="containerd://eab15d981817030be2c2da3e2d3408b922b118052e8ea6c0a1879a76c53ad4a8"} 1
		kube_pod_container_info{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",image="k8s.gcr.io/etcd:3.3.10",image_id="sha256:2c4adeb21b4ff8ed3309d0e42b6b4ae39872399f7b37e0856e673b13c4aba13d",container_id="containerd://9f95c4f129e898a5315dabbb8135847066c40f2072df786348e88668e14af737"} 1
		kube_pod_container_info{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",image="quay.io/coreos/kube-state-metrics:v1.7.0",image_id="sha256:648d05c2fb5ef5ca003afa041c193bba2d1f291bcd6ae3e3f9a3d1a383de4ebd",container_id="containerd://dbe8e310e8046d1262a945403cc56289463f3ce64f6cb0726e5c0f4ac223f146"} 1
		kube_pod_container_info{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",image="k8s.gcr.io/coredns:1.3.1",image_id="sha256:eb516548c180f8a6e0235034ccee2428027896af16a509786da13022fe95fe8c",container_id="containerd://f21970f637634202d549fb6a26ab559fcb671d629e36da9eea120c85ff13a1bd"} 1
		# HELP kube_pod_init_container_info Information about an init container in a pod.
		# TYPE kube_pod_init_container_info gauge
		# HELP kube_pod_container_status_waiting Describes whether the container is currently in waiting state.
		# TYPE kube_pod_container_status_waiting gauge
		kube_pod_container_status_waiting{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="ubuntu",container="ubuntu"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler"} 0
		kube_pod_container_status_waiting{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager"} 0
		# HELP kube_pod_init_container_status_waiting Describes whether the init container is currently in waiting state.
		# TYPE kube_pod_init_container_status_waiting gauge
		# HELP kube_pod_container_status_waiting_reason Describes the reason the container is currently in waiting state.
		# TYPE kube_pod_container_status_waiting_reason gauge
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="InvalidImageName"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="ContainerCreating"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="CrashLoopBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="CreateContainerConfigError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="ErrImagePull"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="ImagePullBackOff"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="CreateContainerError"} 0
		kube_pod_container_status_waiting_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="InvalidImageName"} 0
		# HELP kube_pod_init_container_status_waiting_reason Describes the reason the init container is currently in waiting state.
		# TYPE kube_pod_init_container_status_waiting_reason gauge
		# HELP kube_pod_container_status_running Describes whether the container is currently in running state.
		# TYPE kube_pod_container_status_running gauge
		kube_pod_container_status_running{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="ubuntu",container="ubuntu"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics"} 1
		kube_pod_container_status_running{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns"} 1
		# HELP kube_pod_init_container_status_running Describes whether the init container is currently in running state.
		# TYPE kube_pod_init_container_status_running gauge
		# HELP kube_pod_container_status_terminated Describes whether the container is currently in terminated state.
		# TYPE kube_pod_container_status_terminated gauge
		kube_pod_container_status_terminated{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="ubuntu",container="ubuntu"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver"} 0
		kube_pod_container_status_terminated{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler"} 0
		# HELP kube_pod_init_container_status_terminated Describes whether the init container is currently in terminated state.
		# TYPE kube_pod_init_container_status_terminated gauge
		# HELP kube_pod_container_status_terminated_reason Describes the reason the container is currently in terminated state.
		# TYPE kube_pod_container_status_terminated_reason gauge
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="DeadlineExceeded"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="OOMKilled"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="Completed"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="Error"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="ContainerCannotRun"} 0
		kube_pod_container_status_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="DeadlineExceeded"} 0
		# HELP kube_pod_init_container_status_terminated_reason Describes the reason the init container is currently in terminated state.
		# TYPE kube_pod_init_container_status_terminated_reason gauge
		# HELP kube_pod_container_status_last_terminated_reason Describes the last reason the container was in terminated state.
		# TYPE kube_pod_container_status_last_terminated_reason gauge
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="Error"} 1
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="ubuntu",container="ubuntu",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy",reason="DeadlineExceeded"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="OOMKilled"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="Completed"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="Error"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="ContainerCannotRun"} 0
		kube_pod_container_status_last_terminated_reason{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",reason="DeadlineExceeded"} 0
		# HELP kube_pod_init_container_status_last_terminated_reason Describes the last reason the init container was in terminated state.
		# TYPE kube_pod_init_container_status_last_terminated_reason gauge
		# HELP kube_pod_container_status_ready Describes whether the containers readiness check succeeded.
		# TYPE kube_pod_container_status_ready gauge
		kube_pod_container_status_ready{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="ubuntu",container="ubuntu"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler"} 1
		kube_pod_container_status_ready{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager"} 1
		# HELP kube_pod_init_container_status_ready Describes whether the init containers readiness check succeeded.
		# TYPE kube_pod_init_container_status_ready gauge
		# HELP kube_pod_container_status_restarts_total The number of container restarts per container.
		# TYPE kube_pod_container_status_restarts_total counter
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns"} 0
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="ip-masq-agent-lhrqv",container="ip-masq-agent"} 0
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="ubuntu",container="ubuntu"} 0
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="etcd-kind-control-plane",container="etcd"} 0
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="kube-state-metrics-6696d9cc47-v4rnc",container="kube-state-metrics"} 0
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns"} 0
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="kube-proxy-w26cb",container="kube-proxy"} 0
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager"} 0
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni"} 1
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver"} 0
		kube_pod_container_status_restarts_total{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler"} 0
		# HELP kube_pod_init_container_status_restarts_total The number of restarts for the init container.
		# TYPE kube_pod_init_container_status_restarts_total counter
		# HELP kube_pod_container_resource_requests The number of requested request resource by a container.
		# TYPE kube_pod_container_resource_requests gauge
		kube_pod_container_resource_requests{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",node="kind-control-plane",resource="cpu",unit="core"} 0.1
		kube_pod_container_resource_requests{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",node="kind-control-plane",resource="memory",unit="byte"} 7.340032e+07
		kube_pod_container_resource_requests{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",node="kind-control-plane",resource="cpu",unit="core"} 0.1
		kube_pod_container_resource_requests{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",node="kind-control-plane",resource="memory",unit="byte"} 7.340032e+07
		kube_pod_container_resource_requests{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",node="kind-control-plane",resource="cpu",unit="core"} 0.2
		kube_pod_container_resource_requests{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",node="kind-control-plane",resource="cpu",unit="core"} 0.1
		kube_pod_container_resource_requests{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",node="kind-control-plane",resource="memory",unit="byte"} 5.24288e+07
		kube_pod_container_resource_requests{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",node="kind-control-plane",resource="cpu",unit="core"} 0.25
		kube_pod_container_resource_requests{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",node="kind-control-plane",resource="cpu",unit="core"} 0.1
		# HELP kube_pod_container_resource_limits The number of requested limit resource by a container.
		# TYPE kube_pod_container_resource_limits gauge
		kube_pod_container_resource_limits{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",node="kind-control-plane",resource="cpu",unit="core"} 0.1
		kube_pod_container_resource_limits{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",node="kind-control-plane",resource="memory",unit="byte"} 5.24288e+07
		kube_pod_container_resource_limits{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",node="kind-control-plane",resource="memory",unit="byte"} 1.7825792e+08
		kube_pod_container_resource_limits{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",node="kind-control-plane",resource="memory",unit="byte"} 1.7825792e+08
		# HELP kube_pod_init_container_resource_limits The number of requested limit resource by the init container.
		# TYPE kube_pod_init_container_resource_limits gauge
		# HELP kube_pod_container_resource_requests_cpu_cores The number of requested cpu cores by a container.
		# TYPE kube_pod_container_resource_requests_cpu_cores gauge
		kube_pod_container_resource_requests_cpu_cores{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",node="kind-control-plane"} 0.1
		kube_pod_container_resource_requests_cpu_cores{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",node="kind-control-plane"} 0.1
		kube_pod_container_resource_requests_cpu_cores{namespace="kube-system",pod="kube-apiserver-kind-control-plane",container="kube-apiserver",node="kind-control-plane"} 0.25
		kube_pod_container_resource_requests_cpu_cores{namespace="kube-system",pod="kube-scheduler-kind-control-plane",container="kube-scheduler",node="kind-control-plane"} 0.1
		kube_pod_container_resource_requests_cpu_cores{namespace="kube-system",pod="kube-controller-manager-kind-control-plane",container="kube-controller-manager",node="kind-control-plane"} 0.2
		kube_pod_container_resource_requests_cpu_cores{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",node="kind-control-plane"} 0.1
		# HELP kube_init_pod_container_resource_requests_cpu_cores The number of requested cpu cores by an init container.
		# TYPE kube_init_pod_container_resource_requests_cpu_cores gauge
		# HELP kube_pod_container_resource_requests_memory_bytes The number of requested memory bytes by a container.
		# TYPE kube_pod_container_resource_requests_memory_bytes gauge
		kube_pod_container_resource_requests_memory_bytes{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",node="kind-control-plane"} 7.340032e+07
		kube_pod_container_resource_requests_memory_bytes{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",node="kind-control-plane"} 7.340032e+07
		kube_pod_container_resource_requests_memory_bytes{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",node="kind-control-plane"} 5.24288e+07
		# HELP kube_pod_container_resource_limits_cpu_cores The limit on cpu cores to be used by a container.
		# TYPE kube_pod_container_resource_limits_cpu_cores gauge
		kube_pod_container_resource_limits_cpu_cores{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",node="kind-control-plane"} 0.1
		# HELP kube_pod_container_resource_limits_memory_bytes The limit on memory to be used by a container in bytes.
		# TYPE kube_pod_container_resource_limits_memory_bytes gauge
		kube_pod_container_resource_limits_memory_bytes{namespace="kube-system",pod="coredns-fb8b8dccf-2rzx7",container="coredns",node="kind-control-plane"} 1.7825792e+08
		kube_pod_container_resource_limits_memory_bytes{namespace="kube-system",pod="coredns-fb8b8dccf-lzl45",container="coredns",node="kind-control-plane"} 1.7825792e+08
		kube_pod_container_resource_limits_memory_bytes{namespace="kube-system",pod="kindnet-wzfk8",container="kindnet-cni",node="kind-control-plane"} 5.24288e+07
		# HELP kube_pod_spec_volumes_persistentvolumeclaims_info Information about persistentvolumeclaim volumes in a pod.
		# TYPE kube_pod_spec_volumes_persistentvolumeclaims_info gauge
		# HELP kube_pod_spec_volumes_persistentvolumeclaims_readonly Describes whether a persistentvolumeclaim is mounted read only.
		# TYPE kube_pod_spec_volumes_persistentvolumeclaims_readonly gauge
	*/
	return nil
}

func replicaSetMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_replicaset_created Unix creation timestamp
		# TYPE kube_replicaset_created gauge
		kube_replicaset_created{namespace="kube-system",replicaset="coredns-fb8b8dccf"} 1.563532785e+09
		kube_replicaset_created{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47"} 1.563532907e+09
		# HELP kube_replicaset_status_replicas The number of replicas per ReplicaSet.
		# TYPE kube_replicaset_status_replicas gauge
		kube_replicaset_status_replicas{namespace="kube-system",replicaset="coredns-fb8b8dccf"} 2
		kube_replicaset_status_replicas{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47"} 1
		# HELP kube_replicaset_status_fully_labeled_replicas The number of fully labeled replicas per ReplicaSet.
		# TYPE kube_replicaset_status_fully_labeled_replicas gauge
		kube_replicaset_status_fully_labeled_replicas{namespace="kube-system",replicaset="coredns-fb8b8dccf"} 2
		kube_replicaset_status_fully_labeled_replicas{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47"} 1
		# HELP kube_replicaset_status_ready_replicas The number of ready replicas per ReplicaSet.
		# TYPE kube_replicaset_status_ready_replicas gauge
		kube_replicaset_status_ready_replicas{namespace="kube-system",replicaset="coredns-fb8b8dccf"} 2
		kube_replicaset_status_ready_replicas{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47"} 1
		# HELP kube_replicaset_status_observed_generation The generation observed by the ReplicaSet controller.
		# TYPE kube_replicaset_status_observed_generation gauge
		kube_replicaset_status_observed_generation{namespace="kube-system",replicaset="coredns-fb8b8dccf"} 1
		kube_replicaset_status_observed_generation{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47"} 1
		# HELP kube_replicaset_spec_replicas Number of desired pods for a ReplicaSet.
		# TYPE kube_replicaset_spec_replicas gauge
		kube_replicaset_spec_replicas{namespace="kube-system",replicaset="coredns-fb8b8dccf"} 2
		kube_replicaset_spec_replicas{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47"} 1
		# HELP kube_replicaset_metadata_generation Sequence number representing a specific generation of the desired state.
		# TYPE kube_replicaset_metadata_generation gauge
		kube_replicaset_metadata_generation{namespace="kube-system",replicaset="coredns-fb8b8dccf"} 1
		kube_replicaset_metadata_generation{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47"} 1
		# HELP kube_replicaset_owner Information about the ReplicaSet's owner.
		# TYPE kube_replicaset_owner gauge
		kube_replicaset_owner{namespace="kube-system",replicaset="coredns-fb8b8dccf",owner_kind="Deployment",owner_name="coredns",owner_is_controller="true"} 1
		kube_replicaset_owner{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47",owner_kind="Deployment",owner_name="kube-state-metrics",owner_is_controller="true"} 1
		# HELP kube_replicaset_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_replicaset_labels gauge
		kube_replicaset_labels{namespace="kube-system",replicaset="coredns-fb8b8dccf",label_k8s_app="kube-dns",label_pod_template_hash="fb8b8dccf"} 1
		kube_replicaset_labels{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47",label_k8s_app="kube-state-metrics",label_pod_template_hash="6696d9cc47"} 1
		# HELP kube_replicaset_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_replicaset_annotations gauge
		kube_replicaset_annotations{namespace="kube-system",replicaset="coredns-fb8b8dccf",annotation_deployment_kubernetes_io_desired_replicas="2",annotation_deployment_kubernetes_io_max_replicas="3",annotation_deployment_kubernetes_io_revision="1"} 1
		kube_replicaset_annotations{namespace="kube-system",replicaset="kube-state-metrics-6696d9cc47",annotation_deployment_kubernetes_io_desired_replicas="1",annotation_deployment_kubernetes_io_max_replicas="2",annotation_deployment_kubernetes_io_revision="1"} 1
	*/
	return nil
}

func replicationControllerMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_replicationcontroller_created Unix creation timestamp
		# TYPE kube_replicationcontroller_created gauge
		# HELP kube_replicationcontroller_status_replicas The number of replicas per ReplicationController.
		# TYPE kube_replicationcontroller_status_replicas gauge
		# HELP kube_replicationcontroller_status_fully_labeled_replicas The number of fully labeled replicas per ReplicationController.
		# TYPE kube_replicationcontroller_status_fully_labeled_replicas gauge
		# HELP kube_replicationcontroller_status_ready_replicas The number of ready replicas per ReplicationController.
		# TYPE kube_replicationcontroller_status_ready_replicas gauge
		# HELP kube_replicationcontroller_status_available_replicas The number of available replicas per ReplicationController.
		# TYPE kube_replicationcontroller_status_available_replicas gauge
		# HELP kube_replicationcontroller_status_observed_generation The generation observed by the ReplicationController controller.
		# TYPE kube_replicationcontroller_status_observed_generation gauge
		# HELP kube_replicationcontroller_spec_replicas Number of desired pods for a ReplicationController.
		# TYPE kube_replicationcontroller_spec_replicas gauge
		# HELP kube_replicationcontroller_metadata_generation Sequence number representing a specific generation of the desired state.
		# TYPE kube_replicationcontroller_metadata_generation gauge
		# HELP kube_replicationcontroller_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_replicationcontroller_annotations gauge
	*/
	return nil
}

func resourceQuotaMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_resourcequota_created Unix creation timestamp
		# TYPE kube_resourcequota_created gauge
		# HELP kube_resourcequota Information about resource quota.
		# TYPE kube_resourcequota gauge
		# HELP kube_resourcequota_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_resourcequota_annotations gauge

	*/
	return nil
}

func secretMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_secret_info Information about secret.
		# TYPE kube_secret_info gauge
		kube_secret_info{namespace="kube-system",secret="deployment-controller-token-n8c5t"} 1
		kube_secret_info{namespace="kube-system",secret="namespace-controller-token-zdnml"} 1
		kube_secret_info{namespace="kube-system",secret="attachdetach-controller-token-8jqrt"} 1
		kube_secret_info{namespace="kube-system",secret="generic-garbage-collector-token-dqfpm"} 1
		kube_secret_info{namespace="kube-system",secret="disruption-controller-token-klfts"} 1
		kube_secret_info{namespace="kube-system",secret="statefulset-controller-token-hhx6t"} 1
		kube_secret_info{namespace="kube-system",secret="pv-protection-controller-token-mlmkr"} 1
		kube_secret_info{namespace="kube-public",secret="default-token-5q227"} 1
		kube_secret_info{namespace="kube-system",secret="service-account-controller-token-cbkhv"} 1
		kube_secret_info{namespace="kube-system",secret="bootstrap-signer-token-kkmtr"} 1
		kube_secret_info{namespace="kube-system",secret="node-controller-token-vzf7x"} 1
		kube_secret_info{namespace="kube-system",secret="cronjob-controller-token-lr7tf"} 1
		kube_secret_info{namespace="kube-system",secret="kindnet-token-pstlq"} 1
		kube_secret_info{namespace="kube-system",secret="service-controller-token-mvqks"} 1
		kube_secret_info{namespace="kube-system",secret="kube-proxy-token-wdmjk"} 1
		kube_secret_info{namespace="kube-system",secret="ttl-controller-token-k85qk"} 1
		kube_secret_info{namespace="kube-system",secret="pvc-protection-controller-token-htsp2"} 1
		kube_secret_info{namespace="kube-system",secret="token-cleaner-token-2cpt7"} 1
		kube_secret_info{namespace="kube-system",secret="job-controller-token-j8zjw"} 1
		kube_secret_info{namespace="kube-system",secret="certificate-controller-token-gmzlx"} 1
		kube_secret_info{namespace="kube-system",secret="replicaset-controller-token-ncmp9"} 1
		kube_secret_info{namespace="kube-system",secret="clusterrole-aggregation-controller-token-pl2nh"} 1
		kube_secret_info{namespace="kube-system",secret="daemon-set-controller-token-xm5hp"} 1
		kube_secret_info{namespace="kube-system",secret="coredns-token-dpbwm"} 1
		kube_secret_info{namespace="kube-node-lease",secret="default-token-fwf9l"} 1
		kube_secret_info{namespace="kube-system",secret="horizontal-pod-autoscaler-token-dqrqw"} 1
		kube_secret_info{namespace="kube-system",secret="persistent-volume-binder-token-rcxjh"} 1
		kube_secret_info{namespace="kube-system",secret="kube-state-metrics-token-b7wx5"} 1
		kube_secret_info{namespace="kube-system",secret="endpoint-controller-token-vkw8g"} 1
		kube_secret_info{namespace="default",secret="default-token-j47pd"} 1
		kube_secret_info{namespace="kube-system",secret="resourcequota-controller-token-fwnwk"} 1
		kube_secret_info{namespace="kube-system",secret="replication-controller-token-5hbfj"} 1
		kube_secret_info{namespace="kube-system",secret="bootstrap-token-abcdef"} 1
		kube_secret_info{namespace="kube-system",secret="ip-masq-agent-token-tmdfq"} 1
		kube_secret_info{namespace="kube-system",secret="pod-garbage-collector-token-2wtqr"} 1
		kube_secret_info{namespace="kube-system",secret="default-token-gl89g"} 1
		kube_secret_info{namespace="kube-system",secret="expand-controller-token-jcl6d"} 1
		kube_secret_info{namespace="kube-system",secret="metricbeat-kube-token-nnqc4"} 1
		# HELP kube_secret_type Type about secret.
		# TYPE kube_secret_type gauge
		kube_secret_type{namespace="kube-system",secret="pod-garbage-collector-token-2wtqr",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="default-token-gl89g",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="expand-controller-token-jcl6d",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="metricbeat-kube-token-nnqc4",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="deployment-controller-token-n8c5t",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="namespace-controller-token-zdnml",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="attachdetach-controller-token-8jqrt",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="generic-garbage-collector-token-dqfpm",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="disruption-controller-token-klfts",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="statefulset-controller-token-hhx6t",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="pv-protection-controller-token-mlmkr",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-public",secret="default-token-5q227",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="service-account-controller-token-cbkhv",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="bootstrap-signer-token-kkmtr",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="node-controller-token-vzf7x",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="cronjob-controller-token-lr7tf",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="kindnet-token-pstlq",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="service-controller-token-mvqks",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="kube-proxy-token-wdmjk",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="ttl-controller-token-k85qk",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="pvc-protection-controller-token-htsp2",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="token-cleaner-token-2cpt7",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="job-controller-token-j8zjw",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="certificate-controller-token-gmzlx",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="replicaset-controller-token-ncmp9",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="clusterrole-aggregation-controller-token-pl2nh",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="daemon-set-controller-token-xm5hp",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="coredns-token-dpbwm",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-node-lease",secret="default-token-fwf9l",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="horizontal-pod-autoscaler-token-dqrqw",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="persistent-volume-binder-token-rcxjh",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="kube-state-metrics-token-b7wx5",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="endpoint-controller-token-vkw8g",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="default",secret="default-token-j47pd",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="resourcequota-controller-token-fwnwk",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="replication-controller-token-5hbfj",type="kubernetes.io/service-account-token"} 1
		kube_secret_type{namespace="kube-system",secret="bootstrap-token-abcdef",type="bootstrap.kubernetes.io/token"} 1
		kube_secret_type{namespace="kube-system",secret="ip-masq-agent-token-tmdfq",type="kubernetes.io/service-account-token"} 1
		# HELP kube_secret_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_secret_labels gauge
		kube_secret_labels{namespace="kube-system",secret="bootstrap-signer-token-kkmtr"} 1
		kube_secret_labels{namespace="kube-system",secret="node-controller-token-vzf7x"} 1
		kube_secret_labels{namespace="kube-system",secret="generic-garbage-collector-token-dqfpm"} 1
		kube_secret_labels{namespace="kube-system",secret="disruption-controller-token-klfts"} 1
		kube_secret_labels{namespace="kube-system",secret="statefulset-controller-token-hhx6t"} 1
		kube_secret_labels{namespace="kube-system",secret="pv-protection-controller-token-mlmkr"} 1
		kube_secret_labels{namespace="kube-public",secret="default-token-5q227"} 1
		kube_secret_labels{namespace="kube-system",secret="service-account-controller-token-cbkhv"} 1
		kube_secret_labels{namespace="kube-system",secret="cronjob-controller-token-lr7tf"} 1
		kube_secret_labels{namespace="kube-system",secret="kindnet-token-pstlq"} 1
		kube_secret_labels{namespace="kube-system",secret="service-controller-token-mvqks"} 1
		kube_secret_labels{namespace="kube-system",secret="kube-proxy-token-wdmjk"} 1
		kube_secret_labels{namespace="kube-system",secret="replicaset-controller-token-ncmp9"} 1
		kube_secret_labels{namespace="kube-system",secret="clusterrole-aggregation-controller-token-pl2nh"} 1
		kube_secret_labels{namespace="kube-system",secret="ttl-controller-token-k85qk"} 1
		kube_secret_labels{namespace="kube-system",secret="pvc-protection-controller-token-htsp2"} 1
		kube_secret_labels{namespace="kube-system",secret="token-cleaner-token-2cpt7"} 1
		kube_secret_labels{namespace="kube-system",secret="job-controller-token-j8zjw"} 1
		kube_secret_labels{namespace="kube-system",secret="certificate-controller-token-gmzlx"} 1
		kube_secret_labels{namespace="kube-system",secret="horizontal-pod-autoscaler-token-dqrqw"} 1
		kube_secret_labels{namespace="kube-system",secret="persistent-volume-binder-token-rcxjh"} 1
		kube_secret_labels{namespace="kube-system",secret="daemon-set-controller-token-xm5hp"} 1
		kube_secret_labels{namespace="kube-system",secret="coredns-token-dpbwm"} 1
		kube_secret_labels{namespace="kube-node-lease",secret="default-token-fwf9l"} 1
		kube_secret_labels{namespace="kube-system",secret="resourcequota-controller-token-fwnwk"} 1
		kube_secret_labels{namespace="kube-system",secret="kube-state-metrics-token-b7wx5"} 1
		kube_secret_labels{namespace="kube-system",secret="endpoint-controller-token-vkw8g"} 1
		kube_secret_labels{namespace="default",secret="default-token-j47pd"} 1
		kube_secret_labels{namespace="kube-system",secret="replication-controller-token-5hbfj"} 1
		kube_secret_labels{namespace="kube-system",secret="bootstrap-token-abcdef"} 1
		kube_secret_labels{namespace="kube-system",secret="ip-masq-agent-token-tmdfq"} 1
		kube_secret_labels{namespace="kube-system",secret="metricbeat-kube-token-nnqc4"} 1
		kube_secret_labels{namespace="kube-system",secret="pod-garbage-collector-token-2wtqr"} 1
		kube_secret_labels{namespace="kube-system",secret="default-token-gl89g"} 1
		kube_secret_labels{namespace="kube-system",secret="expand-controller-token-jcl6d"} 1
		kube_secret_labels{namespace="kube-system",secret="deployment-controller-token-n8c5t"} 1
		kube_secret_labels{namespace="kube-system",secret="namespace-controller-token-zdnml"} 1
		kube_secret_labels{namespace="kube-system",secret="attachdetach-controller-token-8jqrt"} 1
		# HELP kube_secret_created Unix creation timestamp
		# TYPE kube_secret_created gauge
		kube_secret_created{namespace="kube-system",secret="service-controller-token-mvqks"} 1.563532775e+09
		kube_secret_created{namespace="kube-system",secret="kube-proxy-token-wdmjk"} 1.563532771e+09
		kube_secret_created{namespace="kube-system",secret="kindnet-token-pstlq"} 1.563532772e+09
		kube_secret_created{namespace="kube-system",secret="job-controller-token-j8zjw"} 1.56353277e+09
		kube_secret_created{namespace="kube-system",secret="certificate-controller-token-gmzlx"} 1.56353277e+09
		kube_secret_created{namespace="kube-system",secret="replicaset-controller-token-ncmp9"} 1.563532772e+09
		kube_secret_created{namespace="kube-system",secret="clusterrole-aggregation-controller-token-pl2nh"} 1.563532771e+09
		kube_secret_created{namespace="kube-system",secret="ttl-controller-token-k85qk"} 1.563532773e+09
		kube_secret_created{namespace="kube-system",secret="pvc-protection-controller-token-htsp2"} 1.563532771e+09
		kube_secret_created{namespace="kube-system",secret="token-cleaner-token-2cpt7"} 1.563532785e+09
		kube_secret_created{namespace="kube-system",secret="coredns-token-dpbwm"} 1.563532771e+09
		kube_secret_created{namespace="kube-node-lease",secret="default-token-fwf9l"} 1.563532785e+09
		kube_secret_created{namespace="kube-system",secret="horizontal-pod-autoscaler-token-dqrqw"} 1.563532772e+09
		kube_secret_created{namespace="kube-system",secret="persistent-volume-binder-token-rcxjh"} 1.563532774e+09
		kube_secret_created{namespace="kube-system",secret="daemon-set-controller-token-xm5hp"} 1.56353277e+09
		kube_secret_created{namespace="kube-system",secret="endpoint-controller-token-vkw8g"} 1.56353277e+09
		kube_secret_created{namespace="default",secret="default-token-j47pd"} 1.563532785e+09
		kube_secret_created{namespace="kube-system",secret="resourcequota-controller-token-fwnwk"} 1.563532774e+09
		kube_secret_created{namespace="kube-system",secret="kube-state-metrics-token-b7wx5"} 1.563532907e+09
		kube_secret_created{namespace="kube-system",secret="bootstrap-token-abcdef"} 1.563532771e+09
		kube_secret_created{namespace="kube-system",secret="ip-masq-agent-token-tmdfq"} 1.563532772e+09
		kube_secret_created{namespace="kube-system",secret="replication-controller-token-5hbfj"} 1.563532785e+09
		kube_secret_created{namespace="kube-system",secret="default-token-gl89g"} 1.563532785e+09
		kube_secret_created{namespace="kube-system",secret="expand-controller-token-jcl6d"} 1.563532785e+09
		kube_secret_created{namespace="kube-system",secret="metricbeat-kube-token-nnqc4"} 1.563532937e+09
		kube_secret_created{namespace="kube-system",secret="pod-garbage-collector-token-2wtqr"} 1.563532772e+09
		kube_secret_created{namespace="kube-system",secret="namespace-controller-token-zdnml"} 1.56353277e+09
		kube_secret_created{namespace="kube-system",secret="attachdetach-controller-token-8jqrt"} 1.563532771e+09
		kube_secret_created{namespace="kube-system",secret="deployment-controller-token-n8c5t"} 1.563532774e+09
		kube_secret_created{namespace="kube-public",secret="default-token-5q227"} 1.563532785e+09
		kube_secret_created{namespace="kube-system",secret="service-account-controller-token-cbkhv"} 1.56353277e+09
		kube_secret_created{namespace="kube-system",secret="bootstrap-signer-token-kkmtr"} 1.563532773e+09
		kube_secret_created{namespace="kube-system",secret="node-controller-token-vzf7x"} 1.56353277e+09
		kube_secret_created{namespace="kube-system",secret="generic-garbage-collector-token-dqfpm"} 1.563532771e+09
		kube_secret_created{namespace="kube-system",secret="disruption-controller-token-klfts"} 1.563532785e+09
		kube_secret_created{namespace="kube-system",secret="statefulset-controller-token-hhx6t"} 1.563532775e+09
		kube_secret_created{namespace="kube-system",secret="pv-protection-controller-token-mlmkr"} 1.563532785e+09
		kube_secret_created{namespace="kube-system",secret="cronjob-controller-token-lr7tf"} 1.563532774e+09
		# HELP kube_secret_metadata_resource_version Resource version representing a specific version of secret.
		# TYPE kube_secret_metadata_resource_version gauge
		kube_secret_metadata_resource_version{namespace="kube-system",secret="deployment-controller-token-n8c5t",resource_version="288"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="namespace-controller-token-zdnml",resource_version="168"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="attachdetach-controller-token-8jqrt",resource_version="208"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="pv-protection-controller-token-mlmkr",resource_version="317"} 1
		kube_secret_metadata_resource_version{namespace="kube-public",secret="default-token-5q227",resource_version="336"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="service-account-controller-token-cbkhv",resource_version="175"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="bootstrap-signer-token-kkmtr",resource_version="275"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="node-controller-token-vzf7x",resource_version="187"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="generic-garbage-collector-token-dqfpm",resource_version="231"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="disruption-controller-token-klfts",resource_version="323"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="statefulset-controller-token-hhx6t",resource_version="297"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="cronjob-controller-token-lr7tf",resource_version="291"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="kindnet-token-pstlq",resource_version="251"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="service-controller-token-mvqks",resource_version="294"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="kube-proxy-token-wdmjk",resource_version="212"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="token-cleaner-token-2cpt7",resource_version="326"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="job-controller-token-j8zjw",resource_version="184"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="certificate-controller-token-gmzlx",resource_version="181"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="replicaset-controller-token-ncmp9",resource_version="238"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="clusterrole-aggregation-controller-token-pl2nh",resource_version="216"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="ttl-controller-token-k85qk",resource_version="270"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="pvc-protection-controller-token-htsp2",resource_version="226"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="daemon-set-controller-token-xm5hp",resource_version="178"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="coredns-token-dpbwm",resource_version="202"} 1
		kube_secret_metadata_resource_version{namespace="kube-node-lease",secret="default-token-fwf9l",resource_version="338"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="horizontal-pod-autoscaler-token-dqrqw",resource_version="250"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="persistent-volume-binder-token-rcxjh",resource_version="279"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="kube-state-metrics-token-b7wx5",resource_version="605"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="endpoint-controller-token-vkw8g",resource_version="171"} 1
		kube_secret_metadata_resource_version{namespace="default",secret="default-token-j47pd",resource_version="346"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="resourcequota-controller-token-fwnwk",resource_version="283"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="replication-controller-token-5hbfj",resource_version="320"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="bootstrap-token-abcdef",resource_version="191"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="ip-masq-agent-token-tmdfq",resource_version="261"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="pod-garbage-collector-token-2wtqr",resource_version="241"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="default-token-gl89g",resource_version="355"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="expand-controller-token-jcl6d",resource_version="314"} 1
		kube_secret_metadata_resource_version{namespace="kube-system",secret="metricbeat-kube-token-nnqc4",resource_version="667"} 1
		# HELP kube_secret_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_secret_annotations gauge
		kube_secret_annotations{namespace="kube-node-lease",secret="default-token-fwf9l",annotation_kubernetes_io_service_account_uid="86735724-aa11-11e9-9d2e-0242ac110002",annotation_kubernetes_io_service_account_name="default"} 1
		kube_secret_annotations{namespace="kube-system",secret="horizontal-pod-autoscaler-token-dqrqw",annotation_kubernetes_io_service_account_name="horizontal-pod-autoscaler",annotation_kubernetes_io_service_account_uid="7eb47b00-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="persistent-volume-binder-token-rcxjh",annotation_kubernetes_io_service_account_name="persistent-volume-binder",annotation_kubernetes_io_service_account_uid="7f6b9d70-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="daemon-set-controller-token-xm5hp",annotation_kubernetes_io_service_account_name="daemon-set-controller",annotation_kubernetes_io_service_account_uid="7d916643-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="coredns-token-dpbwm",annotation_kubernetes_io_service_account_name="coredns",annotation_kubernetes_io_service_account_uid="7db5a0a7-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="default",secret="default-token-j47pd",annotation_kubernetes_io_service_account_name="default",annotation_kubernetes_io_service_account_uid="8673e71e-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="resourcequota-controller-token-fwnwk",annotation_kubernetes_io_service_account_name="resourcequota-controller",annotation_kubernetes_io_service_account_uid="7f91bf73-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="kube-state-metrics-token-b7wx5",annotation_kubernetes_io_service_account_name="kube-state-metrics",annotation_kubernetes_io_service_account_uid="cef546e8-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="endpoint-controller-token-vkw8g",annotation_kubernetes_io_service_account_name="endpoint-controller",annotation_kubernetes_io_service_account_uid="7d8dd89a-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="ip-masq-agent-token-tmdfq",annotation_kubernetes_io_service_account_name="ip-masq-agent",annotation_kubernetes_io_service_account_uid="7eb7a315-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="replication-controller-token-5hbfj",annotation_kubernetes_io_service_account_name="replication-controller",annotation_kubernetes_io_service_account_uid="8666e859-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="bootstrap-token-abcdef"} 1
		kube_secret_annotations{namespace="kube-system",secret="expand-controller-token-jcl6d",annotation_kubernetes_io_service_account_name="expand-controller",annotation_kubernetes_io_service_account_uid="865f5710-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="metricbeat-kube-token-nnqc4",annotation_kubernetes_io_service_account_name="metricbeat-kube",annotation_kubernetes_io_service_account_uid="e0eabb54-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="pod-garbage-collector-token-2wtqr",annotation_kubernetes_io_service_account_name="pod-garbage-collector",annotation_kubernetes_io_service_account_uid="7e937715-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="default-token-gl89g",annotation_kubernetes_io_service_account_name="default",annotation_kubernetes_io_service_account_uid="8674d10e-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="attachdetach-controller-token-8jqrt",annotation_kubernetes_io_service_account_name="attachdetach-controller",annotation_kubernetes_io_service_account_uid="7db8b888-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="deployment-controller-token-n8c5t",annotation_kubernetes_io_service_account_name="deployment-controller",annotation_kubernetes_io_service_account_uid="7fc07099-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="namespace-controller-token-zdnml",annotation_kubernetes_io_service_account_name="namespace-controller",annotation_kubernetes_io_service_account_uid="7d7c29b4-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="service-account-controller-token-cbkhv",annotation_kubernetes_io_service_account_name="service-account-controller",annotation_kubernetes_io_service_account_uid="7d8fa0e4-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="bootstrap-signer-token-kkmtr",annotation_kubernetes_io_service_account_uid="7f457204-aa11-11e9-9d2e-0242ac110002",annotation_kubernetes_io_service_account_name="bootstrap-signer"} 1
		kube_secret_annotations{namespace="kube-system",secret="node-controller-token-vzf7x",annotation_kubernetes_io_service_account_name="node-controller",annotation_kubernetes_io_service_account_uid="7d95fdb3-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="generic-garbage-collector-token-dqfpm",annotation_kubernetes_io_service_account_uid="7e2b2a4e-aa11-11e9-9d2e-0242ac110002",annotation_kubernetes_io_service_account_name="generic-garbage-collector"} 1
		kube_secret_annotations{namespace="kube-system",secret="disruption-controller-token-klfts",annotation_kubernetes_io_service_account_uid="86695faa-aa11-11e9-9d2e-0242ac110002",annotation_kubernetes_io_service_account_name="disruption-controller"} 1
		kube_secret_annotations{namespace="kube-system",secret="statefulset-controller-token-hhx6t",annotation_kubernetes_io_service_account_name="statefulset-controller",annotation_kubernetes_io_service_account_uid="802a5793-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="pv-protection-controller-token-mlmkr",annotation_kubernetes_io_service_account_name="pv-protection-controller",annotation_kubernetes_io_service_account_uid="8663c5a7-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-public",secret="default-token-5q227",annotation_kubernetes_io_service_account_uid="8672ba60-aa11-11e9-9d2e-0242ac110002",annotation_kubernetes_io_service_account_name="default"} 1
		kube_secret_annotations{namespace="kube-system",secret="cronjob-controller-token-lr7tf",annotation_kubernetes_io_service_account_uid="7fde0b4f-aa11-11e9-9d2e-0242ac110002",annotation_kubernetes_io_service_account_name="cronjob-controller"} 1
		kube_secret_annotations{namespace="kube-system",secret="kube-proxy-token-wdmjk",annotation_kubernetes_io_service_account_name="kube-proxy",annotation_kubernetes_io_service_account_uid="7dcf02b9-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="kindnet-token-pstlq",annotation_kubernetes_io_service_account_name="kindnet",annotation_kubernetes_io_service_account_uid="7eb4efc2-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="service-controller-token-mvqks",annotation_kubernetes_io_service_account_name="service-controller",annotation_kubernetes_io_service_account_uid="80043041-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="certificate-controller-token-gmzlx",annotation_kubernetes_io_service_account_name="certificate-controller",annotation_kubernetes_io_service_account_uid="7d92ce8b-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="replicaset-controller-token-ncmp9",annotation_kubernetes_io_service_account_uid="7e8f6d70-aa11-11e9-9d2e-0242ac110002",annotation_kubernetes_io_service_account_name="replicaset-controller"} 1
		kube_secret_annotations{namespace="kube-system",secret="clusterrole-aggregation-controller-token-pl2nh",annotation_kubernetes_io_service_account_uid="7ddee090-aa11-11e9-9d2e-0242ac110002",annotation_kubernetes_io_service_account_name="clusterrole-aggregation-controller"} 1
		kube_secret_annotations{namespace="kube-system",secret="ttl-controller-token-k85qk",annotation_kubernetes_io_service_account_name="ttl-controller",annotation_kubernetes_io_service_account_uid="7f1f5527-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="pvc-protection-controller-token-htsp2",annotation_kubernetes_io_service_account_name="pvc-protection-controller",annotation_kubernetes_io_service_account_uid="7e050740-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="token-cleaner-token-2cpt7",annotation_kubernetes_io_service_account_name="token-cleaner",annotation_kubernetes_io_service_account_uid="866b578a-aa11-11e9-9d2e-0242ac110002"} 1
		kube_secret_annotations{namespace="kube-system",secret="job-controller-token-j8zjw",annotation_kubernetes_io_service_account_name="job-controller",annotation_kubernetes_io_service_account_uid="7d947e54-aa11-11e9-9d2e-0242ac110002"} 1
	*/

	return nil
}

func serviceMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_service_info Information about service.
		# TYPE kube_service_info gauge
		kube_service_info{namespace="kube-system",service="kube-state-metrics",cluster_ip="10.105.144.35",external_name="",load_balancer_ip=""} 1
		kube_service_info{namespace="default",service="kubernetes",cluster_ip="10.96.0.1",external_name="",load_balancer_ip=""} 1
		kube_service_info{namespace="kube-system",service="kube-dns",cluster_ip="10.96.0.10",external_name="",load_balancer_ip=""} 1
		# HELP kube_service_created Unix creation timestamp
		# TYPE kube_service_created gauge
		kube_service_created{namespace="default",service="kubernetes"} 1.563532769e+09
		kube_service_created{namespace="kube-system",service="kube-dns"} 1.563532771e+09
		kube_service_created{namespace="kube-system",service="kube-state-metrics"} 1.563532907e+09
		# HELP kube_service_spec_type Type about service.
		# TYPE kube_service_spec_type gauge
		kube_service_spec_type{namespace="default",service="kubernetes",type="ClusterIP"} 1
		kube_service_spec_type{namespace="kube-system",service="kube-dns",type="ClusterIP"} 1
		kube_service_spec_type{namespace="kube-system",service="kube-state-metrics",type="ClusterIP"} 1
		# HELP kube_service_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_service_labels gauge
		kube_service_labels{namespace="default",service="kubernetes",label_component="apiserver",label_provider="kubernetes"} 1
		kube_service_labels{namespace="kube-system",service="kube-dns",label_k8s_app="kube-dns",label_kubernetes_io_cluster_service="true",label_kubernetes_io_name="KubeDNS"} 1
		kube_service_labels{namespace="kube-system",service="kube-state-metrics",label_k8s_app="kube-state-metrics"} 1
		# HELP kube_service_spec_external_ip Service external ips. One series for each ip
		# TYPE kube_service_spec_external_ip gauge
		# HELP kube_service_status_load_balancer_ingress Service load balancer ingress status
		# TYPE kube_service_status_load_balancer_ingress gauge
		# HELP kube_service_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_service_annotations gauge
		kube_service_annotations{namespace="default",service="kubernetes"} 1
		kube_service_annotations{namespace="kube-system",service="kube-dns",annotation_prometheus_io_scrape="true",annotation_prometheus_io_port="9153"} 1
		kube_service_annotations{namespace="kube-system",service="kube-state-metrics",annotation_prometheus_io_scrape="true"} 1
	*/
	return nil
}

func statefulSetMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_statefulset_created Unix creation timestamp
		# TYPE kube_statefulset_created gauge
		# HELP kube_statefulset_status_replicas The number of replicas per StatefulSet.
		# TYPE kube_statefulset_status_replicas gauge
		# HELP kube_statefulset_status_replicas_current The number of current replicas per StatefulSet.
		# TYPE kube_statefulset_status_replicas_current gauge
		# HELP kube_statefulset_status_replicas_ready The number of ready replicas per StatefulSet.
		# TYPE kube_statefulset_status_replicas_ready gauge
		# HELP kube_statefulset_status_replicas_updated The number of updated replicas per StatefulSet.
		# TYPE kube_statefulset_status_replicas_updated gauge
		# HELP kube_statefulset_status_observed_generation The generation observed by the StatefulSet controller.
		# TYPE kube_statefulset_status_observed_generation gauge
		# HELP kube_statefulset_replicas Number of desired pods for a StatefulSet.
		# TYPE kube_statefulset_replicas gauge
		# HELP kube_statefulset_metadata_generation Sequence number representing a specific generation of the desired state for the StatefulSet.
		# TYPE kube_statefulset_metadata_generation gauge
		# HELP kube_statefulset_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_statefulset_labels gauge
		# HELP kube_statefulset_status_current_revision Indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
		# TYPE kube_statefulset_status_current_revision gauge
		# HELP kube_statefulset_status_update_revision Indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
		# TYPE kube_statefulset_status_update_revision gauge
		# HELP kube_statefulset_annotations Kubernetes annotations converted to Prometheus labels.
		# TYPE kube_statefulset_annotations gauge
	*/
	return nil
}

func storageClassMappings() map[string]prometheus.MetricMap {
	/*
		# HELP kube_storageclass_info Information about storageclass.
		# TYPE kube_storageclass_info gauge
		kube_storageclass_info{storageclass="standard",provisioner="kubernetes.io/host-path",reclaimPolicy="Delete",volumeBindingMode="Immediate"} 1
		# HELP kube_storageclass_created Unix creation timestamp
		# TYPE kube_storageclass_created gauge
		kube_storageclass_created{storageclass="standard"} 1.563532773e+09
		# HELP kube_storageclass_labels Kubernetes labels converted to Prometheus labels.
		# TYPE kube_storageclass_labels gauge
		kube_storageclass_labels{storageclass="standard",label_addonmanager_kubernetes_io_mode="EnsureExists"} 1
	*/
	return nil
}
