package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type casetViewKeymap struct {
	LikeCast    key.Binding
	ViewProfile key.Binding
	ViewChannel key.Binding
	ViewParent  key.Binding
	Comment     key.Binding
	OpenCast    key.Binding
}

func (k casetViewKeymap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.LikeCast,
		k.ViewParent,
		k.Comment,
	}
}
func (k casetViewKeymap) All() []key.Binding {
	return []key.Binding{
		k.LikeCast,
		k.ViewProfile,
		k.ViewChannel,
		k.ViewParent,
		k.Comment,
		k.OpenCast,
	}
}

func (k casetViewKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.All(),
	}
}

func (k casetViewKeymap) HandleMsg(c *CastView, msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, k.LikeCast):
		return c.LikeCast()
	case key.Matches(msg, k.ViewProfile):
		return c.ViewProfile()
	case key.Matches(msg, k.ViewChannel):
		return c.ViewChannel()
	case key.Matches(msg, k.ViewParent):
		return c.ViewParent()
	case key.Matches(msg, k.Comment):
		c.Reply()
		return noOp()
	case key.Matches(msg, k.OpenCast):
		return c.OpenCast()
	}
	return nil
}

var CastViewKeyMap = casetViewKeymap{
	LikeCast: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "like cast"),
	),
	ViewProfile: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "view profile"),
	),
	ViewChannel: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "view channel"),
	),
	ViewParent: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "view parent"),
	),
	Comment: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reply"),
	),
	OpenCast: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "open in browser"),
	),
}

type feedKeymap struct {
	ViewCast    key.Binding
	LikeCast    key.Binding
	ViewProfile key.Binding
	ViewChannel key.Binding
	OpenCast    key.Binding
}

func (k feedKeymap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.LikeCast,
		k.ViewProfile,
		k.ViewChannel,
	}
}
func (k feedKeymap) All() []key.Binding {
	return []key.Binding{
		k.ViewCast,
		k.LikeCast,
		k.ViewProfile,
		k.ViewChannel,
		k.OpenCast,
	}
}

func (k feedKeymap) HandleMsg(f *FeedView, msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, k.ViewCast):
		log.Println("ViewCast")
		return f.SelectCurrentItem()
	case key.Matches(msg, k.LikeCast):
		log.Println("LikeCast")
		return f.LikeCurrentItem()
	case key.Matches(msg, k.ViewProfile):
		log.Println("ViewProfile")
		return f.ViewCurrentProfile()
	case key.Matches(msg, k.ViewChannel):
		log.Println("ViewChannel")
		return f.ViewCurrentChannel()
	case key.Matches(msg, k.OpenCast):
		log.Println("OpenCast")
		return f.OpenCurrentItem()
	}
	return nil
}

var FeedKeyMap = feedKeymap{
	ViewCast: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "view cast"),
	),
	LikeCast: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "like cast"),
	),
	ViewProfile: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "view profile"),
	),
	ViewChannel: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "view channel"),
	),
	OpenCast: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "open in browser "),
	),
}

type navKeymap struct {
	Feed key.Binding

	Publish                 key.Binding
	QuickSelect             key.Binding
	Help                    key.Binding
	ToggleSidebarFocus      key.Binding
	ToggleSidebarVisibility key.Binding
	Previous                key.Binding
	ViewNotifications       key.Binding
}

func (k navKeymap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Feed,
		k.QuickSelect,
		k.ViewNotifications,
		k.Help,
	}
}

func (k navKeymap) All() []key.Binding {
	return []key.Binding{
		k.Feed, k.QuickSelect,
		k.Publish,
		k.ViewNotifications,
		k.Previous,
		k.Help,
		k.ToggleSidebarFocus, k.ToggleSidebarVisibility,
	}
}

var NavKeyMap = navKeymap{
	Feed: key.NewBinding(
		key.WithKeys("F", "1"),
		key.WithHelp("F/1", "feed"),
	),
	Publish: key.NewBinding(
		key.WithKeys("P"),
		key.WithHelp("P", "publish cast"),
	),
	QuickSelect: key.NewBinding(
		key.WithKeys("ctrl+k"),
		key.WithHelp("ctrl+k", "quick select"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	ToggleSidebarFocus: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "toggle sidebar focus"),
	),
	ToggleSidebarVisibility: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "toggle sidebar focus"),
	),
	Previous: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "focus previous"),
	),
	ViewNotifications: key.NewBinding(
		key.WithKeys("N"),
		key.WithHelp("N", "view notifications"),
	),
}

func (k navKeymap) HandleMsg(a *App, msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, k.Feed):
		// TODO cleanup
		// reset params for user's feed
		var cmd tea.Cmd
		a.SetNavName("feed")
		a.sidebar.SetActive(false)
		return tea.Sequence(cmd, a.FocusFeed())

	case key.Matches(msg, k.Publish):
		a.FocusPublish()
		return noOp()

	case key.Matches(msg, k.QuickSelect):
		a.FocusQuickSelect()
		return nil

	case key.Matches(msg, k.Help):
		a.FocusHelp()

	case key.Matches(msg, k.ViewNotifications):
		log.Println("ViewNotifications")
		return a.FocusNotifications()

	case key.Matches(msg, k.Previous):
		return a.FocusPrev()

	case key.Matches(msg, k.ToggleSidebarVisibility):
		if a.showSidebar {
			a.showSidebar = false
			a.sidebar.SetActive(false)
			return noOp()
		}
		a.showSidebar = true
		a.sidebar.SetActive(true)
		return noOp()

	case key.Matches(msg, k.ToggleSidebarFocus):
		if a.quickSelect.Active() {
			_, cmd := a.quickSelect.Update(msg)
			return cmd
		}
		if !a.showSidebar {
			return nil
		}
		a.sidebar.SetActive(!a.sidebar.Active())
	}

	return nil
}

type kmap struct {
	nav  navKeymap
	feed feedKeymap
	cast casetViewKeymap
}

var GlobalKeyMap = kmap{
	nav:  NavKeyMap,
	feed: FeedKeyMap,
	cast: CastViewKeyMap,
}

func (k kmap) ShortHelp() []key.Binding {
	return append(k.nav.ShortHelp(), k.feed.ShortHelp()...)
}

func (k kmap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.nav.All(),
		k.feed.All(),
		k.cast.All(),
	}
}

func noOp() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}
