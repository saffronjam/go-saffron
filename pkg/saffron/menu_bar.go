package saffron

import "github.com/saffronjam/cimgui-go/imgui"

type MenuBarMenu struct {
	Title    string
	RenderUI func()
}

type MenuBar struct {
	Menus []MenuBarMenu
}

func NewMenuBar() *MenuBar {
	return &MenuBar{
		Menus: []MenuBarMenu{},
	}
}

func (mb *MenuBar) AddMenu(title string, renderUI func()) {
	mb.Menus = append(mb.Menus, MenuBarMenu{
		Title:    title,
		RenderUI: renderUI,
	})
}

func (mb *MenuBar) RemoveMenu(title string) {
	for i, menu := range mb.Menus {
		if menu.Title == title {
			mb.Menus = append(mb.Menus[:i], mb.Menus[i+1:]...)
			return
		}
	}
}

func (mb *MenuBar) RenderUI() {
	if imgui.BeginMenuBar() {
		for _, menu := range mb.Menus {
			if imgui.BeginMenu(menu.Title) {
				menu.RenderUI()
				imgui.EndMenu()
			}
		}
		imgui.EndMenuBar()
	}
}
