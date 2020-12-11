// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an RDS Cluster Instance Resource. A Cluster Instance Resource defines
// attributes that are specific to a single instance in a [RDS Cluster](https://www.terraform.io/docs/providers/aws/r/rds_cluster.html),
// specifically running Amazon Aurora.
//
// Unlike other RDS resources that support replication, with Amazon Aurora you do
// not designate a primary and subsequent replicas. Instead, you simply add RDS
// Instances and Aurora manages the replication. You can use the [count](https://www.terraform.io/docs/configuration/resources.html#count)
// meta-parameter to make multiple instances and join them all to the same RDS
// Cluster, or you may specify different Cluster Instance resources with various
// `instanceClass` sizes.
//
// For more information on Amazon Aurora, see [Aurora on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Aurora.html) in the Amazon RDS User Guide.
//
// > **NOTE:** Deletion Protection from the RDS service can only be enabled at the cluster level, not for individual cluster instances. You can still add the [`protect` CustomResourceOption](https://www.pulumi.com/docs/intro/concepts/programming-model/#protect) to this resource configuration if you desire protection from accidental deletion.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := rds.NewCluster(ctx, "_default", &rds.ClusterArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-west-2a"),
// 				pulumi.String("us-west-2b"),
// 				pulumi.String("us-west-2c"),
// 			},
// 			DatabaseName:   pulumi.String("mydb"),
// 			MasterUsername: pulumi.String("foo"),
// 			MasterPassword: pulumi.String("barbut8chars"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		var clusterInstances []*rds.ClusterInstance
// 		for key0, val0 := range 2 {
// 			__res, err := rds.NewClusterInstance(ctx, fmt.Sprintf("clusterInstances-%v", key0), &rds.ClusterInstanceArgs{
// 				Identifier:        pulumi.String(fmt.Sprintf("%v%v", "aurora-cluster-demo-", val0)),
// 				ClusterIdentifier: _default.ID(),
// 				InstanceClass:     pulumi.String("db.r4.large"),
// 				Engine:            _default.Engine,
// 				EngineVersion:     _default.EngineVersion,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 			clusterInstances = append(clusterInstances, __res)
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// RDS Cluster Instances can be imported using the `identifier`, e.g.
//
// ```sh
//  $ pulumi import aws:rds/clusterInstance:ClusterInstance prod_instance_1 aurora-cluster-instance-1
// ```
type ClusterInstance struct {
	pulumi.CustomResourceState

	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolOutput `pulumi:"applyImmediately"`
	// Amazon Resource Name (ARN) of cluster instance
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade pulumi.BoolPtrOutput `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The identifier of the CA certificate for the DB instance.
	CaCertIdentifier pulumi.StringOutput `pulumi:"caCertIdentifier"`
	// The identifier of the `rds.Cluster` in which to launch this instance.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
	CopyTagsToSnapshot pulumi.BoolPtrOutput `pulumi:"copyTagsToSnapshot"`
	// The name of the DB parameter group to associate with this instance.
	DbParameterGroupName pulumi.StringOutput `pulumi:"dbParameterGroupName"`
	// A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` of the attached `rds.Cluster`.
	DbSubnetGroupName pulumi.StringOutput `pulumi:"dbSubnetGroupName"`
	// The region-unique, immutable identifier for the DB instance.
	DbiResourceId pulumi.StringOutput `pulumi:"dbiResourceId"`
	// The DNS address for this instance. May not be writable
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
	// For information on the difference between the available Aurora MySQL engines
	// see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
	// in the Amazon RDS User Guide.
	Engine pulumi.StringPtrOutput `pulumi:"engine"`
	// The database engine version
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringOutput `pulumi:"identifierPrefix"`
	// The instance class to use. For details on CPU
	// and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
	InstanceClass pulumi.StringOutput `pulumi:"instanceClass"`
	// The ARN for the KMS encryption key if one is set to the cluster.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	MonitoringInterval pulumi.IntPtrOutput `pulumi:"monitoringInterval"`
	// The ARN for the IAM role that permits RDS to send
	// enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
	// what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	MonitoringRoleArn pulumi.StringOutput `pulumi:"monitoringRoleArn"`
	// Specifies whether Performance Insights is enabled or not.
	PerformanceInsightsEnabled pulumi.BoolOutput `pulumi:"performanceInsightsEnabled"`
	// The ARN for the KMS key to encrypt Performance Insights data. When specifying `performanceInsightsKmsKeyId`, `performanceInsightsEnabled` needs to be set to true.
	PerformanceInsightsKmsKeyId pulumi.StringOutput `pulumi:"performanceInsightsKmsKeyId"`
	// The database port
	Port pulumi.IntOutput `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled.
	// Eg: "04:00-09:00"
	PreferredBackupWindow pulumi.StringOutput `pulumi:"preferredBackupWindow"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringOutput `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	PromotionTier pulumi.IntPtrOutput `pulumi:"promotionTier"`
	// Bool to control if instance is publicly accessible.
	// Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
	// details on controlling this property.
	PubliclyAccessible pulumi.BoolPtrOutput `pulumi:"publiclyAccessible"`
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted pulumi.BoolOutput `pulumi:"storageEncrypted"`
	// A map of tags to assign to the instance.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer pulumi.BoolOutput `pulumi:"writer"`
}

// NewClusterInstance registers a new resource with the given unique name, arguments, and options.
func NewClusterInstance(ctx *pulumi.Context,
	name string, args *ClusterInstanceArgs, opts ...pulumi.ResourceOption) (*ClusterInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'ClusterIdentifier'")
	}
	if args.InstanceClass == nil {
		return nil, errors.New("invalid value for required argument 'InstanceClass'")
	}
	var resource ClusterInstance
	err := ctx.RegisterResource("aws:rds/clusterInstance:ClusterInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterInstance gets an existing ClusterInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterInstanceState, opts ...pulumi.ResourceOption) (*ClusterInstance, error) {
	var resource ClusterInstance
	err := ctx.ReadResource("aws:rds/clusterInstance:ClusterInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterInstance resources.
type clusterInstanceState struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// Amazon Resource Name (ARN) of cluster instance
	Arn *string `pulumi:"arn"`
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The identifier of the CA certificate for the DB instance.
	CaCertIdentifier *string `pulumi:"caCertIdentifier"`
	// The identifier of the `rds.Cluster` in which to launch this instance.
	ClusterIdentifier *string `pulumi:"clusterIdentifier"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
	CopyTagsToSnapshot *bool `pulumi:"copyTagsToSnapshot"`
	// The name of the DB parameter group to associate with this instance.
	DbParameterGroupName *string `pulumi:"dbParameterGroupName"`
	// A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` of the attached `rds.Cluster`.
	DbSubnetGroupName *string `pulumi:"dbSubnetGroupName"`
	// The region-unique, immutable identifier for the DB instance.
	DbiResourceId *string `pulumi:"dbiResourceId"`
	// The DNS address for this instance. May not be writable
	Endpoint *string `pulumi:"endpoint"`
	// The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
	// For information on the difference between the available Aurora MySQL engines
	// see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
	// in the Amazon RDS User Guide.
	Engine *string `pulumi:"engine"`
	// The database engine version
	EngineVersion *string `pulumi:"engineVersion"`
	// The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
	Identifier *string `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix *string `pulumi:"identifierPrefix"`
	// The instance class to use. For details on CPU
	// and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
	InstanceClass *string `pulumi:"instanceClass"`
	// The ARN for the KMS encryption key if one is set to the cluster.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	MonitoringInterval *int `pulumi:"monitoringInterval"`
	// The ARN for the IAM role that permits RDS to send
	// enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
	// what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	MonitoringRoleArn *string `pulumi:"monitoringRoleArn"`
	// Specifies whether Performance Insights is enabled or not.
	PerformanceInsightsEnabled *bool `pulumi:"performanceInsightsEnabled"`
	// The ARN for the KMS key to encrypt Performance Insights data. When specifying `performanceInsightsKmsKeyId`, `performanceInsightsEnabled` needs to be set to true.
	PerformanceInsightsKmsKeyId *string `pulumi:"performanceInsightsKmsKeyId"`
	// The database port
	Port *int `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled.
	// Eg: "04:00-09:00"
	PreferredBackupWindow *string `pulumi:"preferredBackupWindow"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	PromotionTier *int `pulumi:"promotionTier"`
	// Bool to control if instance is publicly accessible.
	// Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
	// details on controlling this property.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
	// A map of tags to assign to the instance.
	Tags map[string]string `pulumi:"tags"`
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer *bool `pulumi:"writer"`
}

type ClusterInstanceState struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of cluster instance
	Arn pulumi.StringPtrInput
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
	AvailabilityZone pulumi.StringPtrInput
	// The identifier of the CA certificate for the DB instance.
	CaCertIdentifier pulumi.StringPtrInput
	// The identifier of the `rds.Cluster` in which to launch this instance.
	ClusterIdentifier pulumi.StringPtrInput
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
	CopyTagsToSnapshot pulumi.BoolPtrInput
	// The name of the DB parameter group to associate with this instance.
	DbParameterGroupName pulumi.StringPtrInput
	// A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` of the attached `rds.Cluster`.
	DbSubnetGroupName pulumi.StringPtrInput
	// The region-unique, immutable identifier for the DB instance.
	DbiResourceId pulumi.StringPtrInput
	// The DNS address for this instance. May not be writable
	Endpoint pulumi.StringPtrInput
	// The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
	// For information on the difference between the available Aurora MySQL engines
	// see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
	// in the Amazon RDS User Guide.
	Engine pulumi.StringPtrInput
	// The database engine version
	EngineVersion pulumi.StringPtrInput
	// The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringPtrInput
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringPtrInput
	// The instance class to use. For details on CPU
	// and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
	InstanceClass pulumi.StringPtrInput
	// The ARN for the KMS encryption key if one is set to the cluster.
	KmsKeyId pulumi.StringPtrInput
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	MonitoringInterval pulumi.IntPtrInput
	// The ARN for the IAM role that permits RDS to send
	// enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
	// what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	MonitoringRoleArn pulumi.StringPtrInput
	// Specifies whether Performance Insights is enabled or not.
	PerformanceInsightsEnabled pulumi.BoolPtrInput
	// The ARN for the KMS key to encrypt Performance Insights data. When specifying `performanceInsightsKmsKeyId`, `performanceInsightsEnabled` needs to be set to true.
	PerformanceInsightsKmsKeyId pulumi.StringPtrInput
	// The database port
	Port pulumi.IntPtrInput
	// The daily time range during which automated backups are created if automated backups are enabled.
	// Eg: "04:00-09:00"
	PreferredBackupWindow pulumi.StringPtrInput
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	PromotionTier pulumi.IntPtrInput
	// Bool to control if instance is publicly accessible.
	// Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
	// details on controlling this property.
	PubliclyAccessible pulumi.BoolPtrInput
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted pulumi.BoolPtrInput
	// A map of tags to assign to the instance.
	Tags pulumi.StringMapInput
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer pulumi.BoolPtrInput
}

func (ClusterInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterInstanceState)(nil)).Elem()
}

type clusterInstanceArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The identifier of the CA certificate for the DB instance.
	CaCertIdentifier *string `pulumi:"caCertIdentifier"`
	// The identifier of the `rds.Cluster` in which to launch this instance.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
	CopyTagsToSnapshot *bool `pulumi:"copyTagsToSnapshot"`
	// The name of the DB parameter group to associate with this instance.
	DbParameterGroupName *string `pulumi:"dbParameterGroupName"`
	// A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` of the attached `rds.Cluster`.
	DbSubnetGroupName *string `pulumi:"dbSubnetGroupName"`
	// The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
	// For information on the difference between the available Aurora MySQL engines
	// see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
	// in the Amazon RDS User Guide.
	Engine *string `pulumi:"engine"`
	// The database engine version
	EngineVersion *string `pulumi:"engineVersion"`
	// The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
	Identifier *string `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix *string `pulumi:"identifierPrefix"`
	// The instance class to use. For details on CPU
	// and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
	InstanceClass string `pulumi:"instanceClass"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	MonitoringInterval *int `pulumi:"monitoringInterval"`
	// The ARN for the IAM role that permits RDS to send
	// enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
	// what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	MonitoringRoleArn *string `pulumi:"monitoringRoleArn"`
	// Specifies whether Performance Insights is enabled or not.
	PerformanceInsightsEnabled *bool `pulumi:"performanceInsightsEnabled"`
	// The ARN for the KMS key to encrypt Performance Insights data. When specifying `performanceInsightsKmsKeyId`, `performanceInsightsEnabled` needs to be set to true.
	PerformanceInsightsKmsKeyId *string `pulumi:"performanceInsightsKmsKeyId"`
	// The daily time range during which automated backups are created if automated backups are enabled.
	// Eg: "04:00-09:00"
	PreferredBackupWindow *string `pulumi:"preferredBackupWindow"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	PromotionTier *int `pulumi:"promotionTier"`
	// Bool to control if instance is publicly accessible.
	// Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
	// details on controlling this property.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// A map of tags to assign to the instance.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterInstance resource.
type ClusterInstanceArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolPtrInput
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
	AvailabilityZone pulumi.StringPtrInput
	// The identifier of the CA certificate for the DB instance.
	CaCertIdentifier pulumi.StringPtrInput
	// The identifier of the `rds.Cluster` in which to launch this instance.
	ClusterIdentifier pulumi.StringInput
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
	CopyTagsToSnapshot pulumi.BoolPtrInput
	// The name of the DB parameter group to associate with this instance.
	DbParameterGroupName pulumi.StringPtrInput
	// A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` of the attached `rds.Cluster`.
	DbSubnetGroupName pulumi.StringPtrInput
	// The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
	// For information on the difference between the available Aurora MySQL engines
	// see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
	// in the Amazon RDS User Guide.
	Engine pulumi.StringPtrInput
	// The database engine version
	EngineVersion pulumi.StringPtrInput
	// The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringPtrInput
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringPtrInput
	// The instance class to use. For details on CPU
	// and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
	InstanceClass pulumi.StringInput
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	MonitoringInterval pulumi.IntPtrInput
	// The ARN for the IAM role that permits RDS to send
	// enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
	// what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	MonitoringRoleArn pulumi.StringPtrInput
	// Specifies whether Performance Insights is enabled or not.
	PerformanceInsightsEnabled pulumi.BoolPtrInput
	// The ARN for the KMS key to encrypt Performance Insights data. When specifying `performanceInsightsKmsKeyId`, `performanceInsightsEnabled` needs to be set to true.
	PerformanceInsightsKmsKeyId pulumi.StringPtrInput
	// The daily time range during which automated backups are created if automated backups are enabled.
	// Eg: "04:00-09:00"
	PreferredBackupWindow pulumi.StringPtrInput
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	PromotionTier pulumi.IntPtrInput
	// Bool to control if instance is publicly accessible.
	// Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
	// details on controlling this property.
	PubliclyAccessible pulumi.BoolPtrInput
	// A map of tags to assign to the instance.
	Tags pulumi.StringMapInput
}

func (ClusterInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterInstanceArgs)(nil)).Elem()
}

type ClusterInstanceInput interface {
	pulumi.Input

	ToClusterInstanceOutput() ClusterInstanceOutput
	ToClusterInstanceOutputWithContext(ctx context.Context) ClusterInstanceOutput
}

type ClusterInstancePtrInput interface {
	pulumi.Input

	ToClusterInstancePtrOutput() ClusterInstancePtrOutput
	ToClusterInstancePtrOutputWithContext(ctx context.Context) ClusterInstancePtrOutput
}

func (ClusterInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInstance)(nil)).Elem()
}

func (i ClusterInstance) ToClusterInstanceOutput() ClusterInstanceOutput {
	return i.ToClusterInstanceOutputWithContext(context.Background())
}

func (i ClusterInstance) ToClusterInstanceOutputWithContext(ctx context.Context) ClusterInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInstanceOutput)
}

func (i ClusterInstance) ToClusterInstancePtrOutput() ClusterInstancePtrOutput {
	return i.ToClusterInstancePtrOutputWithContext(context.Background())
}

func (i ClusterInstance) ToClusterInstancePtrOutputWithContext(ctx context.Context) ClusterInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInstancePtrOutput)
}

type ClusterInstanceOutput struct {
	*pulumi.OutputState
}

func (ClusterInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInstanceOutput)(nil)).Elem()
}

func (o ClusterInstanceOutput) ToClusterInstanceOutput() ClusterInstanceOutput {
	return o
}

func (o ClusterInstanceOutput) ToClusterInstanceOutputWithContext(ctx context.Context) ClusterInstanceOutput {
	return o
}

type ClusterInstancePtrOutput struct {
	*pulumi.OutputState
}

func (ClusterInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInstance)(nil)).Elem()
}

func (o ClusterInstancePtrOutput) ToClusterInstancePtrOutput() ClusterInstancePtrOutput {
	return o
}

func (o ClusterInstancePtrOutput) ToClusterInstancePtrOutputWithContext(ctx context.Context) ClusterInstancePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterInstanceOutput{})
	pulumi.RegisterOutputType(ClusterInstancePtrOutput{})
}
