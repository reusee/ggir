package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>
*/
import "C"

type AccelGroupEntry C.GtkAccelGroupEntry
type AccelKey C.GtkAccelKey
type ActionEntry C.GtkActionEntry
type BindingArg C.GtkBindingArg
type BindingEntry C.GtkBindingEntry
type BindingSet C.GtkBindingSet
type BindingSignal C.GtkBindingSignal
type Border C.GtkBorder
type CssSection C.GtkCssSection
type FileFilterInfo C.GtkFileFilterInfo
type FixedChild C.GtkFixedChild
type Gradient C.GtkGradient
type IMContextInfo C.GtkIMContextInfo
type IconSet C.GtkIconSet
type IconSource C.GtkIconSource
type LabelSelectionInfo C.GtkLabelSelectionInfo
type PageRange C.GtkPageRange
type PaperSize C.GtkPaperSize
type RadioActionEntry C.GtkRadioActionEntry
type RcContext C.GtkRcContext
type RcProperty C.GtkRcProperty
type RecentData C.GtkRecentData
type RecentFilterInfo C.GtkRecentFilterInfo
type RecentInfo C.GtkRecentInfo
type RequestedSize C.GtkRequestedSize
type Requisition C.GtkRequisition
type SelectionData C.GtkSelectionData
type SettingsValue C.GtkSettingsValue
type StockItem C.GtkStockItem
type SymbolicColor C.GtkSymbolicColor
type TableChild C.GtkTableChild
type TableRowCol C.GtkTableRowCol
type TargetEntry C.GtkTargetEntry
type TargetList C.GtkTargetList
type TargetPair C.GtkTargetPair
type TextAppearance C.GtkTextAppearance
type TextAttributes C.GtkTextAttributes
type TextBTree C.GtkTextBTree
type TextIter C.GtkTextIter
type ThemeEngine C.GtkThemeEngine
type ToggleActionEntry C.GtkToggleActionEntry
type TreeIter C.GtkTreeIter
type TreePath C.GtkTreePath
type TreeRowReference C.GtkTreeRowReference
type WidgetAuxInfo C.GtkWidgetAuxInfo
type WidgetPath C.GtkWidgetPath
type WindowGeometryInfo C.GtkWindowGeometryInfo
