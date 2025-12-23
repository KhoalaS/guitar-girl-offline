package embeds

import "embed"

//go:embed static
var StaticFiles embed.FS

//go:embed game_data.json
var GameData []byte

//go:embed game_data.bz2.base64
var RawGameData []byte
