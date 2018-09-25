---
title: Access OK Log
type: Details
---

To access the OK Log UI, follow these steps:

1. Run the following command to configure port forwarding:

```
kubectl port-forward -n kyma-system svc/core-logging-oklog 7650:7650
```

2. Access the OK Log UI at `http://localhost:7650`.

3. Use a plaintext or regular expression to pull up logs from OK Log. For example, `pod.name:core-kubeless`.