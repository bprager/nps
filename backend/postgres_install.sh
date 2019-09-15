helm install --name database -f postgres_init.yaml stable/postgresql --set postgresqlPassword=secret
