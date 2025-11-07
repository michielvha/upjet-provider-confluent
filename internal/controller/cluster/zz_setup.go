// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connector "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/confluent/connector"
	network "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/confluent/network"
	peering "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/confluent/peering"
	connectorplugin "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/custom/connectorplugin"
	record "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/dns/record"
	environment "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/environment/environment"
	computepool "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/flink/computepool"
	statement "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/flink/statement"
	apikey "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/iam/apikey"
	identitypool "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/iam/identitypool"
	identityprovider "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/iam/identityprovider"
	invitation "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/iam/invitation"
	rolebinding "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/iam/rolebinding"
	serviceaccount "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/iam/serviceaccount"
	acl "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/kafka/acl"
	clientquota "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/kafka/clientquota"
	cluster "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/kafka/cluster"
	clusterconfig "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/kafka/clusterconfig"
	clusterlink "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/kafka/clusterlink"
	mirrortopic "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/kafka/mirrortopic"
	topic "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/kafka/topic"
	clusterksql "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/ksql/cluster"
	linkendpoint "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/network/linkendpoint"
	linkservice "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/network/linkservice"
	linkaccess "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/private/linkaccess"
	providerconfig "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/providerconfig"
	businessmetadata "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/businessmetadata"
	businessmetadatabinding "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/businessmetadatabinding"
	registryclusterconfig "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/registryclusterconfig"
	registryclustermode "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/registryclustermode"
	registrydek "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/registrydek"
	registrykek "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/registrykek"
	schema "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/schema"
	subjectconfig "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/subjectconfig"
	subjectmode "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/subjectmode"
	tag "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/tag"
	tagbinding "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/schema/tagbinding"
	gatewayattachment "github.com/michielvha/upjet-provider-confluent/internal/controller/cluster/transit/gatewayattachment"
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
