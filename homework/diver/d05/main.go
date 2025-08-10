package d05

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// User 用户模型
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:50;not null;uniqueIndex"`
	Email     string `gorm:"size:100;not null;uniqueIndex"`
	Password  string `gorm:"size:255;not null"`
	PostCount int    `gorm:"default:0"`
	Posts     []Post `gorm:"foreignKey:UserID"` // 一对多关系：一个用户可以发布多篇文章
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Post 文章模型
type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"size:200;not null"`
	Content       string    `gorm:"type:text"`
	UserID        uint      `gorm:"not null;index"`        // 外键
	User          User      `gorm:"foreignKey:UserID"`     // 关联用户
	Comments      []Comment `gorm:"foreignKey:PostID"`     // 一对多关系：一篇文章可以有多个评论
	CommentCount  int       `gorm:"default:0"`             // 评论数量字段
	CommentStatus string    `gorm:"size:50;default:'无评论'"` // 评论状态字段
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Comment 评论模型
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   `gorm:"not null;index"`    // 外键关联文章
	Post      Post   `gorm:"foreignKey:PostID"` // 关联文章
	UserID    uint   `gorm:"not null;index"`    // 外键关联用户
	User      User   `gorm:"foreignKey:UserID"` // 关联用户
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate 是 Post 模型的创建前钩子函数
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 设置默认的评论状态
	if p.CommentStatus == "" {
		p.CommentStatus = "无评论"
	}
	return nil
}

// AfterCreate 是 Post 模型的创建后钩子函数
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户的文章数量统计
	err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

// AfterDelete 是 Post 模型的删除后钩子函数
func (p *Post) AfterDelete(tx *gorm.DB) error {
	// 减少用户的文章数量统计
	err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count - ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

// AfterCreate 是 Comment 模型的创建后钩子函数
func (c *Comment) AfterCreate(tx *gorm.DB) error {
	// 更新文章的评论数量和状态
	err := tx.Model(&Post{}).Where("id = ?", c.PostID).Updates(map[string]interface{}{
		"comment_count":  gorm.Expr("comment_count + ?", 1),
		"comment_status": "有评论",
	}).Error

	if err != nil {
		return err
	}
	return nil
}

// AfterDelete 是 Comment 模型的删除后钩子函数
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 获取文章当前的评论数量
	var post Post
	err := tx.Select("id, comment_count").Where("id = ?", c.PostID).First(&post).Error
	if err != nil {
		return err
	}

	// 计算删除后的评论数量
	newCommentCount := post.CommentCount - 1
	if newCommentCount < 0 {
		newCommentCount = 0
	}

	// 准备更新数据
	updates := map[string]interface{}{
		"comment_count": newCommentCount,
	}

	// 如果评论数量为0，更新评论状态为"无评论"
	if newCommentCount == 0 {
		updates["comment_status"] = "无评论"
	} else {
		updates["comment_status"] = "有评论"
	}

	// 更新文章的评论数量和状态
	err = tx.Model(&Post{}).Where("id = ?", c.PostID).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

// InsertInitialData 插入初始化数据
func InsertInitialData(db *gorm.DB) error {
	// 检查是否已有用户数据
	var userCount int64
	db.Model(&User{}).Count(&userCount)

	if userCount > 0 {
		fmt.Println("数据库中已有数据，跳过初始化数据插入")
		return nil
	}

	fmt.Println("正在插入初始化数据...")

	// 创建用户
	users := []User{
		{Username: "alice", Email: "alice@example.com", Password: "password123"},
		{Username: "bob", Email: "bob@example.com", Password: "password456"},
		{Username: "charlie", Email: "charlie@example.com", Password: "password789"},
		{Username: "diana", Email: "diana@example.com", Password: "password000"},
	}

	fmt.Printf("创建 %d 个用户...\n", len(users))
	for i := range users {
		result := db.Create(&users[i])
		if result.Error != nil {
			return fmt.Errorf("创建用户失败: %v", result.Error)
		}
		fmt.Printf("  - 用户 '%s' 创建成功 (ID: %d)\n", users[i].Username, users[i].ID)
	}

	// 创建文章
	posts := []Post{
		// alice的文章
		{Title: "Go语言简介", Content: "Go是一种开源编程语言，具有简单、可靠和高效的特点。它由Google开发，于2009年首次发布。Go语言的设计目标是简洁、高效和并发安全。", UserID: users[0].ID},
		{Title: "GORM使用指南", Content: "GORM是一个优秀的Go语言ORM库，支持多种数据库。它提供了丰富的功能，包括关联查询、预加载、事务处理等。使用GORM可以大大简化数据库操作。", UserID: users[0].ID},
		{Title: "Go并发编程", Content: "Go语言通过goroutine和channel提供了强大的并发支持。goroutine是轻量级线程，channel用于goroutine之间的通信。这种设计使得并发编程变得简单而安全。", UserID: users[0].ID},

		// bob的文章
		{Title: "Web开发实践", Content: "现代Web开发需要掌握多种技术和工具。从前端的HTML/CSS/JavaScript到后端的各种框架和数据库，开发者需要具备全面的技术栈。", UserID: users[1].ID},
		{Title: "RESTful API设计", Content: "RESTful API是一种常用的Web服务设计风格。它使用HTTP方法操作资源，具有无状态、可缓存、分层系统等特点。良好的API设计可以提高系统的可维护性和可扩展性。", UserID: users[1].ID},

		// charlie的文章
		{Title: "数据库设计", Content: "良好的数据库设计是系统成功的关键因素之一。需要考虑数据的完整性、一致性、性能等因素。规范化和反规范化是数据库设计中的重要概念。", UserID: users[2].ID},
		{Title: "SQL优化技巧", Content: "SQL优化是提高数据库性能的重要手段。包括索引优化、查询重写、表结构优化等方面。合理的SQL语句可以显著提高查询效率。", UserID: users[2].ID},

		// diana的文章
		{Title: "微服务架构", Content: "微服务架构是一种将单一应用程序开发为一套小型服务的方法。每个服务运行在自己的进程中，并通过轻量级机制（通常是HTTP资源API）进行通信。", UserID: users[3].ID},
		{Title: "Docker容器化", Content: "Docker是一个开源的应用容器引擎，让开发者可以打包他们的应用以及依赖包到一个可移植的容器中，然后发布到任何流行的Linux机器上，也可以实现虚拟化。", UserID: users[3].ID},
	}

	fmt.Printf("创建 %d 篇文章...\n", len(posts))
	for i := range posts {
		result := db.Create(&posts[i])
		if result.Error != nil {
			return fmt.Errorf("创建文章失败: %v", result.Error)
		}
		fmt.Printf("  - 文章 '%s' 创建成功 (ID: %d)\n", posts[i].Title, posts[i].ID)
	}

	// 创建评论
	comments := []Comment{
		// 对 alice 的 "Go语言简介" 的评论
		{Content: "很好的介绍，谢谢分享！", PostID: 1, UserID: users[1].ID},
		{Content: "对初学者很有帮助。", PostID: 1, UserID: users[2].ID},
		{Content: "希望能有更多实际例子。", PostID: 1, UserID: users[3].ID},

		// 对 alice 的 "GORM使用指南" 的评论
		{Content: "详细介绍了GORM的使用方法。", PostID: 2, UserID: users[1].ID},
		{Content: "期待更多相关内容。", PostID: 2, UserID: users[2].ID},
		{Content: "这个ORM库确实很好用。", PostID: 2, UserID: users[3].ID},

		// 对 alice 的 "Go并发编程" 的评论
		{Content: "goroutine的讲解很清晰。", PostID: 3, UserID: users[1].ID},
		{Content: "channel部分需要更多例子。", PostID: 3, UserID: users[3].ID},

		// 对 bob 的 "Web开发实践" 的评论
		{Content: "实用的开发经验分享。", PostID: 4, UserID: users[0].ID},
		{Content: "内容很全面。", PostID: 4, UserID: users[2].ID},

		// 对 bob 的 "RESTful API设计" 的评论
		{Content: "API设计原则讲得很好。", PostID: 5, UserID: users[0].ID},
		{Content: "可以增加一些安全方面的内容。", PostID: 5, UserID: users[2].ID},
		{Content: "很实用的API设计指南。", PostID: 5, UserID: users[3].ID},

		// 对 charlie 的 "数据库设计" 的评论
		{Content: "数据库设计很重要。", PostID: 6, UserID: users[0].ID},
		{Content: "规范化理论讲得很清楚。", PostID: 6, UserID: users[1].ID},

		// 对 charlie 的 "SQL优化技巧" 的评论
		{Content: "索引优化部分很有用。", PostID: 7, UserID: users[0].ID},
		{Content: "希望能有更多实际案例。", PostID: 7, UserID: users[1].ID},
		{Content: "性能优化的关键技巧。", PostID: 7, UserID: users[3].ID},

		// 对 diana 的 "微服务架构" 的评论
		{Content: "微服务是趋势。", PostID: 8, UserID: users[0].ID},
		{Content: "服务拆分需要注意边界。", PostID: 8, UserID: users[1].ID},
		{Content: "治理复杂度是挑战。", PostID: 8, UserID: users[2].ID},

		// 对 diana 的 "Docker容器化" 的评论
		{Content: "容器化是部署的未来。", PostID: 9, UserID: users[0].ID},
		{Content: "Docker确实方便。", PostID: 9, UserID: users[1].ID},
		{Content: "生产环境使用需要注意安全。", PostID: 9, UserID: users[2].ID},
	}

	fmt.Printf("创建 %d 条评论...\n", len(comments))
	for i := range comments {
		result := db.Create(&comments[i])
		if result.Error != nil {
			return fmt.Errorf("创建评论失败: %v", result.Error)
		}
		fmt.Printf("  - 评论 '%.30s...' 创建成功 (ID: %d)\n", comments[i].Content, comments[i].ID)
	}

	fmt.Println("初始化数据插入成功!")
	fmt.Printf("总计: %d 个用户, %d 篇文章, %d 条评论\n", len(users), len(posts), len(comments))

	return nil
}

// GetUserPostsAndComments 查询指定用户发布的所有文章及其评论信息
func GetUserPostsAndComments(db *gorm.DB, userID uint) (*User, error) {
	var user User

	// 使用Preload预加载用户的文章和文章的评论
	result := db.Preload("Posts").Preload("Posts.Comments").Preload("Posts.Comments.User").First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// GetPostWithMostComments 查询评论数量最多的文章信息
func GetPostWithMostComments(db *gorm.DB) (*Post, error) {
	var post Post

	// 使用子查询和JOIN来查找评论数量最多的文章
	subQuery := db.Model(&Comment{}).
		Select("post_id, COUNT(*) as comment_count").
		Group("post_id").
		Order("comment_count DESC").
		Limit(1)

	// 获取评论数最多的post_id
	var result struct {
		PostID       uint
		CommentCount int
	}

	err := subQuery.Scan(&result).Error
	if err != nil {
		return nil, err
	}

	if result.PostID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// 查询该文章的完整信息
	err = db.Preload("User").Preload("Comments").Preload("Comments.User").First(&post, result.PostID).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetPostWithMostCommentsDetailed 查询评论数量最多的文章详细信息
func GetPostWithMostCommentsDetailed(db *gorm.DB) (map[string]interface{}, error) {
	var post Post
	var commentCount int64

	// 先找到评论数最多的文章ID
	var result struct {
		PostID uint
	}

	err := db.Model(&Comment{}).
		Select("post_id, COUNT(*) as count").
		Group("post_id").
		Order("count DESC").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	if result.PostID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// 获取该文章的评论数量
	err = db.Model(&Comment{}).Where("post_id = ?", result.PostID).Count(&commentCount).Error
	if err != nil {
		return nil, err
	}

	// 获取文章详细信息
	err = db.Preload("User").Preload("Comments").Preload("Comments.User").First(&post, result.PostID).Error
	if err != nil {
		return nil, err
	}

	// 构造返回结果
	resultData := map[string]interface{}{
		"post":          post,
		"comment_count": commentCount,
	}

	return resultData, nil
}

// GetPostWithMostCommentsV2 使用原生SQL查询评论数量最多的文章信息
func GetPostWithMostCommentsV2(db *gorm.DB) (*Post, error) {
	var post Post

	// 使用原生SQL查询评论数量最多的文章
	err := db.Raw(`
        SELECT p.* 
        FROM posts p
        WHERE p.id = (
            SELECT c.post_id 
            FROM comments c 
            GROUP BY c.post_id 
            ORDER BY COUNT(*) DESC 
            LIMIT 1
        )
    `).Scan(&post).Error

	if err != nil {
		return nil, err
	}

	// 检查是否找到了文章
	if post.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// 预加载关联信息
	err = db.Preload("User").Preload("Comments").Preload("Comments.User").First(&post, post.ID).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetPostWithMostCommentsV3 使用GORM的Joins和Group查询评论数量最多的文章
func GetPostWithMostCommentsV3(db *gorm.DB) (*Post, error) {
	var post Post

	// 使用Joins和Group查询
	err := db.Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("left join comments on posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		First(&post).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

// PrintPostWithMostComments 打印评论数最多的文章信息
func PrintPostWithMostComments(resultData map[string]interface{}) {
	post, ok := resultData["post"].(Post)
	if !ok {
		fmt.Println("无法获取文章信息")
		return
	}

	commentCount, ok := resultData["comment_count"].(int64)
	if !ok {
		fmt.Println("无法获取评论数量")
		return
	}

	fmt.Printf("\n=== 评论数量最多的文章 ===\n")
	fmt.Printf("文章标题: %s\n", post.Title)
	fmt.Printf("作者: %s\n", post.User.Username)
	fmt.Printf("评论数量: %d\n", commentCount)
	fmt.Printf("文章内容: %.100s...\n", post.Content)
	fmt.Printf("发布时间: %s\n", post.CreatedAt.Format("2006-01-02 15:04:05"))

	fmt.Printf("\n评论列表:\n")
	for i, comment := range post.Comments {
		fmt.Printf("  [%d] %s (评论者: %s, 时间: %s)\n",
			i+1, comment.Content, comment.User.Username,
			comment.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	fmt.Println("表结构创建成功:")
	fmt.Println("- users 表")
	fmt.Println("- posts 表")
	fmt.Println("- comments 表")
	fmt.Println("博客系统数据库初始化完成！")
	err := InsertInitialData(db)
	if err != nil {
		return
	}

	fmt.Println("开始查询数据...")
	user, err := GetUserPostsAndComments(db, 1)
	if err != nil {
		fmt.Println("查询数据失败:", err)
		return
	}
	fmt.Println("用户:", user.Username)
	fmt.Println("文章:")
	for _, post := range user.Posts {
		fmt.Println("  -", post.Title)
		fmt.Println("    - 评论:")
		for _, comment := range post.Comments {
			fmt.Println("      -", comment.Content)
			fmt.Println("        - 评论者:", comment.User.Username)
		}
	}
	fmt.Println("数据查询完成！")

	// 查询评论数量最多的文章
	fmt.Println("\n=== 查询评论数量最多的文章 ===")
	postWithMostComments, err := GetPostWithMostComments(db)
	if err != nil {
		fmt.Println("查询评论最多的文章失败:", err)
		return
	}

	fmt.Printf("评论数量最多的文章: %s\n", postWithMostComments.Title)
	fmt.Printf("作者: %s\n", postWithMostComments.User.Username)

	// 使用详细查询方法
	fmt.Println("\n=== 使用详细查询方法 ===")
	detailedResult, err := GetPostWithMostCommentsDetailed(db)
	if err != nil {
		fmt.Println("详细查询失败:", err)
		return
	}

	PrintPostWithMostComments(detailedResult)

	fmt.Println("数据查询完成！")

}
