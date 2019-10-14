package resources

import "github.com/awslabs/goformation/v2/cloudformation/policies"

// AWSEFSFileSystem_LifecyclePolicy AWS CloudFormation Resource (AWS::EFS::FileSystem.LifecyclePolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticfilesystem-filesystem-lifecyclepolicy.html
type AWSEFSFileSystem_LifecyclePolicy struct {

	// TransitionToIA AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticfilesystem-filesystem-lifecyclepolicy.html#cfn-elasticfilesystem-filesystem-lifecyclepolicy-transitiontoia
	TransitionToIA string `json:"TransitionToIA,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEFSFileSystem_LifecyclePolicy) AWSCloudFormationType() string {
	return "AWS::EFS::FileSystem.LifecyclePolicy"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEFSFileSystem_LifecyclePolicy) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEFSFileSystem_LifecyclePolicy) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEFSFileSystem_LifecyclePolicy) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEFSFileSystem_LifecyclePolicy) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEFSFileSystem_LifecyclePolicy) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEFSFileSystem_LifecyclePolicy) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
