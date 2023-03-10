// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AssignPublicIP string

const (
	AssignPublicIP_ENABLED  AssignPublicIP = "ENABLED"
	AssignPublicIP_DISABLED AssignPublicIP = "DISABLED"
)

type BatchJobDependencyType string

const (
	BatchJobDependencyType_N_TO_N     BatchJobDependencyType = "N_TO_N"
	BatchJobDependencyType_SEQUENTIAL BatchJobDependencyType = "SEQUENTIAL"
)

type BatchResourceRequirementType string

const (
	BatchResourceRequirementType_GPU    BatchResourceRequirementType = "GPU"
	BatchResourceRequirementType_MEMORY BatchResourceRequirementType = "MEMORY"
	BatchResourceRequirementType_VCPU   BatchResourceRequirementType = "VCPU"
)

type DynamoDBStreamStartPosition string

const (
	DynamoDBStreamStartPosition_TRIM_HORIZON DynamoDBStreamStartPosition = "TRIM_HORIZON"
	DynamoDBStreamStartPosition_LATEST       DynamoDBStreamStartPosition = "LATEST"
)

type ECSEnvironmentFileType string

const (
	ECSEnvironmentFileType_s3 ECSEnvironmentFileType = "s3"
)

type ECSResourceRequirementType string

const (
	ECSResourceRequirementType_GPU                  ECSResourceRequirementType = "GPU"
	ECSResourceRequirementType_InferenceAccelerator ECSResourceRequirementType = "InferenceAccelerator"
)

type KinesisStreamStartPosition string

const (
	KinesisStreamStartPosition_TRIM_HORIZON KinesisStreamStartPosition = "TRIM_HORIZON"
	KinesisStreamStartPosition_LATEST       KinesisStreamStartPosition = "LATEST"
	KinesisStreamStartPosition_AT_TIMESTAMP KinesisStreamStartPosition = "AT_TIMESTAMP"
)

type LaunchType string

const (
	LaunchType_EC2      LaunchType = "EC2"
	LaunchType_FARGATE  LaunchType = "FARGATE"
	LaunchType_EXTERNAL LaunchType = "EXTERNAL"
)

type MSKStartPosition string

const (
	MSKStartPosition_TRIM_HORIZON MSKStartPosition = "TRIM_HORIZON"
	MSKStartPosition_LATEST       MSKStartPosition = "LATEST"
)

type OnPartialBatchItemFailureStreams string

const (
	OnPartialBatchItemFailureStreams_AUTOMATIC_BISECT OnPartialBatchItemFailureStreams = "AUTOMATIC_BISECT"
)

type PipeState string

const (
	PipeState_RUNNING       PipeState = "RUNNING"
	PipeState_STOPPED       PipeState = "STOPPED"
	PipeState_CREATING      PipeState = "CREATING"
	PipeState_UPDATING      PipeState = "UPDATING"
	PipeState_DELETING      PipeState = "DELETING"
	PipeState_STARTING      PipeState = "STARTING"
	PipeState_STOPPING      PipeState = "STOPPING"
	PipeState_CREATE_FAILED PipeState = "CREATE_FAILED"
	PipeState_UPDATE_FAILED PipeState = "UPDATE_FAILED"
	PipeState_START_FAILED  PipeState = "START_FAILED"
	PipeState_STOP_FAILED   PipeState = "STOP_FAILED"
)

type PipeTargetInvocationType string

const (
	PipeTargetInvocationType_REQUEST_RESPONSE PipeTargetInvocationType = "REQUEST_RESPONSE"
	PipeTargetInvocationType_FIRE_AND_FORGET  PipeTargetInvocationType = "FIRE_AND_FORGET"
)

type PlacementConstraintType string

const (
	PlacementConstraintType_distinctInstance PlacementConstraintType = "distinctInstance"
	PlacementConstraintType_memberOf         PlacementConstraintType = "memberOf"
)

type PlacementStrategyType string

const (
	PlacementStrategyType_random  PlacementStrategyType = "random"
	PlacementStrategyType_spread  PlacementStrategyType = "spread"
	PlacementStrategyType_binpack PlacementStrategyType = "binpack"
)

type PropagateTags string

const (
	PropagateTags_TASK_DEFINITION PropagateTags = "TASK_DEFINITION"
)

type RequestedPipeState string

const (
	RequestedPipeState_RUNNING RequestedPipeState = "RUNNING"
	RequestedPipeState_STOPPED RequestedPipeState = "STOPPED"
)

type RequestedPipeStateDescribeResponse string

const (
	RequestedPipeStateDescribeResponse_RUNNING RequestedPipeStateDescribeResponse = "RUNNING"
	RequestedPipeStateDescribeResponse_STOPPED RequestedPipeStateDescribeResponse = "STOPPED"
	RequestedPipeStateDescribeResponse_DELETED RequestedPipeStateDescribeResponse = "DELETED"
)

type SelfManagedKafkaStartPosition string

const (
	SelfManagedKafkaStartPosition_TRIM_HORIZON SelfManagedKafkaStartPosition = "TRIM_HORIZON"
	SelfManagedKafkaStartPosition_LATEST       SelfManagedKafkaStartPosition = "LATEST"
)