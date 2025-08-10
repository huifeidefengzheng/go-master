package d04

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Book 书籍信息结构体，与books表字段对应
type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func QueryExpensiveBooks(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book

	query := "SELECT id, title, author, price FROM books WHERE price > ? ORDER BY price DESC"
	err := db.Select(&books, query, minPrice)

	if err != nil {
		return nil, fmt.Errorf("查询书籍信息失败: %v", err)
	}

	return books, nil
}

// CreateBooksTable 创建books表
func CreateBooksTable(db *sqlx.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS books (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author VARCHAR(255) NOT NULL,
        price DECIMAL(10,2) NOT NULL DEFAULT 0.00
    )`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("创建books表失败: %v", err)
	}

	fmt.Println("books表创建成功")
	return nil
}

// InsertSampleData 插入示例数据
func InsertSampleData(db *sqlx.DB) error {
	// 检查是否已有数据
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM books")
	if err != nil {
		return fmt.Errorf("查询数据条数失败: %v", err)
	}

	// 如果已有数据则不插入
	if count > 0 {
		fmt.Println("books表中已有数据，跳过示例数据插入")
		return nil
	}

	// 插入示例数据
	books := []Book{
		{Title: "红楼梦", Author: "曹雪芹", Price: 89.50},
		{Title: "西游记", Author: "吴承恩", Price: 76.20},
		{Title: "三国演义", Author: "罗贯中", Price: 92.00},
		{Title: "水浒传", Author: "施耐庵", Price: 68.80},
		{Title: "朝花夕拾", Author: "鲁迅", Price: 35.60},
		{Title: "呐喊", Author: "鲁迅", Price: 42.30},
		{Title: "彷徨", Author: "鲁迅", Price: 38.90},
		{Title: "围城", Author: "钱钟书", Price: 56.70},
		{Title: "活着", Author: "余华", Price: 32.50},
		{Title: "平凡的世界", Author: "路遥", Price: 128.00},
		{Title: "白鹿原", Author: "陈忠实", Price: 65.40},
		{Title: "穆斯林的葬礼", Author: "霍达", Price: 48.20},
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败: %v", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO books (title, author, price) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("准备插入语句失败: %v", err)
	}
	defer stmt.Close()

	for _, book := range books {
		_, err := stmt.Exec(book.Title, book.Author, book.Price)
		if err != nil {
			return fmt.Errorf("插入书籍数据失败: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("提交事务失败: %v", err)
	}

	fmt.Printf("成功插入 %d 条示例数据\n", len(books))
	return nil
}

// QueryBooksByAuthor 查询指定作者的书籍
func QueryBooksByAuthor(db *sqlx.DB, author string) ([]Book, error) {
	var books []Book

	query := "SELECT id, title, author, price FROM books WHERE author = ? ORDER BY price DESC"
	err := db.Select(&books, query, author)

	if err != nil {
		return nil, fmt.Errorf("根据作者查询书籍失败: %v", err)
	}

	return books, nil
}

// QueryBooksByPriceRange 查询指定价格范围内的书籍
func QueryBooksByPriceRange(db *sqlx.DB, minPrice, maxPrice float64) ([]Book, error) {
	var books []Book

	query := "SELECT id, title, author, price FROM books WHERE price BETWEEN ? AND ? ORDER BY price"
	err := db.Select(&books, query, minPrice, maxPrice)

	if err != nil {
		return nil, fmt.Errorf("查询价格范围内的书籍失败: %v", err)
	}

	return books, nil
}

// QueryBookByID 根据ID查询单本书籍
func QueryBookByID(db *sqlx.DB, id int) (*Book, error) {
	var book Book

	query := "SELECT id, title, author, price FROM books WHERE id = ?"
	err := db.Get(&book, query, id)

	if err != nil {
		return nil, fmt.Errorf("查询书籍失败: %v", err)
	}

	return &book, nil
}

func Run(db *sqlx.DB) {
	err := CreateBooksTable(db)
	if err != nil {
		fmt.Print("创建表失败：", err)
		return
	}
	err1 := InsertSampleData(db)
	if err1 != nil {
		return
	}

	minPrice := 50.0
	books, err := QueryExpensiveBooks(db, minPrice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("价格高于50元的书籍信息:")
	for _, book := range books {
		fmt.Printf("书籍ID: %d, 标题: %s, 作者: %s, 价格: %.2f\n", book.Id, book.Title, book.Author, book.Price)
	}
}
