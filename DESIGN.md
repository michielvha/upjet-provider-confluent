# Upjet Provider Confluent - Design Document

## Overview

This document outlines the design and implementation strategy for building a custom Crossplane provider for Confluent Cloud using Upjet. The provider will be generated from the official Terraform Confluent provider, ensuring full feature parity with the latest Confluent Cloud API capabilities.

## Problem Statement

The existing community `crossplane-contrib/provider-confluent` (v0.5.0) is outdated and missing critical resources available in the upstream Terraform provider, including:

- Schema Registry Cluster support
- Latest Kafka features (tiered storage, cluster linking)
- Advanced connector configurations
- ksqlDB cluster management
- Flink compute pools and statements
- Enhanced security features (RBAC, service accounts)

Building a custom Upjet-based provider allows us to:

1. **Stay current** with the latest Confluent Terraform provider releases
2. **Include all resources** automatically via Upjet code generation
3. **Customize resource configurations** to fit our specific use cases
4. **Maintain control** over upgrade cycles and features

## Architecture

### High-Level Design

```
┌─────────────────────────────────────────────────────────┐
│                  Crossplane Runtime                      │
└─────────────────────────────────────────────────────────┘
                          ▲
                          │
┌─────────────────────────┼─────────────────────────────┐
│                         │                              │
│         Upjet Provider Confluent                       │
│                                                        │
│  ┌──────────────────────────────────────────────┐    │
│  │   Generated Managed Resources (CRDs)         │    │
│  │  - Kafka Clusters                            │    │
│  │  - Kafka Topics                              │    │
│  │  - Schema Registry                           │    │
│  │  - Connectors                                │    │
│  │  - ksqlDB Clusters                           │    │
│  │  - Service Accounts & ACLs                   │    │
│  │  - And 50+ more resources...                 │    │
│  └──────────────────────────────────────────────┘    │
│                         │                              │
│  ┌──────────────────────┼─────────────────────────┐   │
│  │     Upjet Framework (v2)                       │   │
│  │  - Terraform Provider Invocation               │   │
│  │  - External Name Management                    │   │
│  │  - Cross-Resource References                   │   │
│  │  - Late Initialization                         │   │
│  └────────────────────────────────────────────────┘   │
└───────────────────────┼──────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────┐
│      Terraform Confluent Provider (Go Library)          │
└─────────────────────────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────┐
│              Confluent Cloud REST API                   │
└─────────────────────────────────────────────────────────┘
```

### Provider Configuration

The provider will support both cluster-scoped and namespace-scoped configurations to accommodate different organizational security models.

#### Authentication Methods

Confluent Cloud supports multiple authentication mechanisms:

1. **API Key/Secret** (Primary method)
   - Cloud API Key for organization-level resources
   - Cluster API Key for data plane resources
2. **Service Account** with RBAC
3. **Environment-specific credentials**

#### ProviderConfig Structure

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
  # Optional: Default configuration
  configuration:
    cloudApiKey: ""        # Set from secret
    cloudApiSecret: ""     # Set from secret
    endpoint: "https://api.confluent.cloud"
    
    # For Kafka cluster resources
    kafkaApiKey: ""        # Set from secret
    kafkaApiSecret: ""     # Set from secret
    kafkaRestEndpoint: ""
    
    # For Schema Registry resources
    schemaRegistryApiKey: ""     # Set from secret
    schemaRegistryApiSecret: ""  # Set from secret
    schemaRegistryUrl: ""
```

## Resource Organization

### API Group Strategy

Confluent resources will be organized into logical API groups based on service boundaries:

| API Group | Description | Example Resources |
|-----------|-------------|-------------------|
| `kafka` | Kafka cluster and topic resources | Cluster, Topic, ACL, ClusterLink |
| `schema` | Schema Registry resources | SchemaRegistryCluster, Schema, SchemaRegistryKek |
| `connect` | Kafka Connect resources | Connector, CustomConnectorPlugin |
| `ksql` | ksqlDB resources | KsqlCluster |
| `flink` | Flink compute resources | FlinkComputePool, FlinkStatement |
| `iam` | Identity and access management | ServiceAccount, RoleBinding, ApiKey |
| `network` | Networking resources | Network, PrivateLinkAccess, PeeringConnection |
| `environment` | Organization and environment | Environment, Organization |
| `billing` | Cost management | BillingRecord |

### Resource Naming Convention

- **Terraform Resource**: `confluent_kafka_topic`
- **Generated CRD Kind**: `Topic`
- **API Version**: `kafka.confluent.crossplane.io/v1alpha1`
- **External Name**: Actual Confluent resource ID

## External Name Configuration

External names define how Crossplane resource names map to Confluent resource identifiers.

### Common Patterns

#### 1. Name as Identifier
For resources where the name is the unique identifier:

```go
"confluent_environment": config.NameAsIdentifier,
"confluent_kafka_topic": config.NameAsIdentifier,
```

#### 2. Terraform ID
For resources using Terraform's generated ID:

```go
"confluent_service_account": config.IdentifierFromProvider,
"confluent_kafka_cluster": config.IdentifierFromProvider,
```

#### 3. Templated String
For composite identifiers:

```go
"confluent_kafka_acl": config.TemplatedStringAsIdentifier("", 
  "{{ .parameters.resource_type }}:{{ .parameters.resource_name }}:{{ .parameters.principal }}"),

"confluent_role_binding": config.TemplatedStringAsIdentifier("",
  "{{ .parameters.principal }}:{{ .parameters.role_name }}:{{ .parameters.crn_pattern }}"),
```

### Resource-Specific External Names

```go
var ExternalNameConfigs = map[string]config.ExternalName{
  // Core Infrastructure
  "confluent_environment":              config.NameAsIdentifier,
  "confluent_schema_registry_cluster":  config.IdentifierFromProvider,
  "confluent_kafka_cluster":            config.IdentifierFromProvider,
  "confluent_ksql_cluster":             config.IdentifierFromProvider,
  
  // Kafka Resources
  "confluent_kafka_topic":              config.NameAsIdentifier,
  "confluent_kafka_acl":                config.TemplatedStringAsIdentifier("", 
    "{{ .parameters.kafka_cluster.id }}:{{ .parameters.resource_type }}:{{ .parameters.resource_name }}:{{ .parameters.pattern_type }}:{{ .parameters.principal }}:{{ .parameters.host }}:{{ .parameters.operation }}:{{ .parameters.permission }}"),
  "confluent_kafka_cluster_config":     config.IdentifierFromProvider,
  "confluent_kafka_client_quota":       config.IdentifierFromProvider,
  "confluent_cluster_link":             config.TemplatedStringAsIdentifier("link_name", 
    "{{ .parameters.link_name }}"),
  
  // Schema Registry
  "confluent_schema":                   config.TemplatedStringAsIdentifier("subject_name",
    "{{ .parameters.schema_registry_cluster.id }}:{{ .parameters.subject_name }}"),
  "confluent_subject_mode":             config.TemplatedStringAsIdentifier("subject_name",
    "{{ .parameters.schema_registry_cluster.id }}:{{ .parameters.subject_name }}"),
  "confluent_subject_config":           config.TemplatedStringAsIdentifier("subject_name",
    "{{ .parameters.schema_registry_cluster.id }}:{{ .parameters.subject_name }}"),
  "confluent_schema_registry_kek":      config.NameAsIdentifier,
  "confluent_schema_registry_dek":      config.IdentifierFromProvider,
  
  // Connect
  "confluent_connector":                config.NameAsIdentifier,
  "confluent_custom_connector_plugin":  config.NameAsIdentifier,
  
  // Flink
  "confluent_flink_compute_pool":       config.IdentifierFromProvider,
  "confluent_flink_statement":          config.IdentifierFromProvider,
  
  // IAM
  "confluent_service_account":          config.IdentifierFromProvider,
  "confluent_api_key":                  config.IdentifierFromProvider,
  "confluent_role_binding":             config.IdentifierFromProvider,
  "confluent_identity_provider":        config.IdentifierFromProvider,
  "confluent_identity_pool":            config.IdentifierFromProvider,
  
  // Networking
  "confluent_network":                  config.IdentifierFromProvider,
  "confluent_peering":                  config.IdentifierFromProvider,
  "confluent_private_link_access":      config.IdentifierFromProvider,
  "confluent_transit_gateway_attachment": config.IdentifierFromProvider,
  "confluent_dns_record":               config.IdentifierFromProvider,
  "confluent_network_link_service":     config.IdentifierFromProvider,
  "confluent_network_link_endpoint":    config.IdentifierFromProvider,
  
  // Business Metadata (Schema Registry)
  "confluent_business_metadata":        config.IdentifierFromProvider,
  "confluent_business_metadata_binding": config.IdentifierFromProvider,
  "confluent_tag":                      config.IdentifierFromProvider,
  "confluent_tag_binding":              config.IdentifierFromProvider,
  
  // Invitations
  "confluent_invitation":               config.IdentifierFromProvider,
}
```

## Cross-Resource References

### Reference Patterns

#### Cluster to Environment

```go
r.References["environment"] = config.Reference{
  Type:      "github.com/michielvha/provider-confluent/apis/environment/v1alpha1.Environment",
  Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
}
```

#### Topic to Kafka Cluster

```go
r.References["kafka_cluster.id"] = config.Reference{
  Type:      "github.com/michielvha/provider-confluent/apis/kafka/v1alpha1.Cluster",
  Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
}
```

#### Schema to Schema Registry Cluster

```go
r.References["schema_registry_cluster.id"] = config.Reference{
  Type:      "github.com/michielvha/provider-confluent/apis/schema/v1alpha1.RegistryCluster",
  Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
}
```

#### Connector to Kafka Cluster

```go
r.References["kafka_cluster.id"] = config.Reference{
  Type:      "github.com/michielvha/provider-confluent/apis/kafka/v1alpha1.Cluster",
}
r.References["environment.id"] = config.Reference{
  Type:      "github.com/michielvha/provider-confluent/apis/environment/v1alpha1.Environment",
}
```

### Complex Reference Example: ACL

ACLs reference both clusters and service accounts:

```go
r.References["kafka_cluster.id"] = config.Reference{
  Type: "github.com/michielvha/provider-confluent/apis/kafka/v1alpha1.Cluster",
}
r.References["principal"] = config.Reference{
  Type:      "github.com/michielvha/provider-confluent/apis/iam/v1alpha1.ServiceAccount",
  Extractor: "github.com/michielvha/provider-confluent/config/common.ExtractPrincipalFromServiceAccount()",
}
```

## Terraform Provider Configuration

### Makefile Variables

```makefile
export TERRAFORM_PROVIDER_SOURCE := confluentinc/confluent
export TERRAFORM_PROVIDER_REPO := https://github.com/confluentinc/terraform-provider-confluent
export TERRAFORM_PROVIDER_VERSION := 2.11.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME := terraform-provider-confluent
export TERRAFORM_NATIVE_PROVIDER_BINARY := terraform-provider-confluent_v2.11.0
export TERRAFORM_DOCS_PATH := docs/resources
```

### Version Strategy

- **Initial Version**: 2.11.0 (latest stable as of design)
- **Update Strategy**: Follow Confluent provider releases
- **Testing**: Validate each version upgrade with integration tests

## Implementation Phases

### Phase 1: Foundation (Week 1)

**Goal**: Basic provider infrastructure and core resources

- [x] Repository setup from template
- [ ] Configure Makefile with Confluent provider details
- [ ] Implement ProviderConfig with authentication
- [ ] Generate initial CRDs
- [ ] Implement core resources:
  - Environment
  - Kafka Cluster
  - Kafka Topic
  - Service Account
  - API Key

**Success Criteria**: Can create a Kafka cluster and topic using the provider

### Phase 2: Schema Registry (Week 2)

**Goal**: Schema management capabilities

- [ ] Schema Registry Cluster
- [ ] Schema
- [ ] Subject Mode
- [ ] Subject Config
- [ ] Schema Registry KEK/DEK
- [ ] Business Metadata
- [ ] Tag Management

**Success Criteria**: Full schema lifecycle management with versioning

### Phase 3: Connectivity & Security (Week 3)

**Goal**: Networking, connectors, and access control

- [ ] Kafka Connect resources
- [ ] Custom connector plugins
- [ ] ACL management
- [ ] Role Bindings (RBAC)
- [ ] Network resources
- [ ] Private Link configurations
- [ ] Peering connections

**Success Criteria**: Secure multi-cluster connectivity with proper ACLs

### Phase 4: Advanced Features (Week 4)

**Goal**: ksqlDB, Flink, and advanced Kafka features

- [ ] ksqlDB Cluster
- [ ] Flink Compute Pool
- [ ] Flink Statement
- [ ] Cluster Linking
- [ ] Mirror Topics
- [ ] Kafka Client Quotas
- [ ] DNS Records
- [ ] Identity Providers/Pools

**Success Criteria**: Stream processing and cross-cluster replication

### Phase 5: Testing & Documentation (Week 5)

**Goal**: Comprehensive testing and documentation

- [ ] Unit tests for custom configurations
- [ ] Integration tests with Uptest
- [ ] Example manifests for all resources
- [ ] API documentation generation
- [ ] Migration guide from community provider
- [ ] Best practices guide

**Success Criteria**: Production-ready provider with full documentation

## Custom Resource Configurations

### Directory Structure

```
config/
├── provider.go                    # Main provider configuration
├── external_name.go               # External name configurations
├── common/                        # Shared utilities
│   ├── extractors.go             # Custom field extractors
│   └── validators.go             # Custom validators
├── kafka/                         # Kafka resources
│   └── config.go
├── schema/                        # Schema Registry resources
│   └── config.go
├── connect/                       # Kafka Connect resources
│   └── config.go
├── ksql/                          # ksqlDB resources
│   └── config.go
├── flink/                         # Flink resources
│   └── config.go
├── iam/                           # IAM resources
│   └── config.go
├── network/                       # Networking resources
│   └── config.go
└── environment/                   # Environment resources
    └── config.go
```

### Example Custom Configuration: Kafka Topic

```go
// config/kafka/config.go
package kafka

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
  p.AddResourceConfigurator("confluent_kafka_topic", func(r *config.Resource) {
    r.ShortGroup = "kafka"
    
    // Reference to Kafka Cluster
    r.References["kafka_cluster.id"] = config.Reference{
      Type: "github.com/michielvha/provider-confluent/apis/kafka/v1alpha1.Cluster",
    }
    
    // Sensitive fields
    r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
      conn := map[string][]byte{}
      if v, ok := attr["topic_name"].(string); ok {
        conn["topic_name"] = []byte(v)
      }
      return conn, nil
    }
    
    // Late initialization - don't override user-specified values
    r.LateInitializer = config.LateInitializer{
      IgnoredFields: []string{"config"},
    }
  })
  
  p.AddResourceConfigurator("confluent_kafka_cluster", func(r *config.Resource) {
    r.ShortGroup = "kafka"
    
    // Reference to Environment
    r.References["environment.id"] = config.Reference{
      Type: "github.com/michielvha/provider-confluent/apis/environment/v1alpha1.Environment",
    }
    
    // Mark sensitive outputs
    r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
      conn := map[string][]byte{}
      if v, ok := attr["bootstrap_endpoint"].(string); ok {
        conn["bootstrap_endpoint"] = []byte(v)
      }
      if v, ok := attr["rest_endpoint"].(string); ok {
        conn["rest_endpoint"] = []byte(v)
      }
      return conn, nil
    }
  })
  
  p.AddResourceConfigurator("confluent_kafka_acl", func(r *config.Resource) {
    r.ShortGroup = "kafka"
    
    r.References["kafka_cluster.id"] = config.Reference{
      Type: "github.com/michielvha/provider-confluent/apis/kafka/v1alpha1.Cluster",
    }
    
    // Custom extractor for principal from service account
    r.References["principal"] = config.Reference{
      Type:      "github.com/michielvha/provider-confluent/apis/iam/v1alpha1.ServiceAccount",
      Extractor: "github.com/michielvha/provider-confluent/config/common.ExtractPrincipalFromServiceAccount()",
    }
  })
}
```

### Example Custom Configuration: Schema Registry

```go
// config/schema/config.go
package schema

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
  p.AddResourceConfigurator("confluent_schema_registry_cluster", func(r *config.Resource) {
    r.ShortGroup = "schema"
    r.Kind = "RegistryCluster"
    
    r.References["environment.id"] = config.Reference{
      Type: "github.com/michielvha/provider-confluent/apis/environment/v1alpha1.Environment",
    }
    
    // Expose endpoints in connection details
    r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
      conn := map[string][]byte{}
      if v, ok := attr["rest_endpoint"].(string); ok {
        conn["rest_endpoint"] = []byte(v)
      }
      return conn, nil
    }
  })
  
  p.AddResourceConfigurator("confluent_schema", func(r *config.Resource) {
    r.ShortGroup = "schema"
    
    r.References["schema_registry_cluster.id"] = config.Reference{
      Type: "github.com/michielvha/provider-confluent/apis/schema/v1alpha1.RegistryCluster",
    }
    
    // Credentials reference
    r.References["credentials.key"] = config.Reference{
      Type: "github.com/michielvha/provider-confluent/apis/iam/v1alpha1.ApiKey",
      Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractParamPath(\"id\",false)",
    }
    r.References["credentials.secret"] = config.Reference{
      Type: "github.com/michielvha/provider-confluent/apis/iam/v1alpha1.ApiKey",
      Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractParamPath(\"secret\",true)",
    }
  })
}
```

## Testing Strategy

### Unit Tests

- Custom configurator logic
- External name generators
- Field extractors and validators

### Integration Tests with Uptest

```yaml
# examples-generated/kafka/topic.yaml
apiVersion: kafka.confluent.crossplane.io/v1alpha1
kind: Topic
metadata:
  name: orders-topic
  annotations:
    uptest.upbound.io/timeout: "600"
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

### Test Scenarios

1. **Basic CRUD**: Create, Read, Update, Delete for all resources
2. **Cross-resource references**: Topic → Cluster, Schema → Registry
3. **Multi-cluster**: Resources across different Kafka clusters
4. **Security**: ACLs, API keys, RBAC
5. **Networking**: Private Link, Peering
6. **Import**: Import existing Confluent resources

## Security Considerations

### Credential Management

- Store API keys in Kubernetes Secrets
- Support for external secret stores (e.g., Vault, AWS Secrets Manager)
- Rotation strategy for API keys
- Least privilege principle for service accounts

### Sensitive Data Handling

```go
// Mark fields as sensitive
r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
  conn := map[string][]byte{}
  // Extract sensitive connection details
  return conn, nil
}
```

### RBAC Integration

- Namespace isolation for multi-tenant scenarios
- Integration with Kubernetes RBAC
- Confluent RBAC role binding support

## Migration Path from Community Provider

### Compatibility Considerations

1. **API Group Changes**: Community provider uses different API groups
2. **Resource Naming**: May need to align with new conventions
3. **Field Mappings**: Verify all fields are compatible

### Migration Steps

1. Export existing resources as YAML
2. Transform to new provider format
3. Delete old resources (with `--cascade=orphan` to keep Confluent resources)
4. Apply new resource definitions
5. Verify synchronization

## Monitoring & Observability

### Metrics

- Resource reconciliation times
- API call rates to Confluent
- Error rates by resource type
- Resource count by type

### Logging

- Structured logging for debugging
- Terraform provider debug logs (when needed)
- Audit trail for resource changes

### Alerts

- Failed reconciliations
- API quota warnings
- Certificate expiration (for Private Link)

## Future Enhancements

### Post-v1.0

1. **Composite Resources (XRDs)**
   - Complete Kafka application stack
   - Multi-region cluster setup
   - Disaster recovery configurations

2. **Advanced Features**
   - Automated topic lifecycle policies
   - Schema evolution automation
   - Cost optimization recommendations

3. **Integration**
   - ArgoCD integration examples
   - Flux CD compatibility
   - Backstage catalog integration

4. **Observability**
   - Metrics export to Prometheus
   - Grafana dashboards
   - Alert manager integration

## References

- [Upjet Documentation](https://github.com/crossplane/upjet)
- [Terraform Confluent Provider](https://registry.terraform.io/providers/confluentinc/confluent/latest/docs)
- [Confluent Cloud API](https://docs.confluent.io/cloud/current/api.html)
- [Crossplane Documentation](https://docs.crossplane.io)
- [Upjet Provider Template](https://github.com/crossplane/upjet-provider-template)

## Contributors

- Michiel van Heyst (@michielvha)

## License

Apache-2.0

---

**Document Version**: 1.0  
**Last Updated**: November 7, 2025  
**Status**: Draft
