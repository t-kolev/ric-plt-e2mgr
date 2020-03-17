//
// Copyright 2019 AT&T Intellectual Property
// Copyright 2019 Nokia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//  This source code is part of the near-RT RIC (RAN Intelligent Controller)
//  platform project (RICP).

package models

import (
	"encoding/xml"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
)

type E2SetupRequestMessage struct {
	XMLName xml.Name `xml:"E2SetupRequestMessage"`
	Text    string   `xml:",chardata"`
	E2APPDU struct {
		Text              string `xml:",chardata"`
		InitiatingMessage struct {
			Text          string `xml:",chardata"`
			ProcedureCode string `xml:"procedureCode"`
			Criticality   struct {
				Text   string `xml:",chardata"`
				Reject string `xml:"reject"`
			} `xml:"criticality"`
			Value struct {
				Text           string `xml:",chardata"`
				E2setupRequest struct {
					Text        string `xml:",chardata"`
					ProtocolIEs struct {
						Text              string `xml:",chardata"`
						E2setupRequestIEs []struct {
							Text        string `xml:",chardata"`
							ID          string `xml:"id"`
							Criticality struct {
								Text   string `xml:",chardata"`
								Reject string `xml:"reject"`
							} `xml:"criticality"`
							Value struct {
								Text           string `xml:",chardata"`
								GlobalE2nodeID struct {
									Text string `xml:",chardata"`
									GNB  struct {
										Text        string `xml:",chardata"`
										GlobalGNBID struct {
											Text   string `xml:",chardata"`
											PlmnID string `xml:"plmn-id"`
											GnbID  struct {
												Text  string `xml:",chardata"`
												GnbID string `xml:"gnb-ID"`
											} `xml:"gnb-id"`
										} `xml:"global-gNB-ID"`
									} `xml:"gNB"`
									EnGNB struct {
										Text        string `xml:",chardata"`
										GlobalGNBID struct {
											Text   string `xml:",chardata"`
											PlmnID string `xml:"plmn-id"`
											GnbID  struct {
												Text  string `xml:",chardata"`
												GnbID string `xml:"gnb-ID"`
											} `xml:"gnb-id"`
										} `xml:"global-gNB-ID"`
									} `xml:"en-gNB"`
									NgENB struct {
										Text          string `xml:",chardata"`
										GlobalNgENBID struct {
											Text   string `xml:",chardata"`
											PlmnID string `xml:"plmn-id"`
											GnbID  struct {
												Text  string `xml:",chardata"`
												GnbID string `xml:"gnb-ID"`
											} `xml:"gnb-id"`
										} `xml:"global-ng-eNB-ID"`
									} `xml:"ng-eNB"`
									ENB struct {
										Text        string `xml:",chardata"`
										GlobalENBID struct {
											Text   string `xml:",chardata"`
											PlmnID string `xml:"plmn-id"`
											GnbID  struct {
												Text  string `xml:",chardata"`
												GnbID string `xml:"gnb-ID"`
											} `xml:"gnb-id"`
										} `xml:"global-eNB-ID"`
									} `xml:"eNB"`
								} `xml:"GlobalE2node-ID"`
								RANfunctionsList struct {
									Text                      string `xml:",chardata"`
									ProtocolIESingleContainer []struct {
										Text        string `xml:",chardata"`
										ID          string `xml:"id"`
										Criticality struct {
											Text   string `xml:",chardata"`
											Reject string `xml:"reject"`
										} `xml:"criticality"`
										Value struct {
											Text            string `xml:",chardata"`
											RANfunctionItem struct {
												Text                  string `xml:",chardata"`
												RanFunctionID         string `xml:"ranFunctionID"`
												RanFunctionDefinition string `xml:"ranFunctionDefinition"`
												RanFunctionRevision   string `xml:"ranFunctionRevision"`
											} `xml:"RANfunction-Item"`
										} `xml:"value"`
									} `xml:"ProtocolIE-SingleContainer"`
								} `xml:"RANfunctions-List"`
							} `xml:"value"`
						} `xml:"E2setupRequestIEs"`
					} `xml:"protocolIEs"`
				} `xml:"E2setupRequest"`
			} `xml:"value"`
		} `xml:"initiatingMessage"`
	} `xml:"E2AP-PDU"`
}

func (m *E2SetupRequestMessage) GetNodeType() entities.Node_Type{
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.GNB.GlobalGNBID.PlmnID; id!= ""{
		return entities.Node_GNB
	}
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.EnGNB.GlobalGNBID.PlmnID; id!= ""{
		return entities.Node_GNB
	}
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.ENB.GlobalENBID.PlmnID; id!= ""{
		return entities.Node_ENB
	}
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.NgENB.GlobalNgENBID.PlmnID; id!= ""{
		return entities.Node_GNB
	}
	return entities.Node_UNKNOWN
}

func (m *E2SetupRequestMessage) GetPlmnId() string{
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.GNB.GlobalGNBID.PlmnID; id!= ""{
		return id
	}
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.EnGNB.GlobalGNBID.PlmnID; id!= ""{
		return id
	}
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.ENB.GlobalENBID.PlmnID; id!= ""{
		return id
	}
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.NgENB.GlobalNgENBID.PlmnID; id!= ""{
		return id
	}
	return ""
}

func (m *E2SetupRequestMessage) GetNbId() string{
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.GNB.GlobalGNBID.GnbID.GnbID; id!= ""{
		return id
	}
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.EnGNB.GlobalGNBID.GnbID.GnbID; id!= ""{
		return id
	}
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.ENB.GlobalENBID.GnbID.GnbID; id!= ""{
		return id
	}
	if id := m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID.NgENB.GlobalNgENBID.GnbID.GnbID; id!= ""{
		return id
	}
	return ""
}