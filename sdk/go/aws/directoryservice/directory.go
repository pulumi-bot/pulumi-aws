// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Simple or Managed Microsoft directory in AWS Directory Service.
//
// > **Note:** All arguments including the password and customer username will be stored in the raw state as plain-text.
//
// ## Example Usage
// ### SimpleAD
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directoryservice"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			CidrBlock:        pulumi.String("10.0.1.0/24"),
// 			VpcId:            main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		barSubnet, err := ec2.NewSubnet(ctx, "barSubnet", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String("us-west-2b"),
// 			CidrBlock:        pulumi.String("10.0.2.0/24"),
// 			VpcId:            main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = directoryservice.NewDirectory(ctx, "barDirectory", &directoryservice.DirectoryArgs{
// 			Password: pulumi.String("SuperSecretPassw0rd"),
// 			Size:     pulumi.String("Small"),
// 			Tags: map[string]interface{}{
// 				"Project": "foo",
// 			},
// 			VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
// 				SubnetIds: pulumi.StringArray{
// 					foo.ID(),
// 					barSubnet.ID(),
// 				},
// 				VpcId: main.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Microsoft Active Directory (MicrosoftAD)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directoryservice"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			CidrBlock:        pulumi.String("10.0.1.0/24"),
// 			VpcId:            main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		barSubnet, err := ec2.NewSubnet(ctx, "barSubnet", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String("us-west-2b"),
// 			CidrBlock:        pulumi.String("10.0.2.0/24"),
// 			VpcId:            main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = directoryservice.NewDirectory(ctx, "barDirectory", &directoryservice.DirectoryArgs{
// 			Edition:  pulumi.String("Standard"),
// 			Password: pulumi.String("SuperSecretPassw0rd"),
// 			Tags: map[string]interface{}{
// 				"Project": "foo",
// 			},
// 			Type: pulumi.String("MicrosoftAD"),
// 			VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
// 				SubnetIds: pulumi.StringArray{
// 					foo.ID(),
// 					barSubnet.ID(),
// 				},
// 				VpcId: main.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Microsoft Active Directory Connector (ADConnector)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directoryservice"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			CidrBlock:        pulumi.String("10.0.1.0/24"),
// 			VpcId:            main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bar, err := ec2.NewSubnet(ctx, "bar", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String("us-west-2b"),
// 			CidrBlock:        pulumi.String("10.0.2.0/24"),
// 			VpcId:            main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = directoryservice.NewDirectory(ctx, "connector", &directoryservice.DirectoryArgs{
// 			ConnectSettings: &directoryservice.DirectoryConnectSettingsArgs{
// 				CustomerDnsIps: pulumi.StringArray{
// 					pulumi.String("A.B.C.D"),
// 				},
// 				CustomerUsername: pulumi.String("Admin"),
// 				SubnetIds: pulumi.StringArray{
// 					foo.ID(),
// 					bar.ID(),
// 				},
// 				VpcId: main.ID(),
// 			},
// 			Password: pulumi.String("SuperSecretPassw0rd"),
// 			Size:     pulumi.String("Small"),
// 			Type:     pulumi.String("ADConnector"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Directory struct {
	pulumi.CustomResourceState

	// The access URL for the directory, such as `http://alias.awsapps.com`.
	AccessUrl pulumi.StringOutput `pulumi:"accessUrl"`
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Connector related information about the directory. Fields documented below.
	ConnectSettings DirectoryConnectSettingsPtrOutput `pulumi:"connectSettings"`
	// A textual description for the directory.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of IP addresses of the DNS servers for the directory or connector.
	DnsIpAddresses pulumi.StringArrayOutput `pulumi:"dnsIpAddresses"`
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
	Edition pulumi.StringOutput `pulumi:"edition"`
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso pulumi.BoolPtrOutput `pulumi:"enableSso"`
	// The fully qualified name for the directory, such as `corp.example.com`
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for the directory administrator or connector user.
	Password pulumi.StringOutput `pulumi:"password"`
	// The ID of the security group created by the directory.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The short name of the directory, such as `CORP`.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
	// The size of the directory (`Small` or `Large` are accepted values).
	Size pulumi.StringOutput `pulumi:"size"`
	// A map of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// VPC related information about the directory. Fields documented below.
	VpcSettings DirectoryVpcSettingsPtrOutput `pulumi:"vpcSettings"`
}

// NewDirectory registers a new resource with the given unique name, arguments, and options.
func NewDirectory(ctx *pulumi.Context,
	name string, args *DirectoryArgs, opts ...pulumi.ResourceOption) (*Directory, error) {
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil {
		args = &DirectoryArgs{}
	}
	var resource Directory
	err := ctx.RegisterResource("aws:directoryservice/directory:Directory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectory gets an existing Directory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectoryState, opts ...pulumi.ResourceOption) (*Directory, error) {
	var resource Directory
	err := ctx.ReadResource("aws:directoryservice/directory:Directory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Directory resources.
type directoryState struct {
	// The access URL for the directory, such as `http://alias.awsapps.com`.
	AccessUrl *string `pulumi:"accessUrl"`
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias *string `pulumi:"alias"`
	// Connector related information about the directory. Fields documented below.
	ConnectSettings *DirectoryConnectSettings `pulumi:"connectSettings"`
	// A textual description for the directory.
	Description *string `pulumi:"description"`
	// A list of IP addresses of the DNS servers for the directory or connector.
	DnsIpAddresses []string `pulumi:"dnsIpAddresses"`
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
	Edition *string `pulumi:"edition"`
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso *bool `pulumi:"enableSso"`
	// The fully qualified name for the directory, such as `corp.example.com`
	Name *string `pulumi:"name"`
	// The password for the directory administrator or connector user.
	Password *string `pulumi:"password"`
	// The ID of the security group created by the directory.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The short name of the directory, such as `CORP`.
	ShortName *string `pulumi:"shortName"`
	// The size of the directory (`Small` or `Large` are accepted values).
	Size *string `pulumi:"size"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type *string `pulumi:"type"`
	// VPC related information about the directory. Fields documented below.
	VpcSettings *DirectoryVpcSettings `pulumi:"vpcSettings"`
}

type DirectoryState struct {
	// The access URL for the directory, such as `http://alias.awsapps.com`.
	AccessUrl pulumi.StringPtrInput
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias pulumi.StringPtrInput
	// Connector related information about the directory. Fields documented below.
	ConnectSettings DirectoryConnectSettingsPtrInput
	// A textual description for the directory.
	Description pulumi.StringPtrInput
	// A list of IP addresses of the DNS servers for the directory or connector.
	DnsIpAddresses pulumi.StringArrayInput
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
	Edition pulumi.StringPtrInput
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso pulumi.BoolPtrInput
	// The fully qualified name for the directory, such as `corp.example.com`
	Name pulumi.StringPtrInput
	// The password for the directory administrator or connector user.
	Password pulumi.StringPtrInput
	// The ID of the security group created by the directory.
	SecurityGroupId pulumi.StringPtrInput
	// The short name of the directory, such as `CORP`.
	ShortName pulumi.StringPtrInput
	// The size of the directory (`Small` or `Large` are accepted values).
	Size pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type pulumi.StringPtrInput
	// VPC related information about the directory. Fields documented below.
	VpcSettings DirectoryVpcSettingsPtrInput
}

func (DirectoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryState)(nil)).Elem()
}

type directoryArgs struct {
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias *string `pulumi:"alias"`
	// Connector related information about the directory. Fields documented below.
	ConnectSettings *DirectoryConnectSettings `pulumi:"connectSettings"`
	// A textual description for the directory.
	Description *string `pulumi:"description"`
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
	Edition *string `pulumi:"edition"`
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso *bool `pulumi:"enableSso"`
	// The fully qualified name for the directory, such as `corp.example.com`
	Name *string `pulumi:"name"`
	// The password for the directory administrator or connector user.
	Password string `pulumi:"password"`
	// The short name of the directory, such as `CORP`.
	ShortName *string `pulumi:"shortName"`
	// The size of the directory (`Small` or `Large` are accepted values).
	Size *string `pulumi:"size"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type *string `pulumi:"type"`
	// VPC related information about the directory. Fields documented below.
	VpcSettings *DirectoryVpcSettings `pulumi:"vpcSettings"`
}

// The set of arguments for constructing a Directory resource.
type DirectoryArgs struct {
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias pulumi.StringPtrInput
	// Connector related information about the directory. Fields documented below.
	ConnectSettings DirectoryConnectSettingsPtrInput
	// A textual description for the directory.
	Description pulumi.StringPtrInput
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
	Edition pulumi.StringPtrInput
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso pulumi.BoolPtrInput
	// The fully qualified name for the directory, such as `corp.example.com`
	Name pulumi.StringPtrInput
	// The password for the directory administrator or connector user.
	Password pulumi.StringInput
	// The short name of the directory, such as `CORP`.
	ShortName pulumi.StringPtrInput
	// The size of the directory (`Small` or `Large` are accepted values).
	Size pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type pulumi.StringPtrInput
	// VPC related information about the directory. Fields documented below.
	VpcSettings DirectoryVpcSettingsPtrInput
}

func (DirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryArgs)(nil)).Elem()
}
