/*
Copyright 2023 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/deckhouse/deckhouse/pkg/log"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/ptr"

	"github.com/deckhouse/deckhouse/go_lib/cloud-data/apis/v1alpha1"
)

type Discoverer struct {
	logger *log.Logger
	region string
}

func NewDiscoverer(logger *log.Logger) *Discoverer {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		logger.Fatal("AWS_REGION not found")
	}

	return &Discoverer{
		logger: logger,
		region: region,
	}
}

func (d *Discoverer) CheckCloudConditions(ctx context.Context) ([]v1alpha1.CloudCondition, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(d.region),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize new session: %v", err)
	}

	ec2Client := ec2.New(sess)

	var res []v1alpha1.CloudCondition

	// check permission for DescribeAddressesAttribute
	log.Debug("checking permission for DescribeAddressesAttribute")
	_, err = ec2Client.DescribeAddressesAttribute(&ec2.DescribeAddressesAttributeInput{
		DryRun: aws.Bool(true),
	})
	if err != nil {
		var awsErr awserr.Error
		if !errors.As(err, &awsErr) {
			return nil, fmt.Errorf("DescribeAddressesAttribute AWS IAM permission check error: %v", err)
		}

		if awsErr.Code() != "DryRunOperation" {
			res = append(res, v1alpha1.CloudCondition{
				Name:    "INSUFFICIENT_AWS_SA_PERMISSIONS",
				Message: "DescribeAddressesAttribute is not allowed",
				Ok:      false,
			})
		}
	}

	// check permission for DescribeInstanceTopology
	log.Debug("checking permission for DescribeInstanceTopology")
	_, err = ec2Client.DescribeInstanceTopology(&ec2.DescribeInstanceTopologyInput{
		DryRun: aws.Bool(true),
	})
	if err != nil {
		var awsErr awserr.Error
		if !errors.As(err, &awsErr) {
			return nil, fmt.Errorf("DescribeInstanceTopology AWS IAM permission check error: %v", err)
		}
		if awsErr.Code() != "DryRunOperation" {
			res = append(res, v1alpha1.CloudCondition{
				Name:    "INSUFFICIENT_AWS_SA_PERMISSIONS",
				Message: "DescribeInstanceTopology is not allowed",
				Ok:      false,
			})
		}
	}

	return res, nil
}

func (d *Discoverer) InstanceTypes(_ context.Context) ([]v1alpha1.InstanceType, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(d.region),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to initialize new session: %v", err)
	}

	ec2Client := ec2.New(sess)
	res := make([]v1alpha1.InstanceType, 0)

	var token *string

	for {
		out, err := ec2Client.DescribeInstanceTypes(&ec2.DescribeInstanceTypesInput{
			MaxResults: ptr.To(int64(100)),
			NextToken:  token,
		})
		if err != nil {
			return nil, err
		}

		for _, ins := range out.InstanceTypes {
			if ins.InstanceType == nil {
				return nil, fmt.Errorf("instance type is nil")
			}

			name := *ins.InstanceType

			if ins.VCpuInfo == nil {
				return nil, fmt.Errorf("VCpuInfo is nil for %s", name)
			}

			if ins.VCpuInfo.DefaultVCpus == nil {
				return nil, fmt.Errorf("VCpuInfo.DefaultVCpus is nil for %s", name)
			}

			if ins.MemoryInfo == nil {
				return nil, fmt.Errorf("MemoryInfo is nil for %s", name)
			}

			if ins.MemoryInfo.SizeInMiB == nil {
				return nil, fmt.Errorf("MemoryInfo.SizeInMiB is nil for %s", name)
			}

			res = append(res, v1alpha1.InstanceType{
				Name:     name,
				CPU:      resource.MustParse(strconv.FormatInt(*ins.VCpuInfo.DefaultVCpus, 10)),
				Memory:   resource.MustParse(strconv.FormatInt(*ins.MemoryInfo.SizeInMiB, 10) + "Mi"),
				RootDisk: resource.MustParse("0"),
			})
		}

		if out.NextToken == nil || *out.NextToken == "" {
			break
		} else {
			token = out.NextToken
		}
	}

	return res, nil
}

func (d *Discoverer) DiscoveryData(ctx context.Context, cloudProviderDiscoveryData []byte) ([]byte, error) {
	return nil, nil
}

// NotImplemented
func (d *Discoverer) DisksMeta(ctx context.Context) ([]v1alpha1.DiskMeta, error) {
	return []v1alpha1.DiskMeta{}, nil
}
