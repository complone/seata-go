package constant

const (
	// region State Types
	StateTypeServiceTask          string = "ServiceTask"
	StateTypeChoice               string = "Choice"
	StateTypeFail                 string = "Fail"
	StateTypeSucceed              string = "Succeed"
	StateTypeCompensationTrigger  string = "CompensationTrigger"
	StateTypeSubStateMachine      string = "SubStateMachine"
	StateTypeCompensateSubMachine string = "CompensateSubMachine"
	StateTypeScriptTask           string = "ScriptTask"
	StateTypeLoopStart            string = "LoopStart"
	// end region

	// region Service Types
	ServiceTypeGRPC string = "GRPC"
	// end region

	// region System Variables
	VarNameOutputParams                  string = "outputParams"
	VarNameProcessType                   string = "_ProcessType_"
	VarNameOperationName                 string = "_operation_name_"
	OperationNameStart                   string = "start"
	OperationNameCompensate              string = "compensate"
	VarNameAsyncCallback                 string = "_async_callback_"
	VarNameCurrentExceptionRoute         string = "_current_exception_route_"
	VarNameIsExceptionNotCatch           string = "_is_exception_not_catch_"
	VarNameSubMachineParentId            string = "_sub_machine_parent_id_"
	VarNameCurrentChoice                 string = "_current_choice_"
	VarNameStateMachineInst              string = "_current_statemachine_instance_"
	VarNameStateMachine                  string = "_current_statemachine_"
	VarNameStateMachineEngine            string = "_current_statemachine_engine_"
	VarNameStateMachineConfig            string = "_statemachine_config_"
	VarNameStateMachineContext           string = "context"
	VarNameIsAsyncExecution              string = "_is_async_execution_"
	VarNameStateInst                     string = "_current_state_instance_"
	VarNameBusinesskey                   string = "_business_key_"
	VarNameParentId                      string = "_parent_id_"
	VarNameCurrentException              string = "currentException"
	CompensateSubMachineStateNamePrefix  string = "_compensate_sub_machine_state_"
	DefaultScriptType                    string = "groovy"
	VarNameSyncExeStack                  string = "_sync_execution_stack_"
	VarNameInputParams                   string = "inputParams"
	VarNameIsLoopState                   string = "_is_loop_state_"
	VarNameCurrentCompensateTriggerState string = "_is_compensating_"
	VarNameCurrentCompensationHolder     string = "_current_compensation_holder_"
	VarNameFirstCompensationStateStarted string = "_first_compensation_state_started"
	VarNameCurrentLoopContextHolder      string = "_current_loop_context_holder_"
	// TODO: this lock in process context only has one, try to add more to add concurrent
	VarNameProcessContextMutexLock string = "_current_context_mutex_lock"
	VarNameFailEndStateFlag        string = "_fail_end_state_flag_"
	// end region

	// region of loop
	LoopCounter                string = "loopCounter"
	LoopSemaphore              string = "loopSemaphore"
	LoopResult                 string = "loopResult"
	NumberOfInstances          string = "nrOfInstances"
	NumberOfActiveInstances    string = "nrOfActiveInstances"
	NumberOfCompletedInstances string = "nrOfCompletedInstances"
	// end region

	// region others
	SeqEntityStateMachineInst string = "STATE_MACHINE_INST"
	SeqEntityStateInst        string = "STATE_INST"
	OperationNameForward      string = "forward"
	LoopStateNamePattern      string = "-loop-"
	// end region

	SeperatorParentId string = ":"
)