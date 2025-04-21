package protocols

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsKnownProtocolNumber(t *testing.T) {
	nums := []int{
		Minecraft_1_7, Minecraft_1_7_1, Minecraft_1_7_2, Minecraft_1_7_3, Minecraft_1_7_4, Minecraft_1_7_5, Minecraft_1_7_6, Minecraft_1_7_7, Minecraft_1_7_8, Minecraft_1_7_9, Minecraft_1_7_10,
		Minecraft_1_8, Minecraft_1_8_1, Minecraft_1_8_2, Minecraft_1_8_3, Minecraft_1_8_4, Minecraft_1_8_5, Minecraft_1_8_6, Minecraft_1_8_7, Minecraft_1_8_8, Minecraft_1_8_9,
		Minecraft_1_9, Minecraft_1_9_1, Minecraft_1_9_2, Minecraft_1_9_3, Minecraft_1_9_4,
		Minecraft_1_10, Minecraft_1_10_1, Minecraft_1_10_2,
		Minecraft_1_11, Minecraft_1_11_1, Minecraft_1_11_2,
		Minecraft_1_12, Minecraft_1_12_1, Minecraft_1_12_2,
		Minecraft_1_13, Minecraft_1_13_1, Minecraft_1_13_2,
		Minecraft_1_14, Minecraft_1_14_1, Minecraft_1_14_2, Minecraft_1_14_3, Minecraft_1_14_4,
		Minecraft_1_15, Minecraft_1_15_1, Minecraft_1_15_2,
		Minecraft_1_16, Minecraft_1_16_1, Minecraft_1_16_2, Minecraft_1_16_3, Minecraft_1_16_4, Minecraft_1_16_5,
		Minecraft_1_17, Minecraft_1_17_1,
		Minecraft_1_18, Minecraft_1_18_1, Minecraft_1_18_2,
		Minecraft_1_19, Minecraft_1_19_1, Minecraft_1_19_2, Minecraft_1_19_3, Minecraft_1_19_4,
		Minecraft_1_20, Minecraft_1_20_1, Minecraft_1_20_2, Minecraft_1_20_3, Minecraft_1_20_4, Minecraft_1_20_5, Minecraft_1_20_6,
		Minecraft_1_21, Minecraft_1_21_1, Minecraft_1_21_2, Minecraft_1_21_3, Minecraft_1_21_4, Minecraft_1_21_5,
	}
	t.Run("For known protocol number should return true", func(t0 *testing.T) {
		for _, i := range nums {
			require.True(t0, IsKnownProtocolNumber(i))
		}
	})
}
