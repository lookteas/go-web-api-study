// 数据库基础 - Go语言数据库操作
// 学习目标：掌握Go语言中的数据库连接、操作和最佳实践

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	
	// 数据库驱动（注释掉避免编译错误，实际使用时需要导入）
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

// User 用户结构体
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`  // 不在JSON中显示
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Product 产品结构体
type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Stock       int     `json:"stock" db:"stock"`
	CategoryID  int     `json:"category_id" db:"category_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Category 分类结构体
type Category struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
	SSLMode  string
}

// UserRepository 用户数据访问层
type UserRepository struct {
	db *sql.DB
}

// ProductRepository 产品数据访问层
type ProductRepository struct {
	db *sql.DB
}

func main() {
	fmt.Println("=== Go数据库操作学习 ===")
	
	// 1. 数据库连接
	fmt.Println("\n1. 数据库连接")
	demonstrateDatabaseConnection()
	
	// 2. SQL基础操作
	fmt.Println("\n2. SQL基础操作")
	demonstrateBasicSQL()
	
	// 3. 预处理语句
	fmt.Println("\n3. 预处理语句")
	demonstratePreparedStatements()
	
	// 4. 事务处理
	fmt.Println("\n4. 事务处理")
	demonstrateTransactions()
	
	// 5. 连接池管理
	fmt.Println("\n5. 连接池管理")
	demonstrateConnectionPool()
	
	// 6. ORM基础
	fmt.Println("\n6. ORM基础")
	demonstrateORM()
	
	// 7. 数据库迁移
	fmt.Println("\n7. 数据库迁移")
	demonstrateMigrations()
	
	// 8. 性能优化
	fmt.Println("\n8. 性能优化")
	demonstratePerformanceOptimization()
	
	fmt.Println("\n=== 数据库操作学习完成 ===")
}

// 1. 数据库连接演示
func demonstrateDatabaseConnection() {
	fmt.Println("数据库连接方法：")
	
	// 不同数据库的连接字符串
	connectionStrings := map[string]string{
		"MySQL":      "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		"PostgreSQL": "host=localhost port=5432 user=username dbname=mydb password=password sslmode=disable",
		"SQLite":     "./database.db",
		"SQL Server": "server=localhost;user id=username;password=password;database=dbname",
	}
	
	for db, connStr := range connectionStrings {
		fmt.Printf("%-12s: %s\n", db, connStr)
	}
	
	// 连接示例代码
	fmt.Println("\n连接代码示例：")
	connectionExample := `
// MySQL连接示例
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// 测试连接
err = db.Ping()
if err != nil {
    log.Fatal("无法连接到数据库:", err)
}
`
	fmt.Println(connectionExample)
	
	// 连接配置
	fmt.Println("连接配置参数：")
	configParams := map[string]string{
		"MaxOpenConns":     "最大打开连接数",
		"MaxIdleConns":     "最大空闲连接数",
		"ConnMaxLifetime":  "连接最大生存时间",
		"ConnMaxIdleTime":  "连接最大空闲时间",
	}
	
	for param, description := range configParams {
		fmt.Printf("%-16s: %s\n", param, description)
	}
}

// 2. SQL基础操作演示
func demonstrateBasicSQL() {
	fmt.Println("SQL基础操作：")
	
	// CRUD操作示例
	sqlOperations := map[string]string{
		"CREATE": `
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);`,
		"INSERT": `
INSERT INTO users (username, email, password) 
VALUES (?, ?, ?);`,
		"SELECT": `
SELECT id, username, email, created_at 
FROM users 
WHERE id = ?;`,
		"UPDATE": `
UPDATE users 
SET username = ?, email = ?, updated_at = NOW() 
WHERE id = ?;`,
		"DELETE": `
DELETE FROM users 
WHERE id = ?;`,
	}
	
	for operation, sql := range sqlOperations {
		fmt.Printf("\n%s操作：%s\n", operation, sql)
	}
	
	// Go代码示例
	fmt.Println("\nGo代码实现：")
	goCodeExample := `
// 插入数据
result, err := db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", 
    "john_doe", "john@example.com", "hashed_password")
if err != nil {
    log.Fatal(err)
}

// 获取插入的ID
lastID, err := result.LastInsertId()
if err != nil {
    log.Fatal(err)
}

// 查询数据
var user User
err = db.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = ?", lastID).
    Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
if err != nil {
    if err == sql.ErrNoRows {
        fmt.Println("用户不存在")
    } else {
        log.Fatal(err)
    }
}
`
	fmt.Println(goCodeExample)
}

// 3. 预处理语句演示
func demonstratePreparedStatements() {
	fmt.Println("预处理语句的优势：")
	
	advantages := []string{
		"1. 防止SQL注入攻击",
		"2. 提高执行性能（重复使用）",
		"3. 减少SQL解析开销",
		"4. 支持参数化查询",
	}
	
	for _, advantage := range advantages {
		fmt.Println(advantage)
	}
	
	// 预处理语句示例
	fmt.Println("\n预处理语句示例：")
	preparedExample := `
// 准备语句
stmt, err := db.Prepare("SELECT id, username, email FROM users WHERE username = ? AND email = ?")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// 执行查询
rows, err := stmt.Query("john_doe", "john@example.com")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

// 处理结果
for rows.Next() {
    var user User
    err := rows.Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("用户: %+v\n", user)
}
`
	fmt.Println(preparedExample)
	
	// 批量操作
	fmt.Println("\n批量操作示例：")
	batchExample := `
// 批量插入
stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

users := [][]interface{}{
    {"user1", "user1@example.com", "password1"},
    {"user2", "user2@example.com", "password2"},
    {"user3", "user3@example.com", "password3"},
}

for _, user := range users {
    _, err := stmt.Exec(user...)
    if err != nil {
        log.Printf("插入用户失败: %v", err)
    }
}
`
	fmt.Println(batchExample)
}

// 4. 事务处理演示
func demonstrateTransactions() {
	fmt.Println("事务处理的重要性：")
	
	transactionBenefits := []string{
		"1. 原子性 (Atomicity) - 全部成功或全部失败",
		"2. 一致性 (Consistency) - 数据保持一致状态",
		"3. 隔离性 (Isolation) - 并发事务相互隔离",
		"4. 持久性 (Durability) - 提交后数据永久保存",
	}
	
	for _, benefit := range transactionBenefits {
		fmt.Println(benefit)
	}
	
	// 事务示例
	fmt.Println("\n事务使用示例：")
	transactionExample := `
// 开始事务
tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}

// 使用defer确保事务处理
defer func() {
    if p := recover(); p != nil {
        tx.Rollback()
        panic(p) // 重新抛出panic
    } else if err != nil {
        tx.Rollback()
    } else {
        err = tx.Commit()
    }
}()

// 执行多个操作
_, err = tx.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", 
    "new_user", "new@example.com", "password")
if err != nil {
    return err
}

_, err = tx.Exec("UPDATE products SET stock = stock - 1 WHERE id = ?", productID)
if err != nil {
    return err
}

_, err = tx.Exec("INSERT INTO orders (user_id, product_id, quantity) VALUES (?, ?, ?)", 
    userID, productID, 1)
if err != nil {
    return err
}

// 如果到这里没有错误，事务将在defer中提交
`
	fmt.Println(transactionExample)
	
	// 事务隔离级别
	fmt.Println("\n事务隔离级别：")
	isolationLevels := map[string]string{
		"READ UNCOMMITTED": "读未提交 - 最低级别，可能出现脏读",
		"READ COMMITTED":   "读已提交 - 避免脏读，可能出现不可重复读",
		"REPEATABLE READ":  "可重复读 - 避免脏读和不可重复读",
		"SERIALIZABLE":     "串行化 - 最高级别，完全隔离",
	}
	
	for level, description := range isolationLevels {
		fmt.Printf("%-17s: %s\n", level, description)
	}
}

// 5. 连接池管理演示
func demonstrateConnectionPool() {
	fmt.Println("连接池配置：")
	
	// 连接池参数
	poolParams := map[string]string{
		"MaxOpenConns":     "最大打开连接数 (默认无限制)",
		"MaxIdleConns":     "最大空闲连接数 (默认2)",
		"ConnMaxLifetime":  "连接最大生存时间 (默认无限制)",
		"ConnMaxIdleTime":  "连接最大空闲时间 (Go 1.15+)",
	}
	
	for param, description := range poolParams {
		fmt.Printf("%-16s: %s\n", param, description)
	}
	
	// 连接池配置示例
	fmt.Println("\n连接池配置示例：")
	poolConfigExample := `
// 配置连接池
db.SetMaxOpenConns(25)                 // 最大打开连接数
db.SetMaxIdleConns(25)                 // 最大空闲连接数
db.SetConnMaxLifetime(5 * time.Minute) // 连接最大生存时间
db.SetConnMaxIdleTime(5 * time.Minute) // 连接最大空闲时间

// 获取连接池状态
stats := db.Stats()
fmt.Printf("打开连接数: %d\n", stats.OpenConnections)
fmt.Printf("使用中连接数: %d\n", stats.InUse)
fmt.Printf("空闲连接数: %d\n", stats.Idle)
fmt.Printf("等待连接数: %d\n", stats.WaitCount)
fmt.Printf("等待时长: %v\n", stats.WaitDuration)
`
	fmt.Println(poolConfigExample)
	
	// 最佳实践
	fmt.Println("\n连接池最佳实践：")
	bestPractices := []string{
		"1. 根据应用负载设置合适的连接数",
		"2. 监控连接池状态和性能指标",
		"3. 设置合理的连接生存时间",
		"4. 避免长时间持有连接",
		"5. 使用连接池而不是频繁创建连接",
	}
	
	for _, practice := range bestPractices {
		fmt.Println(practice)
	}
}

// 6. ORM基础演示
func demonstrateORM() {
	fmt.Println("ORM (对象关系映射) 概念：")
	
	// ORM优势
	ormAdvantages := []string{
		"1. 减少SQL编写工作量",
		"2. 提供类型安全的查询",
		"3. 自动处理数据映射",
		"4. 支持数据库迁移",
		"5. 提供关联关系处理",
	}
	
	for _, advantage := range ormAdvantages {
		fmt.Println(advantage)
	}
	
	// 流行的Go ORM
	fmt.Println("\n流行的Go ORM框架：")
	ormFrameworks := map[string]string{
		"GORM":     "功能丰富，易于使用，支持多种数据库",
		"Ent":      "Facebook开发，类型安全，代码生成",
		"SQLBoiler": "代码生成，高性能，类型安全",
		"Xorm":     "简单易用，支持多种数据库",
		"Beego ORM": "Beego框架内置ORM",
	}
	
	for framework, description := range ormFrameworks {
		fmt.Printf("%-10s: %s\n", framework, description)
	}
	
	// GORM示例
	fmt.Println("\nGORM使用示例：")
	gormExample := `
// 定义模型
type User struct {
    ID        uint      ` + "`gorm:\"primaryKey\"`" + `
    Username  string    ` + "`gorm:\"uniqueIndex;size:50\"`" + `
    Email     string    ` + "`gorm:\"uniqueIndex;size:100\"`" + `
    Password  string    ` + "`gorm:\"size:255\"`" + `
    CreatedAt time.Time
    UpdatedAt time.Time
}

// 自动迁移
db.AutoMigrate(&User{})

// 创建记录
user := User{Username: "john", Email: "john@example.com", Password: "hashed"}
result := db.Create(&user)

// 查询记录
var user User
db.First(&user, 1) // 根据主键查询
db.First(&user, "username = ?", "john") // 根据条件查询

// 更新记录
db.Model(&user).Update("Email", "newemail@example.com")
db.Model(&user).Updates(User{Email: "new@example.com", Username: "john_new"})

// 删除记录
db.Delete(&user, 1)
`
	fmt.Println(gormExample)
}

// 7. 数据库迁移演示
func demonstrateMigrations() {
	fmt.Println("数据库迁移的重要性：")
	
	migrationBenefits := []string{
		"1. 版本控制数据库结构变更",
		"2. 团队协作时保持数据库同步",
		"3. 支持回滚到之前版本",
		"4. 自动化部署流程",
		"5. 记录数据库变更历史",
	}
	
	for _, benefit := range migrationBenefits {
		fmt.Println(benefit)
	}
	
	// 迁移工具
	fmt.Println("\n迁移工具：")
	migrationTools := map[string]string{
		"golang-migrate": "专业的数据库迁移工具",
		"GORM AutoMigrate": "GORM内置的自动迁移",
		"Goose": "简单的数据库迁移工具",
		"Atlas": "现代化的数据库迁移工具",
	}
	
	for tool, description := range migrationTools {
		fmt.Printf("%-15s: %s\n", tool, description)
	}
	
	// 迁移文件示例
	fmt.Println("\n迁移文件示例：")
	migrationExample := `
-- 001_create_users_table.up.sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 001_create_users_table.down.sql
DROP TABLE users;

-- 002_add_user_profile.up.sql
ALTER TABLE users ADD COLUMN profile TEXT;
ALTER TABLE users ADD COLUMN avatar_url VARCHAR(255);

-- 002_add_user_profile.down.sql
ALTER TABLE users DROP COLUMN profile;
ALTER TABLE users DROP COLUMN avatar_url;
`
	fmt.Println(migrationExample)
	
	// 迁移命令
	fmt.Println("\n迁移命令示例：")
	migrationCommands := []string{
		"migrate -path ./migrations -database mysql://user:pass@tcp(localhost:3306)/db up",
		"migrate -path ./migrations -database mysql://user:pass@tcp(localhost:3306)/db down 1",
		"migrate -path ./migrations -database mysql://user:pass@tcp(localhost:3306)/db version",
	}
	
	for _, command := range migrationCommands {
		fmt.Printf("  %s\n", command)
	}
}

// 8. 性能优化演示
func demonstratePerformanceOptimization() {
	fmt.Println("数据库性能优化策略：")
	
	// 查询优化
	fmt.Println("\n1. 查询优化：")
	queryOptimizations := []string{
		"使用索引加速查询",
		"避免SELECT *，只查询需要的字段",
		"使用LIMIT限制结果集大小",
		"优化WHERE条件，使用合适的操作符",
		"避免在WHERE子句中使用函数",
	}
	
	for _, optimization := range queryOptimizations {
		fmt.Printf("  - %s\n", optimization)
	}
	
	// 索引策略
	fmt.Println("\n2. 索引策略：")
	indexStrategies := []string{
		"为经常查询的字段创建索引",
		"复合索引的字段顺序很重要",
		"避免过多索引影响写入性能",
		"定期分析和优化索引使用",
		"使用覆盖索引减少回表查询",
	}
	
	for _, strategy := range indexStrategies {
		fmt.Printf("  - %s\n", strategy)
	}
	
	// 连接优化
	fmt.Println("\n3. 连接优化：")
	connectionOptimizations := []string{
		"合理配置连接池大小",
		"使用连接复用减少开销",
		"监控连接池使用情况",
		"设置合适的连接超时时间",
		"避免长时间持有连接",
	}
	
	for _, optimization := range connectionOptimizations {
		fmt.Printf("  - %s\n", optimization)
	}
	
	// 批量操作
	fmt.Println("\n4. 批量操作：")
	batchOptimizations := []string{
		"使用批量插入代替单条插入",
		"使用事务包装批量操作",
		"合理设置批次大小",
		"使用预处理语句提高性能",
		"考虑使用BULK操作",
	}
	
	for _, optimization := range batchOptimizations {
		fmt.Printf("  - %s\n", optimization)
	}
	
	// 性能监控
	fmt.Println("\n5. 性能监控：")
	monitoringAspects := []string{
		"监控慢查询日志",
		"分析查询执行计划",
		"监控数据库连接数",
		"跟踪查询响应时间",
		"监控数据库资源使用",
	}
	
	for _, aspect := range monitoringAspects {
		fmt.Printf("  - %s\n", aspect)
	}
}

// UserRepository 方法实现

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	query := `
		INSERT INTO users (username, email, password, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?)`
	
	now := time.Now()
	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, now, now)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	user.ID = int(id)
	user.CreatedAt = now
	user.UpdatedAt = now
	
	return nil
}

func (r *UserRepository) GetByID(id int) (*User, error) {
	query := `
		SELECT id, username, email, password, created_at, updated_at 
		FROM users WHERE id = ?`
	
	user := &User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	
	return user, nil
}

func (r *UserRepository) GetByUsername(username string) (*User, error) {
	query := `
		SELECT id, username, email, password, created_at, updated_at 
		FROM users WHERE username = ?`
	
	user := &User{}
	err := r.db.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	
	return user, nil
}

func (r *UserRepository) Update(user *User) error {
	query := `
		UPDATE users 
		SET username = ?, email = ?, password = ?, updated_at = ? 
		WHERE id = ?`
	
	user.UpdatedAt = time.Now()
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.UpdatedAt, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) List(limit, offset int) ([]User, error) {
	query := `
		SELECT id, username, email, password, created_at, updated_at 
		FROM users 
		ORDER BY created_at DESC 
		LIMIT ? OFFSET ?`
	
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID, &user.Username, &user.Email, &user.Password,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, rows.Err()
}

// ProductRepository 方法实现

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *Product) error {
	query := `
		INSERT INTO products (name, description, price, stock, category_id, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`
	
	now := time.Now()
	result, err := r.db.Exec(query, product.Name, product.Description, product.Price, 
		product.Stock, product.CategoryID, now, now)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	product.ID = int(id)
	product.CreatedAt = now
	product.UpdatedAt = now
	
	return nil
}

func (r *ProductRepository) GetByID(id int) (*Product, error) {
	query := `
		SELECT id, name, description, price, stock, category_id, created_at, updated_at 
		FROM products WHERE id = ?`
	
	product := &Product{}
	err := r.db.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price,
		&product.Stock, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	
	return product, nil
}

// 工具函数

// CreateDatabaseConnection 创建数据库连接
func CreateDatabaseConnection(config DatabaseConfig) (*sql.DB, error) {
	var dsn string
	
	switch config.Driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Username, config.Password, config.Host, config.Port, config.Database)
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.Password, config.Database, config.SSLMode)
	case "sqlite3":
		dsn = config.Database
	default:
		return nil, fmt.Errorf("不支持的数据库驱动: %s", config.Driver)
	}
	
	db, err := sql.Open(config.Driver, dsn)
	if err != nil {
		return nil, err
	}
	
	// 测试连接
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	
	return db, nil
}

// ConfigureConnectionPool 配置连接池
func ConfigureConnectionPool(db *sql.DB) {
	// 设置最大打开连接数
	db.SetMaxOpenConns(25)
	
	// 设置最大空闲连接数
	db.SetMaxIdleConns(25)
	
	// 设置连接最大生存时间
	db.SetConnMaxLifetime(5 * time.Minute)
	
	// 设置连接最大空闲时间 (Go 1.15+)
	db.SetConnMaxIdleTime(5 * time.Minute)
}

// ExecuteInTransaction 在事务中执行操作
func ExecuteInTransaction(db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	
	err = fn(tx)
	return err
}

/*
学习要点总结：

1. 数据库连接
   - 不同数据库的连接方式
   - 连接字符串格式
   - 连接测试和错误处理

2. SQL基础操作
   - CRUD操作的SQL语句
   - Go语言中的SQL执行
   - 结果处理和错误处理

3. 预处理语句
   - 防止SQL注入
   - 提高执行性能
   - 批量操作优化

4. 事务处理
   - ACID特性理解
   - 事务的使用场景
   - 事务隔离级别

5. 连接池管理
   - 连接池参数配置
   - 性能监控
   - 最佳实践

6. ORM基础
   - ORM框架选择
   - 模型定义
   - 基本操作

7. 数据库迁移
   - 版本控制
   - 迁移工具使用
   - 迁移文件编写

8. 性能优化
   - 查询优化
   - 索引策略
   - 批量操作
   - 性能监控

实践建议：
- 从简单的CRUD操作开始
- 学习使用预处理语句
- 掌握事务的正确使用
- 了解连接池配置
- 尝试使用ORM框架
- 学习数据库设计原则
- 关注性能优化技巧
*/