// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"net"
	"testing"
	"time"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestSessionReportRequest(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Single IE",
			Structured: message.NewSessionReportRequest(
				mp, fo, seid, seq, pri,
				ie.NewReportType(1, 1, 1, 1),
				ie.NewDownlinkDataReport(
					ie.NewPDRID(0xffff),
					ie.NewDownlinkDataServiceInformation(true, true, 0xff, 0xff),
					ie.NewDLDataPacketsSize(0xffff),
				),
				ie.NewUsageReportWithinSessionReportRequest(
					ie.NewURRID(0xffffffff),
					ie.NewURSEQN(0xffffffff),
					ie.NewUsageReportTrigger(0xff, 0xff, 0xff),
					ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewVolumeMeasurement(0x3f, 0x1111111111111111, 0x2222222222222222, 0x3333333333333333, 0x4444444444444444, 0x5555555555555555, 0x6666666666666666),
					ie.NewDurationMeasurement(10*time.Second),
					ie.NewApplicationDetectionInformation(
						ie.NewApplicationID("https://github.com/wmnsk/go-pfcp/"),
						ie.NewApplicationInstanceID("go-pfcp"),
						ie.NewFlowInformation(ie.FlowDirectionDownlink, "go-pfcp"),
						ie.NewPDRID(0xffff),
					),
					ie.NewUEIPAddress(0x02, "127.0.0.1", "", 0, 0),
					ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewUsageInformation(1, 1, 1, 1),
					ie.NewQueryURRReference(0xffffffff),
					ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewEthernetTrafficInformation(
						ie.NewMACAddressesDetected(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
						ie.NewMACAddressesRemoved(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
					),
					ie.NewJoinIPMulticastInformationWithinUsageReport(
						ie.NewIPMulticastAddress(net.ParseIP("127.0.0.1"), nil, net.ParseIP("127.0.0.1"), nil),
						ie.NewSourceIPAddress(net.ParseIP("127.0.0.1"), nil, 24),
					),
					ie.NewLeaveIPMulticastInformationWithinUsageReport(
						ie.NewIPMulticastAddress(net.ParseIP("127.0.0.1"), nil, net.ParseIP("127.0.0.1"), nil),
						ie.NewSourceIPAddress(net.ParseIP("127.0.0.1"), nil, 24),
					),
				),
				ie.NewErrorIndicationReport(
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
				),
				ie.NewLoadControlInformation(ie.NewSequenceNumber(0xffffffff), ie.NewMetric(0x01)),
				ie.NewOverloadControlInformation(
					ie.NewSequenceNumber(0xffffffff),
					ie.NewMetric(0x01),
					ie.NewTimer(20*time.Hour),
					ie.NewOCIFlags(0x01),
				),
				ie.NewAdditionalUsageReportsInformation(0x00ff),
				ie.NewPFCPSRReqFlags(0x01),
				ie.NewFSEID(0x1111111122222222, net.ParseIP("127.0.0.1"), nil, nil),
				ie.NewPacketRateStatusReport(
					ie.NewQERID(0xffffffff),
					ie.NewPacketRateStatus(0x07, 0x1111, 0x2222, 0x3333, 0x4444, time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				),
				ie.NewPortManagementInformationForTSCWithinSessionReportRequest(
					ie.NewPortManagementInformationContainer("go-pfcp"),
				),
				ie.NewSessionReport(
					ie.NewSRRID(255),
					ie.NewAccessAvailabilityControlInformation(
						ie.NewRequestedAccessAvailabilityInformation(1),
					),
					ie.NewQoSMonitoringReport(
						ie.NewQFI(0x01),
						ie.NewQoSMonitoringMeasurement(0x0f, 0x11111111, 0x22222222, 0x33333333),
						ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
						ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					),
				),
			),
			Serialized: []byte{
				0x21, 0x38, 0x02, 0x4f, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x27, 0x00, 0x01, 0x0f,
				0x00, 0x53, 0x00, 0x13,
				0x00, 0x38, 0x00, 0x02, 0xff, 0xff,
				0x00, 0x2d, 0x00, 0x03, 0x03, 0xff, 0xff,
				0x00, 0xfa, 0x00, 0x02, 0xff, 0xff,
				0x00, 0x50, 0x01, 0x5e,
				0x00, 0x51, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x68, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x3f, 0x00, 0x03, 0xff, 0xff, 0xff,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x42, 0x00, 0x31,
				0x3f,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
				0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22,
				0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
				0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
				0x00, 0x43, 0x00, 0x04, 0x00, 0x00, 0x00, 0x0a,
				0x00, 0x44, 0x00, 0x44,
				0x00, 0x18, 0x00, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6d, 0x6e, 0x73, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x2f,
				0x00, 0x5b, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
				0x00, 0x5c, 0x00, 0x0a, 0x01, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
				0x00, 0x38, 0x00, 0x02, 0xff, 0xff,
				0x00, 0x5d, 0x00, 0x05, 0x02, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x45, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x46, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x5a, 0x00, 0x01, 0x0f,
				0x00, 0x7d, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x8f, 0x00, 0x4a,
				0x00, 0x90, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0x91, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0xbd, 0x00, 0x17,
				0x00, 0xbf, 0x00, 0x09,
				0x06,
				0x7f, 0x00, 0x00, 0x01,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0xc0, 0x00, 0x06,
				0x06,
				0x7f, 0x00, 0x00, 0x01, 0x18,
				0x00, 0xbe, 0x00, 0x17,
				0x00, 0xbf, 0x00, 0x09,
				0x06,
				0x7f, 0x00, 0x00, 0x01,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0xc0, 0x00, 0x06,
				0x06,
				0x7f, 0x00, 0x00, 0x01, 0x18,
				0x00, 0x63, 0x00, 0x0d,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x33, 0x00, 0x0d, 0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff, 0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x36, 0x00, 0x17,
				0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x37, 0x00, 0x01, 0x82,
				0x00, 0x6e, 0x00, 0x01, 0x01,
				0x00, 0x7e, 0x00, 0x02, 0x80, 0xff,
				0x00, 0xa1, 0x00, 0x01, 0x01,
				0x00, 0x39, 0x00, 0x0d, 0x02, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0xfc, 0x00, 0x1d,
				0x00, 0x6d, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xc1, 0x00, 0x11,
				0x07,
				0x11, 0x11,
				0x22, 0x22,
				0x33, 0x33,
				0x44, 0x44,
				0x00, 0x00, 0x00, 0x00, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xc9, 0x00, 0x0b,
				0x00, 0xca, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
				0x00, 0xd6, 0x00, 0x38,
				0x00, 0xd7, 0x00, 0x01, 0xff,
				0x00, 0xd8, 0x00, 0x05,
				0x00, 0xd9, 0x00, 0x01, 0x01,
				0x00, 0xf7, 0x00, 0x26,
				0x00, 0x7c, 0x00, 0x01, 0x01,
				0x00, 0xf8, 0x00, 0x0d,
				0x0f,
				0x11, 0x11, 0x11, 0x11,
				0x22, 0x22, 0x22, 0x22,
				0x33, 0x33, 0x33, 0x33,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
			},
		}, {
			Description: "Multiple IEs",
			Structured: message.NewSessionReportRequest(
				mp, fo, seid, seq, pri,
				ie.NewReportType(1, 1, 1, 1),
				ie.NewDownlinkDataReport(
					ie.NewPDRID(0xffff),
					ie.NewDownlinkDataServiceInformation(true, true, 0xff, 0xff),
					ie.NewDLDataPacketsSize(0xffff),
				),
				ie.NewUsageReportWithinSessionReportRequest(
					ie.NewURRID(0xffffffff),
					ie.NewURSEQN(0xffffffff),
					ie.NewUsageReportTrigger(0xff, 0xff, 0xff),
					ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewVolumeMeasurement(0x3f, 0x1111111111111111, 0x2222222222222222, 0x3333333333333333, 0x4444444444444444, 0x5555555555555555, 0x6666666666666666),
					ie.NewDurationMeasurement(10*time.Second),
					ie.NewApplicationDetectionInformation(
						ie.NewApplicationID("https://github.com/wmnsk/go-pfcp/"),
						ie.NewApplicationInstanceID("go-pfcp"),
						ie.NewFlowInformation(ie.FlowDirectionDownlink, "go-pfcp"),
						ie.NewPDRID(0xffff),
					),
					ie.NewUEIPAddress(0x02, "127.0.0.1", "", 0, 0),
					ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewUsageInformation(1, 1, 1, 1),
					ie.NewQueryURRReference(0xffffffff),
					ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewEthernetTrafficInformation(
						ie.NewMACAddressesDetected(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
						ie.NewMACAddressesRemoved(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
					),
					ie.NewJoinIPMulticastInformationWithinUsageReport(
						ie.NewIPMulticastAddress(net.ParseIP("127.0.0.1"), nil, net.ParseIP("127.0.0.1"), nil),
						ie.NewSourceIPAddress(net.ParseIP("127.0.0.1"), nil, 24),
					),
					ie.NewLeaveIPMulticastInformationWithinUsageReport(
						ie.NewIPMulticastAddress(net.ParseIP("127.0.0.1"), nil, net.ParseIP("127.0.0.1"), nil),
						ie.NewSourceIPAddress(net.ParseIP("127.0.0.1"), nil, 24),
					),
				),
				ie.NewErrorIndicationReport(
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
				),
				ie.NewLoadControlInformation(ie.NewSequenceNumber(0xffffffff), ie.NewMetric(0x01)),
				ie.NewOverloadControlInformation(
					ie.NewSequenceNumber(0xffffffff),
					ie.NewMetric(0x01),
					ie.NewTimer(20*time.Hour),
					ie.NewOCIFlags(0x01),
				),
				ie.NewAdditionalUsageReportsInformation(0x00ff),
				ie.NewPFCPSRReqFlags(0x01),
				ie.NewFSEID(0x1111111122222222, net.ParseIP("127.0.0.1"), nil, nil),
				ie.NewPacketRateStatusReport(
					ie.NewQERID(0xffffffff),
					ie.NewPacketRateStatus(0x07, 0x1111, 0x2222, 0x3333, 0x4444, time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				),
				ie.NewPortManagementInformationForTSCWithinSessionReportRequest(
					ie.NewPortManagementInformationContainer("go-pfcp"),
				),
				ie.NewSessionReport(
					ie.NewSRRID(255),
					ie.NewAccessAvailabilityControlInformation(
						ie.NewRequestedAccessAvailabilityInformation(1),
					),
					ie.NewQoSMonitoringReport(
						ie.NewQFI(0x01),
						ie.NewQoSMonitoringMeasurement(0x0f, 0x11111111, 0x22222222, 0x33333333),
						ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
						ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					),
				),
				ie.NewSessionReport(
					ie.NewSRRID(1),
					ie.NewAccessAvailabilityControlInformation(
						ie.NewRequestedAccessAvailabilityInformation(1),
					),
					ie.NewQoSMonitoringReport(
						ie.NewQFI(0x01),
						ie.NewQoSMonitoringMeasurement(0x0f, 0x11111111, 0x22222222, 0x33333333),
						ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
						ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					),
				),
			),
			Serialized: []byte{
				0x21, 0x38, 0x02, 0x8b, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x27, 0x00, 0x01, 0x0f,
				0x00, 0x53, 0x00, 0x13,
				0x00, 0x38, 0x00, 0x02, 0xff, 0xff,
				0x00, 0x2d, 0x00, 0x03, 0x03, 0xff, 0xff,
				0x00, 0xfa, 0x00, 0x02, 0xff, 0xff,
				0x00, 0x50, 0x01, 0x5e,
				0x00, 0x51, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x68, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x3f, 0x00, 0x03, 0xff, 0xff, 0xff,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x42, 0x00, 0x31,
				0x3f,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
				0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22,
				0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
				0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
				0x00, 0x43, 0x00, 0x04, 0x00, 0x00, 0x00, 0x0a,
				0x00, 0x44, 0x00, 0x44,
				0x00, 0x18, 0x00, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6d, 0x6e, 0x73, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x2f,
				0x00, 0x5b, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
				0x00, 0x5c, 0x00, 0x0a, 0x01, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
				0x00, 0x38, 0x00, 0x02, 0xff, 0xff,
				0x00, 0x5d, 0x00, 0x05, 0x02, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x45, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x46, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x5a, 0x00, 0x01, 0x0f,
				0x00, 0x7d, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x8f, 0x00, 0x4a,
				0x00, 0x90, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0x91, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0xbd, 0x00, 0x17,
				0x00, 0xbf, 0x00, 0x09,
				0x06,
				0x7f, 0x00, 0x00, 0x01,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0xc0, 0x00, 0x06,
				0x06,
				0x7f, 0x00, 0x00, 0x01, 0x18,
				0x00, 0xbe, 0x00, 0x17,
				0x00, 0xbf, 0x00, 0x09,
				0x06,
				0x7f, 0x00, 0x00, 0x01,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0xc0, 0x00, 0x06,
				0x06,
				0x7f, 0x00, 0x00, 0x01, 0x18,
				0x00, 0x63, 0x00, 0x0d,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x33, 0x00, 0x0d, 0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff, 0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x36, 0x00, 0x17,
				0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x37, 0x00, 0x01, 0x82,
				0x00, 0x6e, 0x00, 0x01, 0x01,
				0x00, 0x7e, 0x00, 0x02, 0x80, 0xff,
				0x00, 0xa1, 0x00, 0x01, 0x01,
				0x00, 0x39, 0x00, 0x0d, 0x02, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0xfc, 0x00, 0x1d,
				0x00, 0x6d, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xc1, 0x00, 0x11,
				0x07,
				0x11, 0x11,
				0x22, 0x22,
				0x33, 0x33,
				0x44, 0x44,
				0x00, 0x00, 0x00, 0x00, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xc9, 0x00, 0x0b,
				0x00, 0xca, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
				0x00, 0xd6, 0x00, 0x38,
				0x00, 0xd7, 0x00, 0x01, 0xff,
				0x00, 0xd8, 0x00, 0x05,
				0x00, 0xd9, 0x00, 0x01, 0x01,
				0x00, 0xf7, 0x00, 0x26,
				0x00, 0x7c, 0x00, 0x01, 0x01,
				0x00, 0xf8, 0x00, 0x0d,
				0x0f,
				0x11, 0x11, 0x11, 0x11,
				0x22, 0x22, 0x22, 0x22,
				0x33, 0x33, 0x33, 0x33,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xd6, 0x00, 0x38,
				0x00, 0xd7, 0x00, 0x01, 0x01,
				0x00, 0xd8, 0x00, 0x05,
				0x00, 0xd9, 0x00, 0x01, 0x01,
				0x00, 0xf7, 0x00, 0x26,
				0x00, 0x7c, 0x00, 0x01, 0x01,
				0x00, 0xf8, 0x00, 0x0d,
				0x0f,
				0x11, 0x11, 0x11, 0x11,
				0x22, 0x22, 0x22, 0x22,
				0x33, 0x33, 0x33, 0x33,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseSessionReportRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
