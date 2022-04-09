package awsapplyansibleplaybookslogsummarizer

import (
	"fmt"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	ansiblelogparser "github.com/yukihiko-shinoda/go-ansible-log-parser"
	ansiblelogparserforcloudwatch "github.com/yukihiko-shinoda/go-ansible-log-parser-for-cloudwatch"
	"github.com/yukihiko-shinoda/go-aws-api-util/cloudwatch/logs"
	"github.com/yukihiko-shinoda/go-aws-api-util/ec2/instance"
	"github.com/yukihiko-shinoda/go-aws-api-util/ssm/command"
)

type structFailed struct {
	instanceId   string
	tagName      string
	errorMessage string
}

type structSucceed struct {
	instanceId string
	tagName    string
	timestamp  time.Time
	log        string
	playRecap  *ansiblelogparser.StructPlayRecap
}

func ViewLog(ssmClient command.ListCommandsListCommandInvocationsAPI, ec2Client instance.DescribeInstancesAPI, cloudwatchlogsClient logs.GetLogEventsAPI, commandId *string) error {
	commandInvocations, err := command.GetApplyAnsiblePlaybooksInvocations(ssmClient, commandId)
	if err != nil {
		return err
	}
	succeed, failed := checkCloudWatchLogs(ec2Client, cloudwatchlogsClient, commandInvocations)
	report(failed, succeed)
	return nil
}

func checkCloudWatchLogs(ec2Client instance.DescribeInstancesAPI, cloudwatchlogsClient logs.GetLogEventsAPI, commandInvocations []types.CommandInvocation) ([]structSucceed, []structFailed) {
	var failed []structFailed
	var succeed []structSucceed
	for _, commandInvocation := range commandInvocations {
		logStreamName := command.BuildLogStreamNameRunShellScriptStdout(commandInvocation)
		tagName, err := instance.GetNameFromTag(ec2Client, commandInvocation)
		if err != nil {
			errorString := err.Error()
			tagName = &errorString
		}
		events, err := logs.GetAllLogs(cloudwatchlogsClient, "/aws/systems-manager/run-command", logStreamName)
		if err != nil {
			failed = append(failed, structFailed{
				*commandInvocation.InstanceId,
				*tagName,
				err.Error(),
			})
			continue
		}

		playRecap, err := ansiblelogparserforcloudwatch.PickupNumberPlayRecap(events)
		if err != nil {
			failed = append(failed, structFailed{
				*commandInvocation.InstanceId,
				*tagName,
				err.Error(),
			})
			continue
		}
		message := ansiblelogparserforcloudwatch.PickupMessage(events)
		succeed = append(succeed, structSucceed{
			*commandInvocation.InstanceId,
			*tagName,
			time.Unix(*events[0].Timestamp/int64(time.Microsecond), 0).Local(),
			message,
			playRecap,
		})
	}
	return sortSucceed(succeed), sortFailed(failed)
}

func sortSucceed(succeed []structSucceed) []structSucceed {
	sort.Slice(succeed, func(i, j int) bool {
		playRecapI := succeed[i].playRecap
		playRecapJ := succeed[j].playRecap
		if playRecapI == nil {
			return true
		}
		if playRecapJ == nil {
			return false
		}
		if playRecapI.Failed > playRecapJ.Failed {
			return true
		}
		if playRecapI.Failed < playRecapJ.Failed {
			return false
		}
		if playRecapI.Changed > playRecapJ.Changed {
			return true
		}
		if playRecapI.Changed < playRecapJ.Changed {
			return false
		}
		return succeed[i].tagName > succeed[j].tagName
	})
	return succeed
}

func sortFailed(failed []structFailed) []structFailed {
	sort.Slice(failed, func(i, j int) bool {
		return failed[i].tagName > failed[j].tagName
	})
	return failed
}

func report(failed []structFailed, succeed []structSucceed) {
	for _, f := range failed {
		fmt.Println(
			f.instanceId,
			f.tagName,
			f.errorMessage,
		)
	}
	for _, s := range succeed {
		fmt.Printf(
			"%s\t%s\t%s\n%s\n",
			s.instanceId,
			s.tagName,
			s.timestamp,
			s.log,
		)
	}
}
