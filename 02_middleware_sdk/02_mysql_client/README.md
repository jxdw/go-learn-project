# go-sql-driver说明

## 参考资料
https://www.ithome.io/b/a652bcbf-45fd-e01a-f52b-d2e3179a6d92.html
https://www.cnblogs.com/tsiangleo/p/4483657.html
## mysql client的概念模型
### driver（驱动）
### connection（链接）
### statement（声明）
### resultSet（结果集）

## mysql client源码分析
### init函数涉及文件
 github.com\go-sql-driver\mysql@v1.4.1\driver.go
 `````
 func init() {
 	sql.Register("mysql", &MySQLDriver{})
 }
`````

### sql.Open函数
`````
func Open(driverName, dataSourceName string) (*DB, error) {
	driversMu.RLock()
	driveri, ok := drivers[driverName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("sql: unknown driver %q (forgotten import?)", driverName)
	}

	if driverCtx, ok := driveri.(driver.DriverContext); ok {
		connector, err := driverCtx.OpenConnector(dataSourceName)
		if err != nil {
			return nil, err
		}
		return OpenDB(connector), nil
	}

	return OpenDB(dsnConnector{dsn: dataSourceName, driver: driveri}), nil
}
`````
这里是声明了一个DB对象。