package models

type NamedAPIResource struct {
	Name string "json:name"
	Url  string "json:url"
}
type Name struct {
	Name     string           "json:name"
	Language NamedAPIResource "json:language"
}
type EncounterVersionDetails struct {
	Rate    int32            "json:rate"
	Version NamedAPIResource "json:version"
}
type EncounterMethodRate struct {
	Encounter_method NamedAPIResource          "json:encounter_method"
	Version_details  []EncounterVersionDetails "json:version_detail"
}
type EncounterMethod struct {
	Id    int32  "json:id"
	Name  string "json:name"
	Order int32  "json:order"
	Names []Name "json:names"
}
type Encounter struct {
	Min_level        int32              "json:min_level"
	Max_level        int32              "json:max_level"
	Condition_values []NamedAPIResource "json:condition_values"
	Chance           int32              "json:chance"
	Method           NamedAPIResource   "json:method"
}
type VersionEncounterDetail struct {
	Version           NamedAPIResource "json:version"
	max_chance        int32            "json:max_chance"
	Encounter_details []Encounter      "json:encounter_details"
}

type PokemonEncounter struct {
	Pokemon         NamedAPIResource         "json:pokemon"
	Version_details []VersionEncounterDetail "json:version_details"
}
type LocationAreas struct {
	Id                     int32                 "json:id"
	Name                   string                "json:name"
	Game_index             int32                 "json:game_index"
	Encounter_method_rates []EncounterMethodRate "json:encounter_method_rates"
	Location               NamedAPIResource      "json:location"
	Names                  []Name                "json:names"
	Pokemon_encounters     []PokemonEncounter    "json:pokemon_encounters"
}
