package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>
*/
import "C"

type _TraitAboutDialog struct{ CPointer *C.GtkAboutDialog }

type _TraitAccelGroup struct{ CPointer *C.GtkAccelGroup }

type _TraitAccelLabel struct{ CPointer *C.GtkAccelLabel }

type _TraitAccelMap struct{ CPointer *C.GtkAccelMap }

type _TraitAccessible struct{ CPointer *C.GtkAccessible }

type _TraitAction struct{ CPointer *C.GtkAction }

type _TraitActionBar struct{ CPointer *C.GtkActionBar }

type _TraitActionGroup struct{ CPointer *C.GtkActionGroup }

type _TraitAdjustment struct{ CPointer *C.GtkAdjustment }

type _TraitAlignment struct{ CPointer *C.GtkAlignment }

type _TraitAppChooserButton struct{ CPointer *C.GtkAppChooserButton }

type _TraitAppChooserDialog struct{ CPointer *C.GtkAppChooserDialog }

type _TraitAppChooserWidget struct{ CPointer *C.GtkAppChooserWidget }

type _TraitApplication struct{ CPointer *C.GtkApplication }

type _TraitApplicationWindow struct{ CPointer *C.GtkApplicationWindow }

type _TraitArrow struct{ CPointer *C.GtkArrow }

type _TraitAspectFrame struct{ CPointer *C.GtkAspectFrame }

type _TraitAssistant struct{ CPointer *C.GtkAssistant }

type _TraitBin struct{ CPointer *C.GtkBin }

type _TraitBox struct{ CPointer *C.GtkBox }

type _TraitBuilder struct{ CPointer *C.GtkBuilder }

type _TraitButton struct{ CPointer *C.GtkButton }

type _TraitButtonBox struct{ CPointer *C.GtkButtonBox }

type _TraitCalendar struct{ CPointer *C.GtkCalendar }

type _TraitCellArea struct{ CPointer *C.GtkCellArea }

type _TraitCellAreaBox struct{ CPointer *C.GtkCellAreaBox }

type _TraitCellAreaContext struct{ CPointer *C.GtkCellAreaContext }

type _TraitCellRenderer struct{ CPointer *C.GtkCellRenderer }

type _TraitCellRendererAccel struct{ CPointer *C.GtkCellRendererAccel }

type _TraitCellRendererCombo struct{ CPointer *C.GtkCellRendererCombo }

type _TraitCellRendererPixbuf struct{ CPointer *C.GtkCellRendererPixbuf }

type _TraitCellRendererProgress struct{ CPointer *C.GtkCellRendererProgress }

type _TraitCellRendererSpin struct{ CPointer *C.GtkCellRendererSpin }

type _TraitCellRendererSpinner struct{ CPointer *C.GtkCellRendererSpinner }

type _TraitCellRendererText struct{ CPointer *C.GtkCellRendererText }

type _TraitCellRendererToggle struct{ CPointer *C.GtkCellRendererToggle }

type _TraitCellView struct{ CPointer *C.GtkCellView }

type _TraitCheckButton struct{ CPointer *C.GtkCheckButton }

type _TraitCheckMenuItem struct{ CPointer *C.GtkCheckMenuItem }

type _TraitClipboard struct{ CPointer *C.GtkClipboard }

type _TraitColorButton struct{ CPointer *C.GtkColorButton }

type _TraitColorChooserDialog struct{ CPointer *C.GtkColorChooserDialog }

type _TraitColorChooserWidget struct{ CPointer *C.GtkColorChooserWidget }

type _TraitColorSelection struct{ CPointer *C.GtkColorSelection }

type _TraitColorSelectionDialog struct{ CPointer *C.GtkColorSelectionDialog }

type _TraitComboBox struct{ CPointer *C.GtkComboBox }

type _TraitComboBoxText struct{ CPointer *C.GtkComboBoxText }

type _TraitContainer struct{ CPointer *C.GtkContainer }

type _TraitCssProvider struct{ CPointer *C.GtkCssProvider }

type _TraitDialog struct{ CPointer *C.GtkDialog }

type _TraitDrawingArea struct{ CPointer *C.GtkDrawingArea }

type _TraitEntry struct{ CPointer *C.GtkEntry }

type _TraitEntryBuffer struct{ CPointer *C.GtkEntryBuffer }

type _TraitEntryCompletion struct{ CPointer *C.GtkEntryCompletion }

type _TraitEventBox struct{ CPointer *C.GtkEventBox }

type _TraitExpander struct{ CPointer *C.GtkExpander }

type _TraitFileChooserButton struct{ CPointer *C.GtkFileChooserButton }

type _TraitFileChooserDialog struct{ CPointer *C.GtkFileChooserDialog }

type _TraitFileChooserWidget struct{ CPointer *C.GtkFileChooserWidget }

type _TraitFileFilter struct{ CPointer *C.GtkFileFilter }

type _TraitFixed struct{ CPointer *C.GtkFixed }

type _TraitFlowBox struct{ CPointer *C.GtkFlowBox }

type _TraitFlowBoxChild struct{ CPointer *C.GtkFlowBoxChild }

type _TraitFontButton struct{ CPointer *C.GtkFontButton }

type _TraitFontChooserDialog struct{ CPointer *C.GtkFontChooserDialog }

type _TraitFontChooserWidget struct{ CPointer *C.GtkFontChooserWidget }

type _TraitFontSelection struct{ CPointer *C.GtkFontSelection }

type _TraitFontSelectionDialog struct{ CPointer *C.GtkFontSelectionDialog }

type _TraitFrame struct{ CPointer *C.GtkFrame }

type _TraitGrid struct{ CPointer *C.GtkGrid }

type _TraitHBox struct{ CPointer *C.GtkHBox }

type _TraitHButtonBox struct{ CPointer *C.GtkHButtonBox }

type _TraitHPaned struct{ CPointer *C.GtkHPaned }

type _TraitHSV struct{ CPointer *C.GtkHSV }

type _TraitHScale struct{ CPointer *C.GtkHScale }

type _TraitHScrollbar struct{ CPointer *C.GtkHScrollbar }

type _TraitHSeparator struct{ CPointer *C.GtkHSeparator }

type _TraitHandleBox struct{ CPointer *C.GtkHandleBox }

type _TraitHeaderBar struct{ CPointer *C.GtkHeaderBar }

type _TraitIMContext struct{ CPointer *C.GtkIMContext }

type _TraitIMContextSimple struct{ CPointer *C.GtkIMContextSimple }

type _TraitIMMulticontext struct{ CPointer *C.GtkIMMulticontext }

type _TraitIconFactory struct{ CPointer *C.GtkIconFactory }

type _TraitIconInfo struct{ CPointer *C.GtkIconInfo }

type _TraitIconTheme struct{ CPointer *C.GtkIconTheme }

type _TraitIconView struct{ CPointer *C.GtkIconView }

type _TraitImage struct{ CPointer *C.GtkImage }

type _TraitImageMenuItem struct{ CPointer *C.GtkImageMenuItem }

type _TraitInfoBar struct{ CPointer *C.GtkInfoBar }

type _TraitInvisible struct{ CPointer *C.GtkInvisible }

type _TraitLabel struct{ CPointer *C.GtkLabel }

type _TraitLayout struct{ CPointer *C.GtkLayout }

type _TraitLevelBar struct{ CPointer *C.GtkLevelBar }

type _TraitLinkButton struct{ CPointer *C.GtkLinkButton }

type _TraitListBox struct{ CPointer *C.GtkListBox }

type _TraitListBoxRow struct{ CPointer *C.GtkListBoxRow }

type _TraitListStore struct{ CPointer *C.GtkListStore }

type _TraitLockButton struct{ CPointer *C.GtkLockButton }

type _TraitMenu struct{ CPointer *C.GtkMenu }

type _TraitMenuBar struct{ CPointer *C.GtkMenuBar }

type _TraitMenuButton struct{ CPointer *C.GtkMenuButton }

type _TraitMenuItem struct{ CPointer *C.GtkMenuItem }

type _TraitMenuShell struct{ CPointer *C.GtkMenuShell }

type _TraitMenuToolButton struct{ CPointer *C.GtkMenuToolButton }

type _TraitMessageDialog struct{ CPointer *C.GtkMessageDialog }

type _TraitMisc struct{ CPointer *C.GtkMisc }

type _TraitMountOperation struct{ CPointer *C.GtkMountOperation }

type _TraitNotebook struct{ CPointer *C.GtkNotebook }

type _TraitNumerableIcon struct{ CPointer *C.GtkNumerableIcon }

type _TraitOffscreenWindow struct{ CPointer *C.GtkOffscreenWindow }

type _TraitOverlay struct{ CPointer *C.GtkOverlay }

type _TraitPageSetup struct{ CPointer *C.GtkPageSetup }

type _TraitPaned struct{ CPointer *C.GtkPaned }

type _TraitPlacesSidebar struct{ CPointer *C.GtkPlacesSidebar }

type _TraitPlug struct{ CPointer *C.GtkPlug }

type _TraitPopover struct{ CPointer *C.GtkPopover }

type _TraitPrintContext struct{ CPointer *C.GtkPrintContext }

type _TraitPrintOperation struct{ CPointer *C.GtkPrintOperation }

type _TraitPrintSettings struct{ CPointer *C.GtkPrintSettings }

type _TraitProgressBar struct{ CPointer *C.GtkProgressBar }

type _TraitRadioAction struct{ CPointer *C.GtkRadioAction }

type _TraitRadioButton struct{ CPointer *C.GtkRadioButton }

type _TraitRadioMenuItem struct{ CPointer *C.GtkRadioMenuItem }

type _TraitRadioToolButton struct{ CPointer *C.GtkRadioToolButton }

type _TraitRange struct{ CPointer *C.GtkRange }

type _TraitRcStyle struct{ CPointer *C.GtkRcStyle }

type _TraitRecentAction struct{ CPointer *C.GtkRecentAction }

type _TraitRecentChooserDialog struct{ CPointer *C.GtkRecentChooserDialog }

type _TraitRecentChooserMenu struct{ CPointer *C.GtkRecentChooserMenu }

type _TraitRecentChooserWidget struct{ CPointer *C.GtkRecentChooserWidget }

type _TraitRecentFilter struct{ CPointer *C.GtkRecentFilter }

type _TraitRecentManager struct{ CPointer *C.GtkRecentManager }

type _TraitRevealer struct{ CPointer *C.GtkRevealer }

type _TraitScale struct{ CPointer *C.GtkScale }

type _TraitScaleButton struct{ CPointer *C.GtkScaleButton }

type _TraitScrollbar struct{ CPointer *C.GtkScrollbar }

type _TraitScrolledWindow struct{ CPointer *C.GtkScrolledWindow }

type _TraitSearchBar struct{ CPointer *C.GtkSearchBar }

type _TraitSearchEntry struct{ CPointer *C.GtkSearchEntry }

type _TraitSeparator struct{ CPointer *C.GtkSeparator }

type _TraitSeparatorMenuItem struct{ CPointer *C.GtkSeparatorMenuItem }

type _TraitSeparatorToolItem struct{ CPointer *C.GtkSeparatorToolItem }

type _TraitSettings struct{ CPointer *C.GtkSettings }

type _TraitSizeGroup struct{ CPointer *C.GtkSizeGroup }

type _TraitSocket struct{ CPointer *C.GtkSocket }

type _TraitSpinButton struct{ CPointer *C.GtkSpinButton }

type _TraitSpinner struct{ CPointer *C.GtkSpinner }

type _TraitStack struct{ CPointer *C.GtkStack }

type _TraitStackSwitcher struct{ CPointer *C.GtkStackSwitcher }

type _TraitStatusIcon struct{ CPointer *C.GtkStatusIcon }

type _TraitStatusbar struct{ CPointer *C.GtkStatusbar }

type _TraitStyle struct{ CPointer *C.GtkStyle }

type _TraitStyleContext struct{ CPointer *C.GtkStyleContext }

type _TraitStyleProperties struct{ CPointer *C.GtkStyleProperties }

type _TraitSwitch struct{ CPointer *C.GtkSwitch }

type _TraitTable struct{ CPointer *C.GtkTable }

type _TraitTearoffMenuItem struct{ CPointer *C.GtkTearoffMenuItem }

type _TraitTextBuffer struct{ CPointer *C.GtkTextBuffer }

type _TraitTextChildAnchor struct{ CPointer *C.GtkTextChildAnchor }

type _TraitTextMark struct{ CPointer *C.GtkTextMark }

type _TraitTextTag struct{ CPointer *C.GtkTextTag }

type _TraitTextTagTable struct{ CPointer *C.GtkTextTagTable }

type _TraitTextView struct{ CPointer *C.GtkTextView }

type _TraitThemingEngine struct{ CPointer *C.GtkThemingEngine }

type _TraitToggleAction struct{ CPointer *C.GtkToggleAction }

type _TraitToggleButton struct{ CPointer *C.GtkToggleButton }

type _TraitToggleToolButton struct{ CPointer *C.GtkToggleToolButton }

type _TraitToolButton struct{ CPointer *C.GtkToolButton }

type _TraitToolItem struct{ CPointer *C.GtkToolItem }

type _TraitToolItemGroup struct{ CPointer *C.GtkToolItemGroup }

type _TraitToolPalette struct{ CPointer *C.GtkToolPalette }

type _TraitToolbar struct{ CPointer *C.GtkToolbar }

type _TraitTooltip struct{ CPointer *C.GtkTooltip }

type _TraitTreeModelFilter struct{ CPointer *C.GtkTreeModelFilter }

type _TraitTreeModelSort struct{ CPointer *C.GtkTreeModelSort }

type _TraitTreeSelection struct{ CPointer *C.GtkTreeSelection }

type _TraitTreeStore struct{ CPointer *C.GtkTreeStore }

type _TraitTreeView struct{ CPointer *C.GtkTreeView }

type _TraitTreeViewColumn struct{ CPointer *C.GtkTreeViewColumn }

type _TraitUIManager struct{ CPointer *C.GtkUIManager }

type _TraitVBox struct{ CPointer *C.GtkVBox }

type _TraitVButtonBox struct{ CPointer *C.GtkVButtonBox }

type _TraitVPaned struct{ CPointer *C.GtkVPaned }

type _TraitVScale struct{ CPointer *C.GtkVScale }

type _TraitVScrollbar struct{ CPointer *C.GtkVScrollbar }

type _TraitVSeparator struct{ CPointer *C.GtkVSeparator }

type _TraitViewport struct{ CPointer *C.GtkViewport }

type _TraitVolumeButton struct{ CPointer *C.GtkVolumeButton }

type _TraitWidget struct{ CPointer *C.GtkWidget }

type _TraitWindow struct{ CPointer *C.GtkWindow }

type _TraitWindowGroup struct{ CPointer *C.GtkWindowGroup }
