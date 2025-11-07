// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connector "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/confluent/connector"
	network "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/confluent/network"
	peering "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/confluent/peering"
	connectorplugin "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/custom/connectorplugin"
	record "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/dns/record"
	environment "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/environment/environment"
	computepool "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/flink/computepool"
	statement "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/flink/statement"
	apikey "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/iam/apikey"
	identitypool "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/iam/identitypool"
	identityprovider "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/iam/identityprovider"
	invitation "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/iam/invitation"
	rolebinding "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/iam/rolebinding"
	serviceaccount "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/iam/serviceaccount"
	acl "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/kafka/acl"
	clientquota "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/kafka/clientquota"
	cluster "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/kafka/cluster"
	clusterconfig "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/kafka/clusterconfig"
	clusterlink "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/kafka/clusterlink"
	mirrortopic "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/kafka/mirrortopic"
	topic "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/kafka/topic"
	clusterksql "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/ksql/cluster"
	linkendpoint "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/network/linkendpoint"
	linkservice "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/network/linkservice"
	linkaccess "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/private/linkaccess"
	providerconfig "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/providerconfig"
	businessmetadata "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/businessmetadata"
	businessmetadatabinding "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/businessmetadatabinding"
	registryclusterconfig "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/registryclusterconfig"
	registryclustermode "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/registryclustermode"
	registrydek "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/registrydek"
	registrykek "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/registrykek"
	schema "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/schema"
	subjectconfig "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/subjectconfig"
	subjectmode "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/subjectmode"
	tag "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/tag"
	tagbinding "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/schema/tagbinding"
	gatewayattachment "github.com/michielvha/upjet-provider-confluent/internal/controller/namespaced/transit/gatewayattachment"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connector.Setup,
		network.Setup,
		peering.Setup,
		connectorplugin.Setup,
		record.Setup,
		environment.Setup,
		computepool.Setup,
		statement.Setup,
		apikey.Setup,
		identitypool.Setup,
		identityprovider.Setup,
		invitation.Setup,
		rolebinding.Setup,
		serviceaccount.Setup,
		acl.Setup,
		clientquota.Setup,
		cluster.Setup,
		clusterconfig.Setup,
		clusterlink.Setup,
		mirrortopic.Setup,
		topic.Setup,
		clusterksql.Setup,
		linkendpoint.Setup,
		linkservice.Setup,
		linkaccess.Setup,
		providerconfig.Setup,
		businessmetadata.Setup,
		businessmetadatabinding.Setup,
		registryclusterconfig.Setup,
		registryclustermode.Setup,
		registrydek.Setup,
		registrykek.Setup,
		schema.Setup,
		subjectconfig.Setup,
		subjectmode.Setup,
		tag.Setup,
		tagbinding.Setup,
		gatewayattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connector.SetupGated,
		network.SetupGated,
		peering.SetupGated,
		connectorplugin.SetupGated,
		record.SetupGated,
		environment.SetupGated,
		computepool.SetupGated,
		statement.SetupGated,
		apikey.SetupGated,
		identitypool.SetupGated,
		identityprovider.SetupGated,
		invitation.SetupGated,
		rolebinding.SetupGated,
		serviceaccount.SetupGated,
		acl.SetupGated,
		clientquota.SetupGated,
		cluster.SetupGated,
		clusterconfig.SetupGated,
		clusterlink.SetupGated,
		mirrortopic.SetupGated,
		topic.SetupGated,
		clusterksql.SetupGated,
		linkendpoint.SetupGated,
		linkservice.SetupGated,
		linkaccess.SetupGated,
		providerconfig.SetupGated,
		businessmetadata.SetupGated,
		businessmetadatabinding.SetupGated,
		registryclusterconfig.SetupGated,
		registryclustermode.SetupGated,
		registrydek.SetupGated,
		registrykek.SetupGated,
		schema.SetupGated,
		subjectconfig.SetupGated,
		subjectmode.SetupGated,
		tag.SetupGated,
		tagbinding.SetupGated,
		gatewayattachment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
