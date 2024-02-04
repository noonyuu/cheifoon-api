package model

import "gorm.io/gorm"

type Machine struct {
	gorm.Model
	RoomId        int        `json:"room_id" gorm:"type:int;not null"`
	Room          Room       `json:"-" gorm:"foreignKey:RoomId"`
	RecipeId      int        `json:"recipe_id" gorm:"type:int;not null"`
	Recipe        Recipe     `json:"-" gorm:"foreignKey:RecipeId"`
	RecipeName    string     `json:"recipe_name" gorm:"-"`
	SeasoningName string     `json:"seasoning_name" gorm:"-"`
	SeasoningStatus int   `json:"seasoning_status" gorm"-"`
	TableSpoon    int        `json:"table_spoon" gorm:"-"`
	TeaSpoon      int        `json:"tea_spoon" gorm:"-"`
}

func CreateMachine(machine *Machine) error {
	result := db.Create(machine)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// recipeの名前のみ取得
func FindMachineRecipe(roomId int) (map[int]string, error) {
    // レシピIDとレシピ名を取得するクエリを実行
    rows, err := db.Raw("SELECT machines.recipe_id, recipes.recipe_name FROM machines "+
        "LEFT JOIN recipes ON machines.recipe_id = recipes.id "+
        "WHERE machines.room_id = ?", roomId).Rows()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // レシピIDをキーとし、レシピ名を値とするマップを初期化
    recipeIDToName := make(map[int]string)

    // 各行を反復処理し、マップにレシピIDとレシピ名を格納
    for rows.Next() {
        var recipeID int
        var recipeName string
        if err := rows.Scan(&recipeID, &recipeName); err != nil {
            return nil, err
        }
        recipeIDToName[recipeID] = recipeName
    }

    return recipeIDToName, nil
}