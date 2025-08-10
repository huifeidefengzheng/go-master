package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePost 创建文章
func CreatePost(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id").(uint)

	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.UserID = userID

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章创建失败"})
		return
	}

	// 获取完整的文章信息（包括用户信息）
	//var createdPost Post
	//db.Preload("User").First(&createdPost, post.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "文章创建成功",
		"post":    post,
	})
}

// GetPosts 获取所有文章列表
func GetPosts(c *gin.Context, db *gorm.DB) {
	var posts []Post
	result := db.Preload("User").Preload("Comments").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"count": len(posts),
	})
}

// GetPost 获取单个文章详情
func GetPost(c *gin.Context, db *gorm.DB) {
	postID := c.Param("id")
	var post Post

	result := db.Preload("User").Preload("Comments").Preload("Comments.User").First(&post, postID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// UpdatePost 更新文章
func UpdatePost(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id").(uint)
	postID := c.Param("id")

	// 检查文章是否存在
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询文章失败"})
		return
	}

	// 检查文章是否属于当前用户
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限修改此文章"})
		return
	}

	// 更新文章
	var updateData Post
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&post).Updates(map[string]interface{}{
		"title":   updateData.Title,
		"content": updateData.Content,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章更新失败"})
		return
	}

	// 获取更新后的文章信息
	var updatedPost Post
	db.Preload("User").First(&updatedPost, post.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "文章更新成功",
		"post":    updatedPost,
	})
}

// DeletePost 删除文章
func DeletePost(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id").(uint)
	postID := c.Param("id")

	// 检查文章是否存在
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询文章失败"})
		return
	}

	// 检查文章是否属于当前用户
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限删除此文章"})
		return
	}

	// 删除文章
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}

// CreateComment 创建评论
func CreateComment(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id").(uint)
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 检查文章是否存在
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询文章失败"})
		return
	}

	var comment Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserID = userID
	comment.PostID = uint(postID)

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论创建失败"})
		return
	}

	// 获取完整的评论信息（包括用户信息）
	var createdComment Comment
	db.Preload("User").First(&createdComment, comment.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "评论创建成功",
		"comment": createdComment,
	})
}

// GetComments 获取某篇文章的所有评论
func GetComments(c *gin.Context, db *gorm.DB) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 检查文章是否存在
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询文章失败"})
		return
	}

	var comments []Comment
	result := db.Where("post_id = ?", postID).Preload("User").Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
		"count":    len(comments),
	})
}
