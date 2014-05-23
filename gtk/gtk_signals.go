package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>
*/
import "C"
import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/gdk"
import "github.com/reusee/ggir/cairo"
import "unsafe"
import "reflect"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
}

var _ = gobject.UnusedFix_
var _ = gdk.UnusedFix_
var _ = cairo.UnusedFix_

func (self *TraitAboutDialog) OnActivateLink(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-link", func(_self reflect.Value, uri reflect.Value) {
	})
}

func (self *TraitAccelGroup) OnAccelActivate(f func(*gobject.Object, uint, C.GdkModifierType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "accel-activate", func(_self reflect.Value, acceleratable reflect.Value, keyval reflect.Value, modifier reflect.Value) {
	})
}

func (self *TraitAccelGroup) OnAccelChanged(f func(uint, C.GdkModifierType, *gobject.Closure)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "accel-changed", func(_self reflect.Value, keyval reflect.Value, modifier reflect.Value, accel_closure reflect.Value) {
	})
}

func (self *TraitAccelMap) OnChanged(f func(string, uint, C.GdkModifierType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value, accel_path reflect.Value, accel_key reflect.Value, accel_mods reflect.Value) {
	})
}

func (self *TraitAdjustment) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitAdjustment) OnValueChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "value-changed", func(_self reflect.Value) {
	})
}

func (self *TraitAppChooserButton) OnCustomItemActivated(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "custom-item-activated", func(_self reflect.Value, item_name reflect.Value) {
	})
}

func (self *TraitApplication) OnWindowAdded(f func(*Window)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "window-added", func(_self reflect.Value, window reflect.Value) {
	})
}

func (self *TraitApplication) OnWindowRemoved(f func(*Window)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "window-removed", func(_self reflect.Value, window reflect.Value) {
	})
}

func (self *TraitAssistant) OnApply(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "apply", func(_self reflect.Value) {
	})
}

func (self *TraitAssistant) OnCancel(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cancel", func(_self reflect.Value) {
	})
}

func (self *TraitAssistant) OnClose(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "close", func(_self reflect.Value) {
	})
}

func (self *TraitAssistant) OnPrepare(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "prepare", func(_self reflect.Value, page reflect.Value) {
	})
}

func (self *TraitButton) OnActivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value) {
	})
}

func (self *TraitButton) OnClicked(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "clicked", func(_self reflect.Value) {
	})
}

func (self *TraitCalendar) OnDaySelected(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "day-selected", func(_self reflect.Value) {
	})
}

func (self *TraitCalendar) OnDaySelectedDoubleClick(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "day-selected-double-click", func(_self reflect.Value) {
	})
}

func (self *TraitCalendar) OnMonthChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "month-changed", func(_self reflect.Value) {
	})
}

func (self *TraitCalendar) OnNextMonth(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "next-month", func(_self reflect.Value) {
	})
}

func (self *TraitCalendar) OnNextYear(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "next-year", func(_self reflect.Value) {
	})
}

func (self *TraitCalendar) OnPrevMonth(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "prev-month", func(_self reflect.Value) {
	})
}

func (self *TraitCalendar) OnPrevYear(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "prev-year", func(_self reflect.Value) {
	})
}

func (self *TraitCellArea) OnFocusChanged(f func(*CellRenderer, string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "focus-changed", func(_self reflect.Value, renderer reflect.Value, path reflect.Value) {
	})
}

func (self *TraitCellRenderer) OnEditingCanceled(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "editing-canceled", func(_self reflect.Value) {
	})
}

func (self *TraitCellRendererAccel) OnAccelCleared(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "accel-cleared", func(_self reflect.Value, path_string reflect.Value) {
	})
}

func (self *TraitCellRendererAccel) OnAccelEdited(f func(string, uint, C.GdkModifierType, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "accel-edited", func(_self reflect.Value, path_string reflect.Value, accel_key reflect.Value, accel_mods reflect.Value, hardware_keycode reflect.Value) {
	})
}

func (self *TraitCellRendererCombo) OnChanged(f func(string, *TreeIter)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value, path_string reflect.Value, new_iter reflect.Value) {
	})
}

func (self *TraitCellRendererText) OnEdited(f func(string, string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "edited", func(_self reflect.Value, path reflect.Value, new_text reflect.Value) {
	})
}

func (self *TraitCellRendererToggle) OnToggled(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggled", func(_self reflect.Value, path reflect.Value) {
	})
}

func (self *TraitCheckMenuItem) OnToggled(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggled", func(_self reflect.Value) {
	})
}

func (self *TraitClipboard) OnOwnerChange(f func(*gdk.EventOwnerChange)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "owner-change", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitColorButton) OnColorSet(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "color-set", func(_self reflect.Value) {
	})
}

func (self *TraitColorSelection) OnColorChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "color-changed", func(_self reflect.Value) {
	})
}

func (self *TraitComboBox) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitComboBox) OnFormatEntryText(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "format-entry-text", func(_self reflect.Value, path reflect.Value) {
	})
}

func (self *TraitComboBox) OnMoveActive(f func(C.GtkScrollType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-active", func(_self reflect.Value, scroll_type reflect.Value) {
	})
}

func (self *TraitComboBox) OnPopdown(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "popdown", func(_self reflect.Value) {
	})
}

func (self *TraitComboBox) OnPopup(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "popup", func(_self reflect.Value) {
	})
}

func (self *TraitContainer) OnAdd(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "add", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitContainer) OnCheckResize(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "check-resize", func(_self reflect.Value) {
	})
}

func (self *TraitContainer) OnRemove(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "remove", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitContainer) OnSetFocusChild(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "set-focus-child", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitDialog) OnClose(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "close", func(_self reflect.Value) {
	})
}

func (self *TraitDialog) OnResponse(f func(int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "response", func(_self reflect.Value, response_id reflect.Value) {
	})
}

func (self *TraitEntry) OnActivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value) {
	})
}

func (self *TraitEntry) OnBackspace(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "backspace", func(_self reflect.Value) {
	})
}

func (self *TraitEntry) OnCopyClipboard(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "copy-clipboard", func(_self reflect.Value) {
	})
}

func (self *TraitEntry) OnCutClipboard(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cut-clipboard", func(_self reflect.Value) {
	})
}

func (self *TraitEntry) OnDeleteFromCursor(f func(C.GtkDeleteType, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "delete-from-cursor", func(_self reflect.Value, type_ reflect.Value, count reflect.Value) {
	})
}

func (self *TraitEntry) OnIconPress(f func(C.GtkEntryIconPosition, *gdk.EventButton)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "icon-press", func(_self reflect.Value, icon_pos reflect.Value, event reflect.Value) {
	})
}

func (self *TraitEntry) OnIconRelease(f func(C.GtkEntryIconPosition, *gdk.EventButton)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "icon-release", func(_self reflect.Value, icon_pos reflect.Value, event reflect.Value) {
	})
}

func (self *TraitEntry) OnInsertAtCursor(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "insert-at-cursor", func(_self reflect.Value, string_ reflect.Value) {
	})
}

func (self *TraitEntry) OnMoveCursor(f func(C.GtkMovementStep, int, bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-cursor", func(_self reflect.Value, step reflect.Value, count reflect.Value, extend_selection reflect.Value) {
	})
}

func (self *TraitEntry) OnPasteClipboard(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "paste-clipboard", func(_self reflect.Value) {
	})
}

func (self *TraitEntry) OnPopulatePopup(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "populate-popup", func(_self reflect.Value, popup reflect.Value) {
	})
}

func (self *TraitEntry) OnPreeditChanged(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "preedit-changed", func(_self reflect.Value, preedit reflect.Value) {
	})
}

func (self *TraitEntry) OnToggleOverwrite(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-overwrite", func(_self reflect.Value) {
	})
}

func (self *TraitEntryBuffer) OnDeletedText(f func(uint, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "deleted-text", func(_self reflect.Value, position reflect.Value, n_chars reflect.Value) {
	})
}

func (self *TraitEntryBuffer) OnInsertedText(f func(uint, string, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "inserted-text", func(_self reflect.Value, position reflect.Value, chars reflect.Value, n_chars reflect.Value) {
	})
}

func (self *TraitEntryCompletion) OnActionActivated(f func(int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "action-activated", func(_self reflect.Value, index reflect.Value) {
	})
}

func (self *TraitEntryCompletion) OnInsertPrefix(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "insert-prefix", func(_self reflect.Value, prefix reflect.Value) {
	})
}

func (self *TraitExpander) OnActivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserButton) OnFileSet(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "file-set", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnDesktopFolder(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "desktop-folder", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnDownFolder(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "down-folder", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnHomeFolder(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "home-folder", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnLocationPopup(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "location-popup", func(_self reflect.Value, path reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnLocationPopupOnPaste(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "location-popup-on-paste", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnLocationTogglePopup(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "location-toggle-popup", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnQuickBookmark(f func(int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "quick-bookmark", func(_self reflect.Value, bookmark_index reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnRecentShortcut(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "recent-shortcut", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnSearchShortcut(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "search-shortcut", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnShowHidden(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "show-hidden", func(_self reflect.Value) {
	})
}

func (self *TraitFileChooserWidget) OnUpFolder(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "up-folder", func(_self reflect.Value) {
	})
}

func (self *TraitFlowBox) OnActivateCursorChild(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-cursor-child", func(_self reflect.Value) {
	})
}

func (self *TraitFlowBox) OnChildActivated(f func(*FlowBoxChild)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "child-activated", func(_self reflect.Value, child reflect.Value) {
	})
}

func (self *TraitFlowBox) OnMoveCursor(f func(C.GtkMovementStep, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-cursor", func(_self reflect.Value, step reflect.Value, count reflect.Value) {
	})
}

func (self *TraitFlowBox) OnSelectAll(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "select-all", func(_self reflect.Value) {
	})
}

func (self *TraitFlowBox) OnSelectedChildrenChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "selected-children-changed", func(_self reflect.Value) {
	})
}

func (self *TraitFlowBox) OnToggleCursorChild(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-cursor-child", func(_self reflect.Value) {
	})
}

func (self *TraitFlowBox) OnUnselectAll(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "unselect-all", func(_self reflect.Value) {
	})
}

func (self *TraitFlowBoxChild) OnActivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value) {
	})
}

func (self *TraitFontButton) OnFontSet(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "font-set", func(_self reflect.Value) {
	})
}

func (self *TraitHSV) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitHSV) OnMove(f func(C.GtkDirectionType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitIMContext) OnCommit(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "commit", func(_self reflect.Value, str reflect.Value) {
	})
}

func (self *TraitIMContext) OnDeleteSurrounding(f func(int, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "delete-surrounding", func(_self reflect.Value, offset reflect.Value, n_chars reflect.Value) {
	})
}

func (self *TraitIMContext) OnPreeditChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "preedit-changed", func(_self reflect.Value) {
	})
}

func (self *TraitIMContext) OnPreeditEnd(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "preedit-end", func(_self reflect.Value) {
	})
}

func (self *TraitIMContext) OnPreeditStart(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "preedit-start", func(_self reflect.Value) {
	})
}

func (self *TraitIMContext) OnRetrieveSurrounding(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "retrieve-surrounding", func(_self reflect.Value) {
	})
}

func (self *TraitIconTheme) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitIconView) OnActivateCursorItem(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-cursor-item", func(_self reflect.Value) {
	})
}

func (self *TraitIconView) OnItemActivated(f func(*TreePath)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "item-activated", func(_self reflect.Value, path reflect.Value) {
	})
}

func (self *TraitIconView) OnMoveCursor(f func(C.GtkMovementStep, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-cursor", func(_self reflect.Value, step reflect.Value, count reflect.Value) {
	})
}

func (self *TraitIconView) OnSelectAll(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "select-all", func(_self reflect.Value) {
	})
}

func (self *TraitIconView) OnSelectCursorItem(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "select-cursor-item", func(_self reflect.Value) {
	})
}

func (self *TraitIconView) OnSelectionChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "selection-changed", func(_self reflect.Value) {
	})
}

func (self *TraitIconView) OnToggleCursorItem(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-cursor-item", func(_self reflect.Value) {
	})
}

func (self *TraitIconView) OnUnselectAll(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "unselect-all", func(_self reflect.Value) {
	})
}

func (self *TraitInfoBar) OnClose(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "close", func(_self reflect.Value) {
	})
}

func (self *TraitInfoBar) OnResponse(f func(int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "response", func(_self reflect.Value, response_id reflect.Value) {
	})
}

func (self *TraitLabel) OnActivateCurrentLink(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-current-link", func(_self reflect.Value) {
	})
}

func (self *TraitLabel) OnActivateLink(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-link", func(_self reflect.Value, uri reflect.Value) {
	})
}

func (self *TraitLabel) OnCopyClipboard(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "copy-clipboard", func(_self reflect.Value) {
	})
}

func (self *TraitLabel) OnMoveCursor(f func(C.GtkMovementStep, int, bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-cursor", func(_self reflect.Value, step reflect.Value, count reflect.Value, extend_selection reflect.Value) {
	})
}

func (self *TraitLabel) OnPopulatePopup(f func(*Menu)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "populate-popup", func(_self reflect.Value, menu reflect.Value) {
	})
}

func (self *TraitLevelBar) OnOffsetChanged(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "offset-changed", func(_self reflect.Value, name reflect.Value) {
	})
}

func (self *TraitLinkButton) OnActivateLink(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-link", func(_self reflect.Value) {
	})
}

func (self *TraitListBox) OnActivateCursorRow(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-cursor-row", func(_self reflect.Value) {
	})
}

func (self *TraitListBox) OnMoveCursor(f func(C.GtkMovementStep, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-cursor", func(_self reflect.Value, object reflect.Value, p0 reflect.Value) {
	})
}

func (self *TraitListBox) OnRowActivated(f func(*ListBoxRow)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "row-activated", func(_self reflect.Value, row reflect.Value) {
	})
}

func (self *TraitListBox) OnRowSelected(f func(*ListBoxRow)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "row-selected", func(_self reflect.Value, row reflect.Value) {
	})
}

func (self *TraitListBox) OnToggleCursorRow(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-cursor-row", func(_self reflect.Value) {
	})
}

func (self *TraitListBoxRow) OnActivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value) {
	})
}

func (self *TraitMenu) OnMoveScroll(f func(C.GtkScrollType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-scroll", func(_self reflect.Value, scroll_type reflect.Value) {
	})
}

func (self *TraitMenuItem) OnActivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value) {
	})
}

func (self *TraitMenuItem) OnActivateItem(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-item", func(_self reflect.Value) {
	})
}

func (self *TraitMenuItem) OnDeselect(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "deselect", func(_self reflect.Value) {
	})
}

func (self *TraitMenuItem) OnSelect(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "select", func(_self reflect.Value) {
	})
}

func (self *TraitMenuItem) OnToggleSizeAllocate(f func(int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-size-allocate", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitMenuItem) OnToggleSizeRequest(f func(unsafe.Pointer)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-size-request", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitMenuShell) OnActivateCurrent(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-current", func(_self reflect.Value, force_hide reflect.Value) {
	})
}

func (self *TraitMenuShell) OnCancel(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cancel", func(_self reflect.Value) {
	})
}

func (self *TraitMenuShell) OnCycleFocus(f func(C.GtkDirectionType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cycle-focus", func(_self reflect.Value, direction reflect.Value) {
	})
}

func (self *TraitMenuShell) OnDeactivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "deactivate", func(_self reflect.Value) {
	})
}

func (self *TraitMenuShell) OnInsert(f func(*Widget, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "insert", func(_self reflect.Value, child reflect.Value, position reflect.Value) {
	})
}

func (self *TraitMenuShell) OnMoveCurrent(f func(C.GtkMenuDirectionType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-current", func(_self reflect.Value, direction reflect.Value) {
	})
}

func (self *TraitMenuShell) OnMoveSelected(f func(int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-selected", func(_self reflect.Value, distance reflect.Value) {
	})
}

func (self *TraitMenuShell) OnSelectionDone(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "selection-done", func(_self reflect.Value) {
	})
}

func (self *TraitMenuToolButton) OnShowMenu(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "show-menu", func(_self reflect.Value) {
	})
}

func (self *TraitNotebook) OnChangeCurrentPage(f func(int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "change-current-page", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitNotebook) OnCreateWindow(f func(*Widget, int, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "create-window", func(_self reflect.Value, page reflect.Value, x reflect.Value, y reflect.Value) {
	})
}

func (self *TraitNotebook) OnFocusTab(f func(C.GtkNotebookTab)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "focus-tab", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitNotebook) OnMoveFocusOut(f func(C.GtkDirectionType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-focus-out", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitNotebook) OnPageAdded(f func(*Widget, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "page-added", func(_self reflect.Value, child reflect.Value, page_num reflect.Value) {
	})
}

func (self *TraitNotebook) OnPageRemoved(f func(*Widget, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "page-removed", func(_self reflect.Value, child reflect.Value, page_num reflect.Value) {
	})
}

func (self *TraitNotebook) OnPageReordered(f func(*Widget, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "page-reordered", func(_self reflect.Value, child reflect.Value, page_num reflect.Value) {
	})
}

func (self *TraitNotebook) OnReorderTab(f func(C.GtkDirectionType, bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "reorder-tab", func(_self reflect.Value, object reflect.Value, p0 reflect.Value) {
	})
}

func (self *TraitNotebook) OnSelectPage(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "select-page", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitNotebook) OnSwitchPage(f func(*Widget, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "switch-page", func(_self reflect.Value, page reflect.Value, page_num reflect.Value) {
	})
}

func (self *TraitOverlay) OnGetChildPosition(f func(*Widget, *cairo.RectangleInt)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "get-child-position", func(_self reflect.Value, widget reflect.Value, allocation reflect.Value) {
	})
}

func (self *TraitPaned) OnAcceptPosition(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "accept-position", func(_self reflect.Value) {
	})
}

func (self *TraitPaned) OnCancelPosition(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cancel-position", func(_self reflect.Value) {
	})
}

func (self *TraitPaned) OnCycleChildFocus(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cycle-child-focus", func(_self reflect.Value, reversed reflect.Value) {
	})
}

func (self *TraitPaned) OnCycleHandleFocus(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cycle-handle-focus", func(_self reflect.Value, reversed reflect.Value) {
	})
}

func (self *TraitPaned) OnMoveHandle(f func(C.GtkScrollType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-handle", func(_self reflect.Value, scroll_type reflect.Value) {
	})
}

func (self *TraitPaned) OnToggleHandleFocus(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-handle-focus", func(_self reflect.Value) {
	})
}

func (self *TraitPlacesSidebar) OnDragActionAsk(f func(int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-action-ask", func(_self reflect.Value, actions reflect.Value) {
	})
}

func (self *TraitPlacesSidebar) OnDragActionRequested(f func(*gdk.DragContext, *gobject.Object, unsafe.Pointer)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-action-requested", func(_self reflect.Value, context reflect.Value, dest_file reflect.Value, source_file_list reflect.Value) {
	})
}

func (self *TraitPlacesSidebar) OnDragPerformDrop(f func(*gobject.Object, unsafe.Pointer, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-perform-drop", func(_self reflect.Value, dest_file reflect.Value, source_file_list reflect.Value, action reflect.Value) {
	})
}

func (self *TraitPlacesSidebar) OnOpenLocation(f func(*gobject.Object, C.GtkPlacesOpenFlags)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "open-location", func(_self reflect.Value, location reflect.Value, open_flags reflect.Value) {
	})
}

func (self *TraitPlacesSidebar) OnPopulatePopup(f func(*gobject.Object, *gobject.Object, *gobject.Object)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "populate-popup", func(_self reflect.Value, menu reflect.Value, selected_item reflect.Value, selected_volume reflect.Value) {
	})
}

func (self *TraitPlacesSidebar) OnShowConnectToServer(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "show-connect-to-server", func(_self reflect.Value) {
	})
}

func (self *TraitPlacesSidebar) OnShowErrorMessage(f func(string, string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "show-error-message", func(_self reflect.Value, primary reflect.Value, secondary reflect.Value) {
	})
}

func (self *TraitPlug) OnEmbedded(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "embedded", func(_self reflect.Value) {
	})
}

func (self *TraitPopover) OnClosed(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "closed", func(_self reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnBeginPrint(f func(*PrintContext)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "begin-print", func(_self reflect.Value, context reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnCreateCustomWidget(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "create-custom-widget", func(_self reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnCustomWidgetApply(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "custom-widget-apply", func(_self reflect.Value, widget reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnDone(f func(C.GtkPrintOperationResult)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "done", func(_self reflect.Value, result reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnDrawPage(f func(*PrintContext, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "draw-page", func(_self reflect.Value, context reflect.Value, page_nr reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnEndPrint(f func(*PrintContext)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "end-print", func(_self reflect.Value, context reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnPaginate(f func(*PrintContext)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "paginate", func(_self reflect.Value, context reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnRequestPageSetup(f func(*PrintContext, int, *PageSetup)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "request-page-setup", func(_self reflect.Value, context reflect.Value, page_nr reflect.Value, setup reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnStatusChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "status-changed", func(_self reflect.Value) {
	})
}

func (self *TraitPrintOperation) OnUpdateCustomWidget(f func(*Widget, *PageSetup, *PrintSettings)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "update-custom-widget", func(_self reflect.Value, widget reflect.Value, setup reflect.Value, settings reflect.Value) {
	})
}

func (self *TraitRadioButton) OnGroupChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "group-changed", func(_self reflect.Value) {
	})
}

func (self *TraitRadioMenuItem) OnGroupChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "group-changed", func(_self reflect.Value) {
	})
}

func (self *TraitRange) OnAdjustBounds(f func(float64)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "adjust-bounds", func(_self reflect.Value, value reflect.Value) {
	})
}

func (self *TraitRange) OnChangeValue(f func(C.GtkScrollType, float64)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "change-value", func(_self reflect.Value, scroll reflect.Value, value reflect.Value) {
	})
}

func (self *TraitRange) OnMoveSlider(f func(C.GtkScrollType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-slider", func(_self reflect.Value, step reflect.Value) {
	})
}

func (self *TraitRange) OnValueChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "value-changed", func(_self reflect.Value) {
	})
}

func (self *TraitRecentManager) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitScale) OnFormatValue(f func(float64)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "format-value", func(_self reflect.Value, value reflect.Value) {
	})
}

func (self *TraitScaleButton) OnPopdown(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "popdown", func(_self reflect.Value) {
	})
}

func (self *TraitScaleButton) OnPopup(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "popup", func(_self reflect.Value) {
	})
}

func (self *TraitScaleButton) OnValueChanged(f func(float64)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "value-changed", func(_self reflect.Value, value reflect.Value) {
	})
}

func (self *TraitScrolledWindow) OnMoveFocusOut(f func(C.GtkDirectionType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-focus-out", func(_self reflect.Value, direction_type reflect.Value) {
	})
}

func (self *TraitScrolledWindow) OnScrollChild(f func(C.GtkScrollType, bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "scroll-child", func(_self reflect.Value, scroll reflect.Value, horizontal reflect.Value) {
	})
}

func (self *TraitSearchEntry) OnSearchChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "search-changed", func(_self reflect.Value) {
	})
}

func (self *TraitSocket) OnPlugAdded(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "plug-added", func(_self reflect.Value) {
	})
}

func (self *TraitSocket) OnPlugRemoved(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "plug-removed", func(_self reflect.Value) {
	})
}

func (self *TraitSpinButton) OnChangeValue(f func(C.GtkScrollType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "change-value", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitSpinButton) OnInput(f func(unsafe.Pointer)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "input", func(_self reflect.Value, new_value reflect.Value) {
	})
}

func (self *TraitSpinButton) OnOutput(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "output", func(_self reflect.Value) {
	})
}

func (self *TraitSpinButton) OnValueChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "value-changed", func(_self reflect.Value) {
	})
}

func (self *TraitSpinButton) OnWrapped(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "wrapped", func(_self reflect.Value) {
	})
}

func (self *TraitStatusIcon) OnActivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value) {
	})
}

func (self *TraitStatusIcon) OnButtonPressEvent(f func(*gdk.EventButton)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "button-press-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitStatusIcon) OnButtonReleaseEvent(f func(*gdk.EventButton)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "button-release-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitStatusIcon) OnPopupMenu(f func(uint, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "popup-menu", func(_self reflect.Value, button reflect.Value, activate_time reflect.Value) {
	})
}

func (self *TraitStatusIcon) OnQueryTooltip(f func(int, int, bool, *Tooltip)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "query-tooltip", func(_self reflect.Value, x reflect.Value, y reflect.Value, keyboard_mode reflect.Value, tooltip reflect.Value) {
	})
}

func (self *TraitStatusIcon) OnScrollEvent(f func(*gdk.EventScroll)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "scroll-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitStatusIcon) OnSizeChanged(f func(int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "size-changed", func(_self reflect.Value, size reflect.Value) {
	})
}

func (self *TraitStatusbar) OnTextPopped(f func(uint, string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "text-popped", func(_self reflect.Value, context_id reflect.Value, text reflect.Value) {
	})
}

func (self *TraitStatusbar) OnTextPushed(f func(uint, string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "text-pushed", func(_self reflect.Value, context_id reflect.Value, text reflect.Value) {
	})
}

func (self *TraitStyle) OnRealize(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "realize", func(_self reflect.Value) {
	})
}

func (self *TraitStyle) OnUnrealize(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "unrealize", func(_self reflect.Value) {
	})
}

func (self *TraitStyleContext) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitSwitch) OnActivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnApplyTag(f func(*TextTag, *TextIter, *TextIter)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "apply-tag", func(_self reflect.Value, tag reflect.Value, start reflect.Value, end reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnBeginUserAction(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "begin-user-action", func(_self reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnDeleteRange(f func(*TextIter, *TextIter)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "delete-range", func(_self reflect.Value, start reflect.Value, end reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnEndUserAction(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "end-user-action", func(_self reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnInsertChildAnchor(f func(*TextIter, *TextChildAnchor)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "insert-child-anchor", func(_self reflect.Value, location reflect.Value, anchor reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnInsertText(f func(*TextIter, string, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "insert-text", func(_self reflect.Value, location reflect.Value, text reflect.Value, len_ reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnMarkDeleted(f func(*TextMark)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "mark-deleted", func(_self reflect.Value, mark reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnMarkSet(f func(*TextIter, *TextMark)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "mark-set", func(_self reflect.Value, location reflect.Value, mark reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnModifiedChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "modified-changed", func(_self reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnPasteDone(f func(*Clipboard)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "paste-done", func(_self reflect.Value, clipboard reflect.Value) {
	})
}

func (self *TraitTextBuffer) OnRemoveTag(f func(*TextTag, *TextIter, *TextIter)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "remove-tag", func(_self reflect.Value, tag reflect.Value, start reflect.Value, end reflect.Value) {
	})
}

func (self *TraitTextTag) OnEvent(f func(*gobject.Object, *C.GdkEvent, *TextIter)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "event", func(_self reflect.Value, object reflect.Value, event reflect.Value, iter reflect.Value) {
	})
}

func (self *TraitTextTagTable) OnTagAdded(f func(*TextTag)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "tag-added", func(_self reflect.Value, tag reflect.Value) {
	})
}

func (self *TraitTextTagTable) OnTagChanged(f func(*TextTag, bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "tag-changed", func(_self reflect.Value, tag reflect.Value, size_changed reflect.Value) {
	})
}

func (self *TraitTextTagTable) OnTagRemoved(f func(*TextTag)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "tag-removed", func(_self reflect.Value, tag reflect.Value) {
	})
}

func (self *TraitTextView) OnBackspace(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "backspace", func(_self reflect.Value) {
	})
}

func (self *TraitTextView) OnCopyClipboard(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "copy-clipboard", func(_self reflect.Value) {
	})
}

func (self *TraitTextView) OnCutClipboard(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cut-clipboard", func(_self reflect.Value) {
	})
}

func (self *TraitTextView) OnDeleteFromCursor(f func(C.GtkDeleteType, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "delete-from-cursor", func(_self reflect.Value, type_ reflect.Value, count reflect.Value) {
	})
}

func (self *TraitTextView) OnInsertAtCursor(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "insert-at-cursor", func(_self reflect.Value, string_ reflect.Value) {
	})
}

func (self *TraitTextView) OnMoveCursor(f func(C.GtkMovementStep, int, bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-cursor", func(_self reflect.Value, step reflect.Value, count reflect.Value, extend_selection reflect.Value) {
	})
}

func (self *TraitTextView) OnMoveViewport(f func(C.GtkScrollStep, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-viewport", func(_self reflect.Value, step reflect.Value, count reflect.Value) {
	})
}

func (self *TraitTextView) OnPasteClipboard(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "paste-clipboard", func(_self reflect.Value) {
	})
}

func (self *TraitTextView) OnPopulatePopup(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "populate-popup", func(_self reflect.Value, popup reflect.Value) {
	})
}

func (self *TraitTextView) OnPreeditChanged(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "preedit-changed", func(_self reflect.Value, preedit reflect.Value) {
	})
}

func (self *TraitTextView) OnSelectAll(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "select-all", func(_self reflect.Value, select_ reflect.Value) {
	})
}

func (self *TraitTextView) OnSetAnchor(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "set-anchor", func(_self reflect.Value) {
	})
}

func (self *TraitTextView) OnToggleCursorVisible(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-cursor-visible", func(_self reflect.Value) {
	})
}

func (self *TraitTextView) OnToggleOverwrite(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-overwrite", func(_self reflect.Value) {
	})
}

func (self *TraitToggleButton) OnToggled(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggled", func(_self reflect.Value) {
	})
}

func (self *TraitToggleToolButton) OnToggled(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggled", func(_self reflect.Value) {
	})
}

func (self *TraitToolButton) OnClicked(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "clicked", func(_self reflect.Value) {
	})
}

func (self *TraitToolItem) OnCreateMenuProxy(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "create-menu-proxy", func(_self reflect.Value) {
	})
}

func (self *TraitToolItem) OnToolbarReconfigured(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toolbar-reconfigured", func(_self reflect.Value) {
	})
}

func (self *TraitToolbar) OnFocusHomeOrEnd(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "focus-home-or-end", func(_self reflect.Value, focus_home reflect.Value) {
	})
}

func (self *TraitToolbar) OnOrientationChanged(f func(C.GtkOrientation)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "orientation-changed", func(_self reflect.Value, orientation reflect.Value) {
	})
}

func (self *TraitToolbar) OnPopupContextMenu(f func(int, int, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "popup-context-menu", func(_self reflect.Value, x reflect.Value, y reflect.Value, button reflect.Value) {
	})
}

func (self *TraitToolbar) OnStyleChanged(f func(C.GtkToolbarStyle)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "style-changed", func(_self reflect.Value, style reflect.Value) {
	})
}

func (self *TraitTreeSelection) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitTreeView) OnColumnsChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "columns-changed", func(_self reflect.Value) {
	})
}

func (self *TraitTreeView) OnCursorChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cursor-changed", func(_self reflect.Value) {
	})
}

func (self *TraitTreeView) OnExpandCollapseCursorRow(f func(bool, bool, bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "expand-collapse-cursor-row", func(_self reflect.Value, object reflect.Value, p0 reflect.Value, p1 reflect.Value) {
	})
}

func (self *TraitTreeView) OnMoveCursor(f func(C.GtkMovementStep, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-cursor", func(_self reflect.Value, step reflect.Value, direction reflect.Value) {
	})
}

func (self *TraitTreeView) OnRowActivated(f func(*TreePath, *TreeViewColumn)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "row-activated", func(_self reflect.Value, path reflect.Value, column reflect.Value) {
	})
}

func (self *TraitTreeView) OnRowCollapsed(f func(*TreeIter, *TreePath)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "row-collapsed", func(_self reflect.Value, iter reflect.Value, path reflect.Value) {
	})
}

func (self *TraitTreeView) OnRowExpanded(f func(*TreeIter, *TreePath)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "row-expanded", func(_self reflect.Value, iter reflect.Value, path reflect.Value) {
	})
}

func (self *TraitTreeView) OnSelectAll(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "select-all", func(_self reflect.Value) {
	})
}

func (self *TraitTreeView) OnSelectCursorParent(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "select-cursor-parent", func(_self reflect.Value) {
	})
}

func (self *TraitTreeView) OnSelectCursorRow(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "select-cursor-row", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitTreeView) OnStartInteractiveSearch(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "start-interactive-search", func(_self reflect.Value) {
	})
}

func (self *TraitTreeView) OnTestCollapseRow(f func(*TreeIter, *TreePath)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "test-collapse-row", func(_self reflect.Value, iter reflect.Value, path reflect.Value) {
	})
}

func (self *TraitTreeView) OnTestExpandRow(f func(*TreeIter, *TreePath)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "test-expand-row", func(_self reflect.Value, iter reflect.Value, path reflect.Value) {
	})
}

func (self *TraitTreeView) OnToggleCursorRow(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "toggle-cursor-row", func(_self reflect.Value) {
	})
}

func (self *TraitTreeView) OnUnselectAll(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "unselect-all", func(_self reflect.Value) {
	})
}

func (self *TraitTreeViewColumn) OnClicked(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "clicked", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnAccelClosuresChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "accel-closures-changed", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnButtonPressEvent(f func(*gdk.EventButton)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "button-press-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnButtonReleaseEvent(f func(*gdk.EventButton)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "button-release-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnCanActivateAccel(f func(uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "can-activate-accel", func(_self reflect.Value, signal_id reflect.Value) {
	})
}

func (self *TraitWidget) OnChildNotify(f func(*gobject.ParamSpec)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "child-notify", func(_self reflect.Value, child_property reflect.Value) {
	})
}

func (self *TraitWidget) OnCompositedChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "composited-changed", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnConfigureEvent(f func(*gdk.EventConfigure)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "configure-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnDamageEvent(f func(*gdk.EventExpose)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "damage-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnDeleteEvent(f func(*C.GdkEvent)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "delete-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnDestroy(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "destroy", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnDestroyEvent(f func(*C.GdkEvent)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "destroy-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnDirectionChanged(f func(C.GtkTextDirection)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "direction-changed", func(_self reflect.Value, previous_direction reflect.Value) {
	})
}

func (self *TraitWidget) OnDragBegin(f func(*gdk.DragContext)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-begin", func(_self reflect.Value, context reflect.Value) {
	})
}

func (self *TraitWidget) OnDragDataDelete(f func(*gdk.DragContext)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-data-delete", func(_self reflect.Value, context reflect.Value) {
	})
}

func (self *TraitWidget) OnDragDataGet(f func(*gdk.DragContext, *SelectionData, uint, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-data-get", func(_self reflect.Value, context reflect.Value, data reflect.Value, info reflect.Value, time reflect.Value) {
	})
}

func (self *TraitWidget) OnDragDataReceived(f func(*gdk.DragContext, int, int, *SelectionData, uint, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-data-received", func(_self reflect.Value, context reflect.Value, x reflect.Value, y reflect.Value, data reflect.Value, info reflect.Value, time reflect.Value) {
	})
}

func (self *TraitWidget) OnDragDrop(f func(*gdk.DragContext, int, int, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-drop", func(_self reflect.Value, context reflect.Value, x reflect.Value, y reflect.Value, time reflect.Value) {
	})
}

func (self *TraitWidget) OnDragEnd(f func(*gdk.DragContext)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-end", func(_self reflect.Value, context reflect.Value) {
	})
}

func (self *TraitWidget) OnDragFailed(f func(*gdk.DragContext, C.GtkDragResult)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-failed", func(_self reflect.Value, context reflect.Value, result reflect.Value) {
	})
}

func (self *TraitWidget) OnDragLeave(f func(*gdk.DragContext, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-leave", func(_self reflect.Value, context reflect.Value, time reflect.Value) {
	})
}

func (self *TraitWidget) OnDragMotion(f func(*gdk.DragContext, int, int, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "drag-motion", func(_self reflect.Value, context reflect.Value, x reflect.Value, y reflect.Value, time reflect.Value) {
	})
}

func (self *TraitWidget) OnDraw(f func(*cairo.Context)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "draw", func(_self reflect.Value, cr reflect.Value) {
	})
}

func (self *TraitWidget) OnEnterNotifyEvent(f func(*gdk.EventCrossing)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "enter-notify-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnEvent(f func(*C.GdkEvent)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnEventAfter(f func(*C.GdkEvent)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "event-after", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnFocus(f func(C.GtkDirectionType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "focus", func(_self reflect.Value, direction reflect.Value) {
	})
}

func (self *TraitWidget) OnFocusInEvent(f func(*gdk.EventFocus)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "focus-in-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnFocusOutEvent(f func(*gdk.EventFocus)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "focus-out-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnGrabBrokenEvent(f func(*gdk.EventGrabBroken)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "grab-broken-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnGrabFocus(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "grab-focus", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnGrabNotify(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "grab-notify", func(_self reflect.Value, was_grabbed reflect.Value) {
	})
}

func (self *TraitWidget) OnHide(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "hide", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnHierarchyChanged(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "hierarchy-changed", func(_self reflect.Value, previous_toplevel reflect.Value) {
	})
}

func (self *TraitWidget) OnKeyPressEvent(f func(*gdk.EventKey)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "key-press-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnKeyReleaseEvent(f func(*gdk.EventKey)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "key-release-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnKeynavFailed(f func(C.GtkDirectionType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "keynav-failed", func(_self reflect.Value, direction reflect.Value) {
	})
}

func (self *TraitWidget) OnLeaveNotifyEvent(f func(*gdk.EventCrossing)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "leave-notify-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnMap(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "map", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnMapEvent(f func(*gdk.EventAny)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "map-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnMnemonicActivate(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "mnemonic-activate", func(_self reflect.Value, arg1 reflect.Value) {
	})
}

func (self *TraitWidget) OnMotionNotifyEvent(f func(*gdk.EventMotion)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "motion-notify-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnMoveFocus(f func(C.GtkDirectionType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "move-focus", func(_self reflect.Value, direction reflect.Value) {
	})
}

func (self *TraitWidget) OnParentSet(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "parent-set", func(_self reflect.Value, old_parent reflect.Value) {
	})
}

func (self *TraitWidget) OnPopupMenu(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "popup-menu", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnPropertyNotifyEvent(f func(*gdk.EventProperty)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "property-notify-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnProximityInEvent(f func(*gdk.EventProximity)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "proximity-in-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnProximityOutEvent(f func(*gdk.EventProximity)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "proximity-out-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnQueryTooltip(f func(int, int, bool, *Tooltip)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "query-tooltip", func(_self reflect.Value, x reflect.Value, y reflect.Value, keyboard_mode reflect.Value, tooltip reflect.Value) {
	})
}

func (self *TraitWidget) OnRealize(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "realize", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnScreenChanged(f func(*gdk.Screen)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "screen-changed", func(_self reflect.Value, previous_screen reflect.Value) {
	})
}

func (self *TraitWidget) OnScrollEvent(f func(*gdk.EventScroll)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "scroll-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnSelectionClearEvent(f func(*gdk.EventSelection)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "selection-clear-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnSelectionGet(f func(*SelectionData, uint, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "selection-get", func(_self reflect.Value, data reflect.Value, info reflect.Value, time reflect.Value) {
	})
}

func (self *TraitWidget) OnSelectionNotifyEvent(f func(*gdk.EventSelection)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "selection-notify-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnSelectionReceived(f func(*SelectionData, uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "selection-received", func(_self reflect.Value, data reflect.Value, time reflect.Value) {
	})
}

func (self *TraitWidget) OnSelectionRequestEvent(f func(*gdk.EventSelection)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "selection-request-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnShow(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "show", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnShowHelp(f func(C.GtkWidgetHelpType)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "show-help", func(_self reflect.Value, help_type reflect.Value) {
	})
}

func (self *TraitWidget) OnSizeAllocate(f func(*cairo.RectangleInt)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "size-allocate", func(_self reflect.Value, allocation reflect.Value) {
	})
}

func (self *TraitWidget) OnStateFlagsChanged(f func(C.GtkStateFlags)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "state-flags-changed", func(_self reflect.Value, flags reflect.Value) {
	})
}

func (self *TraitWidget) OnStyleUpdated(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "style-updated", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnTouchEvent(f func(*C.GdkEvent)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "touch-event", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitWidget) OnUnmap(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "unmap", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnUnmapEvent(f func(*gdk.EventAny)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "unmap-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWidget) OnUnrealize(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "unrealize", func(_self reflect.Value) {
	})
}

func (self *TraitWidget) OnWindowStateEvent(f func(*gdk.EventWindowState)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "window-state-event", func(_self reflect.Value, event reflect.Value) {
	})
}

func (self *TraitWindow) OnActivateDefault(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-default", func(_self reflect.Value) {
	})
}

func (self *TraitWindow) OnActivateFocus(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate-focus", func(_self reflect.Value) {
	})
}

func (self *TraitWindow) OnKeysChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "keys-changed", func(_self reflect.Value) {
	})
}

func (self *TraitWindow) OnSetFocus(f func(*Widget)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "set-focus", func(_self reflect.Value, object reflect.Value) {
	})
}
