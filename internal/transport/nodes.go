package transport

import (
	"fmt"
	"strings"
)

var nodes = []string{
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SPD.SP.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S2_D",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xRunCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S2_UD2",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xRunCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xBoost",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SPD.SP.ValR[1]",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xRamp",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S2_UD3",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xFwdCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SPD.SP.ValR[2]",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xFwdCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xBof",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.SEQ_S2_UD.SP.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S2_UD4",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.CUR.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xFwd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xRunCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xFault",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xBoost",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.MV.Val[2]",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SPD.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xRev",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xFwdCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xWarning",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.TRQ.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xReady",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xRevCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xWarningExt",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.CUR.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xFaultExt",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xExcluded",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xCtrlRelease",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xRun",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.G11_B2.PV.ValF",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SPD.SP.ValR[2]",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xFwdCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xLimit_1",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xRunRls",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xBof",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xReady",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xRunLim",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xLimit_2",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xRevRls",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xReset",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xRunCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xRevLim",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xFault",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xFwdRls",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.F.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.I.ValA",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.kP.ValC",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDI.ValB",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDU.ValC",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValBC",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValCN",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.SI1.PV.ValF",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_L",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_LR4",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_UD2",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S2_0",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_SB7",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_SBH4",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xSHL_K54",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.Q.LC2.xPU1_SBH4",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G11_UZ1.IN.p0x6303",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G12_UZ1.IN.p0x6306",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_T11_UZ1.IN.p0x6301",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SPD.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xRev",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xFwdCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xWarning",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.TRQ.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SPD.SP.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xReady",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xRevCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xWarningExt",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.CUR.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xFaultExt",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xExcluded",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xCtrlRelease",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xRun",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.POS.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xFwd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xRunCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xFault",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xBoost",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.G12_B1.PV.ValF",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xReady",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xRunLim",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xLimit_2",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xRevRls",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xReset",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xRevLim",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xFault",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xFwdRls",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SPD.SP.ValR[1]",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xRevCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xFwdLim",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xWarning",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xRamp",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.I.ValDif",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.kP.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.P.ValA",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDI.ValC",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDU.ValA",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValCD",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.SEQ_S1_LR.SP.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.TI1.PV.ValF",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_R",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_0",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_UD3",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_SA1",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xSHL_SB1",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_SBH3",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xSHL_K102",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.Q.LC1.xSH1_K_Rst",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.Q.LC2.xSHL_K11",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.Q.LC2.xPU1_SBH3",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G11_UZ1.IN.p0x6306",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_H11_UZ1.IN.p0x6301",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_T11_UZ1.IN.p0x6303",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SPD.SP.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xReady",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xRevCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xWarningExt",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.ER.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.SP.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xFaultExt",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xExcluded",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xCtrlRelease",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xRun",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.POS.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xFwd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xFault",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SPD.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xRev",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xFwdCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xWarning",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.TRQ.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.G12_B2.PV.ValF",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xRunCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xRevLim",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xFault",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xFwdRls",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xRevCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xFwdLim",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xWarning",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SPD.SP.ValR[2]",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xFwdCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xLimit_1",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xRunRls",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xBof",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.I.ValB",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.kP.ValA",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.P.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDI.ValA",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValDif",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValAN",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.SEQ_S1_UD.SP.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.WI1.PV.ValF",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_LR2",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_U",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_UD4",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_SBA",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xSHL_SB2",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_SBH6",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.Q.LC1.xSH1_K_PLC",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.Q.LC2.xPU1_SBH6",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G12_UZ1.IN.p0x6301",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_H11_UZ1.IN.p0x6303",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_T11_UZ1.IN.p0x6306",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xFaultExt",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xExcluded",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xCtrlRelease",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SW.xRun",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.MV.Val[1]",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CUR.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xFwd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xRunCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xFault",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SW.xBoost",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SPD.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xRev",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SW.xWarning",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.TRQ.PV.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SPD.SP.Val",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xReady",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xRevCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SW.xWarningExt",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.G11_B1.PV.ValF",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SPD.SP.ValR[1]",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xRevCmd",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xFwdLim",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xWarning",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SW.xRamp",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xLimit_1",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SW.xRunRls",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xReady",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xRunLim",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xLimit_2",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xRevRls",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SW.xReset",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.I.ValC",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.kP.ValB",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.P.ValR",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDU.ValB",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValAB",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValBN",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xSHL_1SF2",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_LR3",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S1_D",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_S2_U",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_SB8",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xPU1_SBH1",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.IN.LC2.xSHL_K55",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.Q.LC1.xSH1_K_KMM",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.IO.Q.LC2.xPU1_SBH1",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G11_UZ1.IN.p0x6301",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G12_UZ1.IN.p0x6303",
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_H11_UZ1.IN.p0x6306",
}

// handleToName: ClientHandle → имя тега.
//
// ClientHandle = items.id тега в базе plc200_data (NEW_DATA_GOLANG, localhost),
// поэтому он сразу пишется в history.itemid (см. services.go).
// Номера выгружены из таблицы items 2026-10-05. Если items пересоздать или
// писать в другую базу — номера будут другими, их нужно выгрузить заново.
//
// «было N» — старый порядковый номер (позиция в nodes). До этой правки
// данные в history записывались под ним.
var handleToName = map[uint32]string{
	39:  "DATA.CRANE.DRV_H1_1.SPD.SP.Val",      // было 1
	36:  "DATA.IO.IN.LC2.xPU1_S2_D",            // было 2
	35:  "DATA.CRANE.GEAR_H1.SW.xRunCmd",       // было 3
	34:  "DATA.IO.IN.LC2.xPU1_S2_UD2",          // было 4
	33:  "DATA.CRANE.DRV_H1_1.SW.xRunCmd",      // было 5
	32:  "DATA.CRANE.DRV_H1_1.SW.xBoost",       // было 6
	31:  "DATA.CRANE.GEAR_H1.SPD.SP.ValR[1]",   // было 7
	30:  "DATA.CRANE.GEAR_H1.SW.xRamp",         // было 8
	29:  "DATA.IO.IN.LC2.xPU1_S2_UD3",          // было 9
	28:  "DATA.CRANE.DRV_H1_1.SW.xFwdCmd",      // было 10
	25:  "DATA.CRANE.GEAR_H1.SPD.SP.ValR[2]",   // было 11
	22:  "DATA.CRANE.GEAR_H1.SW.xFwdCmd",       // было 12
	21:  "DATA.CRANE.GEAR_H1.SW.xBof",          // было 13
	18:  "DATA.CRANE.SEQ_S2_UD.SP.Val",         // было 14
	15:  "DATA.IO.IN.LC2.xPU1_S2_UD4",          // было 15
	11:  "DATA.CRANE.DRV_G1_1.CUR.PV.Val",      // было 16
	8:   "DATA.CRANE.DRV_G1_1.SW.xFwd",         // было 17
	7:   "DATA.CRANE.DRV_G1_1.SW.xRunCmd",      // было 18
	6:   "DATA.CRANE.DRV_G1_1.SW.xFault",       // было 19
	5:   "DATA.CRANE.DRV_G1_1.SW.xBoost",       // было 20
	102: "DATA.CRANE.DRV_G1_2.CTRL.MV.Val[2]",  // было 21
	99:  "DATA.CRANE.DRV_G1_2.SPD.PV.Val",      // было 22
	96:  "DATA.CRANE.DRV_G1_2.SW.xRev",         // было 23
	95:  "DATA.CRANE.DRV_G1_2.SW.xFwdCmd",      // было 24
	94:  "DATA.CRANE.DRV_G1_2.SW.xWarning",     // было 25
	92:  "DATA.CRANE.DRV_G1_2.TRQ.PV.Val",      // было 26
	88:  "DATA.CRANE.DRV_H1_1.SW.xReady",       // было 27
	87:  "DATA.CRANE.DRV_H1_1.SW.xRevCmd",      // было 28
	86:  "DATA.CRANE.DRV_H1_1.SW.xWarningExt",  // было 29
	85:  "DATA.CRANE.DRV_T1_1.CUR.PV.Val",      // было 30
	82:  "DATA.CRANE.DRV_T1_1.SW.xFaultExt",    // было 31
	81:  "DATA.CRANE.DRV_T1_1.SW.xExcluded",    // было 32
	80:  "DATA.CRANE.DRV_T1_1.SW.xCtrlRelease", // было 33
	79:  "DATA.CRANE.DRV_T1_1.SW.xRun",         // было 34
	76:  "DATA.CRANE.G11_B2.PV.ValF",           // было 35
	73:  "DATA.CRANE.GEAR_G1.SPD.SP.ValR[2]",   // было 36
	70:  "DATA.CRANE.GEAR_G1.SW.xFwdCmd",       // было 37
	69:  "DATA.CRANE.GEAR_G1.SW.xLimit_1",      // было 38
	68:  "DATA.CRANE.GEAR_G1.SW.xRunRls",       // было 39
	67:  "DATA.CRANE.GEAR_G1.SW.xBof",          // было 40
	64:  "DATA.CRANE.GEAR_H1.SW.xReady",        // было 41
	63:  "DATA.CRANE.GEAR_H1.SW.xRunLim",       // было 42
	62:  "DATA.CRANE.GEAR_H1.SW.xLimit_2",      // было 43
	61:  "DATA.CRANE.GEAR_H1.SW.xRevRls",       // было 44
	60:  "DATA.CRANE.GEAR_H1.SW.xReset",        // было 45
	59:  "DATA.CRANE.GEAR_T1.SW.xRunCmd",       // было 46
	58:  "DATA.CRANE.GEAR_T1.SW.xRevLim",       // было 47
	57:  "DATA.CRANE.GEAR_T1.SW.xFault",        // было 48
	56:  "DATA.CRANE.GEAR_T1.SW.xFwdRls",       // было 49
	53:  "DATA.CRANE.PM_U1.F.Val",              // было 50
	51:  "DATA.CRANE.PM_U1.I.ValA",             // было 51
	49:  "DATA.CRANE.PM_U1.kP.ValC",            // было 52
	47:  "DATA.CRANE.PM_U1.THDI.ValB",          // было 53
	45:  "DATA.CRANE.PM_U1.THDU.ValC",          // было 54
	43:  "DATA.CRANE.PM_U1.U.ValBC",            // было 55
	42:  "DATA.CRANE.PM_U1.U.ValCN",            // было 56
	138: "DATA.CRANE.SI1.PV.ValF",              // было 57
	135: "DATA.IO.IN.LC2.xPU1_S1_L",            // было 58
	134: "DATA.IO.IN.LC2.xPU1_S1_LR4",          // было 59
	133: "DATA.IO.IN.LC2.xPU1_S1_UD2",          // было 60
	132: "DATA.IO.IN.LC2.xPU1_S2_0",            // было 61
	131: "DATA.IO.IN.LC2.xPU1_SB7",             // было 62
	130: "DATA.IO.IN.LC2.xPU1_SBH4",            // было 63
	129: "DATA.IO.IN.LC2.xSHL_K54",             // было 64
	128: "DATA.IO.Q.LC2.xPU1_SBH4",             // было 65
	125: "DATA.NET.GK820_G11_UZ1.IN.p0x6303",   // было 66
	122: "DATA.NET.GK820_G12_UZ1.IN.p0x6306",   // было 67
	119: "DATA.NET.GK820_T11_UZ1.IN.p0x6301",   // было 68
	115: "DATA.CRANE.DRV_G1_1.SPD.PV.Val",      // было 69
	112: "DATA.CRANE.DRV_G1_1.SW.xRev",         // было 70
	111: "DATA.CRANE.DRV_G1_1.SW.xFwdCmd",      // было 71
	110: "DATA.CRANE.DRV_G1_1.SW.xWarning",     // было 72
	109: "DATA.CRANE.DRV_G1_1.TRQ.PV.Val",      // было 73
	106: "DATA.CRANE.DRV_G1_2.CTRL.PV.Val",     // было 74
	104: "DATA.CRANE.DRV_G1_2.SPD.SP.Val",      // было 75
	229: "DATA.CRANE.DRV_G1_2.SW.xReady",       // было 76
	228: "DATA.CRANE.DRV_G1_2.SW.xRevCmd",      // было 77
	227: "DATA.CRANE.DRV_G1_2.SW.xWarningExt",  // было 78
	226: "DATA.CRANE.DRV_H1_1.CUR.PV.Val",      // было 79
	221: "DATA.CRANE.DRV_H1_1.SW.xFaultExt",    // было 80
	214: "DATA.CRANE.DRV_H1_1.SW.xExcluded",    // было 81
	200: "DATA.CRANE.DRV_H1_1.SW.xCtrlRelease", // было 82
	197: "DATA.CRANE.DRV_H1_1.SW.xRun",         // было 83
	196: "DATA.CRANE.DRV_T1_1.POS.PV.Val",      // было 84
	193: "DATA.CRANE.DRV_T1_1.SW.xFwd",         // было 85
	175: "DATA.CRANE.DRV_T1_1.SW.xRunCmd",      // было 86
	173: "DATA.CRANE.DRV_T1_1.SW.xFault",       // было 87
	172: "DATA.CRANE.DRV_T1_1.SW.xBoost",       // было 88
	171: "DATA.CRANE.G12_B1.PV.ValF",           // было 89
	168: "DATA.CRANE.GEAR_G1.SW.xReady",        // было 90
	167: "DATA.CRANE.GEAR_G1.SW.xRunLim",       // было 91
	166: "DATA.CRANE.GEAR_G1.SW.xLimit_2",      // было 92
	165: "DATA.CRANE.GEAR_G1.SW.xRevRls",       // было 93
	164: "DATA.CRANE.GEAR_G1.SW.xReset",        // было 94
	163: "DATA.CRANE.GEAR_H1.SW.xRevLim",       // было 95
	162: "DATA.CRANE.GEAR_H1.SW.xFault",        // было 96
	161: "DATA.CRANE.GEAR_H1.SW.xFwdRls",       // было 97
	160: "DATA.CRANE.GEAR_T1.SPD.SP.ValR[1]",   // было 98
	157: "DATA.CRANE.GEAR_T1.SW.xRevCmd",       // было 99
	156: "DATA.CRANE.GEAR_T1.SW.xFwdLim",       // было 100
	155: "DATA.CRANE.GEAR_T1.SW.xWarning",      // было 101
	154: "DATA.CRANE.GEAR_T1.SW.xRamp",         // было 102
	153: "DATA.CRANE.PM_U1.I.ValDif",           // было 103
	152: "DATA.CRANE.PM_U1.kP.Val",             // было 104
	151: "DATA.CRANE.PM_U1.P.ValA",             // было 105
	149: "DATA.CRANE.PM_U1.THDI.ValC",          // было 106
	148: "DATA.CRANE.PM_U1.THDU.ValA",          // было 107
	147: "DATA.CRANE.PM_U1.U.ValCD",            // было 108
	146: "DATA.CRANE.SEQ_S1_LR.SP.Val",         // было 109
	143: "DATA.CRANE.TI1.PV.ValF",              // было 110
	140: "DATA.IO.IN.LC2.xPU1_S1_R",            // было 111
	139: "DATA.IO.IN.LC2.xPU1_S1_0",            // было 112
	223: "DATA.IO.IN.LC2.xPU1_S1_UD3",          // было 113
	220: "DATA.IO.IN.LC2.xPU1_SA1",             // было 114
	219: "DATA.IO.IN.LC2.xSHL_SB1",             // было 115
	218: "DATA.IO.IN.LC2.xPU1_SBH3",            // было 116
	217: "DATA.IO.IN.LC2.xSHL_K102",            // было 117
	216: "DATA.IO.Q.LC1.xSH1_K_Rst",            // было 118
	212: "DATA.IO.Q.LC2.xSHL_K11",              // было 119
	202: "DATA.IO.Q.LC2.xPU1_SBH3",             // было 120
	199: "DATA.NET.GK820_G11_UZ1.IN.p0x6306",   // было 121
	192: "DATA.NET.GK820_H11_UZ1.IN.p0x6301",   // было 122
	189: "DATA.NET.GK820_T11_UZ1.IN.p0x6303",   // было 123
	188: "DATA.CRANE.DRV_G1_1.SPD.SP.Val",      // было 124
	186: "DATA.CRANE.DRV_G1_1.SW.xReady",       // было 125
	185: "DATA.CRANE.DRV_G1_1.SW.xRevCmd",      // было 126
	184: "DATA.CRANE.DRV_G1_1.SW.xWarningExt",  // было 127
	183: "DATA.CRANE.DRV_G1_2.CTRL.ER.Val",     // было 128
	181: "DATA.CRANE.DRV_G1_2.CTRL.SP.Val",     // было 129
	179: "DATA.CRANE.DRV_G1_2.SW.xFaultExt",    // было 130
	178: "DATA.CRANE.DRV_G1_2.SW.xExcluded",    // было 131
	177: "DATA.CRANE.DRV_G1_2.SW.xCtrlRelease", // было 132
	174: "DATA.CRANE.DRV_G1_2.SW.xRun",         // было 133
	278: "DATA.CRANE.DRV_H1_1.POS.PV.Val",      // было 134
	275: "DATA.CRANE.DRV_H1_1.SW.xFwd",         // было 135
	274: "DATA.CRANE.DRV_H1_1.SW.xFault",       // было 136
	273: "DATA.CRANE.DRV_T1_1.SPD.PV.Val",      // было 137
	270: "DATA.CRANE.DRV_T1_1.SW.xRev",         // было 138
	269: "DATA.CRANE.DRV_T1_1.SW.xFwdCmd",      // было 139
	268: "DATA.CRANE.DRV_T1_1.SW.xWarning",     // было 140
	267: "DATA.CRANE.DRV_T1_1.TRQ.PV.Val",      // было 141
	264: "DATA.CRANE.G12_B2.PV.ValF",           // было 142
	261: "DATA.CRANE.GEAR_G1.SW.xRunCmd",       // было 143
	260: "DATA.CRANE.GEAR_G1.SW.xRevLim",       // было 144
	259: "DATA.CRANE.GEAR_G1.SW.xFault",        // было 145
	258: "DATA.CRANE.GEAR_G1.SW.xFwdRls",       // было 146
	257: "DATA.CRANE.GEAR_H1.SW.xRevCmd",       // было 147
	256: "DATA.CRANE.GEAR_H1.SW.xFwdLim",       // было 148
	255: "DATA.CRANE.GEAR_H1.SW.xWarning",      // было 149
	254: "DATA.CRANE.GEAR_T1.SPD.SP.ValR[2]",   // было 150
	253: "DATA.CRANE.GEAR_T1.SW.xFwdCmd",       // было 151
	252: "DATA.CRANE.GEAR_T1.SW.xLimit_1",      // было 152
	251: "DATA.CRANE.GEAR_T1.SW.xRunRls",       // было 153
	250: "DATA.CRANE.GEAR_T1.SW.xBof",          // было 154
	249: "DATA.CRANE.PM_U1.I.ValB",             // было 155
	248: "DATA.CRANE.PM_U1.kP.ValA",            // было 156
	247: "DATA.CRANE.PM_U1.P.Val",              // было 157
	246: "DATA.CRANE.PM_U1.THDI.ValA",          // было 158
	245: "DATA.CRANE.PM_U1.U.ValDif",           // было 159
	244: "DATA.CRANE.PM_U1.U.ValAN",            // было 160
	243: "DATA.CRANE.SEQ_S1_UD.SP.Val",         // было 161
	240: "DATA.CRANE.WI1.PV.ValF",              // было 162
	237: "DATA.IO.IN.LC2.xPU1_S1_LR2",          // было 163
	236: "DATA.IO.IN.LC2.xPU1_S1_U",            // было 164
	235: "DATA.IO.IN.LC2.xPU1_S1_UD4",          // было 165
	234: "DATA.IO.IN.LC2.xPU1_SBA",             // было 166
	233: "DATA.IO.IN.LC2.xSHL_SB2",             // было 167
	232: "DATA.IO.IN.LC2.xPU1_SBH6",            // было 168
	231: "DATA.IO.Q.LC1.xSH1_K_PLC",            // было 169
	230: "DATA.IO.Q.LC2.xPU1_SBH6",             // было 170
	222: "DATA.NET.GK820_G12_UZ1.IN.p0x6301",   // было 171
	213: "DATA.NET.GK820_H11_UZ1.IN.p0x6303",   // было 172
	211: "DATA.NET.GK820_T11_UZ1.IN.p0x6306",   // было 173
	210: "DATA.CRANE.DRV_G1_1.SW.xFaultExt",    // было 174
	209: "DATA.CRANE.DRV_G1_1.SW.xExcluded",    // было 175
	208: "DATA.CRANE.DRV_G1_1.SW.xCtrlRelease", // было 176
	207: "DATA.CRANE.DRV_G1_1.SW.xRun",         // было 177
	206: "DATA.CRANE.DRV_G1_2.CTRL.MV.Val[1]",  // было 178
	205: "DATA.CRANE.DRV_G1_2.CUR.PV.Val",      // было 179
	201: "DATA.CRANE.DRV_G1_2.SW.xFwd",         // было 180
	198: "DATA.CRANE.DRV_G1_2.SW.xRunCmd",      // было 181
	176: "DATA.CRANE.DRV_G1_2.SW.xFault",       // было 182
	324: "DATA.CRANE.DRV_G1_2.SW.xBoost",       // было 183
	323: "DATA.CRANE.DRV_H1_1.SPD.PV.Val",      // было 184
	321: "DATA.CRANE.DRV_H1_1.SW.xRev",         // было 185
	320: "DATA.CRANE.DRV_H1_1.SW.xWarning",     // было 186
	319: "DATA.CRANE.DRV_H1_1.TRQ.PV.Val",      // было 187
	316: "DATA.CRANE.DRV_T1_1.SPD.SP.Val",      // было 188
	314: "DATA.CRANE.DRV_T1_1.SW.xReady",       // было 189
	313: "DATA.CRANE.DRV_T1_1.SW.xRevCmd",      // было 190
	312: "DATA.CRANE.DRV_T1_1.SW.xWarningExt",  // было 191
	311: "DATA.CRANE.G11_B1.PV.ValF",           // было 192
	308: "DATA.CRANE.GEAR_G1.SPD.SP.ValR[1]",   // было 193
	307: "DATA.CRANE.GEAR_G1.SW.xRevCmd",       // было 194
	306: "DATA.CRANE.GEAR_G1.SW.xFwdLim",       // было 195
	305: "DATA.CRANE.GEAR_G1.SW.xWarning",      // было 196
	304: "DATA.CRANE.GEAR_G1.SW.xRamp",         // было 197
	303: "DATA.CRANE.GEAR_H1.SW.xLimit_1",      // было 198
	302: "DATA.CRANE.GEAR_H1.SW.xRunRls",       // было 199
	301: "DATA.CRANE.GEAR_T1.SW.xReady",        // было 200
	300: "DATA.CRANE.GEAR_T1.SW.xRunLim",       // было 201
	299: "DATA.CRANE.GEAR_T1.SW.xLimit_2",      // было 202
	298: "DATA.CRANE.GEAR_T1.SW.xRevRls",       // было 203
	297: "DATA.CRANE.GEAR_T1.SW.xReset",        // было 204
	296: "DATA.CRANE.PM_U1.I.ValC",             // было 205
	295: "DATA.CRANE.PM_U1.kP.ValB",            // было 206
	294: "DATA.CRANE.PM_U1.P.ValR",             // было 207
	293: "DATA.CRANE.PM_U1.THDU.ValB",          // было 208
	292: "DATA.CRANE.PM_U1.U.ValAB",            // было 209
	291: "DATA.CRANE.PM_U1.U.ValBN",            // было 210
	290: "DATA.IO.IN.LC2.xSHL_1SF2",            // было 211
	289: "DATA.IO.IN.LC2.xPU1_S1_LR3",          // было 212
	288: "DATA.IO.IN.LC2.xPU1_S1_D",            // было 213
	287: "DATA.IO.IN.LC2.xPU1_S2_U",            // было 214
	286: "DATA.IO.IN.LC2.xPU1_SB8",             // было 215
	285: "DATA.IO.IN.LC2.xPU1_SBH1",            // было 216
	284: "DATA.IO.IN.LC2.xSHL_K55",             // было 217
	283: "DATA.IO.Q.LC1.xSH1_K_KMM",            // было 218
	282: "DATA.IO.Q.LC2.xPU1_SBH1",             // было 219
	281: "DATA.NET.GK820_G11_UZ1.IN.p0x6301",   // было 220
	280: "DATA.NET.GK820_G12_UZ1.IN.p0x6303",   // было 221
	279: "DATA.NET.GK820_H11_UZ1.IN.p0x6306",   // было 222
}

// plcNodePrefix — общее начало адреса узла OPC UA; дальше идёт имя тега без «DATA.».
const plcNodePrefix = "ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application."

// nodeHandle: адрес узла из nodes → ClientHandle (= items.id).
// Заполняется в init() из handleToName, чтобы номер тега был задан в одном месте.
var nodeHandle = make(map[string]uint32, len(handleToName))

func init() {
	for h, name := range handleToName {
		nodeHandle[plcNodePrefix+strings.TrimPrefix(name, "DATA.")] = h
	}
	// Проверка при запуске: у каждого узла из nodes должен быть номер.
	// Если его нет — сразу останавливаемся, а не пишем данные под чужим id.
	for _, n := range nodes {
		if _, ok := nodeHandle[n]; !ok {
			panic(fmt.Sprintf("nodes.go: нет номера (items.id) в handleToName для узла %s", n))
		}
	}
}

// numericNodes — теги с типом double из tags.json. Deadband можно
// ставить только на числа: на bool сервер фильтр отклонит.
var numericNodes = map[string]bool{
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SPD.SP.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SPD.SP.ValR[1]":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_H1.SPD.SP.ValR[2]":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.SEQ_S2_UD.SP.Val":        true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.CUR.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.MV.Val[2]": true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SPD.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.TRQ.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.CUR.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.G11_B2.PV.ValF":          true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SPD.SP.ValR[2]":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.F.Val":             true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.I.ValA":            true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.kP.ValC":           true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDI.ValB":         true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDU.ValC":         true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValBC":           true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValCN":           true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.SI1.PV.ValF":             true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G11_UZ1.IN.p0x6303":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G12_UZ1.IN.p0x6306":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_T11_UZ1.IN.p0x6301":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SPD.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.TRQ.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.PV.Val":    true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.SPD.SP.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.CUR.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.POS.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.G12_B1.PV.ValF":          true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SPD.SP.ValR[1]":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.I.ValDif":          true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.kP.Val":            true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.P.ValA":            true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDI.ValC":         true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDU.ValA":         true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValCD":           true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.SEQ_S1_LR.SP.Val":        true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.TI1.PV.ValF":             true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G11_UZ1.IN.p0x6306":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_H11_UZ1.IN.p0x6301":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_T11_UZ1.IN.p0x6303":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_1.SPD.SP.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.ER.Val":    true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.SP.Val":    true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.POS.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SPD.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.TRQ.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.G12_B2.PV.ValF":          true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_T1.SPD.SP.ValR[2]":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.I.ValB":            true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.kP.ValA":           true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.P.Val":             true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDI.ValA":         true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValDif":          true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValAN":           true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.SEQ_S1_UD.SP.Val":        true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.WI1.PV.ValF":             true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G12_UZ1.IN.p0x6301":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_H11_UZ1.IN.p0x6303":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_T11_UZ1.IN.p0x6306":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CTRL.MV.Val[1]": true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_G1_2.CUR.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.SPD.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_H1_1.TRQ.PV.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.DRV_T1_1.SPD.SP.Val":     true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.G11_B1.PV.ValF":          true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.GEAR_G1.SPD.SP.ValR[1]":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.I.ValC":            true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.kP.ValB":           true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.P.ValR":            true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.THDU.ValB":         true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValAB":           true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.CRANE.PM_U1.U.ValBN":           true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G11_UZ1.IN.p0x6301":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_G12_UZ1.IN.p0x6303":  true,
	"ns=4;s=|var|PLC210 (Arm32bit + Linux ).Application.NET.GK820_H11_UZ1.IN.p0x6306":  true,
}
