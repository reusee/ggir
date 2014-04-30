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

func NewAboutDialogFromCPointer(p unsafe.Pointer) *AboutDialog {
	return &AboutDialog{
		&_TraitAboutDialog{(*C.GtkAboutDialog)(p)},
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type AccelGroup struct {
	*_TraitAccelGroup
	CPointer unsafe.Pointer
}

func NewAccelGroupFromCPointer(p unsafe.Pointer) *AccelGroup {
	return &AccelGroup{
		&_TraitAccelGroup{(*C.GtkAccelGroup)(p)},
		p,
	}
}

type AccelLabel struct {
	*_TraitAccelLabel
	*_TraitLabel
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewAccelLabelFromCPointer(p unsafe.Pointer) *AccelLabel {
	return &AccelLabel{
		&_TraitAccelLabel{(*C.GtkAccelLabel)(p)},
		&_TraitLabel{(*C.GtkLabel)(p)},
		&_TraitMisc{(*C.GtkMisc)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type AccelMap struct {
	*_TraitAccelMap
	CPointer unsafe.Pointer
}

func NewAccelMapFromCPointer(p unsafe.Pointer) *AccelMap {
	return &AccelMap{
		&_TraitAccelMap{(*C.GtkAccelMap)(p)},
		p,
	}
}

type Accessible struct {
	*_TraitAccessible
	CPointer unsafe.Pointer
}

func NewAccessibleFromCPointer(p unsafe.Pointer) *Accessible {
	return &Accessible{
		&_TraitAccessible{(*C.GtkAccessible)(p)},
		p,
	}
}

type Action struct {
	*_TraitAction
	CPointer unsafe.Pointer
}

func NewActionFromCPointer(p unsafe.Pointer) *Action {
	return &Action{
		&_TraitAction{(*C.GtkAction)(p)},
		p,
	}
}

type ActionBar struct {
	*_TraitActionBar
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewActionBarFromCPointer(p unsafe.Pointer) *ActionBar {
	return &ActionBar{
		&_TraitActionBar{(*C.GtkActionBar)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ActionGroup struct {
	*_TraitActionGroup
	CPointer unsafe.Pointer
}

func NewActionGroupFromCPointer(p unsafe.Pointer) *ActionGroup {
	return &ActionGroup{
		&_TraitActionGroup{(*C.GtkActionGroup)(p)},
		p,
	}
}

type Adjustment struct {
	*_TraitAdjustment
	CPointer unsafe.Pointer
}

func NewAdjustmentFromCPointer(p unsafe.Pointer) *Adjustment {
	return &Adjustment{
		&_TraitAdjustment{(*C.GtkAdjustment)(p)},
		p,
	}
}

type Alignment struct {
	*_TraitAlignment
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewAlignmentFromCPointer(p unsafe.Pointer) *Alignment {
	return &Alignment{
		&_TraitAlignment{(*C.GtkAlignment)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type AppChooserButton struct {
	*_TraitAppChooserButton
	*_TraitComboBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewAppChooserButtonFromCPointer(p unsafe.Pointer) *AppChooserButton {
	return &AppChooserButton{
		&_TraitAppChooserButton{(*C.GtkAppChooserButton)(p)},
		&_TraitComboBox{(*C.GtkComboBox)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewAppChooserDialogFromCPointer(p unsafe.Pointer) *AppChooserDialog {
	return &AppChooserDialog{
		&_TraitAppChooserDialog{(*C.GtkAppChooserDialog)(p)},
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type AppChooserWidget struct {
	*_TraitAppChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewAppChooserWidgetFromCPointer(p unsafe.Pointer) *AppChooserWidget {
	return &AppChooserWidget{
		&_TraitAppChooserWidget{(*C.GtkAppChooserWidget)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Application struct {
	*_TraitApplication
	CPointer unsafe.Pointer
}

func NewApplicationFromCPointer(p unsafe.Pointer) *Application {
	return &Application{
		&_TraitApplication{(*C.GtkApplication)(p)},
		p,
	}
}

type ApplicationWindow struct {
	*_TraitApplicationWindow
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewApplicationWindowFromCPointer(p unsafe.Pointer) *ApplicationWindow {
	return &ApplicationWindow{
		&_TraitApplicationWindow{(*C.GtkApplicationWindow)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Arrow struct {
	*_TraitArrow
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewArrowFromCPointer(p unsafe.Pointer) *Arrow {
	return &Arrow{
		&_TraitArrow{(*C.GtkArrow)(p)},
		&_TraitMisc{(*C.GtkMisc)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type AspectFrame struct {
	*_TraitAspectFrame
	*_TraitFrame
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewAspectFrameFromCPointer(p unsafe.Pointer) *AspectFrame {
	return &AspectFrame{
		&_TraitAspectFrame{(*C.GtkAspectFrame)(p)},
		&_TraitFrame{(*C.GtkFrame)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Assistant struct {
	*_TraitAssistant
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewAssistantFromCPointer(p unsafe.Pointer) *Assistant {
	return &Assistant{
		&_TraitAssistant{(*C.GtkAssistant)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Bin struct {
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewBinFromCPointer(p unsafe.Pointer) *Bin {
	return &Bin{
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Box struct {
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewBoxFromCPointer(p unsafe.Pointer) *Box {
	return &Box{
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Builder struct {
	*_TraitBuilder
	CPointer unsafe.Pointer
}

func NewBuilderFromCPointer(p unsafe.Pointer) *Builder {
	return &Builder{
		&_TraitBuilder{(*C.GtkBuilder)(p)},
		p,
	}
}

type Button struct {
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewButtonFromCPointer(p unsafe.Pointer) *Button {
	return &Button{
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ButtonBox struct {
	*_TraitButtonBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewButtonBoxFromCPointer(p unsafe.Pointer) *ButtonBox {
	return &ButtonBox{
		&_TraitButtonBox{(*C.GtkButtonBox)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Calendar struct {
	*_TraitCalendar
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewCalendarFromCPointer(p unsafe.Pointer) *Calendar {
	return &Calendar{
		&_TraitCalendar{(*C.GtkCalendar)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type CellArea struct {
	*_TraitCellArea
	CPointer unsafe.Pointer
}

func NewCellAreaFromCPointer(p unsafe.Pointer) *CellArea {
	return &CellArea{
		&_TraitCellArea{(*C.GtkCellArea)(p)},
		p,
	}
}

type CellAreaBox struct {
	*_TraitCellAreaBox
	*_TraitCellArea
	CPointer unsafe.Pointer
}

func NewCellAreaBoxFromCPointer(p unsafe.Pointer) *CellAreaBox {
	return &CellAreaBox{
		&_TraitCellAreaBox{(*C.GtkCellAreaBox)(p)},
		&_TraitCellArea{(*C.GtkCellArea)(p)},
		p,
	}
}

type CellAreaContext struct {
	*_TraitCellAreaContext
	CPointer unsafe.Pointer
}

func NewCellAreaContextFromCPointer(p unsafe.Pointer) *CellAreaContext {
	return &CellAreaContext{
		&_TraitCellAreaContext{(*C.GtkCellAreaContext)(p)},
		p,
	}
}

type CellRenderer struct {
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

func NewCellRendererFromCPointer(p unsafe.Pointer) *CellRenderer {
	return &CellRenderer{
		&_TraitCellRenderer{(*C.GtkCellRenderer)(p)},
		p,
	}
}

type CellRendererAccel struct {
	*_TraitCellRendererAccel
	*_TraitCellRendererText
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

func NewCellRendererAccelFromCPointer(p unsafe.Pointer) *CellRendererAccel {
	return &CellRendererAccel{
		&_TraitCellRendererAccel{(*C.GtkCellRendererAccel)(p)},
		&_TraitCellRendererText{(*C.GtkCellRendererText)(p)},
		&_TraitCellRenderer{(*C.GtkCellRenderer)(p)},
		p,
	}
}

type CellRendererCombo struct {
	*_TraitCellRendererCombo
	*_TraitCellRendererText
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

func NewCellRendererComboFromCPointer(p unsafe.Pointer) *CellRendererCombo {
	return &CellRendererCombo{
		&_TraitCellRendererCombo{(*C.GtkCellRendererCombo)(p)},
		&_TraitCellRendererText{(*C.GtkCellRendererText)(p)},
		&_TraitCellRenderer{(*C.GtkCellRenderer)(p)},
		p,
	}
}

type CellRendererPixbuf struct {
	*_TraitCellRendererPixbuf
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

func NewCellRendererPixbufFromCPointer(p unsafe.Pointer) *CellRendererPixbuf {
	return &CellRendererPixbuf{
		&_TraitCellRendererPixbuf{(*C.GtkCellRendererPixbuf)(p)},
		&_TraitCellRenderer{(*C.GtkCellRenderer)(p)},
		p,
	}
}

type CellRendererProgress struct {
	*_TraitCellRendererProgress
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

func NewCellRendererProgressFromCPointer(p unsafe.Pointer) *CellRendererProgress {
	return &CellRendererProgress{
		&_TraitCellRendererProgress{(*C.GtkCellRendererProgress)(p)},
		&_TraitCellRenderer{(*C.GtkCellRenderer)(p)},
		p,
	}
}

type CellRendererSpin struct {
	*_TraitCellRendererSpin
	*_TraitCellRendererText
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

func NewCellRendererSpinFromCPointer(p unsafe.Pointer) *CellRendererSpin {
	return &CellRendererSpin{
		&_TraitCellRendererSpin{(*C.GtkCellRendererSpin)(p)},
		&_TraitCellRendererText{(*C.GtkCellRendererText)(p)},
		&_TraitCellRenderer{(*C.GtkCellRenderer)(p)},
		p,
	}
}

type CellRendererSpinner struct {
	*_TraitCellRendererSpinner
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

func NewCellRendererSpinnerFromCPointer(p unsafe.Pointer) *CellRendererSpinner {
	return &CellRendererSpinner{
		&_TraitCellRendererSpinner{(*C.GtkCellRendererSpinner)(p)},
		&_TraitCellRenderer{(*C.GtkCellRenderer)(p)},
		p,
	}
}

type CellRendererText struct {
	*_TraitCellRendererText
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

func NewCellRendererTextFromCPointer(p unsafe.Pointer) *CellRendererText {
	return &CellRendererText{
		&_TraitCellRendererText{(*C.GtkCellRendererText)(p)},
		&_TraitCellRenderer{(*C.GtkCellRenderer)(p)},
		p,
	}
}

type CellRendererToggle struct {
	*_TraitCellRendererToggle
	*_TraitCellRenderer
	CPointer unsafe.Pointer
}

func NewCellRendererToggleFromCPointer(p unsafe.Pointer) *CellRendererToggle {
	return &CellRendererToggle{
		&_TraitCellRendererToggle{(*C.GtkCellRendererToggle)(p)},
		&_TraitCellRenderer{(*C.GtkCellRenderer)(p)},
		p,
	}
}

type CellView struct {
	*_TraitCellView
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewCellViewFromCPointer(p unsafe.Pointer) *CellView {
	return &CellView{
		&_TraitCellView{(*C.GtkCellView)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewCheckButtonFromCPointer(p unsafe.Pointer) *CheckButton {
	return &CheckButton{
		&_TraitCheckButton{(*C.GtkCheckButton)(p)},
		&_TraitToggleButton{(*C.GtkToggleButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type CheckMenuItem struct {
	*_TraitCheckMenuItem
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewCheckMenuItemFromCPointer(p unsafe.Pointer) *CheckMenuItem {
	return &CheckMenuItem{
		&_TraitCheckMenuItem{(*C.GtkCheckMenuItem)(p)},
		&_TraitMenuItem{(*C.GtkMenuItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Clipboard struct {
	*_TraitClipboard
	CPointer unsafe.Pointer
}

func NewClipboardFromCPointer(p unsafe.Pointer) *Clipboard {
	return &Clipboard{
		&_TraitClipboard{(*C.GtkClipboard)(p)},
		p,
	}
}

type ColorButton struct {
	*_TraitColorButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewColorButtonFromCPointer(p unsafe.Pointer) *ColorButton {
	return &ColorButton{
		&_TraitColorButton{(*C.GtkColorButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewColorChooserDialogFromCPointer(p unsafe.Pointer) *ColorChooserDialog {
	return &ColorChooserDialog{
		&_TraitColorChooserDialog{(*C.GtkColorChooserDialog)(p)},
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ColorChooserWidget struct {
	*_TraitColorChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewColorChooserWidgetFromCPointer(p unsafe.Pointer) *ColorChooserWidget {
	return &ColorChooserWidget{
		&_TraitColorChooserWidget{(*C.GtkColorChooserWidget)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ColorSelection struct {
	*_TraitColorSelection
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewColorSelectionFromCPointer(p unsafe.Pointer) *ColorSelection {
	return &ColorSelection{
		&_TraitColorSelection{(*C.GtkColorSelection)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewColorSelectionDialogFromCPointer(p unsafe.Pointer) *ColorSelectionDialog {
	return &ColorSelectionDialog{
		&_TraitColorSelectionDialog{(*C.GtkColorSelectionDialog)(p)},
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ComboBox struct {
	*_TraitComboBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewComboBoxFromCPointer(p unsafe.Pointer) *ComboBox {
	return &ComboBox{
		&_TraitComboBox{(*C.GtkComboBox)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ComboBoxText struct {
	*_TraitComboBoxText
	*_TraitComboBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewComboBoxTextFromCPointer(p unsafe.Pointer) *ComboBoxText {
	return &ComboBoxText{
		&_TraitComboBoxText{(*C.GtkComboBoxText)(p)},
		&_TraitComboBox{(*C.GtkComboBox)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Container struct {
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewContainerFromCPointer(p unsafe.Pointer) *Container {
	return &Container{
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type CssProvider struct {
	*_TraitCssProvider
	CPointer unsafe.Pointer
}

func NewCssProviderFromCPointer(p unsafe.Pointer) *CssProvider {
	return &CssProvider{
		&_TraitCssProvider{(*C.GtkCssProvider)(p)},
		p,
	}
}

type Dialog struct {
	*_TraitDialog
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewDialogFromCPointer(p unsafe.Pointer) *Dialog {
	return &Dialog{
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type DrawingArea struct {
	*_TraitDrawingArea
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewDrawingAreaFromCPointer(p unsafe.Pointer) *DrawingArea {
	return &DrawingArea{
		&_TraitDrawingArea{(*C.GtkDrawingArea)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Entry struct {
	*_TraitEntry
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewEntryFromCPointer(p unsafe.Pointer) *Entry {
	return &Entry{
		&_TraitEntry{(*C.GtkEntry)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type EntryBuffer struct {
	*_TraitEntryBuffer
	CPointer unsafe.Pointer
}

func NewEntryBufferFromCPointer(p unsafe.Pointer) *EntryBuffer {
	return &EntryBuffer{
		&_TraitEntryBuffer{(*C.GtkEntryBuffer)(p)},
		p,
	}
}

type EntryCompletion struct {
	*_TraitEntryCompletion
	CPointer unsafe.Pointer
}

func NewEntryCompletionFromCPointer(p unsafe.Pointer) *EntryCompletion {
	return &EntryCompletion{
		&_TraitEntryCompletion{(*C.GtkEntryCompletion)(p)},
		p,
	}
}

type EventBox struct {
	*_TraitEventBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewEventBoxFromCPointer(p unsafe.Pointer) *EventBox {
	return &EventBox{
		&_TraitEventBox{(*C.GtkEventBox)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Expander struct {
	*_TraitExpander
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewExpanderFromCPointer(p unsafe.Pointer) *Expander {
	return &Expander{
		&_TraitExpander{(*C.GtkExpander)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type FileChooserButton struct {
	*_TraitFileChooserButton
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewFileChooserButtonFromCPointer(p unsafe.Pointer) *FileChooserButton {
	return &FileChooserButton{
		&_TraitFileChooserButton{(*C.GtkFileChooserButton)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewFileChooserDialogFromCPointer(p unsafe.Pointer) *FileChooserDialog {
	return &FileChooserDialog{
		&_TraitFileChooserDialog{(*C.GtkFileChooserDialog)(p)},
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type FileChooserWidget struct {
	*_TraitFileChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewFileChooserWidgetFromCPointer(p unsafe.Pointer) *FileChooserWidget {
	return &FileChooserWidget{
		&_TraitFileChooserWidget{(*C.GtkFileChooserWidget)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type FileFilter struct {
	*_TraitFileFilter
	CPointer unsafe.Pointer
}

func NewFileFilterFromCPointer(p unsafe.Pointer) *FileFilter {
	return &FileFilter{
		&_TraitFileFilter{(*C.GtkFileFilter)(p)},
		p,
	}
}

type Fixed struct {
	*_TraitFixed
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewFixedFromCPointer(p unsafe.Pointer) *Fixed {
	return &Fixed{
		&_TraitFixed{(*C.GtkFixed)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type FlowBox struct {
	*_TraitFlowBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewFlowBoxFromCPointer(p unsafe.Pointer) *FlowBox {
	return &FlowBox{
		&_TraitFlowBox{(*C.GtkFlowBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type FlowBoxChild struct {
	*_TraitFlowBoxChild
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewFlowBoxChildFromCPointer(p unsafe.Pointer) *FlowBoxChild {
	return &FlowBoxChild{
		&_TraitFlowBoxChild{(*C.GtkFlowBoxChild)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type FontButton struct {
	*_TraitFontButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewFontButtonFromCPointer(p unsafe.Pointer) *FontButton {
	return &FontButton{
		&_TraitFontButton{(*C.GtkFontButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewFontChooserDialogFromCPointer(p unsafe.Pointer) *FontChooserDialog {
	return &FontChooserDialog{
		&_TraitFontChooserDialog{(*C.GtkFontChooserDialog)(p)},
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type FontChooserWidget struct {
	*_TraitFontChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewFontChooserWidgetFromCPointer(p unsafe.Pointer) *FontChooserWidget {
	return &FontChooserWidget{
		&_TraitFontChooserWidget{(*C.GtkFontChooserWidget)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type FontSelection struct {
	*_TraitFontSelection
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewFontSelectionFromCPointer(p unsafe.Pointer) *FontSelection {
	return &FontSelection{
		&_TraitFontSelection{(*C.GtkFontSelection)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewFontSelectionDialogFromCPointer(p unsafe.Pointer) *FontSelectionDialog {
	return &FontSelectionDialog{
		&_TraitFontSelectionDialog{(*C.GtkFontSelectionDialog)(p)},
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Frame struct {
	*_TraitFrame
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewFrameFromCPointer(p unsafe.Pointer) *Frame {
	return &Frame{
		&_TraitFrame{(*C.GtkFrame)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Grid struct {
	*_TraitGrid
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewGridFromCPointer(p unsafe.Pointer) *Grid {
	return &Grid{
		&_TraitGrid{(*C.GtkGrid)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type HBox struct {
	*_TraitHBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewHBoxFromCPointer(p unsafe.Pointer) *HBox {
	return &HBox{
		&_TraitHBox{(*C.GtkHBox)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type HButtonBox struct {
	*_TraitHButtonBox
	*_TraitButtonBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewHButtonBoxFromCPointer(p unsafe.Pointer) *HButtonBox {
	return &HButtonBox{
		&_TraitHButtonBox{(*C.GtkHButtonBox)(p)},
		&_TraitButtonBox{(*C.GtkButtonBox)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type HPaned struct {
	*_TraitHPaned
	*_TraitPaned
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewHPanedFromCPointer(p unsafe.Pointer) *HPaned {
	return &HPaned{
		&_TraitHPaned{(*C.GtkHPaned)(p)},
		&_TraitPaned{(*C.GtkPaned)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type HSV struct {
	*_TraitHSV
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewHSVFromCPointer(p unsafe.Pointer) *HSV {
	return &HSV{
		&_TraitHSV{(*C.GtkHSV)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type HScale struct {
	*_TraitHScale
	*_TraitScale
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewHScaleFromCPointer(p unsafe.Pointer) *HScale {
	return &HScale{
		&_TraitHScale{(*C.GtkHScale)(p)},
		&_TraitScale{(*C.GtkScale)(p)},
		&_TraitRange{(*C.GtkRange)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type HScrollbar struct {
	*_TraitHScrollbar
	*_TraitScrollbar
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewHScrollbarFromCPointer(p unsafe.Pointer) *HScrollbar {
	return &HScrollbar{
		&_TraitHScrollbar{(*C.GtkHScrollbar)(p)},
		&_TraitScrollbar{(*C.GtkScrollbar)(p)},
		&_TraitRange{(*C.GtkRange)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type HSeparator struct {
	*_TraitHSeparator
	*_TraitSeparator
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewHSeparatorFromCPointer(p unsafe.Pointer) *HSeparator {
	return &HSeparator{
		&_TraitHSeparator{(*C.GtkHSeparator)(p)},
		&_TraitSeparator{(*C.GtkSeparator)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type HandleBox struct {
	*_TraitHandleBox
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewHandleBoxFromCPointer(p unsafe.Pointer) *HandleBox {
	return &HandleBox{
		&_TraitHandleBox{(*C.GtkHandleBox)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type HeaderBar struct {
	*_TraitHeaderBar
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewHeaderBarFromCPointer(p unsafe.Pointer) *HeaderBar {
	return &HeaderBar{
		&_TraitHeaderBar{(*C.GtkHeaderBar)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type IMContext struct {
	*_TraitIMContext
	CPointer unsafe.Pointer
}

func NewIMContextFromCPointer(p unsafe.Pointer) *IMContext {
	return &IMContext{
		&_TraitIMContext{(*C.GtkIMContext)(p)},
		p,
	}
}

type IMContextSimple struct {
	*_TraitIMContextSimple
	*_TraitIMContext
	CPointer unsafe.Pointer
}

func NewIMContextSimpleFromCPointer(p unsafe.Pointer) *IMContextSimple {
	return &IMContextSimple{
		&_TraitIMContextSimple{(*C.GtkIMContextSimple)(p)},
		&_TraitIMContext{(*C.GtkIMContext)(p)},
		p,
	}
}

type IMMulticontext struct {
	*_TraitIMMulticontext
	*_TraitIMContext
	CPointer unsafe.Pointer
}

func NewIMMulticontextFromCPointer(p unsafe.Pointer) *IMMulticontext {
	return &IMMulticontext{
		&_TraitIMMulticontext{(*C.GtkIMMulticontext)(p)},
		&_TraitIMContext{(*C.GtkIMContext)(p)},
		p,
	}
}

type IconFactory struct {
	*_TraitIconFactory
	CPointer unsafe.Pointer
}

func NewIconFactoryFromCPointer(p unsafe.Pointer) *IconFactory {
	return &IconFactory{
		&_TraitIconFactory{(*C.GtkIconFactory)(p)},
		p,
	}
}

type IconInfo struct {
	*_TraitIconInfo
	CPointer unsafe.Pointer
}

func NewIconInfoFromCPointer(p unsafe.Pointer) *IconInfo {
	return &IconInfo{
		&_TraitIconInfo{(*C.GtkIconInfo)(p)},
		p,
	}
}

type IconTheme struct {
	*_TraitIconTheme
	CPointer unsafe.Pointer
}

func NewIconThemeFromCPointer(p unsafe.Pointer) *IconTheme {
	return &IconTheme{
		&_TraitIconTheme{(*C.GtkIconTheme)(p)},
		p,
	}
}

type IconView struct {
	*_TraitIconView
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewIconViewFromCPointer(p unsafe.Pointer) *IconView {
	return &IconView{
		&_TraitIconView{(*C.GtkIconView)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Image struct {
	*_TraitImage
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewImageFromCPointer(p unsafe.Pointer) *Image {
	return &Image{
		&_TraitImage{(*C.GtkImage)(p)},
		&_TraitMisc{(*C.GtkMisc)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ImageMenuItem struct {
	*_TraitImageMenuItem
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewImageMenuItemFromCPointer(p unsafe.Pointer) *ImageMenuItem {
	return &ImageMenuItem{
		&_TraitImageMenuItem{(*C.GtkImageMenuItem)(p)},
		&_TraitMenuItem{(*C.GtkMenuItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type InfoBar struct {
	*_TraitInfoBar
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewInfoBarFromCPointer(p unsafe.Pointer) *InfoBar {
	return &InfoBar{
		&_TraitInfoBar{(*C.GtkInfoBar)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Invisible struct {
	*_TraitInvisible
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewInvisibleFromCPointer(p unsafe.Pointer) *Invisible {
	return &Invisible{
		&_TraitInvisible{(*C.GtkInvisible)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Label struct {
	*_TraitLabel
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewLabelFromCPointer(p unsafe.Pointer) *Label {
	return &Label{
		&_TraitLabel{(*C.GtkLabel)(p)},
		&_TraitMisc{(*C.GtkMisc)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Layout struct {
	*_TraitLayout
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewLayoutFromCPointer(p unsafe.Pointer) *Layout {
	return &Layout{
		&_TraitLayout{(*C.GtkLayout)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type LevelBar struct {
	*_TraitLevelBar
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewLevelBarFromCPointer(p unsafe.Pointer) *LevelBar {
	return &LevelBar{
		&_TraitLevelBar{(*C.GtkLevelBar)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type LinkButton struct {
	*_TraitLinkButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewLinkButtonFromCPointer(p unsafe.Pointer) *LinkButton {
	return &LinkButton{
		&_TraitLinkButton{(*C.GtkLinkButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ListBox struct {
	*_TraitListBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewListBoxFromCPointer(p unsafe.Pointer) *ListBox {
	return &ListBox{
		&_TraitListBox{(*C.GtkListBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ListBoxRow struct {
	*_TraitListBoxRow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewListBoxRowFromCPointer(p unsafe.Pointer) *ListBoxRow {
	return &ListBoxRow{
		&_TraitListBoxRow{(*C.GtkListBoxRow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ListStore struct {
	*_TraitListStore
	CPointer unsafe.Pointer
}

func NewListStoreFromCPointer(p unsafe.Pointer) *ListStore {
	return &ListStore{
		&_TraitListStore{(*C.GtkListStore)(p)},
		p,
	}
}

type LockButton struct {
	*_TraitLockButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewLockButtonFromCPointer(p unsafe.Pointer) *LockButton {
	return &LockButton{
		&_TraitLockButton{(*C.GtkLockButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Menu struct {
	*_TraitMenu
	*_TraitMenuShell
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewMenuFromCPointer(p unsafe.Pointer) *Menu {
	return &Menu{
		&_TraitMenu{(*C.GtkMenu)(p)},
		&_TraitMenuShell{(*C.GtkMenuShell)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type MenuBar struct {
	*_TraitMenuBar
	*_TraitMenuShell
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewMenuBarFromCPointer(p unsafe.Pointer) *MenuBar {
	return &MenuBar{
		&_TraitMenuBar{(*C.GtkMenuBar)(p)},
		&_TraitMenuShell{(*C.GtkMenuShell)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewMenuButtonFromCPointer(p unsafe.Pointer) *MenuButton {
	return &MenuButton{
		&_TraitMenuButton{(*C.GtkMenuButton)(p)},
		&_TraitToggleButton{(*C.GtkToggleButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type MenuItem struct {
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewMenuItemFromCPointer(p unsafe.Pointer) *MenuItem {
	return &MenuItem{
		&_TraitMenuItem{(*C.GtkMenuItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type MenuShell struct {
	*_TraitMenuShell
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewMenuShellFromCPointer(p unsafe.Pointer) *MenuShell {
	return &MenuShell{
		&_TraitMenuShell{(*C.GtkMenuShell)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewMenuToolButtonFromCPointer(p unsafe.Pointer) *MenuToolButton {
	return &MenuToolButton{
		&_TraitMenuToolButton{(*C.GtkMenuToolButton)(p)},
		&_TraitToolButton{(*C.GtkToolButton)(p)},
		&_TraitToolItem{(*C.GtkToolItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewMessageDialogFromCPointer(p unsafe.Pointer) *MessageDialog {
	return &MessageDialog{
		&_TraitMessageDialog{(*C.GtkMessageDialog)(p)},
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Misc struct {
	*_TraitMisc
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewMiscFromCPointer(p unsafe.Pointer) *Misc {
	return &Misc{
		&_TraitMisc{(*C.GtkMisc)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type MountOperation struct {
	*_TraitMountOperation
	CPointer unsafe.Pointer
}

func NewMountOperationFromCPointer(p unsafe.Pointer) *MountOperation {
	return &MountOperation{
		&_TraitMountOperation{(*C.GtkMountOperation)(p)},
		p,
	}
}

type Notebook struct {
	*_TraitNotebook
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewNotebookFromCPointer(p unsafe.Pointer) *Notebook {
	return &Notebook{
		&_TraitNotebook{(*C.GtkNotebook)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type NumerableIcon struct {
	*_TraitNumerableIcon
	CPointer unsafe.Pointer
}

func NewNumerableIconFromCPointer(p unsafe.Pointer) *NumerableIcon {
	return &NumerableIcon{
		&_TraitNumerableIcon{(*C.GtkNumerableIcon)(p)},
		p,
	}
}

type OffscreenWindow struct {
	*_TraitOffscreenWindow
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewOffscreenWindowFromCPointer(p unsafe.Pointer) *OffscreenWindow {
	return &OffscreenWindow{
		&_TraitOffscreenWindow{(*C.GtkOffscreenWindow)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Overlay struct {
	*_TraitOverlay
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewOverlayFromCPointer(p unsafe.Pointer) *Overlay {
	return &Overlay{
		&_TraitOverlay{(*C.GtkOverlay)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type PageSetup struct {
	*_TraitPageSetup
	CPointer unsafe.Pointer
}

func NewPageSetupFromCPointer(p unsafe.Pointer) *PageSetup {
	return &PageSetup{
		&_TraitPageSetup{(*C.GtkPageSetup)(p)},
		p,
	}
}

type Paned struct {
	*_TraitPaned
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewPanedFromCPointer(p unsafe.Pointer) *Paned {
	return &Paned{
		&_TraitPaned{(*C.GtkPaned)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type PlacesSidebar struct {
	*_TraitPlacesSidebar
	*_TraitScrolledWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewPlacesSidebarFromCPointer(p unsafe.Pointer) *PlacesSidebar {
	return &PlacesSidebar{
		&_TraitPlacesSidebar{(*C.GtkPlacesSidebar)(p)},
		&_TraitScrolledWindow{(*C.GtkScrolledWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Plug struct {
	*_TraitPlug
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewPlugFromCPointer(p unsafe.Pointer) *Plug {
	return &Plug{
		&_TraitPlug{(*C.GtkPlug)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Popover struct {
	*_TraitPopover
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewPopoverFromCPointer(p unsafe.Pointer) *Popover {
	return &Popover{
		&_TraitPopover{(*C.GtkPopover)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type PrintContext struct {
	*_TraitPrintContext
	CPointer unsafe.Pointer
}

func NewPrintContextFromCPointer(p unsafe.Pointer) *PrintContext {
	return &PrintContext{
		&_TraitPrintContext{(*C.GtkPrintContext)(p)},
		p,
	}
}

type PrintOperation struct {
	*_TraitPrintOperation
	CPointer unsafe.Pointer
}

func NewPrintOperationFromCPointer(p unsafe.Pointer) *PrintOperation {
	return &PrintOperation{
		&_TraitPrintOperation{(*C.GtkPrintOperation)(p)},
		p,
	}
}

type PrintSettings struct {
	*_TraitPrintSettings
	CPointer unsafe.Pointer
}

func NewPrintSettingsFromCPointer(p unsafe.Pointer) *PrintSettings {
	return &PrintSettings{
		&_TraitPrintSettings{(*C.GtkPrintSettings)(p)},
		p,
	}
}

type ProgressBar struct {
	*_TraitProgressBar
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewProgressBarFromCPointer(p unsafe.Pointer) *ProgressBar {
	return &ProgressBar{
		&_TraitProgressBar{(*C.GtkProgressBar)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type RadioAction struct {
	*_TraitRadioAction
	*_TraitToggleAction
	*_TraitAction
	CPointer unsafe.Pointer
}

func NewRadioActionFromCPointer(p unsafe.Pointer) *RadioAction {
	return &RadioAction{
		&_TraitRadioAction{(*C.GtkRadioAction)(p)},
		&_TraitToggleAction{(*C.GtkToggleAction)(p)},
		&_TraitAction{(*C.GtkAction)(p)},
		p,
	}
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

func NewRadioButtonFromCPointer(p unsafe.Pointer) *RadioButton {
	return &RadioButton{
		&_TraitRadioButton{(*C.GtkRadioButton)(p)},
		&_TraitCheckButton{(*C.GtkCheckButton)(p)},
		&_TraitToggleButton{(*C.GtkToggleButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewRadioMenuItemFromCPointer(p unsafe.Pointer) *RadioMenuItem {
	return &RadioMenuItem{
		&_TraitRadioMenuItem{(*C.GtkRadioMenuItem)(p)},
		&_TraitCheckMenuItem{(*C.GtkCheckMenuItem)(p)},
		&_TraitMenuItem{(*C.GtkMenuItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewRadioToolButtonFromCPointer(p unsafe.Pointer) *RadioToolButton {
	return &RadioToolButton{
		&_TraitRadioToolButton{(*C.GtkRadioToolButton)(p)},
		&_TraitToggleToolButton{(*C.GtkToggleToolButton)(p)},
		&_TraitToolButton{(*C.GtkToolButton)(p)},
		&_TraitToolItem{(*C.GtkToolItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Range struct {
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewRangeFromCPointer(p unsafe.Pointer) *Range {
	return &Range{
		&_TraitRange{(*C.GtkRange)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type RcStyle struct {
	*_TraitRcStyle
	CPointer unsafe.Pointer
}

func NewRcStyleFromCPointer(p unsafe.Pointer) *RcStyle {
	return &RcStyle{
		&_TraitRcStyle{(*C.GtkRcStyle)(p)},
		p,
	}
}

type RecentAction struct {
	*_TraitRecentAction
	*_TraitAction
	CPointer unsafe.Pointer
}

func NewRecentActionFromCPointer(p unsafe.Pointer) *RecentAction {
	return &RecentAction{
		&_TraitRecentAction{(*C.GtkRecentAction)(p)},
		&_TraitAction{(*C.GtkAction)(p)},
		p,
	}
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

func NewRecentChooserDialogFromCPointer(p unsafe.Pointer) *RecentChooserDialog {
	return &RecentChooserDialog{
		&_TraitRecentChooserDialog{(*C.GtkRecentChooserDialog)(p)},
		&_TraitDialog{(*C.GtkDialog)(p)},
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type RecentChooserMenu struct {
	*_TraitRecentChooserMenu
	*_TraitMenu
	*_TraitMenuShell
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewRecentChooserMenuFromCPointer(p unsafe.Pointer) *RecentChooserMenu {
	return &RecentChooserMenu{
		&_TraitRecentChooserMenu{(*C.GtkRecentChooserMenu)(p)},
		&_TraitMenu{(*C.GtkMenu)(p)},
		&_TraitMenuShell{(*C.GtkMenuShell)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type RecentChooserWidget struct {
	*_TraitRecentChooserWidget
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewRecentChooserWidgetFromCPointer(p unsafe.Pointer) *RecentChooserWidget {
	return &RecentChooserWidget{
		&_TraitRecentChooserWidget{(*C.GtkRecentChooserWidget)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type RecentFilter struct {
	*_TraitRecentFilter
	CPointer unsafe.Pointer
}

func NewRecentFilterFromCPointer(p unsafe.Pointer) *RecentFilter {
	return &RecentFilter{
		&_TraitRecentFilter{(*C.GtkRecentFilter)(p)},
		p,
	}
}

type RecentManager struct {
	*_TraitRecentManager
	CPointer unsafe.Pointer
}

func NewRecentManagerFromCPointer(p unsafe.Pointer) *RecentManager {
	return &RecentManager{
		&_TraitRecentManager{(*C.GtkRecentManager)(p)},
		p,
	}
}

type Revealer struct {
	*_TraitRevealer
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewRevealerFromCPointer(p unsafe.Pointer) *Revealer {
	return &Revealer{
		&_TraitRevealer{(*C.GtkRevealer)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Scale struct {
	*_TraitScale
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewScaleFromCPointer(p unsafe.Pointer) *Scale {
	return &Scale{
		&_TraitScale{(*C.GtkScale)(p)},
		&_TraitRange{(*C.GtkRange)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ScaleButton struct {
	*_TraitScaleButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewScaleButtonFromCPointer(p unsafe.Pointer) *ScaleButton {
	return &ScaleButton{
		&_TraitScaleButton{(*C.GtkScaleButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Scrollbar struct {
	*_TraitScrollbar
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewScrollbarFromCPointer(p unsafe.Pointer) *Scrollbar {
	return &Scrollbar{
		&_TraitScrollbar{(*C.GtkScrollbar)(p)},
		&_TraitRange{(*C.GtkRange)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ScrolledWindow struct {
	*_TraitScrolledWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewScrolledWindowFromCPointer(p unsafe.Pointer) *ScrolledWindow {
	return &ScrolledWindow{
		&_TraitScrolledWindow{(*C.GtkScrolledWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type SearchBar struct {
	*_TraitSearchBar
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewSearchBarFromCPointer(p unsafe.Pointer) *SearchBar {
	return &SearchBar{
		&_TraitSearchBar{(*C.GtkSearchBar)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type SearchEntry struct {
	*_TraitSearchEntry
	*_TraitEntry
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewSearchEntryFromCPointer(p unsafe.Pointer) *SearchEntry {
	return &SearchEntry{
		&_TraitSearchEntry{(*C.GtkSearchEntry)(p)},
		&_TraitEntry{(*C.GtkEntry)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Separator struct {
	*_TraitSeparator
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewSeparatorFromCPointer(p unsafe.Pointer) *Separator {
	return &Separator{
		&_TraitSeparator{(*C.GtkSeparator)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type SeparatorMenuItem struct {
	*_TraitSeparatorMenuItem
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewSeparatorMenuItemFromCPointer(p unsafe.Pointer) *SeparatorMenuItem {
	return &SeparatorMenuItem{
		&_TraitSeparatorMenuItem{(*C.GtkSeparatorMenuItem)(p)},
		&_TraitMenuItem{(*C.GtkMenuItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type SeparatorToolItem struct {
	*_TraitSeparatorToolItem
	*_TraitToolItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewSeparatorToolItemFromCPointer(p unsafe.Pointer) *SeparatorToolItem {
	return &SeparatorToolItem{
		&_TraitSeparatorToolItem{(*C.GtkSeparatorToolItem)(p)},
		&_TraitToolItem{(*C.GtkToolItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Settings struct {
	*_TraitSettings
	CPointer unsafe.Pointer
}

func NewSettingsFromCPointer(p unsafe.Pointer) *Settings {
	return &Settings{
		&_TraitSettings{(*C.GtkSettings)(p)},
		p,
	}
}

type SizeGroup struct {
	*_TraitSizeGroup
	CPointer unsafe.Pointer
}

func NewSizeGroupFromCPointer(p unsafe.Pointer) *SizeGroup {
	return &SizeGroup{
		&_TraitSizeGroup{(*C.GtkSizeGroup)(p)},
		p,
	}
}

type Socket struct {
	*_TraitSocket
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewSocketFromCPointer(p unsafe.Pointer) *Socket {
	return &Socket{
		&_TraitSocket{(*C.GtkSocket)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type SpinButton struct {
	*_TraitSpinButton
	*_TraitEntry
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewSpinButtonFromCPointer(p unsafe.Pointer) *SpinButton {
	return &SpinButton{
		&_TraitSpinButton{(*C.GtkSpinButton)(p)},
		&_TraitEntry{(*C.GtkEntry)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Spinner struct {
	*_TraitSpinner
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewSpinnerFromCPointer(p unsafe.Pointer) *Spinner {
	return &Spinner{
		&_TraitSpinner{(*C.GtkSpinner)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Stack struct {
	*_TraitStack
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewStackFromCPointer(p unsafe.Pointer) *Stack {
	return &Stack{
		&_TraitStack{(*C.GtkStack)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type StackSwitcher struct {
	*_TraitStackSwitcher
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewStackSwitcherFromCPointer(p unsafe.Pointer) *StackSwitcher {
	return &StackSwitcher{
		&_TraitStackSwitcher{(*C.GtkStackSwitcher)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type StatusIcon struct {
	*_TraitStatusIcon
	CPointer unsafe.Pointer
}

func NewStatusIconFromCPointer(p unsafe.Pointer) *StatusIcon {
	return &StatusIcon{
		&_TraitStatusIcon{(*C.GtkStatusIcon)(p)},
		p,
	}
}

type Statusbar struct {
	*_TraitStatusbar
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewStatusbarFromCPointer(p unsafe.Pointer) *Statusbar {
	return &Statusbar{
		&_TraitStatusbar{(*C.GtkStatusbar)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Style struct {
	*_TraitStyle
	CPointer unsafe.Pointer
}

func NewStyleFromCPointer(p unsafe.Pointer) *Style {
	return &Style{
		&_TraitStyle{(*C.GtkStyle)(p)},
		p,
	}
}

type StyleContext struct {
	*_TraitStyleContext
	CPointer unsafe.Pointer
}

func NewStyleContextFromCPointer(p unsafe.Pointer) *StyleContext {
	return &StyleContext{
		&_TraitStyleContext{(*C.GtkStyleContext)(p)},
		p,
	}
}

type StyleProperties struct {
	*_TraitStyleProperties
	CPointer unsafe.Pointer
}

func NewStylePropertiesFromCPointer(p unsafe.Pointer) *StyleProperties {
	return &StyleProperties{
		&_TraitStyleProperties{(*C.GtkStyleProperties)(p)},
		p,
	}
}

type Switch struct {
	*_TraitSwitch
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewSwitchFromCPointer(p unsafe.Pointer) *Switch {
	return &Switch{
		&_TraitSwitch{(*C.GtkSwitch)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Table struct {
	*_TraitTable
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewTableFromCPointer(p unsafe.Pointer) *Table {
	return &Table{
		&_TraitTable{(*C.GtkTable)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type TearoffMenuItem struct {
	*_TraitTearoffMenuItem
	*_TraitMenuItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewTearoffMenuItemFromCPointer(p unsafe.Pointer) *TearoffMenuItem {
	return &TearoffMenuItem{
		&_TraitTearoffMenuItem{(*C.GtkTearoffMenuItem)(p)},
		&_TraitMenuItem{(*C.GtkMenuItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type TextBuffer struct {
	*_TraitTextBuffer
	CPointer unsafe.Pointer
}

func NewTextBufferFromCPointer(p unsafe.Pointer) *TextBuffer {
	return &TextBuffer{
		&_TraitTextBuffer{(*C.GtkTextBuffer)(p)},
		p,
	}
}

type TextChildAnchor struct {
	*_TraitTextChildAnchor
	CPointer unsafe.Pointer
}

func NewTextChildAnchorFromCPointer(p unsafe.Pointer) *TextChildAnchor {
	return &TextChildAnchor{
		&_TraitTextChildAnchor{(*C.GtkTextChildAnchor)(p)},
		p,
	}
}

type TextMark struct {
	*_TraitTextMark
	CPointer unsafe.Pointer
}

func NewTextMarkFromCPointer(p unsafe.Pointer) *TextMark {
	return &TextMark{
		&_TraitTextMark{(*C.GtkTextMark)(p)},
		p,
	}
}

type TextTag struct {
	*_TraitTextTag
	CPointer unsafe.Pointer
}

func NewTextTagFromCPointer(p unsafe.Pointer) *TextTag {
	return &TextTag{
		&_TraitTextTag{(*C.GtkTextTag)(p)},
		p,
	}
}

type TextTagTable struct {
	*_TraitTextTagTable
	CPointer unsafe.Pointer
}

func NewTextTagTableFromCPointer(p unsafe.Pointer) *TextTagTable {
	return &TextTagTable{
		&_TraitTextTagTable{(*C.GtkTextTagTable)(p)},
		p,
	}
}

type TextView struct {
	*_TraitTextView
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewTextViewFromCPointer(p unsafe.Pointer) *TextView {
	return &TextView{
		&_TraitTextView{(*C.GtkTextView)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ThemingEngine struct {
	*_TraitThemingEngine
	CPointer unsafe.Pointer
}

func NewThemingEngineFromCPointer(p unsafe.Pointer) *ThemingEngine {
	return &ThemingEngine{
		&_TraitThemingEngine{(*C.GtkThemingEngine)(p)},
		p,
	}
}

type ToggleAction struct {
	*_TraitToggleAction
	*_TraitAction
	CPointer unsafe.Pointer
}

func NewToggleActionFromCPointer(p unsafe.Pointer) *ToggleAction {
	return &ToggleAction{
		&_TraitToggleAction{(*C.GtkToggleAction)(p)},
		&_TraitAction{(*C.GtkAction)(p)},
		p,
	}
}

type ToggleButton struct {
	*_TraitToggleButton
	*_TraitButton
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewToggleButtonFromCPointer(p unsafe.Pointer) *ToggleButton {
	return &ToggleButton{
		&_TraitToggleButton{(*C.GtkToggleButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewToggleToolButtonFromCPointer(p unsafe.Pointer) *ToggleToolButton {
	return &ToggleToolButton{
		&_TraitToggleToolButton{(*C.GtkToggleToolButton)(p)},
		&_TraitToolButton{(*C.GtkToolButton)(p)},
		&_TraitToolItem{(*C.GtkToolItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ToolButton struct {
	*_TraitToolButton
	*_TraitToolItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewToolButtonFromCPointer(p unsafe.Pointer) *ToolButton {
	return &ToolButton{
		&_TraitToolButton{(*C.GtkToolButton)(p)},
		&_TraitToolItem{(*C.GtkToolItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ToolItem struct {
	*_TraitToolItem
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewToolItemFromCPointer(p unsafe.Pointer) *ToolItem {
	return &ToolItem{
		&_TraitToolItem{(*C.GtkToolItem)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ToolItemGroup struct {
	*_TraitToolItemGroup
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewToolItemGroupFromCPointer(p unsafe.Pointer) *ToolItemGroup {
	return &ToolItemGroup{
		&_TraitToolItemGroup{(*C.GtkToolItemGroup)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type ToolPalette struct {
	*_TraitToolPalette
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewToolPaletteFromCPointer(p unsafe.Pointer) *ToolPalette {
	return &ToolPalette{
		&_TraitToolPalette{(*C.GtkToolPalette)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Toolbar struct {
	*_TraitToolbar
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewToolbarFromCPointer(p unsafe.Pointer) *Toolbar {
	return &Toolbar{
		&_TraitToolbar{(*C.GtkToolbar)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Tooltip struct {
	*_TraitTooltip
	CPointer unsafe.Pointer
}

func NewTooltipFromCPointer(p unsafe.Pointer) *Tooltip {
	return &Tooltip{
		&_TraitTooltip{(*C.GtkTooltip)(p)},
		p,
	}
}

type TreeModelFilter struct {
	*_TraitTreeModelFilter
	CPointer unsafe.Pointer
}

func NewTreeModelFilterFromCPointer(p unsafe.Pointer) *TreeModelFilter {
	return &TreeModelFilter{
		&_TraitTreeModelFilter{(*C.GtkTreeModelFilter)(p)},
		p,
	}
}

type TreeModelSort struct {
	*_TraitTreeModelSort
	CPointer unsafe.Pointer
}

func NewTreeModelSortFromCPointer(p unsafe.Pointer) *TreeModelSort {
	return &TreeModelSort{
		&_TraitTreeModelSort{(*C.GtkTreeModelSort)(p)},
		p,
	}
}

type TreeSelection struct {
	*_TraitTreeSelection
	CPointer unsafe.Pointer
}

func NewTreeSelectionFromCPointer(p unsafe.Pointer) *TreeSelection {
	return &TreeSelection{
		&_TraitTreeSelection{(*C.GtkTreeSelection)(p)},
		p,
	}
}

type TreeStore struct {
	*_TraitTreeStore
	CPointer unsafe.Pointer
}

func NewTreeStoreFromCPointer(p unsafe.Pointer) *TreeStore {
	return &TreeStore{
		&_TraitTreeStore{(*C.GtkTreeStore)(p)},
		p,
	}
}

type TreeView struct {
	*_TraitTreeView
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewTreeViewFromCPointer(p unsafe.Pointer) *TreeView {
	return &TreeView{
		&_TraitTreeView{(*C.GtkTreeView)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type TreeViewColumn struct {
	*_TraitTreeViewColumn
	CPointer unsafe.Pointer
}

func NewTreeViewColumnFromCPointer(p unsafe.Pointer) *TreeViewColumn {
	return &TreeViewColumn{
		&_TraitTreeViewColumn{(*C.GtkTreeViewColumn)(p)},
		p,
	}
}

type UIManager struct {
	*_TraitUIManager
	CPointer unsafe.Pointer
}

func NewUIManagerFromCPointer(p unsafe.Pointer) *UIManager {
	return &UIManager{
		&_TraitUIManager{(*C.GtkUIManager)(p)},
		p,
	}
}

type VBox struct {
	*_TraitVBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewVBoxFromCPointer(p unsafe.Pointer) *VBox {
	return &VBox{
		&_TraitVBox{(*C.GtkVBox)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type VButtonBox struct {
	*_TraitVButtonBox
	*_TraitButtonBox
	*_TraitBox
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewVButtonBoxFromCPointer(p unsafe.Pointer) *VButtonBox {
	return &VButtonBox{
		&_TraitVButtonBox{(*C.GtkVButtonBox)(p)},
		&_TraitButtonBox{(*C.GtkButtonBox)(p)},
		&_TraitBox{(*C.GtkBox)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type VPaned struct {
	*_TraitVPaned
	*_TraitPaned
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewVPanedFromCPointer(p unsafe.Pointer) *VPaned {
	return &VPaned{
		&_TraitVPaned{(*C.GtkVPaned)(p)},
		&_TraitPaned{(*C.GtkPaned)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type VScale struct {
	*_TraitVScale
	*_TraitScale
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewVScaleFromCPointer(p unsafe.Pointer) *VScale {
	return &VScale{
		&_TraitVScale{(*C.GtkVScale)(p)},
		&_TraitScale{(*C.GtkScale)(p)},
		&_TraitRange{(*C.GtkRange)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type VScrollbar struct {
	*_TraitVScrollbar
	*_TraitScrollbar
	*_TraitRange
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewVScrollbarFromCPointer(p unsafe.Pointer) *VScrollbar {
	return &VScrollbar{
		&_TraitVScrollbar{(*C.GtkVScrollbar)(p)},
		&_TraitScrollbar{(*C.GtkScrollbar)(p)},
		&_TraitRange{(*C.GtkRange)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type VSeparator struct {
	*_TraitVSeparator
	*_TraitSeparator
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewVSeparatorFromCPointer(p unsafe.Pointer) *VSeparator {
	return &VSeparator{
		&_TraitVSeparator{(*C.GtkVSeparator)(p)},
		&_TraitSeparator{(*C.GtkSeparator)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Viewport struct {
	*_TraitViewport
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewViewportFromCPointer(p unsafe.Pointer) *Viewport {
	return &Viewport{
		&_TraitViewport{(*C.GtkViewport)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
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

func NewVolumeButtonFromCPointer(p unsafe.Pointer) *VolumeButton {
	return &VolumeButton{
		&_TraitVolumeButton{(*C.GtkVolumeButton)(p)},
		&_TraitScaleButton{(*C.GtkScaleButton)(p)},
		&_TraitButton{(*C.GtkButton)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Widget struct {
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewWidgetFromCPointer(p unsafe.Pointer) *Widget {
	return &Widget{
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type Window struct {
	*_TraitWindow
	*_TraitBin
	*_TraitContainer
	*_TraitWidget
	CPointer unsafe.Pointer
}

func NewWindowFromCPointer(p unsafe.Pointer) *Window {
	return &Window{
		&_TraitWindow{(*C.GtkWindow)(p)},
		&_TraitBin{(*C.GtkBin)(p)},
		&_TraitContainer{(*C.GtkContainer)(p)},
		&_TraitWidget{(*C.GtkWidget)(p)},
		p,
	}
}

type WindowGroup struct {
	*_TraitWindowGroup
	CPointer unsafe.Pointer
}

func NewWindowGroupFromCPointer(p unsafe.Pointer) *WindowGroup {
	return &WindowGroup{
		&_TraitWindowGroup{(*C.GtkWindowGroup)(p)},
		p,
	}
}
