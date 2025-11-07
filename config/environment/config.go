package environment
package environment

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the environment resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_environment", func(r *config.Resource) {
		r.ShortGroup = "environment"
	})
}
