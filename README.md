# prometheus
a prometheus example repo

### build and deploy the all into the cluster

`
hack/build.sh
helm install testmetric deploy/testmetric
`

### config prometheus and reload the config
```shell script
kubectl edit configmap -n monitoring prometheus-server
 - job_name: test-metric
       kubernetes_sd_configs:
       - role: service
         namespaces:
           names:
           - default
       scheme: http
       metrics_path: /metrics
       relabel_configs:
       - action: keep
         source_labels:
         - __meta_kubernetes_service_label_helm_sh_chart
       - source_labels:
         - __meta_kubernetes_namespace
         target_label: kubernetes_namespace
       - source_labels:
         - __meta_kubernetes_service_name
         target_label: kubernetes_name
```

### check the targets and metric in prometheus
```shell script
[ec2-user@ip-10-184-255-52 monitoring]$ curl -X GET 'http://ad91ea1c352fd11eaaacd0ae8503cc73-d6c7cb6a2c1380ba.elb.cn-northwest-1.amazonaws.com.cn/api/v1/query?query=mike_test_metric'
{"status":"success","data":{"resultType":"vector","result":[{"metric":{"Label1":"A test label","__name__":"mike_test_metric","instance":"testmetric.default.svc:80","job":"test-metric","kubernetes_name":"testmetric","kubernetes_namespace":"default"},"value":[1583922716.036,"331"]}]}}
```