// Code generated by "stringer -type JudgeResult"; DO NOT EDIT

package model

import "fmt"

const _JudgeResult_name = "AcceptCompileErrorWrongAnswerTimeLimitExceededMemoryLimitExceededUnhandledHandlingRuntimeErrorPresentationErrorPanicError"

var _JudgeResult_index = [...]uint8{0, 6, 18, 29, 46, 65, 74, 82, 94, 111, 121}

func (i JudgeResult) String() string {
	if i < 0 || i >= JudgeResult(len(_JudgeResult_index)-1) {
		return fmt.Sprintf("JudgeResult(%d)", i)
	}
	return _JudgeResult_name[_JudgeResult_index[i]:_JudgeResult_index[i+1]]
}