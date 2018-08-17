package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type device struct {
	Limit int `json:"Limit"`
	Value int `json:"Value"`
}

type metadata struct {
	SystemID  int    `json:"SystemID,omitempty"`
	FileID    string `json:"FileID,omitempty"`
	SubModule string `json:"SubModule,omitempty"`
}

type input struct {
	ID      int `json:"id"`
	Payload []struct {
		Result struct {
			Metadata       *metadata `json:"Metadata,omitempty"`
			PowerAndEnergy struct {
				BiochemicalConversion          device `json:"Biochemical Conversion"`
				EnergyStorage                  device `json:"Energy Storage"`
				MHDAndRelatedConversion        device `json:"MHD and Related Conversion"`
				NuclearConversion              device `json:"Nuclear Conversion"`
				PhotovoltaicConversion         device `json:"Photovoltaic Conversion"`
				PowerManagementAndDistribution device `json:"Power Management and Distribution"`
				RenewableEnergy                device `json:"Renewable Energy"`
				ThermodynamicConversion        device `json:"Thermodynamic Conversion"`
				ThermoelectricConversion       device `json:"Thermoelectric Conversion"`
				WirelessDistribution           device `json:"Wireless Distribution"`
			} `json:"Power and Energy"`
			Propulsion struct {
				Aerobrake                    device `json:"Aerobrake"`
				AircraftEngines              device `json:"Aircraft Engines"`
				BeamedEnergy                 device `json:"Beamed Energy"`
				Chemical                     device `json:"Chemical"`
				ElectromagneticThrusters     device `json:"Electromagnetic Thrusters"`
				ElectrostaticThrusters       device `json:"Electrostatic Thrusters"`
				FeedSystemComponents         device `json:"Feed System Components"`
				FundamentalPropulsionPhysics device `json:"Fundamental Propulsion Physics"`
				HighEnergyPropellants        device `json:"High Energy Propellants"`
				LaunchAssist                 device `json:"Launch Assist"`
				MHD                          device `json:"MHD"`
				MicroThrusters               device `json:"Micro Thrusters"`
				Monopropellants              device `json:"Monopropellants"`
				Nuclear                      device `json:"Nuclear"`
				PropellantStorage            device `json:"Propellant Storage"`
				Solar                        device `json:"Solar"`
			} `json:"Propulsion"`
			Robotics struct {
				HumanRoboticInterfaces              device `json:"Human-Robotic Interfaces"`
				IntegratedRoboticConceptsAndSystems device `json:"Integrated Robotic Concepts and Systems"`
				Intelligence                        device `json:"Intelligence"`
				Manipulation                        device `json:"Manipulation"`
				Mobility                            device `json:"Mobility"`
				Perception                          device `json:"Perception"`
				Teleoperation                       device `json:"Teleoperation"`
			} `json:"Robotics"`
			SensorsAndSources struct {
				Biochemical                device `json:"Biochemical"`
				Gravitational              device `json:"Gravitational"`
				HighEnergy                 device `json:"High-Energy"`
				LargeAntennasAndTelescopes device `json:"Large Antennas and Telescopes"`
				Microwave                  device `json:"Microwave"`
				Optical                    device `json:"Optical"`
				ParticleAndFields          device `json:"Particle and Fields"`
			} `json:"Sensors and Sources"`
			Structures struct {
				Airframe                      device `json:"Airframe"`
				Airlocks                      device `json:"Airlocks"`
				ControlsStructuresInteraction device `json:"Controls-Structures Interaction"`
			} `json:"Structures"`
			AvionicsAndAstrionics struct {
				AirportInfrastructureAndSafety    device `json:"Airport Infrastructure and Safety"`
				AttitudeDeterminationAndControl   device `json:"Attitude Determination and Control"`
				Guidance                          device `json:"Guidance"`
				OnBoardComputingAndDataManagement device `json:"On-Board Computing and Data Management"`
				PilotSupportSystems               device `json:"Pilot Support Systems"`
				SpaceportInfrastructureAndSafety  device `json:"Spaceport Infrastructure and Safety"`
				Telemetry                         device `json:"Telemetry"`
			} `json:"Avionics and Astrionics"`
			BioTechnology struct {
				AirRevitalizationAndConditioning device `json:"Air Revitalization and Conditioning"`
				BiomassProductionAndStorage      device `json:"Biomass Production and Storage"`
				BiomolecularSensors              device `json:"Biomolecular Sensors"`
			} `json:"Bio-Technology"`
		} `json:"result"`
	} `json:"payload"`
}

func main() {
	data, err := ioutil.ReadFile("../payload.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	payload := input{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("payload %#v\n", payload)
	payload.Payload[0].Result.Metadata = nil // se metadata é um ponteido
	//payload.Payload[0].Result.Metadata = metadata{} // se metadata NÃO é um ponteido

	b, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}
