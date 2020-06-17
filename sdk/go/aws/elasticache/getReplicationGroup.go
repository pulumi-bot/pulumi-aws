// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about an Elasticache Replication Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elasticache.LookupReplicationGroup(ctx, &elasticache.LookupReplicationGroupArgs{
// 			ReplicationGroupId: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupReplicationGroup(ctx *pulumi.Context, args *LookupReplicationGroupArgs, opts ...pulumi.InvokeOption) (*LookupReplicationGroupResult, error) {
	var rv LookupReplicationGroupResult
	err := ctx.Invoke("aws:elasticache/getReplicationGroup:getReplicationGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReplicationGroup.
type LookupReplicationGroupArgs struct {
	// The identifier for the replication group.
	ReplicationGroupId string `pulumi:"replicationGroupId"`
}

// A collection of values returned by getReplicationGroup.
type LookupReplicationGroupResult struct {
	// A flag that enables using an AuthToken (password) when issuing Redis commands.
	AuthTokenEnabled bool `pulumi:"authTokenEnabled"`
	// A flag whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails.
	AutomaticFailoverEnabled bool `pulumi:"automaticFailoverEnabled"`
	// The configuration endpoint address to allow host discovery.
	ConfigurationEndpointAddress string `pulumi:"configurationEndpointAddress"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The identifiers of all the nodes that are part of this replication group.
	MemberClusters []string `pulumi:"memberClusters"`
	// The cluster node type.
	NodeType string `pulumi:"nodeType"`
	// The number of cache clusters that the replication group has.
	NumberCacheClusters int `pulumi:"numberCacheClusters"`
	// The port number on which the configuration endpoint will accept connections.
	Port int `pulumi:"port"`
	// The endpoint of the primary node in this node group (shard).
	PrimaryEndpointAddress string `pulumi:"primaryEndpointAddress"`
	// The description of the replication group.
	ReplicationGroupDescription string `pulumi:"replicationGroupDescription"`
	// The identifier for the replication group.
	ReplicationGroupId string `pulumi:"replicationGroupId"`
	// The number of days for which ElastiCache retains automatic cache cluster snapshots before deleting them.
	SnapshotRetentionLimit int `pulumi:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).
	SnapshotWindow string `pulumi:"snapshotWindow"`
}
