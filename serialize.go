// Serialization functions

package main

// Serialize string to dest as Pascal string, padding up to len(dest) with zeros (TODO: or space?)
func StoreString(dest []byte, src string) {
	if len(src) > len(dest)-1 {
		src = src[:len(dest)-1]
	}
	dest[0] = byte(len(src))
	copy(dest[1:], src)
	for i := 1 + len(src); i < len(dest); i++ {
		dest[i] = 0
	}
}

// Load serialized string from src and return Go string
func LoadString(src []byte) string {
	length := int(src[0])
	if length > len(src)-1 {
		length = len(src) - 1
	}
	return string(src[1 : 1+length])
}

// Serialize RLE tile to dest
func StoreRleTile(dest []byte, rle TRleTile) {
	dest[0] = rle.Count
	dest[1] = rle.Tile.Element
	dest[2] = rle.Tile.Color
}

// Load serialized RLE tile from src
func LoadRleTile(src []byte) TRleTile {
	return TRleTile{
		Count: src[0],
		Tile: TTile{
			Element: src[1],
			Color:   src[2],
		},
	}
}

// Serialize board info to dest
func StoreBoardInfo(dest []byte, info *TBoardInfo) {
	dest[0] = info.MaxShots
	dest[1] = byte(BoolToInt(info.IsDark))
	copy(dest[2:6], info.NeighborBoards[:])
	dest[6] = byte(BoolToInt(info.ReenterWhenZapped))
	StoreString(dest[7:7+SizeOfBoardInfoMessage], info.Message)
	dest[66] = info.StartPlayerX
	dest[67] = info.StartPlayerY
	StoreInt16(dest[68:70], info.TimeLimitSec)
	for i := 70; i < len(dest); i++ {
		dest[i] = 0
	}
}

// Load serialized board info from src
func LoadBoardInfo(src []byte, info *TBoardInfo) {
	info.MaxShots = src[0]
	info.IsDark = src[1] != 0
	copy(info.NeighborBoards[:], src[2:6])
	info.ReenterWhenZapped = src[6] != 0
	info.Message = LoadString(src[7 : 7+SizeOfBoardInfoMessage])
	info.StartPlayerX = src[66]
	info.StartPlayerY = src[67]
	info.TimeLimitSec = LoadInt16(src[68:70])
}

// Serialize int16 to dest (little endian)
func StoreInt16(dest []byte, n int16) {
	u := uint16(n)
	dest[0] = byte(u & 0xFF)
	dest[1] = byte(u >> 8)
}

// Load serialized int16 from src (little endian)
func LoadInt16(src []byte) int16 {
	u := uint16(src[0]) | uint16(src[1])<<8
	return int16(u)
}

// Load serialized int32 from src (little endian)
func LoadInt32(src []byte) int32 {
	u := uint32(src[0]) | uint32(src[1])<<8 | uint32(src[2])<<16 | uint32(src[3])<<24
	return int32(u)
}

// Serialize stat to dest
func StoreStat(dest []byte, stat *TStat) {
	dest[0] = stat.X
	dest[1] = stat.Y
	StoreInt16(dest[2:4], stat.StepX)
	StoreInt16(dest[4:6], stat.StepY)
	StoreInt16(dest[6:8], stat.Cycle)
	dest[8] = stat.P1
	dest[9] = stat.P2
	dest[10] = stat.P3
	StoreInt16(dest[11:13], stat.Follower)
	StoreInt16(dest[13:15], stat.Leader)
	dest[15] = stat.Under.Element
	dest[16] = stat.Under.Color
	dest[17] = 0
	dest[18] = 0
	dest[19] = 0
	dest[20] = 0
	StoreInt16(dest[21:23], stat.DataPos)
	StoreInt16(dest[23:25], stat.DataLen)
	for i := 25; i < len(dest); i++ {
		dest[i] = 0
	}
}

// Load serialized stat from src
func LoadStat(src []byte, stat *TStat) {
	stat.X = src[0]
	stat.Y = src[1]
	stat.StepX = LoadInt16(src[2:4])
	stat.StepY = LoadInt16(src[4:6])
	stat.Cycle = LoadInt16(src[6:8])
	stat.P1 = src[8]
	stat.P2 = src[9]
	stat.P3 = src[10]
	stat.Follower = LoadInt16(src[11:13])
	stat.Leader = LoadInt16(src[13:15])
	stat.Under.Element = src[15]
	stat.Under.Color = src[16]
	stat.DataPos = LoadInt16(src[21:23])
	stat.DataLen = LoadInt16(src[23:25])
}

// Serialize world info struct into dest
func StoreWorldInfo(dest []byte, info *TWorldInfo) {
	StoreInt16(dest[:2], info.Ammo)
	StoreInt16(dest[2:4], info.Gems)
	for i := range info.Keys {
		dest[4+i] = byte(BoolToInt(info.Keys[i]))
	}
	StoreInt16(dest[11:13], info.Health)
	StoreInt16(dest[13:15], info.CurrentBoard)
	StoreInt16(dest[15:17], info.Torches)
	StoreInt16(dest[17:19], info.TorchTicks)
	StoreInt16(dest[19:21], info.EnergizerTicks)
	dest[21] = 0
	dest[22] = 0
	StoreInt16(dest[23:25], info.Score)
	StoreString(dest[25:46], info.Name)
	StoreInt16(dest[256:258], info.BoardTimeSec)
	StoreInt16(dest[258:260], info.BoardTimeHsec)
	dest[260] = byte(BoolToInt(info.IsSave))
	for i := 261; i < len(dest); i++ {
		dest[i] = 0
	}
}

// Load serialized world info struct from src
func LoadWorldInfo(src []byte, info *TWorldInfo) {
	info.Ammo = LoadInt16(src[:2])
	info.Gems = LoadInt16(src[2:4])
	for i := range info.Keys {
		info.Keys[i] = src[4+i] != 0
	}
	info.Health = LoadInt16(src[11:13])
	info.CurrentBoard = LoadInt16(src[13:15])
	info.Torches = LoadInt16(src[15:17])
	info.TorchTicks = LoadInt16(src[17:19])
	info.EnergizerTicks = LoadInt16(src[19:21])
	info.Score = LoadInt16(src[23:25])
	info.Name = LoadString(src[25:46])
	for i := range info.Flags {
		offset := 46 + i*21
		info.Flags[i] = LoadString(src[offset : offset+21])
	}
	info.BoardTimeSec = LoadInt16(src[256:258])
	info.BoardTimeHsec = LoadInt16(src[258:260])
	info.IsSave = src[260] != 0
}

func LoadResourceDataHeader(src []byte, header *TResourceDataHeader) {
	header.EntryCount = LoadInt16(src[:2])
	src = src[2:]
	for i := 0; i < len(header.Name); i++ {
		header.Name[i] = LoadString(src[:51])
		src = src[51:]
	}
	for i := 0; i < len(header.FileOffset); i++ {
		header.FileOffset[i] = LoadInt32(src[:4])
		src = src[4:]
	}
}

func LoadHighScoreList(src []byte, entries []THighScoreEntry) {
	for i := range entries {
		name := LoadString(src[:51])
		src = src[51:]
		score := LoadInt16(src[:2])
		src = src[2:]
		entries[i] = THighScoreEntry{name, score}
	}
}

func StoreHighScoreList(dest []byte, entries []THighScoreEntry) {
	for _, entry := range entries {
		StoreString(dest[:51], entry.Name)
		dest = dest[51:]
		StoreInt16(dest[:2], entry.Score)
		dest = dest[2:]
	}
}
