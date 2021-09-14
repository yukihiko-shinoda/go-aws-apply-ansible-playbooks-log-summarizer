package cloudwatchlogs

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

type stractArgumentCreateGetLogEventOutput struct {
	changed          int
	failed           int
	nextForwardToken *string
}

type unionType struct {
	success *stractArgumentCreateGetLogEventOutput
	fail    *string
}

type StructMockClient struct{}

func (s *StructMockClient) GetLogEvents(ctx context.Context, params *cloudwatchlogs.GetLogEventsInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogEventsOutput, error) {
	log.Println(*params.LogStreamName)
	instanceId := substr(*params.LogStreamName, 42, 19)
	union := map[string]unionType{
		"i-1234567890abcdef0": unionType{&stractArgumentCreateGetLogEventOutput{1, 0, nil}, nil},
		"i-234567890abcdefg0": unionType{&stractArgumentCreateGetLogEventOutput{1, 1, nil}, nil},
		"i-34567890abcdefgh0": unionType{&stractArgumentCreateGetLogEventOutput{1, 1, nil}, nil},
		"i-4567890abcdefghi0": unionType{&stractArgumentCreateGetLogEventOutput{1, 0, nil}, nil},
		"i-567890abcdefghij0": unionType{&stractArgumentCreateGetLogEventOutput{2, 0, nil}, nil},
		"i-67890abcdefghijk0": unionType{&stractArgumentCreateGetLogEventOutput{1, 0, nil}, nil},
		"i-7890abcdefghijkl0": unionType{&stractArgumentCreateGetLogEventOutput{2, 0, nil}, nil},
		"i-890abcdefghijklm0": unionType{&stractArgumentCreateGetLogEventOutput{1, 0, nil}, nil},
		"i-90abcdefghijklmn0": unionType{nil, aws.String("Some error 1")},
		"i-0abcdefghijklmno0": unionType{nil, aws.String("Some error 2")},
	}[instanceId]
	if union.success != nil {
		return createGetLogEventOutput(union.success.changed, union.success.failed, union.success.nextForwardToken)
	}
	if union.fail != nil {
		return nil, errors.New(*union.fail)
	}
	return nil, errors.New("Unexpected LogStreamName.")
}

func createGetLogEventOutput(changed int, failed int, nextForwardToken *string) (*cloudwatchlogs.GetLogEventsOutput, error) {
	message, err := loadMessage(changed, failed)
	if err != nil {
		return nil, err
	}
	return &cloudwatchlogs.GetLogEventsOutput{
		Events: []types.OutputLogEvent{
			{
				Message:   message,
				Timestamp: aws.Int64(time.Now().UnixNano()),
			},
		},
		NextForwardToken: nextForwardToken,
	}, nil
}

func loadMessage(changed int, failed int) (*string, error) {
	_, file, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(file), "testdata", "log_example_changed_"+strconv.Itoa(changed)+"_failed_"+strconv.Itoa(failed)+".txt")
	messageByte, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	message := string(messageByte)
	return &message, nil
}
