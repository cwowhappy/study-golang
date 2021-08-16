package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Operator interface {
	Name() string
	Trigger() <-chan time.Time
	Refresh() error
}

type OperatorManager interface {
	Register(operator Operator) error
	Start(ctx context.Context) error
}

func main() {
	opManager := GetOperatorManager()
	ctx, ctxCancel := context.WithCancel(context.Background())
	for _, i := range []int{1, 2, 3} {
		op := GetOperator(fmt.Sprintf("op-%d", i))
		_ = opManager.Register(op)
	}

	_ = opManager.Start(ctx)

	fmt.Printf("OperaterManager start\n")
	time.Sleep(1 * time.Minute)
	ctxCancel()
}

type operatorImpl struct {
	name string
	refreshInterval time.Duration
	timer *time.Timer
}

type operatorManagerImpl struct {
	operators map[string]Operator
	status uint32
}

func GetOperator(name string) Operator {
	return &operatorImpl{
		name: name,
		refreshInterval: 10 * time.Second,
		timer: time.NewTimer(10 * time.Second),
	}
}

func GetOperatorManager() OperatorManager {
	return &operatorManagerImpl{
		operators: make(map[string]Operator),
		status: 1,
	}
}

func (operator *operatorImpl) Name() string {
	return operator.name
}

func (operator *operatorImpl) Refresh() error {
	fmt.Printf("func[%v.Refresh] start\n", operator)
	time.Sleep(2 * time.Second)
	operator.timer.Reset(operator.refreshInterval)
	fmt.Printf("func[%v.Refresh] end\n", operator)
	return nil
}

func (operator *operatorImpl) Trigger() <-chan time.Time {
	fmt.Printf("func[%v.Trigger] execute\n", operator)
	return operator.timer.C
}

func (opManager * operatorManagerImpl) Register(operator Operator) error {
	if operator == nil {
		return errors.New("operator is nil")
	}

	if _, ok := opManager.operators[operator.Name()]; ok {
		return errors.New("operator exists")
	}
	opManager.operators[operator.Name()] = operator
	return nil
}

func (opManager * operatorManagerImpl) Start(ctx context.Context) error {
	for name, operator := range opManager.operators {
		go func(n string, op Operator) {
			fmt.Printf("operator [%s] start\n", n)
			flag := true
			for flag {
				select {
				case <-op.Trigger():
					_ = op.Refresh()
				case <-ctx.Done():
					flag = false
				}
			}
			fmt.Printf("context is done\n")
		}(name, operator)
	}
	return nil
}


