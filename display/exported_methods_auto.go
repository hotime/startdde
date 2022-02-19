// Code generated by "dbusutil-gen em -type Manager,Monitor"; DO NOT EDIT.

package display

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (v *Manager) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name: "ApplyChanges",
			Fn:   v.ApplyChanges,
		},
		{
			Name:   "AssociateTouch",
			Fn:     v.AssociateTouch,
			InArgs: []string{"outputName", "touchSerial"},
		},
		{
			Name:   "AssociateTouchByUUID",
			Fn:     v.AssociateTouchByUUID,
			InArgs: []string{"outputName", "touchUUID"},
		},
		{
			Name:    "CanRotate",
			Fn:      v.CanRotate,
			OutArgs: []string{"outArg0"},
		},
		{
			Name:    "CanSetBrightness",
			Fn:      v.CanSetBrightness,
			InArgs:  []string{"outputName"},
			OutArgs: []string{"outArg0"},
		},
		{
			Name:   "ChangeBrightness",
			Fn:     v.ChangeBrightness,
			InArgs: []string{"raised"},
		},
		{
			Name:   "DeleteCustomMode",
			Fn:     v.DeleteCustomMode,
			InArgs: []string{"name"},
		},
		{
			Name:    "GetBrightness",
			Fn:      v.GetBrightness,
			OutArgs: []string{"outArg0"},
		},
		{
			Name:    "GetBuiltinMonitor",
			Fn:      v.GetBuiltinMonitor,
			OutArgs: []string{"outArg0", "outArg1"},
		},
		{
			Name:    "GetRealDisplayMode",
			Fn:      v.GetRealDisplayMode,
			OutArgs: []string{"outArg0"},
		},
		{
			Name:    "ListOutputNames",
			Fn:      v.ListOutputNames,
			OutArgs: []string{"outArg0"},
		},
		{
			Name:    "ListOutputsCommonModes",
			Fn:      v.ListOutputsCommonModes,
			OutArgs: []string{"outArg0"},
		},
		{
			Name:   "ModifyConfigName",
			Fn:     v.ModifyConfigName,
			InArgs: []string{"name", "newName"},
		},
		{
			Name: "RefreshBrightness",
			Fn:   v.RefreshBrightness,
		},
		{
			Name: "Reset",
			Fn:   v.Reset,
		},
		{
			Name: "ResetChanges",
			Fn:   v.ResetChanges,
		},
		{
			Name: "Save",
			Fn:   v.Save,
		},
		{
			Name:   "SetAndSaveBrightness",
			Fn:     v.SetAndSaveBrightness,
			InArgs: []string{"outputName", "value"},
		},
		{
			Name:   "SetBrightness",
			Fn:     v.SetBrightness,
			InArgs: []string{"outputName", "value"},
		},
		{
			Name:   "SetColorTemperature",
			Fn:     v.SetColorTemperature,
			InArgs: []string{"value"},
		},
		{
			Name:   "SetMethodAdjustCCT",
			Fn:     v.SetMethodAdjustCCT,
			InArgs: []string{"adjustMethod"},
		},
		{
			Name:   "SetPrimary",
			Fn:     v.SetPrimary,
			InArgs: []string{"outputName"},
		},
		{
			Name:   "SwitchMode",
			Fn:     v.SwitchMode,
			InArgs: []string{"mode", "name"},
		},
	}
}
func (v *Monitor) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:   "Enable",
			Fn:     v.Enable,
			InArgs: []string{"enabled"},
		},
		{
			Name:   "SetMode",
			Fn:     v.SetMode,
			InArgs: []string{"mode"},
		},
		{
			Name:   "SetModeBySize",
			Fn:     v.SetModeBySize,
			InArgs: []string{"width", "height"},
		},
		{
			Name:   "SetPosition",
			Fn:     v.SetPosition,
			InArgs: []string{"X", "y"},
		},
		{
			Name:   "SetReflect",
			Fn:     v.SetReflect,
			InArgs: []string{"value"},
		},
		{
			Name:   "SetRefreshRate",
			Fn:     v.SetRefreshRate,
			InArgs: []string{"value"},
		},
		{
			Name:   "SetRotation",
			Fn:     v.SetRotation,
			InArgs: []string{"value"},
		},
	}
}
