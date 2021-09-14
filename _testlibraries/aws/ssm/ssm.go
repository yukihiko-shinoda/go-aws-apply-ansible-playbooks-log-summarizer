package ssm

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type StructMockClient struct {
	TimeListCommandInvocations int
}

func (s *StructMockClient) ListCommands(ctx context.Context, params *ssm.ListCommandsInput, optFns ...func(*ssm.Options)) (*ssm.ListCommandsOutput, error) {
	return &ssm.ListCommandsOutput{
		Commands: []types.Command{
			{
				CommandId:         aws.String("12345678-90ab-cdef-ghij-klmn-opqrstuvwxyz"),
				RequestedDateTime: &time.Time{},
			},
		},
	}, nil
}

type stractInstanceIdNextToken struct {
	instanceId string
	nextToken  *string
}

func (s *StructMockClient) ListCommandInvocations(ctx context.Context, params *ssm.ListCommandInvocationsInput, optFns ...func(*ssm.Options)) (*ssm.ListCommandInvocationsOutput, error) {
	s.TimeListCommandInvocations++
	structInstanceIdNextToken := map[int]stractInstanceIdNextToken{
		1:  stractInstanceIdNextToken{"i-1234567890abcdef0", aws.String("next1")},
		2:  stractInstanceIdNextToken{"i-234567890abcdefg0", aws.String("next2")},
		3:  stractInstanceIdNextToken{"i-34567890abcdefgh0", aws.String("next3")},
		4:  stractInstanceIdNextToken{"i-4567890abcdefghi0", aws.String("next4")},
		5:  stractInstanceIdNextToken{"i-567890abcdefghij0", aws.String("next5")},
		6:  stractInstanceIdNextToken{"i-67890abcdefghijk0", aws.String("next6")},
		7:  stractInstanceIdNextToken{"i-7890abcdefghijkl0", aws.String("next7")},
		8:  stractInstanceIdNextToken{"i-890abcdefghijklm0", aws.String("next8")},
		9:  stractInstanceIdNextToken{"i-90abcdefghijklmn0", aws.String("next9")},
		10: stractInstanceIdNextToken{"i-0abcdefghijklmno0", aws.String("next10")},
	}[s.TimeListCommandInvocations]
	return createListCommandInvocationsOutput(structInstanceIdNextToken.instanceId, structInstanceIdNextToken.nextToken)
}

func createListCommandInvocationsOutput(instanceId string, nextToken *string) (*ssm.ListCommandInvocationsOutput, error) {
	return &ssm.ListCommandInvocationsOutput{
		CommandInvocations: []types.CommandInvocation{
			createCommandInvocation(instanceId),
		},
		NextToken: nextToken,
	}, nil
}

func createCommandInvocation(instanceId string) types.CommandInvocation {
	return types.CommandInvocation{
		CommandId:         aws.String("12345678-90ab-cdef-ghij-klmn-opqrstuvwxyz"),
		InstanceId:        aws.String(instanceId),
		InstanceName:      aws.String(""),
		RequestedDateTime: &time.Time{},
	}
}

func NewMockClient() *StructMockClient {
	return &StructMockClient{
		TimeListCommandInvocations: 0,
	}
}

type StructMockClientError struct{}

func (s *StructMockClientError) ListCommands(ctx context.Context, params *ssm.ListCommandsInput, optFns ...func(*ssm.Options)) (*ssm.ListCommandsOutput, error) {
	return nil, errors.New("Some error.")
}

func (s *StructMockClientError) ListCommandInvocations(ctx context.Context, params *ssm.ListCommandInvocationsInput, optFns ...func(*ssm.Options)) (*ssm.ListCommandInvocationsOutput, error) {
	return nil, nil
}
