package solution

import "github.com/kyokomi/emoji/v2"

//goland:noinspection GoUnusedExportedFunction
func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
