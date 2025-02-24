package main

const (
	Minecraft_1_7   = 3
	Minecraft_1_7_1 = 3

	Minecraft_1_7_2 = 4
	Minecraft_1_7_3 = 4
	Minecraft_1_7_4 = 4
	Minecraft_1_7_5 = 4

	Minecraft_1_7_6  = 5
	Minecraft_1_7_7  = 5
	Minecraft_1_7_8  = 5
	Minecraft_1_7_9  = 5
	Minecraft_1_7_10 = 5

	Minecraft_1_8   = 47
	Minecraft_1_8_1 = 47
	Minecraft_1_8_2 = 47
	Minecraft_1_8_3 = 47
	Minecraft_1_8_4 = 47
	Minecraft_1_8_5 = 47
	Minecraft_1_8_6 = 47
	Minecraft_1_8_7 = 47
	Minecraft_1_8_8 = 47
	Minecraft_1_8_9 = 47

	Minecraft_1_9   = 107
	Minecraft_1_9_1 = 108
	Minecraft_1_9_2 = 109
	Minecraft_1_9_3 = 110
	Minecraft_1_9_4 = 110

	Minecraft_1_10   = 210
	Minecraft_1_10_1 = 210
	Minecraft_1_10_2 = 210

	Minecraft_1_11   = 315
	Minecraft_1_11_1 = 316
	Minecraft_1_11_2 = 316

	Minecraft_1_12   = 335
	Minecraft_1_12_1 = 338
	Minecraft_1_12_2 = 340

	Minecraft_1_13   = 393
	Minecraft_1_13_1 = 401
	Minecraft_1_13_2 = 404

	Minecraft_1_14   = 477
	Minecraft_1_14_1 = 480
	Minecraft_1_14_2 = 485
	Minecraft_1_14_3 = 490
	Minecraft_1_14_4 = 498

	Minecraft_1_15   = 573
	Minecraft_1_15_1 = 575
	Minecraft_1_15_2 = 578

	Minecraft_1_16   = 735
	Minecraft_1_16_1 = 736
	Minecraft_1_16_2 = 751
	Minecraft_1_16_3 = 753
	Minecraft_1_16_4 = 754
	Minecraft_1_16_5 = 754

	Minecraft_1_17   = 755
	Minecraft_1_17_1 = 756

	Minecraft_1_18   = 757
	Minecraft_1_18_1 = 757
	Minecraft_1_18_2 = 758

	Minecraft_1_19   = 759
	Minecraft_1_19_1 = 760
	Minecraft_1_19_2 = 760
	Minecraft_1_19_3 = 761
	Minecraft_1_19_4 = 762

	Minecraft_1_20   = 763
	Minecraft_1_20_1 = 763
	Minecraft_1_20_2 = 764
	Minecraft_1_20_3 = 765
	Minecraft_1_20_4 = 765
	Minecraft_1_20_5 = 766
	Minecraft_1_20_6 = 766

	Minecraft_1_21   = 767
	Minecraft_1_21_1 = 767
	Minecraft_1_21_2 = 768
	Minecraft_1_21_3 = 768
	Minecraft_1_21_4 = 769
)

var knownProtocolNumbers = map[int]struct{}{
	Minecraft_1_7_1: {},

	Minecraft_1_7_5: {},

	Minecraft_1_7_10: {},

	Minecraft_1_8_9: {},

	Minecraft_1_9:   {},
	Minecraft_1_9_1: {},
	Minecraft_1_9_2: {},
	Minecraft_1_9_4: {},

	Minecraft_1_10_2: {},

	Minecraft_1_11:   {},
	Minecraft_1_11_2: {},

	Minecraft_1_12:   {},
	Minecraft_1_12_1: {},
	Minecraft_1_12_2: {},

	Minecraft_1_13:   {},
	Minecraft_1_13_1: {},
	Minecraft_1_13_2: {},

	Minecraft_1_14:   {},
	Minecraft_1_14_1: {},
	Minecraft_1_14_2: {},
	Minecraft_1_14_3: {},
	Minecraft_1_14_4: {},

	Minecraft_1_15:   {},
	Minecraft_1_15_1: {},
	Minecraft_1_15_2: {},

	Minecraft_1_16:   {},
	Minecraft_1_16_1: {},
	Minecraft_1_16_2: {},
	Minecraft_1_16_3: {},
	Minecraft_1_16_5: {},

	Minecraft_1_17:   {},
	Minecraft_1_17_1: {},

	Minecraft_1_18_1: {},
	Minecraft_1_18_2: {},

	Minecraft_1_19:   {},
	Minecraft_1_19_2: {},
	Minecraft_1_19_3: {},
	Minecraft_1_19_4: {},

	Minecraft_1_20_1: {},
	Minecraft_1_20_2: {},
	Minecraft_1_20_4: {},
	Minecraft_1_20_6: {},

	Minecraft_1_21_1: {},
	Minecraft_1_21_3: {},
	Minecraft_1_21_4: {},
}

func IsKnownProtocolNumber(n int) bool {
	_, exists := knownProtocolNumbers[n]
	return exists
}
