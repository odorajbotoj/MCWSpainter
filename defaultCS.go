package painter

// Thank you, Python3!
// warning: 本颜色库仅适用于1.17的三维地图画
var DefaultColorSpace = ColorSpace{
	[...]Block{
		// Block{"", , Color{, , }, },
		Block{"slime", 0, Color{90, 126, 39}, 1},
		Block{"slime", 0, Color{109, 153, 48}, 0},
		Block{"slime", 0, Color{127, 178, 56}, -1},
		Block{"end_bricks", 0, Color{175, 165, 115}, 1},
		Block{"end_bricks", 0, Color{212, 200, 140}, 0},
		Block{"end_bricks", 0, Color{247, 233, 163}, -1},
		Block{"brown_mushroom_block", 15, Color{141, 141, 141}, 1},
		Block{"brown_mushroom_block", 15, Color{171, 171, 171}, 0},
		Block{"brown_mushroom_block", 15, Color{199, 199, 199}, -1},
		Block{"redstone_block", 0, Color{181, 0, 0}, 1},
		Block{"redstone_block", 0, Color{219, 0, 0}, 0},
		Block{"redstone_block", 0, Color{255, 0, 0}, -1},
		Block{"blue_ice", 0, Color{113, 113, 181}, 1},
		Block{"blue_ice", 0, Color{137, 137, 219}, 0},
		Block{"blue_ice", 0, Color{160, 160, 255}, -1},
		Block{"iron_block", 0, Color{118, 118, 118}, 1},
		Block{"iron_block", 0, Color{143, 143, 143}, 0},
		Block{"iron_block", 0, Color{167, 167, 167}, -1},
		Block{"leaves", 0, Color{0, 88, 0}, 1},
		Block{"leaves", 0, Color{0, 106, 0}, 0},
		Block{"leaves", 0, Color{0, 124, 0}, -1},
		Block{"concrete", 0, Color{181, 181, 181}, 1},
		Block{"concrete", 0, Color{219, 219, 219}, 0},
		Block{"concrete", 0, Color{255, 255, 255}, -1},
		Block{"clay", 0, Color{116, 119, 130}, 1},
		Block{"clay", 0, Color{141, 144, 158}, 0},
		Block{"clay", 0, Color{164, 168, 184}, -1},
		Block{"stone", 2, Color{107, 77, 54}, 1},
		Block{"stone", 2, Color{129, 93, 66}, 0},
		Block{"stone", 2, Color{151, 109, 77}, -1},
		Block{"cobblestone", 0, Color{79, 79, 79}, 1},
		Block{"cobblestone", 0, Color{96, 96, 96}, 0},
		Block{"cobblestone", 0, Color{112, 112, 112}, -1},
		Block{"planks", 0, Color{101, 84, 51}, 1},
		Block{"planks", 0, Color{122, 102, 61}, 0},
		Block{"planks", 0, Color{143, 119, 72}, -1},
		Block{"quartz_block", 0, Color{181, 178, 173}, 1},
		Block{"quartz_block", 0, Color{219, 216, 210}, 0},
		Block{"quartz_block", 0, Color{255, 252, 245}, -1},
		Block{"concrete", 1, Color{153, 90, 36}, 1},
		Block{"concrete", 1, Color{185, 109, 43}, 0},
		Block{"concrete", 1, Color{216, 127, 51}, -1},
		Block{"concrete", 2, Color{126, 53, 153}, 1},
		Block{"concrete", 2, Color{153, 65, 185}, 0},
		Block{"concrete", 2, Color{178, 76, 216}, -1},
		Block{"concrete", 3, Color{72, 108, 153}, 1},
		Block{"concrete", 3, Color{87, 131, 185}, 0},
		Block{"concrete", 3, Color{102, 153, 216}, -1},
		Block{"concrete", 4, Color{162, 162, 36}, 1},
		Block{"concrete", 4, Color{196, 196, 43}, 0},
		Block{"concrete", 4, Color{229, 229, 51}, -1},
		Block{"concrete", 5, Color{90, 144, 17}, 1},
		Block{"concrete", 5, Color{109, 175, 21}, 0},
		Block{"concrete", 5, Color{127, 204, 25}, -1},
		Block{"concrete", 6, Color{171, 90, 117}, 1},
		Block{"concrete", 6, Color{208, 109, 141}, 0},
		Block{"concrete", 6, Color{242, 127, 165}, -1},
		Block{"concrete", 7, Color{53, 53, 53}, 1},
		Block{"concrete", 7, Color{65, 65, 65}, 0},
		Block{"concrete", 7, Color{76, 76, 76}, -1},
		Block{"concrete", 8, Color{108, 108, 108}, 1},
		Block{"concrete", 8, Color{131, 131, 131}, 0},
		Block{"concrete", 8, Color{153, 153, 153}, -1},
		Block{"concrete", 9, Color{53, 90, 108}, 1},
		Block{"concrete", 9, Color{65, 109, 131}, 0},
		Block{"concrete", 9, Color{76, 127, 153}, -1},
		Block{"concrete", 10, Color{90, 44, 126}, 1},
		Block{"concrete", 10, Color{109, 54, 153}, 0},
		Block{"concrete", 10, Color{127, 63, 178}, -1},
		Block{"concrete", 11, Color{36, 53, 126}, 1},
		Block{"concrete", 11, Color{43, 65, 153}, 0},
		Block{"concrete", 11, Color{51, 76, 178}, -1},
		Block{"concrete", 12, Color{72, 53, 36}, 1},
		Block{"concrete", 12, Color{87, 65, 43}, 0},
		Block{"concrete", 12, Color{102, 76, 51}, -1},
		Block{"concrete", 13, Color{72, 90, 36}, 1},
		Block{"concrete", 13, Color{87, 109, 43}, 0},
		Block{"concrete", 13, Color{102, 127, 51}, -1},
		Block{"concrete", 14, Color{108, 36, 36}, 1},
		Block{"concrete", 14, Color{131, 43, 43}, 0},
		Block{"concrete", 14, Color{153, 51, 51}, -1},
		Block{"concrete", 15, Color{17, 17, 17}, 1},
		Block{"concrete", 15, Color{21, 21, 21}, 0},
		Block{"concrete", 15, Color{25, 25, 25}, -1},
		Block{"gold_block", 0, Color{177, 168, 54}, 1},
		Block{"gold_block", 0, Color{215, 204, 66}, 0},
		Block{"gold_block", 0, Color{250, 238, 77}, -1},
		Block{"diamond_block", 0, Color{65, 155, 151}, 1},
		Block{"diamond_block", 0, Color{79, 188, 183}, 0},
		Block{"diamond_block", 0, Color{92, 219, 213}, -1},
		Block{"lapis_block", 0, Color{52, 90, 181}, 1},
		Block{"lapis_block", 0, Color{63, 110, 219}, 0},
		Block{"lapis_block", 0, Color{74, 128, 255}, -1},
		Block{"emerald_block", 0, Color{0, 154, 41}, 1},
		Block{"emerald_block", 0, Color{0, 186, 49}, 0},
		Block{"emerald_block", 0, Color{0, 217, 58}, -1},
		Block{"planks", 1, Color{91, 61, 34}, 1},
		Block{"planks", 1, Color{110, 73, 42}, 0},
		Block{"planks", 1, Color{129, 86, 49}, -1},
		Block{"netherrack", 0, Color{79, 1, 0}, 1},
		Block{"netherrack", 0, Color{96, 1, 0}, 0},
		Block{"netherrack", 0, Color{112, 2, 0}, -1},
		Block{"stained_hardened_clay", 0, Color{148, 125, 114}, 1},
		Block{"stained_hardened_clay", 0, Color{179, 152, 138}, 0},
		Block{"stained_hardened_clay", 0, Color{209, 177, 161}, -1},
		Block{"stained_hardened_clay", 1, Color{112, 58, 25}, 1},
		Block{"stained_hardened_clay", 1, Color{136, 70, 30}, 0},
		Block{"stained_hardened_clay", 1, Color{159, 82, 36}, -1},
		Block{"stained_hardened_clay", 2, Color{105, 61, 76}, 1},
		Block{"stained_hardened_clay", 2, Color{128, 74, 92}, 0},
		Block{"stained_hardened_clay", 2, Color{149, 87, 108}, -1},
		Block{"stained_hardened_clay", 3, Color{79, 76, 97}, 1},
		Block{"stained_hardened_clay", 3, Color{96, 92, 118}, 0},
		Block{"stained_hardened_clay", 3, Color{112, 108, 138}, -1},
		Block{"stained_hardened_clay", 4, Color{132, 94, 25}, 1},
		Block{"stained_hardened_clay", 4, Color{159, 114, 30}, 0},
		Block{"stained_hardened_clay", 4, Color{186, 133, 36}, -1},
		Block{"stained_hardened_clay", 5, Color{73, 83, 37}, 1},
		Block{"stained_hardened_clay", 5, Color{88, 100, 45}, 0},
		Block{"stained_hardened_clay", 5, Color{103, 117, 53}, -1},
		Block{"stained_hardened_clay", 6, Color{113, 54, 55}, 1},
		Block{"stained_hardened_clay", 6, Color{137, 66, 67}, 0},
		Block{"stained_hardened_clay", 6, Color{160, 77, 78}, -1},
		Block{"stained_hardened_clay", 7, Color{40, 29, 24}, 1},
		Block{"stained_hardened_clay", 7, Color{49, 35, 30}, 0},
		Block{"stained_hardened_clay", 7, Color{57, 41, 35}, -1},
		Block{"stained_hardened_clay", 8, Color{95, 75, 69}, 1},
		Block{"stained_hardened_clay", 8, Color{116, 92, 84}, 0},
		Block{"stained_hardened_clay", 8, Color{135, 107, 98}, -1},
		Block{"stained_hardened_clay", 9, Color{61, 65, 65}, 1},
		Block{"stained_hardened_clay", 9, Color{74, 79, 79}, 0},
		Block{"stained_hardened_clay", 9, Color{87, 92, 92}, -1},
		Block{"stained_hardened_clay", 10, Color{86, 51, 62}, 1},
		Block{"stained_hardened_clay", 10, Color{104, 62, 75}, 0},
		Block{"stained_hardened_clay", 10, Color{122, 73, 88}, -1},
		Block{"stained_hardened_clay", 11, Color{53, 44, 65}, 1},
		Block{"stained_hardened_clay", 11, Color{65, 53, 79}, 0},
		Block{"stained_hardened_clay", 11, Color{76, 62, 92}, -1},
		Block{"stained_hardened_clay", 12, Color{53, 35, 24}, 1},
		Block{"stained_hardened_clay", 12, Color{65, 43, 30}, 0},
		Block{"stained_hardened_clay", 12, Color{76, 50, 35}, -1},
		Block{"stained_hardened_clay", 13, Color{53, 58, 29}, 1},
		Block{"stained_hardened_clay", 13, Color{65, 70, 36}, 0},
		Block{"stained_hardened_clay", 13, Color{76, 82, 42}, -1},
		Block{"stained_hardened_clay", 14, Color{100, 42, 32}, 1},
		Block{"stained_hardened_clay", 14, Color{122, 51, 39}, 0},
		Block{"stained_hardened_clay", 14, Color{142, 60, 46}, -1},
		Block{"stained_hardened_clay", 15, Color{26, 15, 11}, 1},
		Block{"stained_hardened_clay", 15, Color{31, 18, 13}, 0},
		Block{"stained_hardened_clay", 15, Color{37, 22, 16}, -1},
		Block{"crimson_nylium", 0, Color{134, 34, 34}, 1},
		Block{"crimson_nylium", 0, Color{162, 41, 42}, 0},
		Block{"crimson_nylium", 0, Color{189, 48, 49}, -1},
		Block{"crimson_planks", 0, Color{105, 44, 68}, 1},
		Block{"crimson_planks", 0, Color{127, 54, 83}, 0},
		Block{"crimson_planks", 0, Color{148, 63, 97}, -1},
		Block{"crimson_hyphae", 0, Color{65, 17, 20}, 1},
		Block{"crimson_hyphae", 0, Color{79, 21, 24}, 0},
		Block{"crimson_hyphae", 0, Color{92, 25, 29}, -1},
		Block{"oxidized_copper", 0, Color{15, 89, 95}, 1},
		Block{"oxidized_copper", 0, Color{18, 108, 115}, 0},
		Block{"oxidized_copper", 0, Color{22, 126, 134}, -1},
		Block{"weathered_copper", 0, Color{41, 100, 99}, 1},
		Block{"weathered_copper", 0, Color{49, 122, 120}, 0},
		Block{"weathered_copper", 0, Color{58, 142, 140}, -1},
		Block{"warped_hyphae", 0, Color{61, 31, 44}, 1},
		Block{"warped_hyphae", 0, Color{73, 37, 53}, 0},
		Block{"warped_hyphae", 0, Color{86, 44, 62}, -1},
		Block{"warped_wart_block", 0, Color{14, 127, 94}, 1},
		Block{"warped_wart_block", 0, Color{17, 154, 114}, 0},
		Block{"warped_wart_block", 0, Color{20, 180, 133}, -1},
		Block{"deepslate", 0, Color{61, 61, 61}, 1},
		Block{"deepslate", 0, Color{73, 73, 73}, 0},
		Block{"deepslate", 0, Color{86, 86, 86}, -1},
		Block{"raw_iron_block", 0, Color{132, 106, 89}, 1},
		Block{"raw_iron_block", 0, Color{159, 129, 108}, 0},
		Block{"raw_iron_block", 0, Color{186, 150, 126}, -1},
	},
}