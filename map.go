package main

var (
	toneMap = map[string][2]float64{
		"20":  {910, 1344},   // Ashford #20
		"71":  {472.5, 1530}, // Eastford #71
		"92":  {2465, 1530},  // Brooklyn - East Brooklyn #92
		"90":  {2465, 871},   // Brooklyn - Mortlake #90
		"590": {2465, 1800},  // Brooklyn - Mortlake Ambulance #90
		"93":  {2465, 950},   // Canterbury #93
		"09":  {2807, 910},   // Chaplin #09
		"12":  {2807, 950},   // Hampton #12
		"516": {2807, 2361},  // Hampton/Chaplin Ambulance
		"62":  {1180, 871},   // Killingly - Attawaugan #62
		"61":  {1180, 1901},  // Killingly - Danielson #61
		"63":  {1180, 799},   // Killingly - Dayville #63
		"64":  {1180, 992},   // Killingly - East Killingly #64
		"561": {1180, 1598},  // Killingly - KB Ambulance
		"65":  {1180, 2465},  // Killingly - South Killingly #65
		"60":  {1180, 732},   // Killingly - Williamsville #60
		"68":  {1180, 435.3}, // Oneco #68
		"95":  {2465, 992},   // Plainfield #95
		"96":  {2465, 1232},  // Plainfield - Atwood Hose #96
		"97":  {2465, 399.2}, // Plainfield - Central Village #97
		"94":  {2465, 435.3}, // Plainfield - Moosup #94
		"594": {2465, 701},   // Plainfield - Moosup Ambulance #94
		"70":  {2465, 731},   // Pomfret #70
		"78":  {472.5, 990},  // Putnam #78
		"79":  {1598, 1669},  // Putnam - East Putnam #79
		"961": {2465, 1901},  // QV Medic
		"16":  {950, 1232},   // Scotland #16
		"67":  {1180, 399.2}, // Sterling #67
		"81":  {1598, 399},   // Thompson - Community #81
		"581": {1598, 435.3}, // Thompson - Community Ambulance
		"85":  {1598, 799},   // Thompson - East Thompson #85
		"83":  {1598, 701},   // Thompson - Quinebaug #83
		"84":  {1598, 871},   // Thompson - Thompson Hill #84
		"82":  {1598, 731},   // Thompson - West Thompson #82
		"76":  {474.8, 799},  // Woodstock #76
		"77":  {474.8, 1820}, // Woodstock - Bungay #77
		"576": {474.8, 950},  // Woodstock - Woodstock Ambulance #76
		"75":  {474.8, 1598}, // Woodstock - Muddy Brook #75
	}
)
