package generate

const appConfigTemplate = `# 应用配置
ADDR="localhost:8080"

# 日志格式
LOG_DIR=""# 默认 runtime/log
`

const cacheConfigTemplate = `# ******** redis config start ********

# 默认redis
DEFAULT_REDIS=true
DEFAULT_REDIS_ADDR="199.199.199.199:6379"
DEFAULT_REDIS_PASSWORD="123456"
DEFAULT_REDIS_DB="1"

# 认证redis
AUTH_REDIS=true
AUTH_REDIS_ADDR="199.199.199.199:6379"
AUTH_REDIS_PASSWORD="123456"
AUTH_REDIS_DB="12"

# ******** redis config end ********
`


const dbConfigTemplate = `# ******** database config start ********
# 默认数据
DEFAULT_DB=true # 连接默认数据
DEFAULT_DB_USERNAME="root3"
DEFAULT_DB_PASSWORD="root3"
DEFAULT_DB_HOST="199.199.199.199"
DEFAULT_DB_PORT="3306"
DEFAULT_DB_DATABASE="gapp"
DEFAULT_DB_CHARSET="UTF8"

# 公共数据
COMMON_DB=true # 连接公共数据库
COMMON_DB_USERNAME="root3"
COMMON_DB_PASSWORD="root3"
COMMON_DB_HOST="199.199.199.199"
COMMON_DB_PORT="3306"
COMMON_DB_DATABASE="gapp"
COMMON_DB_CHARSET="UTF8"

# 业务数据库
BUSINESS_DB=true # 连接业务库
BUSINESS_DB_USERNAME="root3"
BUSINESS_DB_PASSWORD="root3"
BUSINESS_DB_HOST="199.199.199.199"
BUSINESS_DB_PORT="3306"
BUSINESS_DB_DATABASE="gapp"
BUSINESS_DB_CHARSET="UTF8"
# ******** database config end ********
`