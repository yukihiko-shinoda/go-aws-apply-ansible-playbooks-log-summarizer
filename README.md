# go-aws-apply-ansible-playbooks-log-summarizer

[![PkgGoDev](https://pkg.go.dev/badge/github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer)](https://pkg.go.dev/github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer)
[![Test](https://github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer/workflows/Test/badge.svg)](https://github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer/actions?query=workflow%3ATest)
[![Go Report Card](https://goreportcard.com/badge/github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer)](https://goreportcard.com/report/github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer)
[![Test Coverage](https://api.codeclimate.com/v1/badges/6e60db0f480bda718d35/test_coverage)](https://codeclimate.com/github/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/6e60db0f480bda718d35/maintainability)](https://codeclimate.com/github/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer/maintainability)
[![Code Climate technical debt](https://img.shields.io/codeclimate/tech-debt/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer)](https://codeclimate.com/github/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer)](https://github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer/blob/main/go.mod)
[![Twitter URL](https://img.shields.io/twitter/url?url=https%3A%2F%2Fgithub.com%2Fyukihiko-shinoda%2Fgo-aws-apply-ansible-playbooks-log-summarizer)](http://twitter.com/share?text=go-aws-apply-ansible-playbooks-log-summarizer&url=https://github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer&hashtags=golang)

AWS-ApplyAnsiblePlaybooks log summarizer for Golang.

## Quickstart

```console
go get -d github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer
```

<!-- markdownlint-disable no-hard-tabs -->
```go
package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	awsapplyansibleplaybookslogsummarizer "github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer"
)

func main() {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	ssmClient := ssm.NewFromConfig(config)
	ec2Client := ec2.NewFromConfig(config)
	cloudwatchlogsClient := cloudwatchlogs.NewFromConfig(config)

	err := awsapplyansibleplaybookslogsummarizer.ViewLog(ssmClient, ec2Client, cloudwatchlogsClient)
	if err != nil {
		log.Fatal(err)
	}
}
```
<!-- markdownlint-enable no-hard-tabs -->

Example output:

```console
i-0abcdefghijklmno0 EC2InstanceName10 Some error 2
i-90abcdefghijklmn0  Some error 1
i-34567890abcdefgh0     EC2InstanceName3        51700219-01-14 11:59:43 +0000 UTC
TASK [file] *********************************************************************************************************************************************************************************************************************
PLAY RECAP **********************************************************************************************************************************************************************************************************************
localhost                  : ok=3    changed=1    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   


i-234567890abcdefg0     EC2InstanceName2        51700219-01-14 11:58:05 +0000 UTC
TASK [file] *********************************************************************************************************************************************************************************************************************
PLAY RECAP **********************************************************************************************************************************************************************************************************************
localhost                  : ok=3    changed=1    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   


i-7890abcdefghijkl0     EC2InstanceName7        51700219-01-14 12:08:45 +0000 UTC
TASK [file] *********************************************************************************************************************************************************************************************************************
TASK [file] *********************************************************************************************************************************************************************************************************************
PLAY RECAP **********************************************************************************************************************************************************************************************************************
localhost                  : ok=4    changed=2    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   


i-567890abcdefghij0     EC2InstanceName5        51700219-01-14 12:04:29 +0000 UTC
TASK [file] *********************************************************************************************************************************************************************************************************************
TASK [file] *********************************************************************************************************************************************************************************************************************
PLAY RECAP **********************************************************************************************************************************************************************************************************************
localhost                  : ok=4    changed=2    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   


i-890abcdefghijklm0     EC2InstanceName8        51700219-01-14 12:11:15 +0000 UTC
TASK [file] *********************************************************************************************************************************************************************************************************************
PLAY RECAP **********************************************************************************************************************************************************************************************************************
localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   


i-67890abcdefghijk0     EC2InstanceName6        51700219-01-14 12:06:28 +0000 UTC
TASK [file] *********************************************************************************************************************************************************************************************************************
PLAY RECAP **********************************************************************************************************************************************************************************************************************
localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   


i-4567890abcdefghi0     EC2InstanceName4        51700219-01-14 12:01:46 +0000 UTC
TASK [file] *********************************************************************************************************************************************************************************************************************
PLAY RECAP **********************************************************************************************************************************************************************************************************************
localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   


i-1234567890abcdef0     EC2InstanceName1        51700219-01-14 11:55:50 +0000 UTC
TASK [file] *********************************************************************************************************************************************************************************************************************
PLAY RECAP **********************************************************************************************************************************************************************************************************************
localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
```