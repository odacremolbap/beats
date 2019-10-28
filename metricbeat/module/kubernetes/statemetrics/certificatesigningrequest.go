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

package statemetrics

import "github.com/elastic/beats/metricbeat/helper/prometheus"

func getCertificateSigningRequestMapping() (mm map[string]prometheus.MetricMap, lm map[string]prometheus.LabelMap) {
	return map[string]prometheus.MetricMap{
			"kube_certificatesigningrequest_cert_length": prometheus.Metric("certificatesigningrequest.cert.length"),
			"kube_certificatesigningrequest_labels": prometheus.ExtendedInfoMetric(
				prometheus.Configuration{StoreNonMappedLabels: true, NonMappedLabelsPlacement: "labels"},
			),
			"kube_certificatesigningrequest_condition": prometheus.Metric("certificatesigningrequest.condition.count"),
			"kube_certificatesigningrequest_created":   prometheus.Metric("certificatesigningrequest.created", prometheus.OpUnixTimestampValue()),
		},
		map[string]prometheus.LabelMap{
			"certificatesigningrequest": prometheus.KeyLabel("certificatesigningrequest.name"),
			"condition":                 prometheus.KeyLabel("certificatesigningrequest.condition.status"),
		}
}
