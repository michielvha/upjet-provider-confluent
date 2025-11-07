package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Core Infrastructure
	"confluent_environment":             config.NameAsIdentifier,
	"confluent_kafka_cluster":           config.IdentifierFromProvider,
	"confluent_ksql_cluster":            config.IdentifierFromProvider,

	// Kafka Resources
	"confluent_kafka_topic": config.NameAsIdentifier,
	"confluent_kafka_acl": config.TemplatedStringAsIdentifier("",
		"{{ .parameters.kafka_cluster.id }}:{{ .parameters.resource_type }}:{{ .parameters.resource_name }}:{{ .parameters.pattern_type }}:{{ .parameters.principal }}:{{ .parameters.host }}:{{ .parameters.operation }}:{{ .parameters.permission }}"),
	"confluent_kafka_cluster_config": config.IdentifierFromProvider,
	"confluent_kafka_client_quota":   config.IdentifierFromProvider,
	"confluent_cluster_link": config.TemplatedStringAsIdentifier("link_name",
		"{{ .parameters.link_name }}"),
	"confluent_kafka_mirror_topic": config.IdentifierFromProvider,

	// Schema Registry
	"confluent_schema": config.TemplatedStringAsIdentifier("subject_name",
		"{{ .parameters.schema_registry_cluster.id }}:{{ .parameters.subject_name }}"),
	"confluent_subject_mode": config.TemplatedStringAsIdentifier("subject_name",
		"{{ .parameters.schema_registry_cluster.id }}:{{ .parameters.subject_name }}"),
	"confluent_subject_config": config.TemplatedStringAsIdentifier("subject_name",
		"{{ .parameters.schema_registry_cluster.id }}:{{ .parameters.subject_name }}"),
	"confluent_schema_registry_kek":            config.NameAsIdentifier,
	"confluent_schema_registry_dek":            config.IdentifierFromProvider,
	"confluent_schema_registry_cluster_mode":   config.IdentifierFromProvider,
	"confluent_schema_registry_cluster_config": config.IdentifierFromProvider,

	// Connect
	"confluent_connector":               config.NameAsIdentifier,
	"confluent_custom_connector_plugin": config.NameAsIdentifier,

	// Flink
	"confluent_flink_compute_pool": config.IdentifierFromProvider,
	"confluent_flink_statement":    config.IdentifierFromProvider,

	// IAM
	"confluent_service_account":   config.IdentifierFromProvider,
	"confluent_api_key":           config.IdentifierFromProvider,
	"confluent_role_binding":      config.IdentifierFromProvider,
	"confluent_identity_provider": config.IdentifierFromProvider,
	"confluent_identity_pool":     config.IdentifierFromProvider,
	"confluent_invitation":        config.IdentifierFromProvider,

	// Networking
	"confluent_network":                    config.IdentifierFromProvider,
	"confluent_peering":                    config.IdentifierFromProvider,
	"confluent_private_link_access":        config.IdentifierFromProvider,
	"confluent_transit_gateway_attachment": config.IdentifierFromProvider,
	"confluent_dns_record":                 config.IdentifierFromProvider,
	"confluent_network_link_service":       config.IdentifierFromProvider,
	"confluent_network_link_endpoint":      config.IdentifierFromProvider,

	// Business Metadata (Schema Registry)
	"confluent_business_metadata":         config.IdentifierFromProvider,
	"confluent_business_metadata_binding": config.IdentifierFromProvider,
	"confluent_tag":                       config.IdentifierFromProvider,
	"confluent_tag_binding":               config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
