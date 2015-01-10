package naif

//go:generate go run gennames.go

// Barycenters
const (
	BarycenterSolarSystem Code = 0
	BarycenterMercury     Code = 1
	BarycenterVenus       Code = 2
	BarycenterEarth       Code = 3
	BarycenterMars        Code = 4
	BarycenterJupiter     Code = 5
	BarycenterSaturn      Code = 6
	BarycenterUranus      Code = 7
	BarycenterNeptune     Code = 8
	BarycenterPluto       Code = 9
	Sun                   Code = 10
)

// Planets
const (
	Mercury Code = 199
	Venus   Code = 299
	Earth   Code = 399
	Mars    Code = 499
	Jupiter Code = 599
	Saturn  Code = 699
	Uranus  Code = 799
	Neptune Code = 899
	Pluto   Code = 999
)

// Moons
const (
	Moon Code = 301

	Phobos Code = 401
	Deimos Code = 402

	Io       Code = 501
	Europa   Code = 502
	Ganymede Code = 503
	Callisto Code = 504
	Amalthea Code = 505
	Himalia  Code = 506
	Elara    Code = 507
	Pasiphae Code = 508
	Sinope   Code = 509
	Lysithea Code = 510
	Carme    Code = 511
	Ananke   Code = 512
	Leda     Code = 513
	Thebe    Code = 514
	Adrastea Code = 515

	Metis      Code = 516
	Mimas      Code = 601
	Enceladus  Code = 602
	Tethys     Code = 603
	Dione      Code = 604
	Rhea       Code = 605
	Titan      Code = 606
	Hyperion   Code = 607
	Iapetus    Code = 608
	Phoebe     Code = 609
	Janus      Code = 610
	Epimetheus Code = 611
	Helene     Code = 612
	Telesto    Code = 613
	Calypso    Code = 614
	Atlas      Code = 615
	Prometheus Code = 616
	Pandora    Code = 617
	Pan        Code = 618

	Ariel     Code = 701
	Umbriel   Code = 702
	Titania   Code = 703
	Oberon    Code = 704
	Miranda   Code = 705
	Cordelia  Code = 706
	Ophelia   Code = 707
	Bianca    Code = 708
	Cressida  Code = 709
	Desdemona Code = 710
	Juliet    Code = 711
	Portia    Code = 712
	Rosalind  Code = 713
	Belinda   Code = 714
	Puck      Code = 715

	Triton   Code = 801
	Nereid   Code = 802
	Naiad    Code = 803
	Thalassa Code = 804
	Despina  Code = 805
	Galatea  Code = 806
	Larissa  Code = 807
	Proteus  Code = 808

	Charon   Code = 901
	Nix      Code = 902
	Hydra    Code = 903
	Kerberos Code = 904
	Styx     Code = 905
)

// Spacecraft tracked by the Deep Space Network.
const (
	Pioneer12               Code = -12
	Magellan                Code = -18
	Viking1Orbiter          Code = -27
	Viking2Orbiter          Code = -30
	Voyager1                Code = -31
	Voyager2                Code = -32
	Clementine              Code = -40
	Sakigake                Code = -46
	Suisei                  Code = -47
	MarsPathfinder          Code = -53
	AsteroidRetrieval       Code = -54
	Ulysses                 Code = -55
	Vsop                    Code = -58
	Radioastron             Code = -59
	OsirisREx               Code = -64
	Vega1                   Code = -66
	Vega2                   Code = -67
	MercuryPlanetaryOrbiter Code = -69
	MarsScienceLab          Code = -76
	GalileoOrbiter          Code = -77
	Giotto                  Code = -78
	Spitzer                 Code = -79
	CassiniItl              Code = -81
	Cassini                 Code = -82
	CassiniSimulation       Code = -90
	NearRendezvous          Code = -93
	MarsGlobalSurveyor      Code = -94
	MGSSimulation           Code = -95
	Ice                     Code = -112
	Hayabusa                Code = -130
	CassiniHuygens          Code = -150
	Maven                   Code = -202
	GalileoProbe            Code = -344
	Mars96                  Code = -550
)

// Spacecraft orbiting Earth
//
// The standard NAIF numbering system for spacecraft orbiting the Earth
// (no Deep Space Network ID) is based on the tracking ID from NORAD:
//
//	NAIF Code = -100000 - NORAD ID code
//
// There are many. One is here as an example.
const (
	NOAA9 Code = -115427
)

// Tracking Stations
const (
	NewNorcia Code = 398990
	Goldstone Code = 399001
	Canberra  Code = 399002
	Madrid    Code = 399003
	Usuda     Code = 399004
	Parkes    Code = 399005
	DSS12     Code = 399012
	DSS13     Code = 399013
	DSS14     Code = 399014
	DSS15     Code = 399015
	DSS16     Code = 399016
	DSS17     Code = 399017
	DSS23     Code = 399023
	DSS24     Code = 399024
	DSS25     Code = 399025
	DSS26     Code = 399026
	DSS27     Code = 399027
	DSS28     Code = 399028
)

// Asteroids
//
// The standard NAIF numbering system for asteroids is based on the
// JPL asteroid number:
//
//	NAIF Code = 2000000 + JPL Asteroid number
//
// There are three exceptions to this: Gaspra, Ida, and Dactyl.
const (
	Gaspra           Code = 9511010
	Ida              Code = 2431010
	Dactyl           Code = 2431011
	Ceres            Code = 2000001
	Pallas           Code = 2000002
	Vesta            Code = 2000004
	Lutetia          Code = 2000021
	Kleopatra        Code = 2000216
	Eros             Code = 2000433
	Davida           Code = 2000511
	Mathilde         Code = 2000253
	Steins           Code = 2002867
	Braille          Code = 2009969
	WilsonHarrington Code = 2004015
	Toutatis         Code = 2004179
	Itokawa          Code = 2025143
)

// Comets
//
// This is subset of the small bodies tracked in the solar system.
// See http://ssd.jpl.nasa.gov/sbdb.cgi for all the comet IDs.
const (
	Arend                 Code = 1000001
	ArendRigaux           Code = 1000002
	AshbrookJackson       Code = 1000003
	Boethin               Code = 1000004
	Borrelly              Code = 1000005
	BowellSkiff           Code = 1000006
	Bradfield             Code = 1000007
	Brooks2               Code = 1000008
	BrorsenMetcalf        Code = 1000009
	Bus                   Code = 1000010
	Chernykh              Code = 1000011
	ChuryumovGerasimenko  Code = 1000012
	Ciffreo               Code = 1000013
	Clark                 Code = 1000014
	ComasSola             Code = 1000015
	Crommelin             Code = 1000016
	Arrest                Code = 1000017
	Daniel                Code = 1000018
	DeVicoSwift           Code = 1000019
	DenningFujikawa       Code = 1000020
	Dutoit1               Code = 1000021
	DutoitHartley         Code = 1000022
	DutoitNeujminDelporte Code = 1000023
	Dubiago               Code = 1000024
	Encke                 Code = 1000025
	Faye                  Code = 1000026
	Finlay                Code = 1000027
	Forbes                Code = 1000028
	Gehrels1              Code = 1000029
	Gehrels2              Code = 1000030
	Gehrels3              Code = 1000031
	GiacobiniZinner       Code = 1000032
	Giclas                Code = 1000033
	GriggSkjellerup       Code = 1000034
	Gunn                  Code = 1000035
	Halley                Code = 1000036
	HanedaCampos          Code = 1000037
	Harrington            Code = 1000038
	HarringtonAbell       Code = 1000039
	Hartley1              Code = 1000040
	Hartley2              Code = 1000041
	HartleyIras           Code = 1000042
	HerschelRigollet      Code = 1000043
	Holmes                Code = 1000044
	HondaMrkosPajdusakova Code = 1000045
	Howell                Code = 1000046
	Iras                  Code = 1000047
	JacksonNeujmin        Code = 1000048
	Johnson               Code = 1000049
	KearnsKwee            Code = 1000050
	Klemola               Code = 1000051
	Kohoutek              Code = 1000052
	Kojima                Code = 1000053
	Kopff                 Code = 1000054
	Kowal1                Code = 1000055
	Kowal2                Code = 1000056
	KowalMrkos            Code = 1000057
	KowalVavrova          Code = 1000058
	Longmore              Code = 1000059
	Lovas1                Code = 1000060
	Machholz              Code = 1000061
	Maury                 Code = 1000062
	Neujmin1              Code = 1000063
	Neujmin2              Code = 1000064
	Neujmin3              Code = 1000065
	Olbers                Code = 1000066
	PetersHartley         Code = 1000067
	PonsBrooks            Code = 1000068
	PonsWinnecke          Code = 1000069
	Reinmuth1             Code = 1000070
	Reinmuth2             Code = 1000071
	Russell1              Code = 1000072
	Russell2              Code = 1000073
	Russell3              Code = 1000074
	Russell4              Code = 1000075
	Sanguin               Code = 1000076
	Schaumasse            Code = 1000077
	Schuster              Code = 1000078
	SchwassmannWachmann1  Code = 1000079
	SchwassmannWachmann2  Code = 1000080
	SchwassmannWachmann3  Code = 1000081
	ShajnSchaldach        Code = 1000082
	Shoemaker1            Code = 1000083
	Shoemaker2            Code = 1000084
	Shoemaker3            Code = 1000085
	SingerBrewster        Code = 1000086
	SlaughterBurnham      Code = 1000087
	SmirnovaChernykh      Code = 1000088
	StephanOterma         Code = 1000089
	SwiftGehrels          Code = 1000090
	Takamizawa            Code = 1000091
	Taylor                Code = 1000092
	Tempel1               Code = 1000093
	Tempel2               Code = 1000094
	TempelTuttle          Code = 1000095
	Tritton               Code = 1000096
	Tsuchinshan1          Code = 1000097
	Tsuchinshan2          Code = 1000098
	Tuttle                Code = 1000099
	TuttleGiacobiniKresak Code = 1000100
	Vaisala1              Code = 1000101
	VanBiesbroeck         Code = 1000102
	VanHouten             Code = 1000103
	WestKohoutekIkemura   Code = 1000104
	Whipple               Code = 1000105
	Wild1                 Code = 1000106
	Wild2                 Code = 1000107
	Wild3                 Code = 1000108
	Wirtanen              Code = 1000109
	Wolf                  Code = 1000110
	WolfHarrington        Code = 1000111
	Lovas2                Code = 1000112
	UrataNiijima          Code = 1000113
	WisemanSkiff          Code = 1000114
	Helin                 Code = 1000115
	Mueller               Code = 1000116
	ShoemakerHolt1        Code = 1000117
	HelinRomanCrockett    Code = 1000118
	Hartley3              Code = 1000119
	ParkerHartley         Code = 1000120
	HelinRomanAlu1        Code = 1000121
	Wild4                 Code = 1000122
	Mueller2              Code = 1000123
	Mueller3              Code = 1000124
	ShoemakerLevy1        Code = 1000125
	ShoemakerLevy2        Code = 1000126
	HoltOlmstead          Code = 1000127
	MetcalfBrewington     Code = 1000128
	Levy                  Code = 1000129
	ShoemakerLevy9        Code = 1000130
	Hyakutake             Code = 1000131
	HaleBopp              Code = 1000132
)

// The comet Shoemaker Levy 9 was torn apart by Jupiter and while the
// remnants were in orbit they were tracked.
const (
	ShoemakerLevy9W  Code = 50000001
	ShoemakerLevy9V  Code = 50000002
	ShoemakerLevy9U  Code = 50000003
	ShoemakerLevy9T  Code = 50000004
	ShoemakerLevy9S  Code = 50000005
	ShoemakerLevy9R  Code = 50000006
	ShoemakerLevy9Q  Code = 50000007
	ShoemakerLevy9P  Code = 50000008
	ShoemakerLevy9N  Code = 50000009
	ShoemakerLevy9M  Code = 50000010
	ShoemakerLevy9L  Code = 50000011
	ShoemakerLevy9K  Code = 50000012
	ShoemakerLevy9J  Code = 50000013
	ShoemakerLevy9H  Code = 50000014
	ShoemakerLevy9G  Code = 50000015
	ShoemakerLevy9F  Code = 50000016
	ShoemakerLevy9E  Code = 50000017
	ShoemakerLevy9D  Code = 50000018
	ShoemakerLevy9C  Code = 50000019
	ShoemakerLevy9B  Code = 50000020
	ShoemakerLevy9A  Code = 50000021
	ShoemakerLevy9Q1 Code = 50000022
	ShoemakerLevy9P2 Code = 50000023
)
