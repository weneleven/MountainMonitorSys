package dao

import (
	"fmt"
	"mountain/global"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
)

func CreateProject(project *model.Project) int {
	result := global.DBEngine.Create(project)
	if result.Error != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

func DeleteProject(id int) int {
	// 开始数据库事务
	tx := global.DBEngine.Begin()
	// 删除项目
	result := tx.Where("id = ?", id).Delete(&model.Project{})
	if result.Error != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return errmessage.ERROR
	}
	tx.Commit()
	return errmessage.SUCCESS
}


func GetProjects(pageSize int, pageNum int) ([]model.Project, int64) {
	var projects []model.Project
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := global.DBEngine.Limit(pageSize).Offset(offset).Find(&projects).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return projects, total
}
func GetProjectByID(id int) (*model.Project, int) {
	var project model.Project
	result := global.DBEngine.Where("id = ?", id).First(&project)
	if result.Error != nil {
		return nil, errmessage.ERROR
	}

	return &project, errmessage.SUCCESS
}
func GetProjectByName(name string) (*model.Project, int){
	var project model.Project
	fmt.Println(name)
	//result := global.DBEngine.Where("project_name = ?", name).First(&project)
	// 使用 LIKE 进行模糊查询
	result := global.DBEngine.Where("project_name LIKE ?", "%" + name + "%").Find(&project)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, errmessage.ERROR
	}
	return &project, errmessage.SUCCESS
}


func UpdateProject(project *model.Project,id int) int {
	result := global.DBEngine.Model(&model.Project{}).Where("id = ?", id).Updates(project)
	//result := global.DBEngine.Save(project)
	if result.Error != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 删除项目下传感器
func DeleteProjectSensor(id int) int {
	var sensor model.Sensor
	var count int64
	// 首先检查是否存在匹配的记录,即使不存在,也返回200
	global.DBEngine.Model(&sensor).Where("project_id = ?", id).Count(&count)
	if count == 0 {
		return errmessage.SUCCESS
	}
	// 存在匹配的记录，执行删除操作
	result := global.DBEngine.Where("project_id = ?", id).Delete(&sensor)
	if result.Error != nil {
		return errmessage.ERROR
	}

	return errmessage.SUCCESS
}