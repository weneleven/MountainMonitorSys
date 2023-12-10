package dao

import (
	"mountain/global"
	"mountain/internal/model"
)

func CreateProject(project *model.Project) error {
	result := global.DBEngine.Create(project)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func DeleteProjectAndSensors(id uint) error {
//     // 开始数据库事务
//     tx := global.DBEngine.Begin()

//     // 删除项目
//     result := tx.Where("id = ?", id).Delete(&model.Project{})
//     if result.Error != nil {
//         // 回滚事务并返回错误
//         tx.Rollback()
//         return result.Error
//     }

//     // 删除与项目关联的传感器数据
//     result = tx.Where("project_id = ?", id).Delete(&model.Sensor{})
//     if result.Error != nil {
//         // 回滚事务并返回错误
//         tx.Rollback()
//         return result.Error
//     }

//     // 提交事务
//     tx.Commit()

//     return nil
// }

func GetProjects() ([]model.Project, error) {
	var projects []model.Project

	result := global.DBEngine.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

func GetProjectByID(id uint) (*model.Project, error) {
	var project model.Project
	result := global.DBEngine.Where("id = ?", id).First(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}

func UpdateProject(project *model.Project) error {
	result := global.DBEngine.Save(project)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
