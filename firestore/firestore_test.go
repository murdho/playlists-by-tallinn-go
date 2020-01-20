// +build integration

package firestore

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc/status"
)

func TestMain(m *testing.M) {
	if _, ok := os.LookupEnv("FIRESTORE_EMULATOR_HOST"); !ok {
		fmt.Println("ERROR\tFirestore emulator and environment variable FIRESTORE_EMULATOR_HOST are required for `firestore` tests.")
		os.Exit(1)
	}

	os.Exit(m.Run())
}

type dat struct {
	X string `firestore:"x"`
}

func TestNew(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fs, err := New(ctx, Project("a"), Collection("b"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	//
	// err = fs.Set(ctx, "a", dat{"b"})
	// if err != nil {
	// 	t.Error(err)
	// 	t.FailNow()
	// }

	var d dat
	err = fs.Get(ctx, &d, "c")
	if err != nil {
		fmt.Println(status.Code(err))
		t.Error(err)
		t.FailNow()
	}

	fmt.Printf("%+v\n", d)

}
