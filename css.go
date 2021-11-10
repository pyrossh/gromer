package wapp

type M map[string]interface{}
type MS map[string]string
type Arr []interface{}

type KeyValues struct {
	Keys   M
	Values MS
}

var colors = KeyValues{
	Keys: M{
		"bg":       "background-color",
		"text":     "color",
		"divide":   "border-color",
		"border":   "border-color",
		"ring":     "--tw-ring-color",
		"border-l": "border-left-color",
		"border-r": "border-right-color",
		"border-t": "border-top-color",
		"border-b": "border-bottom-color",
	},
	Values: MS{
		"transparent": "transparent",
		"current":     "currentColor",
		"black":       "rgba(0, 0, 0, 1)",
		"white":       "rgba(255, 255, 255, 1)",
		"gray-50":     "rgba(249, 250, 251, 1)",
		"gray-100":    "rgba(243, 244, 246, 1)",
		"gray-200":    "rgba(229, 231, 235, 1)",
		"gray-300":    "rgba(209, 213, 219, 1)",
		"gray-400":    "rgba(156, 163, 175, 1)",
		"gray-500":    "rgba(107, 114, 128, 1)",
		"gray-600":    "rgba(75, 85, 99, 1)",
		"gray-700":    "rgba(55, 65, 81, 1)",
		"gray-800":    "rgba(31, 41, 55, 1)",
		"gray-900":    "rgba(17, 24, 39, 1)",
		"red-50":      "rgba(254, 242, 242, 1)",
		"red-100":     "rgba(254, 226, 226, 1)",
		"red-200":     "rgba(254, 202, 202, 1)",
		"red-300":     "rgba(252, 165, 165, 1)",
		"red-400":     "rgba(248, 113, 113, 1)",
		"red-500":     "rgba(239, 68, 68, 1)",
		"red-600":     "rgba(220, 38, 38, 1)",
		"red-700":     "rgba(185, 28, 28, 1)",
		"red-800":     "rgba(153, 27, 27, 1)",
		"red-900":     "rgba(127, 29, 29, 1)",
		"yellow-50":   "rgba(255, 251, 235, 1)",
		"yellow-100":  "rgba(254, 243, 199, 1)",
		"yellow-200":  "rgba(253, 230, 138, 1)",
		"yellow-300":  "rgba(252, 211, 77, 1)",
		"yellow-400":  "rgba(251, 191, 36, 1)",
		"yellow-500":  "rgba(245, 158, 11, 1)",
		"yellow-600":  "rgba(217, 119, 6, 1)",
		"yellow-700":  "rgba(180, 83, 9, 1)",
		"yellow-800":  "rgba(146, 64, 14, 1)",
		"yellow-900":  "rgba(120, 53, 15, 1)",
		"green-50":    "rgba(236, 253, 245, 1)",
		"green-100":   "rgba(209, 250, 229, 1)",
		"green-200":   "rgba(167, 243, 208, 1)",
		"green-300":   "rgba(110, 231, 183, 1)",
		"green-400":   "rgba(52, 211, 153, 1)",
		"green-500":   "rgba(16, 185, 129, 1)",
		"green-600":   "rgba(5, 150, 105, 1)",
		"green-700":   "rgba(4, 120, 87, 1)",
		"green-800":   "rgba(6, 95, 70, 1)",
		"green-900":   "rgba(6, 78, 59, 1)",
		"blue-50":     "rgba(239, 246, 255, 1)",
		"blue-100":    "rgba(219, 234, 254, 1)",
		"blue-200":    "rgba(191, 219, 254, 1)",
		"blue-300":    "rgba(147, 197, 253, 1)",
		"blue-400":    "rgba(96, 165, 250, 1)",
		"blue-500":    "rgba(59, 130, 246, 1)",
		"blue-600":    "rgba(37, 99, 235, 1)",
		"blue-700":    "rgba(29, 78, 216, 1)",
		"blue-800":    "rgba(30, 64, 175, 1)",
		"blue-900":    "rgba(30, 58, 138, 1)",
		"indigo-50":   "rgba(238, 242, 255, 1)",
		"indigo-100":  "rgba(224, 231, 255, 1)",
		"indigo-200":  "rgba(199, 210, 254, 1)",
		"indigo-300":  "rgba(165, 180, 252, 1)",
		"indigo-400":  "rgba(129, 140, 248, 1)",
		"indigo-500":  "rgba(99, 102, 241, 1)",
		"indigo-600":  "rgba(79, 70, 229, 1)",
		"indigo-700":  "rgba(67, 56, 202, 1)",
		"indigo-800":  "rgba(55, 48, 163, 1)",
		"indigo-900":  "rgba(49, 46, 129, 1)",
		"purple-50":   "rgba(245, 243, 255, 1)",
		"purple-100":  "rgba(237, 233, 254, 1)",
		"purple-200":  "rgba(221, 214, 254, 1)",
		"purple-300":  "rgba(196, 181, 253, 1)",
		"purple-400":  "rgba(167, 139, 250, 1)",
		"purple-500":  "rgba(139, 92, 246, 1)",
		"purple-600":  "rgba(124, 58, 237, 1)",
		"purple-700":  "rgba(109, 40, 217, 1)",
		"purple-800":  "rgba(91, 33, 182, 1)",
		"purple-900":  "rgba(76, 29, 149, 1)",
		"pink-50":     "rgba(253, 242, 248, 1)",
		"pink-100":    "rgba(252, 231, 243, 1)",
		"pink-200":    "rgba(251, 207, 232, 1)",
		"pink-300":    "rgba(249, 168, 212, 1)",
		"pink-400":    "rgba(244, 114, 182, 1)",
		"pink-500":    "rgba(236, 72, 153, 1)",
		"pink-600":    "rgba(219, 39, 119, 1)",
		"pink-700":    "rgba(190, 24, 93, 1)",
		"pink-800":    "rgba(157, 23, 77, 1)",
		"pink-900":    "rgba(131, 24, 67, 1)",
	},
}

var spacing = KeyValues{
	Keys: M{
		"mr": "margin-right",
		"ml": "margin-left",
		"mt": "margin-top",
		"mb": "margin-bottom",
		"mx": Arr{
			"margin-left",
			"margin-right",
		},
		"my": Arr{
			"margin-top",
			"margin-bottom",
		},
		"m":  "margin",
		"pr": "padding-right",
		"pl": "padding-left",
		"pt": "padding-top",
		"pb": "padding-bottom",
		"px": Arr{
			"padding-left",
			"padding-right",
		},
		"py": Arr{
			"padding-top",
			"padding-bottom",
		},
		"p": "padding",
	},
	Values: MS{
		"0":    "0px",
		"1":    "0.25rem",
		"2":    "0.5rem",
		"3":    "0.75rem",
		"4":    "1rem",
		"5":    "1.25rem",
		"6":    "1.5rem",
		"7":    "1.75rem",
		"8":    "2rem",
		"9":    "2.25rem",
		"10":   "2.5rem",
		"11":   "2.75rem",
		"12":   "3rem",
		"14":   "3.5rem",
		"16":   "4rem",
		"20":   "5rem",
		"24":   "6rem",
		"28":   "7rem",
		"32":   "8rem",
		"36":   "9rem",
		"40":   "10rem",
		"44":   "11rem",
		"48":   "12rem",
		"52":   "13rem",
		"56":   "14rem",
		"60":   "15rem",
		"64":   "16rem",
		"72":   "18rem",
		"80":   "20rem",
		"96":   "24rem",
		"auto": "auto",
		"px":   "1px",
		"0.5":  "0.125rem",
		"1.5":  "0.375rem",
		"2.5":  "0.625rem",
		"3.5":  "0.875rem",
	},
}

var radius = KeyValues{
	Keys: M{
		"rounded":   "border-radius",
		"rounded-t": "border-top-radius",
		"rounded-r": "border-right-radius",
		"rounded-l": "border-left-radius",
		"rounded-b": "border-bottom-radius",
		"rounded-tl": Arr{
			"border-top-radius",
			"border-left-radius",
		},
		"rounded-tr": Arr{
			"border-top-radius",
			"border-right-radius",
		},
		"rounded-bl": Arr{
			"border-bottom-radius",
			"border-left-radius",
		},
		"rounded-br": Arr{
			"border-bottom-radius",
			"border-right-radius",
		},
	},
	Values: MS{
		"none": "0px",
		"sm":   "0.125rem",
		"":     "0.25rem",
		"md":   "0.375rem",
		"lg":   "0.5rem",
		"xl":   "0.75rem",
		"2xl":  "1rem",
		"3xl":  "1.5rem",
		"full": "9999px",
	},
}

var borders = KeyValues{
	Keys: M{
		"border":   "border-width",
		"border-l": "border-left-width",
		"border-r": "border-right-width",
		"border-t": "border-top-width",
		"border-b": "border-bottom-width",
	},
	Values: MS{
		"0": "0px",
		"2": "2px",
		"4": "4px",
		"8": "8px",
		"":  "1px",
	},
}

var sizes = KeyValues{
	Keys: M{
		"h":      "height",
		"w":      "width",
		"top":    "top",
		"left":   "left",
		"bottom": "bottom",
		"right":  "right",
		"minh":   "min-height",
		"minw":   "min-width",
		"maxh":   "max-height",
		"maxw":   "max-width",
	},
	Values: MS{
		"auto":  "auto",
		"min":   "min-content",
		"max":   "max-content",
		"0":     "0px",
		"1":     "0.25rem",
		"2":     "0.5rem",
		"3":     "0.75rem",
		"4":     "1rem",
		"5":     "1.25rem",
		"6":     "1.5rem",
		"7":     "1.75rem",
		"8":     "2rem",
		"9":     "2.25rem",
		"10":    "2.5rem",
		"11":    "2.75rem",
		"12":    "3rem",
		"14":    "3.5rem",
		"16":    "4rem",
		"20":    "5rem",
		"24":    "6rem",
		"28":    "7rem",
		"32":    "8rem",
		"36":    "9rem",
		"40":    "10rem",
		"44":    "11rem",
		"48":    "12rem",
		"52":    "13rem",
		"56":    "14rem",
		"60":    "15rem",
		"64":    "16rem",
		"72":    "18rem",
		"80":    "20rem",
		"96":    "24rem",
		"px":    "1px",
		"0.5":   "0.125rem",
		"1.5":   "0.375rem",
		"2.5":   "0.625rem",
		"3.5":   "0.875rem",
		"1/2":   "50%",
		"1/3":   "33.33%",
		"2/3":   "66.66%",
		"1/4":   "25%",
		"2/4":   "50%",
		"3/4":   "75%",
		"1/5":   "20%",
		"2/5":   "40%",
		"3/5":   "60%",
		"4/5":   "80%",
		"1/6":   "16.66%",
		"2/6":   "33.33%",
		"3/6":   "50%",
		"4/6":   "66.66%",
		"5/6":   "83.33%",
		"1/12":  "8.33%",
		"2/12":  "16.66%",
		"3/12":  "25%",
		"4/12":  "33.33%",
		"5/12":  "41.66%",
		"6/12":  "50%",
		"7/12":  "58.33%",
		"8/12":  "66.66%",
		"9/12":  "75%",
		"10/12": "83.33%",
		"11/12": "91.66%",
		"full":  "100%",
	},
}

var twClassLookup = map[string]string{
	"flex":                "display: flex;",
	"inline-flex":         "display: inline-flex;",
	"block":               "display: block;",
	"inline-block":        "display: inline-block;",
	"inline":              "display: inline;",
	"table":               "display: table;",
	"inline-table":        "display: inline-table;",
	"grid":                "display: grid;",
	"inline-grid":         "display: inline-grid;",
	"contents":            "display: contents;",
	"list-item":           "display: list-item;",
	"hidden":              "display: none;",
	"flex-1":              "flex: 1;",
	"flex-row":            "flex-direction: row;",
	"flex-col":            "flex-direction: column;",
	"flex-wrap":           "flex-wrap: wrap;",
	"flex-nowrap":         "flex-wrap: nowrap;",
	"flex-wrap-reverse":   "flex-wrap: wrap-reverse;",
	"items-baseline":      "align-items: baseline;",
	"items-start":         "align-items: flex-start;",
	"items-center":        "align-items: center;",
	"items-end":           "align-items: flex-end;",
	"items-stretch":       "align-items: stretch;",
	"justify-start":       "justify-content: flex-start;",
	"justify-end":         "justify-content: flex-end;",
	"justify-center":      "justify-content: center;",
	"justify-between":     "justify-content: space-between;",
	"justify-around":      "justify-content: space-around;",
	"justify-evenly":      "justify-content: space-evenly;",
	"uppercase":           "text-transform: uppercase",
	"lowercase":           "text-transform: lowercase",
	"capitalize":          "text-transform: capitalize",
	"normal-case":         "text-transform: normal-case",
	"text-left":           "text-align: left;",
	"text-center":         "text-align: center;",
	"text-right":          "text-align: right;",
	"text-justify":        "text-align: justify;",
	"underline":           "text-decoration: underline;",
	"line-through":        "text-decoration: line-through;",
	"no-underline":        "text-decoration: none;",
	"whitespace-normal":   "white-space: normal;",
	"whitespace-nowrap":   "white-space: nowrap;",
	"whitespace-pre":      "white-space: pre;",
	"whitespace-pre-line": "white-space: pre-line;",
	"whitespace-pre-wrap": "white-space: pre-wrap;",
	"break-normal":        "word-break: normal; overflow-wrap: normal;",
	"break-words":         "word-break: break-word;",
	"break-all":           "word-break: break-all;",
	"font-sans":           "font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, \"Segoe UI\", Roboto, \"Helvetica Neue\", Arial, \"Noto Sans\", sans-serif, \"Apple Color Emoji\", \"Segoe UI Emoji\", \"Segoe UI Symbol\", \"Noto Color Emoji\";",
	"font-serif":          "font-family: ui-serif, Georgia, Cambria, \"Times New Roman\", Times, serif;",
	"font-mono":           "font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, \"Liberation Mono\", \"Courier New\", monospace;",
	"font-thin":           "font-weight: 100;",
	"font-extralight":     "font-weight: 200;",
	"font-light":          "font-weight: 300;",
	"font-normal":         "font-weight: 400;",
	"font-medium":         "font-weight: 500;",
	"font-semibold":       "font-weight: 600;",
	"font-bold":           "font-weight: 700;",
	"font-extrabold":      "font-weight: 800;",
	"font-black":          "font-weight: 900;",
	"text-xs":             "font-size: 0.75rem; line-height: 1rem;",
	"text-sm":             "font-size: 0.875rem; line-height: 1.25rem;",
	"text-base":           "font-size: 1rem; line-height: 1.5rem;",
	"text-lg":             "font-size: 1.125rem; line-height: 1.75rem;",
	"text-xl":             "font-size: 1.25rem; line-height: 1.75rem;",
	"text-2xl":            "font-size: 1.5rem; line-height: 2rem;",
	"text-3xl":            "font-size: 1.875rem; line-height: 2.25rem;",
	"text-4xl":            "font-size: 2.25rem; line-height: 2.5rem;",
	"text-5xl":            "font-size: 3rem; line-height: 1;",
	"text-6xl":            "font-size: 3.75rem;; line-height: 1;",
	"text-7xl":            "font-size: 4.5rem; line-height: 1;",
	"text-8xl":            "font-size: 6rem; line-height: 1;",
	"text-9xl":            "font-size: 8rem; line-height: 1;",
	"cursor-auto":         "cursor: auto;",
	"cursor-default":      "cursor: default;",
	"cursor-pointer":      "cursor: pointer;",
	"cursor-wait":         "cursor: wait;",
	"cursor-text":         "cursor: text;",
	"cursor-move":         "cursor: move;",
	"cursor-help":         "cursor: help;",
	"cursor-not-allowed":  "cursor: not-allowed;",
	"pointer-events-none": "pointer-events: none;",
	"pointer-events-auto": "pointer-events: auto;",
	"select-none":         "user-select: none;",
	"select-text":         "user-select: text;",
	"select-all":          "user-select: all;",
	"select-auto":         "user-select: auto;",
	"w-screen":            "100vw",
	"h-screen":            "100vh",
	"static":              "position: static;",
	"fixed":               "position: fixed;",
	"absolute":            "position: absolute;",
	"relative":            "position: relative;",
	"sticky":              "position: sticky;",
	"overflow-auto":       "overflow: auto;",
	"overflow-hidden":     "overflow: hidden;",
	"overflow-visible":    "overflow: visible;",
	"overflow-scroll":     "overflow: scroll;",
	"overflow-x-auto":     "overflow-x: auto;",
	"overflow-y-auto":     "overflow-y: auto;",
	"overflow-x-hidden":   "overflow-x: hidden;",
	"overflow-y-hidden":   "overflow-y: hidden;",
	"overflow-x-visible":  "overflow-x: visible;",
	"overflow-y-visible":  "overflow-y: visible;",
	"overflow-x-scroll":   "overflow-x: scroll;",
	"overflow-y-scroll":   "overflow-y: scroll;",
	"origin-center":       "transform-origin: center;",
	"origin-top":          "transform-origin: top;",
	"origin-top-right":    "transform-origin: top right;",
	"origin-right":        "transform-origin: right;",
	"origin-bottom-right": "transform-origin: bottom right;",
	"origin-bottom":       "transform-origin: bottom;",
	"origin-bottom-left":  "transform-origin: bottom left;",
	"origin-left":         "transform-origin: left;",
	"origin-top-left":     "transform-origin: top left;",
	"shadow-sm":           "box-shadow: 0 0 #0000, 0 0 #0000, 0 1px 2px 0 rgba(0, 0, 0, 0.05);",
	"shadow":              "box-shadow: 0 0 #0000, 0 0 #0000, 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);",
	"shadow-md":           "box-shadow: 0 0 #0000, 0 0 #0000, 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);",
	"shadow-lg":           "box-shadow: 0 0 #0000, 0 0 #0000, 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);",
	"shadow-xl":           "box-shadow: 0 0 #0000, 0 0 #0000, 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);",
	"shadow-2xl":          "box-shadow: 0 0 #0000, 0 0 #0000, 0 25px 50px -12px rgba(0, 0, 0, 0.25);",
	"shadow-inner":        "box-shadow: 0 0 #0000, 0 0 #0000, inset 0 2px 4px 0 rgba(0, 0, 0, 0.06);",
	"shadow-none":         "box-shadow: 0 0 #0000, 0 0 #0000, 0 0 #0000;",
	"ring-inset":          "--tw-ring-inset: insest;",
	"ring-0":              "box-shadow:  0 0 0 calc(0px + 0px) rgba(59, 130, 246, 0.5);",
	"ring-1":              "box-shadow:  0 0 0 calc(1px + 0px) rgba(59, 130, 246, 0.5);",
	"ring-2":              "box-shadow:  0 0 0 calc(2px + 0px) rgba(59, 130, 246, 0.5);",
	"ring-4":              "box-shadow:  0 0 0 calc(4px + 0px) rgba(59, 130, 246, 0.5);",
	"ring-8":              "box-shadow:  0 0 0 calc(8px + 0px) rgba(59, 130, 246, 0.5);",
	"ring":                "box-shadow:  0 0 0 calc(3px + 0px) rgba(59, 130, 246, 0.5);",
}

func init() {
	mapApply(sizes)
	mapApply(spacing)
	mapApply(colors)
	mapApply(borders)
	mapApply(radius)
}

func mapApply(obj KeyValues) {
	for key, v := range obj.Keys {
		for vkey, vv := range obj.Values {
			suffix := ""
			if vkey != "" {
				suffix = "-" + vkey
			}
			className := key + suffix
			if vstring, ok := v.(string); ok {
				twClassLookup[className] = vstring + ": " + vv + ";"
			}
			if varr, ok := v.(Arr); ok {
				for _, kk := range varr {
					twClassLookup[className] = kk.(string) + ": " + vv + ";"
				}
			}
		}
	}
}

var normalizeStyles = `
			*, ::before, ::after { box-sizing: border-box; }
			html { -moz-tab-size: 4; -o-tab-size: 4; tab-size: 4; line-height: 1.15; -webkit-text-size-adjust: 100%; }
			body { margin: 0; font-family: system-ui, -apple-system, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji'; }
			hr { height: 0; color: inherit; }
			abbr[title] { -webkit-text-decoration: underline dotted; text-decoration: underline dotted; }
			b, strong { font-weight: bolder; }
			code, kbd, samp, pre { font-family: ui-monospace, SFMono-Regular, Consolas, 'Liberation Mono', Menlo, monospace; font-size: 1em; }
			small { font-size: 80%; }
			sub, sup { font-size: 75%; line-height: 0; position: relative; vertical-align: baseline; }
			sub { bottom: -0.25em; }
			sup { top: -0.5em; }
			table { text-indent: 0; border-color: inherit; }
			button, input, optgroup, select, textarea { font-family: inherit; font-size: 100%; line-height: 1.15; margin: 0; }
			button, select { text-transform: none; }
			button, [type='button'], [type='reset'], [type='submit'] { -webkit-appearance: button; }
			::-moz-focus-inner { border-style: none; padding: 0; }
			:-moz-focusring { outline: 1px dotted ButtonText; outline: auto; }
			:-moz-ui-invalid { box-shadow: none; }
			legend { padding: 0; }
			progress { vertical-align: baseline; }
			::-webkit-inner-spin-button, ::-webkit-outer-spin-button { height: auto; }
			[type='search'] { -webkit-appearance: textfield; outline-offset: -2px; }
			::-webkit-search-decoration { -webkit-appearance: none; }
			::-webkit-file-upload-button { -webkit-appearance: button; font: inherit; }
			summary { display: list-item; }
			blockquote, dl, dd, h1, h2, h3, h4, h5, h6, hr, figure, p, pre { margin: 0; }
			button { background-color: transparent; background-image: none; }
			fieldset { margin: 0; padding: 0; }
			ol, ul { list-style: none; margin: 0; padding: 0; }
			html { font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"; line-height: 1.5; }
			body { font-family: inherit; line-height: inherit; }
			*, ::before, ::after { box-sizing: border-box; border-width: 0; border-style: solid; border-color: currentColor; }
			hr { border-top-width: 1px; }
			img { border-style: solid; }
			textarea { resize: vertical; }
			input::-moz-placeholder, textarea::-moz-placeholder { opacity: 1; color: #9ca3af; }
			input:-ms-input-placeholder, textarea:-ms-input-placeholder { opacity: 1; color: #9ca3af; }
			input::placeholder, textarea::placeholder { opacity: 1; color: #9ca3af; }
			button, [role="button"] { cursor: pointer; }
			table { border-collapse: collapse; }
			h1, h2, h3, h4, h5, h6 { 	font-size: inherit; 	font-weight: inherit; }
			a { color: inherit; text-decoration: inherit; }
			button, input, optgroup, select, textarea { padding: 0; line-height: inherit; color: inherit; }
			pre, code, kbd, samp { font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace; }
			img, svg, video, canvas, audio, iframe, embed, object { display: block; vertical-align: middle; }
			img, video { max-width: 100%; height: auto; }
			[hidden] { display: none; }
			*, ::before, ::after { --tw-border-opacity: 1; border-color: rgba(229, 231, 235, var(--tw-border-opacity)); }
`
