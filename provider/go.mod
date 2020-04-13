module github.com/pulumi/pulumi-aws/provider

go 1.13

require (
	github.com/hashicorp/aws-sdk-go-base v0.4.0
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pulumi/pulumi-terraform-bridge v1.8.5-0.20200402004340-25b33bb489e0
	github.com/pulumi/pulumi/sdk v1.14.1-0.20200402003901-9e377419164b
	github.com/terraform-providers/terraform-provider-aws v0.0.0-20191010190908-1261a98537f2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-aws => github.com/pulumi/terraform-provider-aws v1.38.1-0.20200410125707-78b003cab40e
)

replace github.com/pulumi/pulumi-terraform-bridge => ../../pulumi-terraform-bridge
