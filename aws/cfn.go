package aws

const cfnTemplate string = `
Description: AWS S3 bucket for tenant

Parameters:

    S3Path:
        Type: String

    S3Bucket:
        Type: String

    RoleName:
        Type: String

Resources:
    RootBucket:
        Type: AWS::S3::Bucket
        Properties:
            BucketName: !Ref S3Bucket

    Role:
        Type: AWS::IAM::Role
        Properties:
            Path: /
            RoleName: !Ref RoleName
            AssumeRolePolicyDocument: |
                {
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "ec2.amazonaws.com""
                        }
                    }]
                }
            Policies:
                - PolicyName: !Sub ${AWS::StackName}-s3-policy
                  PolicyDocument:
                    {
                        "Version": "2012-10-17",
                        "Statement": [
                            {
                                "Sid":"AllowListingOfNamespaceFolders",
                                "Effect": "Allow",
                                "Action": [
                                    "s3:ListBucket",
                                    "s3:ListBucketVersions"
                                ],
                                "Resource": !Join [ "", ["arn:aws:s3:::", !Ref S3Bucket] ],
                                "Condition": {
                                    "StringLike": {
                                        "s3:prefix": [
                                            !Join [ "", [!Ref S3Path, /*] ]
                                        ]
                                    }
                                }
                            },
                            {
                                "Sid":"AllowAllS3ActionsInNamespaceFolder",
                                "Effect": "Allow",
                                "Action": [
                                    "s3:*"
                                ],
                                "Resource": !Join [ "", ["arn:aws:s3:::", !Ref S3Bucket, /, !Ref S3Path, /*] ]
                            }
                        ]
                    }



Outputs:
    Role:
        Description: The name of the iam role to be used for the tenant
        Value: !Ref Role

    S3EndPoint:
        Description: The S3 endpoint
        Value: !Sub 'https://s3.${AWS::Region}.amazonaws.com'

    S3Bucket:
        Description: The S3 bucket
        Value: !Ref S3Bucket

    S3Path:
        Description: The S3 path under S3 bucket
        Value: !Ref S3Path
`
