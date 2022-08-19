package fpl_test

var BootstrapStaticPayload = `
{
	"events": [
		{
			"id": 22,
			"name": "Gameweek 22",
			"deadline_time": "2022-01-14T18:30:00Z",
			"average_entry_score": 51,
			"finished": true,
			"data_checked": true,
			"highest_scoring_entry": 3532161,
			"deadline_time_epoch": 1642185000,
			"deadline_time_game_offset": 0,
			"highest_score": 145,
			"is_previous": true,
			"is_current": false,
			"is_next": false,
			"cup_leagues_created": true,
			"h2h_ko_matches_created": true,
			"chip_plays": [
				{
				  "chip_name": "bboost",
				  "num_played": 57393
				},
				{
				  "chip_name": "freehit",
				  "num_played": 677118
				},
				{
				  "chip_name": "wildcard",
				  "num_played": 585135
				},
				{
				  "chip_name": "3xc",
				  "num_played": 56338
				}
			],
			"most_selected": 237,
			"most_transferred_in": 240,
			"top_element": 277,
			"top_element_info": {
				"id": 277,
				"points": 23
			},
			"transfers_made": 23424094,
			"most_captained": 233,
			"most_vice_captained": 579
		},
		{
			"id": 23,
			"name": "Gameweek 23",
			"deadline_time": "2022-01-21T18:30:00Z",
			"average_entry_score": 30,
			"finished": false,
			"data_checked": false,
			"highest_scoring_entry": 6083359,
			"deadline_time_epoch": 1642789800,
			"deadline_time_game_offset": 0,
			"highest_score": 100,
			"is_previous": false,
			"is_current": true,
			"is_next": false,
			"cup_leagues_created": true,
			"h2h_ko_matches_created": true,
			"chip_plays": [
				{
				  "chip_name": "bboost",
				  "num_played": 56752
				},
				{
				  "chip_name": "freehit",
				  "num_played": 78202
				},
				{
				  "chip_name": "wildcard",
				  "num_played": 259651
				},
				{
				  "chip_name": "3xc",
				  "num_played": 36867
				}
			],
			"most_selected": 237,
			"most_transferred_in": 681,
			"top_element": 426,
			"top_element_info": {
				"id": 426,
				"points": 13
			},
			"transfers_made": 12666677,
			"most_captained": 233,
			"most_vice_captained": 579
		}
	],
	"teams": [
		{
			"code": 3,
			"draw": 0,
			"form": null,
			"id": 1,
			"loss": 0,
			"name": "Arsenal",
			"played": 0,
			"points": 0,
			"position": 0,
			"short_name": "ARS",
			"strength": 4,
			"team_division": null,
			"unavailable": false,
			"win": 0,
			"strength_overall_home": 1190,
			"strength_overall_away": 1250,
			"strength_attack_home": 1110,
			"strength_attack_away": 1140,
			"strength_defence_home": 1110,
			"strength_defence_away": 1170,
			"pulse_id": 1
		},
		{
			"code": 7,
			"draw": 0,
			"form": null,
			"id": 2,
			"loss": 0,
			"name": "Aston Villa",
			"played": 0,
			"points": 0,
			"position": 0,
			"short_name": "AVL",
			"strength": 3,
			"team_division": null,
			"unavailable": false,
			"win": 0,
			"strength_overall_home": 1130,
			"strength_overall_away": 1160,
			"strength_attack_home": 1110,
			"strength_attack_away": 1120,
			"strength_defence_home": 1130,
			"strength_defence_away": 1160,
			"pulse_id": 2
		}
	]
}
`

var FixtureResponsePayload = `[
	{
		"code": 2210485,
		"event": 22,
		"finished": true,
		"finished_provisional": true,
		"id": 215,
		"kickoff_time": "2022-01-15T12:30:00Z",
		"minutes": 90,
		"provisional_start_time": false,
		"started": true,
		"team_a": 6,
		"team_a_score": 0,
		"team_h": 12,
		"team_h_score": 1,
		"stats": [
		  {
			"identifier": "goals_scored",
			"a": [],
			"h": [
			  {
				"value": 1,
				"element": 251
			  }
			]
		  },
		  {
			"identifier": "assists",
			"a": [],
			"h": [
			  {
				"value": 1,
				"element": 256
			  }
			]
		  },
		  {
			"identifier": "own_goals",
			"a": [],
			"h": []
		  },
		  {
			"identifier": "penalties_saved",
			"a": [],
			"h": []
		  },
		  {
			"identifier": "penalties_missed",
			"a": [],
			"h": []
		  },
		  {
			"identifier": "yellow_cards",
			"a": [
			  {
				"value": 1,
				"element": 122
			  },
			  {
				"value": 1,
				"element": 125
			  }
			],
			"h": []
		  },
		  {
			"identifier": "red_cards",
			"a": [],
			"h": []
		  },
		  {
			"identifier": "saves",
			"a": [
			  {
				"value": 5,
				"element": 129
			  }
			],
			"h": [
			  {
				"value": 1,
				"element": 257
			  }
			]
		  },
		  {
			"identifier": "bonus",
			"a": [],
			"h": [
			  {
				"value": 3,
				"element": 251
			  },
			  {
				"value": 2,
				"element": 256
			  },
			  {
				"value": 1,
				"element": 252
			  },
			  {
				"value": 1,
				"element": 259
			  }
			]
		  },
		  {
			"identifier": "bps",
			"a": [
			  {
				"value": 21,
				"element": 501
			  },
			  {
				"value": 20,
				"element": 129
			  },
			  {
				"value": 18,
				"element": 119
			  },
			  {
				"value": 17,
				"element": 125
			  },
			  {
				"value": 15,
				"element": 121
			  },
			  {
				"value": 13,
				"element": 127
			  },
			  {
				"value": 13,
				"element": 130
			  },
			  {
				"value": 9,
				"element": 122
			  },
			  {
				"value": 7,
				"element": 137
			  },
			  {
				"value": 4,
				"element": 131
			  },
			  {
				"value": 4,
				"element": 138
			  },
			  {
				"value": 3,
				"element": 140
			  },
			  {
				"value": 3,
				"element": 529
			  },
			  {
				"value": 2,
				"element": 134
			  }
			],
			"h": [
			  {
				"value": 34,
				"element": 251
			  },
			  {
				"value": 30,
				"element": 256
			  },
			  {
				"value": 27,
				"element": 252
			  },
			  {
				"value": 27,
				"element": 259
			  },
			  {
				"value": 22,
				"element": 257
			  },
			  {
				"value": 21,
				"element": 249
			  },
			  {
				"value": 14,
				"element": 266
			  },
			  {
				"value": 11,
				"element": 265
			  },
			  {
				"value": 8,
				"element": 33
			  },
			  {
				"value": 8,
				"element": 261
			  },
			  {
				"value": 5,
				"element": 250
			  },
			  {
				"value": 3,
				"element": 263
			  }
			]
		  }
		],
		"team_h_difficulty": 4,
		"team_a_difficulty": 5,
		"pulse_id": 66556
	}, 
	{
		"code": 2210486,
		"event": 22,
		"finished": false,
		"finished_provisional": true,
		"id": 216,
		"kickoff_time": "2022-01-15T15:00:00Z",
		"minutes": 70,
		"provisional_start_time": false,
		"started": false,
		"team_a": 18,
		"team_a_score": 2,
		"team_h": 14,
		"team_h_score": 3,
		"stats": [
		  {
			"identifier": "goals_scored",
			"a": [
			  {
				"value": 1,
				"element": 404
			  }
			],
			"h": [
			  {
				"value": 1,
				"element": 307
			  }
			]
		  },
		  {
			"identifier": "assists",
			"a": [
			  {
				"value": 1,
				"element": 382
			  }
			],
			"h": []
		  },
		  {
			"identifier": "own_goals",
			"a": [],
			"h": []
		  },
		  {
			"identifier": "penalties_saved",
			"a": [],
			"h": []
		  },
		  {
			"identifier": "penalties_missed",
			"a": [],
			"h": []
		  },
		  {
			"identifier": "yellow_cards",
			"a": [
			  {
				"value": 1,
				"element": 354
			  },
			  {
				"value": 1,
				"element": 677
			  }
			],
			"h": [
			  {
				"value": 1,
				"element": 291
			  },
			  {
				"value": 1,
				"element": 298
			  },
			  {
				"value": 1,
				"element": 301
			  }
			]
		  },
		  {
			"identifier": "red_cards",
			"a": [],
			"h": []
		  },
		  {
			"identifier": "saves",
			"a": [],
			"h": [
			  {
				"value": 4,
				"element": 295
			  }
			]
		  },
		  {
			"identifier": "bonus",
			"a": [
			  {
				"value": 2,
				"element": 404
			  },
			  {
				"value": 1,
				"element": 679
			  }
			],
			"h": [
			  {
				"value": 3,
				"element": 307
			  }
			]
		  },
		  {
			"identifier": "bps",
			"a": [
			  {
				"value": 24,
				"element": 404
			  },
			  {
				"value": 21,
				"element": 679
			  },
			  {
				"value": 18,
				"element": 379
			  },
			  {
				"value": 18,
				"element": 676
			  },
			  {
				"value": 14,
				"element": 382
			  },
			  {
				"value": 11,
				"element": 450
			  },
			  {
				"value": 11,
				"element": 677
			  },
			  {
				"value": 8,
				"element": 400
			  },
			  {
				"value": 8,
				"element": 495
			  },
			  {
				"value": 7,
				"element": 376
			  },
			  {
				"value": 3,
				"element": 472
			  },
			  {
				"value": 2,
				"element": 381
			  },
			  {
				"value": -1,
				"element": 468
			  }
			],
			"h": [
			  {
				"value": 27,
				"element": 307
			  },
			  {
				"value": 17,
				"element": 295
			  },
			  {
				"value": 13,
				"element": 309
			  },
			  {
				"value": 12,
				"element": 678
			  },
			  {
				"value": 10,
				"element": 299
			  },
			  {
				"value": 9,
				"element": 298
			  },
			  {
				"value": 9,
				"element": 301
			  },
			  {
				"value": 9,
				"element": 305
			  },
			  {
				"value": 8,
				"element": 310
			  },
			  {
				"value": 6,
				"element": 291
			  },
			  {
				"value": 3,
				"element": 304
			  },
			  {
				"value": 3,
				"element": 308
			  }
			]
		  }
		],
		"team_h_difficulty": 2,
		"team_a_difficulty": 2,
		"pulse_id": 66557
	}
]`
