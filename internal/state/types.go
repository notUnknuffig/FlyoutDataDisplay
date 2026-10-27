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
	Count       int     `json:"Count"`
	Type        string  `json:"Type"`
	Name        string  `json:"Name"`
	MaxDistance float32 `json:"MaxDistance"`
}

type Radar struct {
	Range  float32 `json:"Range"`
	Mode   string  `json:"Mode"`
	SteerX float32 `json:"SteerX"`
	SteerY float32 `json:"SteerY"`
	SteerZ float32 `json:"SteerZ"`
	Size   float32 `json:"Size"`
}

type Target struct {
	HasTarget bool    `json:"HasTarget"`
	ID        int     `json:"ID"`
	Heading   float32 `json:"Heading"`
	RelX      float32 `json:"RelX"`
	RelY      float32 `json:"RelY"`
	RelZ      float32 `json:"RelZ"`
}

type Input struct {
	FunctionKeyMask int `json:"FunctionKeyMask"` // f20 f19 f18 f17 f16 f15 f14 f13 f12 f11 f10 f9 f8 f7 f6 f5 f4 f3 f2 f1 <- At leadst 20 bits
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
	Beta          float32        `json:"Beta"`
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
	Radar         Radar          `json:"Radar"`
	Target        Target         `json:"Target"`
	Flaps         int32          `json:"Flaps"`
	HasHook       bool           `json:"HasHook"`
	Hook          bool           `json:"Hook"`
	HasApu        bool           `json:"HasApu"`
	Apu           bool           `json:"Apu"`
	HasAirbrake   bool           `json:"HasAirbrake"`
	Airbrake      float32        `json:"Airbrake"`
	LandingGear   bool           `json:"LandingGear"`
	Input         Input          `json:"Input"`
}

type Options struct {
	ResolutionX int
	ResolutionY int
	AspectRatio string
	Fullscreen  bool
	Scale       float32
	Units       *Units
}

type Units struct {
	DistanceUnit       string
	DistanceConversion float64
	SpeedUnit          string
	SpeedConversion    float64
	HightUnit          string
	HightConversion    float64
	ClimbUnit          string
	ClimbConversion    float64
	HeatUnit           string
	HeatConversion     float64
	HeatFreezingPoint  float64
	WeightUnit         string
	WeightConversion   float64
	TonUnit            string // Useless because its going to be "t"
	TonConversion      float64
	FlowRateUnit       string
	FlowRateConversion float64
}
