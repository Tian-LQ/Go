package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"default:Mr.T"`
	Age  uint8  `gorm:"comment:年龄"`
}

func AutoMigrate() {
	GlobalDB.AutoMigrate(&User{})
}

func CreateRecord() {
	// 1.根据数据指针创建
	user1 := User{
		Name: "Mr.X",
		Age:  18,
	}
	ret := GlobalDB.Create(&user1)
	if ret.Error != nil {
		fmt.Println(ret.Error)
	}
	fmt.Printf("create record user1: %+v\n", user1)

	// 2.用指定的字段创建记录
	user2 := User{
		Name: "Mr.T",
		Age:  24,
	}
	// Select方法表示只使用user2当中的Name字段去创建记录
	// 对应的Omit方法表示使用user2当中除了Name以外的字段去创建记录
	ret = GlobalDB.Select("Name").Create(&user2)
	if ret.Error != nil {
		fmt.Println(ret.Error)
	}
	fmt.Printf("create record user2: %+v\n", user2)

	// 3.使用切片批量创建记录
	users := []User{{Name: "tianlq", Age: 25}, {Name: "yuwd", Age: 24}, {Name: "maw", Age: 30}}
	ret = GlobalDB.Create(&users)
	if ret.Error != nil {
		fmt.Println(ret.Error)
	}
	fmt.Printf("batch create records users: %+v\n", users)

	// 4.使用Map创建记录[批量]
	ret = GlobalDB.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "wangrs", "Age": 24},
		{"Name": "shuq", "Age": 23},
	})
	if ret.Error != nil {
		fmt.Println(ret.Error)
	}
}

func FindRecord() {
	user := User{}
	// 1.获取第一条记录（主键升序）
	ret := GlobalDB.First(&user)
	if ret.Error != nil {
		if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
			fmt.Println("record not found")
		} else {
			fmt.Println(ret.Error)
		}
	}
	fmt.Printf("find first record order by ID: %+v\n", user)
	result := map[string]interface{}{}
	// 1> Model方法指定model
	GlobalDB.Model(&User{}).First(&result)
	fmt.Printf("find first record order by ID: %+v\n", result)

	// 2.获取一条记录，没有指定排序字段
	GlobalDB.Take(&user)
	fmt.Printf("find first record: %+v\n", user)

	// 3.获取最后一条记录（主键降序）
	GlobalDB.Last(&user)
	fmt.Printf("find first record order by ID desc: %+v\n", user)

	// 4.主键检索
	users := make([]User, 0, 3)
	GlobalDB.Find(&users, []int{1, 2, 3})
	fmt.Printf("find record by primary key: %+v\n", users)

	// 5.条件检索
	// 1> String 条件
	GlobalDB.Where("name = ?", "Mr.T").Find(&users)
	fmt.Printf("find records where name = 'Mr.T': %+v\n", users)
	// 2> Struct & Map 条件
	// 当使用结构作为条件查询时，GORM 只会查询非零值字段
	GlobalDB.Where(&User{Name: "Mr.T"}).Find(&users)
	fmt.Printf("find records where name = 'Mr.T': %+v\n", users)
	// 如果想要包含零值查询条件，你可以使用 map，其会包含所有 key-value 的查询条件
	GlobalDB.Where(map[string]interface{}{"Name": "Mr.T"}).Find(&users)
	fmt.Printf("find records where name = 'Mr.T': %+v\n", users)
	GlobalDB.Where([]uint{1, 2, 3}).Find(&users)
	fmt.Printf("find records where id in (1, 2, 3): %+v\n", users)
}

func UpdateRecord() {
	// 1.使用 Save 保存所有字段
	// 1> 按照主键进行更新，如果主键为空，此时 Save 会创建新的 record
	users := make([]User, 0)
	GlobalDB.Where(&User{Age: 24}).Find(&users)
	for i, _ := range users {
		users[i].Age += 1
	}
	if len(users) != 0 {
		GlobalDB.Save(users)
	}

	// 2.使用 Update 更新单列
	GlobalDB.Model(&User{}).Where("name like ?", "%Mr%").Update("age", 15)

	// 3.使用 Updates 更新多列
	// 1> 当通过 struct 更新时，GORM 只会更新非零字段
	// 2> 可以通过 map 来完成零值字段的更新
	GlobalDB.Model(&User{}).Where("name <> ?", "Mr.T").Updates(User{Name: "Hello", Age: 0})
	GlobalDB.Model(&User{}).Where("name <> ?", "Mr.T").Updates(map[string]interface{}{"name": "World", "age": 0})
}

func DeleteRecord() {
	// 1.软删除
	users := make([]User, 0)
	GlobalDB.Where("name <> ?", "Mr.X").Delete(&users)

	// 2.硬删除
	GlobalDB.Unscoped().Where("name <> ?", "Mr.X").Delete(&users)
}

func SQL() {
	// 1.原生 SQL
	users := make([]User, 0)
	GlobalDB.Raw("select id, name, age from users where name = ?", "Mr.T").Scan(&users)
	GlobalDB.Exec("update users set age = ?", 8)
}

func Transaction() {
	// 1.嵌套事务
	GlobalDB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&User{
			Name: "张三",
			Age:  18,
		})
		tx.Create(&User{
			Name: "李四",
			Age:  18,
		})
		tx.Create(&User{
			Name: "王五",
			Age:  18,
		})
		tx.Transaction(func(tx *gorm.DB) error {
			tx.Create(&User{
				Name: "刘备",
				Age:  18,
			})
			return nil
		})
		return nil
	})

	// 2.手动事务
	tx := GlobalDB.Begin()
	tx.Create(&User{
		Name: "张三",
		Age:  18,
	})
	tx.Create(&User{
		Name: "李四",
		Age:  18,
	})
	tx.SavePoint("sp1")
	tx.Create(&User{
		Name: "王五",
		Age:  18,
	})
	tx.RollbackTo("sp1")
	tx.Commit()
}
