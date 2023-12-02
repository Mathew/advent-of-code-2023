package main

/*
:%s/\(\w\+\)/Game{ \r\tGameId:/
:%s/\(\d\+\)/\1,\r\tRounds: []map[string]int{\r\t\t
:%s/\(\d\+\) \(\w\+\)/\2: \1/g
:%s/\(\t\: \)/\t{\r\t\t\t
:%s/;/, \r\t\t},{\r\t\t\t/g
:%s/\(\d\)$/\1,\r\t\t},\r\t},
:%s/\(red\|blue\|green\)/"\1"/g
:%s/Game{/},Game {
*/

var INPUT = []Game{
	Game{
		GameId: 1,
		Rounds: []map[string]int{
			{
				"red": 7, "blue": 14,
			}, {
				"blue": 2, "red": 3, "green": 3,
			}, {
				"green": 4, "blue": 12, "red": 15,
			}, {
				"green": 3, "blue": 12, "red": 3,
			}, {
				"red": 11, "green": 2,
			},
		},
	}, Game{
		GameId: 2,
		Rounds: []map[string]int{
			{
				"blue": 16, "red": 9, "green": 5,
			}, {
				"red": 8,
			}, {
				"blue": 8, "green": 5, "red": 12,
			}, {
				"blue": 11, "green": 8, "red": 17,
			},
		},
	}, Game{
		GameId: 3,
		Rounds: []map[string]int{
			{
				"green": 8, "blue": 1, "red": 7,
			}, {
				"red": 12, "blue": 6, "green": 9,
			}, {
				"blue": 2, "red": 1, "green": 14,
			}, {
				"green": 9, "red": 4,
			}, {
				"red": 2, "blue": 1, "green": 8,
			},
		},
	}, Game{
		GameId: 4,
		Rounds: []map[string]int{
			{
				"blue": 1, "green": 3,
			}, {
				"green": 2, "blue": 1, "red": 1,
			}, {
				"red": 1, "green": 3,
			},
		},
	}, Game{
		GameId: 5,
		Rounds: []map[string]int{
			{
				"red": 6, "blue": 1,
			}, {
				"green": 1,
			}, {
				"red": 5, "green": 2,
			}, {
				"red": 1, "blue": 1, "green": 3,
			},
		},
	}, Game{
		GameId: 6,
		Rounds: []map[string]int{
			{
				"green": 3, "red": 4, "blue": 1,
			}, {
				"blue": 2, "green": 5, "red": 2,
			}, {
				"green": 12, "blue": 3, "red": 2,
			}, {
				"blue": 4, "green": 1, "red": 4,
			}, {
				"green": 11, "red": 6,
			}, {
				"green": 5, "red": 10, "blue": 3,
			},
		},
	}, Game{
		GameId: 7,
		Rounds: []map[string]int{
			{
				"blue": 2, "green": 3, "red": 16,
			}, {
				"blue": 1, "red": 3,
			}, {
				"green": 2, "red": 13,
			}, {
				"red": 18, "blue": 2, "green": 1,
			}, {
				"red": 3, "blue": 1,
			},
		},
	}, Game{
		GameId: 8,
		Rounds: []map[string]int{
			{
				"red": 4, "blue": 3, "green": 8,
			}, {
				"red": 2, "green": 16,
			}, {
				"red": 2, "blue": 1,
			},
		},
	}, Game{
		GameId: 9,
		Rounds: []map[string]int{
			{
				"green": 4, "blue": 14, "red": 8,
			}, {
				"blue": 17, "red": 3, "green": 5,
			}, {
				"green": 2, "red": 4, "blue": 6,
			}, {
				"red": 7, "green": 2, "blue": 18,
			}, {
				"red": 3, "blue": 19, "green": 4,
			}, {
				"green": 4, "red": 8, "blue": 6,
			},
		},
	}, Game{
		GameId: 10,
		Rounds: []map[string]int{
			{
				"green": 12, "red": 7, "blue": 1,
			}, {
				"red": 6, "green": 12,
			}, {
				"red": 6, "green": 7, "blue": 1,
			}, {
				"red": 1, "blue": 1, "green": 18,
			}, {
				"green": 11, "blue": 1,
			},
		},
	}, Game{
		GameId: 11,
		Rounds: []map[string]int{
			{
				"green": 10, "red": 3, "blue": 13,
			}, {
				"blue": 13, "green": 6, "red": 8,
			}, {
				"blue": 12, "green": 4, "red": 8,
			}, {
				"green": 9, "red": 9, "blue": 3,
			}, {
				"blue": 6, "green": 7, "red": 6,
			}, {
				"blue": 11, "green": 13,
			},
		},
	}, Game{
		GameId: 12,
		Rounds: []map[string]int{
			{
				"green": 9, "blue": 2,
			}, {
				"green": 4, "blue": 1, "red": 7,
			}, {
				"green": 2, "blue": 1, "red": 5,
			},
		},
	}, Game{
		GameId: 13,
		Rounds: []map[string]int{
			{
				"green": 1,
			}, {
				"blue": 7, "red": 1, "green": 2,
			}, {
				"blue": 8, "green": 2,
			},
		},
	}, Game{
		GameId: 14,
		Rounds: []map[string]int{
			{
				"red": 8, "green": 3,
			}, {
				"red": 1, "green": 8,
			}, {
				"blue": 1, "green": 10,
			},
		},
	}, Game{
		GameId: 15,
		Rounds: []map[string]int{
			{
				"blue": 1, "green": 6, "red": 14,
			}, {
				"red": 3, "blue": 1, "green": 6,
			}, {
				"green": 4,
			}, {
				"blue": 1, "green": 5, "red": 2,
			}, {
				"blue": 2, "green": 1, "red": 6,
			}, {
				"red": 4, "green": 8, "blue": 1,
			},
		},
	}, Game{
		GameId: 16,
		Rounds: []map[string]int{
			{
				"green": 1, "red": 6, "blue": 8,
			}, {
				"green": 1, "blue": 11, "red": 1,
			}, {
				"blue": 7, "green": 3, "red": 4,
			}, {
				"green": 2, "red": 6, "blue": 12,
			},
		},
	}, Game{
		GameId: 17,
		Rounds: []map[string]int{
			{
				"blue": 3, "red": 4, "green": 4,
			}, {
				"blue": 18, "red": 6, "green": 11,
			}, {
				"green": 2, "red": 6, "blue": 17,
			}, {
				"green": 12, "blue": 3, "red": 5,
			},
		},
	}, Game{
		GameId: 18,
		Rounds: []map[string]int{
			{
				"green": 3, "blue": 2, "red": 10,
			}, {
				"green": 7, "blue": 10,
			}, {
				"blue": 9, "red": 7, "green": 14,
			}, {
				"green": 18, "blue": 10, "red": 11,
			}, {
				"red": 10, "blue": 2, "green": 16,
			},
		},
	}, Game{
		GameId: 19,
		Rounds: []map[string]int{
			{
				"red": 6, "green": 1, "blue": 18,
			}, {
				"red": 2, "blue": 1,
			}, {
				"blue": 7, "red": 3, "green": 2,
			}, {
				"blue": 18, "green": 2, "red": 1,
			}, {
				"red": 7, "blue": 10,
			},
		},
	}, Game{
		GameId: 20,
		Rounds: []map[string]int{
			{
				"blue": 13, "red": 2,
			}, {
				"green": 2, "red": 2,
			}, {
				"green": 1, "blue": 9,
			},
		},
	}, Game{
		GameId: 21,
		Rounds: []map[string]int{
			{
				"blue": 4, "red": 1,
			}, {
				"red": 2, "blue": 4, "green": 1,
			}, {
				"red": 3,
			}, {
				"green": 4, "red": 1, "blue": 1,
			}, {
				"green": 3, "blue": 9, "red": 1,
			},
		},
	}, Game{
		GameId: 22,
		Rounds: []map[string]int{
			{
				"blue": 7, "green": 5, "red": 14,
			}, {
				"red": 15, "blue": 9, "green": 11,
			}, {
				"blue": 10, "red": 5, "green": 11,
			}, {
				"red": 14, "blue": 10, "green": 13,
			},
		},
	}, Game{
		GameId: 23,
		Rounds: []map[string]int{
			{
				"red": 10, "blue": 6,
			}, {
				"red": 1, "blue": 4, "green": 3,
			}, {
				"green": 3, "blue": 2,
			}, {
				"red": 5, "green": 3,
			}, {
				"green": 3, "blue": 4, "red": 5,
			}, {
				"green": 3, "red": 7, "blue": 6,
			},
		},
	}, Game{
		GameId: 24,
		Rounds: []map[string]int{
			{
				"red": 4, "green": 8,
			}, {
				"red": 1, "green": 10,
			}, {
				"red": 2, "green": 1,
			}, {
				"green": 2, "blue": 1,
			}, {
				"red": 4, "green": 12,
			}, {
				"green": 3,
			},
		},
	}, Game{
		GameId: 25,
		Rounds: []map[string]int{
			{
				"red": 5, "blue": 2, "green": 6,
			}, {
				"red": 4, "blue": 3, "green": 8,
			}, {
				"green": 11, "red": 4, "blue": 1,
			},
		},
	}, Game{
		GameId: 26,
		Rounds: []map[string]int{
			{
				"blue": 5, "red": 1,
			}, {
				"blue": 18, "green": 4,
			}, {
				"green": 9, "red": 3, "blue": 17,
			}, {
				"green": 6, "blue": 10, "red": 1,
			}, {
				"blue": 3, "green": 7,
			}, {
				"blue": 4, "red": 3, "green": 5,
			},
		},
	}, Game{
		GameId: 27,
		Rounds: []map[string]int{
			{
				"red": 13, "blue": 2,
			}, {
				"blue": 7, "green": 2, "red": 12,
			}, {
				"green": 1, "blue": 9, "red": 9,
			}, {
				"red": 4, "green": 4, "blue": 8,
			}, {
				"red": 13, "blue": 6,
			}, {
				"red": 3, "blue": 9, "green": 3,
			},
		},
	}, Game{
		GameId: 28,
		Rounds: []map[string]int{
			{
				"blue": 1, "green": 12, "red": 1,
			}, {
				"blue": 1, "green": 12, "red": 2,
			}, {
				"red": 2, "green": 8, "blue": 1,
			}, {
				"green": 5, "red": 2,
			}, {
				"blue": 1, "green": 9, "red": 6,
			}, {
				"blue": 1, "green": 13,
			},
		},
	}, Game{
		GameId: 29,
		Rounds: []map[string]int{
			{
				"blue": 5, "red": 5, "green": 11,
			}, {
				"blue": 15, "red": 5, "green": 10,
			}, {
				"red": 2, "green": 11, "blue": 19,
			}, {
				"blue": 19, "green": 3, "red": 6,
			},
		},
	}, Game{
		GameId: 30,
		Rounds: []map[string]int{
			{
				"blue": 1, "red": 12, "green": 1,
			}, {
				"blue": 12, "red": 1, "green": 2,
			}, {
				"red": 12, "green": 5,
			}, {
				"red": 2, "green": 2, "blue": 5,
			}, {
				"red": 5, "green": 2, "blue": 6,
			},
		},
	}, Game{
		GameId: 31,
		Rounds: []map[string]int{
			{
				"green": 20, "red": 1, "blue": 16,
			}, {
				"green": 3, "red": 1, "blue": 7,
			}, {
				"red": 6, "blue": 18, "green": 8,
			},
		},
	}, Game{
		GameId: 32,
		Rounds: []map[string]int{
			{
				"green": 5,
			}, {
				"blue": 1, "red": 2, "green": 5,
			}, {
				"blue": 1, "red": 2, "green": 5,
			}, {
				"green": 2, "red": 2,
			},
		},
	}, Game{
		GameId: 33,
		Rounds: []map[string]int{
			{
				"blue": 6, "green": 5,
			}, {
				"blue": 6, "red": 3, "green": 1,
			}, {
				"green": 4, "blue": 3, "red": 2,
			}, {
				"red": 1, "blue": 6, "green": 5,
			}, {
				"blue": 1, "red": 2, "green": 5,
			}, {
				"red": 4, "blue": 3,
			},
		},
	}, Game{
		GameId: 34,
		Rounds: []map[string]int{
			{
				"red": 12, "green": 12,
			}, {
				"red": 12, "green": 7,
			}, {
				"blue": 1, "red": 12, "green": 11,
			}, {
				"red": 7, "green": 7, "blue": 2,
			},
		},
	}, Game{
		GameId: 35,
		Rounds: []map[string]int{
			{
				"red": 3, "blue": 3, "green": 1,
			}, {
				"green": 1, "red": 5, "blue": 5,
			}, {
				"green": 8, "red": 2, "blue": 14,
			}, {
				"green": 8,
			}, {
				"blue": 6, "red": 3, "green": 6,
			}, {
				"red": 1, "blue": 1, "green": 12,
			},
		},
	}, Game{
		GameId: 36,
		Rounds: []map[string]int{
			{
				"red": 13, "blue": 5,
			}, {
				"blue": 13, "green": 10, "red": 6,
			}, {
				"red": 10, "green": 5, "blue": 10,
			}, {
				"blue": 10, "green": 8, "red": 6,
			}, {
				"green": 1, "red": 14, "blue": 2,
			}, {
				"green": 8, "blue": 4,
			},
		},
	}, Game{
		GameId: 37,
		Rounds: []map[string]int{
			{
				"green": 6, "red": 1,
			}, {
				"blue": 1, "green": 4, "red": 1,
			}, {
				"red": 1, "green": 14,
			}, {
				"red": 1, "green": 9,
			}, {
				"green": 1, "red": 1,
			}, {
				"green": 9, "red": 1,
			},
		},
	}, Game{
		GameId: 38,
		Rounds: []map[string]int{
			{
				"green": 1, "blue": 4, "red": 17,
			}, {
				"red": 13, "blue": 9, "green": 12,
			}, {
				"green": 7, "blue": 11,
			},
		},
	}, Game{
		GameId: 39,
		Rounds: []map[string]int{
			{
				"green": 18, "blue": 9, "red": 2,
			}, {
				"red": 11, "blue": 1, "green": 4,
			}, {
				"red": 9, "green": 4, "blue": 10,
			}, {
				"blue": 9, "red": 5, "green": 2,
			},
		},
	}, Game{
		GameId: 40,
		Rounds: []map[string]int{
			{
				"green": 3, "red": 8,
			}, {
				"green": 2, "red": 6,
			}, {
				"green": 1, "red": 9, "blue": 4,
			}, {
				"blue": 1, "red": 6,
			}, {
				"green": 2, "blue": 2, "red": 3,
			},
		},
	}, Game{
		GameId: 41,
		Rounds: []map[string]int{
			{
				"red": 3, "green": 15, "blue": 3,
			}, {
				"green": 19, "red": 2, "blue": 5,
			}, {
				"blue": 8, "green": 7, "red": 4,
			}, {
				"blue": 3, "red": 4, "green": 5,
			}, {
				"blue": 1,
			}, {
				"blue": 6, "green": 15, "red": 3,
			},
		},
	}, Game{
		GameId: 42,
		Rounds: []map[string]int{
			{
				"red": 2, "blue": 18, "green": 6,
			}, {
				"green": 3, "blue": 2, "red": 8,
			}, {
				"blue": 9, "green": 1, "red": 5,
			}, {
				"red": 12, "blue": 3, "green": 8,
			},
		},
	}, Game{
		GameId: 43,
		Rounds: []map[string]int{
			{
				"blue": 3, "green": 1, "red": 3,
			}, {
				"blue": 8, "green": 3, "red": 1,
			}, {
				"red": 3, "blue": 5,
			}, {
				"green": 3, "red": 3, "blue": 7,
			}, {
				"blue": 6, "green": 1, "red": 2,
			}, {
				"blue": 7, "green": 2, "red": 5,
			},
		},
	}, Game{
		GameId: 44,
		Rounds: []map[string]int{
			{
				"green": 2, "blue": 5, "red": 1,
			}, {
				"red": 9, "blue": 16,
			}, {
				"blue": 4, "green": 2, "red": 12,
			}, {
				"red": 13, "blue": 5, "green": 10,
			}, {
				"green": 4, "blue": 3, "red": 11,
			},
		},
	}, Game{
		GameId: 45,
		Rounds: []map[string]int{
			{
				"blue": 6, "red": 3, "green": 13,
			}, {
				"green": 10, "blue": 13, "red": 12,
			}, {
				"red": 7, "blue": 19, "green": 16,
			}, {
				"blue": 15, "red": 4, "green": 11,
			}, {
				"red": 1, "green": 4,
			},
		},
	}, Game{
		GameId: 46,
		Rounds: []map[string]int{
			{
				"red": 11, "green": 2,
			}, {
				"blue": 5, "red": 2, "green": 2,
			}, {
				"green": 3, "red": 8, "blue": 6,
			}, {
				"blue": 3, "green": 10, "red": 8,
			},
		},
	}, Game{
		GameId: 47,
		Rounds: []map[string]int{
			{
				"green": 6, "red": 16,
			}, {
				"blue": 4, "red": 4, "green": 2,
			}, {
				"blue": 3, "green": 1, "red": 12,
			}, {
				"red": 2, "blue": 4, "green": 4,
			}, {
				"blue": 2, "red": 16,
			}, {
				"blue": 5, "green": 5, "red": 5,
			},
		},
	}, Game{
		GameId: 48,
		Rounds: []map[string]int{
			{
				"red": 8, "blue": 1,
			}, {
				"green": 1, "blue": 2, "red": 6,
			}, {
				"red": 11, "green": 6, "blue": 2,
			},
		},
	}, Game{
		GameId: 49,
		Rounds: []map[string]int{
			{
				"green": 5, "blue": 16, "red": 2,
			}, {
				"red": 2, "blue": 20, "green": 6,
			}, {
				"red": 1, "blue": 3, "green": 5,
			}, {
				"green": 7, "blue": 4,
			},
		},
	}, Game{
		GameId: 50,
		Rounds: []map[string]int{
			{
				"red": 9, "green": 8,
			}, {
				"green": 11, "red": 10, "blue": 1,
			}, {
				"red": 9, "green": 5,
			}, {
				"blue": 1, "green": 12, "red": 8,
			}, {
				"blue": 1, "green": 5, "red": 2,
			}, {
				"green": 6, "blue": 1, "red": 2,
			},
		},
	}, Game{
		GameId: 51,
		Rounds: []map[string]int{
			{
				"red": 1, "blue": 4,
			}, {
				"red": 1, "green": 3, "blue": 3,
			}, {
				"green": 1, "red": 1, "blue": 2,
			},
		},
	}, Game{
		GameId: 52,
		Rounds: []map[string]int{
			{
				"red": 11, "blue": 4,
			}, {
				"green": 1, "blue": 6, "red": 10,
			}, {
				"blue": 8, "red": 13,
			},
		},
	}, Game{
		GameId: 53,
		Rounds: []map[string]int{
			{
				"green": 6, "red": 9,
			}, {
				"blue": 4, "red": 13, "green": 2,
			}, {
				"red": 10, "green": 5, "blue": 3,
			}, {
				"red": 11, "blue": 3, "green": 4,
			},
		},
	}, Game{
		GameId: 54,
		Rounds: []map[string]int{
			{
				"red": 16, "blue": 9, "green": 8,
			}, {
				"red": 9, "blue": 1,
			}, {
				"green": 12, "red": 2, "blue": 13,
			}, {
				"blue": 5, "green": 14, "red": 15,
			}, {
				"green": 3, "red": 2, "blue": 2,
			},
		},
	}, Game{
		GameId: 55,
		Rounds: []map[string]int{
			{
				"green": 3, "blue": 4, "red": 5,
			}, {
				"red": 3, "green": 9, "blue": 1,
			}, {
				"blue": 3, "green": 4, "red": 5,
			}, {
				"green": 4, "blue": 3, "red": 7,
			}, {
				"red": 5, "blue": 2,
			}, {
				"blue": 2, "red": 8, "green": 5,
			},
		},
	}, Game{
		GameId: 56,
		Rounds: []map[string]int{
			{
				"red": 3, "green": 5, "blue": 3,
			}, {
				"red": 15, "green": 3, "blue": 15,
			}, {
				"green": 3, "blue": 1, "red": 10,
			}, {
				"blue": 15, "red": 1, "green": 2,
			}, {
				"red": 6, "blue": 16, "green": 6,
			}, {
				"red": 19, "blue": 16,
			},
		},
	}, Game{
		GameId: 57,
		Rounds: []map[string]int{
			{
				"blue": 5, "red": 1, "green": 5,
			}, {
				"blue": 8, "green": 16,
			}, {
				"green": 13, "blue": 5, "red": 3,
			}, {
				"blue": 1, "red": 1, "green": 13,
			}, {
				"green": 12, "red": 3, "blue": 4,
			}, {
				"blue": 8, "red": 3, "green": 1,
			},
		},
	}, Game{
		GameId: 58,
		Rounds: []map[string]int{
			{
				"blue": 5, "green": 4,
			}, {
				"blue": 7, "red": 1, "green": 10,
			}, {
				"red": 1, "green": 13, "blue": 4,
			}, {
				"blue": 7, "green": 12, "red": 4,
			}, {
				"red": 4, "green": 13, "blue": 5,
			}, {
				"green": 2, "blue": 1, "red": 12,
			},
		},
	}, Game{
		GameId: 59,
		Rounds: []map[string]int{
			{
				"red": 2, "blue": 11, "green": 6,
			}, {
				"green": 1, "blue": 8, "red": 3,
			}, {
				"red": 4, "blue": 6,
			},
		},
	}, Game{
		GameId: 60,
		Rounds: []map[string]int{
			{
				"green": 4, "red": 1,
			}, {
				"blue": 3, "red": 15, "green": 2,
			}, {
				"red": 13, "blue": 16, "green": 2,
			}, {
				"green": 6, "blue": 13, "red": 10,
			}, {
				"blue": 15, "red": 11, "green": 5,
			}, {
				"blue": 7, "green": 4,
			},
		},
	}, Game{
		GameId: 61,
		Rounds: []map[string]int{
			{
				"red": 14, "green": 2, "blue": 14,
			}, {
				"green": 1, "red": 15, "blue": 3,
			}, {
				"green": 2, "blue": 8,
			},
		},
	}, Game{
		GameId: 62,
		Rounds: []map[string]int{
			{
				"green": 13, "blue": 13,
			}, {
				"red": 1, "green": 6, "blue": 1,
			}, {
				"blue": 13, "green": 16,
			}, {
				"blue": 3, "red": 1, "green": 2,
			},
		},
	}, Game{
		GameId: 63,
		Rounds: []map[string]int{
			{
				"blue": 10, "red": 3, "green": 4,
			}, {
				"red": 15,
			}, {
				"blue": 10, "green": 10, "red": 14,
			}, {
				"blue": 9, "green": 6,
			}, {
				"blue": 3, "green": 7, "red": 13,
			},
		},
	}, Game{
		GameId: 64,
		Rounds: []map[string]int{
			{
				"red": 2, "green": 4,
			}, {
				"blue": 1, "red": 9,
			}, {
				"green": 1, "blue": 2, "red": 10,
			}, {
				"red": 9, "blue": 1, "green": 5,
			}, {
				"green": 6, "red": 6,
			},
		},
	}, Game{
		GameId: 65,
		Rounds: []map[string]int{
			{
				"blue": 10, "green": 4,
			}, {
				"green": 4, "red": 2, "blue": 9,
			}, {
				"red": 11, "green": 1, "blue": 10,
			}, {
				"green": 14, "blue": 19, "red": 3,
			}, {
				"red": 12, "blue": 5, "green": 11,
			}, {
				"blue": 14, "green": 12, "red": 11,
			},
		},
	}, Game{
		GameId: 66,
		Rounds: []map[string]int{
			{
				"blue": 5, "red": 2,
			}, {
				"blue": 5, "green": 1, "red": 7,
			}, {
				"red": 14, "green": 1, "blue": 2,
			}, {
				"red": 8, "blue": 7,
			}, {
				"red": 2, "blue": 4, "green": 1,
			}, {
				"blue": 2, "red": 18,
			},
		},
	}, Game{
		GameId: 67,
		Rounds: []map[string]int{
			{
				"red": 6, "blue": 1,
			}, {
				"green": 5, "blue": 2, "red": 1,
			}, {
				"red": 2, "green": 3, "blue": 3,
			}, {
				"green": 8, "blue": 4, "red": 6,
			},
		},
	}, Game{
		GameId: 68,
		Rounds: []map[string]int{
			{
				"blue": 4, "green": 1,
			}, {
				"blue": 12, "red": 2, "green": 3,
			}, {
				"green": 2, "blue": 7,
			}, {
				"red": 1, "blue": 19, "green": 3,
			},
		},
	}, Game{
		GameId: 69,
		Rounds: []map[string]int{
			{
				"green": 6, "red": 11, "blue": 2,
			}, {
				"blue": 1, "green": 7, "red": 6,
			}, {
				"blue": 1, "red": 8,
			}, {
				"red": 17, "blue": 3, "green": 5,
			},
		},
	}, Game{
		GameId: 70,
		Rounds: []map[string]int{
			{
				"green": 2, "red": 6, "blue": 4,
			}, {
				"green": 2, "red": 7, "blue": 1,
			}, {
				"blue": 11, "green": 1,
			},
		},
	}, Game{
		GameId: 71,
		Rounds: []map[string]int{
			{
				"blue": 10, "red": 9,
			}, {
				"red": 3, "blue": 10,
			}, {
				"red": 1, "blue": 8, "green": 2,
			}, {
				"blue": 7, "green": 4, "red": 5,
			}, {
				"red": 6, "blue": 2, "green": 7,
			}, {
				"red": 5, "blue": 2, "green": 4,
			},
		},
	}, Game{
		GameId: 72,
		Rounds: []map[string]int{
			{
				"green": 1, "blue": 12, "red": 8,
			}, {
				"red": 9, "blue": 3,
			}, {
				"red": 2, "green": 2, "blue": 10,
			},
		},
	}, Game{
		GameId: 73,
		Rounds: []map[string]int{
			{
				"red": 7, "green": 3, "blue": 11,
			}, {
				"green": 4, "blue": 7,
			}, {
				"blue": 6, "green": 13, "red": 9,
			}, {
				"green": 11, "blue": 4,
			}, {
				"blue": 12, "red": 3, "green": 2,
			}, {
				"green": 9,
			},
		},
	}, Game{
		GameId: 74,
		Rounds: []map[string]int{
			{
				"blue": 5, "red": 2,
			}, {
				"red": 6, "blue": 1, "green": 8,
			}, {
				"green": 6, "blue": 5, "red": 16,
			}, {
				"green": 1, "red": 9, "blue": 3,
			}, {
				"green": 12, "red": 1, "blue": 1,
			}, {
				"blue": 2, "green": 7, "red": 13,
			},
		},
	}, Game{
		GameId: 75,
		Rounds: []map[string]int{
			{
				"green": 5, "red": 20,
			}, {
				"red": 7, "green": 6, "blue": 2,
			}, {
				"green": 4, "blue": 2,
			}, {
				"blue": 2, "green": 1, "red": 3,
			}, {
				"blue": 2, "green": 2, "red": 12,
			}, {
				"red": 6, "green": 6,
			},
		},
	}, Game{
		GameId: 76,
		Rounds: []map[string]int{
			{
				"red": 9, "green": 12, "blue": 3,
			}, {
				"blue": 2, "red": 1, "green": 6,
			}, {
				"green": 13, "blue": 2,
			}, {
				"red": 2, "green": 7, "blue": 3,
			}, {
				"red": 7, "green": 4, "blue": 2,
			}, {
				"red": 2, "blue": 3, "green": 3,
			},
		},
	}, Game{
		GameId: 77,
		Rounds: []map[string]int{
			{
				"blue": 2, "red": 6,
			}, {
				"red": 4, "green": 15, "blue": 1,
			}, {
				"green": 7, "blue": 5, "red": 6,
			}, {
				"red": 4, "blue": 5,
			},
		},
	}, Game{
		GameId: 78,
		Rounds: []map[string]int{
			{
				"blue": 5, "red": 3, "green": 1,
			}, {
				"green": 2, "red": 7, "blue": 3,
			}, {
				"blue": 3, "red": 5, "green": 5,
			},
		},
	}, Game{
		GameId: 79,
		Rounds: []map[string]int{
			{
				"red": 6, "blue": 9, "green": 1,
			}, {
				"green": 9, "red": 8, "blue": 7,
			}, {
				"blue": 1, "green": 12, "red": 13,
			}, {
				"red": 7, "blue": 14, "green": 2,
			}, {
				"blue": 13, "green": 4, "red": 9,
			}, {
				"blue": 4, "green": 2,
			},
		},
	}, Game{
		GameId: 80,
		Rounds: []map[string]int{
			{
				"green": 4, "blue": 2,
			}, {
				"green": 5, "red": 3, "blue": 8,
			}, {
				"blue": 9, "red": 11, "green": 4,
			}, {
				"blue": 2, "green": 3, "red": 4,
			}, {
				"red": 5,
			},
		},
	}, Game{
		GameId: 81,
		Rounds: []map[string]int{
			{
				"red": 8, "blue": 3, "green": 4,
			}, {
				"blue": 13, "red": 8, "green": 1,
			}, {
				"blue": 6, "green": 1,
			}, {
				"green": 18, "red": 6, "blue": 10,
			}, {
				"green": 17, "blue": 8, "red": 3,
			}, {
				"red": 6, "green": 5, "blue": 12,
			},
		},
	}, Game{
		GameId: 82,
		Rounds: []map[string]int{
			{
				"red": 3, "blue": 7,
			}, {
				"red": 4, "blue": 6, "green": 14,
			}, {
				"blue": 9, "green": 2, "red": 3,
			},
		},
	}, Game{
		GameId: 83,
		Rounds: []map[string]int{
			{
				"blue": 1, "red": 2,
			}, {
				"green": 5, "red": 16,
			}, {
				"red": 12, "green": 1,
			}, {
				"green": 8, "red": 8,
			},
		},
	}, Game{
		GameId: 84,
		Rounds: []map[string]int{
			{
				"red": 3, "green": 9, "blue": 1,
			}, {
				"red": 3, "blue": 6, "green": 7,
			}, {
				"red": 5, "green": 8, "blue": 8,
			}, {
				"red": 5, "blue": 3, "green": 11,
			}, {
				"green": 3, "blue": 4,
			}, {
				"green": 4, "blue": 1, "red": 2,
			},
		},
	}, Game{
		GameId: 85,
		Rounds: []map[string]int{
			{
				"red": 4, "blue": 6, "green": 1,
			}, {
				"red": 7, "blue": 6,
			}, {
				"red": 9, "green": 1,
			}, {
				"blue": 1, "green": 1, "red": 10,
			}, {
				"red": 2, "blue": 2, "green": 1,
			}, {
				"blue": 5, "red": 7,
			},
		},
	}, Game{
		GameId: 86,
		Rounds: []map[string]int{
			{
				"blue": 4, "green": 5, "red": 6,
			}, {
				"red": 9, "blue": 3,
			}, {
				"green": 5, "red": 3, "blue": 10,
			}, {
				"green": 3, "blue": 7, "red": 3,
			}, {
				"red": 4,
			}, {
				"green": 4, "blue": 1, "red": 8,
			},
		},
	}, Game{
		GameId: 87,
		Rounds: []map[string]int{
			{
				"red": 3, "green": 3,
			}, {
				"blue": 3, "green": 1,
			}, {
				"red": 3, "green": 3,
			}, {
				"red": 3, "blue": 1, "green": 3,
			}, {
				"green": 2, "red": 1,
			},
		},
	}, Game{
		GameId: 88,
		Rounds: []map[string]int{
			{
				"red": 1, "green": 13, "blue": 3,
			}, {
				"blue": 17, "green": 14, "red": 5,
			}, {
				"red": 3, "blue": 19, "green": 13,
			}, {
				"green": 7, "blue": 19,
			}, {
				"red": 5, "green": 13, "blue": 17,
			}, {
				"blue": 13, "green": 8, "red": 2,
			},
		},
	}, Game{
		GameId: 89,
		Rounds: []map[string]int{
			{
				"blue": 3, "red": 4,
			}, {
				"green": 2, "red": 15, "blue": 1,
			}, {
				"green": 3, "blue": 3, "red": 13,
			}, {
				"blue": 3, "red": 9, "green": 2,
			}, {
				"red": 8,
			},
		},
	}, Game{
		GameId: 90,
		Rounds: []map[string]int{
			{
				"red": 2, "green": 2, "blue": 1,
			}, {
				"blue": 3, "green": 2,
			}, {
				"blue": 1, "green": 2, "red": 4,
			}, {
				"blue": 3,
			},
		},
	}, Game{
		GameId: 91,
		Rounds: []map[string]int{
			{
				"blue": 13, "green": 5, "red": 4,
			}, {
				"blue": 17, "red": 8, "green": 11,
			}, {
				"green": 1, "red": 6, "blue": 19,
			}, {
				"blue": 12, "green": 6,
			}, {
				"green": 7, "red": 2,
			},
		},
	}, Game{
		GameId: 92,
		Rounds: []map[string]int{
			{
				"red": 6, "green": 4,
			}, {
				"blue": 2, "red": 11,
			}, {
				"green": 4, "blue": 7,
			}, {
				"red": 2, "blue": 12, "green": 2,
			},
		},
	}, Game{
		GameId: 93,
		Rounds: []map[string]int{
			{
				"blue": 3, "red": 2,
			}, {
				"blue": 2, "red": 11, "green": 1,
			}, {
				"red": 7, "green": 1,
			}, {
				"red": 1, "blue": 2,
			}, {
				"red": 13, "blue": 3,
			},
		},
	}, Game{
		GameId: 94,
		Rounds: []map[string]int{
			{
				"blue": 2, "red": 1, "green": 20,
			}, {
				"red": 1, "blue": 4, "green": 10,
			}, {
				"red": 1, "green": 20, "blue": 13,
			}, {
				"green": 20,
			},
		},
	}, Game{
		GameId: 95,
		Rounds: []map[string]int{
			{
				"blue": 6, "green": 1,
			}, {
				"red": 3, "green": 11,
			}, {
				"blue": 4,
			},
		},
	}, Game{
		GameId: 96,
		Rounds: []map[string]int{
			{
				"red": 4, "green": 4, "blue": 3,
			}, {
				"green": 4, "blue": 17, "red": 3,
			}, {
				"red": 3, "blue": 3, "green": 13,
			}, {
				"red": 8, "blue": 7, "green": 6,
			},
		},
	}, Game{
		GameId: 97,
		Rounds: []map[string]int{
			{
				"blue": 5, "green": 9,
			}, {
				"green": 4, "blue": 4,
			}, {
				"red": 4, "green": 19,
			}, {
				"red": 2, "green": 3,
			}, {
				"green": 19, "blue": 3, "red": 4,
			}, {
				"red": 3, "green": 10,
			},
		},
	}, Game{
		GameId: 98,
		Rounds: []map[string]int{
			{
				"blue": 4, "red": 10, "green": 8,
			}, {
				"red": 2, "green": 3,
			}, {
				"red": 5, "blue": 4, "green": 10,
			},
		},
	}, Game{
		GameId: 99,
		Rounds: []map[string]int{
			{
				"blue": 9, "red": 12,
			}, {
				"blue": 9, "red": 11, "green": 13,
			}, {
				"blue": 9, "red": 1, "green": 13,
			}, {
				"blue": 4, "green": 12,
			}, {
				"blue": 10, "red": 17, "green": 8,
			},
		},
	}, Game{
		GameId: 100,
		Rounds: []map[string]int{
			{
				"red": 8, "green": 3,
			}, {
				"green": 4, "blue": 1, "red": 15,
			}, {
				"red": 10, "green": 8, "blue": 1,
			},
		},
	},
}
