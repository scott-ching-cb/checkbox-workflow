// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: execute_workflow.proto

package workflow

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// WorkflowFormData is the form data, input by the user, to perform weather workflow alert check based on a threshold.
type WorkflowFormData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	City          string                 `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Operator      string                 `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Threshold     float64                `protobuf:"fixed64,5,opt,name=threshold,proto3" json:"threshold,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowFormData) Reset() {
	*x = WorkflowFormData{}
	mi := &file_execute_workflow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowFormData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowFormData) ProtoMessage() {}

func (x *WorkflowFormData) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowFormData.ProtoReflect.Descriptor instead.
func (*WorkflowFormData) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowFormData) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *WorkflowFormData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *WorkflowFormData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkflowFormData) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *WorkflowFormData) GetThreshold() float64 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

// *
// ExecuteWorkflowRequest is the request body for updating and executing a workflow.
// Defined in spec section 'Detailed Requirements' sub-section 2.
type ExecuteWorkflowRequest struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Condition     *ExecuteWorkflowRequest_Condition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	FormData      *WorkflowFormData                 `protobuf:"bytes,2,opt,name=form_data,json=formData,proto3" json:"form_data,omitempty"`
	WorkflowEdges []*Edge                           `protobuf:"bytes,3,rep,name=workflow_edges,json=workflowEdges,proto3" json:"workflow_edges,omitempty"`
	WorkflowNodes []*Node                           `protobuf:"bytes,4,rep,name=workflow_nodes,json=workflowNodes,proto3" json:"workflow_nodes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteWorkflowRequest) Reset() {
	*x = ExecuteWorkflowRequest{}
	mi := &file_execute_workflow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteWorkflowRequest) ProtoMessage() {}

func (x *ExecuteWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteWorkflowRequest.ProtoReflect.Descriptor instead.
func (*ExecuteWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteWorkflowRequest) GetCondition() *ExecuteWorkflowRequest_Condition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *ExecuteWorkflowRequest) GetFormData() *WorkflowFormData {
	if x != nil {
		return x.FormData
	}
	return nil
}

func (x *ExecuteWorkflowRequest) GetWorkflowEdges() []*Edge {
	if x != nil {
		return x.WorkflowEdges
	}
	return nil
}

func (x *ExecuteWorkflowRequest) GetWorkflowNodes() []*Node {
	if x != nil {
		return x.WorkflowNodes
	}
	return nil
}

// *
// ExecuteError is the error returned if the validation or persistence of the workflow fails.
type ExecuteError struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteError) Reset() {
	*x = ExecuteError{}
	mi := &file_execute_workflow_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteError) ProtoMessage() {}

func (x *ExecuteError) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteError.ProtoReflect.Descriptor instead.
func (*ExecuteError) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// *
// ExecutionStep summary of the execution step for a specific workflow node.
type ExecutionStep struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Description   string                 `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Duration      *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Error         *string                `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"`
	Label         string                 `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	NodeId        string                 `protobuf:"bytes,5,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeType      string                 `protobuf:"bytes,6,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	Output        *ExecutionStep_Output  `protobuf:"bytes,7,opt,name=output,proto3" json:"output,omitempty"`
	Status        string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	StepNumber    int32                  `protobuf:"varint,9,opt,name=step_number,json=stepNumber,proto3" json:"step_number,omitempty"`
	Timestamp     string                 `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecutionStep) Reset() {
	*x = ExecutionStep{}
	mi := &file_execute_workflow_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutionStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionStep) ProtoMessage() {}

func (x *ExecutionStep) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionStep.ProtoReflect.Descriptor instead.
func (*ExecutionStep) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *ExecutionStep) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ExecutionStep) GetDuration() *wrapperspb.Int64Value {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *ExecutionStep) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

func (x *ExecutionStep) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ExecutionStep) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ExecutionStep) GetNodeType() string {
	if x != nil {
		return x.NodeType
	}
	return ""
}

func (x *ExecutionStep) GetOutput() *ExecutionStep_Output {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *ExecutionStep) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ExecutionStep) GetStepNumber() int32 {
	if x != nil {
		return x.StepNumber
	}
	return 0
}

func (x *ExecutionStep) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

// *
// ExecutionResults contains a summary response to the client upon completion or failure of workflow execution.
type ExecutionResults struct {
	state         protoimpl.MessageState              `protogen:"open.v1"`
	EndTime       string                              `protobuf:"bytes,1,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	ExecutionId   string                              `protobuf:"bytes,2,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	Metadata      *ExecutionResults_ExecutionMetadata `protobuf:"bytes,3,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	StartTime     string                              `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Status        string                              `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Steps         []*ExecutionStep                    `protobuf:"bytes,6,rep,name=steps,proto3" json:"steps,omitempty"`
	TotalDuration *wrapperspb.Int64Value              `protobuf:"bytes,7,opt,name=total_duration,json=totalDuration,proto3,oneof" json:"total_duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecutionResults) Reset() {
	*x = ExecutionResults{}
	mi := &file_execute_workflow_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutionResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionResults) ProtoMessage() {}

func (x *ExecutionResults) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionResults.ProtoReflect.Descriptor instead.
func (*ExecutionResults) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{4}
}

func (x *ExecutionResults) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *ExecutionResults) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

func (x *ExecutionResults) GetMetadata() *ExecutionResults_ExecutionMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ExecutionResults) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ExecutionResults) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ExecutionResults) GetSteps() []*ExecutionStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *ExecutionResults) GetTotalDuration() *wrapperspb.Int64Value {
	if x != nil {
		return x.TotalDuration
	}
	return nil
}

type ExecuteWorkflowRequest_Condition struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operator      string                 `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Threshold     float64                `protobuf:"fixed64,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteWorkflowRequest_Condition) Reset() {
	*x = ExecuteWorkflowRequest_Condition{}
	mi := &file_execute_workflow_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteWorkflowRequest_Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteWorkflowRequest_Condition) ProtoMessage() {}

func (x *ExecuteWorkflowRequest_Condition) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteWorkflowRequest_Condition.ProtoReflect.Descriptor instead.
func (*ExecuteWorkflowRequest_Condition) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ExecuteWorkflowRequest_Condition) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *ExecuteWorkflowRequest_Condition) GetThreshold() float64 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

type ExecutionStep_Output struct {
	state           protoimpl.MessageState                `protogen:"open.v1"`
	ApiResponse     *ExecutionStep_Output_ApiResponse     `protobuf:"bytes,1,opt,name=api_response,json=apiResponse,proto3,oneof" json:"api_response,omitempty"`
	ConditionResult *ExecutionStep_Output_ConditionResult `protobuf:"bytes,2,opt,name=condition_result,json=conditionResult,proto3,oneof" json:"condition_result,omitempty"`
	Details         map[string]*structpb.Value            `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	EmailContent    *ExecutionStep_Output_EmailContent    `protobuf:"bytes,4,opt,name=email_content,json=emailContent,proto3,oneof" json:"email_content,omitempty"`
	FormData        *WorkflowFormData                     `protobuf:"bytes,5,opt,name=form_data,json=formData,proto3,oneof" json:"form_data,omitempty"`
	Message         string                                `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ExecutionStep_Output) Reset() {
	*x = ExecutionStep_Output{}
	mi := &file_execute_workflow_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutionStep_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionStep_Output) ProtoMessage() {}

func (x *ExecutionStep_Output) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionStep_Output.ProtoReflect.Descriptor instead.
func (*ExecutionStep_Output) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ExecutionStep_Output) GetApiResponse() *ExecutionStep_Output_ApiResponse {
	if x != nil {
		return x.ApiResponse
	}
	return nil
}

func (x *ExecutionStep_Output) GetConditionResult() *ExecutionStep_Output_ConditionResult {
	if x != nil {
		return x.ConditionResult
	}
	return nil
}

func (x *ExecutionStep_Output) GetDetails() map[string]*structpb.Value {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ExecutionStep_Output) GetEmailContent() *ExecutionStep_Output_EmailContent {
	if x != nil {
		return x.EmailContent
	}
	return nil
}

func (x *ExecutionStep_Output) GetFormData() *WorkflowFormData {
	if x != nil {
		return x.FormData
	}
	return nil
}

func (x *ExecutionStep_Output) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ExecutionStep_Output_ApiResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *structpb.Value        `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Endpoint      string                 `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Method        string                 `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	StatusCode    int32                  `protobuf:"varint,4,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecutionStep_Output_ApiResponse) Reset() {
	*x = ExecutionStep_Output_ApiResponse{}
	mi := &file_execute_workflow_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutionStep_Output_ApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionStep_Output_ApiResponse) ProtoMessage() {}

func (x *ExecutionStep_Output_ApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionStep_Output_ApiResponse.ProtoReflect.Descriptor instead.
func (*ExecutionStep_Output_ApiResponse) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{3, 0, 0}
}

func (x *ExecutionStep_Output_ApiResponse) GetData() *structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ExecutionStep_Output_ApiResponse) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ExecutionStep_Output_ApiResponse) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ExecutionStep_Output_ApiResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type ExecutionStep_Output_ConditionResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Expression    string                 `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
	Operator      string                 `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Result        bool                   `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	Temperature   float64                `protobuf:"fixed64,4,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Threshold     float64                `protobuf:"fixed64,5,opt,name=threshold,proto3" json:"threshold,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecutionStep_Output_ConditionResult) Reset() {
	*x = ExecutionStep_Output_ConditionResult{}
	mi := &file_execute_workflow_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutionStep_Output_ConditionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionStep_Output_ConditionResult) ProtoMessage() {}

func (x *ExecutionStep_Output_ConditionResult) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionStep_Output_ConditionResult.ProtoReflect.Descriptor instead.
func (*ExecutionStep_Output_ConditionResult) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{3, 0, 1}
}

func (x *ExecutionStep_Output_ConditionResult) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *ExecutionStep_Output_ConditionResult) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *ExecutionStep_Output_ConditionResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *ExecutionStep_Output_ConditionResult) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *ExecutionStep_Output_ConditionResult) GetThreshold() float64 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

type ExecutionStep_Output_EmailContent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Body          string                 `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Subject       string                 `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Timestamp     *string                `protobuf:"bytes,3,opt,name=timestamp,proto3,oneof" json:"timestamp,omitempty"`
	To            string                 `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecutionStep_Output_EmailContent) Reset() {
	*x = ExecutionStep_Output_EmailContent{}
	mi := &file_execute_workflow_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutionStep_Output_EmailContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionStep_Output_EmailContent) ProtoMessage() {}

func (x *ExecutionStep_Output_EmailContent) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionStep_Output_EmailContent.ProtoReflect.Descriptor instead.
func (*ExecutionStep_Output_EmailContent) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{3, 0, 2}
}

func (x *ExecutionStep_Output_EmailContent) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ExecutionStep_Output_EmailContent) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ExecutionStep_Output_EmailContent) GetTimestamp() string {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return ""
}

func (x *ExecutionStep_Output_EmailContent) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type ExecutionResults_ExecutionMetadata struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Environment     *string                `protobuf:"bytes,1,opt,name=environment,proto3,oneof" json:"environment,omitempty"`
	TriggeredBy     *string                `protobuf:"bytes,2,opt,name=triggered_by,json=triggeredBy,proto3,oneof" json:"triggered_by,omitempty"`
	WorkflowVersion *string                `protobuf:"bytes,3,opt,name=workflow_version,json=workflowVersion,proto3,oneof" json:"workflow_version,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ExecutionResults_ExecutionMetadata) Reset() {
	*x = ExecutionResults_ExecutionMetadata{}
	mi := &file_execute_workflow_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutionResults_ExecutionMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionResults_ExecutionMetadata) ProtoMessage() {}

func (x *ExecutionResults_ExecutionMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_execute_workflow_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionResults_ExecutionMetadata.ProtoReflect.Descriptor instead.
func (*ExecutionResults_ExecutionMetadata) Descriptor() ([]byte, []int) {
	return file_execute_workflow_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ExecutionResults_ExecutionMetadata) GetEnvironment() string {
	if x != nil && x.Environment != nil {
		return *x.Environment
	}
	return ""
}

func (x *ExecutionResults_ExecutionMetadata) GetTriggeredBy() string {
	if x != nil && x.TriggeredBy != nil {
		return *x.TriggeredBy
	}
	return ""
}

func (x *ExecutionResults_ExecutionMetadata) GetWorkflowVersion() string {
	if x != nil && x.WorkflowVersion != nil {
		return *x.WorkflowVersion
	}
	return ""
}

var File_execute_workflow_proto protoreflect.FileDescriptor

const file_execute_workflow_proto_rawDesc = "" +
	"\n" +
	"\x16execute_workflow.proto\x12\x05proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\n" +
	"node.proto\x1a\n" +
	"edge.proto\"\x8a\x01\n" +
	"\x10WorkflowFormData\x12\x12\n" +
	"\x04city\x18\x01 \x01(\tR\x04city\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x1a\n" +
	"\boperator\x18\x04 \x01(\tR\boperator\x12\x1c\n" +
	"\tthreshold\x18\x05 \x01(\x01R\tthreshold\"\xc4\x02\n" +
	"\x16ExecuteWorkflowRequest\x12E\n" +
	"\tcondition\x18\x01 \x01(\v2'.proto.ExecuteWorkflowRequest.ConditionR\tcondition\x124\n" +
	"\tform_data\x18\x02 \x01(\v2\x17.proto.WorkflowFormDataR\bformData\x122\n" +
	"\x0eworkflow_edges\x18\x03 \x03(\v2\v.proto.EdgeR\rworkflowEdges\x122\n" +
	"\x0eworkflow_nodes\x18\x04 \x03(\v2\v.proto.NodeR\rworkflowNodes\x1aE\n" +
	"\tCondition\x12\x1a\n" +
	"\boperator\x18\x01 \x01(\tR\boperator\x12\x1c\n" +
	"\tthreshold\x18\x02 \x01(\x01R\tthreshold\"(\n" +
	"\fExecuteError\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\xdf\n" +
	"\n" +
	"\rExecutionStep\x12 \n" +
	"\vdescription\x18\x01 \x01(\tR\vdescription\x127\n" +
	"\bduration\x18\x02 \x01(\v2\x1b.google.protobuf.Int64ValueR\bduration\x12\x19\n" +
	"\x05error\x18\x03 \x01(\tH\x00R\x05error\x88\x01\x01\x12\x14\n" +
	"\x05label\x18\x04 \x01(\tR\x05label\x12\x17\n" +
	"\anode_id\x18\x05 \x01(\tR\x06nodeId\x12\x1b\n" +
	"\tnode_type\x18\x06 \x01(\tR\bnodeType\x123\n" +
	"\x06output\x18\a \x01(\v2\x1b.proto.ExecutionStep.OutputR\x06output\x12\x16\n" +
	"\x06status\x18\b \x01(\tR\x06status\x12\x1f\n" +
	"\vstep_number\x18\t \x01(\x05R\n" +
	"stepNumber\x12\x1c\n" +
	"\ttimestamp\x18\n" +
	" \x01(\tR\ttimestamp\x1a\xf5\a\n" +
	"\x06Output\x12O\n" +
	"\fapi_response\x18\x01 \x01(\v2'.proto.ExecutionStep.Output.ApiResponseH\x00R\vapiResponse\x88\x01\x01\x12[\n" +
	"\x10condition_result\x18\x02 \x01(\v2+.proto.ExecutionStep.Output.ConditionResultH\x01R\x0fconditionResult\x88\x01\x01\x12B\n" +
	"\adetails\x18\x03 \x03(\v2(.proto.ExecutionStep.Output.DetailsEntryR\adetails\x12R\n" +
	"\remail_content\x18\x04 \x01(\v2(.proto.ExecutionStep.Output.EmailContentH\x02R\femailContent\x88\x01\x01\x129\n" +
	"\tform_data\x18\x05 \x01(\v2\x17.proto.WorkflowFormDataH\x03R\bformData\x88\x01\x01\x12\x18\n" +
	"\amessage\x18\x06 \x01(\tR\amessage\x1a\x8e\x01\n" +
	"\vApiResponse\x12*\n" +
	"\x04data\x18\x01 \x01(\v2\x16.google.protobuf.ValueR\x04data\x12\x1a\n" +
	"\bendpoint\x18\x02 \x01(\tR\bendpoint\x12\x16\n" +
	"\x06method\x18\x03 \x01(\tR\x06method\x12\x1f\n" +
	"\vstatus_code\x18\x04 \x01(\x05R\n" +
	"statusCode\x1a\xa5\x01\n" +
	"\x0fConditionResult\x12\x1e\n" +
	"\n" +
	"expression\x18\x01 \x01(\tR\n" +
	"expression\x12\x1a\n" +
	"\boperator\x18\x02 \x01(\tR\boperator\x12\x16\n" +
	"\x06result\x18\x03 \x01(\bR\x06result\x12 \n" +
	"\vtemperature\x18\x04 \x01(\x01R\vtemperature\x12\x1c\n" +
	"\tthreshold\x18\x05 \x01(\x01R\tthreshold\x1a}\n" +
	"\fEmailContent\x12\x12\n" +
	"\x04body\x18\x01 \x01(\tR\x04body\x12\x18\n" +
	"\asubject\x18\x02 \x01(\tR\asubject\x12!\n" +
	"\ttimestamp\x18\x03 \x01(\tH\x00R\ttimestamp\x88\x01\x01\x12\x0e\n" +
	"\x02to\x18\x04 \x01(\tR\x02toB\f\n" +
	"\n" +
	"_timestamp\x1aR\n" +
	"\fDetailsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12,\n" +
	"\x05value\x18\x02 \x01(\v2\x16.google.protobuf.ValueR\x05value:\x028\x01B\x0f\n" +
	"\r_api_responseB\x13\n" +
	"\x11_condition_resultB\x10\n" +
	"\x0e_email_contentB\f\n" +
	"\n" +
	"_form_dataB\b\n" +
	"\x06_error\"\xb3\x04\n" +
	"\x10ExecutionResults\x12\x19\n" +
	"\bend_time\x18\x01 \x01(\tR\aendTime\x12!\n" +
	"\fexecution_id\x18\x02 \x01(\tR\vexecutionId\x12J\n" +
	"\bmetadata\x18\x03 \x01(\v2).proto.ExecutionResults.ExecutionMetadataH\x00R\bmetadata\x88\x01\x01\x12\x1d\n" +
	"\n" +
	"start_time\x18\x04 \x01(\tR\tstartTime\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\x12*\n" +
	"\x05steps\x18\x06 \x03(\v2\x14.proto.ExecutionStepR\x05steps\x12G\n" +
	"\x0etotal_duration\x18\a \x01(\v2\x1b.google.protobuf.Int64ValueH\x01R\rtotalDuration\x88\x01\x01\x1a\xc8\x01\n" +
	"\x11ExecutionMetadata\x12%\n" +
	"\venvironment\x18\x01 \x01(\tH\x00R\venvironment\x88\x01\x01\x12&\n" +
	"\ftriggered_by\x18\x02 \x01(\tH\x01R\vtriggeredBy\x88\x01\x01\x12.\n" +
	"\x10workflow_version\x18\x03 \x01(\tH\x02R\x0fworkflowVersion\x88\x01\x01B\x0e\n" +
	"\f_environmentB\x0f\n" +
	"\r_triggered_byB\x13\n" +
	"\x11_workflow_versionB\v\n" +
	"\t_metadataB\x11\n" +
	"\x0f_total_durationB\vZ\t/workflowb\x06proto3"

var (
	file_execute_workflow_proto_rawDescOnce sync.Once
	file_execute_workflow_proto_rawDescData []byte
)

func file_execute_workflow_proto_rawDescGZIP() []byte {
	file_execute_workflow_proto_rawDescOnce.Do(func() {
		file_execute_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_execute_workflow_proto_rawDesc), len(file_execute_workflow_proto_rawDesc)))
	})
	return file_execute_workflow_proto_rawDescData
}

var file_execute_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_execute_workflow_proto_goTypes = []any{
	(*WorkflowFormData)(nil),                     // 0: proto.WorkflowFormData
	(*ExecuteWorkflowRequest)(nil),               // 1: proto.ExecuteWorkflowRequest
	(*ExecuteError)(nil),                         // 2: proto.ExecuteError
	(*ExecutionStep)(nil),                        // 3: proto.ExecutionStep
	(*ExecutionResults)(nil),                     // 4: proto.ExecutionResults
	(*ExecuteWorkflowRequest_Condition)(nil),     // 5: proto.ExecuteWorkflowRequest.Condition
	(*ExecutionStep_Output)(nil),                 // 6: proto.ExecutionStep.Output
	(*ExecutionStep_Output_ApiResponse)(nil),     // 7: proto.ExecutionStep.Output.ApiResponse
	(*ExecutionStep_Output_ConditionResult)(nil), // 8: proto.ExecutionStep.Output.ConditionResult
	(*ExecutionStep_Output_EmailContent)(nil),    // 9: proto.ExecutionStep.Output.EmailContent
	nil, // 10: proto.ExecutionStep.Output.DetailsEntry
	(*ExecutionResults_ExecutionMetadata)(nil), // 11: proto.ExecutionResults.ExecutionMetadata
	(*Edge)(nil),                  // 12: proto.Edge
	(*Node)(nil),                  // 13: proto.Node
	(*wrapperspb.Int64Value)(nil), // 14: google.protobuf.Int64Value
	(*structpb.Value)(nil),        // 15: google.protobuf.Value
}
var file_execute_workflow_proto_depIdxs = []int32{
	5,  // 0: proto.ExecuteWorkflowRequest.condition:type_name -> proto.ExecuteWorkflowRequest.Condition
	0,  // 1: proto.ExecuteWorkflowRequest.form_data:type_name -> proto.WorkflowFormData
	12, // 2: proto.ExecuteWorkflowRequest.workflow_edges:type_name -> proto.Edge
	13, // 3: proto.ExecuteWorkflowRequest.workflow_nodes:type_name -> proto.Node
	14, // 4: proto.ExecutionStep.duration:type_name -> google.protobuf.Int64Value
	6,  // 5: proto.ExecutionStep.output:type_name -> proto.ExecutionStep.Output
	11, // 6: proto.ExecutionResults.metadata:type_name -> proto.ExecutionResults.ExecutionMetadata
	3,  // 7: proto.ExecutionResults.steps:type_name -> proto.ExecutionStep
	14, // 8: proto.ExecutionResults.total_duration:type_name -> google.protobuf.Int64Value
	7,  // 9: proto.ExecutionStep.Output.api_response:type_name -> proto.ExecutionStep.Output.ApiResponse
	8,  // 10: proto.ExecutionStep.Output.condition_result:type_name -> proto.ExecutionStep.Output.ConditionResult
	10, // 11: proto.ExecutionStep.Output.details:type_name -> proto.ExecutionStep.Output.DetailsEntry
	9,  // 12: proto.ExecutionStep.Output.email_content:type_name -> proto.ExecutionStep.Output.EmailContent
	0,  // 13: proto.ExecutionStep.Output.form_data:type_name -> proto.WorkflowFormData
	15, // 14: proto.ExecutionStep.Output.ApiResponse.data:type_name -> google.protobuf.Value
	15, // 15: proto.ExecutionStep.Output.DetailsEntry.value:type_name -> google.protobuf.Value
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_execute_workflow_proto_init() }
func file_execute_workflow_proto_init() {
	if File_execute_workflow_proto != nil {
		return
	}
	file_node_proto_init()
	file_edge_proto_init()
	file_execute_workflow_proto_msgTypes[3].OneofWrappers = []any{}
	file_execute_workflow_proto_msgTypes[4].OneofWrappers = []any{}
	file_execute_workflow_proto_msgTypes[6].OneofWrappers = []any{}
	file_execute_workflow_proto_msgTypes[9].OneofWrappers = []any{}
	file_execute_workflow_proto_msgTypes[11].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_execute_workflow_proto_rawDesc), len(file_execute_workflow_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_execute_workflow_proto_goTypes,
		DependencyIndexes: file_execute_workflow_proto_depIdxs,
		MessageInfos:      file_execute_workflow_proto_msgTypes,
	}.Build()
	File_execute_workflow_proto = out.File
	file_execute_workflow_proto_goTypes = nil
	file_execute_workflow_proto_depIdxs = nil
}
