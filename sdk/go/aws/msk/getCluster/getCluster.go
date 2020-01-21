// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getCluster

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get information on an Amazon MSK Cluster.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/msk_cluster.html.markdown.
func GetCluster(ctx *pulumi.Context, args *GetClusterArgs, opts ...pulumi.InvokeOption) (*GetClusterResult, error) {
	var rv GetClusterResult
	err := ctx.Invoke("aws:msk/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type GetClusterArgs struct {
	// Name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	Tags map[string]interface{} `pulumi:"tags"`
}


// A collection of values returned by getCluster.
type GetClusterResult struct {
	// Amazon Resource Name (ARN) of the MSK cluster.
	Arn string `pulumi:"arn"`
	// A comma separated list of one or more hostname:port pairs of Kafka brokers suitable to boostrap connectivity to the Kafka cluster.
	BootstrapBrokers string `pulumi:"bootstrapBrokers"`
	// A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster.
	BootstrapBrokersTls string `pulumi:"bootstrapBrokersTls"`
	ClusterName string `pulumi:"clusterName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Apache Kafka version.
	KafkaVersion string `pulumi:"kafkaVersion"`
	// Number of broker nodes in the cluster.
	NumberOfBrokerNodes int `pulumi:"numberOfBrokerNodes"`
	// Map of key-value pairs assigned to the cluster.
	Tags map[string]interface{} `pulumi:"tags"`
	// A comma separated list of one or more IP:port pairs to use to connect to the Apache Zookeeper cluster.
	ZookeeperConnectString string `pulumi:"zookeeperConnectString"`
}

