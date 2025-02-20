package queue

import (
	"errors"
	"testing"
)

/*
implement a queue with maximum capacity.
*/

func TestQueue(t *testing.T) {

	q := NewQueue(5)

	/* --------------- */

	_, err := q.Dequeue()
	if err == nil {
		t.Fatal("expected error because of Dequeue operation on empty queue")
	} else if !errors.Is(err, ErrQueueIsEmpty) {
		t.Fatalf("expected error %q; got %q", ErrQueueIsEmpty.Error(), err.Error())
	}

	/* --------------- */

	err = q.Enqueue(1)
	if err != nil {
		t.Fatal(err)
	}

	err = q.Enqueue(2)
	if err != nil {
		t.Fatal(err)
	}

	err = q.Enqueue(3)
	if err != nil {
		t.Fatal(err)
	}

	err = q.Enqueue(4)
	if err != nil {
		t.Fatal(err)
	}

	err = q.Enqueue(5)
	if err != nil {
		t.Fatal(err)
	}

	/* --------------- */

	err = q.Enqueue(5)
	if err == nil {
		t.Fatal("expected error because of Enqueue operation on full queue")
	} else if !errors.Is(err, ErrQueueIsFull) {
		t.Fatalf("expected error %q; got %q", ErrQueueIsFull.Error(), err.Error())
	}

	/* --------------- */

	elem, err := q.Dequeue()
	if err != nil {
		t.Fatal(err)
	} else if value := 1; elem != value {
		t.Fatalf("expected to Dequeue value %d; got %d", value, elem)
	}

	elem, err = q.Dequeue()
	if err != nil {
		t.Fatal(err)
	} else if value := 2; elem != value {
		t.Fatalf("expected to Dequeue value %d; got %d", value, elem)
	}

	elem, err = q.Dequeue()
	if err != nil {
		t.Fatal(err)
	} else if value := 3; elem != value {
		t.Fatalf("expected to Dequeue value %d; got %d", value, elem)
	}

	elem, err = q.Dequeue()
	if err != nil {
		t.Fatal(err)
	} else if value := 4; elem != value {
		t.Fatalf("expected to Dequeue value %d; got %d", value, elem)
	}

	elem, err = q.Dequeue()
	if err != nil {
		t.Fatal(err)
	} else if value := 5; elem != value {
		t.Fatalf("expected to Dequeue value %d; got %d", value, elem)
	}

	/* --------------- */

	_, err = q.Dequeue()
	if err == nil {
		t.Fatal("expected error because of Dequeue operation on empty queue")
	} else if !errors.Is(err, ErrQueueIsEmpty) {
		t.Fatalf("expected error %q; got %q", ErrQueueIsEmpty.Error(), err.Error())
	}

	/* --------------- */

}
