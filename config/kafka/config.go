package kafka

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the kafka resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_kafka_cluster", func(r *config.Resource) {
		r.ShortGroup = "kafka"
		r.Kind = "Cluster"

		// Reference to Environment
		r.References["environment.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/environment/v1alpha1.Environment",
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

	p.AddResourceConfigurator("confluent_kafka_topic", func(r *config.Resource) {
		r.ShortGroup = "kafka"
		r.Kind = "Topic"

		// Reference to Kafka Cluster
		r.References["kafka_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/kafka/v1alpha1.Cluster",
		}

		// Sensitive fields
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["topic_name"].(string); ok {
				conn["topic_name"] = []byte(v)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("confluent_kafka_acl", func(r *config.Resource) {
		r.ShortGroup = "kafka"
		r.Kind = "ACL"

		r.References["kafka_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/kafka/v1alpha1.Cluster",
		}
	})

	p.AddResourceConfigurator("confluent_kafka_cluster_config", func(r *config.Resource) {
		r.ShortGroup = "kafka"
		r.Kind = "ClusterConfig"

		r.References["kafka_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/kafka/v1alpha1.Cluster",
		}
	})

	p.AddResourceConfigurator("confluent_kafka_client_quota", func(r *config.Resource) {
		r.ShortGroup = "kafka"
		r.Kind = "ClientQuota"

		r.References["kafka_cluster.id"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/kafka/v1alpha1.Cluster",
		}
	})

	p.AddResourceConfigurator("confluent_cluster_link", func(r *config.Resource) {
		r.ShortGroup = "kafka"
		r.Kind = "ClusterLink"
	})

	p.AddResourceConfigurator("confluent_kafka_mirror_topic", func(r *config.Resource) {
		r.ShortGroup = "kafka"
		r.Kind = "MirrorTopic"

		r.References["cluster_link.link_name"] = config.Reference{
			Type: "github.com/michielvha/upjet-provider-confluent/apis/cluster/kafka/v1alpha1.ClusterLink",
		}
	})
}
