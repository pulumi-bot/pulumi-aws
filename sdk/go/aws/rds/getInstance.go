// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about an RDS instance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := rds.LookupInstance(ctx, &rds.LookupInstanceArgs{
// 			DbInstanceIdentifier: "my-test-database",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("aws:rds/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	// The name of the RDS instance
	DbInstanceIdentifier string            `pulumi:"dbInstanceIdentifier"`
	Tags                 map[string]string `pulumi:"tags"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	// The hostname of the RDS instance. See also `endpoint` and `port`.
	Address string `pulumi:"address"`
	// Specifies the allocated storage size specified in gigabytes.
	AllocatedStorage int `pulumi:"allocatedStorage"`
	// Indicates that minor version patches are applied automatically.
	AutoMinorVersionUpgrade bool `pulumi:"autoMinorVersionUpgrade"`
	// Specifies the name of the Availability Zone the DB instance is located in.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Specifies the number of days for which automatic DB snapshots are retained.
	BackupRetentionPeriod int `pulumi:"backupRetentionPeriod"`
	// Specifies the identifier of the CA certificate for the DB instance.
	CaCertIdentifier string `pulumi:"caCertIdentifier"`
	// If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.
	DbClusterIdentifier string `pulumi:"dbClusterIdentifier"`
	// The Amazon Resource Name (ARN) for the DB instance.
	DbInstanceArn string `pulumi:"dbInstanceArn"`
	// Contains the name of the compute and memory capacity class of the DB instance.
	DbInstanceClass      string `pulumi:"dbInstanceClass"`
	DbInstanceIdentifier string `pulumi:"dbInstanceIdentifier"`
	// Specifies the port that the DB instance listens on.
	DbInstancePort int `pulumi:"dbInstancePort"`
	// Contains the name of the initial database of this instance that was provided at create time, if one was specified when the DB instance was created. This same name is returned for the life of the DB instance.
	DbName string `pulumi:"dbName"`
	// Provides the list of DB parameter groups applied to this DB instance.
	DbParameterGroups []string `pulumi:"dbParameterGroups"`
	// Provides List of DB security groups associated to this DB instance.
	DbSecurityGroups []string `pulumi:"dbSecurityGroups"`
	// Specifies the name of the subnet group associated with the DB instance.
	DbSubnetGroup string `pulumi:"dbSubnetGroup"`
	// List of log types to export to cloudwatch.
	EnabledCloudwatchLogsExports []string `pulumi:"enabledCloudwatchLogsExports"`
	// The connection endpoint in `address:port` format.
	Endpoint string `pulumi:"endpoint"`
	// Provides the name of the database engine to be used for this DB instance.
	Engine string `pulumi:"engine"`
	// Indicates the database engine version.
	EngineVersion string `pulumi:"engineVersion"`
	// The canonical hosted zone ID of the DB instance (to be used in a Route 53 Alias record).
	HostedZoneId string `pulumi:"hostedZoneId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specifies the Provisioned IOPS (I/O operations per second) value.
	Iops int `pulumi:"iops"`
	// If StorageEncrypted is true, the KMS key identifier for the encrypted DB instance.
	KmsKeyId string `pulumi:"kmsKeyId"`
	// License model information for this DB instance.
	LicenseModel string `pulumi:"licenseModel"`
	// Contains the master username for the DB instance.
	MasterUsername string `pulumi:"masterUsername"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
	MonitoringInterval int `pulumi:"monitoringInterval"`
	// The ARN for the IAM role that permits RDS to send Enhanced Monitoring metrics to CloudWatch Logs.
	MonitoringRoleArn string `pulumi:"monitoringRoleArn"`
	// Specifies if the DB instance is a Multi-AZ deployment.
	MultiAz bool `pulumi:"multiAz"`
	// Provides the list of option group memberships for this DB instance.
	OptionGroupMemberships []string `pulumi:"optionGroupMemberships"`
	// The database port.
	Port int `pulumi:"port"`
	// Specifies the daily time range during which automated backups are created.
	PreferredBackupWindow string `pulumi:"preferredBackupWindow"`
	// Specifies the weekly time range during which system maintenance can occur in UTC.
	PreferredMaintenanceWindow string `pulumi:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for the DB instance.
	PubliclyAccessible bool `pulumi:"publiclyAccessible"`
	// The identifier of the source DB that this is a replica of.
	ReplicateSourceDb string `pulumi:"replicateSourceDb"`
	// The RDS Resource ID of this instance.
	ResourceId string `pulumi:"resourceId"`
	// Specifies whether the DB instance is encrypted.
	StorageEncrypted bool `pulumi:"storageEncrypted"`
	// Specifies the storage type associated with DB instance.
	StorageType string            `pulumi:"storageType"`
	Tags        map[string]string `pulumi:"tags"`
	// The time zone of the DB instance.
	Timezone string `pulumi:"timezone"`
	// Provides a list of VPC security group elements that the DB instance belongs to.
	VpcSecurityGroups []string `pulumi:"vpcSecurityGroups"`
}
