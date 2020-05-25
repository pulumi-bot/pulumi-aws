module github.com/pulumi/pulumi-aws/provider/v2

go 1.13

require (
	github.com/hashicorp/aws-sdk-go-base v0.4.0
	github.com/hashicorp/terraform-plugin-sdk v1.13.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.3.3
	github.com/pulumi/pulumi/sdk/v2 v2.2.2-0.20200519081838-f66100ce28b0
	github.com/terraform-providers/terraform-provider-aws v0.0.0-20191010190908-1261a98537f2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-aws => github.com/pulumi/terraform-provider-aws v1.38.1-0.20200522142732-42ec5c9dfcc6
)

replace github.com/pulumi/pulumi-terraform-bridge/v2 => ../../pulumi-terraform-bridge
