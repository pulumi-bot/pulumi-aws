// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Glue User Defined Function Resource.
type UserDefinedFunction struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId pulumi.StringPtrOutput `pulumi:"catalogId"`
	// The Java class that contains the function code.
	ClassName  pulumi.StringOutput `pulumi:"className"`
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The name of the Database to create the Function.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The name of the function.
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner of the function.
	OwnerName pulumi.StringOutput `pulumi:"ownerName"`
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType pulumi.StringOutput `pulumi:"ownerType"`
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris UserDefinedFunctionResourceUriArrayOutput `pulumi:"resourceUris"`
}

// NewUserDefinedFunction registers a new resource with the given unique name, arguments, and options.
func NewUserDefinedFunction(ctx *pulumi.Context,
	name string, args *UserDefinedFunctionArgs, opts ...pulumi.ResourceOption) (*UserDefinedFunction, error) {
	if args == nil || args.ClassName == nil {
		return nil, errors.New("missing required argument 'ClassName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.OwnerName == nil {
		return nil, errors.New("missing required argument 'OwnerName'")
	}
	if args == nil || args.OwnerType == nil {
		return nil, errors.New("missing required argument 'OwnerType'")
	}
	if args == nil {
		args = &UserDefinedFunctionArgs{}
	}
	var resource UserDefinedFunction
	err := ctx.RegisterResource("aws:glue/userDefinedFunction:UserDefinedFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserDefinedFunction gets an existing UserDefinedFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserDefinedFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserDefinedFunctionState, opts ...pulumi.ResourceOption) (*UserDefinedFunction, error) {
	var resource UserDefinedFunction
	err := ctx.ReadResource("aws:glue/userDefinedFunction:UserDefinedFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserDefinedFunction resources.
type userDefinedFunctionState struct {
	Arn *string `pulumi:"arn"`
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId *string `pulumi:"catalogId"`
	// The Java class that contains the function code.
	ClassName  *string `pulumi:"className"`
	CreateTime *string `pulumi:"createTime"`
	// The name of the Database to create the Function.
	DatabaseName *string `pulumi:"databaseName"`
	// The name of the function.
	Name *string `pulumi:"name"`
	// The owner of the function.
	OwnerName *string `pulumi:"ownerName"`
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType *string `pulumi:"ownerType"`
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris []UserDefinedFunctionResourceUri `pulumi:"resourceUris"`
}

type UserDefinedFunctionState struct {
	Arn pulumi.StringPtrInput
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId pulumi.StringPtrInput
	// The Java class that contains the function code.
	ClassName  pulumi.StringPtrInput
	CreateTime pulumi.StringPtrInput
	// The name of the Database to create the Function.
	DatabaseName pulumi.StringPtrInput
	// The name of the function.
	Name pulumi.StringPtrInput
	// The owner of the function.
	OwnerName pulumi.StringPtrInput
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType pulumi.StringPtrInput
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris UserDefinedFunctionResourceUriArrayInput
}

func (UserDefinedFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*userDefinedFunctionState)(nil)).Elem()
}

type userDefinedFunctionArgs struct {
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId *string `pulumi:"catalogId"`
	// The Java class that contains the function code.
	ClassName string `pulumi:"className"`
	// The name of the Database to create the Function.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the function.
	Name *string `pulumi:"name"`
	// The owner of the function.
	OwnerName string `pulumi:"ownerName"`
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType string `pulumi:"ownerType"`
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris []UserDefinedFunctionResourceUri `pulumi:"resourceUris"`
}

// The set of arguments for constructing a UserDefinedFunction resource.
type UserDefinedFunctionArgs struct {
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId pulumi.StringPtrInput
	// The Java class that contains the function code.
	ClassName pulumi.StringInput
	// The name of the Database to create the Function.
	DatabaseName pulumi.StringInput
	// The name of the function.
	Name pulumi.StringPtrInput
	// The owner of the function.
	OwnerName pulumi.StringInput
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType pulumi.StringInput
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris UserDefinedFunctionResourceUriArrayInput
}

func (UserDefinedFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userDefinedFunctionArgs)(nil)).Elem()
}
