package game

type Menu interface {
	GetId() int // Get the ID of the menu
	GetPrevMenu() Menu // Get the previous menu of the menu
	SetPrevMenu(Menu) // Set the previous menu of the menu
	DisplayChoices(*Game) // Display the choices for the player
	Choice(*Game, int) any // Use the input of the player to make a choice
	Quit(*Game) // Quit the menu
}

func InitMenus(game *Game) {
	Menus = make(map[int]Menu, 0)

	Menus[INVENTORY] = InitInvMenu(game) 
	Menus[MAIN] = InitMainMenu(game)
	Menus[SHOPKEEPER] = InitNpcMenu(Npcs["Shopkeeper"], SHOPKEEPER)
	Menus[BLACKSMITH] = InitNpcMenu(Npcs["Blacksmith"], BLACKSMITH)
	Menus[SKILL_TREE] = InitSkillMenu(game)
}
