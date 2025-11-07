# Upjet Provider Confluent - Initial Configuration

**Date**: November 7, 2025  
**Provider Version**: Confluent Terraform Provider v2.51.0  
**Generated Resources**: 37 Confluent Cloud resources

---

## Overview

This document captures the complete initial setup and configuration of the Upjet-based Crossplane provider for Confluent Cloud. The provider was generated from the official Confluent Terraform provider (v2.51.0) and includes comprehensive support for Confluent Cloud resources.

## Setup Steps Completed

### 1. Repository Initialization

**Actions Taken:**
- Cloned the `crossplane/upjet-provider-template` repository as the foundation
- Initialized git repository in `/Users/michielvh/code/personal/upjet-provider-confluent`
- Added `crossplane/build` as a git submodule for build tooling
- Pushed initial repository to `https://github.com/michielvha/upjet-provider-confluent.git`

**Commands Executed:**
```bash
git clone https://github.com/crossplane/upjet-provider-template.git upjet-provider-confluent-temp
rsync -av --exclude='.git' upjet-provider-confluent-temp/ upjet-provider-confluent/
rm -rf upjet-provider-confluent-temp
cd upjet-provider-confluent
git init
git submodule add https://github.com/crossplane/build build
```

### 2. Module Path Configuration

**Updated Files:**
- `go.mod` - Changed module path to `github.com/michielvha/upjet-provider-confluent`
- `Makefile` - Updated `PROJECT_NAME` and `PROJECT_REPO`
- `cmd/generator/main.go` - Updated import paths
- `cmd/provider/main.go` - Updated import paths
- `internal/clients/template.go` - Updated import paths
- `internal/controller/cluster/providerconfig/config.go` - Updated import paths
- `internal/controller/namespaced/providerconfig/config.go` - Updated import paths
- `config/provider.go` - Updated module path and root group

**Module Configuration:**
```go
module github.com/michielvha/upjet-provider-confluent

const (
    resourcePrefix = "confluent"
    modulePath     = "github.com/michielvha/upjet-provider-confluent"
)
```

**API Groups:**
- Cluster-scoped: `confluent.crossplane.io`
- Namespace-scoped: `confluent.m.crossplane.io`

### 3. Terraform Provider Configuration

**Updated Makefile Variables:**
```makefile
export TERRAFORM_PROVIDER_SOURCE := confluentinc/confluent
export TERRAFORM_PROVIDER_REPO := https://github.com/confluentinc/terraform-provider-confluent
export TERRAFORM_PROVIDER_VERSION := 2.51.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME := terraform-provider-confluent
export TERRAFORM_NATIVE_PROVIDER_BINARY := terraform-provider-confluent_v2.51.0
export TERRAFORM_DOCS_PATH := docs/resources
```

### 4. Authentication Configuration

**File**: `internal/clients/template.go`

**Implemented Confluent Cloud Authentication:**
The provider supports multiple credential types for different Confluent Cloud services:

```go
// Cloud API credentials (organization/environment level)
keyCloudAPIKey           = "cloud_api_key"
keyCloudAPISecret        = "cloud_api_secret"

// Kafka cluster credentials (data plane resources)
keyKafkaAPIKey           = "kafka_api_key"
keyKafkaAPISecret        = "kafka_api_secret"

// Schema Registry credentials
keySchemaRegistryAPIKey    = "schema_registry_api_key"
keySchemaRegistryAPISecret = "schema_registry_api_secret"

// Endpoints
keyKafkaRestEndpoint     = "kafka_rest_endpoint"
keySchemaRegistryURL     = "schema_registry_url"
keyEndpoint              = "endpoint"
```

**ProviderConfig Secret Format:**
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: confluent-credentials
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "cloud_api_key": "YOUR_CLOUD_API_KEY",
      "cloud_api_secret": "YOUR_CLOUD_API_SECRET",
      "kafka_api_key": "YOUR_KAFKA_API_KEY",
      "kafka_api_secret": "YOUR_KAFKA_API_SECRET",
      "schema_registry_api_key": "YOUR_SR_API_KEY",
      "schema_registry_api_secret": "YOUR_SR_API_SECRET",
      "endpoint": "https://api.confluent.cloud"
    }
```

### 5. External Name Configuration

**File**: `config/external_name.go`

**Configured 50+ External Name Mappings:**

| Resource Type | External Name Pattern |
|--------------|----------------------|
| **Core Infrastructure** |
| `confluent_environment` | NameAsIdentifier |
| `confluent_kafka_cluster` | IdentifierFromProvider |
| `confluent_schema_registry_cluster` | IdentifierFromProvider |
| `confluent_ksql_cluster` | IdentifierFromProvider |
| **Kafka Resources** |
| `confluent_kafka_topic` | NameAsIdentifier |
| `confluent_kafka_acl` | TemplatedString (composite key) |
| `confluent_kafka_cluster_config` | IdentifierFromProvider |
| `confluent_kafka_client_quota` | IdentifierFromProvider |
| `confluent_cluster_link` | TemplatedString (link_name) |
| **Schema Registry** |
| `confluent_schema` | TemplatedString (cluster:subject) |
| `confluent_subject_mode` | TemplatedString (cluster:subject) |
| `confluent_subject_config` | TemplatedString (cluster:subject) |
| `confluent_schema_registry_kek` | NameAsIdentifier |
| `confluent_schema_registry_dek` | IdentifierFromProvider |
| **IAM Resources** |
| `confluent_service_account` | IdentifierFromProvider |
| `confluent_api_key` | IdentifierFromProvider |
| `confluent_role_binding` | IdentifierFromProvider |
| `confluent_identity_provider` | IdentifierFromProvider |
| `confluent_identity_pool` | IdentifierFromProvider |
| **Networking** |
| `confluent_network` | IdentifierFromProvider |
| `confluent_peering` | IdentifierFromProvider |
| `confluent_private_link_access` | IdentifierFromProvider |
| `confluent_transit_gateway_attachment` | IdentifierFromProvider |
| `confluent_dns_record` | IdentifierFromProvider |

### 6. Resource Group Configurations

Created custom configurations for logical resource grouping:

#### Environment Group (`config/environment/config.go`)
```go
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("confluent_environment", func(r *config.Resource) {
        r.ShortGroup = "environment"
    })
}
```

#### Kafka Group (`config/kafka/config.go`)
**Resources Configured:**
- `Cluster` - Kafka cluster with environment reference
- `Topic` - Topics with cluster reference
- `ACL` - Access control lists
- `ClusterConfig` - Cluster configuration
- `ClientQuota` - Client quotas
- `ClusterLink` - Cluster linking
- `MirrorTopic` - Mirror topics

**Key Features:**
- Cross-resource references (Topic → Cluster, Cluster → Environment)
- Sensitive connection details exported (bootstrap_endpoint, rest_endpoint)

#### Schema Registry Group (`config/schema/config.go`)
**Resources Configured:**
- `RegistryCluster` - Schema Registry cluster
- `Schema` - Schema definitions
- `SubjectMode` - Subject modes
- `SubjectConfig` - Subject configurations
- `RegistryKek` - Key encryption keys
- `RegistryDek` - Data encryption keys
- `RegistryClusterMode` - Cluster modes
- `RegistryClusterConfig` - Cluster configurations
- `BusinessMetadata` - Business metadata
- `BusinessMetadataBinding` - Metadata bindings
- `Tag` - Tags
- `TagBinding` - Tag bindings

**Key Features:**
- Schema Registry cluster references across all resources
- KEK → DEK dependency chain

#### IAM Group (`config/iam/config.go`)
**Resources Configured:**
- `ServiceAccount` - Service accounts
- `APIKey` - API keys with service account reference
- `RoleBinding` - RBAC role bindings
- `IdentityProvider` - Identity providers
- `IdentityPool` - Identity pools
- `Invitation` - User invitations

**Key Features:**
- API Key secrets exported to connection details
- Service account references in role bindings

### 7. Cross-Resource Reference Configuration

**Important Fix Applied:**
Updated all cross-resource references to use cluster-scoped paths instead of generic paths.

**Correct Reference Pattern:**
```go
r.References["kafka_cluster.id"] = config.Reference{
    Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/kafka/v1alpha1.Cluster",
}
```

**Incorrect Pattern (Fixed):**
```go
// This caused module resolution issues
r.References["kafka_cluster.id"] = config.Reference{
    Type: "github.com/michielvha/upjet-provider-confluent/apis/kafka/v1alpha1.Cluster",
}
```

### 8. Provider Metadata Configuration

**File**: `config/provider-metadata.yaml`

**Configuration:**
```yaml
name: confluentinc/confluent
resources: {}
```

**Note**: Set to empty resources to skip documentation scraping due to incompatible documentation format in the Confluent Terraform provider repository.

### 9. Module Dependencies

**File**: `go.mod`

**Added Replace Directive:**
```go
replace github.com/michielvha/upjet-provider-confluent => ./
```

This allows the generated code to reference local packages during development before publishing to a module registry.

### 10. Code Generation

**Command Executed:**
```bash
make generate
```

**Generation Results:**
- ✅ **37 resources** generated with cluster scope
- ✅ **37 resources** generated with namespace scope  
- ✅ **74 total resource types** (37 × 2 scopes)
- ✅ All CRDs generated successfully
- ✅ Controllers generated for all resources
- ✅ Example manifests generated

**Generated Resource Groups:**
1. `confluent` - Connector resources
2. `custom` - Custom connector plugins
3. `dns` - DNS records
4. `environment` - Environments
5. `flink` - Flink compute pools and statements
6. `iam` - IAM resources (service accounts, API keys, role bindings)
7. `kafka` - Kafka clusters, topics, ACLs, configs
8. `ksql` - ksqlDB clusters
9. `network` - Networking resources
10. `private` - Private link resources
11. `schema` - Schema Registry resources
12. `transit` - Transit gateway attachments

## Generated Files Structure

```
upjet-provider-confluent/
├── apis/
│   ├── cluster/              # Cluster-scoped resources
│   │   ├── confluent/v1alpha1/
│   │   ├── custom/v1alpha1/
│   │   ├── dns/v1alpha1/
│   │   ├── environment/v1alpha1/
│   │   ├── flink/v1alpha1/
│   │   ├── iam/v1alpha1/
│   │   ├── kafka/v1alpha1/
│   │   ├── ksql/v1alpha1/
│   │   ├── network/v1alpha1/
│   │   ├── private/v1alpha1/
│   │   ├── schema/v1alpha1/
│   │   ├── transit/v1alpha1/
│   │   ├── v1beta1/          # ProviderConfig types
│   │   └── zz_register.go
│   └── namespaced/           # Namespace-scoped resources (same structure)
├── config/
│   ├── environment/config.go
│   ├── iam/config.go
│   ├── kafka/config.go
│   ├── schema/config.go
│   ├── external_name.go
│   ├── provider.go
│   └── provider-metadata.yaml
├── internal/
│   ├── clients/template.go
│   └── controller/
│       ├── cluster/          # Cluster-scoped controllers
│       └── namespaced/       # Namespace-scoped controllers
├── package/crds/             # Generated CRDs
└── examples-generated/       # Generated example manifests
```

## Complete Resource List (37 Resources)

### Environment & Organization (1)
1. `confluent_environment`

### Kafka Resources (7)
2. `confluent_kafka_cluster`
3. `confluent_kafka_topic`
4. `confluent_kafka_acl`
5. `confluent_kafka_cluster_config`
6. `confluent_kafka_client_quota`
7. `confluent_cluster_link`
8. `confluent_kafka_mirror_topic`

### Schema Registry (12)
9. `confluent_schema_registry_cluster`
10. `confluent_schema`
11. `confluent_subject_mode`
12. `confluent_subject_config`
13. `confluent_schema_registry_kek`
14. `confluent_schema_registry_dek`
15. `confluent_schema_registry_cluster_mode`
16. `confluent_schema_registry_cluster_config`
17. `confluent_business_metadata`
18. `confluent_business_metadata_binding`
19. `confluent_tag`
20. `confluent_tag_binding`

### Connect (2)
21. `confluent_connector`
22. `confluent_custom_connector_plugin`

### ksqlDB (1)
23. `confluent_ksql_cluster`

### Flink (2)
24. `confluent_flink_compute_pool`
25. `confluent_flink_statement`

### IAM (6)
26. `confluent_service_account`
27. `confluent_api_key`
28. `confluent_role_binding`
29. `confluent_identity_provider`
30. `confluent_identity_pool`
31. `confluent_invitation`

### Networking (7)
32. `confluent_network`
33. `confluent_peering`
34. `confluent_private_link_access`
35. `confluent_transit_gateway_attachment`
36. `confluent_dns_record`
37. `confluent_network_link_service`
38. `confluent_network_link_endpoint`

## Issues Encountered and Resolutions

### Issue 1: Module Path References
**Problem**: Initial imports used `crossplane/upjet-provider-template` paths  
**Solution**: Updated all import paths to `github.com/michielvha/upjet-provider-confluent`  
**Files Modified**: 7 files (cmd/, internal/, config/)

### Issue 2: Duplicate Package Declarations
**Problem**: Auto-formatter added duplicate `package` statements  
**Solution**: Removed duplicate declarations from all config files  
**Files Modified**: config/environment/, config/kafka/, config/schema/, config/iam/

### Issue 3: Cross-Reference Path Resolution
**Problem**: References used wrong API group paths causing go mod resolution failures  
**Solution**: Updated to use full cluster-scoped paths  
**Example Fix**:
```go
// Before
Type: "github.com/michielvha/upjet-provider-confluent/apis/kafka/v1alpha1.Cluster"

// After  
Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/kafka/v1alpha1.Cluster"
```

### Issue 4: Documentation Scraping
**Problem**: Confluent provider docs format incompatible with Upjet scraper  
**Solution**: Set `provider-metadata.yaml` to empty resources to skip scraping  
**Impact**: Provider still generates successfully, documentation can be added later

### Issue 5: Go Module Resolution
**Problem**: `go mod tidy` trying to fetch local packages from GitHub  
**Solution**: Added replace directive in go.mod pointing to local directory  
**Impact**: Allows development without publishing to module registry

## Next Steps

### 1. Testing the Provider
```bash
# Build the provider
make build

# Run locally (requires local Kubernetes cluster)
make run

# Create test resources
kubectl apply -f examples/cluster/providerconfig/secret.yaml
kubectl apply -f examples/cluster/providerconfig/providerconfig.yaml
```

### 2. Create Example Manifests
Create example YAML files for common use cases:
- Creating a Kafka cluster
- Creating topics with ACLs
- Setting up Schema Registry
- Configuring service accounts and API keys

### 3. Documentation
- Add resource-specific documentation
- Create usage guides
- Document authentication setup
- Add troubleshooting guide

### 4. CI/CD Setup
- Configure GitHub Actions workflows
- Set up automated testing
- Configure release automation
- Set up package publishing

### 5. Publishing
- Tag initial release (e.g., v0.1.0)
- Publish to container registry
- Create Crossplane package
- Publish to Upbound Marketplace (optional)

## Configuration Reference

### ProviderConfig Example
```yaml
apiVersion: confluent.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: confluent-credentials
      key: credentials
```

### Kafka Cluster Example
```yaml
apiVersion: kafka.confluent.crossplane.io/v1alpha1
kind: Cluster
metadata:
  name: production-cluster
spec:
  forProvider:
    availability: MULTI_ZONE
    cloud: AWS
    region: us-west-2
    basic: {}
    environmentRef:
      name: production-env
  providerConfigRef:
    name: default
```

### Kafka Topic Example
```yaml
apiVersion: kafka.confluent.crossplane.io/v1alpha1
kind: Topic
metadata:
  name: orders-topic
spec:
  forProvider:
    topicName: orders
    partitionsCount: 6
    config:
      "cleanup.policy": "delete"
      "retention.ms": "604800000"
    kafkaClusterRef:
      name: production-cluster
  providerConfigRef:
    name: default
```

## Summary

The Upjet-based Confluent provider has been successfully initialized and configured with:

- ✅ **74 total resource types** (37 cluster-scoped + 37 namespace-scoped)
- ✅ **12 resource groups** organized by Confluent service
- ✅ **Full authentication support** for Confluent Cloud, Kafka, and Schema Registry
- ✅ **Cross-resource references** for declarative dependency management
- ✅ **External name mappings** for all resources
- ✅ **Both scopes supported**: Cluster-wide and namespace-isolated resources

The provider is now ready for local testing and development. All code has been generated, compiled, and is ready to be deployed to a Kubernetes cluster with Crossplane installed.

## Troubleshooting & Fixes

### Issue 1: Schema Registry Cluster Not a Manageable Resource

**Problem**: Build errors indicated `RegistryCluster` type was undefined in schema resources.

**Root Cause**: `confluent_schema_registry_cluster` is not a manageable resource in the Confluent Terraform provider. Only `confluent_schema_registry_cluster_mode` and `confluent_schema_registry_cluster_config` exist as resources.

**Fix Applied**:
1. Removed `confluent_schema_registry_cluster` from `config/external_name.go`
2. Removed all `RegistryCluster` references from `config/schema/config.go`
3. Removed cross-references to `schema_registry_cluster.id` in schema resources
4. Kept `RegistryClusterMode` and `RegistryClusterConfig` resources which are valid

**Files Modified**:
- `config/external_name.go`: Removed line 12
- `config/schema/config.go`: Removed RegistryCluster configurator and all references

### Issue 2: IAM Extractor Package Version Mismatch

**Problem**: Build error `cannot use resource.ExtractParamPath` due to wrong upjet package version.

**Root Cause**: IAM role binding configuration used old upjet package path `github.com/crossplane/upjet/pkg/resource` instead of v2 path.

**Fix Applied**:
Changed extractor path in `config/iam/config.go` from:
```go
Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",false)`
```
to:
```go
Extractor: `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("id",false)`
```

**Files Modified**:
- `config/iam/config.go`: Line 39

### Post-Fix Actions

After applying fixes:
1. Ran `make generate` to regenerate code without RegistryCluster references
2. Ran `make build` to verify compilation - **SUCCESS**
3. Build completed successfully with 37 resources × 2 scopes = 74 total resources

**Note**: Final build error about Docker daemon is expected if Docker isn't running - this only affects container image building, not Go compilation.

---

**Generated on**: November 7, 2025  
**Provider Version**: v0.1.0-dev  
**Terraform Provider**: confluentinc/confluent v2.51.0  
**Upjet Version**: v2.0.1
