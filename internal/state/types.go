package state

type State interface {
	Input() State
	Draw()
}

type JetEngine struct {
	Throttle            float32 `json:"Throttle"`
	IdleThrottle        float32 `json:"IdleThrottle"`
	HasAfterburner      bool    `json:"HasAfterburner"`
	AfterburnerThrottle float32 `json:"AfterburnerThrottle"`
	AlternatorPower     float32 `json:"AlternatorPower"`
	HydraulicPower      float32 `json:"HydraulicPower"`
	NetThrust           float32 `json:"NetThrust"`
	FuelFlow            float32 `json:"FuelFlow"`
}

type PistonEngine struct {
	Throttle     float32 `json:"Throttle"`
	IdleThrottle float32 `json:"IdleThrottle"`
	RPM          float32 `json:"RPM"`
	FuelFlow     float32 `json:"FuelFlow"`
	Temperature  float32 `json:"Temperature"`
	Power        float32 `json:"Power"`
}

type FuelTank struct {
	Priority    int     `json:"Priority"`
	IsEmpty     bool    `json:"IsEmpty"`
	Fuel        float32 `json:"Fuel"`
	FuelPercent float32 `json:"FuelPercent"`
	Capacity    float32 `json:"Capacity"`
}

type Missile struct {
	Count int    `json:"Count"`
	Type  string `json:"Type"`
	Name  string `json:"Name"`
}

type FlightData struct {
	Time          float32        `json:"Time"`
	Name          string         `json:"Name"`
	Altitude      float32        `json:"Altitude"`
	Airspeed      float32        `json:"Airspeed"`
	Pitch         float32        `json:"Pitch"`
	Roll          float32        `json:"Roll"`
	G             float32        `json:"G"`
	AGL           float32        `json:"AGL"`
	Heading       float32        `json:"Heading"`
	Climb         float32        `json:"Climb"`
	Mass          float32        `json:"Mass"`
	Mach          float32        `json:"Mach"`
	Alpha         float32        `json:"Alpha"`
	Latitude      float32        `json:"Latitude"`
	Longitude     float32        `json:"Longitude"`
	JetEngines    []JetEngine    `json:"JetEngines"`
	PistonEngines []PistonEngine `json:"PistonEngines"`
	Fuel          float32        `json:"Fuel"`
	FuelCapacity  float32        `json:"FuelCapacity"`
	TimeToEmpty   float32        `json:"TimeToEmpty"`
	FuelRatio     float32        `json:"FuelRatio"`
	FuelTanks     []FuelTank     `json:"FuelTanks"`
	Missiles      []Missile      `json:"Missiles"`
	ActiveMissile string         `json:"ActiveMissile"`
}

type Options struct {
	ResolutionX int
	ResolutionY int
	AspectRatio string
	Fullscreen  bool
	Scale       float32
}
