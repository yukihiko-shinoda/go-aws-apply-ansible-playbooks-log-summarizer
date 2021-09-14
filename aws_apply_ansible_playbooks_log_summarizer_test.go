package awsapplyansibleplaybookslogsummarizer

import (
	"testing"

	"github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer/_testlibraries/aws/cloudwatchlogs"
	"github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer/_testlibraries/aws/ec2"
	"github.com/yukihiko-shinoda/go-aws-apply-ansible-playbooks-log-summarizer/_testlibraries/aws/ssm"
)

func TestViewLog(t *testing.T) {
	ssmClient := ssm.NewMockClient()
	ec2Client := &ec2.StructMockClient{}
	cloudwatchlogsClient := &cloudwatchlogs.StructMockClient{}
	err := ViewLog(ssmClient, ec2Client, cloudwatchlogsClient)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestViewLogError(t *testing.T) {
	ssmClient := &ssm.StructMockClientError{}
	ec2Client := &ec2.StructMockClient{}
	cloudwatchlogsClient := &cloudwatchlogs.StructMockClient{}
	err := ViewLog(ssmClient, ec2Client, cloudwatchlogsClient)
	if err == nil {
		t.Errorf("Error is not returned")
	}
}
