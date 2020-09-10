// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudhsmv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	ClusterCertificates    ClusterClusterCertificateArrayOutput `pulumi:"clusterCertificates"`
	ClusterId              pulumi.StringOutput                  `pulumi:"clusterId"`
	ClusterState           pulumi.StringOutput                  `pulumi:"clusterState"`
	HsmType                pulumi.StringOutput                  `pulumi:"hsmType"`
	SecurityGroupId        pulumi.StringOutput                  `pulumi:"securityGroupId"`
	SourceBackupIdentifier pulumi.StringPtrOutput               `pulumi:"sourceBackupIdentifier"`
	SubnetIds              pulumi.StringArrayOutput             `pulumi:"subnetIds"`
	Tags                   pulumi.StringMapOutput               `pulumi:"tags"`
	VpcId                  pulumi.StringOutput                  `pulumi:"vpcId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.HsmType == nil {
		return nil, errors.New("missing required argument 'HsmType'")
	}
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("aws:cloudhsmv2/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("aws:cloudhsmv2/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	ClusterCertificates    []ClusterClusterCertificate `pulumi:"clusterCertificates"`
	ClusterId              *string                     `pulumi:"clusterId"`
	ClusterState           *string                     `pulumi:"clusterState"`
	HsmType                *string                     `pulumi:"hsmType"`
	SecurityGroupId        *string                     `pulumi:"securityGroupId"`
	SourceBackupIdentifier *string                     `pulumi:"sourceBackupIdentifier"`
	SubnetIds              []string                    `pulumi:"subnetIds"`
	Tags                   map[string]string           `pulumi:"tags"`
	VpcId                  *string                     `pulumi:"vpcId"`
}

type ClusterState struct {
	ClusterCertificates    ClusterClusterCertificateArrayInput
	ClusterId              pulumi.StringPtrInput
	ClusterState           pulumi.StringPtrInput
	HsmType                pulumi.StringPtrInput
	SecurityGroupId        pulumi.StringPtrInput
	SourceBackupIdentifier pulumi.StringPtrInput
	SubnetIds              pulumi.StringArrayInput
	Tags                   pulumi.StringMapInput
	VpcId                  pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	HsmType                string            `pulumi:"hsmType"`
	SourceBackupIdentifier *string           `pulumi:"sourceBackupIdentifier"`
	SubnetIds              []string          `pulumi:"subnetIds"`
	Tags                   map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	HsmType                pulumi.StringInput
	SourceBackupIdentifier pulumi.StringPtrInput
	SubnetIds              pulumi.StringArrayInput
	Tags                   pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
