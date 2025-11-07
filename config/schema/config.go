package schema
package schema

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the schema resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_schema_registry_cluster", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "RegistryCluster"
		
		r.References["environment.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/environment/v1alpha1.Environment",
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
		r.Kind = "Schema"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
	
	p.AddResourceConfigurator("confluent_subject_mode", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "SubjectMode"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
	
	p.AddResourceConfigurator("confluent_subject_config", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "SubjectConfig"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
	
	p.AddResourceConfigurator("confluent_schema_registry_kek", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "RegistryKek"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
	
	p.AddResourceConfigurator("confluent_schema_registry_dek", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "RegistryDek"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
		
		r.References["kek_name"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryKek",
		}
	})
	
	p.AddResourceConfigurator("confluent_schema_registry_cluster_mode", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "RegistryClusterMode"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
	
	p.AddResourceConfigurator("confluent_schema_registry_cluster_config", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "RegistryClusterConfig"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
	
	p.AddResourceConfigurator("confluent_business_metadata", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "BusinessMetadata"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
	
	p.AddResourceConfigurator("confluent_business_metadata_binding", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "BusinessMetadataBinding"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
	
	p.AddResourceConfigurator("confluent_tag", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "Tag"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
	
	p.AddResourceConfigurator("confluent_tag_binding", func(r *config.Resource) {
		r.ShortGroup = "schema"
		r.Kind = "TagBinding"
		
		r.References["schema_registry_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/schema/v1alpha1.RegistryCluster",
		}
	})
}
