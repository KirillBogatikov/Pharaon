package service

import "log"

type Task func() error

type Transaction struct {
	forCancel []Task
}

func OpenTransaction() *Transaction {
	return &Transaction{make([]Task, 0)}
}

func (t *Transaction) Exec(task Task, cancel Task) error {
	err := task()
	if err != nil {
		t.forCancel = append(t.forCancel, cancel)
	}

	log.Printf("WARNING: Transaction task #%d failed: %s\n", len(t.forCancel)+1, err)

	return err
}

func (t *Transaction) Cancel() (errors []error) {
	for i, task := range t.forCancel {
		err := task()
		if err != nil {
			log.Printf("Can not cancel transaction task #%d: %s\n", i, err)
			errors = append(errors, err)
		}
	}

	return
}
