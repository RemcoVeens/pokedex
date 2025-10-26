package models

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}
type EncounterVersionDetails struct {
	Rate    int32            `json:"rate"`
	Version NamedAPIResource `json:"version"`
}
type EncounterMethodRate struct {
	Encounter_method NamedAPIResource          `json:"encounter_method"`
	Version_details  []EncounterVersionDetails `json:"version_detail"`
}
type EncounterMethod struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Order int32  `json:"order"`
	Names []Name `json:"names"`
}
type Encounter struct {
	Min_level        int32              `json:"min_level"`
	Max_level        int32              `json:"max_level"`
	Condition_values []NamedAPIResource `json:"condition_values"`
	Chance           int32              `json:"chance"`
	Method           NamedAPIResource   `json:"method"`
}
type VersionEncounterDetail struct {
	Version           NamedAPIResource `json:"version"`
	max_chance        int32            `json:"max_chance"`
	Encounter_details []Encounter      `json:"encounter_details"`
}

type PokemonEncounter struct {
	Pokemon         NamedAPIResource         `json:"pokemon"`
	Version_details []VersionEncounterDetail `json:"version_details"`
}
type LocationAreas struct {
	Id                     int32                 `json:"id"`
	Name                   string                `json:"name"`
	Game_index             int32                 `json:"game_index"`
	Encounter_method_rates []EncounterMethodRate `json:"encounter_method_rates"`
	Location               NamedAPIResource      `json:"location"`
	Names                  []Name                `json:"names"`
	Pokemon_encounters     []PokemonEncounter    `json:"pokemon_encounters"`
}
type Pokemon struct {
	ID                     int                    `json:"id"`
	Name                   string                 `json:"name"`
	BaseExperience         int                    `json:"base_experience"`
	Height                 int                    `json:"height"`
	IsDefault              bool                   `json:"is_default"`
	Order                  int                    `json:"order"`
	Weight                 int                    `json:"weight"`
	Abilities              []AbilityContainer     `json:"abilities"`
	Forms                  []NamedAPIResource     `json:"forms"`
	GameIndices            []GameIndex            `json:"game_indices"`
	HeldItems              []HeldItemContainer    `json:"held_items"`
	LocationAreaEncounters string                 `json:"location_area_encounters"`
	Moves                  []MoveContainer        `json:"moves"`
	Species                NamedAPIResource       `json:"species"`
	Sprites                Sprites                `json:"sprites"`
	Cries                  Cries                  `json:"cries"`
	Stats                  []StatContainer        `json:"stats"`
	Types                  []TypeContainer        `json:"types"`
	PastTypes              []PastTypeContainer    `json:"past_types"`
	PastAbilities          []PastAbilityContainer `json:"past_abilities"`
}

// AbilityContainer wraps the ability resource and related data.
type AbilityContainer struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

// GameIndex links a Pokemon to its index in a specific game version.
type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

// HeldItemContainer wraps the held item resource and version details.
type HeldItemContainer struct {
	Item           NamedAPIResource        `json:"item"`
	VersionDetails []HeldItemVersionDetail `json:"version_details"`
}

// HeldItemVersionDetail provides rarity information for an item in a specific version.
type HeldItemVersionDetail struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

// MoveContainer wraps the move resource and version group details.
type MoveContainer struct {
	Move                NamedAPIResource         `json:"move"`
	VersionGroupDetails []MoveVersionGroupDetail `json:"version_group_details"`
}

// MoveVersionGroupDetail provides details on how a move is learned.
type MoveVersionGroupDetail struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	Order           int              `json:"order"`
}

// Sprites contains various sprite URLs for the Pokemon.
type Sprites struct {
	BackDefault      string       `json:"back_default"`
	BackFemale       *string      `json:"back_female"`
	BackShiny        string       `json:"back_shiny"`
	BackShinyFemale  *string      `json:"back_shiny_female"`
	FrontDefault     string       `json:"front_default"`
	FrontFemale      *string      `json:"front_female"`
	FrontShiny       string       `json:"front_shiny"`
	FrontShinyFemale *string      `json:"front_shiny_female"`
	Other            OtherSprites `json:"other"`
	Versions         Versions     `json:"versions"`
}

// OtherSprites contains grouped sprites for specific styles.
type OtherSprites struct {
	DreamWorld      DreamWorldSprites      `json:"dream_world"`
	Home            HomeSprites            `json:"home"`
	OfficialArtwork OfficialArtworkSprites `json:"official-artwork"`
	Showdown        ShowdownSprites        `json:"showdown"`
}

// DreamWorldSprites holds dream world style sprites.
type DreamWorldSprites struct {
	FrontDefault string  `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

// HomeSprites holds home style sprites.
type HomeSprites struct {
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// OfficialArtworkSprites holds official artwork style sprites.
type OfficialArtworkSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// ShowdownSprites holds showdown style sprites (often GIF format).
type ShowdownSprites struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// Versions contains sprites grouped by generation.
type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationII   GenerationII   `json:"generation-ii"`
	GenerationIII  GenerationIII  `json:"generation-iii"`
	GenerationIV   GenerationIV   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVI   GenerationVI   `json:"generation-vi"`
	GenerationVII  GenerationVII  `json:"generation-vii"`
	GenerationVIII GenerationVIII `json:"generation-viii"`
}

// --- Generation-Specific Sprite Structs (Example for Gen I) ---
// Note: Many of these sub-structs follow a similar pattern.

// GenerationI holds all Generation I version sprites.
type GenerationI struct {
	RedBlue RedBlueSprites `json:"red-blue"`
	Yellow  YellowSprites  `json:"yellow"`
}

// RedBlueSprites holds Red/Blue version sprites.
type RedBlueSprites struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

// YellowSprites holds Yellow version sprites.
type YellowSprites struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

// ... other Generation structs (GenerationII, GenerationIII, etc.) would be defined similarly ...

// Cries contains the URLs for the Pokemon's cries.
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

// StatContainer wraps a stat resource and related data.
type StatContainer struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

// TypeContainer wraps a type resource and slot.
type TypeContainer struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

// PastTypeContainer includes type information for a past generation.
type PastTypeContainer struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []TypeContainer  `json:"types"`
}

// PastAbilityContainer includes ability information for a past generation.
type PastAbilityContainer struct {
	Generation NamedAPIResource    `json:"generation"`
	Abilities  []PastAbilityDetail `json:"abilities"`
}

// PastAbilityDetail provides details for an ability in a past generation.
type PastAbilityDetail struct {
	Ability  *NamedAPIResource `json:"ability"`
	IsHidden bool              `json:"is_hidden"`
	Slot     int               `json:"slot"`
}
