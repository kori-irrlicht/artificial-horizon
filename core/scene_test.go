package core

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type scene struct {
	name     string
	isInput  bool
	isUpdate bool
	isRender bool
}

func (s *scene) Update()      { s.isUpdate = true }
func (s *scene) Input()       { s.isInput = true }
func (s *scene) Render()      { s.isRender = true }
func (s *scene) Name() string { return s.name }

func TestSceneManager(t *testing.T) {
	Convey("A new sceneManager", t, func() {
		sm := NewSceneManager("abc").(*sceneManager)
		Convey("Setting a name", func() {
			So(sm.Name(), ShouldEqual, "abc")
		})
		Convey("Changing to a non-existing scene", func() {
			err := sm.Change("nope")
			So(err, ShouldNotBeNil)
		})
		Convey("And a new scene", func() {
			s1 := scene{name: "s1"}
			Convey("Adding the scene", func() {
				sm.Add(&s1)
				So(sm.scenes[s1.Name()], ShouldEqual, &s1)
				Convey("And changing to it", func() {
					err := sm.Change(s1.Name())
					So(sm.current, ShouldEqual, &s1)
					So(err, ShouldBeNil)
					Convey("Input is called", func() {
						sm.Input()
						So(s1.isInput, ShouldBeTrue)
					})
					Convey("Update is called", func() {
						sm.Update()
						So(s1.isUpdate, ShouldBeTrue)
					})
					Convey("Render is called", func() {
						sm.Render()
						So(s1.isRender, ShouldBeTrue)
					})
				})
				Convey("Adding another scene", func() {
					s2 := scene{name: "s2"}
					sm.Add(&s2)
					So(sm.scenes[s2.Name()], ShouldEqual, &s2)
					Convey("And changing to it", func() {
						sm.Change(s2.Name())
						So(sm.current, ShouldEqual, &s2)
					})
				})
			})
		})
	})
}
