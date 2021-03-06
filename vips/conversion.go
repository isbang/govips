package vips

// #cgo pkg-config: vips
// #include "conversion.h"
import "C"

type BlendMode int

const (
	BlendModeClear      BlendMode = C.VIPS_BLEND_MODE_CLEAR
	BlendModeSource     BlendMode = C.VIPS_BLEND_MODE_SOURCE
	BlendModeOver       BlendMode = C.VIPS_BLEND_MODE_OVER
	BlendModeIn         BlendMode = C.VIPS_BLEND_MODE_IN
	BlendModeOut        BlendMode = C.VIPS_BLEND_MODE_OUT
	BlendModeAtop       BlendMode = C.VIPS_BLEND_MODE_ATOP
	BlendModeDest       BlendMode = C.VIPS_BLEND_MODE_DEST
	BlendModeDestOver   BlendMode = C.VIPS_BLEND_MODE_DEST_OVER
	BlendModeDestIn     BlendMode = C.VIPS_BLEND_MODE_DEST_IN
	BlendModeDestOut    BlendMode = C.VIPS_BLEND_MODE_DEST_OUT
	BlendModeDestAtop   BlendMode = C.VIPS_BLEND_MODE_DEST_ATOP
	BlendModeXOR        BlendMode = C.VIPS_BLEND_MODE_XOR
	BlendModeAdd        BlendMode = C.VIPS_BLEND_MODE_ADD
	BlendModeSaturate   BlendMode = C.VIPS_BLEND_MODE_SATURATE
	BlendModeMultiply   BlendMode = C.VIPS_BLEND_MODE_MULTIPLY
	BlendModeScreen     BlendMode = C.VIPS_BLEND_MODE_SCREEN
	BlendModeOverlay    BlendMode = C.VIPS_BLEND_MODE_OVERLAY
	BlendModeDarken     BlendMode = C.VIPS_BLEND_MODE_DARKEN
	BlendModeLighten    BlendMode = C.VIPS_BLEND_MODE_LIGHTEN
	BlendModeColorDodge BlendMode = C.VIPS_BLEND_MODE_COLOUR_DODGE
	BlendModeColorBurn  BlendMode = C.VIPS_BLEND_MODE_COLOUR_BURN
	BlendModeHardLight  BlendMode = C.VIPS_BLEND_MODE_HARD_LIGHT
	BlendModeSoftLight  BlendMode = C.VIPS_BLEND_MODE_SOFT_LIGHT
	BlendModeDifference BlendMode = C.VIPS_BLEND_MODE_DIFFERENCE
	BlendModeExclusion  BlendMode = C.VIPS_BLEND_MODE_EXCLUSION
)

// Direction represents VIPS_DIRECTION type
type Direction int

// Direction enum
const (
	DirectionHorizontal Direction = C.VIPS_DIRECTION_HORIZONTAL
	DirectionVertical   Direction = C.VIPS_DIRECTION_VERTICAL
)

// Angle represents VIPS_ANGLE type
type Angle int

// Angle enum
const (
	Angle0   Angle = C.VIPS_ANGLE_D0
	Angle90  Angle = C.VIPS_ANGLE_D90
	Angle180 Angle = C.VIPS_ANGLE_D180
	Angle270 Angle = C.VIPS_ANGLE_D270
)

// Angle45 represents VIPS_ANGLE45 type
type Angle45 int

// Angle45 enum
const (
	Angle45_0   Angle45 = C.VIPS_ANGLE45_D0
	Angle45_45  Angle45 = C.VIPS_ANGLE45_D45
	Angle45_90  Angle45 = C.VIPS_ANGLE45_D90
	Angle45_135 Angle45 = C.VIPS_ANGLE45_D135
	Angle45_180 Angle45 = C.VIPS_ANGLE45_D180
	Angle45_225 Angle45 = C.VIPS_ANGLE45_D225
	Angle45_270 Angle45 = C.VIPS_ANGLE45_D270
	Angle45_315 Angle45 = C.VIPS_ANGLE45_D315
)

// ExtendStrategy represents VIPS_EXTEND type
type ExtendStrategy int

// ExtendStrategy enum
const (
	ExtendBlack      ExtendStrategy = C.VIPS_EXTEND_BLACK
	ExtendCopy       ExtendStrategy = C.VIPS_EXTEND_COPY
	ExtendRepeat     ExtendStrategy = C.VIPS_EXTEND_REPEAT
	ExtendMirror     ExtendStrategy = C.VIPS_EXTEND_MIRROR
	ExtendWhite      ExtendStrategy = C.VIPS_EXTEND_WHITE
	ExtendBackground ExtendStrategy = C.VIPS_EXTEND_BACKGROUND
)

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-copy
func vipsCopyImage(in *C.VipsImage) (*C.VipsImage, error) {
	var out *C.VipsImage

	if err := C.copy_image(in, &out); int(err) != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-embed
func vipsEmbed(in *C.VipsImage, left, top, width, height int, extend ExtendStrategy) (*C.VipsImage, error) {
	incOpCounter("embed")
	var out *C.VipsImage

	if err := C.embed_image(in, &out, C.int(left), C.int(top), C.int(width), C.int(height), C.int(extend), 0, 0, 0); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-flip
func vipsFlip(in *C.VipsImage, direction Direction) (*C.VipsImage, error) {
	incOpCounter("flip")
	var out *C.VipsImage

	if err := C.flip_image(in, &out, C.int(direction)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-extract-area
func vipsExtractArea(in *C.VipsImage, left, top, width, height int) (*C.VipsImage, error) {
	incOpCounter("extractArea")
	var out *C.VipsImage

	if err := C.extract_image_area(in, &out, C.int(left), C.int(top), C.int(width), C.int(height)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-extract-band
func vipsExtractBand(in *C.VipsImage, band, num int) (*C.VipsImage, error) {
	incOpCounter("extractBand")
	var out *C.VipsImage

	if err := C.extract_band(in, &out, C.int(band), C.int(num)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-rot
func vipsRotate(in *C.VipsImage, angle Angle) (*C.VipsImage, error) {
	incOpCounter("rot")
	var out *C.VipsImage

	if err := C.rot_image(in, &out, C.VipsAngle(angle)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-autorot
func vipsAutoRotate(in *C.VipsImage) (*C.VipsImage, error) {
	incOpCounter("autorot")
	var out *C.VipsImage

	if err := C.autorot_image(in, &out); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-zoom
func vipsZoom(in *C.VipsImage, xFactor, yFactor int) (*C.VipsImage, error) {
	incOpCounter("zoom")
	var out *C.VipsImage

	if err := C.zoom_image(in, &out, C.int(xFactor), C.int(yFactor)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-bandjoin
func vipsBandJoin(ins []*C.VipsImage) (*C.VipsImage, error) {
	incOpCounter("bandjoin")
	var out *C.VipsImage

	if err := C.bandjoin(&ins[0], &out, C.int(len(ins))); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-flatten
func vipsFlatten(in *C.VipsImage, color *Color) (*C.VipsImage, error) {
	incOpCounter("flatten")
	var out *C.VipsImage

	err := C.flatten_image(in, &out, C.double(color.R), C.double(color.G), C.double(color.B))
	if int(err) != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

func vipsAddAlpha(in *C.VipsImage) (*C.VipsImage, error) {
	incOpCounter("addAlpha")
	var out *C.VipsImage

	if err := C.add_alpha(in, &out); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-premultiply
func vipsPremultiplyAlpha(in *C.VipsImage) (*C.VipsImage, error) {
	incOpCounter("premultiplyAlpha")
	var out *C.VipsImage

	if err := C.premultiply_alpha(in, &out); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-unpremultiply
func vipsUnpremultiplyAlpha(in *C.VipsImage) (*C.VipsImage, error) {
	incOpCounter("unpremultiplyAlpha")
	var out *C.VipsImage

	if err := C.unpremultiply_alpha(in, &out); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

func vipsCast(in *C.VipsImage, bandFormat BandFormat) (*C.VipsImage, error) {
	incOpCounter("cast")
	var out *C.VipsImage

	if err := C.cast(in, &out, C.int(bandFormat)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-conversion.html#vips-composite2
func vipsComposite2(base *C.VipsImage, overlay *C.VipsImage, mode BlendMode, x, y int) (*C.VipsImage, error) {
	incOpCounter("composite")
	var out *C.VipsImage

	if err := C.composite2_image(base, overlay, &out, C.int(mode), C.gint(x), C.gint(y)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}
