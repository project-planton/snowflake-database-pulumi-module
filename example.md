# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: SnowflakeDatabase
metadata:
  name: my-snowflake-kafka-cluster
spec:
  snowflakeCredentialId: my-snowflake-credential-id
```

# Example with Environment Information

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: SnowflakeDatabase
metadata:
  name: my-env-snowflake-kafka-cluster
spec:
  environmentInfo:
    envId: production-environment
  snowflakeCredentialId: prod-snowflake-credential-id
```

# Example with Stack Job Settings

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: SnowflakeDatabase
metadata:
  name: advanced-snowflake-kafka-cluster
spec:
  stackJobSettings:
    pulumiBackendCredentialId: my-pulumi-backend-credential
    stackJobRunnerId: my-stack-job-runner
  snowflakeCredentialId: advanced-snowflake-credential-id
```
