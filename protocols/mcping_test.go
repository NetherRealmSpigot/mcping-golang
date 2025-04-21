package protocols

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestDeliberateFailure(t *testing.T) {
	_, _, _, _, _, err := Ping("", 25565, "", Minecraft_1_8, 5)
	require.ErrorContains(t, err, "invalid host string")
	_, _, _, _, _, err = Ping("::1", 25565, "", Minecraft_1_8, 5)
	require.ErrorContains(t, err, "invalid host IP")

	_, _, _, _, _, err = Ping("127.0.0.1", 25565, "", Minecraft_1_7-1, 5)
	require.ErrorContains(t, err, "unknown protocol number")
	_, _, _, _, _, err = Ping("127.0.0.1", 25565, "", Minecraft_1_21_5+1, 5)
	require.ErrorContains(t, err, "unknown protocol number")

	_, _, _, _, _, err = Ping("127.0.0.1", 25565, "", Minecraft_1_8, 0)
	require.ErrorContains(t, err, "timeout in seconds must be bigger than 0")
	_, _, _, _, _, err = Ping("127.0.0.1", 25565, "", Minecraft_1_8, -1)
	require.ErrorContains(t, err, "timeout in seconds must be bigger than 0")

	_, _, _, _, _, err = Ping("doesntexist.local", 25565, "", Minecraft_1_8, 2)
	assert.NotNil(t, err)

	_, _, _, _, _, err = Ping("169.254.169.254", 25565, "", Minecraft_1_8, 2)
	require.ErrorContains(t, err, "all ip addresses are not reachable, giving up")
}

func TestPingDeliberateFailure(t *testing.T) {
	host, port, fakeHost, _, buf, err := Ping("play.cubecraft.net", 25566, "", Minecraft_1_21_5, 1)
	assert.Equal(t, "play.cubecraft.net", host)
	assert.Equal(t, uint16(25566), port)
	assert.Equal(t, "play.cubecraft.net", fakeHost)
	assert.Nil(t, buf)
	assert.NotNil(t, err)
}

func TestPing(t *testing.T) {
	host, port, fakeHost, protocol, buf, err := Ping("mc.hypixel.net", 0, "", Minecraft_1_8, 5)
	assert.Equal(t, "mc.hypixel.net", host)
	assert.Equal(t, uint16(25565), port)
	assert.Equal(t, "mc.hypixel.net", fakeHost)
	assert.Equal(t, Minecraft_1_8, protocol)
	assert.NotNil(t, buf)
	assert.Nil(t, err)

	host, port, fakeHost, protocol, buf, err = Ping("play.cubecraft.net", 25565, "cubecraft", Minecraft_1_21_5, 5)
	assert.Equal(t, "play.cubecraft.net", host)
	assert.Equal(t, uint16(25565), port)
	assert.Equal(t, "cubecraft", fakeHost)
	assert.Equal(t, Minecraft_1_21_5, protocol)
	assert.NotNil(t, buf)
	assert.Nil(t, err)
}

func TestPingSRV(t *testing.T) {
	if os.Getenv("GITHUB_ACTION") == "" {
		t.Skip("Run this test in CI environment!")
	}
	host, port, fakeHost, protocol, buf, err := Ping("rebellious.nekoli.dev", 0, "", Minecraft_1_21_4, 5)
	assert.Equal(t, "rebellious.nekoli.dev", host)
	assert.Equal(t, uint16(2025), port)
	assert.Equal(t, "rebellious.nekoli.dev", fakeHost)
	assert.Equal(t, Minecraft_1_21_4, protocol)
	assert.NotNil(t, buf)
	assert.Nil(t, err)
}
