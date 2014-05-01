package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>
#include <glib-object.h>
#cgo pkg-config: gobject-2.0
*/
import "C"
import "unsafe"
import "../gobject"
import "runtime"

func init() {
	_ = unsafe.Pointer(nil)
	_ = runtime.Compiler
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
	ret := &AboutDialog{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AboutDialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type AccelGroup struct {
	*TraitAccelGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAccelGroupFromCPointer(p unsafe.Pointer) *AccelGroup {
	ret := &AccelGroup{
		NewTraitAccelGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AccelGroup) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &AccelLabel{
		NewTraitAccelLabel(p),
		NewTraitLabel(p),
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AccelLabel) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type AccelMap struct {
	*TraitAccelMap
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAccelMapFromCPointer(p unsafe.Pointer) *AccelMap {
	ret := &AccelMap{
		NewTraitAccelMap(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AccelMap) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Accessible struct {
	*TraitAccessible
	CPointer unsafe.Pointer
}

func NewAccessibleFromCPointer(p unsafe.Pointer) *Accessible {
	ret := &Accessible{
		NewTraitAccessible(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Accessible) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Action struct {
	*TraitAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewActionFromCPointer(p unsafe.Pointer) *Action {
	ret := &Action{
		NewTraitAction(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Action) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ActionBar{
		NewTraitActionBar(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ActionBar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ActionGroup struct {
	*TraitActionGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewActionGroupFromCPointer(p unsafe.Pointer) *ActionGroup {
	ret := &ActionGroup{
		NewTraitActionGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ActionGroup) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Adjustment struct {
	*TraitAdjustment
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAdjustmentFromCPointer(p unsafe.Pointer) *Adjustment {
	ret := &Adjustment{
		NewTraitAdjustment(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Adjustment) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Alignment{
		NewTraitAlignment(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Alignment) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &AppChooserButton{
		NewTraitAppChooserButton(p),
		NewTraitComboBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AppChooserButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &AppChooserDialog{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AppChooserDialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &AppChooserWidget{
		NewTraitAppChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AppChooserWidget) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Application struct {
	*TraitApplication
	CPointer unsafe.Pointer
}

func NewApplicationFromCPointer(p unsafe.Pointer) *Application {
	ret := &Application{
		NewTraitApplication(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Application) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ApplicationWindow{
		NewTraitApplicationWindow(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ApplicationWindow) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Arrow{
		NewTraitArrow(p),
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Arrow) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &AspectFrame{
		NewTraitAspectFrame(p),
		NewTraitFrame(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AspectFrame) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Assistant{
		NewTraitAssistant(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Assistant) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Bin{
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Bin) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Box{
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Box) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Builder struct {
	*TraitBuilder
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewBuilderFromCPointer(p unsafe.Pointer) *Builder {
	ret := &Builder{
		NewTraitBuilder(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Builder) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Button{
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Button) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ButtonBox{
		NewTraitButtonBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ButtonBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Calendar struct {
	*TraitCalendar
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCalendarFromCPointer(p unsafe.Pointer) *Calendar {
	ret := &Calendar{
		NewTraitCalendar(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Calendar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellArea struct {
	*TraitCellArea
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellAreaFromCPointer(p unsafe.Pointer) *CellArea {
	ret := &CellArea{
		NewTraitCellArea(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellArea) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellAreaBox struct {
	*TraitCellAreaBox
	*TraitCellArea
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellAreaBoxFromCPointer(p unsafe.Pointer) *CellAreaBox {
	ret := &CellAreaBox{
		NewTraitCellAreaBox(p),
		NewTraitCellArea(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellAreaBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellAreaContext struct {
	*TraitCellAreaContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellAreaContextFromCPointer(p unsafe.Pointer) *CellAreaContext {
	ret := &CellAreaContext{
		NewTraitCellAreaContext(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellAreaContext) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellRenderer struct {
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererFromCPointer(p unsafe.Pointer) *CellRenderer {
	ret := &CellRenderer{
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellRenderer) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &CellRendererAccel{
		NewTraitCellRendererAccel(p),
		NewTraitCellRendererText(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellRendererAccel) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &CellRendererCombo{
		NewTraitCellRendererCombo(p),
		NewTraitCellRendererText(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellRendererCombo) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellRendererPixbuf struct {
	*TraitCellRendererPixbuf
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererPixbufFromCPointer(p unsafe.Pointer) *CellRendererPixbuf {
	ret := &CellRendererPixbuf{
		NewTraitCellRendererPixbuf(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellRendererPixbuf) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellRendererProgress struct {
	*TraitCellRendererProgress
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererProgressFromCPointer(p unsafe.Pointer) *CellRendererProgress {
	ret := &CellRendererProgress{
		NewTraitCellRendererProgress(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellRendererProgress) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &CellRendererSpin{
		NewTraitCellRendererSpin(p),
		NewTraitCellRendererText(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellRendererSpin) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellRendererSpinner struct {
	*TraitCellRendererSpinner
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererSpinnerFromCPointer(p unsafe.Pointer) *CellRendererSpinner {
	ret := &CellRendererSpinner{
		NewTraitCellRendererSpinner(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellRendererSpinner) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellRendererText struct {
	*TraitCellRendererText
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererTextFromCPointer(p unsafe.Pointer) *CellRendererText {
	ret := &CellRendererText{
		NewTraitCellRendererText(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellRendererText) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellRendererToggle struct {
	*TraitCellRendererToggle
	*TraitCellRenderer
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellRendererToggleFromCPointer(p unsafe.Pointer) *CellRendererToggle {
	ret := &CellRendererToggle{
		NewTraitCellRendererToggle(p),
		NewTraitCellRenderer(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellRendererToggle) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CellView struct {
	*TraitCellView
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCellViewFromCPointer(p unsafe.Pointer) *CellView {
	ret := &CellView{
		NewTraitCellView(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CellView) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &CheckButton{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CheckButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &CheckMenuItem{
		NewTraitCheckMenuItem(p),
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CheckMenuItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Clipboard struct {
	*TraitClipboard
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewClipboardFromCPointer(p unsafe.Pointer) *Clipboard {
	ret := &Clipboard{
		NewTraitClipboard(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Clipboard) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ColorButton{
		NewTraitColorButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ColorButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ColorChooserDialog{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ColorChooserDialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ColorChooserWidget{
		NewTraitColorChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ColorChooserWidget) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ColorSelection{
		NewTraitColorSelection(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ColorSelection) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ColorSelectionDialog{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ColorSelectionDialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ComboBox{
		NewTraitComboBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ComboBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ComboBoxText{
		NewTraitComboBoxText(p),
		NewTraitComboBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ComboBoxText) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Container struct {
	*TraitContainer
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewContainerFromCPointer(p unsafe.Pointer) *Container {
	ret := &Container{
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Container) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CssProvider struct {
	*TraitCssProvider
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCssProviderFromCPointer(p unsafe.Pointer) *CssProvider {
	ret := &CssProvider{
		NewTraitCssProvider(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CssProvider) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Dialog{
		NewTraitDialog(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Dialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DrawingArea struct {
	*TraitDrawingArea
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDrawingAreaFromCPointer(p unsafe.Pointer) *DrawingArea {
	ret := &DrawingArea{
		NewTraitDrawingArea(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DrawingArea) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Entry struct {
	*TraitEntry
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewEntryFromCPointer(p unsafe.Pointer) *Entry {
	ret := &Entry{
		NewTraitEntry(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Entry) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type EntryBuffer struct {
	*TraitEntryBuffer
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewEntryBufferFromCPointer(p unsafe.Pointer) *EntryBuffer {
	ret := &EntryBuffer{
		NewTraitEntryBuffer(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *EntryBuffer) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type EntryCompletion struct {
	*TraitEntryCompletion
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewEntryCompletionFromCPointer(p unsafe.Pointer) *EntryCompletion {
	ret := &EntryCompletion{
		NewTraitEntryCompletion(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *EntryCompletion) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &EventBox{
		NewTraitEventBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *EventBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Expander{
		NewTraitExpander(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Expander) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FileChooserButton{
		NewTraitFileChooserButton(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileChooserButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FileChooserDialog{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileChooserDialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FileChooserWidget{
		NewTraitFileChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileChooserWidget) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FileFilter struct {
	*TraitFileFilter
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileFilterFromCPointer(p unsafe.Pointer) *FileFilter {
	ret := &FileFilter{
		NewTraitFileFilter(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileFilter) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Fixed{
		NewTraitFixed(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Fixed) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FlowBox{
		NewTraitFlowBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FlowBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FlowBoxChild{
		NewTraitFlowBoxChild(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FlowBoxChild) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FontButton{
		NewTraitFontButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FontButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FontChooserDialog{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FontChooserDialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FontChooserWidget{
		NewTraitFontChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FontChooserWidget) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FontSelection{
		NewTraitFontSelection(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FontSelection) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &FontSelectionDialog{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FontSelectionDialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Frame{
		NewTraitFrame(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Frame) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Grid{
		NewTraitGrid(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Grid) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &HBox{
		NewTraitHBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &HButtonBox{
		NewTraitHButtonBox(p),
		NewTraitButtonBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HButtonBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &HPaned{
		NewTraitHPaned(p),
		NewTraitPaned(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HPaned) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type HSV struct {
	*TraitHSV
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHSVFromCPointer(p unsafe.Pointer) *HSV {
	ret := &HSV{
		NewTraitHSV(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HSV) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &HScale{
		NewTraitHScale(p),
		NewTraitScale(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HScale) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &HScrollbar{
		NewTraitHScrollbar(p),
		NewTraitScrollbar(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HScrollbar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &HSeparator{
		NewTraitHSeparator(p),
		NewTraitSeparator(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HSeparator) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &HandleBox{
		NewTraitHandleBox(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HandleBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &HeaderBar{
		NewTraitHeaderBar(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HeaderBar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type IMContext struct {
	*TraitIMContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIMContextFromCPointer(p unsafe.Pointer) *IMContext {
	ret := &IMContext{
		NewTraitIMContext(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *IMContext) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type IMContextSimple struct {
	*TraitIMContextSimple
	*TraitIMContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIMContextSimpleFromCPointer(p unsafe.Pointer) *IMContextSimple {
	ret := &IMContextSimple{
		NewTraitIMContextSimple(p),
		NewTraitIMContext(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *IMContextSimple) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type IMMulticontext struct {
	*TraitIMMulticontext
	*TraitIMContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIMMulticontextFromCPointer(p unsafe.Pointer) *IMMulticontext {
	ret := &IMMulticontext{
		NewTraitIMMulticontext(p),
		NewTraitIMContext(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *IMMulticontext) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type IconFactory struct {
	*TraitIconFactory
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIconFactoryFromCPointer(p unsafe.Pointer) *IconFactory {
	ret := &IconFactory{
		NewTraitIconFactory(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *IconFactory) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type IconInfo struct {
	*TraitIconInfo
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIconInfoFromCPointer(p unsafe.Pointer) *IconInfo {
	ret := &IconInfo{
		NewTraitIconInfo(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *IconInfo) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type IconTheme struct {
	*TraitIconTheme
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIconThemeFromCPointer(p unsafe.Pointer) *IconTheme {
	ret := &IconTheme{
		NewTraitIconTheme(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *IconTheme) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &IconView{
		NewTraitIconView(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *IconView) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Image{
		NewTraitImage(p),
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Image) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ImageMenuItem{
		NewTraitImageMenuItem(p),
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ImageMenuItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &InfoBar{
		NewTraitInfoBar(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *InfoBar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Invisible struct {
	*TraitInvisible
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewInvisibleFromCPointer(p unsafe.Pointer) *Invisible {
	ret := &Invisible{
		NewTraitInvisible(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Invisible) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Label{
		NewTraitLabel(p),
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Label) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Layout{
		NewTraitLayout(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Layout) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type LevelBar struct {
	*TraitLevelBar
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewLevelBarFromCPointer(p unsafe.Pointer) *LevelBar {
	ret := &LevelBar{
		NewTraitLevelBar(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *LevelBar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &LinkButton{
		NewTraitLinkButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *LinkButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ListBox{
		NewTraitListBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ListBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ListBoxRow{
		NewTraitListBoxRow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ListBoxRow) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ListStore struct {
	*TraitListStore
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewListStoreFromCPointer(p unsafe.Pointer) *ListStore {
	ret := &ListStore{
		NewTraitListStore(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ListStore) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &LockButton{
		NewTraitLockButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *LockButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Menu{
		NewTraitMenu(p),
		NewTraitMenuShell(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Menu) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &MenuBar{
		NewTraitMenuBar(p),
		NewTraitMenuShell(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MenuBar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &MenuButton{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MenuButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &MenuItem{
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MenuItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &MenuShell{
		NewTraitMenuShell(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MenuShell) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &MenuToolButton{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MenuToolButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &MessageDialog{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MessageDialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Misc struct {
	*TraitMisc
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMiscFromCPointer(p unsafe.Pointer) *Misc {
	ret := &Misc{
		NewTraitMisc(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Misc) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type MountOperation struct {
	*TraitMountOperation
	CPointer unsafe.Pointer
}

func NewMountOperationFromCPointer(p unsafe.Pointer) *MountOperation {
	ret := &MountOperation{
		NewTraitMountOperation(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MountOperation) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Notebook{
		NewTraitNotebook(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Notebook) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type NumerableIcon struct {
	*TraitNumerableIcon
	CPointer unsafe.Pointer
}

func NewNumerableIconFromCPointer(p unsafe.Pointer) *NumerableIcon {
	ret := &NumerableIcon{
		NewTraitNumerableIcon(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *NumerableIcon) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &OffscreenWindow{
		NewTraitOffscreenWindow(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *OffscreenWindow) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Overlay{
		NewTraitOverlay(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Overlay) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PageSetup struct {
	*TraitPageSetup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPageSetupFromCPointer(p unsafe.Pointer) *PageSetup {
	ret := &PageSetup{
		NewTraitPageSetup(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PageSetup) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Paned{
		NewTraitPaned(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Paned) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &PlacesSidebar{
		NewTraitPlacesSidebar(p),
		NewTraitScrolledWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PlacesSidebar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Plug{
		NewTraitPlug(p),
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Plug) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Popover{
		NewTraitPopover(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Popover) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PrintContext struct {
	*TraitPrintContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPrintContextFromCPointer(p unsafe.Pointer) *PrintContext {
	ret := &PrintContext{
		NewTraitPrintContext(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PrintContext) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PrintOperation struct {
	*TraitPrintOperation
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPrintOperationFromCPointer(p unsafe.Pointer) *PrintOperation {
	ret := &PrintOperation{
		NewTraitPrintOperation(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PrintOperation) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PrintSettings struct {
	*TraitPrintSettings
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPrintSettingsFromCPointer(p unsafe.Pointer) *PrintSettings {
	ret := &PrintSettings{
		NewTraitPrintSettings(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PrintSettings) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ProgressBar struct {
	*TraitProgressBar
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewProgressBarFromCPointer(p unsafe.Pointer) *ProgressBar {
	ret := &ProgressBar{
		NewTraitProgressBar(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ProgressBar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type RadioAction struct {
	*TraitRadioAction
	*TraitToggleAction
	*TraitAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRadioActionFromCPointer(p unsafe.Pointer) *RadioAction {
	ret := &RadioAction{
		NewTraitRadioAction(p),
		NewTraitToggleAction(p),
		NewTraitAction(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RadioAction) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &RadioButton{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RadioButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &RadioMenuItem{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RadioMenuItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &RadioToolButton{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RadioToolButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Range struct {
	*TraitRange
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRangeFromCPointer(p unsafe.Pointer) *Range {
	ret := &Range{
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Range) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type RcStyle struct {
	*TraitRcStyle
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRcStyleFromCPointer(p unsafe.Pointer) *RcStyle {
	ret := &RcStyle{
		NewTraitRcStyle(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RcStyle) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type RecentAction struct {
	*TraitRecentAction
	*TraitAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRecentActionFromCPointer(p unsafe.Pointer) *RecentAction {
	ret := &RecentAction{
		NewTraitRecentAction(p),
		NewTraitAction(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RecentAction) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &RecentChooserDialog{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RecentChooserDialog) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &RecentChooserMenu{
		NewTraitRecentChooserMenu(p),
		NewTraitMenu(p),
		NewTraitMenuShell(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RecentChooserMenu) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &RecentChooserWidget{
		NewTraitRecentChooserWidget(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RecentChooserWidget) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type RecentFilter struct {
	*TraitRecentFilter
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRecentFilterFromCPointer(p unsafe.Pointer) *RecentFilter {
	ret := &RecentFilter{
		NewTraitRecentFilter(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RecentFilter) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type RecentManager struct {
	*TraitRecentManager
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewRecentManagerFromCPointer(p unsafe.Pointer) *RecentManager {
	ret := &RecentManager{
		NewTraitRecentManager(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *RecentManager) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Revealer{
		NewTraitRevealer(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Revealer) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Scale{
		NewTraitScale(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Scale) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ScaleButton{
		NewTraitScaleButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ScaleButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Scrollbar{
		NewTraitScrollbar(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Scrollbar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ScrolledWindow{
		NewTraitScrolledWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ScrolledWindow) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &SearchBar{
		NewTraitSearchBar(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SearchBar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &SearchEntry{
		NewTraitSearchEntry(p),
		NewTraitEntry(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SearchEntry) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Separator struct {
	*TraitSeparator
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSeparatorFromCPointer(p unsafe.Pointer) *Separator {
	ret := &Separator{
		NewTraitSeparator(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Separator) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &SeparatorMenuItem{
		NewTraitSeparatorMenuItem(p),
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SeparatorMenuItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &SeparatorToolItem{
		NewTraitSeparatorToolItem(p),
		NewTraitToolItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SeparatorToolItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Settings struct {
	*TraitSettings
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSettingsFromCPointer(p unsafe.Pointer) *Settings {
	ret := &Settings{
		NewTraitSettings(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Settings) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SizeGroup struct {
	*TraitSizeGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSizeGroupFromCPointer(p unsafe.Pointer) *SizeGroup {
	ret := &SizeGroup{
		NewTraitSizeGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SizeGroup) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Socket{
		NewTraitSocket(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Socket) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &SpinButton{
		NewTraitSpinButton(p),
		NewTraitEntry(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SpinButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Spinner struct {
	*TraitSpinner
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSpinnerFromCPointer(p unsafe.Pointer) *Spinner {
	ret := &Spinner{
		NewTraitSpinner(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Spinner) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Stack{
		NewTraitStack(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Stack) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &StackSwitcher{
		NewTraitStackSwitcher(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *StackSwitcher) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type StatusIcon struct {
	*TraitStatusIcon
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStatusIconFromCPointer(p unsafe.Pointer) *StatusIcon {
	ret := &StatusIcon{
		NewTraitStatusIcon(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *StatusIcon) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Statusbar{
		NewTraitStatusbar(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Statusbar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Style struct {
	*TraitStyle
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStyleFromCPointer(p unsafe.Pointer) *Style {
	ret := &Style{
		NewTraitStyle(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Style) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type StyleContext struct {
	*TraitStyleContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStyleContextFromCPointer(p unsafe.Pointer) *StyleContext {
	ret := &StyleContext{
		NewTraitStyleContext(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *StyleContext) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type StyleProperties struct {
	*TraitStyleProperties
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewStylePropertiesFromCPointer(p unsafe.Pointer) *StyleProperties {
	ret := &StyleProperties{
		NewTraitStyleProperties(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *StyleProperties) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Switch struct {
	*TraitSwitch
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSwitchFromCPointer(p unsafe.Pointer) *Switch {
	ret := &Switch{
		NewTraitSwitch(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Switch) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Table{
		NewTraitTable(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Table) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &TearoffMenuItem{
		NewTraitTearoffMenuItem(p),
		NewTraitMenuItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TearoffMenuItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TextBuffer struct {
	*TraitTextBuffer
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextBufferFromCPointer(p unsafe.Pointer) *TextBuffer {
	ret := &TextBuffer{
		NewTraitTextBuffer(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TextBuffer) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TextChildAnchor struct {
	*TraitTextChildAnchor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextChildAnchorFromCPointer(p unsafe.Pointer) *TextChildAnchor {
	ret := &TextChildAnchor{
		NewTraitTextChildAnchor(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TextChildAnchor) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TextMark struct {
	*TraitTextMark
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextMarkFromCPointer(p unsafe.Pointer) *TextMark {
	ret := &TextMark{
		NewTraitTextMark(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TextMark) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TextTag struct {
	*TraitTextTag
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextTagFromCPointer(p unsafe.Pointer) *TextTag {
	ret := &TextTag{
		NewTraitTextTag(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TextTag) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TextTagTable struct {
	*TraitTextTagTable
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTextTagTableFromCPointer(p unsafe.Pointer) *TextTagTable {
	ret := &TextTagTable{
		NewTraitTextTagTable(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TextTagTable) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &TextView{
		NewTraitTextView(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TextView) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ThemingEngine struct {
	*TraitThemingEngine
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewThemingEngineFromCPointer(p unsafe.Pointer) *ThemingEngine {
	ret := &ThemingEngine{
		NewTraitThemingEngine(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ThemingEngine) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ToggleAction struct {
	*TraitToggleAction
	*TraitAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewToggleActionFromCPointer(p unsafe.Pointer) *ToggleAction {
	ret := &ToggleAction{
		NewTraitToggleAction(p),
		NewTraitAction(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ToggleAction) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ToggleButton{
		NewTraitToggleButton(p),
		NewTraitButton(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ToggleButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ToggleToolButton{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ToggleToolButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ToolButton{
		NewTraitToolButton(p),
		NewTraitToolItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ToolButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ToolItem{
		NewTraitToolItem(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ToolItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ToolItemGroup{
		NewTraitToolItemGroup(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ToolItemGroup) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &ToolPalette{
		NewTraitToolPalette(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ToolPalette) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Toolbar{
		NewTraitToolbar(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Toolbar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Tooltip struct {
	*TraitTooltip
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTooltipFromCPointer(p unsafe.Pointer) *Tooltip {
	ret := &Tooltip{
		NewTraitTooltip(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Tooltip) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TreeModelFilter struct {
	*TraitTreeModelFilter
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeModelFilterFromCPointer(p unsafe.Pointer) *TreeModelFilter {
	ret := &TreeModelFilter{
		NewTraitTreeModelFilter(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TreeModelFilter) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TreeModelSort struct {
	*TraitTreeModelSort
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeModelSortFromCPointer(p unsafe.Pointer) *TreeModelSort {
	ret := &TreeModelSort{
		NewTraitTreeModelSort(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TreeModelSort) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TreeSelection struct {
	*TraitTreeSelection
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeSelectionFromCPointer(p unsafe.Pointer) *TreeSelection {
	ret := &TreeSelection{
		NewTraitTreeSelection(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TreeSelection) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TreeStore struct {
	*TraitTreeStore
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeStoreFromCPointer(p unsafe.Pointer) *TreeStore {
	ret := &TreeStore{
		NewTraitTreeStore(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TreeStore) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &TreeView{
		NewTraitTreeView(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TreeView) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TreeViewColumn struct {
	*TraitTreeViewColumn
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTreeViewColumnFromCPointer(p unsafe.Pointer) *TreeViewColumn {
	ret := &TreeViewColumn{
		NewTraitTreeViewColumn(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TreeViewColumn) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type UIManager struct {
	*TraitUIManager
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUIManagerFromCPointer(p unsafe.Pointer) *UIManager {
	ret := &UIManager{
		NewTraitUIManager(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *UIManager) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &VBox{
		NewTraitVBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *VBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &VButtonBox{
		NewTraitVButtonBox(p),
		NewTraitButtonBox(p),
		NewTraitBox(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *VButtonBox) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &VPaned{
		NewTraitVPaned(p),
		NewTraitPaned(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *VPaned) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &VScale{
		NewTraitVScale(p),
		NewTraitScale(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *VScale) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &VScrollbar{
		NewTraitVScrollbar(p),
		NewTraitScrollbar(p),
		NewTraitRange(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *VScrollbar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &VSeparator{
		NewTraitVSeparator(p),
		NewTraitSeparator(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *VSeparator) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Viewport{
		NewTraitViewport(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Viewport) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &VolumeButton{
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
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *VolumeButton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Widget struct {
	*TraitWidget
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWidgetFromCPointer(p unsafe.Pointer) *Widget {
	ret := &Widget{
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Widget) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
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
	ret := &Window{
		NewTraitWindow(p),
		NewTraitBin(p),
		NewTraitContainer(p),
		NewTraitWidget(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Window) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type WindowGroup struct {
	*TraitWindowGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWindowGroupFromCPointer(p unsafe.Pointer) *WindowGroup {
	ret := &WindowGroup{
		NewTraitWindowGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *WindowGroup) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}
