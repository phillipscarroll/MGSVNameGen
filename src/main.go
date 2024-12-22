package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {

	var firstName [98]string
	var lastName [76]string

	//The following names are inspired by Metal Gear Solid 5: The Phantom Pain

	firstName[0] = "Armored"
	firstName[1] = "Bastard"
	firstName[2] = "Biting"
	firstName[3] = "Bitter"
	firstName[4] = "Black"
	firstName[5] = "Blazing"
	firstName[6] = "Bloody"
	firstName[7] = "Brass"
	firstName[8] = "Brutal"
	firstName[9] = "Bullet"
	firstName[10] = "Cannibal"
	firstName[11] = "Charging"
	firstName[12] = "Code"
	firstName[13] = "Copper"
	firstName[14] = "Crawling"
	firstName[15] = "Creeping"
	firstName[16] = "Crimson"
	firstName[17] = "Crying"
	firstName[18] = "Crystal"
	firstName[19] = "Cunning"
	firstName[20] = "Dark"
	firstName[21] = "Death"
	firstName[22] = "Devil"
	firstName[23] = "Depressed"
	firstName[24] = "Dire"
	firstName[25] = "Dizzy"
	firstName[26] = "Doom"
	firstName[27] = "Exotic"
	firstName[28] = "Evil"
	firstName[29] = "Fire"
	firstName[30] = "Flaming"
	firstName[31] = "Frantic"
	firstName[32] = "Frigid"
	firstName[33] = "Garnet"
	firstName[34] = "Glacier"
	firstName[35] = "Golden"
	firstName[36] = "Gray"
	firstName[37] = "Greedy"
	firstName[38] = "Green"
	firstName[39] = "Grizzly"
	firstName[40] = "Growling"
	firstName[41] = "Hinting"
	firstName[42] = "Hissing"
	firstName[43] = "Howling"
	firstName[44] = "Hulking"
	firstName[45] = "Hungry"
	firstName[46] = "Hunting"
	firstName[47] = "Iron"
	firstName[48] = "Jade"
	firstName[49] = "Killer"
	firstName[50] = "Komodo"
	firstName[51] = "Lonely"
	firstName[52] = "Mad"
	firstName[53] = "Midnight"
	firstName[54] = "Night"
	firstName[55] = "Northern"
	firstName[56] = "Ochre"
	firstName[57] = "Panzer"
	firstName[58] = "Pirate"
	firstName[59] = "Poison"
	firstName[60] = "Poucing"
	firstName[61] = "Pouncing"
	firstName[62] = "Prowling"
	firstName[63] = "Punching"
	firstName[64] = "Rabid"
	firstName[65] = "Raging"
	firstName[66] = "Rampant"
	firstName[67] = "Rancid"
	firstName[68] = "Raving"
	firstName[69] = "Razor"
	firstName[70] = "Roaring"
	firstName[71] = "Rumble"
	firstName[72] = "Running"
	firstName[73] = "Sadistic"
	firstName[74] = "Scowling"
	firstName[75] = "Seething"
	firstName[76] = "Sexy"
	firstName[77] = "Shining"
	firstName[78] = "Silent"
	firstName[79] = "Sinister"
	firstName[80] = "Sly"
	firstName[81] = "Smoking"
	firstName[82] = "Solid"
	firstName[83] = "Spitting"
	firstName[84] = "Spying"
	firstName[85] = "Stalking"
	firstName[86] = "Steel"
	firstName[87] = "Stiker"
	firstName[88] = "Stone"
	firstName[89] = "Stubborn"
	firstName[90] = "Stupid"
	firstName[91] = "Sunny"
	firstName[92] = "Thunder"
	firstName[93] = "Titanium"
	firstName[94] = "Vampire"
	firstName[95] = "Vengeful"
	firstName[96] = "Vile"
	firstName[97] = "Wild"

	lastName[0] = "Armadillo"
	lastName[1] = "Axolotl"
	lastName[2] = "Badger"
	lastName[3] = "Barracuda"
	lastName[4] = "Basilisk"
	lastName[5] = "Bat"
	lastName[6] = "Bear"
	lastName[7] = "Beetle"
	lastName[8] = "Bison"
	lastName[9] = "Boa"
	lastName[10] = "Buffalo"
	lastName[11] = "Bull"
	lastName[12] = "Canine"
	lastName[13] = "Capybara"
	lastName[14] = "Cat"
	lastName[15] = "Centipede"
	lastName[16] = "Chameleon"
	lastName[17] = "Cobra"
	lastName[18] = "Crab"
	lastName[19] = "Crow"
	lastName[20] = "Dingo"
	lastName[21] = "Dragon"
	lastName[22] = "Eagle"
	lastName[23] = "Eel"
	lastName[24] = "Gator"
	lastName[25] = "Gecko"
	lastName[26] = "Gibbon"
	lastName[27] = "Goat"
	lastName[28] = "Harrier"
	lastName[29] = "Hawk"
	lastName[30] = "Hippo"
	lastName[31] = "Hound"
	lastName[32] = "Husky"
	lastName[33] = "Hyena"
	lastName[34] = "Iguana"
	lastName[35] = "Jackal"
	lastName[36] = "Kangaroo"
	lastName[37] = "Lion"
	lastName[38] = "Mammoth"
	lastName[39] = "Mantis"
	lastName[40] = "Mastodon"
	lastName[41] = "Mongoose"
	lastName[42] = "Moose"
	lastName[43] = "Mosquito"
	lastName[44] = "Moth"
	lastName[45] = "Mustang"
	lastName[46] = "Ocelot"
	lastName[47] = "Octopus"
	lastName[48] = "Ox"
	lastName[49] = "Panther"
	lastName[50] = "Parrot"
	lastName[51] = "Platypus"
	lastName[52] = "Python"
	lastName[53] = "Ram"
	lastName[54] = "Raptor"
	lastName[55] = "Raven"
	lastName[56] = "Rhino"
	lastName[57] = "Roach"
	lastName[58] = "Rooster"
	lastName[59] = "Serpent"
	lastName[60] = "Snake"
	lastName[61] = "Shark"
	lastName[62] = "Sloth"
	lastName[63] = "Slug"
	lastName[64] = "Squirrel"
	lastName[65] = "Stallion"
	lastName[66] = "Tarantula"
	lastName[67] = "Tiger"
	lastName[68] = "Tree Frog"
	lastName[69] = "Viper"
	lastName[70] = "Vulture"
	lastName[71] = "Wallaby"
	lastName[72] = "Wasp"
	lastName[73] = "Whale"
	lastName[74] = "Worm"
	lastName[75] = "Wolf"

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	first := rand.Intn(98)
	last := rand.Intn(76)

	fmt.Println(firstName[first] + " " + lastName[last])
}