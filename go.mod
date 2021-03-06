module github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer

go 1.16

require (
	github.com/aws/aws-sdk-go-v2 v1.16.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.15.4
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.34.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.24.1
	github.com/fzipp/gocyclo v0.3.1
	github.com/golangci/golangci-lint v1.42.1
	github.com/yukihiko-shinoda/go-ansible-log-parser v0.0.0-20210908183457-076b1c7353ca
	github.com/yukihiko-shinoda/go-ansible-log-parser-for-cloudwatch v0.0.0-20210909181638-b1c18aaef0de
	github.com/yukihiko-shinoda/go-aws-api-util v0.0.0-20220409075707-b13d6d49d9d4
)
