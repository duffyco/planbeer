package main

import (
	"encoding/xml"
	"io/ioutil"
	"golang.org/x/net/html/charset"
	"bytes"
)

func RecipeFile ( filename string ) XmlRecipe {
    data, _ := ioutil.ReadFile(filename)

	reader := bytes.NewReader( data )
	decoder := xml.NewDecoder( reader )
	decoder.CharsetReader = charset.NewReaderLabel
 
	recipesXml := &XmlRecipeWrapper{}
	err := decoder.Decode( &recipesXml )
	if (err != nil) {
		panic( err )
	} 

	return recipesXml.Recipe
}

func ConvertToRecipe( id int, rec XmlRecipe ) Recipe {
	var recipe Recipe 

	recipe.Abv = -1
	recipe.ID = id
	recipe.Ibu = -1
	recipe.Kind = 0
	recipe.Name = rec.Name
	recipe.Uri = nil

	return recipe
}

type XmlRecipeWrapper struct {
	Recipe XmlRecipe `xml:"RECIPE"`
}

type XmlRecipe struct {
	Version      int `xml:"VERSION"`
	Type      string `xml:"TYPE"`
	Brewer      string `xml:"BREWER"`
	BatchSize      float32 `xml:"BATCH_SIZE"`
	DisplayBatchSize      string `xml:"DISPLAY_BATCH_SIZE"`
	BoilSize      float32 `xml:"BOIL_SIZE"`
	DisplayBoilSize      string `xml:"DISPLAY_BOIL_SIZE"`
	BoilTime      int `xml:"BOIL_TIME"`
	Efficiency   int `xml:"EFFICIENCY"`
	TasteNotes      string `xml:"TASTE_NOTES"`
	Date      string `xml:"DATE"`
	OG      float32 `xml:"OG"`
	EstOG      float32 `xml:"EST_OG"`
	FG      float32 `xml:"FG"`
	EstFG      float32 `xml:"EST_FG"`
	IBU      int `xml:"IBU"`
	Color      int `xml:"COLOR"`
	EstColor      int `xml:"EST_COLOR"`
	ABV      float32 `xml:"ABV"`
	EstABV      float32 `xml:"EST_ABV"`
	PrimaryAge      int `xml:"PRIMARY_AGE"`
	PrimaryTemp      int `xml:"PRIMARY_TEMP"`
	SecondaryAge      int `xml:"SECONDARY_AGE"`
	SecondaryTemp      int `xml:"SECONDARY_TEMP"`
	FermentationStages      int `xml:"FERMENTATION_STAGES"`
	Style      XmlStyle `xml:"STYLE"`
	Equipment XmlEquipment `xml:"EQUIPMENT"`
	Zymatic  XmlZymatic	`xml:"ZYMATIC"`
	Mash XmlMash `xml:"MASH"`
	Hops      XmlHopWrapper `xml:"HOPS"`
	Waters []XmlWaterWrapper  `xml:"WATERS"`
	Fermentables []XmlFermentableWrapper `xml:"FERMENTABLES"`
	Yeasts []XmlYeastWrapper `xml:"YEASTS"`
	KegSmartTag FermSteps `xml:"KEGSMART"` 
	Name      string `xml:"NAME"`
}

type XmlStyle struct {
	CategoryNumber      int `xml:"CATEGORY_NUMBER"`
	StyleLetter      string `xml:"STYLE_LETTER"`
	Name      string `xml:"NAME"`
	Version      int `xml:"VERSION"`
	StyleGuide      string `xml:"STYLE_GUIDE"`
	OGMin      float32 `xml:"OG_MIN"`
	OGMax      float32 `xml:"OG_MAX"`
	FGMin      float32 `xml:"FG_MIN"`
	FBMax      float32 `xml:"FG_MAX"`
	IBUMin      int `xml:"IBU_MIN"`
	IBUMax      int `xml:"IBU_MAX"`
	ColorMin      int `xml:"COLOR_MIN"`
	ColorMax      int `xml:"COLOR_MAX"`
}

type XmlEquipment struct {
	Name      string `xml:"NAME"`
	Version      int `xml:"VERSION"`
}

type XmlZymatic struct {
	MashTemp      int `xml:"MASH_TEMP"`
	MashTime      int `xml:"MASH_TIME"`
	BoilTemp      int `xml:"BOIL_TEMP"`
	Steps		 []XmlZymaticStep `xml:"STEP"`
}

type XmlZymaticStep struct {
	Name      string `xml:"NAME"`
	Temp      int `xml:"TEMP"`
	Time      int `xml:"TIME"`
	Location      string `xml:"LOCATION"`
	Drain      int `xml:"DRAIN"`
}

type XmlMash struct {
	Name      string `xml:"NAME"`
	Version      int `xml:"VERSION"`
	GrainTemp      float32 `xml:"GRAIN_TEMP"`
	Steps []XmlMashWrapper `xml:"MASH_STEPS"`
}

type XmlMashStep struct {
	Name      string `xml:"NAME"`
	Version      int `xml:"VERSION"`
	StepTemp      int `xml:"STEP_TEMP"`
	StepTime      int `xml:"STEP_TIME"`
}

type XmlMashWrapper struct {
	Step XmlMashStep `xml:"MASH_STEP"`
}

type XmlHopWrapper struct {
	Hop XmlHop `xml:"HOP"`
}

type XmlHop struct {
	Name      string `xml:"NAME"`
	Version      int `xml:"VERSION"`
	Alpha      float32 `xml:"ALPHA"`
	Amount      float32 `xml:"AMOUNT"`
	Use      string `xml:"USE"`
	Time      int `xml:"TIME"`
}

type XmlWaterWrapper struct {
	Water XmlWater `xml:"WATER"`
}

type XmlWater struct {
	Version      int `xml:"VERSION"`
	Amount      float32 `xml:"AMOUNT"`
}

type XmlFermentableWrapper struct {
	Fermentable XmlFermentable `xml:"FERMENTABLE"`
}

type XmlFermentable struct {
	Name      string `xml:"NAME"`
	Version      int `xml:"VERSION"`
	Amount      float32 `xml:"AMOUNT"`
	Type      string `xml:"TYPE"`
	Yield      float32 `xml:"YIELD"`
	Color      float32 `xml:"COLOR"`
}


type XmlYeastWrapper struct {
	Yeast XmlYeast `xml:"YEAST"`
}

type XmlYeast struct {
	Name      string `xml:"NAME"`
	Version      int `xml:"VERSION"`
	Amount      int `xml:"AMOUNT"`
	Form      string `xml:"FORM"`
	Lab      string `xml:"LABORATORY"`
	ProductID  string `xml:"PRODUCT_ID"`
	MinTemp      int  `xml:"MIN_TEMPERATURE"`
	MaxTemp      int  `xml:"MAX_TEMPERATURE"`
	Attenuation      int  `xml:"ATTENUATION"`
}

type FermSteps struct {
	Steps []FermStepWrapper `xml:"STEPS"`
}

type FermStepWrapper struct {
	Step XmlFermStep `xml:"STEP"`
}

type XmlFermStep struct {
	Number      int `xml:"NUMBER"`
	Name      string `xml:"NAME"`
	Time      int  `xml:"TIME"`
	Temp      int  `xml:"TEMP"`
}
