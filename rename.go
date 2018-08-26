package DeepSort

import (
	"strings"
	"path/filepath"
	"crypto/md5"
	"encoding/hex"
)

var replaceSpace = strings.NewReplacer(" ", "_")
var replaceComma = strings.NewReplacer(",", "")
var replaceDoubleQuote = strings.NewReplacer("\"", "")

// FormatFileName takes an old path of an image,
// its content as bytes and its tags (e.g. from Classify)
// and returns a new file name for an image along with
// the formatted tag portion of the new name
func FormatFileName(path string, image []byte, tags []string) (fullPath string, tagPart string) {
	tagPart = formatTags(tags)

	hashBytes := md5.Sum(image)
	hash := hex.EncodeToString(hashBytes[:])
	absPath, _ := filepath.Abs(path)
	dirPath := filepath.Dir(absPath)
	extension := path[len(path)-4:]
	newName := tagPart + "_" + hash + extension
	fullPath = filepath.Join(dirPath, newName)

	return
}

// Formats tags to be used in a file name
func formatTags(class []string) string {
	result := strings.Join(class, " ")
	result = replaceSpace.Replace(result)
	result = replaceComma.Replace(result)
	result = replaceDoubleQuote.Replace(result)
	return result
}
