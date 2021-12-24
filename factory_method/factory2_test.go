package factory_method

import (
	"testing"
)

func TestGetActionsSet(t *testing.T) {
	catF := &CatActions{}
	catActionsSet := GetActionsSet(catF, "Mr.cat")
	catActionsSet.Jump()
	catActionsSet.Run()
	catActionsSet.DoSomething()

	dogF := &DogActions{}
	dogActionsSet := GetActionsSet(dogF, "Mr.dog")
	dogActionsSet.Jump()
	dogActionsSet.Run()
	dogActionsSet.DoSomething()
}

func TestGetActionsSet2(t *testing.T) {
	catActionsSet2 := GetActionsSet2(0, "Lady cat")
	catActionsSet2.Jump()
	catActionsSet2.Run()
	catActionsSet2.DoSomething()

	dogActionsSet2 := GetActionsSet2(1, "Lady dog")
	dogActionsSet2.Jump()
	dogActionsSet2.Run()
	dogActionsSet2.DoSomething()
}
