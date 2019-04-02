// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudhsmv2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a CloudHSM v2 cluster
func LookupCluster(ctx *pulumi.Context, args *GetClusterArgs) (*GetClusterResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["clusterId"] = args.ClusterId
		inputs["clusterState"] = args.ClusterState
	}
	outputs, err := ctx.Invoke("aws:cloudhsmv2/getCluster:getCluster", inputs)
	if err != nil {
		return nil, err
	}
	return &GetClusterResult{
		ClusterCertificates: outputs["clusterCertificates"],
		ClusterId: outputs["clusterId"],
		ClusterState: outputs["clusterState"],
		SecurityGroupId: outputs["securityGroupId"],
		SubnetIds: outputs["subnetIds"],
		VpcId: outputs["vpcId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getCluster.
type GetClusterArgs struct {
	// The id of Cloud HSM v2 cluster.
	ClusterId interface{}
	// The state of the cluster to be found.
	ClusterState interface{}
}

// A collection of values returned by getCluster.
type GetClusterResult struct {
	// The list of cluster certificates.
	// * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
	// * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in UNINITIALIZED state.
	// * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
	// * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
	// * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
	// The number of available cluster certificates may vary depending on state of the cluster.
	ClusterCertificates interface{}
	ClusterId interface{}
	ClusterState interface{}
	// The ID of the security group associated with the CloudHSM cluster.
	SecurityGroupId interface{}
	// The IDs of subnets in which cluster operates.
	SubnetIds interface{}
	// The id of the VPC that the CloudHSM cluster resides in.
	VpcId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
