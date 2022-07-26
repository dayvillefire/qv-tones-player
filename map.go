package main

var (
	toneMap = map[string][2]float64{
		"20":  {910, 1344},   // Ashford #20
		"71":  {472.5, 1530}, // Eastford #71
		"34":  {1082, 910},   // Bolton #34
		"26":  {2361, 2688},  // Bozrah #26
		"92":  {2465, 1530},  // Brooklyn - East Brooklyn #92
		"90":  {2465, 871},   // Brooklyn - Mortlake #90
		"590": {2465, 1800},  // Brooklyn - Mortlake Ambulance #90
		"93":  {2465, 950},   // Canterbury #93
		"09":  {2807, 910},   // Chaplin #09
		"05":  {1344, 992},   // Columbia #05
		"43":  {1130, 1082},  // Ellington #43
		"25":  {2807, 1985},  // Franklin #25
		"55":  {2361, 910},   // Griswold #55
		"56":  {2361, 765},   // Griswold - Jewett City #56
		"12":  {2807, 950},   // Hampton #12
		"516": {2807, 2361},  // Hampton/Chaplin Ambulance
		"10":  {1082, 2807},  // Hebron #10
		"62":  {1180, 871},   // Killingly - Attawaugan #62
		"61":  {1180, 1901},  // Killingly - Danielson #61
		"63":  {1180, 799},   // Killingly - Dayville #63
		"64":  {1180, 992},   // Killingly - East Killingly #64
		"561": {1180, 1598},  // Killingly - KB Ambulance
		"65":  {1180, 2465},  // Killingly - South Killingly #65
		"60":  {1180, 732},   // Killingly - Williamsville #60
		"06":  {2049, 1344},  // Lebanon #06
		"54":  {2361, 643},   // Lisbon #54
		"07":  {992, 2807},   // Mansfield #07
		"M":   {1180, 335.6}, // Mashantucket Pequot (Foxwoods)
		"NS":  {1530 - 910},  // North Stonington
		"59":  {1465, 950},   // Norwich - East Great Plain
		"52":  {1465, 992},   // Norwich - Laurel Hill #52
		"35":  {1465, 672},   // Norwich - Taftville #35
		"36":  {1465, 732},   // Norwich - Yantic #36
		"68":  {1180, 435.3}, // Oneco #68
		"95":  {2465, 992},   // Plainfield #95
		"96":  {2465, 1232},  // Plainfield - Atwood Hose #96
		"97":  {2465, 399.2}, // Plainfield - Central Village #97
		"94":  {2465, 435.3}, // Plainfield - Moosup #94
		"594": {2465, 701},   // Plainfield - Moosup Ambulance #94
		"70":  {2465, 731},   // Pomfret #70
		"57":  {2361, 1232},  // Preston #57
		"78":  {472.5, 990},  // Putnam #78
		"578": {472.5, 871},  // Putnam - Putnam EMS #578
		"79":  {1598, 1669},  // Putnam - East Putnam #79
		"961": {2465, 1901},  // QV Medic
		"21":  {2575, 2361},  // Salem #21
		"16":  {950, 1232},   // Scotland #16
		"46":  {799, 2043.8}, // Somers #46
		"24":  {2361, 1695},  // Sprague - Baltic #24
		"67":  {1180, 399.2}, // Sterling #67
		"81":  {1598, 399},   // Thompson - Community #81
		"581": {1598, 435.3}, // Thompson - Community Ambulance
		"85":  {1598, 799},   // Thompson - East Thompson #85
		"83":  {1598, 701},   // Thompson - Quinebaug #83
		"84":  {1598, 871},   // Thompson - Thompson Hill #84
		"82":  {1598, 731},   // Thompson - West Thompson #82
		"72":  {2575, 1344},  // Union #72
		"41":  {992, 910},    // Vernon #41
		"53":  {2361, 992},   // Voluntown #53
		"501": {871, 1130},   // Windham - Windham Hospital Medics
		"02":  {950, 1130},   // Windham - North Windham #02
		"03":  {2932, 1130},  // Windham - Windham Center #03
		"76":  {474.8, 799},  // Woodstock #76
		"77":  {474.8, 1820}, // Woodstock - Bungay #77
		"576": {474.8, 950},  // Woodstock - Woodstock Ambulance #76
		"75":  {474.8, 1598}, // Woodstock - Muddy Brook #75
	}
)
