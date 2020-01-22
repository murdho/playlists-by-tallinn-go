// +build integration

package firestore

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

const contextTimeout = time.Second

func TestMain(m *testing.M) {
	if _, ok := os.LookupEnv("FIRESTORE_EMULATOR_HOST"); !ok {
		fmt.Println("ERROR\tFirestore emulator and environment variable FIRESTORE_EMULATOR_HOST are required for `firestore` tests.")
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	fs, err := New(ctx, Project("a"), Collection("b"))
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}
	defer func() {
		if err := fs.Close(); err != nil {
			t.Errorf("unexpected close error: %+v", err)
		}
	}()

	if fs.client == nil {
		t.Errorf("firestore client:\ngot  %+v\nwant %+v", fs.client, "(not nil)")
	}

	if fs.projectID != "a" {
		t.Errorf("project ID:\ngot  %+v\nwant %+v", fs.projectID, "a")
	}

	if fs.collection != "b" {
		t.Errorf("collection:\ngot  %+v\nwant %+v", fs.collection, "b")
	}
}

func TestFirestore_Get(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	fs, err := New(ctx, Project("a"), Collection("b"))
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}
	defer func() {
		if err := fs.Close(); err != nil {
			t.Errorf("unexpected close error: %+v", err)
		}
	}()

	name := strconv.Itoa(int(time.Now().Unix()))
	setData := struct {
		Name string `firestore:"name"`
	}{
		Name: name,
	}

	err = fs.Set(ctx, "x", setData)
	if err != nil {
		t.Errorf("unexpected set error: %+v", err)
	}

	var getData struct {
		Name string `firestore:"name"`
	}

	err = fs.Get(ctx, &getData, "x")
	if err != nil {
		t.Errorf("unexpected get error: %+v", err)
	}

	if getData.Name != name {
		t.Errorf("get data name:\ngot  %+v\nwant %+v", getData.Name, name)
	}
}

func TestFirestore_GetErrNotFound(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	fs, err := New(ctx, Project("a"), Collection("b"))
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}
	defer func() {
		if err := fs.Close(); err != nil {
			t.Errorf("unexpected close error: %+v", err)
		}
	}()

	var data struct{}
	err = fs.Get(ctx, &data, "z")
	if !reflect.DeepEqual(err, ErrNotFound) {
		t.Errorf("not found error:\ngot  %+v\nwant %+v", err, ErrNotFound)
	}
}

func TestFirestore_GetErr(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	fs, err := New(ctx, Project("a"), Collection("b"))
	if err != nil {
		t.Errorf("unexpected new error: %+v", err)
	}

	err = fs.Close()
	if err != nil {
		t.Errorf("unexpected close error: %+v", err)
	}

	var data struct{}
	err = fs.Get(ctx, &data, "z")
	if err == nil {
		t.Errorf("get error:\ngot  %+v\nwant %+v", err, nil)
	}

	if err == ErrNotFound {
		t.Errorf("get error:\ngot  %+v\nwant %+v", err, ErrNotFound)
	}
}
