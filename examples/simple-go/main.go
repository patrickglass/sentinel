package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	fmt.Println("Starting service...")
	// Register the workflow
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(Workflow)
	w.RegisterActivity(Activity)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

type VPC struct {
	Name string
}

type Cluster struct {
	name          string
	owner         string
	storageGB     int
	engine        string
	engineVersion string
	size          string
	username      string
	password      string
	isProtected   bool
}

func (c *Cluster) CreateNetwork(ctx context.Context, vpc VPC) error {
	fmt.Println("Creating network...")
	return nil
}

// Activity is a sample activity function
func Activity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello %s!", name), nil
}

// Workflow is a sample workflow function
func Workflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	db := Cluster{
		"testdb",
		"patrick.glass@databricks.com",
		10,
		"mysql",
		"5.7",
		"db.xlarge",
		"root",
		"password",
		false,
	}
	vpc := VPC{"us-west-2-484u39849"}

	var result string
	err := workflow.ExecuteActivity(ctx, Activity, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, db.CreateNetwork, vpc).Get(ctx, nil)
	if err != nil {
		return "", err
	}
	return result, err
}
