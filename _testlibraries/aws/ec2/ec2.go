package ec2

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type unionType struct {
	success *string
	fail    *string
}

type StructMockClient struct{}

func (s *StructMockClient) DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	union := map[string]unionType{
		"i-1234567890abcdef0": unionType{aws.String("EC2InstanceName1"), nil},
		"i-234567890abcdefg0": unionType{aws.String("EC2InstanceName2"), nil},
		"i-34567890abcdefgh0": unionType{aws.String("EC2InstanceName3"), nil},
		"i-4567890abcdefghi0": unionType{aws.String("EC2InstanceName4"), nil},
		"i-567890abcdefghij0": unionType{aws.String("EC2InstanceName5"), nil},
		"i-67890abcdefghijk0": unionType{aws.String("EC2InstanceName6"), nil},
		"i-7890abcdefghijkl0": unionType{aws.String("EC2InstanceName7"), nil},
		"i-890abcdefghijklm0": unionType{aws.String("EC2InstanceName8"), nil},
		"i-90abcdefghijklmn0": unionType{nil, aws.String("Some error 1")},
		"i-0abcdefghijklmno0": unionType{aws.String("EC2InstanceName10"), nil},
	}[params.InstanceIds[0]]
	if union.success != nil {
		return createDescribeInstancesOutput(*union.success)
	}
	if union.fail != nil {
		return nil, errors.New(*union.fail)
	}
	return nil, errors.New("Unexpected InstanceId.")
}

func createDescribeInstancesOutput(tagName string) (*ec2.DescribeInstancesOutput, error) {
	return &ec2.DescribeInstancesOutput{
		Reservations: []types.Reservation{
			{
				Instances: []types.Instance{
					{
						Tags: []types.Tag{
							{
								Key:   aws.String("Name"),
								Value: aws.String(tagName),
							},
						},
					},
					{},
				},
			},
			{},
		},
	}, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
