package task

import (
	"fmt"
)

type Response []error

func (r Response) Success() bool {
	err := r.Error()
	return err == nil
}

func (r Response) Error() error {
	for i, err := range r {
		if err != nil {
			return fmt.Errorf("task:%d error:%s", i, err)
		}
	}
	return nil
}
