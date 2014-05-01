package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "../gobject"

func init() {
	_ = unsafe.Pointer(nil)
}

type AboutDialog struct {
	*TraitAboutDialog
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAboutDialogFromCPointer(p unsafe.Pointer) *AboutDialog {
	return &AboutDialog{
		NewTraitAboutDialog(p),
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type AccelGroup struct {
	*TraitAccelGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAccelGroupFromCPointer(p unsafe.Pointer) *AccelGroup {
	return &AccelGroup{
		NewTraitAccelGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type AccelLabel struct {
	*TraitAccelLabel
	*TraitLabel
	*TraitMisc
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAccelLabelFromCPointer(p unsafe.Pointer) *AccelLabel {
	return &AccelLabel{
		NewTraitAccelLabel(p),
		NewTraitLabel(p),
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type AccelMap struct {
	*TraitAccelMap
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAccelMapFromCPointer(p unsafe.Pointer) *AccelMap {
	return &AccelMap{
		NewTraitAccelMap(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Accessible struct {
	*TraitAccessible
	CPointer unsafe.Pointer
}

func NewAccessibleFromCPointer(p unsafe.Pointer) *Accessible {
	return &Accessible{
		NewTraitAccessible(p),
		p,
	}
}

type Action struct {
	*TraitAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewActionFromCPointer(p unsafe.Pointer) *Action {
	return &Action{
		NewTraitAction(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ActionBar struct {
	*TraitActionBar
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewActionBarFromCPointer(p unsafe.Pointer) *ActionBar {
	return &ActionBar{
		NewTraitActionBar(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ActionGroup struct {
	*TraitActionGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewActionGroupFromCPointer(p unsafe.Pointer) *ActionGroup {
	return &ActionGroup{
		NewTraitActionGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Adjustment struct {
	*TraitAdjustment
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAdjustmentFromCPointer(p unsafe.Pointer) *Adjustment {
	return &Adjustment{
		NewTraitAdjustment(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Alignment struct {
	*TraitAlignment
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAlignmentFromCPointer(p unsafe.Pointer) *Alignment {
	return &Alignment{
		NewTraitAlignment(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type AppChooserButton struct {
	*TraitAppChooserButton
	*TraitComboBox
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAppChooserButtonFromCPointer(p unsafe.Pointer) *AppChooserButton {
	return &AppChooserButton{
		NewTraitAppChooserButton(p),
		NewTraitComboBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type AppChooserDialog struct {
	*TraitAppChooserDialog
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAppChooserDialogFromCPointer(p unsafe.Pointer) *AppChooserDialog {
	return &AppChooserDialog{
		NewTraitAppChooserDialog(p),
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type AppChooserWidget struct {
	*TraitAppChooserWidget
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAppChooserWidgetFromCPointer(p unsafe.Pointer) *AppChooserWidget {
	return &AppChooserWidget{
		NewTraitAppChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Application struct {
	*TraitApplication
	CPointer unsafe.Pointer
}

func NewApplicationFromCPointer(p unsafe.Pointer) *Application {
	return &Application{
		NewTraitApplication(p),
		p,
	}
}

type ApplicationWindow struct {
	*TraitApplicationWindow
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewApplicationWindowFromCPointer(p unsafe.Pointer) *ApplicationWindow {
	return &ApplicationWindow{
		NewTraitApplicationWindow(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Arrow struct {
	*TraitArrow
	*TraitMisc
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewArrowFromCPointer(p unsafe.Pointer) *Arrow {
	return &Arrow{
		NewTraitArrow(p),
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type AspectFrame struct {
	*TraitAspectFrame
	*TraitFrame
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAspectFrameFromCPointer(p unsafe.Pointer) *AspectFrame {
	return &AspectFrame{
		NewTraitAspectFrame(p),
		NewTraitFrame(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Assistant struct {
	*TraitAssistant
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAssistantFromCPointer(p unsafe.Pointer) *Assistant {
	return &Assistant{
		NewTraitAssistant(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Bin struct {
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewBinFromCPointer(p unsafe.Pointer) *Bin {
	return &Bin{
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Box struct {
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewBoxFromCPointer(p unsafe.Pointer) *Box {
	return &Box{
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Builder struct {
	*TraitBuilder
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewBuilderFromCPointer(p unsafe.Pointer) *Builder {
	return &Builder{
		NewTraitBuilder(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Button struct {
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewButtonFromCPointer(p unsafe.Pointer) *Button {
	return &Button{
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ButtonBox struct {
	*TraitButtonBox
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewButtonBoxFromCPointer(p unsafe.Pointer) *ButtonBox {
	return &ButtonBox{
		NewTraitButtonBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Calendar struct {
	*TraitCalendar
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCalendarFromCPointer(p unsafe.Pointer) *Calendar {
	return &Calendar{
		NewTraitCalendar(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellArea struct {
	*TraitCellArea
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellAreaFromCPointer(p unsafe.Pointer) *CellArea {
	return &CellArea{
		NewTraitCellArea(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellAreaBox struct {
	*TraitCellAreaBox
	*TraitCellArea
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellAreaBoxFromCPointer(p unsafe.Pointer) *CellAreaBox {
	return &CellAreaBox{
		NewTraitCellAreaBox(p),
		NewTraitCellArea(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellAreaContext struct {
	*TraitCellAreaContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellAreaContextFromCPointer(p unsafe.Pointer) *CellAreaContext {
	return &CellAreaContext{
		NewTraitCellAreaContext(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellRenderer struct {
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererFromCPointer(p unsafe.Pointer) *CellRenderer {
	return &CellRenderer{
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellRendererAccel struct {
	*TraitCellRendererAccel
	*TraitCellRendererText
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererAccelFromCPointer(p unsafe.Pointer) *CellRendererAccel {
	return &CellRendererAccel{
		NewTraitCellRendererAccel(p),
		NewTraitCellRendererText(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellRendererCombo struct {
	*TraitCellRendererCombo
	*TraitCellRendererText
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererComboFromCPointer(p unsafe.Pointer) *CellRendererCombo {
	return &CellRendererCombo{
		NewTraitCellRendererCombo(p),
		NewTraitCellRendererText(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellRendererPixbuf struct {
	*TraitCellRendererPixbuf
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererPixbufFromCPointer(p unsafe.Pointer) *CellRendererPixbuf {
	return &CellRendererPixbuf{
		NewTraitCellRendererPixbuf(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellRendererProgress struct {
	*TraitCellRendererProgress
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererProgressFromCPointer(p unsafe.Pointer) *CellRendererProgress {
	return &CellRendererProgress{
		NewTraitCellRendererProgress(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellRendererSpin struct {
	*TraitCellRendererSpin
	*TraitCellRendererText
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererSpinFromCPointer(p unsafe.Pointer) *CellRendererSpin {
	return &CellRendererSpin{
		NewTraitCellRendererSpin(p),
		NewTraitCellRendererText(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellRendererSpinner struct {
	*TraitCellRendererSpinner
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererSpinnerFromCPointer(p unsafe.Pointer) *CellRendererSpinner {
	return &CellRendererSpinner{
		NewTraitCellRendererSpinner(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellRendererText struct {
	*TraitCellRendererText
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererTextFromCPointer(p unsafe.Pointer) *CellRendererText {
	return &CellRendererText{
		NewTraitCellRendererText(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellRendererToggle struct {
	*TraitCellRendererToggle
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererToggleFromCPointer(p unsafe.Pointer) *CellRendererToggle {
	return &CellRendererToggle{
		NewTraitCellRendererToggle(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CellView struct {
	*TraitCellView
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellViewFromCPointer(p unsafe.Pointer) *CellView {
	return &CellView{
		NewTraitCellView(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CheckButton struct {
	*TraitCheckButton
	*TraitToggleButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCheckButtonFromCPointer(p unsafe.Pointer) *CheckButton {
	return &CheckButton{
		NewTraitCheckButton(p),
		NewTraitToggleButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CheckMenuItem struct {
	*TraitCheckMenuItem
	*TraitMenuItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCheckMenuItemFromCPointer(p unsafe.Pointer) *CheckMenuItem {
	return &CheckMenuItem{
		NewTraitCheckMenuItem(p),
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Clipboard struct {
	*TraitClipboard
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewClipboardFromCPointer(p unsafe.Pointer) *Clipboard {
	return &Clipboard{
		NewTraitClipboard(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ColorButton struct {
	*TraitColorButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewColorButtonFromCPointer(p unsafe.Pointer) *ColorButton {
	return &ColorButton{
		NewTraitColorButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ColorChooserDialog struct {
	*TraitColorChooserDialog
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewColorChooserDialogFromCPointer(p unsafe.Pointer) *ColorChooserDialog {
	return &ColorChooserDialog{
		NewTraitColorChooserDialog(p),
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ColorChooserWidget struct {
	*TraitColorChooserWidget
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewColorChooserWidgetFromCPointer(p unsafe.Pointer) *ColorChooserWidget {
	return &ColorChooserWidget{
		NewTraitColorChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ColorSelection struct {
	*TraitColorSelection
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewColorSelectionFromCPointer(p unsafe.Pointer) *ColorSelection {
	return &ColorSelection{
		NewTraitColorSelection(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ColorSelectionDialog struct {
	*TraitColorSelectionDialog
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewColorSelectionDialogFromCPointer(p unsafe.Pointer) *ColorSelectionDialog {
	return &ColorSelectionDialog{
		NewTraitColorSelectionDialog(p),
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ComboBox struct {
	*TraitComboBox
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewComboBoxFromCPointer(p unsafe.Pointer) *ComboBox {
	return &ComboBox{
		NewTraitComboBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ComboBoxText struct {
	*TraitComboBoxText
	*TraitComboBox
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewComboBoxTextFromCPointer(p unsafe.Pointer) *ComboBoxText {
	return &ComboBoxText{
		NewTraitComboBoxText(p),
		NewTraitComboBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Container struct {
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewContainerFromCPointer(p unsafe.Pointer) *Container {
	return &Container{
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type CssProvider struct {
	*TraitCssProvider
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCssProviderFromCPointer(p unsafe.Pointer) *CssProvider {
	return &CssProvider{
		NewTraitCssProvider(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Dialog struct {
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDialogFromCPointer(p unsafe.Pointer) *Dialog {
	return &Dialog{
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type DrawingArea struct {
	*TraitDrawingArea
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDrawingAreaFromCPointer(p unsafe.Pointer) *DrawingArea {
	return &DrawingArea{
		NewTraitDrawingArea(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Entry struct {
	*TraitEntry
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewEntryFromCPointer(p unsafe.Pointer) *Entry {
	return &Entry{
		NewTraitEntry(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type EntryBuffer struct {
	*TraitEntryBuffer
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewEntryBufferFromCPointer(p unsafe.Pointer) *EntryBuffer {
	return &EntryBuffer{
		NewTraitEntryBuffer(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type EntryCompletion struct {
	*TraitEntryCompletion
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewEntryCompletionFromCPointer(p unsafe.Pointer) *EntryCompletion {
	return &EntryCompletion{
		NewTraitEntryCompletion(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type EventBox struct {
	*TraitEventBox
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewEventBoxFromCPointer(p unsafe.Pointer) *EventBox {
	return &EventBox{
		NewTraitEventBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Expander struct {
	*TraitExpander
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewExpanderFromCPointer(p unsafe.Pointer) *Expander {
	return &Expander{
		NewTraitExpander(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FileChooserButton struct {
	*TraitFileChooserButton
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileChooserButtonFromCPointer(p unsafe.Pointer) *FileChooserButton {
	return &FileChooserButton{
		NewTraitFileChooserButton(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FileChooserDialog struct {
	*TraitFileChooserDialog
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileChooserDialogFromCPointer(p unsafe.Pointer) *FileChooserDialog {
	return &FileChooserDialog{
		NewTraitFileChooserDialog(p),
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FileChooserWidget struct {
	*TraitFileChooserWidget
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileChooserWidgetFromCPointer(p unsafe.Pointer) *FileChooserWidget {
	return &FileChooserWidget{
		NewTraitFileChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FileFilter struct {
	*TraitFileFilter
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileFilterFromCPointer(p unsafe.Pointer) *FileFilter {
	return &FileFilter{
		NewTraitFileFilter(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Fixed struct {
	*TraitFixed
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFixedFromCPointer(p unsafe.Pointer) *Fixed {
	return &Fixed{
		NewTraitFixed(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FlowBox struct {
	*TraitFlowBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFlowBoxFromCPointer(p unsafe.Pointer) *FlowBox {
	return &FlowBox{
		NewTraitFlowBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FlowBoxChild struct {
	*TraitFlowBoxChild
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFlowBoxChildFromCPointer(p unsafe.Pointer) *FlowBoxChild {
	return &FlowBoxChild{
		NewTraitFlowBoxChild(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FontButton struct {
	*TraitFontButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFontButtonFromCPointer(p unsafe.Pointer) *FontButton {
	return &FontButton{
		NewTraitFontButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FontChooserDialog struct {
	*TraitFontChooserDialog
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFontChooserDialogFromCPointer(p unsafe.Pointer) *FontChooserDialog {
	return &FontChooserDialog{
		NewTraitFontChooserDialog(p),
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FontChooserWidget struct {
	*TraitFontChooserWidget
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFontChooserWidgetFromCPointer(p unsafe.Pointer) *FontChooserWidget {
	return &FontChooserWidget{
		NewTraitFontChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FontSelection struct {
	*TraitFontSelection
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFontSelectionFromCPointer(p unsafe.Pointer) *FontSelection {
	return &FontSelection{
		NewTraitFontSelection(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type FontSelectionDialog struct {
	*TraitFontSelectionDialog
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFontSelectionDialogFromCPointer(p unsafe.Pointer) *FontSelectionDialog {
	return &FontSelectionDialog{
		NewTraitFontSelectionDialog(p),
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Frame struct {
	*TraitFrame
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFrameFromCPointer(p unsafe.Pointer) *Frame {
	return &Frame{
		NewTraitFrame(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Grid struct {
	*TraitGrid
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewGridFromCPointer(p unsafe.Pointer) *Grid {
	return &Grid{
		NewTraitGrid(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type HBox struct {
	*TraitHBox
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHBoxFromCPointer(p unsafe.Pointer) *HBox {
	return &HBox{
		NewTraitHBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type HButtonBox struct {
	*TraitHButtonBox
	*TraitButtonBox
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHButtonBoxFromCPointer(p unsafe.Pointer) *HButtonBox {
	return &HButtonBox{
		NewTraitHButtonBox(p),
		NewTraitButtonBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type HPaned struct {
	*TraitHPaned
	*TraitPaned
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHPanedFromCPointer(p unsafe.Pointer) *HPaned {
	return &HPaned{
		NewTraitHPaned(p),
		NewTraitPaned(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type HSV struct {
	*TraitHSV
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHSVFromCPointer(p unsafe.Pointer) *HSV {
	return &HSV{
		NewTraitHSV(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type HScale struct {
	*TraitHScale
	*TraitScale
	*TraitRange
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHScaleFromCPointer(p unsafe.Pointer) *HScale {
	return &HScale{
		NewTraitHScale(p),
		NewTraitScale(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type HScrollbar struct {
	*TraitHScrollbar
	*TraitScrollbar
	*TraitRange
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHScrollbarFromCPointer(p unsafe.Pointer) *HScrollbar {
	return &HScrollbar{
		NewTraitHScrollbar(p),
		NewTraitScrollbar(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type HSeparator struct {
	*TraitHSeparator
	*TraitSeparator
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHSeparatorFromCPointer(p unsafe.Pointer) *HSeparator {
	return &HSeparator{
		NewTraitHSeparator(p),
		NewTraitSeparator(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type HandleBox struct {
	*TraitHandleBox
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHandleBoxFromCPointer(p unsafe.Pointer) *HandleBox {
	return &HandleBox{
		NewTraitHandleBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type HeaderBar struct {
	*TraitHeaderBar
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHeaderBarFromCPointer(p unsafe.Pointer) *HeaderBar {
	return &HeaderBar{
		NewTraitHeaderBar(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type IMContext struct {
	*TraitIMContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIMContextFromCPointer(p unsafe.Pointer) *IMContext {
	return &IMContext{
		NewTraitIMContext(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type IMContextSimple struct {
	*TraitIMContextSimple
	*TraitIMContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIMContextSimpleFromCPointer(p unsafe.Pointer) *IMContextSimple {
	return &IMContextSimple{
		NewTraitIMContextSimple(p),
		NewTraitIMContext(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type IMMulticontext struct {
	*TraitIMMulticontext
	*TraitIMContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIMMulticontextFromCPointer(p unsafe.Pointer) *IMMulticontext {
	return &IMMulticontext{
		NewTraitIMMulticontext(p),
		NewTraitIMContext(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type IconFactory struct {
	*TraitIconFactory
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIconFactoryFromCPointer(p unsafe.Pointer) *IconFactory {
	return &IconFactory{
		NewTraitIconFactory(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type IconInfo struct {
	*TraitIconInfo
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIconInfoFromCPointer(p unsafe.Pointer) *IconInfo {
	return &IconInfo{
		NewTraitIconInfo(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type IconTheme struct {
	*TraitIconTheme
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIconThemeFromCPointer(p unsafe.Pointer) *IconTheme {
	return &IconTheme{
		NewTraitIconTheme(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type IconView struct {
	*TraitIconView
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIconViewFromCPointer(p unsafe.Pointer) *IconView {
	return &IconView{
		NewTraitIconView(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Image struct {
	*TraitImage
	*TraitMisc
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewImageFromCPointer(p unsafe.Pointer) *Image {
	return &Image{
		NewTraitImage(p),
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ImageMenuItem struct {
	*TraitImageMenuItem
	*TraitMenuItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewImageMenuItemFromCPointer(p unsafe.Pointer) *ImageMenuItem {
	return &ImageMenuItem{
		NewTraitImageMenuItem(p),
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type InfoBar struct {
	*TraitInfoBar
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewInfoBarFromCPointer(p unsafe.Pointer) *InfoBar {
	return &InfoBar{
		NewTraitInfoBar(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Invisible struct {
	*TraitInvisible
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewInvisibleFromCPointer(p unsafe.Pointer) *Invisible {
	return &Invisible{
		NewTraitInvisible(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Label struct {
	*TraitLabel
	*TraitMisc
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewLabelFromCPointer(p unsafe.Pointer) *Label {
	return &Label{
		NewTraitLabel(p),
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Layout struct {
	*TraitLayout
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewLayoutFromCPointer(p unsafe.Pointer) *Layout {
	return &Layout{
		NewTraitLayout(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type LevelBar struct {
	*TraitLevelBar
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewLevelBarFromCPointer(p unsafe.Pointer) *LevelBar {
	return &LevelBar{
		NewTraitLevelBar(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type LinkButton struct {
	*TraitLinkButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewLinkButtonFromCPointer(p unsafe.Pointer) *LinkButton {
	return &LinkButton{
		NewTraitLinkButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ListBox struct {
	*TraitListBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewListBoxFromCPointer(p unsafe.Pointer) *ListBox {
	return &ListBox{
		NewTraitListBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ListBoxRow struct {
	*TraitListBoxRow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewListBoxRowFromCPointer(p unsafe.Pointer) *ListBoxRow {
	return &ListBoxRow{
		NewTraitListBoxRow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ListStore struct {
	*TraitListStore
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewListStoreFromCPointer(p unsafe.Pointer) *ListStore {
	return &ListStore{
		NewTraitListStore(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type LockButton struct {
	*TraitLockButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewLockButtonFromCPointer(p unsafe.Pointer) *LockButton {
	return &LockButton{
		NewTraitLockButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Menu struct {
	*TraitMenu
	*TraitMenuShell
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuFromCPointer(p unsafe.Pointer) *Menu {
	return &Menu{
		NewTraitMenu(p),
		NewTraitMenuShell(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type MenuBar struct {
	*TraitMenuBar
	*TraitMenuShell
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuBarFromCPointer(p unsafe.Pointer) *MenuBar {
	return &MenuBar{
		NewTraitMenuBar(p),
		NewTraitMenuShell(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type MenuButton struct {
	*TraitMenuButton
	*TraitToggleButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuButtonFromCPointer(p unsafe.Pointer) *MenuButton {
	return &MenuButton{
		NewTraitMenuButton(p),
		NewTraitToggleButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type MenuItem struct {
	*TraitMenuItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuItemFromCPointer(p unsafe.Pointer) *MenuItem {
	return &MenuItem{
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type MenuShell struct {
	*TraitMenuShell
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuShellFromCPointer(p unsafe.Pointer) *MenuShell {
	return &MenuShell{
		NewTraitMenuShell(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type MenuToolButton struct {
	*TraitMenuToolButton
	*TraitToolButton
	*TraitToolItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuToolButtonFromCPointer(p unsafe.Pointer) *MenuToolButton {
	return &MenuToolButton{
		NewTraitMenuToolButton(p),
		NewTraitToolButton(p),
		NewTraitToolItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type MessageDialog struct {
	*TraitMessageDialog
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMessageDialogFromCPointer(p unsafe.Pointer) *MessageDialog {
	return &MessageDialog{
		NewTraitMessageDialog(p),
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Misc struct {
	*TraitMisc
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMiscFromCPointer(p unsafe.Pointer) *Misc {
	return &Misc{
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type MountOperation struct {
	*TraitMountOperation
	CPointer unsafe.Pointer
}

func NewMountOperationFromCPointer(p unsafe.Pointer) *MountOperation {
	return &MountOperation{
		NewTraitMountOperation(p),
		p,
	}
}

type Notebook struct {
	*TraitNotebook
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewNotebookFromCPointer(p unsafe.Pointer) *Notebook {
	return &Notebook{
		NewTraitNotebook(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type NumerableIcon struct {
	*TraitNumerableIcon
	CPointer unsafe.Pointer
}

func NewNumerableIconFromCPointer(p unsafe.Pointer) *NumerableIcon {
	return &NumerableIcon{
		NewTraitNumerableIcon(p),
		p,
	}
}

type OffscreenWindow struct {
	*TraitOffscreenWindow
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewOffscreenWindowFromCPointer(p unsafe.Pointer) *OffscreenWindow {
	return &OffscreenWindow{
		NewTraitOffscreenWindow(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Overlay struct {
	*TraitOverlay
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewOverlayFromCPointer(p unsafe.Pointer) *Overlay {
	return &Overlay{
		NewTraitOverlay(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type PageSetup struct {
	*TraitPageSetup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPageSetupFromCPointer(p unsafe.Pointer) *PageSetup {
	return &PageSetup{
		NewTraitPageSetup(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Paned struct {
	*TraitPaned
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPanedFromCPointer(p unsafe.Pointer) *Paned {
	return &Paned{
		NewTraitPaned(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type PlacesSidebar struct {
	*TraitPlacesSidebar
	*TraitScrolledWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPlacesSidebarFromCPointer(p unsafe.Pointer) *PlacesSidebar {
	return &PlacesSidebar{
		NewTraitPlacesSidebar(p),
		NewTraitScrolledWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Plug struct {
	*TraitPlug
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPlugFromCPointer(p unsafe.Pointer) *Plug {
	return &Plug{
		NewTraitPlug(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Popover struct {
	*TraitPopover
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPopoverFromCPointer(p unsafe.Pointer) *Popover {
	return &Popover{
		NewTraitPopover(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type PrintContext struct {
	*TraitPrintContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPrintContextFromCPointer(p unsafe.Pointer) *PrintContext {
	return &PrintContext{
		NewTraitPrintContext(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type PrintOperation struct {
	*TraitPrintOperation
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPrintOperationFromCPointer(p unsafe.Pointer) *PrintOperation {
	return &PrintOperation{
		NewTraitPrintOperation(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type PrintSettings struct {
	*TraitPrintSettings
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPrintSettingsFromCPointer(p unsafe.Pointer) *PrintSettings {
	return &PrintSettings{
		NewTraitPrintSettings(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ProgressBar struct {
	*TraitProgressBar
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewProgressBarFromCPointer(p unsafe.Pointer) *ProgressBar {
	return &ProgressBar{
		NewTraitProgressBar(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RadioAction struct {
	*TraitRadioAction
	*TraitToggleAction
	*TraitAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRadioActionFromCPointer(p unsafe.Pointer) *RadioAction {
	return &RadioAction{
		NewTraitRadioAction(p),
		NewTraitToggleAction(p),
		NewTraitAction(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RadioButton struct {
	*TraitRadioButton
	*TraitCheckButton
	*TraitToggleButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRadioButtonFromCPointer(p unsafe.Pointer) *RadioButton {
	return &RadioButton{
		NewTraitRadioButton(p),
		NewTraitCheckButton(p),
		NewTraitToggleButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RadioMenuItem struct {
	*TraitRadioMenuItem
	*TraitCheckMenuItem
	*TraitMenuItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRadioMenuItemFromCPointer(p unsafe.Pointer) *RadioMenuItem {
	return &RadioMenuItem{
		NewTraitRadioMenuItem(p),
		NewTraitCheckMenuItem(p),
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RadioToolButton struct {
	*TraitRadioToolButton
	*TraitToggleToolButton
	*TraitToolButton
	*TraitToolItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRadioToolButtonFromCPointer(p unsafe.Pointer) *RadioToolButton {
	return &RadioToolButton{
		NewTraitRadioToolButton(p),
		NewTraitToggleToolButton(p),
		NewTraitToolButton(p),
		NewTraitToolItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Range struct {
	*TraitRange
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRangeFromCPointer(p unsafe.Pointer) *Range {
	return &Range{
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RcStyle struct {
	*TraitRcStyle
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRcStyleFromCPointer(p unsafe.Pointer) *RcStyle {
	return &RcStyle{
		NewTraitRcStyle(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RecentAction struct {
	*TraitRecentAction
	*TraitAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRecentActionFromCPointer(p unsafe.Pointer) *RecentAction {
	return &RecentAction{
		NewTraitRecentAction(p),
		NewTraitAction(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RecentChooserDialog struct {
	*TraitRecentChooserDialog
	*TraitDialog
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRecentChooserDialogFromCPointer(p unsafe.Pointer) *RecentChooserDialog {
	return &RecentChooserDialog{
		NewTraitRecentChooserDialog(p),
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RecentChooserMenu struct {
	*TraitRecentChooserMenu
	*TraitMenu
	*TraitMenuShell
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRecentChooserMenuFromCPointer(p unsafe.Pointer) *RecentChooserMenu {
	return &RecentChooserMenu{
		NewTraitRecentChooserMenu(p),
		NewTraitMenu(p),
		NewTraitMenuShell(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RecentChooserWidget struct {
	*TraitRecentChooserWidget
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRecentChooserWidgetFromCPointer(p unsafe.Pointer) *RecentChooserWidget {
	return &RecentChooserWidget{
		NewTraitRecentChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RecentFilter struct {
	*TraitRecentFilter
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRecentFilterFromCPointer(p unsafe.Pointer) *RecentFilter {
	return &RecentFilter{
		NewTraitRecentFilter(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type RecentManager struct {
	*TraitRecentManager
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRecentManagerFromCPointer(p unsafe.Pointer) *RecentManager {
	return &RecentManager{
		NewTraitRecentManager(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Revealer struct {
	*TraitRevealer
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRevealerFromCPointer(p unsafe.Pointer) *Revealer {
	return &Revealer{
		NewTraitRevealer(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Scale struct {
	*TraitScale
	*TraitRange
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewScaleFromCPointer(p unsafe.Pointer) *Scale {
	return &Scale{
		NewTraitScale(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ScaleButton struct {
	*TraitScaleButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewScaleButtonFromCPointer(p unsafe.Pointer) *ScaleButton {
	return &ScaleButton{
		NewTraitScaleButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Scrollbar struct {
	*TraitScrollbar
	*TraitRange
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewScrollbarFromCPointer(p unsafe.Pointer) *Scrollbar {
	return &Scrollbar{
		NewTraitScrollbar(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ScrolledWindow struct {
	*TraitScrolledWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewScrolledWindowFromCPointer(p unsafe.Pointer) *ScrolledWindow {
	return &ScrolledWindow{
		NewTraitScrolledWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type SearchBar struct {
	*TraitSearchBar
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSearchBarFromCPointer(p unsafe.Pointer) *SearchBar {
	return &SearchBar{
		NewTraitSearchBar(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type SearchEntry struct {
	*TraitSearchEntry
	*TraitEntry
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSearchEntryFromCPointer(p unsafe.Pointer) *SearchEntry {
	return &SearchEntry{
		NewTraitSearchEntry(p),
		NewTraitEntry(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Separator struct {
	*TraitSeparator
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSeparatorFromCPointer(p unsafe.Pointer) *Separator {
	return &Separator{
		NewTraitSeparator(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type SeparatorMenuItem struct {
	*TraitSeparatorMenuItem
	*TraitMenuItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSeparatorMenuItemFromCPointer(p unsafe.Pointer) *SeparatorMenuItem {
	return &SeparatorMenuItem{
		NewTraitSeparatorMenuItem(p),
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type SeparatorToolItem struct {
	*TraitSeparatorToolItem
	*TraitToolItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSeparatorToolItemFromCPointer(p unsafe.Pointer) *SeparatorToolItem {
	return &SeparatorToolItem{
		NewTraitSeparatorToolItem(p),
		NewTraitToolItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Settings struct {
	*TraitSettings
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSettingsFromCPointer(p unsafe.Pointer) *Settings {
	return &Settings{
		NewTraitSettings(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type SizeGroup struct {
	*TraitSizeGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSizeGroupFromCPointer(p unsafe.Pointer) *SizeGroup {
	return &SizeGroup{
		NewTraitSizeGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Socket struct {
	*TraitSocket
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSocketFromCPointer(p unsafe.Pointer) *Socket {
	return &Socket{
		NewTraitSocket(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type SpinButton struct {
	*TraitSpinButton
	*TraitEntry
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSpinButtonFromCPointer(p unsafe.Pointer) *SpinButton {
	return &SpinButton{
		NewTraitSpinButton(p),
		NewTraitEntry(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Spinner struct {
	*TraitSpinner
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSpinnerFromCPointer(p unsafe.Pointer) *Spinner {
	return &Spinner{
		NewTraitSpinner(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Stack struct {
	*TraitStack
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStackFromCPointer(p unsafe.Pointer) *Stack {
	return &Stack{
		NewTraitStack(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type StackSwitcher struct {
	*TraitStackSwitcher
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStackSwitcherFromCPointer(p unsafe.Pointer) *StackSwitcher {
	return &StackSwitcher{
		NewTraitStackSwitcher(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type StatusIcon struct {
	*TraitStatusIcon
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStatusIconFromCPointer(p unsafe.Pointer) *StatusIcon {
	return &StatusIcon{
		NewTraitStatusIcon(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Statusbar struct {
	*TraitStatusbar
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStatusbarFromCPointer(p unsafe.Pointer) *Statusbar {
	return &Statusbar{
		NewTraitStatusbar(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Style struct {
	*TraitStyle
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStyleFromCPointer(p unsafe.Pointer) *Style {
	return &Style{
		NewTraitStyle(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type StyleContext struct {
	*TraitStyleContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStyleContextFromCPointer(p unsafe.Pointer) *StyleContext {
	return &StyleContext{
		NewTraitStyleContext(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type StyleProperties struct {
	*TraitStyleProperties
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStylePropertiesFromCPointer(p unsafe.Pointer) *StyleProperties {
	return &StyleProperties{
		NewTraitStyleProperties(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Switch struct {
	*TraitSwitch
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSwitchFromCPointer(p unsafe.Pointer) *Switch {
	return &Switch{
		NewTraitSwitch(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Table struct {
	*TraitTable
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTableFromCPointer(p unsafe.Pointer) *Table {
	return &Table{
		NewTraitTable(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TearoffMenuItem struct {
	*TraitTearoffMenuItem
	*TraitMenuItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTearoffMenuItemFromCPointer(p unsafe.Pointer) *TearoffMenuItem {
	return &TearoffMenuItem{
		NewTraitTearoffMenuItem(p),
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TextBuffer struct {
	*TraitTextBuffer
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextBufferFromCPointer(p unsafe.Pointer) *TextBuffer {
	return &TextBuffer{
		NewTraitTextBuffer(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TextChildAnchor struct {
	*TraitTextChildAnchor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextChildAnchorFromCPointer(p unsafe.Pointer) *TextChildAnchor {
	return &TextChildAnchor{
		NewTraitTextChildAnchor(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TextMark struct {
	*TraitTextMark
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextMarkFromCPointer(p unsafe.Pointer) *TextMark {
	return &TextMark{
		NewTraitTextMark(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TextTag struct {
	*TraitTextTag
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextTagFromCPointer(p unsafe.Pointer) *TextTag {
	return &TextTag{
		NewTraitTextTag(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TextTagTable struct {
	*TraitTextTagTable
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextTagTableFromCPointer(p unsafe.Pointer) *TextTagTable {
	return &TextTagTable{
		NewTraitTextTagTable(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TextView struct {
	*TraitTextView
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextViewFromCPointer(p unsafe.Pointer) *TextView {
	return &TextView{
		NewTraitTextView(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ThemingEngine struct {
	*TraitThemingEngine
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewThemingEngineFromCPointer(p unsafe.Pointer) *ThemingEngine {
	return &ThemingEngine{
		NewTraitThemingEngine(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ToggleAction struct {
	*TraitToggleAction
	*TraitAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewToggleActionFromCPointer(p unsafe.Pointer) *ToggleAction {
	return &ToggleAction{
		NewTraitToggleAction(p),
		NewTraitAction(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ToggleButton struct {
	*TraitToggleButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewToggleButtonFromCPointer(p unsafe.Pointer) *ToggleButton {
	return &ToggleButton{
		NewTraitToggleButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ToggleToolButton struct {
	*TraitToggleToolButton
	*TraitToolButton
	*TraitToolItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewToggleToolButtonFromCPointer(p unsafe.Pointer) *ToggleToolButton {
	return &ToggleToolButton{
		NewTraitToggleToolButton(p),
		NewTraitToolButton(p),
		NewTraitToolItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ToolButton struct {
	*TraitToolButton
	*TraitToolItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewToolButtonFromCPointer(p unsafe.Pointer) *ToolButton {
	return &ToolButton{
		NewTraitToolButton(p),
		NewTraitToolItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ToolItem struct {
	*TraitToolItem
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewToolItemFromCPointer(p unsafe.Pointer) *ToolItem {
	return &ToolItem{
		NewTraitToolItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ToolItemGroup struct {
	*TraitToolItemGroup
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewToolItemGroupFromCPointer(p unsafe.Pointer) *ToolItemGroup {
	return &ToolItemGroup{
		NewTraitToolItemGroup(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type ToolPalette struct {
	*TraitToolPalette
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewToolPaletteFromCPointer(p unsafe.Pointer) *ToolPalette {
	return &ToolPalette{
		NewTraitToolPalette(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Toolbar struct {
	*TraitToolbar
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewToolbarFromCPointer(p unsafe.Pointer) *Toolbar {
	return &Toolbar{
		NewTraitToolbar(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Tooltip struct {
	*TraitTooltip
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTooltipFromCPointer(p unsafe.Pointer) *Tooltip {
	return &Tooltip{
		NewTraitTooltip(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TreeModelFilter struct {
	*TraitTreeModelFilter
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeModelFilterFromCPointer(p unsafe.Pointer) *TreeModelFilter {
	return &TreeModelFilter{
		NewTraitTreeModelFilter(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TreeModelSort struct {
	*TraitTreeModelSort
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeModelSortFromCPointer(p unsafe.Pointer) *TreeModelSort {
	return &TreeModelSort{
		NewTraitTreeModelSort(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TreeSelection struct {
	*TraitTreeSelection
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeSelectionFromCPointer(p unsafe.Pointer) *TreeSelection {
	return &TreeSelection{
		NewTraitTreeSelection(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TreeStore struct {
	*TraitTreeStore
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeStoreFromCPointer(p unsafe.Pointer) *TreeStore {
	return &TreeStore{
		NewTraitTreeStore(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TreeView struct {
	*TraitTreeView
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeViewFromCPointer(p unsafe.Pointer) *TreeView {
	return &TreeView{
		NewTraitTreeView(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type TreeViewColumn struct {
	*TraitTreeViewColumn
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeViewColumnFromCPointer(p unsafe.Pointer) *TreeViewColumn {
	return &TreeViewColumn{
		NewTraitTreeViewColumn(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type UIManager struct {
	*TraitUIManager
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUIManagerFromCPointer(p unsafe.Pointer) *UIManager {
	return &UIManager{
		NewTraitUIManager(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type VBox struct {
	*TraitVBox
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVBoxFromCPointer(p unsafe.Pointer) *VBox {
	return &VBox{
		NewTraitVBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type VButtonBox struct {
	*TraitVButtonBox
	*TraitButtonBox
	*TraitBox
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVButtonBoxFromCPointer(p unsafe.Pointer) *VButtonBox {
	return &VButtonBox{
		NewTraitVButtonBox(p),
		NewTraitButtonBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type VPaned struct {
	*TraitVPaned
	*TraitPaned
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVPanedFromCPointer(p unsafe.Pointer) *VPaned {
	return &VPaned{
		NewTraitVPaned(p),
		NewTraitPaned(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type VScale struct {
	*TraitVScale
	*TraitScale
	*TraitRange
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVScaleFromCPointer(p unsafe.Pointer) *VScale {
	return &VScale{
		NewTraitVScale(p),
		NewTraitScale(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type VScrollbar struct {
	*TraitVScrollbar
	*TraitScrollbar
	*TraitRange
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVScrollbarFromCPointer(p unsafe.Pointer) *VScrollbar {
	return &VScrollbar{
		NewTraitVScrollbar(p),
		NewTraitScrollbar(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type VSeparator struct {
	*TraitVSeparator
	*TraitSeparator
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVSeparatorFromCPointer(p unsafe.Pointer) *VSeparator {
	return &VSeparator{
		NewTraitVSeparator(p),
		NewTraitSeparator(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Viewport struct {
	*TraitViewport
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewViewportFromCPointer(p unsafe.Pointer) *Viewport {
	return &Viewport{
		NewTraitViewport(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type VolumeButton struct {
	*TraitVolumeButton
	*TraitScaleButton
	*TraitButton
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVolumeButtonFromCPointer(p unsafe.Pointer) *VolumeButton {
	return &VolumeButton{
		NewTraitVolumeButton(p),
		NewTraitScaleButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Widget struct {
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWidgetFromCPointer(p unsafe.Pointer) *Widget {
	return &Widget{
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type Window struct {
	*TraitWindow
	*TraitBin
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWindowFromCPointer(p unsafe.Pointer) *Window {
	return &Window{
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
}

type WindowGroup struct {
	*TraitWindowGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWindowGroupFromCPointer(p unsafe.Pointer) *WindowGroup {
	return &WindowGroup{
		NewTraitWindowGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
}
