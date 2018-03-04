package aws

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"
    "time"

    "github.com/cnf/structhash"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/arn"
    "github.com/aws/aws-sdk-go/service/cloudformation"
    "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/ec2/ec2iface"
    "github.com/aws/aws-sdk-go/service/ecs"
    "github.com/aws/aws-sdk-go/service/ecs/ecsiface"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3iface"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/request"
    "github.com/aws/aws-sdk-go/service/iam"
    "github.com/aws/aws-sdk-go/service/iam/iamiface"
    "github.com/aws/aws-sdk-go/service/sqs/sqsiface"
    "github.com/aws/aws-sdk-go/service/sqs"
)

type context=aws.context

const policyDoc = ``

                // - PolicyName: !Sub ${AWS::StackName}-sqs-policy
                //   PolicyDocument:
                //     {
                //         "Version": "2012-10-17",
                //         "Statement": [{
                //             "Effect": "Allow",
                //             "Action": [
                //                 "sqs:*"
                //             ],
                //             "Resource": !Join [ "", ["arn:aws:sqs:*:*:", !GetAtt Queue.QueueName] ],
                //         }]
                //     }
                // - PolicyName: !Sub ${AWS::StackName}-dead-sqs-policy
                //   PolicyDocument:
                //     {
                //         "Version": "2012-10-17",
                //         "Statement": [{
                //             "Effect": "Allow",
                //             "Action": [
                //                 "sqs:*",
                //             ],
                //             "Resource": !Join [ "", ["arn:aws:sqs:*:*:", !GetAtt DeadLetterQueue.QueueName] ],
                //         }]
                //     }

type SQS struct {
    Role
    Queue
    QueueName
    QueueUrl
    KMSKey
}

type struct Client {
    //ctx context
    // iamRole
    // kms

    iamSvc iamIface.IAMAPI,
    sqsSvc sqsiface.SQSAPI
}

func New(iam iamIface.IAMAPI, sqs sqsiface.SQSAPI) *Client{
    return &SQS{
        iamSvc: iam,
        sqsSvc: sqs,
    }
}

func (c *Client) CreateQueue(name string, kmsKey, role string) (*SQS, error){
    q, err := c.sqsSvc.CreateQueue(&sqs.CreateQueueInput{
        Attributes: map[string]string{
            "VisibilityTimeout":
            "MessageRetentionPeriod":
            "KmsMasterKeyId":
            "RedrivePolicy":
            },
        QueueName: aws.String(name(),
        })

    return q, err
}

func (c *Client) DeleteQueue(url string) error {

    err := c.sqsSvc.DeleteQueue(&sqs.DeleteQueueInput{
        QueueUrl: aws.String(url),
        })
    return err
}

func (c* Client) createPolicy() (*iam.Policy, error) {
    p, err := c.iamSvc.CreatePolicy(&iam.CreatePolicyInput{
    Description: aws.string(),
    Path: aws.string(),
    PolicyDocument: aws.string(policyDoc),
    PolicyName:aws.string(name),
        })
    return p, err
}

func (c *Client) deletePolicy(arn string) error {
    _, err := c.iamSvc.DeletePolicy(&iam.DeletePolicyInput{
    PolicyArn: aws.string(arn),
        })
    return err
}

func (c *Client) detachPolicy(arn string) error {
    _, err := c.iamSvc.DetachPolicy(&iam.DeletePolicyInput{
    PolicyArn: aws.string(arn),
        })
    return err
}

func createQueueUrl() string {
    return fmt.Sprintf("", ...)
}