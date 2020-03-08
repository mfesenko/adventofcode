package arcade

// TileID represents a type of the tile
type TileID int

const (
	// EmptyTile represents an empty tile. No game object appear in this tile.
	EmptyTile TileID = 0
	// WallTile represents a wall tile. Walls are indestructible barriers.
	WallTile TileID = 1
	// BlockTile represents a block tile. Blocks can be broken by the ball.
	BlockTile TileID = 2
	// HorizontalPaddleTile represents a horizontal paddle tile. The paddle is indestructible.
	HorizontalPaddleTile TileID = 3
	// BallTile represents a ball tile. The ball moves diagonally and bounces off objects.
	BallTile TileID = 4
)
