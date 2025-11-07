package schema

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the schema resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_schema", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "Schema"
	})

	p.AddResourceConfigurator("confluent_subject_mode", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "SubjectMode"
	})

	p.AddResourceConfigurator("confluent_subject_config", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "SubjectConfig"
	})

	p.AddResourceConfigurator("confluent_schema_registry_kek", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "RegistryKek"
	})

	p.AddResourceConfigurator("confluent_schema_registry_dek", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "RegistryDek"

		r.References["kek_name"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/schema/v1alpha1.RegistryKek",
		}
	})

	p.AddResourceConfigurator("confluent_schema_registry_cluster_mode", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "RegistryClusterMode"
	})

	p.AddResourceConfigurator("confluent_schema_registry_cluster_config", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "RegistryClusterConfig"
	})

	p.AddResourceConfigurator("confluent_business_metadata", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "BusinessMetadata"
	})

	p.AddResourceConfigurator("confluent_business_metadata_binding", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "BusinessMetadataBinding"
	})

	p.AddResourceConfigurator("confluent_tag", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "Tag"
	})

	p.AddResourceConfigurator("confluent_tag_binding", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "TagBinding"
	})
}
