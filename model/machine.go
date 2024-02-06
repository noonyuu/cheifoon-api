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
	SeasoningStatus int   	 `json:"seasoning_status" gorm"-"`
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

func FindMachine(roomId int, recipeId int) (map[uint][]Machine, error) {
    rows, err := db.Raw("SELECT machines.id, machines.room_id, machines.recipe_id, recipes.recipe_name, admin_seasonings.seasoning_name, seasoning_statuses.id, recipe_infos.table_spoon, recipe_infos.tea_spoon FROM machines "+
        "LEFT JOIN recipes ON machines.recipe_id = recipes.id "+
        "LEFT JOIN recipe_infos ON machines.recipe_id = recipe_infos.recipe_id "+
        "JOIN admin_seasonings ON admin_seasonings.id = recipe_infos.seasoning_id "+
				"JOIN seasoning_statuses ON seasoning_statuses.id = admin_seasonings.status_id "+
        "WHERE machines.room_id = ? AND machines.recipe_id = ?", roomId,recipeId).Rows()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

		// IDをキーとしたマップ
    machineMap := make(map[uint][]Machine) 

    for rows.Next() {
        var machine Machine
        err := rows.Scan(&machine.ID, &machine.RoomId, &machine.RecipeId, &machine.RecipeName, &machine.SeasoningName,&machine.SeasoningStatus, &machine.TableSpoon, &machine.TeaSpoon)
        if err != nil {
            return nil, err
        }
				// machineをIDをキーとしてマップに追加
        machineMap[machine.ID] = append(machineMap[machine.ID], machine) 
    }

    return machineMap, nil
}

// recipeの名前のみ取得
func FindMachineRecipe(roomId int) (map[int]string, error) {
    // レシピIDとレシピ名を取得するクエリを実行
    rows, err := db.Raw("SELECT machines.recipe_id, recipes.recipe_name FROM machines "+
        "LEFT JOIN recipes ON machines.recipe_id = recipes.id "+
        "WHERE machines.room_id = ? AND machines.deleted_at IS NULL", roomId).Rows()
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

func FindMachineMobile(roomId int) ([]map[string]interface{}, error) {
    rows, err := db.Raw("SELECT machines.recipe_id, recipes.recipe_name, recipes.recipe_image FROM machines "+
        "LEFT JOIN recipes ON machines.recipe_id = recipes.id "+
        "WHERE machines.room_id = ? AND machines.deleted_at IS NULL", roomId).Rows()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var machines []map[string]interface{}

    // 各行を反復処理し、機械の情報を取得し、mapに追加
    for rows.Next() {
        var recipeID int
        var recipeName, recipeImage string
        if err := rows.Scan(&recipeID, &recipeName, &recipeImage); err != nil {
            return nil, err
        }
        machine := map[string]interface{}{
            "ID":           recipeID,
            "recipe_name":  recipeName,
            "recipe_image": recipeImage,
        }
        machines = append(machines, machine)
    }

    return machines, nil
}

func DeleteRecipe(recipeId int) (error) {
	err := db.Where("recipe_id = ?", recipeId).
			Delete(&Machine{}).Error

	if err != nil {
		return err
	}
	return nil
}