package routes

import (
	repo "CRM_System/app/storage/db/repository"
	"CRM_System/app/storage/models"
	"fmt"
)

// ----------------------information--------------------------
func Inf_AllGroup() ([]models.EinfGroup, error) {
	result, err := repo.FetchAllGrp()

	if err != nil {
		fmt.Printf("[Routes][Group] - Ошибка при получении всех групп: %v\n", err)
		return nil, err
	}

	fmt.Println("[Routes][Group] - Получение всех групп: успешное")
	return result, nil
}

func Inf_AllGroupAndSubject() ([]models.InFGroupAndSubject, error) {
	Data, err := repo.InfAllGrpWithSubjects()

	if err != nil {
		return nil, err
	}

	return Data, nil
}

func InfGroupId_GroupIdByInfo(course byte, groduates byte, speciality string, groupNum int) (int, error) {
	result, err := repo.GetGroupIDByParams(course, groduates, speciality, groupNum)

	if err != nil {
		fmt.Printf("[Routes][Group] - Ошибка при получении ID группы с параметрами: course=%d, groduates=%d, speciality=%s, groupNum=%d: %v\n", course, groduates, speciality, groupNum, err)
		return 0, err
	} else {
		fmt.Println("[Routes][Group] - Успешное получение ID группы")
	}

	return result, nil
}

// ----------------------Manager--------------------------
func Create_Group(course byte, groduates byte, speciality string, groupNum int) (int, error) {
	Id, err := repo.CrtGrp(course, groduates, speciality, groupNum)

	if err != nil {
		fmt.Printf("[Routes][Group] - Ошибка при создании группы: %v\n", err)
		return 0, err
	} else {
		fmt.Println("[Routes][Group] - Успешное создание группы")
	}

	return Id, nil
}

func Update_GroupById(groupId int, newCourse byte, newGroduates byte, newSpeciality string, newGroupNum int) (bool, error) {
	result, err := repo.UpdateGrp(groupId, newCourse, newGroduates, newSpeciality, newGroupNum)

	if err != nil {
		fmt.Printf("[Routes][Group] - Ошибка при обновлении группы с ID %d: %v\n", groupId, err)
		return false, err
	} else {
		fmt.Printf("[Routes][Group] - Успешное обновление группы с ID %d\n", groupId)
	}

	return result, nil
}

func Delete_GroupById(GroupId int) (bool, error) {
	result, err := repo.DelGrp(GroupId)

	if err != nil {
		fmt.Printf("[Routes][Group] - Ошибка при удалении группы %v\n", err)
		return false, err
	} else {
		fmt.Printf("[Routes][Group] - Успешное удаление группы")
	}

	return result, nil
}

// ----------------------Дублирование группы----------------------
func DublicateGroupAllData(GroupId int) (models.InFGroupAndSubject, error) {
	Group, err := repo.FetchGroupById(GroupId)
	if err != nil {
		return models.InFGroupAndSubject{}, err
	}

	NewNumber, err := repo.MaxNumberByParams(Group.Course, Group.Groudates, Group.Speciality)
	if err != nil {
		return models.InFGroupAndSubject{}, err
	}

	newGroupIds, err := repo.CrtGrp(Group.Course, Group.Groudates, Group.Speciality, NewNumber+1)
	if err != nil {
		return models.InFGroupAndSubject{}, err
	}

	err = repo.CopyDisciplinesBetweenGroups(GroupId, newGroupIds)
	if err != nil {
		return models.InFGroupAndSubject{}, err
	}

	Result, err := repo.FetchGrpWithSubjectsById(newGroupIds)
	if err != nil {
		return models.InFGroupAndSubject{}, err
	}

	return Result, nil
}
