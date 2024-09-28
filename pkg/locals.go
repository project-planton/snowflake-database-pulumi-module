package pkg

import (
	"github.com/plantoncloud/project-planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/snowflake/snowflakedatabase"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	SnowflakeDatabase *snowflakedatabase.SnowflakeDatabase
}

func initializeLocals(ctx *pulumi.Context, stackInput *snowflakedatabase.SnowflakeDatabaseStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.SnowflakeDatabase = stackInput.Target

	return locals
}
