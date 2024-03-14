# How to Run

With a Kubernetes environment setup (e.g. minikube), ensure `kubectl` can access the environment.

Then do these steps in order.

Build Docker image for myworkflow service:
```
docker build -t myworkflow .
```

Install Temporal services:
```shell
# Install PostgreSQL DB
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install postgres bitnami/postgresql \
--set auth.username=temporal \
--set auth.password=temporal \
--set auth.postgresPassword=postgres
# Install Temporal services
helm install temporal helm-charts/temporal --timeout=10m
```

For the first time installation (no need to do it in future upgrades), create the default namespace in Temporal Server:
```shell
# This is a one-time step because the namespace is persisted in DB.
kubectl exec -it services/temporal-admintools /bin/bash
tctl --namespace default namespace desc
```

Install myworkflow service:
```shell
helm install myworkflow helm-charts/myworkflow
```
