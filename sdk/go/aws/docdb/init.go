// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"fmt"
	"os"

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
	case "aws:docdb/cluster:Cluster":
		r = &Cluster{}
	case "aws:docdb/clusterInstance:ClusterInstance":
		r = &ClusterInstance{}
	case "aws:docdb/clusterParameterGroup:ClusterParameterGroup":
		r = &ClusterParameterGroup{}
	case "aws:docdb/clusterSnapshot:ClusterSnapshot":
		r = &ClusterSnapshot{}
	case "aws:docdb/subnetGroup:SubnetGroup":
		r = &SubnetGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"docdb/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"docdb/clusterInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"docdb/clusterParameterGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"docdb/clusterSnapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"docdb/subnetGroup",
		&module{version},
	)
}
