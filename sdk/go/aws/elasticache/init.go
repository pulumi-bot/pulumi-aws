// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:elasticache/cluster:Cluster":
		r = &Cluster{}
	case "aws:elasticache/globalReplicationGroup:GlobalReplicationGroup":
		r = &GlobalReplicationGroup{}
	case "aws:elasticache/parameterGroup:ParameterGroup":
		r = &ParameterGroup{}
	case "aws:elasticache/replicationGroup:ReplicationGroup":
		r = &ReplicationGroup{}
	case "aws:elasticache/securityGroup:SecurityGroup":
		r = &SecurityGroup{}
	case "aws:elasticache/subnetGroup:SubnetGroup":
		r = &SubnetGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version := aws.PkgVersion()
	pulumi.RegisterResourceModule(
		"aws",
		"elasticache/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticache/globalReplicationGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticache/parameterGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticache/replicationGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticache/securityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticache/subnetGroup",
		&module{version},
	)
}
