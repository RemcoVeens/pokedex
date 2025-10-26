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
	MaxChance         int32            `json:"max_chance"`
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
	PastTypes              []PastTypeContainer    `json:"past_types"`
	PastAbilities          []PastAbilityContainer `json:"past_abilities"`
	Sprites                Sprites                `json:"sprites"`
	Cries                  Cries                  `json:"cries"`
	Species                NamedAPIResource       `json:"species"`
	Stats                  []PokemonStat          `json:"stats"`
	Types                  []PokemonType          `json:"types"`
}
type AbilityContainer struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}
type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}
type HeldItemContainer struct {
	Item           NamedAPIResource        `json:"item"`
	VersionDetails []HeldItemVersionDetail `json:"version_details"`
}
type HeldItemVersionDetail struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}
type MoveContainer struct {
	Move                NamedAPIResource         `json:"move"`
	VersionGroupDetails []MoveVersionGroupDetail `json:"version_group_details"`
}
type MoveVersionGroupDetail struct {
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	LevelLearnedAt  int              `json:"level_learned_at"`
	Order           int              `json:"order"`
}
type Sprites struct {
	FrontDefault     string  `json:"front_default"`
	FrontShiny       string  `json:"front_shiny"`
	FrontFemale      *string `json:"front_female"`
	FrontShinyFemale *string `json:"front_shiny_female"`
	BackDefault      string  `json:"back_default"`
	BackShiny        string  `json:"back_shiny"`
	BackFemale       *string `json:"back_female"`
	BackShinyFemale  *string `json:"back_shiny_female"`
}

// Cries contains the URLs for the Pokemon's cries.
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

// StatContainer wraps a stat resource and related data.
type PokemonStat struct {
	Stat     NamedAPIResource `json:"stat"`
	Effort   int              `json:"effort"`
	BaseStat int              `json:"base_stat"`
}

// TypeContainer wraps a type resource and slot.
type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

// PastTypeContainer includes type information for a past generation.
type PastTypeContainer struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []PokemonType    `json:"types"`
}

// PastAbilityContainer includes ability information for a past generation.
type PastAbilityContainer struct {
	Generation NamedAPIResource `json:"generation"`
	Abilities  []PokemonAbility `json:"abilities"`
}

type PokemonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int32            `json:"slot"`
	Ability  NamedAPIResource `json:"ablilty"`
}
