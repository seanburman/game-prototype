package constants

const (
	SPRITE_SCALE      = 3
	SPRITE_BITS       = 32
	GAME_WINDOW_SCALE = 12
)

func GetDimensions() int {
	return SPRITE_SCALE * SPRITE_BITS * GAME_WINDOW_SCALE
}
