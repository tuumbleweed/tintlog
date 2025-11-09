package logger

// Colorizer applies ANSI color to a string (per line). Nil Fn = no color.
type Colorizer struct {
	Name string
	Fn   func(string) string
}

// Apply safely applies the colorizer if present.
func (c Colorizer) Apply(s string) string {
	if c.Fn == nil {
		return s
	}
	return c.Fn(s)
}

// --- Helpers to build colorizers ---

// FgColorizer returns a per-line foreground colorizer.
func FgColorizer(name string, fg Color) Colorizer {
	return Colorizer{
		Name: name,
		Fn:   func(s string) string { return FgLines(s, fg) },
	}
}

// (Future-friendly) Foreground+background variant if you add BG colors later.
// func FgBgColorizer(name string, fg, bg Color) Colorizer {
// 	return Colorizer{
// 		Name: name,
// 		Fn:   func(s string) string { return WithFgBgLines(s, fg, bg) },
// 	}
// }

// --- Registry of reusable colorizers ---

var Colorizers = map[string]Colorizer{
	// Base hues
	"Red":    FgColorizer("Red", Red),
	"Orange": FgColorizer("Orange", Orange),
	"Yellow": FgColorizer("Yellow", Yellow),
	"Green":  FgColorizer("Green", Green),
	"Cyan":   FgColorizer("Cyan", Cyan),
	"Blue":   FgColorizer("Blue", Blue),
	"Purple": FgColorizer("Purple", Purple),
	"Gray":   FgColorizer("Gray", Gray),

	// Bright tints
	"BrightRed":    FgColorizer("BrightRed", BrightRed),
	"BrightOrange": FgColorizer("BrightOrange", BrightOrange),
	"BrightYellow": FgColorizer("BrightYellow", BrightYellow),
	"BrightGreen":  FgColorizer("BrightGreen", BrightGreen),
	"BrightCyan":   FgColorizer("BrightCyan", BrightCyan),
	"BrightBlue":   FgColorizer("BrightBlue", BrightBlue),
	"BrightPurple": FgColorizer("BrightPurple", BrightPurple),
	"BrightGray":   FgColorizer("BrightGray", BrightGray),

	// Dim shades
	"DimRed":    FgColorizer("DimRed", DimRed),
	"DimOrange": FgColorizer("DimOrange", DimOrange),
	"DimYellow": FgColorizer("DimYellow", DimYellow),
	"DimGreen":  FgColorizer("DimGreen", DimGreen),
	"DimCyan":   FgColorizer("DimCyan", DimCyan),
	"DimBlue":   FgColorizer("DimBlue", DimBlue),
	"DimPurple": FgColorizer("DimPurple", DimPurple),
	"DimGray":   FgColorizer("DimGray", DimGray),

	// No color
	"NoColor": {Name: "NoColor", Fn: nil},
}

// Convenience aliases (import-friendly)
var (
	RedText        = Colorizers["Red"]
	OrangeText     = Colorizers["Orange"]
	YellowText     = Colorizers["Yellow"]
	GreenText      = Colorizers["Green"]
	CyanText       = Colorizers["Cyan"]
	BlueText       = Colorizers["Blue"]
	PurpleText     = Colorizers["Purple"]
	GrayText       = Colorizers["Gray"]

	BrightRedText    = Colorizers["BrightRed"]
	BrightOrangeText = Colorizers["BrightOrange"]
	BrightYellowText = Colorizers["BrightYellow"]
	BrightGreenText  = Colorizers["BrightGreen"]
	BrightCyanText   = Colorizers["BrightCyan"]
	BrightBlueText   = Colorizers["BrightBlue"]
	BrightPurpleText = Colorizers["BrightPurple"]
	BrightGrayText   = Colorizers["BrightGray"]

	DimRedText    = Colorizers["DimRed"]
	DimOrangeText = Colorizers["DimOrange"]
	DimYellowText = Colorizers["DimYellow"]
	DimGreenText  = Colorizers["DimGreen"]
	DimCyanText   = Colorizers["DimCyan"]
	DimBlueText   = Colorizers["DimBlue"]
	DimPurpleText = Colorizers["DimPurple"]
	DimGrayText   = Colorizers["DimGray"]

	NoColor = Colorizers["NoColor"]
)

// RegisterColorizer adds/overwrites a colorizer in the registry.
func RegisterColorizer(name string, fn func(string) string) Colorizer {
	c := Colorizer{Name: name, Fn: fn}
	Colorizers[name] = c
	return c
}

// RegisterFg is a shorthand to register a foreground-only colorizer.
func RegisterFg(name string, fg Color) Colorizer {
	return RegisterColorizer(name, func(s string) string { return FgLines(s, fg) })
}
