package colors

import "image/color"

var Colors *dfault_colors

type dfault_colors struct {
	White           color.RGBA
	Black           color.RGBA
	Gray_primary    color.RGBA
	Gray_secondary  color.RGBA
	Red_primary     color.RGBA
	Red_secondary   color.RGBA
	Green_primary   color.RGBA
	Green_secondary color.RGBA
}

func init() {
	Colors = &dfault_colors{
		White:           color.RGBA{255, 255, 255, 255}, // White color with full opacity
		Black:           color.RGBA{0, 0, 0, 255},       // Black color with full opacity
		Gray_primary:    color.RGBA{128, 128, 128, 255}, // Gray primary
		Gray_secondary:  color.RGBA{192, 192, 192, 255}, // Lighter gray secondary
		Red_primary:     color.RGBA{255, 0, 0, 255},     // Red primary
		Red_secondary:   color.RGBA{200, 0, 0, 255},     // Darker red secondary
		Green_primary:   color.RGBA{0, 255, 0, 255},     // Green primary
		Green_secondary: color.RGBA{0, 200, 0, 255},     // Darker green secondary
	}
}
