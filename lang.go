package langdetect

// Lang represents a language following ISO 639-3 standard.
type Lang int

const (
	Afr Lang = iota
	Aka
	Amh
	Arb
	Azj
	Bel
	Ben
	Bho
	Bul
	Ceb
	Ces
	Cmn
	Dan
	Deu
	Ell
	Eng
	Epo
	Est
	Fin
	Fra
	Guj
	Hat
	Hau
	Heb
	Hin
	Hrv
	Hun
	Ibo
	Ilo
	Ind
	Ita
	Jav
	Jpn
	Kan
	Kat
	Khm
	Kin
	Kor
	Kur
	Lav
	Lit
	Mai
	Mal
	Mar
	Mkd
	Mlg
	Mya
	Nep
	Nld
	Nno
	Nob
	Nya
	Ori
	Orm
	Pan
	Pes
	Pol
	Por
	Ron
	Run
	Rus
	Sin
	Skr
	Slv
	Sna
	Som
	Spa
	Srp
	Swe
	Tam
	Tel
	Tgl
	Tha
	Tir
	Tuk
	Tur
	Uig
	Ukr
	Urd
	Uzb
	Vie
	Ydd
	Yor
	Zul
)

func (l Lang) Name() string {
	langMap := map[Lang]string{
		Afr: "Afrikaans",
		Aka: "Akan",
		Amh: "Amharic",
		Arb: "Arabic",
		Azj: "Azerbaijani",
		Bel: "Belarusian",
		Ben: "Bengali",
		Bho: "Bhojpuri",
		Bul: "Bulgarian",
		Ceb: "Cebuano",
		Ces: "Czech",
		Cmn: "Mandarin",
		Dan: "Danish",
		Deu: "German",
		Ell: "Greek",
		Eng: "English",
		Epo: "Esperanto",
		Est: "Estonian",
		Fin: "Finnish",
		Fra: "French",
		Guj: "Gujarati",
		Hat: "Haitian Creole",
		Hau: "Hausa",
		Heb: "Hebrew",
		Hin: "Hindi",
		Hrv: "Croatian",
		Hun: "Hungarian",
		Ibo: "Igbo",
		Ilo: "Ilocano",
		Ind: "Indonesian",
		Ita: "Italian",
		Jav: "Javanese",
		Jpn: "Japanese",
		Kan: "Kannada",
		Kat: "Georgian",
		Khm: "Khmer",
		Kin: "Kinyarwanda",
		Kor: "Korean",
		Kur: "Kurdish",
		Lav: "Latvian",
		Lit: "Lithuanian",
		Mai: "Maithili",
		Mal: "Malayalam",
		Mar: "Marathi",
		Mkd: "Macedonian",
		Mlg: "Malagasy",
		Mya: "Burmese",
		Nep: "Nepali",
		Nld: "Dutch",
		Nno: "Nynorsk",
		Nob: "Bokmal",
		Nya: "Chewa",
		Ori: "Oriya",
		Orm: "Oromo",
		Pan: "Punjabi",
		Pes: "Persian",
		Pol: "Polish",
		Por: "Portuguese",
		Ron: "Romanian",
		Run: "Rundi",
		Rus: "Russian",
		Sin: "Sinhalese",
		Skr: "Saraiki",
		Slv: "Slovene",
		Sna: "Shona",
		Som: "Somali",
		Spa: "Spanish",
		Srp: "Serbian",
		Swe: "Swedish",
		Tam: "Tamil",
		Tel: "Telugu",
		Tgl: "Tagalog",
		Tha: "Thai",
		Tir: "Tigrinya",
		Tuk: "Turkmen",
		Tur: "Turkish",
		Uig: "Uyghur",
		Ukr: "Ukrainian",
		Urd: "Urdu",
		Uzb: "Uzbek",
		Vie: "Vietnamese",
		Ydd: "Yiddish",
		Yor: "Yoruba",
		Zul: "Zulu",
	}

	if val, ok := langMap[l]; ok {
		return val
	}

	return ""
}

func (l Lang) ISO6393() string {
	langMap := map[Lang]string{
		Afr: "afr",
		Aka: "aka",
		Amh: "amh",
		Arb: "arb",
		Azj: "azj",
		Bel: "bel",
		Ben: "ben",
		Bho: "bho",
		Bul: "bul",
		Ceb: "ceb",
		Ces: "ces",
		Cmn: "cmn",
		Dan: "dan",
		Deu: "deu",
		Ell: "ell",
		Eng: "eng",
		Epo: "epo",
		Est: "est",
		Fin: "fin",
		Fra: "fra",
		Guj: "guj",
		Hat: "hat",
		Hau: "hau",
		Heb: "heb",
		Hin: "hin",
		Hrv: "hrv",
		Hun: "hun",
		Ibo: "ibo",
		Ilo: "ilo",
		Ind: "ind",
		Ita: "ita",
		Jav: "jav",
		Jpn: "jpn",
		Kan: "kan",
		Kat: "kat",
		Khm: "khm",
		Kin: "kin",
		Kor: "kor",
		Kur: "kur",
		Lav: "lav",
		Lit: "lit",
		Mai: "mai",
		Mal: "mal",
		Mar: "mar",
		Mkd: "mkd",
		Mlg: "mlg",
		Mya: "mya",
		Nep: "nep",
		Nld: "nld",
		Nno: "nno",
		Nob: "nob",
		Nya: "nya",
		Ori: "ori",
		Orm: "orm",
		Pan: "pan",
		Pes: "pes",
		Pol: "pol",
		Por: "por",
		Ron: "ron",
		Run: "run",
		Rus: "rus",
		Sin: "sin",
		Skr: "skr",
		Slv: "slv",
		Sna: "sna",
		Som: "som",
		Spa: "spa",
		Srp: "srp",
		Swe: "swe",
		Tam: "tam",
		Tel: "tel",
		Tgl: "tgl",
		Tha: "tha",
		Tir: "tir",
		Tuk: "tuk",
		Tur: "tur",
		Uig: "uig",
		Ukr: "ukr",
		Urd: "urd",
		Uzb: "uzb",
		Vie: "vie",
		Ydd: "ydd",
		Yor: "yor",
		Zul: "zul",
	}

	if val, ok := langMap[l]; ok {
		return val
	}

	return ""
}

func (l Lang) ISO6391() string {
	langMap := map[Lang]string{
		Afr: "af",
		Aka: "ak",
		Amh: "am",
		Arb: "ar",
		Azj: "az", // Azerbaijani iso 639-3 is aze, iso 639-1 az
		Bel: "be",
		Ben: "bn",
		Bho: "bh",
		Bul: "bg",
		Ceb: "", // No iso 639-1 code
		Ces: "cs",
		Cmn: "zh", // No iso 639-1, but http://www.loc.gov/standards/iso639-2/faq.html#24
		Dan: "da",
		Deu: "de",
		Ell: "el",
		Eng: "en",
		Epo: "eo",
		Est: "et",
		Fin: "fi",
		Fra: "fr",
		Guj: "gu",
		Hat: "ht",
		Hau: "ha",
		Heb: "he",
		Hin: "hi",
		Hrv: "hr",
		Hun: "hu",
		Ibo: "ig",
		Ilo: "", // No iso639-1
		Ind: "id",
		Ita: "it",
		Jav: "jv",
		Jpn: "ja",
		Kan: "kn",
		Kat: "ka",
		Khm: "km",
		Kin: "rw",
		Kor: "ko",
		Kur: "ku",
		Lav: "lv",
		Lit: "lt",
		Mai: "", // No iso639-1
		Mal: "ml",
		Mar: "mr",
		Mkd: "mk",
		Mlg: "mg",
		Mya: "my",
		Nep: "ne",
		Nld: "nl",
		Nno: "nn",
		Nob: "nb",
		Nya: "ny",
		Ori: "or",
		Orm: "om",
		Pan: "pa",
		Pes: "", // No iso639-1
		Pol: "pl",
		Por: "pt",
		Ron: "ro",
		Run: "rn",
		Rus: "ru",
		Sin: "si",
		Skr: "", // No iso639-1
		Slv: "sl",
		Sna: "sn",
		Som: "so",
		Spa: "es",
		Srp: "sr",
		Swe: "sv",
		Tam: "ta",
		Tel: "te",
		Tgl: "tl",
		Tha: "th",
		Tir: "ti",
		Tuk: "tk",
		Tur: "tr",
		Uig: "ug",
		Ukr: "uk",
		Urd: "ur",
		Uzb: "uz",
		Vie: "vi",
		Ydd: "", // No iso639-1
		Yor: "yo",
		Zul: "zu",
	}

	if val, ok := langMap[l]; ok {
		return val
	}

	return ""
}
