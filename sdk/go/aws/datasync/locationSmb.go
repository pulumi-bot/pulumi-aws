// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a SMB Location within AWS DataSync.
//
// > **NOTE:** The DataSync Agents must be available before creating this resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/datasync"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = datasync.NewLocationSmb(ctx, "example", &datasync.LocationSmbArgs{
// 			AgentArns: pulumi.StringArray{
// 				pulumi.String(aws_datasync_agent.Example.Arn),
// 			},
// 			Password:       pulumi.String("ANotGreatPassword"),
// 			ServerHostname: pulumi.String("smb.example.com"),
// 			Subdirectory:   pulumi.String("/exported/path"),
// 			User:           pulumi.String("Guest"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocationSmb struct {
	pulumi.CustomResourceState

	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns pulumi.StringArrayOutput `pulumi:"agentArns"`
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the Windows domain the SMB server belongs to.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions LocationSmbMountOptionsPtrOutput `pulumi:"mountOptions"`
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password pulumi.StringOutput `pulumi:"password"`
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname pulumi.StringOutput `pulumi:"serverHostname"`
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory pulumi.StringOutput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.MapOutput    `pulumi:"tags"`
	Uri  pulumi.StringOutput `pulumi:"uri"`
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewLocationSmb registers a new resource with the given unique name, arguments, and options.
func NewLocationSmb(ctx *pulumi.Context,
	name string, args *LocationSmbArgs, opts ...pulumi.ResourceOption) (*LocationSmb, error) {
	if args == nil || args.AgentArns == nil {
		return nil, errors.New("missing required argument 'AgentArns'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.ServerHostname == nil {
		return nil, errors.New("missing required argument 'ServerHostname'")
	}
	if args == nil || args.Subdirectory == nil {
		return nil, errors.New("missing required argument 'Subdirectory'")
	}
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	if args == nil {
		args = &LocationSmbArgs{}
	}
	var resource LocationSmb
	err := ctx.RegisterResource("aws:datasync/locationSmb:LocationSmb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocationSmb gets an existing LocationSmb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocationSmb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocationSmbState, opts ...pulumi.ResourceOption) (*LocationSmb, error) {
	var resource LocationSmb
	err := ctx.ReadResource("aws:datasync/locationSmb:LocationSmb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocationSmb resources.
type locationSmbState struct {
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `pulumi:"agentArns"`
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `pulumi:"arn"`
	// The name of the Windows domain the SMB server belongs to.
	Domain *string `pulumi:"domain"`
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions *LocationSmbMountOptions `pulumi:"mountOptions"`
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password *string `pulumi:"password"`
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname *string `pulumi:"serverHostname"`
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]interface{} `pulumi:"tags"`
	Uri  *string                `pulumi:"uri"`
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User *string `pulumi:"user"`
}

type LocationSmbState struct {
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns pulumi.StringArrayInput
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringPtrInput
	// The name of the Windows domain the SMB server belongs to.
	Domain pulumi.StringPtrInput
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions LocationSmbMountOptionsPtrInput
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password pulumi.StringPtrInput
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname pulumi.StringPtrInput
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.MapInput
	Uri  pulumi.StringPtrInput
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User pulumi.StringPtrInput
}

func (LocationSmbState) ElementType() reflect.Type {
	return reflect.TypeOf((*locationSmbState)(nil)).Elem()
}

type locationSmbArgs struct {
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `pulumi:"agentArns"`
	// The name of the Windows domain the SMB server belongs to.
	Domain *string `pulumi:"domain"`
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions *LocationSmbMountOptions `pulumi:"mountOptions"`
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password string `pulumi:"password"`
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname string `pulumi:"serverHostname"`
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]interface{} `pulumi:"tags"`
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a LocationSmb resource.
type LocationSmbArgs struct {
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns pulumi.StringArrayInput
	// The name of the Windows domain the SMB server belongs to.
	Domain pulumi.StringPtrInput
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions LocationSmbMountOptionsPtrInput
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password pulumi.StringInput
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname pulumi.StringInput
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory pulumi.StringInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.MapInput
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User pulumi.StringInput
}

func (LocationSmbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locationSmbArgs)(nil)).Elem()
}
