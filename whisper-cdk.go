package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type WhisperCdkStackProps struct {
	awscdk.StackProps
}

func NewWhisperCdkStack(scope constructs.Construct, id string, props *WhisperCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	usersTable := awsdynamodb.NewTable(stack, jsii.String("dbwhisper"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("groupID"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("convoGroups"),
	})

	lambdaFunc := awslambda.NewFunction(stack, jsii.String("WhisperCdkFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("lambda/function.zip"), nil),
		Handler: jsii.String("main"),
	})

	api := awsapigateway.NewRestApi(stack, jsii.String("whisperApi"), &awsapigateway.RestApiProps{
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
			AllowMethods: jsii.Strings("GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"),
			AllowOrigins: jsii.Strings("*"),
		},
	})

	integration := awsapigateway.NewLambdaIntegration(lambdaFunc, nil)

	//Define Routes
	sampleSource := api.Root().AddResource(jsii.String("sample"), nil)
	sampleSource.AddMethod(jsii.String("GET"), integration, nil)

	groupResource := api.Root().AddResource(jsii.String("groups"), nil)
	groupResource.AddMethod(jsii.String("POST"), integration, nil)

	groupIncrementResource := groupResource.AddResource(jsii.String("increment"), nil)
	groupIDResource := groupIncrementResource.AddResource(jsii.String("{id}"), nil)
	groupIDResource.AddMethod(jsii.String("POST"), integration, nil)

	//Grant Table r/w
	usersTable.GrantReadWriteData(lambdaFunc)

	//Creating WebSocket Connection
	wsApi := awsapigatewayv2.NewWebSocketApi(stack, jsii.String("whisperws"), &awsapigatewayv2.WebSocketApiProps{
		ApiName: jsii.String("whisperwebsocket"),
	})

	awsapigatewayv2.NewWebSocketStage(stack, jsii.String("whispersocketstage"), &awsapigatewayv2.WebSocketStageProps{
		AutoDeploy:   jsii.Bool(true),
		StageName:    jsii.String("dev"),
		WebSocketApi: wsApi,
	})

	wsIntegration := awsapigatewayv2integrations.NewWebSocketLambdaIntegration(jsii.String("WsIntegration"), lambdaFunc, nil)
	wsApi.AddRoute(jsii.String("$connect"), &awsapigatewayv2.WebSocketRouteOptions{
		Integration: wsIntegration,
	})
	wsApi.AddRoute(jsii.String("$disconnect"), &awsapigatewayv2.WebSocketRouteOptions{
		Integration: wsIntegration,
	})
	wsApi.AddRoute(jsii.String("$default"), &awsapigatewayv2.WebSocketRouteOptions{
		Integration: wsIntegration,
	})
	wsApi.AddRoute(jsii.String("sendMessage"), &awsapigatewayv2.WebSocketRouteOptions{
		Integration: wsIntegration,
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewWhisperCdkStack(app, "WhisperCdkStack", &WhisperCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
