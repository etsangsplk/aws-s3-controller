package common

import (
    "strings"
    "fmt"

    "github.com/aws/aws-sdk-go/aws/arn"
)

func isStringEmpty(s string) bool {
    return strings.TrimSpace(s) == ""
}

func validateStrings(a []string) error {
    for s in a {
        if isStringEmpty(s) {
            return fmt.Errorf("%v cannot be empty", s)
        }
    }
}

func parseRoleArn(arn string) (*arn.Arn, error) {
    a, err:= arn.Parse(arn)
    if err != nil {
        return nil, err
    }
    return a, err
}