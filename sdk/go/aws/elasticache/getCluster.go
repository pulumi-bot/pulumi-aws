// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about an Elasticache Cluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticache"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elasticache.LookupCluster(ctx, &elasticache.LookupClusterArgs{
// 			ClusterId: "my-cluster-id",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("aws:elasticache/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// Group identifier.
	ClusterId string `pulumi:"clusterId"`
	// The tags assigned to the resource
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	Arn string `pulumi:"arn"`
	// The Availability Zone for the cache cluster.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// List of node objects including `id`, `address`, `port` and `availabilityZone`.
	// Referenceable e.g. as `${data.aws_elasticache_cluster.bar.cache_nodes.0.address}`
	CacheNodes []GetClusterCacheNode `pulumi:"cacheNodes"`
	// (Memcached only) The DNS name of the cache cluster without the port appended.
	ClusterAddress string `pulumi:"clusterAddress"`
	ClusterId      string `pulumi:"clusterId"`
	// (Memcached only) The configuration endpoint to allow host discovery.
	ConfigurationEndpoint string `pulumi:"configurationEndpoint"`
	// Name of the cache engine.
	Engine string `pulumi:"engine"`
	// Version number of the cache engine.
	EngineVersion string `pulumi:"engineVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed.
	MaintenanceWindow string `pulumi:"maintenanceWindow"`
	// The cluster node type.
	NodeType string `pulumi:"nodeType"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic that ElastiCache notifications get sent to.
	NotificationTopicArn string `pulumi:"notificationTopicArn"`
	// The number of cache nodes that the cache cluster has.
	NumCacheNodes int `pulumi:"numCacheNodes"`
	// Name of the parameter group associated with this cache cluster.
	ParameterGroupName string `pulumi:"parameterGroupName"`
	// The port number on which each of the cache nodes will
	// accept connections.
	Port int `pulumi:"port"`
	// The replication group to which this cache cluster belongs.
	ReplicationGroupId string `pulumi:"replicationGroupId"`
	// List VPC security groups associated with the cache cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// List of security group names associated with this cache cluster.
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them.
	SnapshotRetentionLimit int `pulumi:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of the cache cluster.
	SnapshotWindow string `pulumi:"snapshotWindow"`
	// Name of the subnet group associated to the cache cluster.
	SubnetGroupName string `pulumi:"subnetGroupName"`
	// The tags assigned to the resource
	Tags map[string]string `pulumi:"tags"`
}
