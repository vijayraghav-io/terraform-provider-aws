// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecs

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/private/protocol/json/jsonutil"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func resourceTaskDefinitionMigrateState(v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	ctx := context.Background()
	conn := meta.(*conns.AWSClient).ECSConn(ctx)

	switch v {
	case 0:
		log.Println("[INFO] Found AWS ECS Task Definition State v0; migrating to v1")
		return migrateTaskDefinitionStateV0toV1(is, conn)
	default:
		return is, fmt.Errorf("Unexpected schema version: %d", v)
	}
}

func migrateTaskDefinitionStateV0toV1(is *terraform.InstanceState, conn *ecs.ECS) (*terraform.InstanceState, error) {
	arn := is.Attributes[names.AttrARN]

	ctx := context.TODO() // nosemgrep:ci.semgrep.migrate.context-todo
	// We need to pull definitions from the API b/c they're unrecoverable from the checksum
	td, err := conn.DescribeTaskDefinitionWithContext(ctx, &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(arn),
	})
	if err != nil {
		return nil, err
	}

	b, err := jsonutil.BuildJSON(td.TaskDefinition.ContainerDefinitions)
	if err != nil {
		return nil, err
	}

	is.Attributes["container_definitions"] = string(b)

	return is, nil
}
