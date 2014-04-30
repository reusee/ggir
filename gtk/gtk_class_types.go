package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type AboutDialog struct {
	*_TraitAboutDialog
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type AccelGroup struct {
	*_TraitAccelGroup
	CPointer unsafe.Pointer
}

type AccelLabel struct {
	*_TraitAccelLabel
	*_TraitLabel
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

type AccelMap struct {
	*_TraitAccelMap
	CPointer unsafe.Pointer
}

type Accessible struct {
	*_TraitAccessible
	CPointer unsafe.Pointer
}

type Action struct {
	*_TraitAction
	CPointer unsafe.Pointer
}

type ActionBar struct {
	*_TraitActionBar
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ActionGroup struct {
	*_TraitActionGroup
	CPointer unsafe.Pointer
}

type Adjustment struct {
	*_TraitAdjustment
	CPointer unsafe.Pointer
}

type Alignment struct {
	*_TraitAlignment
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type AppChooserButton struct {
	*_TraitAppChooserButton
	*_TraitComboBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type AppChooserDialog struct {
	*_TraitAppChooserDialog
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type AppChooserWidget struct {
	*_TraitAppChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Application struct {
	*_TraitApplication
	CPointer unsafe.Pointer
}

type ApplicationWindow struct {
	*_TraitApplicationWindow
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Arrow struct {
	*_TraitArrow
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

type AspectFrame struct {
	*_TraitAspectFrame
	*_TraitFrame
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Assistant struct {
	*_TraitAssistant
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Bin struct {
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Box struct {
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Builder struct {
	*_TraitBuilder
	CPointer unsafe.Pointer
}

type Button struct {
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ButtonBox struct {
	*_TraitButtonBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Calendar struct {
	*_TraitCalendar
	*_TraitWidget
	CPointer unsafe.Pointer
}

type CellArea struct {
	*_TraitCellArea
	CPointer unsafe.Pointer
}

type CellAreaBox struct {
	*_TraitCellAreaBox
	*_TraitCellArea
	CPointer unsafe.Pointer
}

type CellAreaContext struct {
	*_TraitCellAreaContext
	CPointer unsafe.Pointer
}

type CellRenderer struct {
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

type CellRendererAccel struct {
	*_TraitCellRendererAccel
	*_TraitCellRendererText
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

type CellRendererCombo struct {
	*_TraitCellRendererCombo
	*_TraitCellRendererText
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

type CellRendererPixbuf struct {
	*_TraitCellRendererPixbuf
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

type CellRendererProgress struct {
	*_TraitCellRendererProgress
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

type CellRendererSpin struct {
	*_TraitCellRendererSpin
	*_TraitCellRendererText
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

type CellRendererSpinner struct {
	*_TraitCellRendererSpinner
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

type CellRendererText struct {
	*_TraitCellRendererText
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

type CellRendererToggle struct {
	*_TraitCellRendererToggle
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

type CellView struct {
	*_TraitCellView
	*_TraitWidget
	CPointer unsafe.Pointer
}

type CheckButton struct {
	*_TraitCheckButton
	*_TraitToggleButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type CheckMenuItem struct {
	*_TraitCheckMenuItem
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Clipboard struct {
	*_TraitClipboard
	CPointer unsafe.Pointer
}

type ColorButton struct {
	*_TraitColorButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ColorChooserDialog struct {
	*_TraitColorChooserDialog
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ColorChooserWidget struct {
	*_TraitColorChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ColorSelection struct {
	*_TraitColorSelection
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ColorSelectionDialog struct {
	*_TraitColorSelectionDialog
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ComboBox struct {
	*_TraitComboBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ComboBoxText struct {
	*_TraitComboBoxText
	*_TraitComboBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Container struct {
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type CssProvider struct {
	*_TraitCssProvider
	CPointer unsafe.Pointer
}

type Dialog struct {
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type DrawingArea struct {
	*_TraitDrawingArea
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Entry struct {
	*_TraitEntry
	*_TraitWidget
	CPointer unsafe.Pointer
}

type EntryBuffer struct {
	*_TraitEntryBuffer
	CPointer unsafe.Pointer
}

type EntryCompletion struct {
	*_TraitEntryCompletion
	CPointer unsafe.Pointer
}

type EventBox struct {
	*_TraitEventBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Expander struct {
	*_TraitExpander
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FileChooserButton struct {
	*_TraitFileChooserButton
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FileChooserDialog struct {
	*_TraitFileChooserDialog
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FileChooserWidget struct {
	*_TraitFileChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FileFilter struct {
	*_TraitFileFilter
	CPointer unsafe.Pointer
}

type Fixed struct {
	*_TraitFixed
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FlowBox struct {
	*_TraitFlowBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FlowBoxChild struct {
	*_TraitFlowBoxChild
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FontButton struct {
	*_TraitFontButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FontChooserDialog struct {
	*_TraitFontChooserDialog
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FontChooserWidget struct {
	*_TraitFontChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FontSelection struct {
	*_TraitFontSelection
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type FontSelectionDialog struct {
	*_TraitFontSelectionDialog
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Frame struct {
	*_TraitFrame
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Grid struct {
	*_TraitGrid
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type HBox struct {
	*_TraitHBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type HButtonBox struct {
	*_TraitHButtonBox
	*_TraitButtonBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type HPaned struct {
	*_TraitHPaned
	*_TraitPaned
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type HSV struct {
	*_TraitHSV
	*_TraitWidget
	CPointer unsafe.Pointer
}

type HScale struct {
	*_TraitHScale
	*_TraitScale
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

type HScrollbar struct {
	*_TraitHScrollbar
	*_TraitScrollbar
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

type HSeparator struct {
	*_TraitHSeparator
	*_TraitSeparator
	*_TraitWidget
	CPointer unsafe.Pointer
}

type HandleBox struct {
	*_TraitHandleBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type HeaderBar struct {
	*_TraitHeaderBar
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type IMContext struct {
	*_TraitIMContext
	CPointer unsafe.Pointer
}

type IMContextSimple struct {
	*_TraitIMContextSimple
	*_TraitIMContext
	CPointer unsafe.Pointer
}

type IMMulticontext struct {
	*_TraitIMMulticontext
	*_TraitIMContext
	CPointer unsafe.Pointer
}

type IconFactory struct {
	*_TraitIconFactory
	CPointer unsafe.Pointer
}

type IconInfo struct {
	*_TraitIconInfo
	CPointer unsafe.Pointer
}

type IconTheme struct {
	*_TraitIconTheme
	CPointer unsafe.Pointer
}

type IconView struct {
	*_TraitIconView
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Image struct {
	*_TraitImage
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ImageMenuItem struct {
	*_TraitImageMenuItem
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type InfoBar struct {
	*_TraitInfoBar
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Invisible struct {
	*_TraitInvisible
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Label struct {
	*_TraitLabel
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Layout struct {
	*_TraitLayout
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type LevelBar struct {
	*_TraitLevelBar
	*_TraitWidget
	CPointer unsafe.Pointer
}

type LinkButton struct {
	*_TraitLinkButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ListBox struct {
	*_TraitListBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ListBoxRow struct {
	*_TraitListBoxRow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ListStore struct {
	*_TraitListStore
	CPointer unsafe.Pointer
}

type LockButton struct {
	*_TraitLockButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Menu struct {
	*_TraitMenu
	*_TraitMenuShell
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type MenuBar struct {
	*_TraitMenuBar
	*_TraitMenuShell
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type MenuButton struct {
	*_TraitMenuButton
	*_TraitToggleButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type MenuItem struct {
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type MenuShell struct {
	*_TraitMenuShell
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type MenuToolButton struct {
	*_TraitMenuToolButton
	*_TraitToolButton
	*_TraitToolItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type MessageDialog struct {
	*_TraitMessageDialog
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Misc struct {
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

type MountOperation struct {
	*_TraitMountOperation
	CPointer unsafe.Pointer
}

type Notebook struct {
	*_TraitNotebook
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type NumerableIcon struct {
	*_TraitNumerableIcon
	CPointer unsafe.Pointer
}

type OffscreenWindow struct {
	*_TraitOffscreenWindow
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Overlay struct {
	*_TraitOverlay
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type PageSetup struct {
	*_TraitPageSetup
	CPointer unsafe.Pointer
}

type Paned struct {
	*_TraitPaned
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type PlacesSidebar struct {
	*_TraitPlacesSidebar
	*_TraitScrolledWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Plug struct {
	*_TraitPlug
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Popover struct {
	*_TraitPopover
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type PrintContext struct {
	*_TraitPrintContext
	CPointer unsafe.Pointer
}

type PrintOperation struct {
	*_TraitPrintOperation
	CPointer unsafe.Pointer
}

type PrintSettings struct {
	*_TraitPrintSettings
	CPointer unsafe.Pointer
}

type ProgressBar struct {
	*_TraitProgressBar
	*_TraitWidget
	CPointer unsafe.Pointer
}

type RadioAction struct {
	*_TraitRadioAction
	*_TraitToggleAction
	*_TraitAction
	CPointer unsafe.Pointer
}

type RadioButton struct {
	*_TraitRadioButton
	*_TraitCheckButton
	*_TraitToggleButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type RadioMenuItem struct {
	*_TraitRadioMenuItem
	*_TraitCheckMenuItem
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type RadioToolButton struct {
	*_TraitRadioToolButton
	*_TraitToggleToolButton
	*_TraitToolButton
	*_TraitToolItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Range struct {
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

type RcStyle struct {
	*_TraitRcStyle
	CPointer unsafe.Pointer
}

type RecentAction struct {
	*_TraitRecentAction
	*_TraitAction
	CPointer unsafe.Pointer
}

type RecentChooserDialog struct {
	*_TraitRecentChooserDialog
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type RecentChooserMenu struct {
	*_TraitRecentChooserMenu
	*_TraitMenu
	*_TraitMenuShell
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type RecentChooserWidget struct {
	*_TraitRecentChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type RecentFilter struct {
	*_TraitRecentFilter
	CPointer unsafe.Pointer
}

type RecentManager struct {
	*_TraitRecentManager
	CPointer unsafe.Pointer
}

type Revealer struct {
	*_TraitRevealer
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Scale struct {
	*_TraitScale
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ScaleButton struct {
	*_TraitScaleButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Scrollbar struct {
	*_TraitScrollbar
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ScrolledWindow struct {
	*_TraitScrolledWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type SearchBar struct {
	*_TraitSearchBar
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type SearchEntry struct {
	*_TraitSearchEntry
	*_TraitEntry
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Separator struct {
	*_TraitSeparator
	*_TraitWidget
	CPointer unsafe.Pointer
}

type SeparatorMenuItem struct {
	*_TraitSeparatorMenuItem
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type SeparatorToolItem struct {
	*_TraitSeparatorToolItem
	*_TraitToolItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Settings struct {
	*_TraitSettings
	CPointer unsafe.Pointer
}

type SizeGroup struct {
	*_TraitSizeGroup
	CPointer unsafe.Pointer
}

type Socket struct {
	*_TraitSocket
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type SpinButton struct {
	*_TraitSpinButton
	*_TraitEntry
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Spinner struct {
	*_TraitSpinner
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Stack struct {
	*_TraitStack
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type StackSwitcher struct {
	*_TraitStackSwitcher
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type StatusIcon struct {
	*_TraitStatusIcon
	CPointer unsafe.Pointer
}

type Statusbar struct {
	*_TraitStatusbar
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Style struct {
	*_TraitStyle
	CPointer unsafe.Pointer
}

type StyleContext struct {
	*_TraitStyleContext
	CPointer unsafe.Pointer
}

type StyleProperties struct {
	*_TraitStyleProperties
	CPointer unsafe.Pointer
}

type Switch struct {
	*_TraitSwitch
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Table struct {
	*_TraitTable
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type TearoffMenuItem struct {
	*_TraitTearoffMenuItem
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type TextBuffer struct {
	*_TraitTextBuffer
	CPointer unsafe.Pointer
}

type TextChildAnchor struct {
	*_TraitTextChildAnchor
	CPointer unsafe.Pointer
}

type TextMark struct {
	*_TraitTextMark
	CPointer unsafe.Pointer
}

type TextTag struct {
	*_TraitTextTag
	CPointer unsafe.Pointer
}

type TextTagTable struct {
	*_TraitTextTagTable
	CPointer unsafe.Pointer
}

type TextView struct {
	*_TraitTextView
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ThemingEngine struct {
	*_TraitThemingEngine
	CPointer unsafe.Pointer
}

type ToggleAction struct {
	*_TraitToggleAction
	*_TraitAction
	CPointer unsafe.Pointer
}

type ToggleButton struct {
	*_TraitToggleButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ToggleToolButton struct {
	*_TraitToggleToolButton
	*_TraitToolButton
	*_TraitToolItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ToolButton struct {
	*_TraitToolButton
	*_TraitToolItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ToolItem struct {
	*_TraitToolItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ToolItemGroup struct {
	*_TraitToolItemGroup
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type ToolPalette struct {
	*_TraitToolPalette
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Toolbar struct {
	*_TraitToolbar
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Tooltip struct {
	*_TraitTooltip
	CPointer unsafe.Pointer
}

type TreeModelFilter struct {
	*_TraitTreeModelFilter
	CPointer unsafe.Pointer
}

type TreeModelSort struct {
	*_TraitTreeModelSort
	CPointer unsafe.Pointer
}

type TreeSelection struct {
	*_TraitTreeSelection
	CPointer unsafe.Pointer
}

type TreeStore struct {
	*_TraitTreeStore
	CPointer unsafe.Pointer
}

type TreeView struct {
	*_TraitTreeView
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type TreeViewColumn struct {
	*_TraitTreeViewColumn
	CPointer unsafe.Pointer
}

type UIManager struct {
	*_TraitUIManager
	CPointer unsafe.Pointer
}

type VBox struct {
	*_TraitVBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type VButtonBox struct {
	*_TraitVButtonBox
	*_TraitButtonBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type VPaned struct {
	*_TraitVPaned
	*_TraitPaned
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type VScale struct {
	*_TraitVScale
	*_TraitScale
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

type VScrollbar struct {
	*_TraitVScrollbar
	*_TraitScrollbar
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

type VSeparator struct {
	*_TraitVSeparator
	*_TraitSeparator
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Viewport struct {
	*_TraitViewport
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type VolumeButton struct {
	*_TraitVolumeButton
	*_TraitScaleButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Widget struct {
	*_TraitWidget
	CPointer unsafe.Pointer
}

type Window struct {
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

type WindowGroup struct {
	*_TraitWindowGroup
	CPointer unsafe.Pointer
}
