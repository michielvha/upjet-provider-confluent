package iam

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the iam resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_service_account", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.Kind = "ServiceAccount"
	})

	p.AddResourceConfigurator("confluent_api_key", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.Kind = "APIKey"

		r.References["owner.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/iam/v1alpha1.ServiceAccount",
		}

		// Mark secret as sensitive
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["id"].(string); ok {
				conn["id"] = []byte(v)
			}
			if v, ok := attr["secret"].(string); ok {
				conn["secret"] = []byte(v)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("confluent_role_binding", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.Kind = "RoleBinding"

		r.References["principal"] = config.Reference{
			Type:      "github.com/michielvha/upjet-provider-confluent/apis/cluster/iam/v1alpha1.ServiceAccount",
			Extractor: `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("id",false)`,
		}
	})

	p.AddResourceConfigurator("confluent_identity_provider", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.Kind = "IdentityProvider"
	})

	p.AddResourceConfigurator("confluent_identity_pool", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.Kind = "IdentityPool"

		r.References["identity_provider.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/iam/v1alpha1.IdentityProvider",
		}
	})

	p.AddResourceConfigurator("confluent_invitation", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.Kind = "Invitation"
	})
}
