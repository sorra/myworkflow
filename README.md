# How to Run

1. With a Kubernetes environment setup (e.g. minikube), ensure `kubectl` can access the environment.

2. Build Docker image for myworkflow service:
    ```shell
    docker build -t myworkflow .
    ```

3. Install Temporal services:
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

4. For the first time installation (no need to do it in future upgrades), create the default namespace in Temporal Server:
    ```shell
    # This is a one-time step because the namespace is persisted in DB.
    kubectl exec -it services/temporal-admintools /bin/bash
    tctl --namespace default namespace desc
    ```

5. Install myworkflow service (5 workers, 1 of them is leader):
    ```shell
    helm install myworkflow helm-charts/myworkflow
    ```
    Now the workflows have been started, you can check myworkflow container logs.