package selector

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gogf/gf/container/gtree"
	gfutil "github.com/gogf/gf/util/gutil"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
)

type ItemList struct {
	List *gtree.RedBlackTree
}

func NewItemList() (li *ItemList) {
	li = &ItemList{
		List: gtree.NewRedBlackTree(gfutil.ComparatorString),
	}
	return
}

func (that *ItemList) Add(key string, value interface{}) {
	that.List.Set(key, value)
}

func (that *ItemList) Remove(key string) interface{} {
	return that.List.Remove(key)
}

func (that *ItemList) Get(key string) interface{} {
	return that.List.Get(key)
}

func (that *ItemList) Clear() {
	that.List.Clear()
}

func (that *ItemList) Keys() (r []Item) {
	for _, key := range that.List.Keys() {
		k := key.(string)
		r = append(r, Item(k))
	}
	return
}

type Selector struct {
	Program  *tea.Program
	model    *SelectorModel
	itemList *ItemList
}

func NewSelector(itemList *ItemList, opts ...SOption) (sl *Selector) {
	sl = &Selector{itemList: itemList}
	items := itemList.Keys()
	sl.model = NewSelectorModel(items, opts...)
	sl.Program = tea.NewProgram(sl.model)
	return
}

func (that *Selector) Run() {
	if _, err := that.Program.Run(); err != nil {
		gprint.PrintError("%+v", err)
	}
}

func (that *Selector) Value() interface{} {
	choice := that.model.Choice()
	if that.itemList != nil {
		return that.itemList.Get(choice)
	}
	return nil
}
