# Kube apiserver Stats

## Version history

- Unknown, unknown
- June 2019, `v1.14.3`: this version drops `apiserver_request_count` in favor of `apiserver_request_total`. There is a new label for `dry_run`, which is rarely used. Current dashboards should be unaffected. Metric `apiserver_request_latencies` is deprecated and should be removed at next major version.

## Resources

`apiserver_request_total`
    - client
    - code
    - component 
    - contentType
    - dry_run
    - resource
    - scope
    - subresource
    - verb
    - version

`apiserver_longrunning_gauge`
    - component
    - group
    - resource
    - scope
    - subresource
    - verb
    - version

`etcd_object_counts`
    - resource

`apiserver_current_inflight_requests`
    - requestKind

`apiserver_audit_event_total`

`apiserver_audit_requests_rejected_total`





### Not included

Not including these buckets one for now, it is massive because of the different labels, and it might be not provide info that's worth it at visualizations
`apiserver_request_duration_seconds_bucket`
    - component
    - dry_run
    - group
    - resource
    - scope
    - subresource
    - verb
    - version

Not including these buckets one for now, it is massive because of the different labels, and it might be not provide info that's worth it at visualizations
`apiserver_response_sizes_bucket`
    - component
    - group
    - resource
    - scope
    - subresource
    - verb
    - version

Metrics from `apiserver_longrunning_gauge` seems to be more interesting than this one
`apiserver_registered_watchers`
    - group
    - kind
    - version